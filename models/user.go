package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    FirebaseUID  string `gorm:"uniqueIndex;size:128;not null" json:"firebase_uid"`
    Email        string `gorm:"uniqueIndex;size:255;not null" json:"email"`
    Name         string `gorm:"size:100" json:"name"`
    Role         string `gorm:"size:20;default:user" json:"role"`
    EmailVerified bool  `gorm:"default:false" json:"email_verified"`
    LastLoginAt  *int64 `gorm:"index" json:"last_login_at,omitempty"`
}

type CreateProductRequest struct {
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description"`
    Price       float64 `json:"price" binding:"required"`
    Stock       int     `json:"stock"`
    Category    string  `json:"category"`
    ImageURL    string  `json:"image_url"`
}

type UpdateProductRequest struct {
    Name        *string  `json:"name"`
    Description *string  `json:"description"`
    Price       *float64 `json:"price"`
    Stock       *int     `json:"stock"`
    Category    *string  `json:"category"`
    ImageURL    *string  `json:"image_url"`
}