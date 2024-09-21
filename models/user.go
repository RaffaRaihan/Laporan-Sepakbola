package models

import (
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// Fungsi untuk hashing password
func (u *User) HashPassword(password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    return nil
}

// Fungsi untuk memeriksa kecocokan password
func (u *User) CheckPassword(providedPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
}
