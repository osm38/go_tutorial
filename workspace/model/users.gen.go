// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID              int32     `gorm:"column:id" json:"id"`
	Name            string    `gorm:"column:name" json:"name"`
	Age             int32     `gorm:"column:age" json:"age"`
	Password        string    `gorm:"column:password" json:"password"`
	LastLogin       time.Time `gorm:"column:last_login" json:"last_login"`
	Deleted         bool      `gorm:"column:deleted" json:"deleted"`
	CreateUser      string    `gorm:"column:create_user" json:"create_user"`
	CreateTimestamp time.Time `gorm:"column:create_timestamp" json:"create_timestamp"`
	UpdateUser      string    `gorm:"column:update_user" json:"update_user"`
	UpdateTimestamp time.Time `gorm:"column:update_timestamp" json:"update_timestamp"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
