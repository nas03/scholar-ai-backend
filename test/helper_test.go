package test

import (
	"context"
	"strings"
	"testing"

	"github.com/nas03/scholar-ai/backend/internal/helper"
	"github.com/nas03/scholar-ai/backend/internal/models"
)

var template string = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<title>Verification Code - ScholarAI</title>
		<style>
			/* Base Reset */
			body {
				margin: 0;
				padding: 0;
				background-color: #f3f3f5;
				font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto,
					Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji',
					'Segoe UI Symbol';
				-webkit-font-smoothing: antialiased;
				-moz-osx-font-smoothing: grayscale;
			}
			table {
				border-collapse: collapse;
				mso-table-lspace: 0pt;
				mso-table-rspace: 0pt;
			}
			td {
				padding: 0;
			}
			img {
				border: 0;
				display: block;
				outline: none;
				text-decoration: none;
			}
			a {
				text-decoration: none;
				color: #030213;
			}

			/* Responsive adjustments */
			@media only screen and (max-width: 600px) {
				.wrapper {
					padding: 20px !important;
				}
				.main-table {
					width: 100% !important;
				}
				.content-padding {
					padding: 24px !important;
				}
				.otp-code {
					font-size: 36px !important;
					letter-spacing: 8px !important;
				}
			}
		</style>
		<!--[if mso]>
			<noscript>
				<xml>
					<o:OfficeDocumentSettings>
						<o:PixelsPerInch>96</o:PixelsPerInch>
					</o:OfficeDocumentSettings>
				</xml>
			</noscript>
		<![endif]-->
	</head>
	<body style="background-color: #f3f3f5; margin: 0; padding: 0">
		<center
			class="wrapper"
			style="
				width: 100%;
				table-layout: fixed;
				background-color: #f3f3f5;
				padding-top: 40px;
				padding-bottom: 40px;
			">
			<div
				style="
					max-width: 540px;
					background-color: #ffffff;
					border-radius: 16px;
					overflow: hidden;
					box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
				">
				<!-- Header -->
				<table
					width="100%"
					cellpadding="0"
					cellspacing="0"
					border="0"
					style="background-color: #ffffff; border-bottom: 1px solid #f0f0f0">
					<tr>
						<td align="center" style="padding: 24px">
							<a
								href="#"
								style="
									font-size: 24px;
									font-weight: 700;
									color: #030213;
									text-decoration: none;
									display: inline-block;
								">
								<span
									style="
										font-size: 24px;
										vertical-align: middle;
										margin-right: 8px;
									"
									>🎓</span
								>
								<span style="vertical-align: middle">ScholarAI</span>
							</a>
						</td>
					</tr>
				</table>

				<!-- Main Content -->
				<table
					width="100%"
					cellpadding="0"
					cellspacing="0"
					border="0"
					style="background-color: #ffffff">
					<tr>
						<td
							class="content-padding"
							style="padding: 40px 40px 32px 40px; text-align: center">
							<h1
								style="
									margin: 0 0 16px 0;
									font-size: 24px;
									font-weight: 600;
									color: #030213;
									letter-spacing: -0.5px;
								">
								Verify your identity
							</h1>
							<p
								style="
									margin: 0 0 32px 0;
									font-size: 16px;
									line-height: 1.6;
									color: #717182;
								">
								Enter the code below to verify your email address and complete
								your account setup.
							</p>

							<!-- OTP Box -->
							<div
								style="
									background-color: #f8f9fa;
									border-radius: 12px;
									padding: 24px;
									margin-bottom: 32px;
									border: 1px solid #ececf0;
								">
								<div
									class="otp-code"
									style="
										font-family: 'SF Mono', 'Courier New', Courier, monospace;
										font-size: 42px;
										font-weight: 700;
										color: #030213;
										letter-spacing: 12px;
										line-height: 1;
									">
									{{.otp}}
								</div>
								<div style="font-size: 13px; color: #717182; margin-top: 12px">
									This code will expire in 10 minutes
								</div>
							</div>

							<p
								style="
									margin: 0;
									font-size: 14px;
									line-height: 1.6;
									color: #717182;
								">
								If you didn't request this code, you can safely ignore this
								email.
							</p>
						</td>
					</tr>
				</table>

				<!-- Device Details Section -->
				<table
					width="100%"
					cellpadding="0"
					cellspacing="0"
					border="0"
					style="background-color: #ffffff">
					<tr>
						<td style="padding: 0 40px 40px 40px">
							<div
								style="
									background-color: #fcfcfc;
									border: 1px solid #f0f0f0;
									border-radius: 8px;
									padding: 20px;
								">
								<table width="100%" cellpadding="0" cellspacing="0" border="0">
									<tr>
										<td
											colspan="2"
											style="
												padding-bottom: 12px;
												border-bottom: 1px solid #f0f0f0;
												margin-bottom: 12px;
											">
											<p
												style="
													margin: 0;
													font-size: 12px;
													font-weight: 600;
													color: #717182;
													text-transform: uppercase;
													letter-spacing: 0.5px;
												">
												Request Details
											</p>
										</td>
									</tr>
									<tr>
										<td
											style="
												padding-top: 12px;
												width: 50%;
												vertical-align: top;
											">
											<p
												style="
													margin: 0 0 4px 0;
													font-size: 11px;
													color: #717182;
												">
												Time
											</p>
											<p
												style="
													margin: 0;
													font-size: 13px;
													color: #030213;
													font-weight: 500;
												">
												{{.Time}}
											</p>
										</td>
										<td
											style="
												padding-top: 12px;
												width: 50%;
												vertical-align: top;
											">
											<p
												style="
													margin: 0 0 4px 0;
													font-size: 11px;
													color: #717182;
												">
												Device
											</p>
											<p
												style="
													margin: 0;
													font-size: 13px;
													color: #030213;
													font-weight: 500;
												">
												{{.OperatingSystem}}
											</p>
										</td>
									</tr>
									<tr>
										<td
											style="
												padding-top: 16px;
												width: 50%;
												vertical-align: top;
											">
											<p
												style="
													margin: 0 0 4px 0;
													font-size: 11px;
													color: #717182;
												">
												Browser
											</p>
											<p
												style="
													margin: 0;
													font-size: 13px;
													color: #030213;
													font-weight: 500;
												">
												{{.Browser}}
											</p>
										</td>
										<td
											style="
												padding-top: 16px;
												width: 50%;
												vertical-align: top;
											">
											<p
												style="
													margin: 0 0 4px 0;
													font-size: 11px;
													color: #717182;
												">
												Location
											</p>
											<p
												style="
													margin: 0;
													font-size: 13px;
													color: #030213;
													font-weight: 500;
												">
												{{.Location}}
											</p>
										</td>
									</tr>
								</table>
							</div>
						</td>
					</tr>
				</table>

				<!-- Footer -->
				<table
					width="100%"
					cellpadding="0"
					cellspacing="0"
					border="0"
					style="background-color: #f9f9f9; border-top: 1px solid #f0f0f0">
					<tr>
						<td align="center" style="padding: 24px 40px">
							<p
								style="
									margin: 0 0 12px 0;
									font-size: 12px;
									color: #717182;
									line-height: 1.5;
								">
								This email was sent to
								<span style="color: #030213; font-weight: 500"
									>{{.RecipientEmail}}</span
								>
							</p>
							<p
								style="
									margin: 0;
									font-size: 12px;
									color: #9ca3af;
									line-height: 1.5;
								">
								&copy; {{.Year}} ScholarAI. All rights reserved.
							</p>
							<div style="margin-top: 16px">
								<a
									href="#"
									style="
										font-size: 12px;
										color: #717182;
										text-decoration: underline;
										margin: 0 8px;
									"
									>Privacy Policy</a
								>
								<a
									href="#"
									style="
										font-size: 12px;
										color: #717182;
										text-decoration: underline;
										margin: 0 8px;
									"
									>Terms of Service</a
								>
								<a
									href="#"
									style="
										font-size: 12px;
										color: #717182;
										text-decoration: underline;
										margin: 0 8px;
									"
									>Support</a
								>
							</div>
						</td>
					</tr>
				</table>
			</div>
		</center>
	</body>
