package main

import (
	"fmt"
)

const maxFilm = 100

type Film struct {
	Judul  string
	Genre  string
	Rating float64
	Status string // "Ditonton" atau "Belum"
}

var daftarFilm [maxFilm]Film
var jumlahFilm int

func main() {
	for {
		tampilkanMenu()
		pilihMenu()
	}
}

func tampilkanMenu() {
	fmt.Println("\n=== APLIKASI MANAJEMEN FILM ===")
	fmt.Println("1. Tambah Film")
	fmt.Println("2. Ubah Film")
	fmt.Println("3. Hapus Film")
	fmt.Println("4. Tampilkan Semua Film")
	fmt.Println("5. Cari Film by Judul")
	fmt.Println("6. Cari Film by Genre")
	fmt.Println("7. Cari Film by Rating")
	fmt.Println("8. Urutkan Film")
	fmt.Println("0. Keluar")
	fmt.Print("Pilihan: ")
}

func pilihMenu() {
	var pilihan int
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		tambahFilm()
	} else if pilihan == 2 {
		ubahFilm()
	} else if pilihan == 3 {
		hapusFilm()
	} else if pilihan == 4 {
		tampilkanFilm()
	} else if pilihan == 5 {
		cariByJudul()
	} else if pilihan == 6 {
		cariByGenre()
	} else if pilihan == 7 {
		cariByRating()
	} else if pilihan == 8 {
		urutkanFilm()
	} else if pilihan == 0 {
		fmt.Println("Terima kasih!")
		return
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

func tambahFilm() {
	if jumlahFilm >= maxFilm {
		fmt.Println("Maaf, kapasitas film penuh!")
		return
	}

	fmt.Print("\nMasukkan data film:\n")
	fmt.Print("Judul: ")
	fmt.Scan(&daftarFilm[jumlahFilm].Judul)
	fmt.Print("Genre: ")
	fmt.Scan(&daftarFilm[jumlahFilm].Genre)
	fmt.Print("Rating (0-10): ")
	fmt.Scan(&daftarFilm[jumlahFilm].Rating)
	fmt.Print("Status (Ditonton/Belum): ")
	fmt.Scan(&daftarFilm[jumlahFilm].Status)

	jumlahFilm++
	fmt.Println("Film berhasil ditambahkan!")
}

func ubahFilm() {
	var judul string
	fmt.Print("\nMasukkan judul film yang akan diubah: ")
	fmt.Scan(&judul)

	index := cariFilmByJudul(judul)
	if index == -1 {
		fmt.Println("Film tidak ditemukan!")
		return
	}

	fmt.Println("\nData saat ini:")
	fmt.Printf("Judul: %s\nGenre: %s\nRating: %.1f\nStatus: %s\n",
		daftarFilm[index].Judul, daftarFilm[index].Genre,
		daftarFilm[index].Rating, daftarFilm[index].Status)

	fmt.Print("\nMasukkan data baru:\n")
	fmt.Print("Judul: ")
	fmt.Scan(&daftarFilm[index].Judul)
	fmt.Print("Genre: ")
	fmt.Scan(&daftarFilm[index].Genre)
	fmt.Print("Rating (0-10): ")
	fmt.Scan(&daftarFilm[index].Rating)
	fmt.Print("Status (Ditonton/Belum): ")
	fmt.Scan(&daftarFilm[index].Status)

	fmt.Println("Film berhasil diubah!")
}

func hapusFilm() {
	var judul string
	fmt.Print("\nMasukkan judul film yang akan dihapus: ")
	fmt.Scan(&judul)

	index := cariFilmByJudul(judul)
	if index == -1 {
		fmt.Println("Film tidak ditemukan!")
		return
	}

	for i := index; i < jumlahFilm-1; i++ {
		daftarFilm[i] = daftarFilm[i+1]
	}
	jumlahFilm--

	fmt.Println("Film berhasil dihapus!")
}

func tampilkanFilm() {
	if jumlahFilm == 0 {
		fmt.Println("\nBelum ada data film!")
		return
	}

	fmt.Println("\nDaftar Film:")
	for i := 0; i < jumlahFilm; i++ {
		fmt.Printf("%d. %s (%s) - Rating: %.1f - Status: %s\n",
			i+1, daftarFilm[i].Judul, daftarFilm[i].Genre,
			daftarFilm[i].Rating, daftarFilm[i].Status)
	}
}

func cariFilmByJudul(judul string) int {
	for i := 0; i < jumlahFilm; i++ {
		if daftarFilm[i].Judul == judul {
			return i
		}
	}
	return -1
}

func cariByJudul() {
	var judul string
	fmt.Print("\nMasukkan judul yang dicari: ")
	fmt.Scan(&judul)

	index := cariFilmByJudul(judul)
	if index != -1 {
		fmt.Printf("\nDitemukan:\nJudul: %s\nGenre: %s\nRating: %.1f\nStatus: %s\n",
			daftarFilm[index].Judul, daftarFilm[index].Genre,
			daftarFilm[index].Rating, daftarFilm[index].Status)
	} else {
		fmt.Println("Film tidak ditemukan!")
	}
}

func cariByGenre() {
	var genre string
	fmt.Print("\nMasukkan genre yang dicari: ")
	fmt.Scan(&genre)

	fmt.Println("\nHasil Pencarian:")
	found := false
	for i := 0; i < jumlahFilm; i++ {
		if daftarFilm[i].Genre == genre {
			fmt.Printf("- %s (Rating: %.1f)\n", daftarFilm[i].Judul, daftarFilm[i].Rating)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada film dengan genre tersebut")
	}
}

func cariByRating() {
	var minRating float64
	fmt.Print("\nMasukkan rating minimum: ")
	fmt.Scan(&minRating)

	// Step 1: Urutkan berdasarkan rating ascending
	for i := 0; i < jumlahFilm-1; i++ {
		for j := 0; j < jumlahFilm-i-1; j++ {
			if daftarFilm[j].Rating > daftarFilm[j+1].Rating {
				daftarFilm[j], daftarFilm[j+1] = daftarFilm[j+1], daftarFilm[j]
			}
		}
	}

	// Step 2: Binary search untuk cari batas pertama rating >= minRating
	low := 0
	high := jumlahFilm - 1
	foundIndex := -1

	for low <= high {
		mid := (low + high) / 2
		if daftarFilm[mid].Rating >= minRating {
			foundIndex = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	// Step 3: Tampilkan hasil
	if foundIndex == -1 {
		fmt.Println("Tidak ada film dengan rating di atas", minRating)
	} else {
		fmt.Println("\nFilm dengan rating di atas", minRating, ":")
		for i := foundIndex; i < jumlahFilm; i++ {
			if daftarFilm[i].Rating >= minRating {
				fmt.Printf("- %s (%s) - Rating: %.1f\n",
					daftarFilm[i].Judul, daftarFilm[i].Genre, daftarFilm[i].Rating)
			}
		}
	}
}

func urutkanFilm() {
	var kriteria, urutan string
	fmt.Print("\nUrutkan berdasarkan (judul/rating): ")
	fmt.Scan(&kriteria)
	fmt.Print("Urutan (asc/desc): ")
	fmt.Scan(&urutan)

	for i := 0; i < jumlahFilm-1; i++ {
		for j := 0; j < jumlahFilm-i-1; j++ {
			shouldSwap := false

			if kriteria == "judul" {
				if urutan == "asc" {
					shouldSwap = daftarFilm[j].Judul > daftarFilm[j+1].Judul
				} else {
					shouldSwap = daftarFilm[j].Judul < daftarFilm[j+1].Judul
				}
			} else if kriteria == "rating" {
				if urutan == "asc" {
					shouldSwap = daftarFilm[j].Rating > daftarFilm[j+1].Rating
				} else {
					shouldSwap = daftarFilm[j].Rating < daftarFilm[j+1].Rating
				}
			}

			if shouldSwap {
				temp := daftarFilm[j]
				daftarFilm[j] = daftarFilm[j+1]
				daftarFilm[j+1] = temp
			}
		}
	}

	fmt.Println("Daftar film telah diurutkan!")
	tampilkanFilm()
}