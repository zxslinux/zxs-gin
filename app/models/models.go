package models

import (
    "gorm.io/gorm"
    "time"
)

// 自增ID主键
type ID struct {
    ID uint `json:"id" gorm:"primaryKey"`
}

// 创建、更新时间
type Timestamps struct {
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// 软删除
type SoftDeletes struct {
    DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Env struct {
    ID
    Name string `json:"name" gorm:"not null;comment:环境英文名称"`
    DisplayName string `json:"display_name" gorm:"not null;index;comment:环境中文名称"`
    Timestamps
    SoftDeletes
}