package models

import "time"

type AuditorialBase struct {
	UserRegister string
	DateRegister time.Time
	UserUpdate   string
	DateUpdate   time.Time
}
