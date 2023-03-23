package main






type indexPageData struck (
	Title string
	Subtitle string
)

func index(w http.ResponseWriter, r *http.Request) 
	ts, err := template.ParseFiles("pages/index.html")
	if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return 
	}

