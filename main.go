package main

import (
	"bom-ascii-radar/bom"
	"flag"
	"fmt"
	"os/exec"
	"time"

	"golang.org/x/term"
)

var (
	productCode    = flag.String("productCode", "IDR763", "The product code to fetch images for")
	cacheDir       = flag.String("cacheDir", "/tmp/radar", "The directory to store cached images")
	timezone       = flag.String("timezone", "Australia/Hobart", "The timezone to use for the radar image timestamps")
	backgroundPath = flag.String("backgroundPath", "./background.png", "The path to the background image")
)

func main() {
	flag.Parse()

	loc, err := time.LoadLocation(*timezone)
	if err != nil {
		panic("Failed to load timezone: " + err.Error())
	}

	imageNames, err := bom.FetchImageNames(*productCode)
	if err != nil {
		panic("Failed to fetch image names: " + err.Error())
	}

	cache := make(map[string]string)

	width, height, err := term.GetSize(0)
	if err != nil {
		fmt.Printf("Failed to get terminal size: %v", err)
		panic(err)
	}

	// full height cause weird jumping issues on image change
	height -= 1

	fmt.Print("\033[2J") // Clear screen
	fmt.Print("\033[?25l")

	for {
		for _, theImageName := range imageNames {
			ascii := getRenderedImage(cache, theImageName, width, height)

			fmt.Printf("\033[0;0H") // Set cursor position
			fmt.Printf("%v", ascii)

			timestamp, _ := bom.GetDateTimeForImagePath(theImageName)
			tsFormat := "2006-01-02 15:04"

			fmt.Printf("\033[%d;%dH", height, width-len(tsFormat)+1) // set cursor to bottom right
			fmt.Print(timestamp.In(loc).Format(tsFormat))

			time.Sleep(1 * time.Second)
		}

		time.Sleep(2 * time.Second)
	}
}

func getRenderedImage(cache map[string]string, theImageName string, width int, height int) string {
	if cachedAscii, ok := cache[theImageName]; ok {
		return cachedAscii
	}

	f, err := bom.GetRadarImagePath(bom.FetchImageOptions{
		TheImageName: theImageName,
		CacheDir:     *cacheDir,
	})

	if err != nil {
		panic(err)
	}

	ascii, err := imageToAscii(f, width, height)
	if err != nil {
		fmt.Printf("Failed to convert image to ascii: %v", err)
		panic(err)
	}

	cache[theImageName] = ascii

	return ascii
}

func imageToAscii(filename string, width, height int) (string, error) {
	innerCmd := fmt.Sprintf(
		"composite %s %s - | ascii-image-converter - -d %d,%d -b -C",
		*backgroundPath,
		filename,
		width,
		height,
	)

	cmd := exec.Command("bash", "-c", innerCmd)

	out, err := cmd.CombinedOutput()
	return string(out), err
}
