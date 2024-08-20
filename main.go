package main

import (
	"test_sagara/src/config"
	"test_sagara/src/handler"
	"test_sagara/src/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.InitGorm()
	repo := repository.NewPostRepository(db)

	r.POST("/api", handler.CreateBaju(repo))
	r.GET("/api", handler.FindAllBaju(repo))
	r.GET("/api/warna/:warna", handler.FindByWarna(repo))
	r.GET("/api/ukuran/:ukuran", handler.FindByUkuran(repo))
	r.PUT("/bajus/tambah-stok/:id/:jumlah", handler.TambahStok(repo))
	r.PUT("/bajus/kurang-stok/:id/:jumlah", handler.KurangStok(repo))
	r.Run(":3000")
}
