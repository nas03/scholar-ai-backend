package repositories

import (
	"context"

	"github.com/nas03/scholar-ai/backend/internal/models"
	"gorm.io/gorm"
)

type IMailRepository interface {
	GetMailTemplate(ctx context.Context, id int) (*models.Mail, error)
}

type MailRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository with the given database connection.
func NewMailRepository(db *gorm.DB) IMailRepository {
	return &MailRepository{db: db}
}

func (r *MailRepository) GetMailTemplate(ctx context.Context, id int) (*models.Mail, error) {
	var mailTemplate models.Mail
	err := r.db.WithContext(ctx).Select("id, subject, header, body, footer").First(&mailTemplate).Error
	if err != nil {
		return nil, err
	}

	return &mailTemplate, nil

}
