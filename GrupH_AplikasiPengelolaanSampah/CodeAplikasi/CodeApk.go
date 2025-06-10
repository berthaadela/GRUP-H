package main

import "fmt"

const MAX_SAMPAH = 100
const MAX_METODE = 100

type Sampah struct {
	KodeSampah      string
	JenisSampah     string
	Jumlah          int
	DaurUlang       int
	MetodeDaurUlang string
}

type Metode struct {
	KodeSampah      string
	JenisSampah     string
	MetodeDaurUlang string
}

type DaftarSampah [MAX_SAMPAH]Sampah

var rekomMetode = [MAX_METODE]Metode{
	{"P01", "plastik", "Recycle"},
	{"P02", "kaca", "Recycle"},
	{"P03", "kertas", "Reuse"},
	{"P04", "logam", "Recycle"},
	{"P05", "karet", "Recycle"},
	{"P06", "kain", "Reuse"},
	{"P07", "elektronik", "Recycle"},
	{"P08", "kayu", "Reuse"},
	{"P09", "organik", "Kompos"},
	{"P10", "baterai", "Recycle"},
	{"P11", "botol_plastik", "Recycle"},
	{"P12", "kardus", "Reuse"},
	{"P13", "aluminium", "Recycle"},
	{"P14", "koran", "Reuse"},
	{"P15", "styrofoam", "Daur Ulang Khusus"},
	{"P16", "kaleng", "Recycle"},
	{"P17", "minyak_jelantah", "Pengolahan Ulang"},
	{"P18", "tekstil", "Reuse"},
	{"P19", "pvc", "Recycle"},
}

func CariInfoLengkapSampah(T [MAX_METODE]Metode, kodeAtauJenis string) (kode, jenis, metode string, ditemukan bool) {
	for i := 0; i < MAX_METODE; i++ {
		if T[i].KodeSampah != "" && T[i].KodeSampah == kodeAtauJenis {
			return T[i].KodeSampah, T[i].JenisSampah, T[i].MetodeDaurUlang, true
		}
	}
	for i := 0; i < MAX_METODE; i++ {
		if T[i].JenisSampah != "" && T[i].JenisSampah == kodeAtauJenis {
			return T[i].KodeSampah, T[i].JenisSampah, T[i].MetodeDaurUlang, true
		}
	}
	return "", "", "Tidak Ada", false
}

func CariSampahByKodeAtauJenis(T DaftarSampah, n int, query string) (bool, int) {
	for i := 0; i < n; i++ {
		if T[i].KodeSampah == query || T[i].JenisSampah == query {
			return true, i
		}
	}
	return false, -1
}

func TambahSampah(T *DaftarSampah, n *int, kode, jenis string, jumlah int, daurUlang int, metode string) {
	if *n >= MAX_SAMPAH {
		fmt.Println("Daftar sampah penuh. Tidak dapat menambah sampah baru.")
	} else {
		ditemukan, idx := CariSampahByKodeAtauJenis(*T, *n, kode)
		if ditemukan {
			T[idx].Jumlah += jumlah
			T[idx].DaurUlang += daurUlang
			fmt.Printf("Jumlah untuk sampah [Kode: %s, Jenis: %s] telah diperbarui.\n", kode, jenis)
		} else {
			T[*n] = Sampah{KodeSampah: kode, JenisSampah: jenis, Jumlah: jumlah, DaurUlang: daurUlang, MetodeDaurUlang: metode}
			(*n)++
			fmt.Printf("Sampah [Kode: %s, Jenis: %s] telah ditambahkan.\n", kode, jenis)
		}
	}
}

func TampilkanStatistik(T DaftarSampah, n int) {
	var totalSampah, totalDaurUlang int
	fmt.Println("--------------------------------------STATISTIK SAMPAH-----------------------------------------")
	fmt.Println("================================================================================================")
	fmt.Printf("| %-3s | %-10s | %-20s | %-10s | %-12s | %-20s |\n", "No", "Kode", "Jenis Sampah", "Jumlah", "Daur Ulang", "Metode Daur Ulang")
	fmt.Println("================================================================================================")
	for i := 0; i < n; i++ {
		totalSampah += T[i].Jumlah
		totalDaurUlang += T[i].DaurUlang
		fmt.Printf("| %-3d | %-10s | %-20s | %-10d | %-12d | %-20s |\n",
			i+1, T[i].KodeSampah, T[i].JenisSampah, T[i].Jumlah, T[i].DaurUlang, T[i].MetodeDaurUlang)
	}
	fmt.Println("================================================================================================")
	fmt.Printf("| %-39s | %-10d | %-12d | %-20s |\n", "TOTAL", totalSampah, totalDaurUlang, "-")
	fmt.Println("================================================================================================")
}

