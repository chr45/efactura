package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.linuxit.ro/chr45/e-factura/models"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	c.HTML(http.StatusOK, "uploadform.html", gin.H{})
}
func Convert(c *gin.Context) {
	file, _, err := c.Request.FormFile("xml")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer file.Close()
	invoice := models.Invoice{}
	data, err := ioutil.ReadAll(file)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	_ = xml.Unmarshal(data, &invoice)
	invoice.GetUM()
	invoice.HasDiscount()
	//toJson(invoice)
	//c.JSON(http.StatusOK, invoice)
	c.HTML(http.StatusOK, "efactura.html", gin.H{
		"invoice": invoice,
	})
}

func UnitatiMasura(c *gin.Context) {
	unitati := openUM()
	c.JSON(http.StatusOK, unitati.Option)
	toJson(unitati.Option)
}

func openUM() models.UnitatiMasura {
	data, _ := ioutil.ReadFile("unitati.xml")

	ds := models.UnitatiMasura{}

	_ = xml.Unmarshal([]byte(data), &ds)
	return ds
}

//func openXML() models.Invoice {
//	openfile := os.Getenv("FileToOpen")
//	data, _ := ioutil.ReadFile(openfile)

//	ds := models.Invoice{}

//	_ = xml.Unmarshal([]byte(data), &ds)
//	return ds
//}

func toJson(obj any) {
	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(b))
}
