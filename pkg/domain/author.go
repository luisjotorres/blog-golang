package domain

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Author struct {
	ID        uint          `json:"id" ,gorm:"primaryKey,autoIncrement"`
	Name      string        `json:"name" ,gorm:"not null"`
	Email     string        `json:"email" ,gorm:"unique"`
	Password  string        `json:"password,omitempty"`
	CreatedAt time.Time     `json:"created_at" ,gorm:"autoCreateTime"`
	UpdatedAt time.Time     `json:"updated_at" ,gorm:"autoUpdateTime"`
	Profile   AuthorProfile `json:"profile"`
	Posts     []Post        `json:"posts"`
}

type AuthorProfile struct {
	gorm.Model            //Esto incluye ID, createdAt, updatedAt, deletedAt
	Description    string `json:"description" gorm:"size:700"`
	ProfilePicture string `json:"profile_picture"`
	AuthorID       uint   `json:"author_id"`
}

func (a *Author) BeforeCreate(tx *gorm.DB) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(a.Password), 8)
	a.Password = string(hashed)
	return nil
}

func (a *Author) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	return err == nil
}
