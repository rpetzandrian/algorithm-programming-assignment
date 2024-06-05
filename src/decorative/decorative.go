package decorative

import (
	"github.com/fatih/color"
)

func PrintLine() {
	color.New(color.FgHiWhite).Println("╔══════════════════════════════════════════════════════════╗")
}

// Fungsi untuk mencetak garis bawah dengan karakter khusus
func PrintBottomLine() {
	color.New(color.FgHiWhite).Println("╚══════════════════════════════════════════════════════════╝")
}

// Fungsi untuk mencetak judul dengan warna dan gaya, dan karakter khusus
func PrintTitle(title string) {
	color.New(color.FgHiCyan, color.Bold).Printf("║ %-56s ║\n", title)
}

// Fungsi untuk mencetak subtitle dengan warna dan gaya, dan karakter khusus
func PrintSubtitle(subtitle string) {
	color.New(color.FgHiGreen).Printf("║ %-56s ║\n", subtitle)
}

// Fungsi untuk mencetak nama pembuat dengan warna dan gaya, dan karakter khusus
func PrintAuthor(author string) {
	color.New(color.FgHiMagenta).Printf("║ %-56s ║\n", author)
}

// Fungsi untuk mencetak baris kosong dengan karakter khusus
func PrintEmptyLine() {
	color.New(color.FgHiWhite).Println("║                                                          ║")
}

// Fungsi untuk mencetak menu dengan nomor urut dan karakter khusus
func PrintMenu(number int, menu string) {
	color.New(color.FgHiYellow).Printf("║ %d. %-53s ║\n", number, menu)
}

// Fungsi untuk mencetak dekorasi tambahan
func PrintDecorativeLine() {
	color.New(color.FgHiBlue).Println("╠══════════════════════════════════════════════════════════╣")
}

// Fungsi untuk mencetak kalimat petunjuk

func PrintInstruction(instruction string) {
	color.New(color.FgHiWhite).Printf("║ %-56s ║\n", instruction)
}

// Fungsi untuk mencetak kalimat Alert
func PrintAlert(alert string) {
	color.New(color.FgHiRed, color.Bold).Printf("%s \n", alert)
}

// Fungsi untuk mencetak kalimat Info
func PrintInfo(info string) {
	color.New(color.FgHiGreen).Printf("%s \n", info)
}
