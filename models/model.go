// Package models @Author Gopher
// @Date 2025/1/30 10:32:00
// @Desc
package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type User struct {
	gorm.Model

	Username string `json:"username"`
	Password string `json:"password"`
}

type Post struct {
	gorm.Model
	Title   string
	Content string `gorm:"type:text"`
	Tag     string
}

type Manager interface {
	AddUser(user *User)
	Login(username string) User

	AddPost(post *Post)
	GetAllPost() []Post
	GetPost(pid int) Post
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:wangchen0912@tcp(127.0.0.1:3306)/test_xorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
}

func (mgr *manager) AddUser(user *User) {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) User {
	var user User
	mgr.db.Where("username=?", username).First(&user)
	return user
}

func (mgr *manager) AddPost(post *Post) {
	mgr.db.Create(post)
}

func (mgr *manager) GetAllPost() []Post {
	var posts = make([]Post, 10)
	mgr.db.Find(&posts)
	return posts
}

func (mgr *manager) GetPost(pid int) Post {
	var post Post
	mgr.db.First(&post, pid)
	return post
}
