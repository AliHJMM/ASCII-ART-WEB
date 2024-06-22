package server




func Submit(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		renderErrorPage(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {

		renderErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, "internal server error", http.StatusInternalServerError)
		return
	}
	text := r.Form.Get("text")

	format := r.Form.Get("format")
	output, err := asciiart.Ascii(text, format)
	if err != nil {
		renderErrorPage(w, "Internal Server", http.StatusInternalServerError)
		return
	}

	data := PageData{
		OutputText: output,
	}

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

func renderErrorPage(w http.ResponseWriter, errMsg string, statusCode int) {
	w.WriteHeader(statusCode)
	data := PageData{
		ErrorMessage: errMsg,
		StatusCode:   strconv.Itoa(statusCode),
	}

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
