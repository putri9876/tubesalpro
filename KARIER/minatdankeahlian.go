package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Karier struct {
	Nama       string
	Kategori   string
	Keahlian   []string
	Gaji       int
	Kecocokan  float64
}

var daftarKarier = []Karier{
	{"Ilmuwan Data", "Teknologi", []string{"Python", "Statistik"}, 15000000, 0},
	{"Perancang UX", "Desain", []string{"Desain UI", "Figma"}, 12000000, 0},
	{"Akuntan", "Keuangan", []string{"Akuntansi", "Excel"}, 10000000, 0},
	{"Pengembang Web", "Teknologi", []string{"HTML", "CSS", "JavaScript"}, 13000000, 0},
}

var keahlianPengguna []string

func masukan(teks string) string {
	fmt.Print(teks)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func tambahKeahlian() {
	keahlian := masukan("Masukkan keahlian: ")
	for _, k := range keahlianPengguna {
		if strings.EqualFold(k, keahlian) {
			fmt.Println("! Keahlian sudah ada.")
			return
		}
	}
	keahlianPengguna = append(keahlianPengguna, keahlian)
	fmt.Println("✓ Keahlian ditambahkan.")
}

func tampilkanKeahlian() {
	fmt.Println("== Daftar Keahlian ==")
	for i, k := range keahlianPengguna {
		fmt.Printf("%d. %s\n", i+1, k)
	}
}

func ubahKeahlian() {
	tampilkanKeahlian()
	indeks, _ := strconv.Atoi(masukan("Pilih nomor keahlian yang ingin diubah: "))
	if indeks > 0 && indeks <= len(keahlianPengguna) {
		keahlianPengguna[indeks-1] = masukan("Masukkan keahlian baru: ")
		fmt.Println("✓ Keahlian diubah.")
	} else {
		fmt.Println("! Nomor tidak valid.")
	}
}

func hapusKeahlian() {
	tampilkanKeahlian()
	indeks, _ := strconv.Atoi(masukan("Pilih nomor keahlian yang ingin dihapus: "))
	if indeks > 0 && indeks <= len(keahlianPengguna) {
		keahlianPengguna = append(keahlianPengguna[:indeks-1], keahlianPengguna[indeks:]...)
		fmt.Println("✓ Keahlian dihapus.")
	} else {
		fmt.Println("! Nomor tidak valid.")
	}
}

func hitungKecocokan() []Karier {
	var hasil []Karier
	for _, k := range daftarKarier {
		jumlahCocok := 0
		for _, keahlian := range k.Keahlian {
			for _, punya := range keahlianPengguna {
				if strings.EqualFold(keahlian, punya) {
					jumlahCocok++
				}
			}
		}
		persen := float64(jumlahCocok) / float64(len(k.Keahlian)) * 100
		k.Kecocokan = persen
		hasil = append(hasil, k)
	}
	return hasil
}

func tampilkanRekomendasi() {
	hasil := hitungKecocokan()
	fmt.Println("Urut berdasarkan: 1. Kecocokan  2. Gaji")
	pilihan := masukan("Pilih: ")
	if pilihan == "1" {
		urutKecocokan(hasil)
	} else if pilihan == "2" {
		urutGaji(hasil)
	}
	for _, k := range hasil {
		fmt.Printf("- %s (%s) → Kecocokan: %.2f%%, Gaji: Rp%d\n", k.Nama, k.Kategori, k.Kecocokan, k.Gaji)
	}
}

func urutKecocokan(data []Karier) {
	for i := 0; i < len(data)-1; i++ {
		maks := i
		for j := i + 1; j < len(data); j++ {
			if data[j].Kecocokan > data[maks].Kecocokan {
				maks = j
			}
		}
		data[i], data[maks] = data[maks], data[i]
	}
}

func urutGaji(data []Karier) {
	for i := 1; i < len(data); i++ {
		kunci := data[i]
		j := i - 1
		for j >= 0 && data[j].Gaji < kunci.Gaji {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = kunci
	}
}

func cariNama(nama string) *Karier {
	for _, k := range daftarKarier {
		if strings.EqualFold(k.Nama, nama) {
			return &k
		}
	}
	return nil
}

func cariKategori(kategori string) *Karier {
	sort.Slice(daftarKarier, func(i, j int) bool {
		return daftarKarier[i].Kategori < daftarKarier[j].Kategori
	})
	awal, akhir := 0, len(daftarKarier)-1
	for awal <= akhir {
		tengah := (awal + akhir) / 2
		if strings.EqualFold(daftarKarier[tengah].Kategori, kategori) {
			return &daftarKarier[tengah]
		} else if daftarKarier[tengah].Kategori < kategori {
			awal = tengah + 1
		} else {
			akhir = tengah - 1
		}
	}
	return nil
}

func tampilkanStatistik() {
	fmt.Println("== Statistik Kecocokan ==")
	rekom := hitungKecocokan()
	for _, k := range rekom {
		fmt.Printf("%s → %.2f%% cocok\n", k.Nama, k.Kecocokan)
	}
}
func main() {
	for {
		fmt.Println("\n== MENU UTAMA ==")
		fmt.Println("1. Tambah Keahlian")
		fmt.Println("2. Ubah Keahlian")
		fmt.Println("3. Hapus Keahlian")
		fmt.Println("4. Lihat Rekomendasi Karier")
		fmt.Println("5. Cari Karier (Nama/Kategori)")
		fmt.Println("6. Statistik Kecocokan")
		fmt.Println("0. Keluar")

		pilihan := masukan("Pilih menu: ")
		switch pilihan {
		case "1":
			tambahKeahlian()
		case "2":
			ubahKeahlian()
		case "3":
			hapusKeahlian()
		case "4":
			tampilkanRekomendasi()
		case "5":
			mode := masukan("Cari berdasarkan (1. Nama / 2. Kategori): ")
			if mode == "1" {
				nama := masukan("Masukkan nama karier: ")
				hasil := cariNama(nama)
				if hasil != nil {
					fmt.Printf("✓ Ditemukan: %s (%s)\n", hasil.Nama, hasil.Kategori)
				} else {
					fmt.Println("! Karier tidak ditemukan.")
				}
			} else if mode == "2" {
				kat := masukan("Masukkan kategori industri: ")
				hasil := cariKategori(kat)
				if hasil != nil {
					fmt.Printf("✓ Ditemukan: %s (%s)\n", hasil.Nama, hasil.Kategori)
				} else {
					fmt.Println("! Kategori tidak ditemukan.")
				}
			}
		case "6":
			tampilkanStatistik()
		case "0":
			fmt.Println("Terima kasih telah menggunakan aplikasi kami.")
			return
		default:
			fmt.Println("! Menu tidak valid.")
		}
	}
}
