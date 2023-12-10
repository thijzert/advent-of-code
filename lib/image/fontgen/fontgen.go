package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/thijzert/advent-of-code/lib/image"
)

func main() {
	for {
		d, _ := os.Getwd()
		if d == "" || d == "/" {
			log.Fatal("can't find repository root")
		}
		if f, err := os.Open(".git/HEAD"); err == nil {
			f.Close()
			break
		}
		os.Chdir("..")
	}

	fontfiles, err := filepath.Glob("lib/image/fontgen/data/*.txt")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("lib/image/ocrfonts.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "// Code generated by fontgen. DO NOT EDIT.\n")
	fmt.Fprintf(f, "//lint:file-ignore U1000 This file is automatically generated; ignore unused code checks\n")
	fmt.Fprintf(f, "package image\n\n")

	fmt.Fprintf(f, "var (\n")
	for _, fontfile := range fontfiles {
		basename := fontfile[23 : len(fontfile)-4]
		log.Printf("Found font '%s'", basename)
		fmt.Fprintf(f, "\tFont_%s  OCRFont\n", basename)
	}
	fmt.Fprintf(f, ")\n\nvar allFonts []OCRFont\n\n")

	fmt.Fprintf(f, "func init() {\n")
	for _, fontfile := range fontfiles {
		basename := fontfile[23 : len(fontfile)-4]
		fmt.Fprintf(f, "\tFont_%s = OCRFont{\n", basename)

		g, err := os.ReadFile(fontfile)
		if err != nil {
			log.Fatal(err)
		}
		chars := strings.Split(strings.TrimRight(string(g), "\n"), "\n\n")
		fmt.Fprintf(f, "\t\tcharacters: []ocrChar{\n")
		for _, char := range chars {
			r, i := utf8.DecodeRuneInString(char)
			img := image.ReadImage(strings.Split(char[i+1:], "\n"), charmask)
			imageSt := fmt.Sprintf("%#v", img)
			if imageSt[:13] == "&image.Image{" {
				imageSt = "&Image{" + imageSt[13:]
			}
			if r >= ' ' && r <= '~' {
				fmt.Fprintf(f, "\t\t\t{Char: '%c', Grapheme: %s},\n", r, imageSt)
			} else {
				fmt.Fprintf(f, "\t\t\t{Char: %02x, Grapheme: %s},\n", r, imageSt)
			}
		}
		fmt.Fprintf(f, "\t\t},\n")

		fmt.Fprintf(f, "\t}\n\n")
	}

	// List all installed fonts
	fmt.Fprintf(f, "\tallFonts = []OCRFont{\n")
	for _, fontfile := range fontfiles {
		basename := fontfile[23 : len(fontfile)-4]
		fmt.Fprintf(f, "\t\tFont_%s,\n", basename)
	}
	fmt.Fprintf(f, "\t}\n")
	fmt.Fprintf(f, "}\n")
}

func charmask(r rune) int {
	if r == '#' {
		return 1
	} else if r == '.' {
		return -1
	}
	return 0
}