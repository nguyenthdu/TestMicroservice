package main

import (
	eureka "github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang-service1/controller"
	"golang-service1/db"
	"log"
	"time"
)

func main() {
	db.Init()

	// Đăng ký CategoryService với Eureka
	client := eureka.NewClient([]string{
		"http://localhost:8761/eureka",
	})
	instance := eureka.NewInstanceInfo("category-service", "category-service", "localhost", 8081, 30, false)
	client.RegisterInstance("CATEGORY-SERVICE", instance)
	go func() {
		for {
			client.SendHeartbeat(instance.App, instance.HostName)
			time.Sleep(10 * time.Second)
		}
	}()

	r := gin.Default()
	r.POST("/categories", controller.CreateCategory)
	r.GET("/categories/:id", controller.GetCategory)
	r.GET("/metrics", gin.WrapH(promhttp.Handler())) // Prometheus metrics endpoint

	log.Fatal(r.Run(":8081"))
}

//
//func main() {
//	db.Init()
//
//	r := gin.Default()
//	r.POST("/categories", controller.CreateCategory)
//	r.GET("/categories/:id", controller.GetCategory)
//
//	r.Run(":8081") // Run on port 8081
//}
