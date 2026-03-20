package service

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"homework03/config"
	"homework03/models"
	"homework03/utils"
	"log"
	"os"
	"path/filepath"
	"testing"
)

// Data JSON 文件
type Data struct {
	Users    []models.User    `json:"users"`
	Posts    []models.Post    `json:"posts"`
	Comments []models.Comment `json:"comments"`
}

func initTestData(db *gorm.DB) {

	rootDir, e := utils.GetRootPath()
	if e != nil {
		log.Fatal(" Not Find Root Path:\n", e)
	}
	filePath := filepath.Join(rootDir, "data.json")
	file, err := os.Open(filePath) // 假设 JSON 文件名为 data.json
	if err != nil {
		log.Fatal("无法打开文件:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("无法打开文件:", err)
		}
	}(file)

	var data Data
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		log.Fatal("JSON 解析失败:", err)
	}

	// 4. 按依赖顺序插入数据（先用户，再帖子，最后评论）
	if len(data.Users) > 0 {
		if err := db.Create(&data.Users).Error; err != nil {
			log.Fatal("插入用户失败:", err)
		}
		log.Printf("已插入 %d 个用户", len(data.Users))
	}

	if len(data.Posts) > 0 {
		if err := db.Create(&data.Posts).Error; err != nil {
			log.Fatal("插入帖子失败:", err)
		}
		log.Printf("已插入 %d 篇帖子", len(data.Posts))
	}

	if len(data.Comments) > 0 {
		if err := db.Create(&data.Comments).Error; err != nil {
			log.Fatal("插入评论失败:", err)
		}
		log.Printf("已插入 %d 条评论", len(data.Comments))
	}

	log.Println("数据导入成功！")
}

func TestWork1(t *testing.T) {
	log.Printf(" db init......")
	db := config.GetDB(true)
	if db == nil {
		log.Printf("db is nil")
	}
	//initTestData(db)
}

func TestWork2(t *testing.T) {
	log.Printf(" db init......")
	db := config.GetDB(false)
	if db == nil {
		log.Printf("db is nil")
	}

	// 查询某个用户发布的所有文章及其对应的评论信息
	log.Printf("查询某个用户发布的所有文章及其对应的评论信息")
	var user models.User
	// if ue := db.Where("id=?", 5).First(&user).Error; ue != nil {
	//if ue := db.First(&user, 4).Error; ue != nil {
	//	if errors.Is(ue, gorm.ErrRecordNotFound) {
	//		t.Fatalf("no active user found")
	//	}
	//	t.Fatalf("query first active user: %v", ue)
	//}
	//t.Logf("first active user: %+v", user)

	if pe := db.Preload("Posts").Preload("Comments").First(&user, 4).Error; pe != nil {
		if errors.Is(pe, gorm.ErrRecordNotFound) {
			t.Fatalf("no active user found")
		}
		t.Fatalf("query first active user: %v", pe)
	}
	t.Logf("first active user: %+v", user)
}

func TestWork3(t *testing.T) {

	db := config.GetDB(false)
	if db == nil {
		log.Printf("db is nil")
	}
	log.Printf("查询评论数量最多的文章信息")
	type CommentCount struct {
		ID      uint
		UserID  uint
		Title   string
		Content string
		Total   int64
	}
	cc := CommentCount{}
	if er := db.Model(&models.Post{}).
		Select(" posts.id,posts.user_id,posts.title,posts.content,count(comments.id) total").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("total desc").
		Limit(1).
		Scan(&cc).
		Error; er != nil {
		if errors.Is(er, gorm.ErrRecordNotFound) {
			t.Fatalf("no active user found")
		}
		t.Fatalf("query first active user: %v", er)
	}
	t.Logf("first active user: %+v", cc)
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段

func TestWork4(t *testing.T) {

	db := config.GetDB(false)
	if db == nil {
		log.Printf("db is nil")
	}

	p := models.Post{Title: " add new 1", Content: " add new Content", UserID: 5}
	if err := db.Create(&p).Error; err != nil {
		t.Fatalf("create user: %v", err)
	}
	t.Logf("add new post: %+v", p)

}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"

func TestWork5(t *testing.T) {
	db := config.GetDB(false)
	if db == nil {
		log.Printf("db is nil")
	}

	//me := db.AutoMigrate(&models.Post{})
	//if me != nil {
	//	log.Fatal("Failed to migrate Post:", me)
	//}

	if err := db.Where("post_id =?", 1).First(&models.Comment{}).Error; err != nil {
		t.Fatalf(" post comment not found %v", err)
	} else {
		if err := db.Delete(&models.Comment{ID: 2}).Error; err != nil {
			t.Fatalf("delete: %v", err)
		}
		var post models.Post
		if e := db.First(&post, 1).Error; e != nil {
			t.Fatalf(" get post err: %v", err)
		}
		t.Logf("add new post: %+v", post)
	}
}
