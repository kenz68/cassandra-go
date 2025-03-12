package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kenz68/cassandra-go/src/cmd/utils"
	"github.com/kenz68/cassandra-go/src/pkg/client/cassandra"
	"github.com/kenz68/cassandra-go/src/pkg/handler"
	"github.com/kenz68/cassandra-go/src/pkg/repository/db"
)

// Initilize FIN Web Framework and prepare routing, connect to Cassandra Client, Register to Repository and Handler
func Initilize(config utils.Configuration) {
	// Creates a gin router with default middleware:
	router := gin.Default()

	fmt.Printf("%+v\n", config)

	//register database ,repositories and handlers
	session := cassandra.ConnectDatabase(config.Database.Url, config.Database.Keyspace)

	repository := db.NewTodoRepository(session)
	orderHandler := handler.NewTodoHandler(&repository)

	router.GET("/ping", orderHandler.HealthCheck)

	router.POST("api/v1/todo/", orderHandler.CreateTodo)
	router.GET("api/v1/todo/:id", orderHandler.GetTodoById)

	//run the server :8080
	router.Run(":" + config.Host.Port + "")
}
