package models

import (
	"database/sql"
	"time"
)

// TableCommon provides common timestamp fields for all models.
// GORM auto-manages CreatedAt/UpdatedAt.
type TableCommon struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	UserID          string         `gorm:"primaryKey;type:char(36)" json:"user_id"`
	Username        string         `gorm:"uniqueIndex;not null;size:255" json:"username"`
	Email           string         `gorm:"uniqueIndex;not null;size:255" json:"email"`
	Password        string         `gorm:"type:text;not null" json:"-"` // Never expose password in JSON
	PhoneNumber     sql.NullString `gorm:"size:10" json:"phone_number,omitempty"`
	AccountStatus   int8           `gorm:"not null;default:0" json:"account_status"`    // account status (0=inactive, 1=active)
	IsEmailVerified int8           `gorm:"not null;default:0" json:"is_email_verified"` // email verification (0=unverified, 1=verified)
	IsPhoneVerified int8           `gorm:"not null;default:0" json:"is_phone_verified"` // phone verification (0=unverified, 1=verified)
	TableCommon

	// Relationships (one-to-many)
	Courses []Course `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"courses,omitempty"`
}

func (User) TableName() string {
	return "users"
}

type Course struct {
	ID          int            `gorm:"primaryKey;autoIncrement" json:"id"`
	CourseID    string         `gorm:"not null;index;size:255" json:"course_id"` // Course identifier (e.g., "CS101")
	CourseName  string         `gorm:"not null;size:255" json:"course_name"`
	UserID      string         `gorm:"not null;index;type:char(36)" json:"user_id"`
	Description sql.NullString `gorm:"type:text" json:"description,omitempty"`
	Lecturers   string         `gorm:"type:text;not null" json:"lecturers"` // Comma-separated lecturer names
	Credits     int            `gorm:"not null" json:"credits"`
	GPA         float32        `gorm:"not null;default:0" json:"gpa"`
	SemesterID  int            `gorm:"not null;index" json:"semester_id"`
	TableCommon

	// Relationships
	// Don't include User back-ref to avoid circular JSON; fetch separately if needed
	Semester Semester `gorm:"foreignKey:SemesterID;constraint:OnDelete:RESTRICT" json:"semester,omitempty"`
	Tags     []Tag    `gorm:"many2many:course_tags;" json:"tags,omitempty"`
}

func (Course) TableName() string {
	return "courses"
}

type Semester struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"not null;size:255" json:"name"`
	StartDate time.Time `gorm:"not null;index" json:"start_date"` // Index added per SQL schema
	EndDate   time.Time `gorm:"not null" json:"end_date"`
	TableCommon

	// Relationships (one-to-many)
	Courses []Course `gorm:"foreignKey:SemesterID;constraint:OnDelete:RESTRICT" json:"courses,omitempty"`
}

func (Semester) TableName() string {
	return "semesters"
}

type Tag struct {
	ID    int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"not null;index;size:255" json:"name"`          // Index added per SQL schema
	Color string `gorm:"not null;default:#808080;size:7" json:"color"` // hex color
	TableCommon

	// Relationships (many-to-many)
	Courses []Course `gorm:"many2many:course_tags;" json:"courses,omitempty"`
}

func (Tag) TableName() string {
	return "tags"
}

// CourseTag is the explicit join table for the many2many relationship.
type CourseTag struct {
	CourseID int `gorm:"primaryKey;index"`
	TagID    int `gorm:"primaryKey;index"`
}

func (CourseTag) TableName() string {
	return "course_tags"
}

type Mail struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Subject string `gorm:"not null;type:text" json:"subject"`
	Header  string `gorm:"type:text" json:"header"`
	Body    string `gorm:"not null;type:text" json:"body"`
	Footer  string `gorm:"type:text" json:"footer"`
	TableCommon
}

func (Mail) TableName() string {
	return "mail"
}
