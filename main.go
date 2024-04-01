package main

import (
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	apiTokens := os.Getenv("API_TOKENS")
	tokens := parseTokens(apiTokens)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log(r)
		fmt.Fprintf(w, "<h1>Ingress Nginx Cache!</h1>")
	})

	http.HandleFunc("/public", func(w http.ResponseWriter, r *http.Request) {
		log(r)
		currentTime := time.Now().Format("2006-01-02T15:04:05Z07:00")
		fmt.Fprintf(w, fmt.Sprintf("<h1>Public Page</h1><p>%s</p>", currentTime))
	})

	http.HandleFunc("/private", func(w http.ResponseWriter, r *http.Request) {
		log(r)
		token := r.Header.Get("Authorization")
		if !slices.Contains(tokens, token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		currentTime := time.Now().Format("2006-01-02T15:04:05Z07:00")
		fmt.Fprintf(w, fmt.Sprintf("<h1>Private Page</h1><p>%s</p>", currentTime))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func parseTokens(tokens string) []string {
	tokenList := make([]string, 0)
	for _, token := range strings.Split(tokens, ",") {
		tokenList = append(tokenList, strings.TrimSpace(token))
	}
	return tokenList
}

func log(r *http.Request) {
	fmt.Printf("%s %s %s %s\n", time.Now(), r.RemoteAddr, r.Method, r.URL.Path)
}
