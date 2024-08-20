package models

import "context"

type Baju struct {
	Id     uint    `json:"id"`
	Warna  string  `json:"warna"`
	Ukuran string  `json:"ukuran"`
	Harga  float64 `json:"harga"`
	Stok   int     `json:"stok"`
}

type BajuListResponse struct {
	Bajus []Baju `json:"bajus"`
}

type PostRepository interface {
	Create(ctx context.Context, baju *Baju) (*Baju, error)
	FindAll(ctx context.Context) ([]*Baju, error)
	FindByWarna(ctx context.Context, warna string) (*Baju, error)
	FindByUkuran(ctx context.Context, ukuran string) (*Baju, error)
	FindLowStock(ctx context.Context) ([]*Baju, error)
	Update(ctx context.Context, baju *Baju)(*Baju, error)
	TambahStock(ctx context.Context, id uint, jumlah int) (*Baju, error)
	KurangStock(ctx context.Context, id uint, jumlah int) (*Baju, error)
	FindOutOfStock(ctx context.Context) ([]*Baju, error)
	Delete(ctx context.Context, id uint) error
}
