package repository

import (
	"context"
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
	var list_baju []*models.Baju
	// Set value directly in the query
	if err := r.DB.WithContext(ctx).Find(&list_baju).Error; err != nil {
		return nil, err
	}
	return list_baju, nil
}

func (r *PostRepositoryImpl) FindByWarna(ctx context.Context, warna string) ([]*models.Baju, error) {
	var list_baju []*models.Baju
	if err := r.DB.WithContext(ctx).Where("warna = ?", warna).Find(&list_baju).Error; err != nil {
		return nil, err
	}
	return list_baju, nil
}

func (r *PostRepositoryImpl) FindByUkuran(ctx context.Context, ukuran string) ([]*models.Baju, error) {
	var list_baju []*models.Baju
	if err := r.DB.WithContext(ctx).Where("ukuran = ?", ukuran).Find(&list_baju).Error; err != nil {
		return nil, err
	}
	return list_baju, nil
}

// FindLowStock menampilkan baju dengan stok kurang dari 5
func (r *PostRepositoryImpl) FindLowStock(ctx context.Context) ([]*models.Baju, error) {
	var list_baju []*models.Baju
	if err := r.DB.WithContext(ctx).Where("stok < ?", 5).Find(&list_baju).Error; err != nil {
		return nil, err
	}
	return list_baju, nil
}

// Update memperbarui Baju yang ada
func (r *PostRepositoryImpl) Update(ctx context.Context, baju *models.Baju) (*models.Baju, error) {
	var existingBaju models.Baju
	if err := r.DB.WithContext(ctx).First(&existingBaju, baju.Id).Error; err != nil {
		return nil, err
	}

	// Mengatur stok menjadi 0 jika stok pada baju yang diterima adalah 0
	if baju.Stok == 0 {
		baju.Stok = 0
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

	// Pastikan stok tidak menjadi negatif
	if baju.Stok < jumlah {
		jumlah = baju.Stok
	}

	baju.Stok -= jumlah

	if err := r.DB.WithContext(ctx).Save(&baju).Error; err != nil {
		return nil, err
	}
	return &baju, nil
}

// FindOutOfStock mencari semua baju dengan stok 0
func (r *PostRepositoryImpl) FindOutOfStock(ctx context.Context) ([]*models.Baju, error) {
	var list_baju []*models.Baju
	if err := r.DB.WithContext(ctx).Where("stok = 0").Find(&list_baju).Error; err != nil {
		return nil, err
	}
	return list_baju, nil
}

// Delete menghapus baju berdasarkan id
func (r *PostRepositoryImpl) Delete(ctx context.Context, id uint) error {
	if err := r.DB.WithContext(ctx).Delete(&models.Baju{}, id).Error; err != nil {
		return err
	}
	return nil
}
