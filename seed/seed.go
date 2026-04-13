package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yaseumine/coba-golang/config"
	"github.com/yaseumine/coba-golang/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env tidak ditemukan, lanjut memakai environment yang tersedia")
	}

	config.InitDatabase()

	products := []models.Product{
		{
			Name:        "Nasi Goreng Spesial",
			Price:       25000,
			Category:    "Makanan",
			Stock:       50,
			Description: "Nasi goreng dengan telur dan ayam",
			ImageURL:    "https://picsum.photos/400",
		},
		{
			Name:        "Sate Ayam 10 Tusuk",
			Price:       20000,
			Category:    "Makanan",
			Stock:       100,
			Description: "Sate ayam dengan bumbu kacang",
			ImageURL:    "https://picsum.photos/401",
		},
		{
			Name:        "Es Teh Manis",
			Price:       8000,
			Category:    "Minuman",
			Stock:       200,
			Description: "Es teh manis segar",
			ImageURL:    "https://picsum.photos/402",
		},
		{
			Name:        "Kopi Susu",
			Price:       15000,
			Category:    "Minuman",
			Stock:       150,
			Description: "Kopi susu kekinian",
			ImageURL:    "https://picsum.photos/403",
		},
		{
			Name:        "Ayam Bakar",
			Price:       30000,
			Category:    "Makanan",
			Stock:       30,
			Description: "Ayam bakar dengan sambal",
			ImageURL:    "https://picsum.photos/404",
		},
	}

	for _, p := range products {
		if err := config.DB.Create(&p).Error; err != nil {
			log.Fatalf("gagal insert product %q: %v", p.Name, err)
		}
	}

	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}