package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Todo is todo item
type Todo struct {
	ID        uint `gorm:"primary_key"`
	Text      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	dbInit()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		todos := getAllTodo()
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	router.Run()
}

func dbInit() {
	db := connectDB()
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

func createTodo(todo *Todo) {
	db := connectDB()
	db.Create(todo)
}

func getAllTodo() []Todo {
	db := connectDB()
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost user=todo_owner dbname=todo sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
