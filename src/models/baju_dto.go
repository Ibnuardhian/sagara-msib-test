// models/baju_dto.go
package models

// CreateBajuRequest digunakan untuk membuat Baju baru
type BajuRequest struct {
	Warna  string  `json:"warna"`
	Ukuran string  `json:"ukuran"`
	Harga  float64 `json:"harga"`
	Stok   int     `json:"stok"`
}

// BajuResponse digunakan untuk mengembalikan data Baju
type BajuResponse struct {
	Id     uint    `json:"id"`
	Warna  string  `json:"warna"`
	Ukuran string  `json:"ukuran"`
	Harga  float64 `json:"harga"`
	Stok   int     `json:"stok"`
}

// FindAllBajuResponse digunakan untuk mengembalikan daftar Baju
type FindAllBajuResponse struct {
	Bajus []*BajuResponse `json:"bajus"`
}

// TambahStokReq digunakan untuk menambah stok
type TambahStokReq struct {
	Id     uint `json:"id"`
	Jumlah int  `json:"jumlah"`
}

// KurangStokReq digunakan untuk mengurangi stok
type KurangStokReq struct {
	Id     uint `json:"id"`
	Jumlah int  `json:"jumlah"`
}
