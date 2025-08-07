package asciiart

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"strings"
)

const (
	ValidChars = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// `Generate` generates an ASCII art from an image.
func Generate(dest image.Image, threshold int) string {
	srcBounds := dest.Bounds()

	asciiArt := strings.Builder{}
	asciiArt.WriteString("\n")

	for y := 0; y < srcBounds.Max.Y; y += 2 {
		line := strings.Builder{}

		for x := range srcBounds.Max.X {
			c := color.GrayModel.Convert(dest.At(x, y))
			gray, _ := c.(color.Gray)
			if gray.Y < uint8(threshold) {
				line.WriteString(" ")
			} else {
				line.WriteString(randomChars(ValidChars))
			}
		}

		if len(strings.Fields(line.String())) != 0 {
			asciiArt.WriteString(line.String() + "\n")
		}
	}

	return asciiArt.String()
}

func randomChars(chars string) string {
	r := rand.Intn(len(chars))
	return string(chars[r])
}

// `CalcOTSUThreshold` calculates the threshold value using the OTSU's method.
func CalcOTSUThreshold(dest image.Image, ySize, xSize int) int {
	histogram := make([]int, 256) // 0 - 255 histogram

	for y := range ySize {
		for x := range xSize {
			c := color.GrayModel.Convert(dest.At(x, y))
			gray, _ := c.(color.Gray)
			histogram[gray.Y]++
		}
	}

	t := 0
	max := 0.0

	for i := range 256 {
		w1, w2 := 0, 0     // pixel number
		sum1, sum2 := 0, 0 // total gray value
		m1, m2 := 0.0, 0.0 // average gray value

		for j := range i {
			w1 += histogram[j]
			sum1 += histogram[j] * j
		}

		for j := i; j < 256; j++ {
			w2 += histogram[j]
			sum2 += j * histogram[j]
		}

		if 0 < w1 {
			m1 = float64(sum1) / float64(w1)
		}

		if 0 < w2 {
			m2 = float64(sum2) / float64(w2)
		}

		tmp := float64(w1) * float64(w2) * math.Pow(m1-m2, 2)

		if max < tmp {
			max = tmp
			t = i
		}
	}

	return t
}
