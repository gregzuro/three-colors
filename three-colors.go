//
// given a list of images file URLs, write a csv file with the RGB values for the three most prevalent colors from each image
//
// - assume that input list is arbitrarily long, so it's read by a goroutine with single results (file URLs) picked as needed
// - assuming literal RGB colors rather than some clustered / interpolated / average colors
// - there's no super-accurate way to limit memory usage to a specific amount, so
// 	- we read files from network into memory until we go over some maximum amount of heap,
// 		then we pause and trigger a GC for good measure until the heap usage has fallen

package main

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png" // one of the test files - supposedly a jpg - didn't work until this was added...
)

// keep a map of results so that we can short-circuit on duplicates
var results map[string][]string
var resultsLock = sync.RWMutex{}
var wg sync.WaitGroup // to wait for goroutines

func main() {
	var memStats runtime.MemStats

	results = make(map[string][]string)

	fileURLs := make(chan string)
	go getFileURLsFromURL("https://gist.githubusercontent.com/ehmo/e736c827ca73d84581d812b3a27bb132/raw/77680b283d7db4e7447dbf8903731bb63bf43258/input.txt", fileURLs)

	totalStart := time.Now()

	for true {

		fileURL := <-fileURLs // get the next fileURL from the channel
		if fileURL == "*DONE*" {
			wg.Wait()
			log.Println("totalTime: ", time.Since(totalStart))
			return
		}

		// check to see if we've already done this one, since there are a bunch of duplicates in the prototype list
		resultsLock.RLock()
		if _, ok := results[fileURL]; ok {
			// don't print results for duplicates
			resultsLock.RUnlock()
		} else {
			resultsLock.RUnlock()
			wg.Add(1)
			go get(fileURL)

			time.Sleep(time.Millisecond * 10) // pace yourself
		}

		for {
			runtime.ReadMemStats(&memStats)

			// if we're using too much memory
			if bToMb(memStats.Alloc) > 128 {
				log.Printf("using too much memory. pausing. (Alloc = %v MiB)", bToMb(memStats.Alloc))
				time.Sleep(time.Millisecond * 500)
				runtime.GC() // need this to avoid deadlock in some cases
			} else {
				break
			}
		}
	}
	return
}

// get gets the image file from the URL and decodes it
func get(url string) {

	// indicate in-progress files but adding a (blank) result to the map
	resultsLock.Lock()
	results[url] = []string{""}
	resultsLock.Unlock()
	_, fileContents, err := getFileByURL(url)
	if err != nil {
		fmt.Println(err)
	}

	image, _, err := image.Decode(fileContents)
	fileContents.Close()
	if err != nil {
		log.Printf("error decoding %s: %v", url, err)
	} else {
		wg.Add(1)
		go count(url, image)
	}
	wg.Done()
}

// count does the countng of colors and finding the three most common
func count(url string, image image.Image) {
	colorCounts := countColorsFromImage(image)
	topThree := findTopThreeFromCounts(colorCounts)

	printResults(url, topThree, colorCounts, false)

	// cache these results
	resultsLock.Lock()
	results[url] = topThree
	resultsLock.Unlock()

	wg.Done()
}

// getFileURLs fetches file names from the given URL, returning them one by one on a channel
func getFileURLsFromURL(url string, fn chan string) {

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fn <- scanner.Text()
	}

	fn <- "*DONE*"
	return
}

// getFileByURL gets the contents of the file, returning a single blob
func getFileByURL(url string) (int64, io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, nil, err
	}
	//defer resp.Body.Close()

	return resp.ContentLength, resp.Body, nil
}

func countColorsFromImage(image image.Image) map[string]uint {

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

// findTopThreeFromCounts finds three most common colors from a colorCounts-style map
// 	are we likely to want some number other (greater) than three?
// 	if so, then do this a different way
func findTopThreeFromCounts(colorCounts map[string]uint) []string {

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

// printResults prints out the results in CSV style
func printResults(fileURL string, topThree []string, colorCounts map[string]uint, duplicate bool) {

	fmt.Printf("%s,", fileURL)
	for i, c := range topThree {
		if c != "" {
			if duplicate {
				fmt.Printf("#%02x%02x%02x", c[0], c[1], c[2])
			} else {
				//				fmt.Printf("#%02x%02x%02x (%v)", c[0], c[1], c[2], colorCounts[c])
				fmt.Printf("#%02x%02x%02x", c[0], c[1], c[2])
			}
			if i < len(topThree)-1 {
				fmt.Printf(",")
			} else {
				fmt.Printf("\n")
			}
		}
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
