package server




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
