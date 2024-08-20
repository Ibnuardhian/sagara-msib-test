package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"test_sagara/src/models"

	"github.com/gin-gonic/gin"
)

func CreateBaju(repo models.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.BajuRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Convert string fields to lower case
		baju := &models.Baju{
			Warna:  strings.ToLower(req.Warna),
			Ukuran: strings.ToLower(req.Ukuran),
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

		bajus, err := repo.FindAll(context.Background())
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
		warna = strings.ToLower(warna)

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
		ukuran = strings.ToLower(ukuran)

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

// UpdateBaju memperbarui data baju
func UpdateBaju(repo models.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var req models.BajuRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		baju := &models.Baju{
			Id:     uint(id),
			Warna:  req.Warna,
			Ukuran: req.Ukuran,
			Harga:  req.Harga,
			Stok:   req.Stok,
		}

		updatedBaju, err := repo.Update(context.Background(), baju)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.BajuResponse{
			Id:     updatedBaju.Id,
			Warna:  updatedBaju.Warna,
			Ukuran: updatedBaju.Ukuran,
			Harga:  updatedBaju.Harga,
			Stok:   updatedBaju.Stok,
		})
	}
}

// FindLowStock menampilkan baju dengan stok kurang dari 5
func FindLowStock(repo models.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		bajus, err := repo.FindLowStock(context.Background())
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

// TambahStok adds stock to an existing Baju
func TambahStok(repo models.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id parameter"})
			return
		}

		var req models.TambahStokReq
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		baju, err := repo.TambahStock(context.Background(), uint(id), req.Jumlah)
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

// KurangStok reduces stock from an existing Baju
func KurangStok(repo models.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
			return
		}

		var req models.KurangStokReq
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		baju, err := repo.KurangStock(context.Background(), uint(id), req.Jumlah)
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

// FindOutOfStock menampilkan baju dengan stok 0
func FindOutOfStock(repo models.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		bajus, err := repo.FindOutOfStock(context.Background())
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

// DeleteBaju menghapus baju berdasarkan id
func DeleteBaju(repo models.PostRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := repo.Delete(context.Background(), uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Baju deleted successfully"})
	}
}
