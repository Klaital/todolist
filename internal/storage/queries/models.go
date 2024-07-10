// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package queries

import (
	"time"
)

type Chore struct {
	ID            int64
	OwnerID       int32
	Name          string
	LastCompleted time.Time
	NextDeadline  time.Time
	FrequencyDays int32
}

type ListItem struct {
	ID         int64
	TodolistID int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     int32
	Descr      string
}

type Todolist struct {
	ID        int64
	OwnerID   int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID             int64
	Authprovider   string
	AuthproviderID string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UsersPermission struct {
	ID     int64
	UserID int32
	Role   string
}
