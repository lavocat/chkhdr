package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func main() {
	var url = flag.String("url", "http://localhost:8080", "URL to check")
	flag.Parse()
	checkURL(*url)

}

func checkURL(url string) {
	statusPass := color.New(color.Bold, color.FgGreen).PrintFunc()
	statusFail := color.New(color.Bold, color.FgRed).PrintFunc()
	var secHeaders = [11]string{
		"Access-Control-Allow-Origin",
		"Content-Type",
		"Content-Security-Policy",
		"Cross-Origin-Embedder-Policy",
		"Cross-Origin-Resource-Policy",
		"Cross-Origin-Opener-Policy",
		"Strict-Transport-Security",
		"Referrer-Policy",
		"X-Content-Type-Options",
		"X-Frame-Options",
		"X-XSS-Protection",
	}

	headers := getResponse(url)

	for _, header := range secHeaders {
		if val, ok := headers[header]; ok {
			fmt.Printf("\t%s: ", header)
			statusPass(val)
			fmt.Println()
		} else {
			fmt.Printf("\t%s: ", header)
			statusFail("Header Missing!")
			fmt.Println()
		}
	}
}

func getResponse(url string) http.Header {
	statusPass := color.New(color.Bold, color.FgGreen).PrintFunc()
	statusFail := color.New(color.Bold, color.FgRed).PrintFunc()
	resp, err := http.Get(url)
	if err != nil {
		statusFail(err)
		fmt.Println()
	}
	fmt.Printf("Checking URL: %s (", url)
	statusPass(resp.Status)
	fmt.Println(")")
	defer resp.Body.Close()
	return resp.Header
}
