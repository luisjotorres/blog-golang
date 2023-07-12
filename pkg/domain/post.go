package domain

import (
	"gorm.io/gorm"
	"time"
)

type OwnModel struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

type Post struct {
	OwnModel
	Title      string     `json:"title"`
	Content    string     `json:"content" gorm:"size:7000"`
	AuthorID   uint       `json:"author_id"`
	CategoryId uint       `json:"category_id"`
	Reactions  Reactions  `json:"reactions"`
	Comments   []Comments `json:"comments,omitempty"`
}

type Category struct {
	gorm.Model
	Name string `json:"name"`
	Post Post   `json:"post"`
}

type Reactions struct {
	OwnModel
	Likes  uint `json:"likes"`
	Views  uint `json:"views"`
	PostID uint `json:"post_id"`
}

type Comments struct {
	gorm.Model
	Comment string `json:"comment"`
	PostID  uint   `json:"post_id"`
}
