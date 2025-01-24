package models

import (
	"LemonMovie/global"
	"gorm.io/gorm"
)

type MovieModel struct {
	gorm.Model
	Name        string // 名字
	Description string `gorm:"type:text"` // 描述
	Category    int    // 分类
	TypePid     int    // 子分类
	Class       string // 标签
	CoverURL    string // 封面
	Director    string
	Actor       string
	Area        string
	Language    string
	Year        string

	PlayFrom string
	PlayURL  string `gorm:"type:text"`

	VodID int // 网站视频id
}

// ==================================================Movie========================================
func NewMovieModel() *MovieModel {
	return &MovieModel{}
}
func (m *MovieModel) TableName() string {
	return "movie_models"
}

func (m *MovieModel) FindAllMovieList() []MovieModel {
	var movieList []MovieModel
	global.DB.Limit(20).Offset(0).Order("id desc").Find(&movieList)
	return movieList
}
