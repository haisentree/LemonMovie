package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

var DB *gorm.DB

const BASEURL = "https://cj.lziapi.com/api.php/provide/vod/"

type MovieDetail struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Page      string `json:"page"`
	Pagecount int    `json:"pagecount"`
	Limit     string `json:"limit"`
	Total     int    `json:"total"`
	List      []struct {
		VodID            int    `json:"vod_id"`
		TypeID           int    `json:"type_id"`
		TypeID1          int    `json:"type_id_1"`
		GroupID          int    `json:"group_id"`
		VodName          string `json:"vod_name"`
		VodSub           string `json:"vod_sub"`
		VodEn            string `json:"vod_en"`
		VodStatus        int    `json:"vod_status"`
		VodLetter        string `json:"vod_letter"`
		VodColor         string `json:"vod_color"`
		VodTag           string `json:"vod_tag"`
		VodClass         string `json:"vod_class"`
		VodPic           string `json:"vod_pic"`
		VodPicThumb      string `json:"vod_pic_thumb"`
		VodPicSlide      string `json:"vod_pic_slide"`
		VodPicScreenshot string `json:"vod_pic_screenshot"`
		VodActor         string `json:"vod_actor"`
		VodDirector      string `json:"vod_director"`
		VodWriter        string `json:"vod_writer"`
		VodBehind        string `json:"vod_behind"`
		VodBlurb         string `json:"vod_blurb"`
		VodRemarks       string `json:"vod_remarks"`
		VodPubdate       string `json:"vod_pubdate"`
		VodTotal         int    `json:"vod_total"`
		VodSerial        string `json:"vod_serial"`
		VodTv            string `json:"vod_tv"`
		VodWeekday       string `json:"vod_weekday"`
		VodArea          string `json:"vod_area"`
		VodLang          string `json:"vod_lang"`
		VodYear          string `json:"vod_year"`
		VodVersion       string `json:"vod_version"`
		VodState         string `json:"vod_state"`
		VodAuthor        string `json:"vod_author"`
		VodJumpurl       string `json:"vod_jumpurl"`
		VodTpl           string `json:"vod_tpl"`
		VodTplPlay       string `json:"vod_tpl_play"`
		VodTplDown       string `json:"vod_tpl_down"`
		VodIsend         int    `json:"vod_isend"`
		VodLock          int    `json:"vod_lock"`
		VodLevel         int    `json:"vod_level"`
		VodCopyright     int    `json:"vod_copyright"`
		VodPoints        int    `json:"vod_points"`
		VodPointsPlay    int    `json:"vod_points_play"`
		VodPointsDown    int    `json:"vod_points_down"`
		VodHits          int    `json:"vod_hits"`
		VodHitsDay       int    `json:"vod_hits_day"`
		VodHitsWeek      int    `json:"vod_hits_week"`
		VodHitsMonth     int    `json:"vod_hits_month"`
		VodDuration      string `json:"vod_duration"`
		VodUp            int    `json:"vod_up"`
		VodDown          int    `json:"vod_down"`
		VodScore         string `json:"vod_score"`
		VodScoreAll      int    `json:"vod_score_all"`
		VodScoreNum      int    `json:"vod_score_num"`
		VodTime          string `json:"vod_time"`
		VodTimeAdd       int    `json:"vod_time_add"`
		VodTimeHits      int    `json:"vod_time_hits"`
		VodTimeMake      int    `json:"vod_time_make"`
		VodTrysee        int    `json:"vod_trysee"`
		VodDoubanID      int    `json:"vod_douban_id"`
		VodDoubanScore   string `json:"vod_douban_score"`
		VodReurl         string `json:"vod_reurl"`
		VodRelVod        string `json:"vod_rel_vod"`
		VodRelArt        string `json:"vod_rel_art"`
		VodPwd           string `json:"vod_pwd"`
		VodPwdURL        string `json:"vod_pwd_url"`
		VodPwdPlay       string `json:"vod_pwd_play"`
		VodPwdPlayURL    string `json:"vod_pwd_play_url"`
		VodPwdDown       string `json:"vod_pwd_down"`
		VodPwdDownURL    string `json:"vod_pwd_down_url"`
		VodContent       string `json:"vod_content"`
		VodPlayFrom      string `json:"vod_play_from"`
		VodPlayServer    string `json:"vod_play_server"`
		VodPlayNote      string `json:"vod_play_note"`
		VodPlayURL       string `json:"vod_play_url"`
		VodDownFrom      string `json:"vod_down_from"`
		VodDownServer    string `json:"vod_down_server"`
		VodDownNote      string `json:"vod_down_note"`
		VodDownURL       string `json:"vod_down_url"`
		VodPlot          int    `json:"vod_plot"`
		VodPlotName      string `json:"vod_plot_name"`
		VodPlotDetail    string `json:"vod_plot_detail"`
		TypeName         string `json:"type_name"`
	} `json:"list"`
}

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

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	err = db.AutoMigrate(&MovieModel{})
	if err != nil {
		fmt.Println(err)
	}
	DB = db
}

func main() {
	for i := 1; i <= 4938; i++ {
		page := strconv.Itoa(i)
		GetMovie(page)
	}
}

func GetMovie(page string) {
	http.Header{}.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")
	params := url.Values{}
	URL, err := url.Parse(BASEURL)
	if err != nil {
		return
	}
	params.Set("ac", "detail")
	params.Set("pg", page)
	URL.RawQuery = params.Encode()
	urlPath := URL.String()
	fmt.Println(urlPath)

	resp, err := http.Get(urlPath)
	if err != nil {
		fmt.Println(err)
	}
	movieDetail := &MovieDetail{}

	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, movieDetail)

	var movieList []MovieModel
	for _, val := range movieDetail.List {
		movie := MovieModel{
			Name:        val.VodName,
			Description: val.VodBlurb,
			Category:    val.TypeID1,
			TypePid:     val.TypeID,
			Class:       val.VodClass,
			CoverURL:    val.VodPic,
			Director:    val.VodDirector,
			Actor:       val.VodActor,
			Area:        val.VodArea,
			Language:    val.VodLang,
			Year:        val.VodYear,
			PlayFrom:    val.VodPlayFrom,
			PlayURL:     val.VodPlayURL,
			VodID:       val.VodID,
		}
		movieList = append(movieList, movie)
	}
	DB.Create(&movieList)
}
