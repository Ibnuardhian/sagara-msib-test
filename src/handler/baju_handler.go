package handler

import (
	"context"
	"net/http"
	"strconv"
	"test_sagara/src/models"

	"github.com/gin-gonic/gin"
)

func CreateBaju(repo models.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CreateBajuRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		baju := &models.Baju{
			Warna:  req.Warna,
			Ukuran: req.Ukuran,
			Harga:  req.Harga,
			Stok:   req.Stok,
		}

		createdBaju, err := repo.Create(context.Background(), baju)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.BajuResponse{
			Id:     createdBaju.Id,
			Warna:  createdBaju.Warna,
			Ukuran: createdBaju.Ukuran,
			Harga:  createdBaju.Harga,
			Stok:   createdBaju.Stok,
		})
	}
}

func FindAllBaju(repo models.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		stok := c.Query("stok") // Assuming you pass stok as a query parameter
		stokValue, err := strconv.Atoi(stok)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stok parameter"})
			return
		}

		bajus, err := repo.FindAll(context.Background(), stokValue)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var response []*models.BajuResponse
		for _, baju := range bajus {
			response = append(response, &models.BajuResponse{
				Id:     baju.Id,
				Warna:  baju.Warna,
				Ukuran: baju.Ukuran,
				Harga:  baju.Harga,
				Stok:   baju.Stok,
			})
		}

		c.JSON(http.StatusOK, models.FindAllBajuResponse{
			Bajus: response,
		})
	}
}

func FindByWarna(repo models.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		warna := c.Param("warna")

		baju, err := repo.FindByWarna(context.Background(), warna)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.BajuResponse{
			Id:     baju.Id,
			Warna:  baju.Warna,
			Ukuran: baju.Ukuran,
			Harga:  baju.Harga,
			Stok:   baju.Stok,
		})
	}
}

func FindByUkuran(repo models.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		ukuran := c.Param("ukuran")

		baju, err := repo.FindByUkuran(context.Background(), ukuran)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.BajuResponse{
			Id:     baju.Id,
			Warna:  baju.Warna,
			Ukuran: baju.Ukuran,
			Harga:  baju.Harga,
			Stok:   baju.Stok,
		})
	}
}

func TambahStok(repo models.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		jumlahStr := c.Param("jumlah")

		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
			return
		}

		jumlah, err := strconv.Atoi(jumlahStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid jumlah parameter"})
			return
		}

		baju, err := repo.TambahStok(context.Background(), uint(id), jumlah)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.BajuResponse{
			Id:     baju.Id,
			Warna:  baju.Warna,
			Ukuran: baju.Ukuran,
			Harga:  baju.Harga,
			Stok:   baju.Stok,
		})
	}
}

func KurangStok(repo models.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		jumlahStr := c.Param("jumlah")

		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
			return
		}

		jumlah, err := strconv.Atoi(jumlahStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid jumlah parameter"})
			return
		}

		baju, err := repo.KurangStok(context.Background(), uint(id), jumlah)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.BajuResponse{
			Id:     baju.Id,
			Warna:  baju.Warna,
			Ukuran: baju.Ukuran,
			Harga:  baju.Harga,
			Stok:   baju.Stok,
		})
	}
}