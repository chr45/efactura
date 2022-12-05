package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"git.linuxit.ro/chr45/e-factura/models"
)

var funcMap = template.FuncMap{
	"inc": func(i int) int {
		return i + 1
	},
}

func Upload(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "uploadform.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Convert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println(r.URL)
		file, _, err := r.FormFile("xml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		invoice := models.Invoice{}
		data, err := io.ReadAll(file)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer file.Close()
		_ = xml.Unmarshal(data, &invoice)
		invoice.GetUM()
		invoice.HasDiscount()
		fp := path.Join("templates", "efactura.html")
		tmpl := template.New("").Funcs(template.FuncMap(funcMap))
		template.Must(tmpl.ParseFiles(fp))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var htmldata = make(map[string]interface{})
		htmldata["invoice"] = invoice
		err = tmpl.Funcs(funcMap).ExecuteTemplate(w, "efactura.html", htmldata)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func UnitatiMasura(w http.ResponseWriter, r *http.Request) {
	unitati := openUM()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := unitati.Option
	w.Write(toJsonb(data))
}

func openUM() models.UnitatiMasura {
	data, _ := os.ReadFile("unitati.xml")
	ds := models.UnitatiMasura{}
	_ = xml.Unmarshal([]byte(data), &ds)
	return ds
}

/*func toJson(obj any) string {
	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err.Error()
	}
	return string(b)
}*/

func toJsonb(obj any) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil
	}
	return b
}
