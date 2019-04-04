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

	"github.com/gregzuro/three-colors/count"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png" // one of the test files - supposedly a jpg - didn't work until this was added...
)

// keep a map of results so that we can short-circuit on duplicates
var results map[string][]string
var resultsLock = sync.RWMutex{}
var wg sync.WaitGroup // to wait for goroutines

const maxMB uint64 = 128

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
			if bToMb(memStats.Alloc) > maxMB {
				runtime.GC() // need this to avoid deadlock in some cases
				runtime.ReadMemStats(&memStats)
				if bToMb(memStats.Alloc) > maxMB {
					time.Sleep(time.Millisecond * 500)
					log.Printf("using too much memory. pausing. (Alloc = %v MiB)", bToMb(memStats.Alloc))
				}
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
		go doCount(url, image)
	}
	wg.Done()
}

// count does the countng of colors and finding the three most common
func doCount(url string, image image.Image) {

	colorCounts := count.CountColorsFromImage(image)
	topThree := count.FindTopThreeFromCounts(colorCounts)

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
