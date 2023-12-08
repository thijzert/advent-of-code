package image

import "fmt"

type ocrChar struct {
	Char     rune
	Grapheme *Image
}

type OCRFont struct {
	characters []ocrChar
}

func (i *Image) OCRWithFont(font OCRFont) (string, error) {
	rv := ""
	x := 0
	for x < i.Width {
		bestChar, maxCon, offset := rune(0), 1024, 1
		for _, char := range font.characters {
			ok, confidence := charAt(i, char.Grapheme, x, 0)
			if ok && confidence < maxCon {
				bestChar = char.Char
				maxCon = confidence
				offset = char.Grapheme.Width
			}
		}
		if bestChar != 0 {
			rv = fmt.Sprintf("%s%c", rv, bestChar)
		}
		x += offset
	}

	return rv, nil
}

func (i *Image) OCR() (string, error) {
	rv := ""
	var lastError error = nil
	for _, font := range allFonts {
		s, err := i.OCRWithFont(font)
		if len(s) > len(rv) {
			rv, lastError = s, err
		} else if rv == "" {
			lastError = err
		}
	}
	return rv, lastError
}

func charAt(i *Image, grapheme *Image, x, y int) (bool, int) {
	rv, size := 0, 0
	for b := 0; b < grapheme.Height; b++ {
		for a := 0; a < grapheme.Width; a++ {
			c := grapheme.At(a, b)
			if c < 0 {
				continue
			}
			size++
			d := i.At(x+a, y+b)
			if (c == 0 && d == 0) || (c > 0 && d > 0) {
				rv++
			} else {
				rv--
			}
		}
	}
	return rv+(size>>3) >= size, size - rv
}
