package persistence

import "github.com/yourname/ddd-app/domain/entities"

func SaveUser(u *entities.User) error { /* persist to DB */ return nil }
