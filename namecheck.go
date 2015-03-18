package static

import (
	"fmt"
	"html/template"
	"io/ioutil"
	
	"net/http"
	
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/results", myResult)
	
}

func root(w http.ResponseWriter, r *http.Request) {
	indexPage, err := ioutil.ReadFile("htmlpage/index.html")
	if err != nil {
		fmt.Fprint(w, "404 Not Found")
	}
	fmt.Fprint(w, string(indexPage))
}


var output, _ = ioutil.ReadFile("htmlpage/checkResult.html")
var ouputTemplate = template.Must(template.New("output").Parse(string(output)))


func myResult(w http.ResponseWriter, r *http.Request) {
	strEntered := r.FormValue("name")
	if strEntered == "Jagmeet" {
		err := ouputTemplate.Execute(w, "Yes, you are right.")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		err := ouputTemplate.Execute(w, "Sorry wrong !!!")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
