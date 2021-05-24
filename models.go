package main

import "database/sql"

type Client struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	CreatedAt sql.NullString `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
	DeletedAt sql.NullString `json:"deleted_at"`
}

type Project struct {
	ID         int     `json:"id"`
	ClientID   int     `json:"client_id"`
	Name       string  `json:"name"`
	PayPerHour float64 `json:"pph,omitempty"`
	Complete   bool    `json:"complete"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	DeletedAt  string  `json:"deleted_at"`
}

type Task struct {
	ID                int     `json:"id"`
	ProjectID         string  `json:"project_id"`
	Name              string  `json:"name"`
	StartDateTime     string  `json:"start_datetime"`
	EndDateTime       string  `json:"end_datetime"`
	LastPauseDateTime string  `json:"last_pause_datetime"`
	Complete          bool    `json:"complete"`
	CurrentlyActive   bool    `json:"currently_active"`
	TotalTime         float64 `json:"total_time,omitempty"`
	TotalPay          float64 `json:"total_pay,omitempty"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
	DeletedAt         string  `json:"deleted_at"`
}

type Note struct {
	ID        int    `json:"id"`
	TableID   int    `json:"table_id"`
	TableName string `json:"table_name"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
