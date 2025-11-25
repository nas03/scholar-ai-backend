package helper

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/resend/resend-go/v2"
)

type IMailHelper interface {
	SendMail(ctx context.Context, to, subject, body string) (string, error)
	ReplaceParameters(ctx context.Context, html string, data any) string
}

type MailHelper struct {
	client *resend.Client
}

func NewMailHelper() IMailHelper {
	return &MailHelper{
		client: global.Mail,
	}
}

func (h *MailHelper) SendMail(ctx context.Context, to, subject, html string) (string, error) {
	params := &resend.SendEmailRequest{
		From:    global.Config.Resend.From,
		To:      []string{to},
		Subject: subject,
		Html:    html,
	}

	sent, err := h.client.Emails.SendWithContext(ctx, params)

	if err != nil {
		return "", fmt.Errorf("failed to send email to '%s': %w", to, err)
	}

	return sent.Id, nil
}

func (h *MailHelper) ReplaceParameters(ctx context.Context, html string, data any) string {
	if data == nil {
		return html
	}

	result := html
	v := reflect.ValueOf(data)

	// Handle different types of data
	switch v.Kind() {
	case reflect.Pointer:
		if v.IsNil() {
			return html
		}
		v = v.Elem()
		fallthrough
	case reflect.Struct:
		// Iterate through struct fields
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i)

			// Get field name (use json tag if available, otherwise use field name)
			fieldName := field.Name
			if jsonTag := field.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
				// Extract json tag name (before comma if there are options)
				if parts := strings.Split(jsonTag, ","); len(parts) > 0 && parts[0] != "" {
					fieldName = parts[0]
				}
			}

			// Convert field value to string
			var valueStr string
			if fieldValue.Kind() == reflect.Pointer {
				if fieldValue.IsNil() {
					valueStr = ""
				} else {
					valueStr = fmt.Sprintf("%v", fieldValue.Elem().Interface())
				}
			} else {
				valueStr = fmt.Sprintf("%v", fieldValue.Interface())
			}

			// Replace {{FieldName}} and {{.FieldName}} with the value
			placeholder1 := fmt.Sprintf("{{%s}}", fieldName)
			placeholder2 := fmt.Sprintf("{{.%s}}", fieldName)
			result = strings.ReplaceAll(result, placeholder1, valueStr)
			result = strings.ReplaceAll(result, placeholder2, valueStr)
		}
	case reflect.Map:
		// Handle map[string]any or similar
		for _, key := range v.MapKeys() {
			keyStr := fmt.Sprintf("%v", key.Interface())
			value := v.MapIndex(key)
			valueStr := fmt.Sprintf("%v", value.Interface())

			// Replace {{KeyName}} and {{.KeyName}} with the value
			placeholder1 := fmt.Sprintf("{{%s}}", keyStr)
			placeholder2 := fmt.Sprintf("{{.%s}}", keyStr)
			result = strings.ReplaceAll(result, placeholder1, valueStr)
			result = strings.ReplaceAll(result, placeholder2, valueStr)
		}
	}

	return result
}
