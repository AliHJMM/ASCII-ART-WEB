package server

import (
	"html/template"
	"net/http"
	"strconv"

	asciiart "ascii-web/ascii-art"
)

// PageData defines the structure for data to be passed to templates.
type PageData struct {
	OutputText   string // OutputText contains the generated ASCII art or other output.
	ErrorMessage string // ErrorMessage holds any error message to display.
	StatusCode   string // StatusCode represents the HTTP status code as a string.
}

// HomeHandler handles the home page and serves the index.html template or static files.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Handle only GET requests
	if r.Method != "GET" {
		renderErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Serve the main page (index.html)
	if r.URL.Path == "/" {
		tmpl, err := template.ParseFiles("./statics/index.html")
		if err != nil {
			renderErrorPage(w, "internal server error", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, map[string]string{"OutputText": ""})
		if err != nil {
			renderErrorPage(w, "internal server error", http.StatusInternalServerError)
			return
		}
		return
	}

	// Serve style.css
	if r.URL.Path == "/style.css" {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "./statics/style.css")
		return
	}

	// Serve error.css
	if r.URL.Path == "/error.css" {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "./statics/error.css")
		return
	}

	// Handle other paths with a 404 error
	renderErrorPage(w, "Not found", http.StatusNotFound)
}

// Submit handles the POST request to generate ASCII art based on user input.
func Submit(w http.ResponseWriter, r *http.Request) {
	// Handle only /ascii-art path
	if r.URL.Path != "/ascii-art" {
		renderErrorPage(w, "Not found", http.StatusNotFound)
		return
	}

	// Handle only POST requests
	if r.Method != "POST" {
		renderErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, "internal server error", http.StatusInternalServerError)
		return
	}
	text := r.Form.Get("text")
	format := r.Form.Get("format")

	// Generate ASCII art based on input
	output, err := asciiart.Ascii(text, format)
	if err != nil {
		renderErrorPage(w, "Internal Server", http.StatusInternalServerError)
		return
	}

	// Prepare data for rendering the template
	data := PageData{
		OutputText: output,
	}

	// Render index.html template with generated data
	tmpl, err := template.ParseFiles("./statics/index.html")
	if err != nil {
		renderErrorPage(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		renderErrorPage(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

// renderErrorPage renders an error page with the specified message and status code.
func renderErrorPage(w http.ResponseWriter, errMsg string, statusCode int) {
	w.WriteHeader(statusCode)
	data := PageData{
		ErrorMessage: errMsg,
		StatusCode:   strconv.Itoa(statusCode),
	}

	// Render error.html template with error data
	tmpl, err := template.ParseFiles("./statics/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
