package main

import (
	"test_sagara/src/config"
	"test_sagara/src/handler"
	"test_sagara/src/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	db := config.InitGorm()
	repo := repository.NewPostRepository(db)

	r.POST("/api", handler.CreateBaju(repo))
	r.GET("/api", handler.FindAllBaju(repo))
	r.GET("/api/warna/:warna", handler.FindByWarna(repo))
	r.GET("/api/ukuran/:ukuran", handler.FindByUkuran(repo))
	r.GET("/api/baju/lowstock", handler.FindLowStock(repo))
	r.GET("/api/baju/outstock", handler.FindOutOfStock(repo))
	r.PUT("/api/baju/:id", handler.UpdateBaju(repo))
	r.PUT("/api/baju/tambah/:id", handler.TambahStok(repo))
	r.PUT("/api/baju/kurang/:id", handler.KurangStok(repo))
	r.DELETE("/api/delete/:id", handler.DeleteBaju(repo))
	r.Run(":3000")
}
