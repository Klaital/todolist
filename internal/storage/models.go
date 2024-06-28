package storage

import "time"

type User struct {
	Id             uint64
	AuthPlatform   string
	AuthPlatformId string
}

type TodoItem struct {
	Id        uint64
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type TodoList struct {
	Id        uint64
	Items     []TodoItem
	CreatedAt time.Time
	UpdatedAt time.Time
}
