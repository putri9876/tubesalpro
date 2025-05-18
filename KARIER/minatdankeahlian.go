package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Career struct {
	Nama      string
	Kategori  string
	Keahlian  []string
	Gaji      int
	Kecocokan float64
}

var careers = []Career{
	{"Data Scientist", "Teknologi", []string{"Python", "Statistik"}, 15000000, 0},
	{"UX Designer", "Desain", []string{"Desain UI", "Figma"}, 12000000, 0},
	{"Akuntan", "Keuangan", []string{"Akuntansi", "Excel"}, 10000000, 0},
	{"Web Developer", "Teknologi", []string{"HTML", "CSS", "JavaScript"}, 13000000, 0},
}

var userSkills []string

func input(teks string) string {
	fmt.Print(teks)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func tambahKeahlian() {
	skill := input("Masukkan keahlian: ")
	for _, s := range userSkills {
		if strings.EqualFold(s, skill) {
			fmt.Println("! Keahlian sudah ada.")
			return
		}
	}
	userSkills = append(userSkills, skill)
	fmt.Println("✓ Keahlian ditambahkan.")
}

func lihatKeahlian() {
	fmt.Println("== Daftar Keahlian ==")
	for i, s := range userSkills {
		fmt.Printf("%d. %s\n", i+1, s)
	}
}

func editKeahlian() {
	lihatKeahlian()
	idx, _ := strconv.Atoi(input("Pilih nomor keahlian yang ingin diubah: "))
	if idx > 0 && idx <= len(userSkills) {
		userSkills[idx-1] = input("Masukkan keahlian baru: ")
		fmt.Println("✓ Keahlian diubah.")
	} else {
		fmt.Println("! Nomor tidak valid.")
	}
}

func hapusKeahlian() {
	lihatKeahlian()
	idx, _ := strconv.Atoi(input("Pilih nomor keahlian yang ingin dihapus: "))
	if idx > 0 && idx <= len(userSkills) {
		userSkills = append(userSkills[:idx-1], userSkills[idx:]...)
		fmt.Println("✓ Keahlian dihapus.")
	} else {
		fmt.Println("! Nomor tidak valid.")
	}
}

func hitungRekomendasi() []Career {
	var hasil []Career
	for _, c := range careers {
		kecocokan := 0
		for _, skill := range c.Keahlian {
			for _, uskill := range userSkills {
				if strings.EqualFold(skill, uskill) {
					kecocokan++
				}
			}
		}
		persen := float64(kecocokan) / float64(len(c.Keahlian)) * 100
		c.Kecocokan = persen
		hasil = append(hasil, c)
	}
	return hasil
}

func tampilkanRekomendasi() {
	hasil := hitungRekomendasi()
	fmt.Println("Urut berdasarkan: 1. Kecocokan  2. Gaji")
	pilihan := input("Pilih: ")
	if pilihan == "1" {
		selectionSortByKecocokan(hasil)
	} else if pilihan == "2" {
		insertionSortByGaji(hasil)
	}
	for _, c := range hasil {
		fmt.Printf("- %s (%s) → Kecocokan: %.2f%%, Gaji: Rp%d\n", c.Nama, c.Kategori, c.Kecocokan, c.Gaji)
	}
}

func selectionSortByKecocokan(data []Career) {
	for i := 0; i < len(data)-1; i++ {
		maxIdx := i
		for j := i + 1; j < len(data); j++ {
			if data[j].Kecocokan > data[maxIdx].Kecocokan {
				maxIdx = j
			}
		}
		data[i], data[maxIdx] = data[maxIdx], data[i]
	}
}

func insertionSortByGaji(data []Career) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].Gaji < key.Gaji {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func sequentialSearch(nama string) *Career {
	for _, c := range careers {
		if strings.EqualFold(c.Nama, nama) {
			return &c
		}
	}
	return nil
}
func binarySearchKategori(kategori string) *Career {
	sort.Slice(careers, func(i, j int) bool {
		return careers[i].Kategori < careers[j].Kategori
	})
	low, high := 0, len(careers)-1
	for low <= high {
		mid := (low + high) / 2
		if strings.EqualFold(careers[mid].Kategori, kategori) {
			return &careers[mid]
		} else if careers[mid].Kategori < kategori {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}
func statistik() {
	fmt.Println("== Statistik Kecocokan ==")
	rekom := hitungRekomendasi()
	for _, c := range rekom {
		fmt.Printf("%s → %.2f%% cocok\n", c.Nama, c.Kecocokan)
	}
}
func main() {
	for {
		fmt.Println("\n== MENU UTAMA ==")
		fmt.Println("1. Tambah Keahlian")
		fmt.Println("2. Edit Keahlian")
		fmt.Println("3. Hapus Keahlian")
		fmt.Println("4. Lihat Rekomendasi Karier")
		fmt.Println("5. Cari Karier (Nama/Kategori)")
		fmt.Println("6. Statistik Kecocokan")
		fmt.Println("0. Keluar")

		pilih := input("Pilih menu: ")
		switch pilih {
		case "1":
			tambahKeahlian()
		case "2":
			editKeahlian()
		case "3":
			hapusKeahlian()
		case "4":
			tampilkanRekomendasi()
		case "5":
			mode := input("Cari berdasarkan (1. Nama / 2. Kategori): ")
			if mode == "1" {
				nama := input("Masukkan nama karier: ")
				hasil := sequentialSearch(nama)
				if hasil != nil {
					fmt.Printf("✓ Ditemukan: %s (%s)\n", hasil.Nama, hasil.Kategori)
				} else {
					fmt.Println("! Karier tidak ditemukan.")
				}
			} else if mode == "2" {
				kat := input("Masukkan kategori industri: ")
				hasil := binarySearchKategori(kat)
				if hasil != nil {
					fmt.Printf("✓ Ditemukan: %s (%s)\n", hasil.Nama, hasil.Kategori)
				} else {
					fmt.Println("! Kategori tidak ditemukan.")
				}
			}
		case "6":
			statistik()
		case "0":
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("! Menu tidak valid.")
		}
	}
}
