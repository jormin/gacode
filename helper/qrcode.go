package helper

import (
	"fmt"
	"image"
	"image/color"
	"os"
)

// PrintQRCodeSize the size of QR code to print
const PrintQRCodeSize = 50

// ExportQRCodeSize the size of QR code to export
const ExportQRCodeSize = 256

// PrintQRCode print the QR code to terminal
func PrintQRCode(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}
	var points [PrintQRCodeSize][PrintQRCodeSize]int
	gray := image.NewGray(image.Rect(0, 0, PrintQRCodeSize, PrintQRCodeSize))
	for x := 0; x < PrintQRCodeSize; x++ {
		for y := 0; y < PrintQRCodeSize; y++ {
			r32, g32, b32, _ := img.At(x, y).RGBA()
			r, g, b := int(r32>>8), int(g32>>8), int(b32>>8)
			if (r+g+b)/3 > 180 {
				points[y][x] = 0
				gray.Set(x, y, color.Gray{Y: uint8(255)})
			} else {
				points[y][x] = 1
				gray.Set(x, y, color.Gray{Y: uint8(0)})
			}
		}
	}
	for x := 0; x < PrintQRCodeSize; x++ {
		for y := 0; y < PrintQRCodeSize; y++ {
			if points[x][y] == 1 {
				fmt.Print("\033[40;40m  \033[0m")
			} else {
				fmt.Print("\033[47;30m  \033[0m")
			}
		}
		fmt.Println()
	}
	return nil
}