func TampilkanDaftarRekomMetode() {
	if len(rekomMetode) == 0 {
		fmt.Println("Daftar metode daur ulang kosong!")
	} else {
		fmt.Println("============================================================")
		fmt.Printf("| %-10s | %-20s | %-20s |\n", "Kode", "Jenis Sampah", "Metode Daur Ulang")
		fmt.Println("============================================================")
		for _, m := range rekomMetode {
			if m.JenisSampah != "" {
				fmt.Printf("| %-10s | %-20s | %-20s |\n", m.KodeSampah, m.JenisSampah, m.MetodeDaurUlang)
			}
		}
		fmt.Println("============================================================")
	}
}

func HapusSampah(T *DaftarSampah, n *int, kode string) {
	ditemukan, idx := CariSampahByKodeAtauJenis(*T, *n, kode)
	if !ditemukan {
		fmt.Printf("Sampah dengan kode atau jenis '%s' tidak ditemukan.\n", kode)
	} else {
		kodeDihapus := T[idx].KodeSampah
		jenisDihapus := T[idx].JenisSampah
		for i := idx; i < *n-1; i++ {
			T[i] = T[i+1]
		}
		(*n)--
		fmt.Printf("Sampah [Kode: %s, Jenis: %s] telah dihapus.\n", kodeDihapus, jenisDihapus)
	}
}

func UrutkanJumlahDesc(T *DaftarSampah, n int) {
	for i := 0; i < n-1; i++ {
		idx_max := i
		for j := i + 1; j < n; j++ {
			if T[j].Jumlah > T[idx_max].Jumlah {
				idx_max = j
			}
		}
		T[i], T[idx_max] = T[idx_max], T[i]
	}
}

func UrutkanJumlahAsc(T *DaftarSampah, n int) {
	for i := 0; i < n-1; i++ {
		idx_min := i
		for j := i + 1; j < n; j++ {
			if T[j].Jumlah < T[idx_min].Jumlah {
				idx_min = j
			}
		}
		T[i], T[idx_min] = T[idx_min], T[i]
	}
}

