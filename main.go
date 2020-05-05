package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Todo is todo item
type Todo struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
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

	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		createTodo(&Todo{Text: text, Status: status})
		ctx.Redirect(http.StatusSeeOther, "/")
	})

	router.GET("/todos/:id/delete_check", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := findTodo(id)
		ctx.HTML(http.StatusOK, "delete.html", gin.H{"todo": todo})
	})

	router.POST("/todos/:id/delete", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		deleteTodo(id)
		ctx.Redirect(http.StatusSeeOther, "/")
	})

	router.Run()
}

func dbInit() {
	db := connectDB()
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

func findTodo(id int) Todo {
	db := connectDB()
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

func createTodo(todo *Todo) {
	db := connectDB()
	db.Create(todo)
}

func deleteTodo(id int) {
	db := connectDB()
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
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
