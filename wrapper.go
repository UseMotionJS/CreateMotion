package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// SetupWrapper downloads relevant wrapper files from the server
// and places them in src/wrappers/react or src/wrappers/vue based on user choice.
func SetupWrapper(wrapperChoice string) error {
	switch wrapperChoice {
	case "react":
		if err := os.MkdirAll("src/wrappers/react", 0755); err != nil {
			return err
		}
		fmt.Println(ColorCyan + "Downloading React wrapper files..." + ColorReset)
		// Example file list for React. Replace with your real URLs.
		reactFiles := []struct {
			URL      string
			FileName string
		}{
			{
				URL:      "https://install.mot.mindity.net/files/react-wrapper.js",
				FileName: "src/wrappers/react/Wrapper.js",
			},
			{
				URL:      "https://install.mot.mindity.net/files/react-helper.js",
				FileName: "src/wrappers/react/Helper.js",
			},
		}
		for _, f := range reactFiles {
			if err := downloadFile(f.URL, f.FileName); err != nil {
				return err
			}
		}

	case "vue":
		if err := os.MkdirAll("src/wrappers/vue", 0755); err != nil {
			return err
		}
		fmt.Println(ColorCyan + "Downloading Vue wrapper files..." + ColorReset)
		// Example file list for Vue. Replace with your real URLs.
		vueFiles := []struct {
			URL      string
			FileName string
		}{
			{
				URL:      "https://install.mot.mindity.net/files/vue-wrapper.js",
				FileName: "src/wrappers/vue/Wrapper.js",
			},
			{
				URL:      "https://install.mot.mindity.net/files/vue-helper.js",
				FileName: "src/wrappers/vue/Helper.js",
			},
		}
		for _, f := range vueFiles {
			if err := downloadFile(f.URL, f.FileName); err != nil {
				return err
			}
		}

	default:
		fmt.Println(ColorYellow + "Unknown wrapper. Skipping wrapper setup." + ColorReset)
	}
	return nil
}

// downloadFile fetches a file from a given URL and saves it to the specified local path.
func downloadFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download %s: %s", url, resp.Status)
	}

	// Ensure the local directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}
	fmt.Printf(ColorGreen+"Downloaded %s -> %s\n"+ColorReset, url, filePath)
	return nil
}
