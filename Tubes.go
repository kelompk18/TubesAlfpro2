package main

import (
	"fmt"
	"strings"
	"time"
)

const NMAX = 20

type Akun struct {
	Username string
}

type Mining struct {
	Koin  string
	Nonce int
	Waktu time.Duration
}

var akun Akun
var isLoggedIn bool
var riwayat [NMAX]Mining
var nData int
var idCounter int = 1

func main() {
	var pilihan int

	for {
		if !isLoggedIn {
			tampilkanHeader("🟡 CryptoMiner - Simulasi Mining Sederhana")
			fmt.Println("1. Buat Akun")
			fmt.Println("2. Login")
			fmt.Println("3. Keluar")
			fmt.Print(">> Pilihan Anda: ")
			fmt.Scanln(&pilihan)

			// menggunakan fungsi switch case
			switch pilihan {
			case 1:
				buatAkun()
			case 2:
				login()
			case 3:
				fmt.Println("👋 Terima kasih telah mencoba CryptoMiner!")
				return
			default:
				fmt.Println("🚫 Pilihan tidak valid.\n")
			}
		} else {
			menuUtama()
		}
	}
}

// menggunakan strings.Repeat
func tampilkanHeader(judul string) {
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println(judul)
	fmt.Println(strings.Repeat("=", 40))
}

// buat akun
func buatAkun() {
	fmt.Println("\n🆕 Buat Akun Baru")
	fmt.Print("Username: ")
	fmt.Scanln(&akun.Username)
	fmt.Println("✅ Akun berhasil dibuat!\n")
}

// login akun
func login() {
	var usn string
	fmt.Println("\n🔐 Login")
	fmt.Print("Username: ")
	fmt.Scanln(&usn)

	if usn == akun.Username {
		isLoggedIn = true
		fmt.Println("✅ Login berhasil!\n")
	} else {
		fmt.Println("🚫 Username salah\n")
	}
}

// menu utama
func menuUtama() {
	var pilihan int
	tampilkanHeader("⛏️ Menu Utama CryptoMiner")
	fmt.Println("1. Mulai Mining")
	fmt.Println("2. Lihat Riwayat Mining")
	fmt.Println("3. Urutkan Riwayat (Waktu tercepat)")
	fmt.Println("4. Urutkan Riwayat (Nonce terkecil)")
	fmt.Println("5.Logout")
	fmt.Print(">> Pilihan Anda: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		mining()
	case 2:
		tampilkanRiwayat()
	case 3:
		sortByWaktu()
	case 4:
		sortByNonce()
	case 5:
		isLoggedIn = false
		fmt.Println("✅ Logout berhasil.\n")
	default:
		fmt.Println("🚫 Pilihan tidak valid.\n")
	}
}


// mining koin
func mining() {
	if nData >= NMAX {
		fmt.Println("🚫 Riwayat penuh, tidak bisa mining lagi.")
		return
	}

	var namaKoin string
	var attempts int
	fmt.Println("\n🔨 Mulai Mining")
	fmt.Print("Nama Koin: ")
	fmt.Scanln(&namaKoin)

	fmt.Print("Berapa banyak percobaan mining: ")
	fmt.Scanln(&attempts)

	start := time.Now()

	// menggunakan fungsi for-loop dan time.Sleep
	for i := 0; i < attempts; i++ {
		time.Sleep(100 * time.Millisecond)
	}

	elapsed := time.Since(start)

	riwayat[nData] = Mining{
		Koin:  namaKoin,
		Nonce: attempts,
		Waktu: elapsed,
	}
	idCounter++
	nData++

	fmt.Printf("✅ Mining selesai! Percobaan: %d | Waktu: %v\n\n", attempts, elapsed)
}

// menampilkan riwayat mining
func tampilkanRiwayat() {
	tampilkanHeader("📄 Riwayat Mining")

	if nData == 0 {
		fmt.Println("Belum ada riwayat mining.\n")
		return
	}

	fmt.Printf("%-10s %-10s %-15s \n", "Koin", "Nonce", "Waktu")
	fmt.Println(strings.Repeat("-", 40))
	for i := 0; i < nData; i++ {
		m := riwayat[i]
		fmt.Printf("%-10s %-10d %-15v \n", m.Koin, m.Nonce, m.Waktu)
	}
	fmt.Println()
}
// menggunakan insertion sort, mengurutkan berdasarkan nonce (jumlah percobaan)
func sortByNonce() {
	var i int 
	for i = 1; i < nData; i++ {
		temp := riwayat[i]
		j := i - 1
		for j >= 0 && riwayat[j].Nonce > temp.Nonce {
			riwayat[j+1] = riwayat[j]
			j--
		}
		riwayat[j+1] = temp
	}
	fmt.Println("Riwayat diurutkan berdasarkan jumlah nonce (percobaan).\n")
	tampilkanRiwayat()
}


// menggunakan fungsi selection sort
func sortByWaktu() {
	expectedWaktu := 2 * time.Second
	for i := 0; i < nData-1; i++ {
		minIdx := i
		for j := i + 1; j < nData; j++ {
			if riwayat[j].Waktu < riwayat[minIdx].Waktu {
				minIdx = j
			}
		}
		riwayat[i], riwayat[minIdx] = riwayat[minIdx], riwayat[i]

		if riwayat[i].Waktu > expectedWaktu {
			fmt.Println("⚠️  Waktu terlambat untuk entri", i)
		}
	}
	fmt.Println("Riwayat diurutkan berdasarkan waktu tercepat.\n")
	tampilkanRiwayat()
}
