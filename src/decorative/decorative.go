package decorative

import (
	"email-app/src/util"
	"fmt"
	"strings"

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

func PrintNothing() {
	fmt.Print()
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

func HeaderTemplate() {
	// Mencetak tampilan dengan dekorasi dan informasi proyek yang lebih menarik
	PrintLine()
	PrintTitle(" Alpro Assignment ")
	PrintDecorativeLine()
	PrintSubtitle(" EMAIL APP ")
	PrintEmptyLine()
	PrintSubtitle(" Created by: ")
	PrintAuthor(" Rico x Daffa ")
	PrintBottomLine()
}

const maxLineLength = 49 // Panjang maksimum garis

// Fungsi untuk membungkus teks jika lebih panjang dari maxLineLength
func wrapText(text string, maxLineLength int) []string {
	var wrappedLines []string
	words := strings.Fields(text)
	currentLine := ""

	for _, word := range words {
		if len(currentLine)+len(word)+1 > maxLineLength {
			wrappedLines = append(wrappedLines, currentLine)
			currentLine = word
		} else {
			if currentLine == "" {
				currentLine = word
			} else {
				currentLine += " " + word
			}
		}
	}

	if currentLine != "" {
		wrappedLines = append(wrappedLines, currentLine)
	}

	return wrappedLines
}

func PrintStatus(status string, text string) {
	var colorStatus color.Attribute
	var symbolStatus string
	if status == util.PRINT_STATUS_SUCCESS {
		colorStatus = color.FgHiGreen
		symbolStatus = "✔"
	} else if status == util.PRINT_STATUS_ERROR {
		colorStatus = color.FgHiRed
		symbolStatus = "!"
	}
	lines := wrapText(text, maxLineLength)
	boxWidth := maxLineLength + 9 // Lebar kotak termasuk simbol tambahan
	border := "╔" + strings.Repeat("═", boxWidth) + "╗"

	color.New(colorStatus).Println(border)

	for _, line := range lines {
		if len(line) > maxLineLength {
			line = line[:maxLineLength]
		}
		padding := (maxLineLength - len(line)) / 2
		leftPadding := strings.Repeat(" ", padding)
		rightPadding := strings.Repeat(" ", maxLineLength-len(line)-padding)
		color.New(colorStatus).Printf("║ %s ║%s%s%s ║ %s ║\n", symbolStatus, leftPadding, line, rightPadding, symbolStatus)
	}

	color.New(colorStatus).Println("╚" + strings.Repeat("═", boxWidth) + "╝")
}

func ResetPrintStatus(printStatus *string, printText *string) {
	*printStatus = util.PRINT_STATUS_NOTHING
	*printText = ""
}

// Fungsi untuk mencetak kalimat Alert
func PrintAlert(alert string) {
	color.New(color.FgHiRed, color.Bold).Printf("%s \n", alert)
}

// Fungsi untuk mencetak kalimat Info
func PrintInfo(info string) {
	color.New(color.FgHiGreen).Printf("%s \n", info)
}

func infoPage(info string) {
	PrintLine()
	PrintSubtitle(info)
	PrintBottomLine()
}
