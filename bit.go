package main

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	b, err := os.Open("./files/100.atr")
	check(err)
	c := make([]byte, 600)
	b.Read(c)
	fmt.Printf("%b", c)
}

func Read212(path string) []int {
	// Setting up the containers ...
	output := make([]int, 0) // No simple way to find the size ...

	// According to WFDB docs in a 212 format two 12-bit data is stored in a 3 byte size
	bytes := make([]byte, 3)
	chars := make([]string, 3)

	file, err := os.Open("./files/100.dat")
	check(err)
	defer file.Close() // Remember to always close the files ...

	for {
		// I know it's not the usual way ... but we needed the little one
		err = binary.Read(file, binary.LittleEndian, bytes)
		if err == io.EOF {
			break
		}
		check(err)

		// Changing bits to bits ... but as strings
		for a := range bytes {
			g := fmt.Sprintf("%b", bytes[a])
			for h := len(g); h < 8; h++ {
				g = "0" + g
			}
			chars[a] = g
		}

		// It's Binary ... so Base 2 and it's 12-bit data so ... 12 bits
		c1, err := strconv.ParseInt(chars[1][:4]+chars[0], 2, 12)
		check(err)
		c2, err := strconv.ParseInt(chars[1][4:]+chars[2], 2, 12)
		check(err)

		output = append(output, int(c1), int(c2))
	}
	return output
}

// It's here ... but please don't try it ... it's horrible ...
// Try a webframework data visualization instead ... MUCH cleaner
// Try Svelte ...
// You know what? ... I'll make one later
func Visualize(data []int) {
	// Normalizing ...
	min, max := data[0], data[0]

	for _, a := range data {
		if a > max {
			max = a
		} else if a < min {
			min = a
		}
	}

	img := image.NewRGBA(image.Rect(0, min, len(data), max))
	for a, b := range data {
		img.Set(a, b, color.White)
	}

	file, err := os.Create("hello.png")
	check(err)
	if err = png.Encode(file, img); err != nil {
		check(err)
	}
}

// Always check for errors ... SO important ... I'm serious ...
func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
