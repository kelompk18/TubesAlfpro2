package main

import (
	"fmt"
	"strings"
	"time"
)

const NMAX = 20

type Akun struct {
	Username string
	Password string
}

type Mining struct {
	ID       int
	Koin     string
	Nonce    int
	Waktu    time.Duration
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

func tampilkanHeader(judul string) {
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println(judul)
	fmt.Println(strings.Repeat("=", 40))
}

func buatAkun() {
	fmt.Println("\n🆕 Buat Akun Baru")
	fmt.Print("Username: ")
	fmt.Scanln(&akun.Username)
	fmt.Println("✅ Akun berhasil dibuat!\n")
}

func login() {
	var usn string
	fmt.Println("\n🔐 Login")
	fmt.Print("Username: ")
	fmt.Scanln(&usn)

	if usn == akun.Username  {
		isLoggedIn = true
		fmt.Println("✅ Login berhasil!\n")
	} else {
		fmt.Println("🚫 Username salah\n")
	}
}

func menuUtama() {
	var pilihan int
	tampilkanHeader("⛏️ Menu Utama CryptoMiner")
	fmt.Println("1. Mulai Mining")
	fmt.Println("2. Lihat Riwayat Mining")
	fmt.Println("3. Urutkan Riwayat (Waktu tercepat)")
	fmt.Println("4. Urutkan Riwayat (ID terkecil)")
	fmt.Println("5. Cari Riwayat Berdasarkan ID")
	fmt.Println("6. Logout")
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
		isLoggedIn = false
		fmt.Println("✅ Logout berhasil.\n")
	default:
		fmt.Println("🚫 Pilihan tidak valid.\n")
	}
}

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

	// Simulate mining attempts
	for i := 0; i < attempts; i++ {
		// Simulate some processing time for each attempt
		time.Sleep(100 * time.Millisecond) // Simulate work being done
	}

	elapsed := time.Since(start)

	riwayat[nData] = Mining{
		ID:       idCounter,
		Koin:     namaKoin,
		Nonce:    attempts, // Store the number of attempts as the nonce
		Waktu:    elapsed,
	}
	idCounter++
	nData++

	fmt.Printf("✅ Mining selesai! Percobaan: %d | Waktu: %v\n\n", attempts, elapsed)
}


func tampilkanRiwayat() {
	tampilkanHeader("📄 Riwayat Mining")

	if nData == 0 {
		fmt.Println("Belum ada riwayat mining.\n")
		return
	}

	fmt.Printf("%-5s %-10s %-10s %-15s \n", "ID", "Koin", "Nonce", "Waktu")
	fmt.Println(strings.Repeat("-", 40))
	for i := 0; i < nData; i++ {
		m := riwayat[i]
		fmt.Printf("%-5d %-10s %-10d %-15v \n", m.ID, m.Koin, m.Nonce, m.Waktu)
	}
	fmt.Println()
}

func sortByWaktu() {
   	var i int 
	for i = 0; i < nData-1; i++ {
        minIdx := i
        for j := i + 1; j < nData; j++ {
            // Compare Waktu and handle the late condition
            if riwayat[j].Waktu < riwayat[minIdx].Waktu {
                minIdx = j
            }
        }
        // Swap the elements
        riwayat[i], riwayat[minIdx] = riwayat[minIdx], riwayat[i]

        // Output if an entry is late
        if riwayat[i].Waktu > expectedWaktu {
            fmt.Println("Waktu terlambat untuk entri", i)
        }
    }
    fmt.Println("Riwayat diurutkan berdasarkan waktu tercepat.\n")
    tampilkanRiwayat()
}

