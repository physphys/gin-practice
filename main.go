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
	Status    TodoStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TodoStatus express status of todo
type TodoStatus int

const (
	undone TodoStatus = iota
	doing
	done
)

func main() {
	dbInit()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
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

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost user=todo_owner dbname=todo sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
