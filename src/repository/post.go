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

func (r *PostRepositoryImpl) FindAll(ctx context.Context) ([]*models.Baju, error) {
	var bajus []*models.Baju
	// Set value directly in the query
	if err := r.DB.WithContext(ctx).Where("stok < ?", 5).Find(&bajus).Error; err != nil {
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

// FindLowStock menampilkan baju dengan stok kurang dari 5
func (r *PostRepositoryImpl) FindLowStock(ctx context.Context) ([]*models.Baju, error) {
	var bajus []*models.Baju
	if err := r.DB.WithContext(ctx).Where("stok < ?", 5).Find(&bajus).Error; err != nil {
		return nil, err
	}
	return bajus, nil
}

// Update memperbarui Baju yang ada
func (r *PostRepositoryImpl) Update(ctx context.Context, baju *models.Baju) (*models.Baju, error) {
	var existingBaju models.Baju
	if err := r.DB.WithContext(ctx).First(&existingBaju, baju.Id).Error; err != nil {
		return nil, err
	}

	if err := r.DB.WithContext(ctx).Model(&existingBaju).Updates(baju).Error; err != nil {
		return nil, err
	}
	return &existingBaju, nil
}

func (r *PostRepositoryImpl) TambahStock(ctx context.Context, id uint, jumlah int) (*models.Baju, error) {
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

func (r *PostRepositoryImpl) KurangStock(ctx context.Context, id uint, jumlah int) (*models.Baju, error) {
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

// Delete menghapus baju berdasarkan id
func (r *PostRepositoryImpl) Delete(ctx context.Context, id uint) error {
	if err := r.DB.WithContext(ctx).Delete(&models.Baju{}, id).Error; err != nil {
		return err
	}
	return nil
}
