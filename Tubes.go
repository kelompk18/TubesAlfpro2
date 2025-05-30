package main

import (
	"fmt"
	"math/rand"
	
)

const NMAX int = 1000
type hashCollected [NMAX]string
var history = [NMAX]string{}

var crypto = [62]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func main() {
	/*
		IS:  - Variabel `pilihan` belum berisi input dari user.
		     - Variabel `username` masih kosong, belum menyimpan nama pengguna.
		     - Variabel `loggedIn` bernilai false, artinya belum ada pengguna yang login.
		     - Array `HashColl` dan `history` masih kosong, belum ada data hasil mining.

		FS:  - Menampilkan menu utama kepada user.
		     - Menerima input pilihan dari user.
		     - Melakukan proses login dan menyimpan status login.
		     - Melakukan proses mining berdasarkan tingkat kesulitan (1–3).
		     - Menyimpan hasil mining ke dalam `history` (berhasil/gagal).
		     - Menampilkan riwayat hasil mining.
		     - Keluar dari program jika dipilih oleh pengguna.
	*/

	var pilihan int
	var HashColl hashCollected
	var i, j, N int = 0, 0, 0
	var isFound, isEmpty = false, false
	var targetHash, tempHash string
	var username string
	var loggedIn bool = false

	for {
		fmt.Println("\n=== APLIKASI MINING ===")
		fmt.Println("1. Login")
		fmt.Println("2. Mulai Mining")
		fmt.Println("3. Lihat History")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("Masukkan username: ")
			fmt.Scan(&username)
			loggedIn = true
			fmt.Println("Login berhasil!")
		case 2:
		if !loggedIn {
			fmt.Println("Silakan login terlebih dahulu!")
			continue
	}

		var tingkatKesulitan int
		fmt.Print("Masukkan tingkat kesulitan (1-5): ")
		fmt.Scan(&tingkatKesulitan)

		if tingkatKesulitan < 1 || tingkatKesulitan > 5 {
		fmt.Println("Tingkat kesulitan hanya boleh antara 1 sampai 5!")
		continue
	}

	// Tambahkan logika berdasarkan tingkat kesulitan, jika ada.
	fmt.Printf("Tingkat kesulitan %d dipilih.\n", tingkatKesulitan)


			//mining

			fmt.Scan(&N)
			fmt.Scan(&targetHash)

			for !isFound && i <= N {

				hashSederhana(HashColl, &tempHash)

				if tempHash != targetHash && i == N {
					fmt.Println("Waktu habis! Tidak berhasil menemukan hash.")
					for j < NMAX && !isEmpty {
						if len(history[j]) == 0 {
							isEmpty = true
						} else {
							j++
						}
					}
					history[j] = "Gagal"
				} else if tempHash == targetHash {
					fmt.Printf("Percobaan: %d | Hash: %s\n", i, tempHash)
					for j < NMAX && !isEmpty {
						if len(history[j]) == 0 {
							isEmpty = true
						} else {
							j++
						}
					}
					history[j] = "Berhasil"
					isFound = true
				} else {
					i++
				}
			}
		}
	}

}

func hashSederhana(A hashCollected, data *string) {
	/*
	   IS:  - Parameter `data` adalah pointer ke string kosong.
	        - Array `A` berisi hash-hash yang telah dihasilkan sebelumnya (untuk menghindari duplikasi).

	   FS:  - Menghasilkan hash acak sepanjang 8 karakter yang terdiri dari karakter alfanumerik.
	        - Hash yang dihasilkan tidak boleh sama dengan elemen manapun di array `A`.
	        - Nilai hash disimpan dalam variabel `*data`.
	*/
	var addKode string
	var i, j int = 0, 0

	for len(A) != 8 {
		for j = 0; j < 8; j++ {
			addKode = crypto[rand.Intn(62)]
			*data = *data + addKode
		}

		for i < NMAX && *data != "" {
			if A[i] == *data {
				*data = ""
			}
			i++
		}
	}
}

func menujuString(num int) string {
	/*
		IS:  Parameter `num` bertipe integer sudah berisi angka yang akan dikonversi.
		     Variabel `result` masih kosong, `n` menyimpan salinan dari `num`.

		FS:  Fungsi mengembalikan representasi string dari angka `num`.
		     Jika `num = 123`, maka hasilnya adalah "123".
	*/
	if num == 0 {
		return "0"
	}
	var result string = ""
	var n int = num
	var d int
	var digitChar string
	for n > 0 {
		d = n % 10
		digitChar = string('0' + byte(d))
		result = digitChar + result
		n = n / 10
	}
	return result
}
