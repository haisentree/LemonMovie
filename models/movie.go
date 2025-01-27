package models

import (
	"LemonMovie/global"
	"gorm.io/gorm"
	"strconv"
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

// 分页查询所有movie
func (m *MovieModel) FindAllMovieList() ([]MovieModel, int64) {
	var movieList []MovieModel
	var MaxCount int64
	global.DB.Limit(40).Offset(0).Order("id desc").Find(&movieList)
	global.DB.Model(&MovieModel{}).Count(&MaxCount)
	return movieList, MaxCount
}

// 分页查询所有movie
func (m *MovieModel) FindMovieByPage(page int64) ([]MovieModel, int64) {
	page = page - 1
	var movieList []MovieModel
	var MaxCount int64
	strInt := strconv.FormatInt(page, 10)
	pageInt, _ := strconv.Atoi(strInt)

	global.DB.Limit(global.PageLimiter).Offset(global.PageLimiter * pageInt).Order("id desc").Find(&movieList)
	global.DB.Model(&MovieModel{}).Count(&MaxCount)
	return movieList, MaxCount
}

// 根据分类查询movie
func (m *MovieModel) FindByCategory(id int) []MovieModel {
	var movieList []MovieModel
	global.DB.Limit(40).Offset(0).Order("id desc").Where("type_pid = ?", id).Find(&movieList)
	return movieList
}

// 通过ID查询movie
func (m *MovieModel) FindByID(id int) MovieModel {
	var movie MovieModel
	global.DB.Where("id = ?", id).Find(&movie)
	return movie
}
