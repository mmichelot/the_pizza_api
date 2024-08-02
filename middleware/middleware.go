package middleware

import "gorm.io/gorm"

type Middleware struct {
	Db *gorm.DB
}

func NewMiddleware(Db *gorm.DB) *Middleware {
	return &Middleware{Db: Db}
}