</html>`

// TestReplaceParameters tests the ReplaceParameters function with OTP verification mail template
func TestReplaceParameters(t *testing.T) {
	mailHelper := helper.NewMailHelper()
	ctx := context.Background()

	// Test data with OTP
	testData := models.OTPVerificationMail{
		OTP: 123456,
	}

	// Replace parameters in template
	result := mailHelper.ReplaceParameters(ctx, template, testData)

	// Verify that {{.otp}} was replaced (using JSON tag name)
	expectedOTP := "123456"
	if !strings.Contains(result, expectedOTP) {
		t.Errorf("Expected OTP %s to be in result, but it wasn't found", expectedOTP)
	}

	// Verify that {{.otp}} placeholder is gone
	if strings.Contains(result, "{{.otp}}") {
		t.Error("Template placeholder {{.otp}} was not replaced")
	}

	// Print result for manual inspection
	t.Logf("Generated HTML length: %d characters", len(result))
}

// TestReplaceParametersWithMap tests ReplaceParameters with a map instead of struct
func TestReplaceParametersWithMap(t *testing.T) {
	mailHelper := helper.NewMailHelper()
	ctx := context.Background()

	// Test with map data (using lowercase key to match JSON tag)
	testData := map[string]interface{}{
		"otp": 123456,
	}

	result := mailHelper.ReplaceParameters(ctx, template, testData)

	// Verify replacement
	if !strings.Contains(result, "123456") {
		t.Error("OTP value was not replaced in template")
	}

	if strings.Contains(result, "{{.otp}}") {
		t.Error("Template placeholder {{.otp}} was not replaced")
	}
}

// MailHelperTest is a convenience function for manual testing
// You can call this from a main function or another test
func MailHelperTest() string {
	mailHelper := helper.NewMailHelper()
	ctx := context.Background()
	html := mailHelper.ReplaceParameters(ctx, template, models.OTPVerificationMail{OTP: 123456})
	return html
}
