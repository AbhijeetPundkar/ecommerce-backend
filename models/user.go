package models

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Name     string `json:"name" gorm:"type:varchar(100);not null"`
    Username string `json:"username" gorm:"type:varchar(100);unique;not null"`
    Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
    Password string `json:"-" gorm:"not null"`
}

func (u *User) HashPassword(password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    return nil
}

func (u *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}

func (u *User) Create(db *gorm.DB) error {
    return db.Create(u).Error
}

func (u *User) Validate() error {
    if u.Name == "" {
        return errors.New("name is required")
    }

    if u.Username == "" {
        return errors.New("username is required")
    }

    if u.Email == "" {
        return errors.New("email is required")
    }

    if !isValidEmail(u.Email) {
        return errors.New("invalid email format")
    }

    if u.Password == "" {
        return errors.New("password is required")
    }

    if len(u.Password) < 6 {
        return errors.New("password must be at least 6 characters long")
    }

    return nil
}

func isValidEmail(email string) bool {
    re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
    return re.MatchString(email)
}