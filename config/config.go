package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	ExecPath     string
	JsonData     configData
	ServerConfig serverConfig
	DB           *gorm.DB
)

type systemConfig struct {
	SiteName      string `json:"site_name"`
	SiteLogo      string `json:"site_logo"`
	SiteIcp       string `json:"site_icp"`
	SiteCopyright string `json:"site_copyright"`
	BaseUrl       string `json:"base_url"`
	AdminUri      string `json:"admin_uri"`
	SiteClose     int    `json:"site_close"`
	SiteCloseTips string `json:"site_close_tips"`
}

type contentConfig struct {
	RemoteDownload int    `json:"remote_download"`
	FilterOutlink  int    `json:"filter_outlink"`
	ResizeImage    int    `json:"resize_image"`
	ResizeWidth    int    `json:"resize_width"`
	ThumbCrop      int    `json:"thumb_crop"`
	ThumbWidth     int    `json:"thumb_width"`
	ThumbHeight    int    `json:"thumb_height"`
	DefaultThumb   string `json:"default_thumb"`
}

type indexConfig struct {
	SeoTitle       string `json:"seo_title"`
	SeoKeywords    string `json:"seo_keywords"`
	SeoDescription string `json:"seo_description"`
}

type configData struct {
	DB     mysqlConfig  `json:"mysql"`
	Server serverConfig `json:"server"`
	// setting
	System  systemConfig  `json:"system"`
	Content contentConfig `json:"content"`
	Index   indexConfig   `json:"index"`
	// plugin
	PluginPush    pluginPushConfig    `json:"plugin_push"`
	PluginSitemap pluginSitemapConfig `json:"plugin_sitemap"`
}

func initPath() {
	// 目录路径分隔符，这里为\\
	sep := string(os.PathSeparator)
	// 执行命令时所在目录
	ExecPath, _ = os.Getwd()
	length := utf8.RuneCountInString(ExecPath)
	lastchar := ExecPath[length-1:]
	if lastchar != sep {
		ExecPath = ExecPath + sep
	}
}

func initJSON() {
	jsonfile := fmt.Sprintf("%sconfig.json", ExecPath)
	rawConfig, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		// 未初始化
		rawConfig = []byte("{\"db\":{},\"server\":{\"site_name\":\"irisweb 博客\",\"env\": \"development\",\"port\": 8001,\"log_level\":\"debug\"}}")
	}
	if err := json.Unmarshal(rawConfig, &JsonData); err != nil {
		fmt.Println("Invalid Config: ", err.Error())
		os.Exit(-1)
	}
}

func initServer() {
	ServerConfig = JsonData.Server
}

func initDB(setting *mysqlConfig) error {
	var db *gorm.DB
	var err error
	conurl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", setting.User, setting.Password, setting.Host, setting.Port, setting.Database)
	setting.Url = conurl
	db, err = gorm.Open(mysql.Open(conurl), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(1000)
	sqlDB.SetMaxOpenConns(10000)
	sqlDB.SetConnMaxLifetime(-1)
	// db.AutoMigrate(&model.Admin{}, &model.Article{}, &model.ArticleData{}, &model.Attachment{}, &model.Category{})
	DB = db
	return nil
}

// 程序首先执行
func init() {
	initPath()
	initJSON()
	initServer()
	err := initDB(&JsonData.DB)
	if err != nil {
		fmt.Println("Failed To Connect Database: ", err.Error())
		os.Exit(-1)
	}
}
