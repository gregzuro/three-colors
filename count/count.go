package count

import "image"

// CountColorsFromImage counts the colors in the given image, returning a map with the counts for all colors
func CountColorsFromImage(image image.Image) map[string]uint {

	bounds := image.Bounds()

	// count the colors
	colorCounts := make(map[string]uint) // this can have ~4MM entries - colors that jpeg can have (https://www.hackerfactor.com/blog/index.php?/archives/250-Showing-JPEGs-True-Color.html)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := image.At(x, y).RGBA()
			c := string([]byte{byte(r >> 8), byte(g >> 8), byte(b >> 8)}) // I notice an occasional off=by-one-count compared with photoshop-created images on (at least) the Blue channel)
			// this is likely related to: https://www.hackerfactor.com/blog/index.php?/archives/250-Showing-JPEGs-True-Color.html as well...
			colorCounts[c]++
		}
	}

	return colorCounts
}

// FindTopThreeFromCounts finds three most common colors from a colorCounts-style map
// 	are we likely to want some number other (greater) than three?
// 	if so, then do this a different way
func FindTopThreeFromCounts(colorCounts map[string]uint) []string {

	topThree := make([]string, 3)
	// ifCount := 0
	for c, n := range colorCounts {
		// ifCount++
		if n > colorCounts[topThree[2]] { // first check to see if this one is in the top three at all
			// ifCount++
			if n > colorCounts[topThree[0]] { // then start check from most common (first) down to least (third)
				// this couples the least number of reassighments with the least number of conditionals...
				topThree[2] = topThree[1]
				topThree[1] = topThree[0]
				topThree[0] = c
			} else {
				// ifCount++
				if n > colorCounts[topThree[1]] {
					topThree[2] = topThree[1]
					topThree[1] = c
				} else {
					topThree[2] = c
				}
			}
		}
	}
	return topThree
}