func main() {
	var daftar DaftarSampah
	var jumlahItemSaatIni int
	var pilihan int
	var kodeAtauJenis string
	var jumlahJenisSampah, jumlahDaurUlang int

	fmt.Printf("Aplikasi Pengelolaan Sampah dan Daur Ulang (Array Statis MAX: %d)\n\n", MAX_SAMPAH)

	for {
		fmt.Println("===========Daftar Sampah Saat Ini===========")
		if jumlahItemSaatIni == 0 {
			fmt.Printf("Belum ada sampah\n\n")
		} else {
			for i := 0; i < jumlahItemSaatIni; i++ {
				fmt.Printf("%d. Kode: %s, Jenis: %s, Jumlah: %d, Daur Ulang: %d, Metode: %s\n",
					i+1, daftar[i].KodeSampah, daftar[i].JenisSampah, daftar[i].Jumlah, daftar[i].DaurUlang, daftar[i].MetodeDaurUlang)
			}
			fmt.Println()
		}

		fmt.Println("Menu:")
		fmt.Println("[1] Tambah Sampah")
		fmt.Println("[2] Cari Sampah")
		fmt.Println("[3] Urutkan Sampah (Jumlah Terbanyak)")
		fmt.Println("[4] Urutkan Sampah (Jumlah Tersedikit)")
		fmt.Println("[5] Hapus Sampah")
		fmt.Println("[6] Tampilkan Statistik")
		fmt.Println("[7] Tampilkan Daftar Rekomendasi Metode")
		fmt.Println("[8] Edit Jumlah Sampah")
		fmt.Println("[0] Keluar")
		fmt.Print("Pilih (0-8): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
			case 0:
				fmt.Println("Terima kasih telah menggunakan aplikasi!")
				return
			case 1:
				fmt.Println("\n--- Tambah Sampah ---")
				fmt.Print("Masukkan kode atau jenis sampah (contoh: P19 atau pvc): ")
				fmt.Scanln(&kodeAtauJenis)

				kodeSampah, jenisSampah, metode, ditemukan := CariInfoLengkapSampah(rekomMetode, kodeAtauJenis)
				if ditemukan {
					fmt.Printf("Info Sampah -> Kode: %s, Jenis: %s, Metode Daur Ulang: %s\n", kodeSampah, jenisSampah, metode)
					fmt.Print("Masukkan jumlah sampah: ")
					fmt.Scanln(&jumlahJenisSampah)
					fmt.Print("Masukkan jumlah sampah yang didaur ulang: ")
					fmt.Scanln(&jumlahDaurUlang)
					if jumlahDaurUlang > jumlahJenisSampah {
						fmt.Println("Error: Jumlah daur ulang tidak boleh melebihi jumlah sampah. Gagal menambahkan.")
					} else {
						TambahSampah(&daftar, &jumlahItemSaatIni, kodeSampah, jenisSampah, jumlahJenisSampah, jumlahDaurUlang, metode)
					}
				} else {
					fmt.Printf("Sampah dengan kode atau jenis '%s' tidak ditemukan dalam daftar rekomendasi.\n\n", kodeAtauJenis)
				}
				fmt.Println()
			case 2:
				fmt.Println("\n--- Cari Jenis Sampah ---")
				if jumlahItemSaatIni == 0 {
					fmt.Printf("Belum ada sampah untuk dicari.\n\n")
				} else {
					var queryCari string
					fmt.Print("Masukkan kode atau jenis sampah (contoh: P19 atau pvc): ")
					fmt.Scanln(&queryCari)

					ditemukan, idx := CariSampahByKodeAtauJenis(daftar, jumlahItemSaatIni, queryCari)
					if ditemukan {
						if daftar[idx].KodeSampah == queryCari {
							fmt.Printf("Sampah dengan kode '%s' ditemukan pada urutan ke-%d.\n\n", queryCari, idx+1)
						} else {
							fmt.Printf("Sampah dengan jenis '%s' ditemukan pada urutan ke-%d.\n\n", queryCari, idx+1)
						}
					} else {
						fmt.Printf("Sampah dengan kode atau jenis '%s' tidak ditemukan dalam daftar.\n\n", queryCari)
					}
				}
			case 3:
				UrutkanJumlahDesc(&daftar, jumlahItemSaatIni)
				fmt.Printf("Data sampah berhasil diurutkan berdasarkan jumlah (terbanyak ke tersedikit).\n\n")
			case 4:
				UrutkanJumlahAsc(&daftar, jumlahItemSaatIni)
				fmt.Printf("Data sampah berhasil diurutkan berdasarkan jumlah (tersedikit ke terbanyak).\n\n")
			case 5:
				if jumlahItemSaatIni == 0 {
					fmt.Printf("Tidak ada sampah untuk dihapus.\n\n")
				} else {
					fmt.Print("Masukkan kode atau jenis sampah (contoh: P19 atau pvc): ")
					fmt.Scanln(&kodeAtauJenis)
					HapusSampah(&daftar, &jumlahItemSaatIni, kodeAtauJenis)
					fmt.Println()
				}
			case 6:
				if jumlahItemSaatIni == 0 {
					fmt.Printf("Tidak ada data untuk ditampilkan.\n\n")
				} else {
					TampilkanStatistik(daftar, jumlahItemSaatIni)
					fmt.Println()
				}
			case 7:
				fmt.Printf("\n----------- Daftar Rekomendasi Metode Daur Ulang -----------\n")
				TampilkanDaftarRekomMetode()
				fmt.Println()
			case 8:
				fmt.Printf("\n--- Edit Jumlah Sampah ---\n")
				if jumlahItemSaatIni == 0 {
					fmt.Printf("Belum ada data sampah untuk diedit.\n\n")
				} else {
					var kodeUntukEdit string
					fmt.Print("Masukkan kode atau jenis sampah (contoh: P19 atau pvc): ")
					fmt.Scanln(&kodeUntukEdit)

					ditemukan, idx := CariSampahByKodeAtauJenis(daftar, jumlahItemSaatIni, kodeUntukEdit)
					if !ditemukan {
						fmt.Printf("Sampah dengan kode atau jenis '%s' tidak ditemukan.\n\n", kodeUntukEdit)
					} else {
						sampahTarget := daftar[idx]
						fmt.Printf("Data saat ini -> Jenis: %s, Jumlah: %d, Daur Ulang: %d\n", sampahTarget.JenisSampah, sampahTarget.Jumlah, sampahTarget.DaurUlang)

						var jumlahBaru, daurUlangBaru int
						fmt.Print("Masukkan jumlah sampah baru: ")
						fmt.Scanln(&jumlahBaru)
						fmt.Print("Masukkan jumlah daur ulang baru: ")
						fmt.Scanln(&daurUlangBaru)

						if daurUlangBaru > jumlahBaru {
							fmt.Printf("Error: Jumlah daur ulang tidak boleh melebihi jumlah sampah.\n")
							fmt.Printf("Proses edit dibatalkan.\n\n")
						} else {
							daftar[idx].Jumlah = jumlahBaru
							daftar[idx].DaurUlang = daurUlangBaru
							fmt.Printf("Data sampah berhasil diperbarui.\n")
							fmt.Printf("Data baru -> Kode: %s, Jenis: %s, Jumlah: %d, Daur Ulang: %d\n\n", daftar[idx].KodeSampah, daftar[idx].JenisSampah, daftar[idx].Jumlah, daftar[idx].DaurUlang)
						}
					}
				}
			default:
				fmt.Printf("Pilihan tidak valid. Silakan coba lagi.\n\n")
			}
		}
}
