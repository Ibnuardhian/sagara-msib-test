package repository

import (
	"context"
	"errors"
	"test_sagara/src/models"

	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) models.PostRepository {
	return &PostRepositoryImpl{DB: db}
}

func (r *PostRepositoryImpl) Create(ctx context.Context, baju *models.Baju) (*models.Baju, error) {
	if err := r.DB.WithContext(ctx).Create(baju).Error; err != nil {
		return nil, err
	}
	return baju, nil
}

func (r *PostRepositoryImpl) FindAll(ctx context.Context, stok int) ([]*models.Baju, error) {
	var bajus []*models.Baju
	if err := r.DB.WithContext(ctx).Where("stok > ?", stok).Find(&bajus).Error; err != nil {
		return nil, err
	}
	return bajus, nil
}


func (r *PostRepositoryImpl) FindByWarna(ctx context.Context, warna string) (*models.Baju, error) {
	var baju models.Baju
	if err := r.DB.WithContext(ctx).Where("warna = ?", warna).First(&baju).Error; err != nil {
		return nil, err
	}
	return &baju, nil
}

func (r *PostRepositoryImpl) FindByUkuran(ctx context.Context, ukuran string) (*models.Baju, error) {
	var baju models.Baju
	if err := r.DB.WithContext(ctx).Where("ukuran = ?", ukuran).First(&baju).Error; err != nil {
		return nil, err
	}
	return &baju, nil
}

func (r *PostRepositoryImpl) TambahStok(ctx context.Context, id uint, jumlah int) (*models.Baju, error) {
	var baju models.Baju
	if err := r.DB.WithContext(ctx).First(&baju, id).Error; err != nil {
		return nil, err
	}
	baju.Stok += jumlah
	if err := r.DB.WithContext(ctx).Save(&baju).Error; err != nil {
		return nil, err
	}
	return &baju, nil
}

func (r *PostRepositoryImpl) KurangStok(ctx context.Context, id uint, jumlah int) (*models.Baju, error) {
	var baju models.Baju
	if err := r.DB.WithContext(ctx).First(&baju, id).Error; err != nil {
		return nil, err
	}
	if baju.Stok < jumlah {
		return nil, errors.New("insufficient stock")
	}
	baju.Stok -= jumlah
	if err := r.DB.WithContext(ctx).Save(&baju).Error; err != nil {
		return nil, err
	}
	return &baju, nil
}
