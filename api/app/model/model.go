package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Wine struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Name        string     `json:"name"`
	Brand       string     `json:"brand"`
	Description string     `gorm:"type:text" json:"description"`
	Year        int        `json:"year"`
	Country     string     `json:"country"`
	Quantity    int        `json:"quantity"`
	Status      bool       `json:"status"`
}

func (e *Wine) Disable() {
	e.Status = false
}

func (p *Wine) Enable() {
	p.Status = true
}

// DBMigrate will create and migrate the tables
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Wine{})
	return db
}
