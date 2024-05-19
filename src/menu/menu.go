package menu

import "github.com/fatih/color"

func PrintStartMenu() {
	// Fungsi untuk mencetak garis dekoratif atas dan bawah dengan karakter khusus
	printLine := func() {
		color.New(color.FgHiWhite).Println("╔══════════════════════════════════════════════════════════╗")
	}

	// Fungsi untuk mencetak garis bawah dengan karakter khusus
	printBottomLine := func() {
		color.New(color.FgHiWhite).Println("╚══════════════════════════════════════════════════════════╝")
	}

	// Fungsi untuk mencetak judul dengan warna dan gaya, dan karakter khusus
	printTitle := func(title string) {
		color.New(color.FgHiCyan, color.Bold).Printf("║ %-56s ║\n", title)
	}

	// Fungsi untuk mencetak subtitle dengan warna dan gaya, dan karakter khusus
	printSubtitle := func(subtitle string) {
		color.New(color.FgHiGreen).Printf("║ %-56s ║\n", subtitle)
	}

	// Fungsi untuk mencetak nama pembuat dengan warna dan gaya, dan karakter khusus
	printAuthor := func(author string) {
		color.New(color.FgHiMagenta).Printf("║ %-56s ║\n", author)
	}

	// Fungsi untuk mencetak baris kosong dengan karakter khusus
	printEmptyLine := func() {
		color.New(color.FgHiWhite).Println("║                                                          ║")
	}

	// Fungsi untuk mencetak menu dengan nomor urut dan karakter khusus
	printMenu := func(number int, menu string) {
		color.New(color.FgHiYellow).Printf("║ %d. %-53s ║\n", number, menu)
	}

	// Fungsi untuk mencetak dekorasi tambahan
	printDecorativeLine := func() {
		color.New(color.FgHiBlue).Println("╠══════════════════════════════════════════════════════════╣")
	}

	// Fungsi untuk mencetak kalimat petunjuk
	printInstruction := func(instruction string) {
		color.New(color.FgHiWhite).Printf("║ %-56s ║\n", instruction)
	}

	// Mencetak tampilan dengan dekorasi dan informasi proyek yang lebih menarik
	printLine()
	printTitle(" Project Tugas Besar Alpro ")
	printDecorativeLine()
	printSubtitle(" Tentang: Email App ")
	printEmptyLine()
	printSubtitle(" Dibuat oleh: ")
	printAuthor(" Rico bersama Daffa ")
	printBottomLine()

	// Menambahkan menu user dan admin
	printLine()
	printTitle(" Menu Utama ")
	printDecorativeLine()
	printMenu(1, "User Menu")
	printMenu(2, "Admin Menu")
	printDecorativeLine()
	printInstruction(" Pilih menu dengan memasukkan angka ")
	printBottomLine()

	// Menambahkan pesan penutup dengan warna yang berbeda
	color.New(color.FgHiYellow, color.Bold).Println("\n🌟 Terima kasih telah menggunakan Email App! 🌟")
}
