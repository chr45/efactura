package models

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

type Invoice struct {
	XMLName                   xml.Name                  `xml:"Invoice"`
	Text                      string                    `xml:",chardata"`
	Xmlns                     string                    `xml:"xmlns,attr"`
	Cbc                       string                    `xml:"cbc,attr"`
	Cac                       string                    `xml:"cac,attr"`
	Ns4                       string                    `xml:"ns4,attr"`
	Xsi                       string                    `xml:"xsi,attr"`
	SchemaLocation            string                    `xml:"schemaLocation,attr"`
	CustomizationID           string                    `xml:"CustomizationID"`
	ID                        string                    `xml:"ID"`
	IssueDate                 string                    `xml:"IssueDate"`
	DueDate                   string                    `xml:"DueDate"`
	InvoiceTypeCode           string                    `xml:"InvoiceTypeCode"`
	Note                      string                    `xml:"Note"`
	TaxPointDate              string                    `xml:"TaxPointDate"`
	DocumentCurrencyCode      string                    `xml:"DocumentCurrencyCode"`
	OrderReference            OrderReference            `xml:"OrderReference"`
	ContractDocumentReference ContractDocumentReference `xml:"ContractDocumentReference"`
	AccountingSupplierParty   AccountingParty           `xml:"AccountingSupplierParty"`
	AccountingCustomerParty   AccountingParty           `xml:"AccountingCustomerParty"`
	Delivery                  Delivery                  `xml:"Delivery"`
	PaymentMeans              PaymentMeans              `xml:"PaymentMeans"`
	TaxTotal                  TaxTotal                  `xml:"TaxTotal"`
	LegalMonetaryTotal        LegalMonetaryTotal        `xml:"LegalMonetaryTotal"`
	InvoiceLine               []InvoiceLine             `xml:"InvoiceLine"`
	InvoiceLineNew            []InvoiceLine
	WithDiscount              string
}

func (i *Invoice) GetUM() {
	unitati := openUM()
	for _, l := range i.InvoiceLine {
		for _, um := range unitati.Option {
			if um.Value == l.InvoicedQuantity.UnitCode {
				l.UM = um.Text
				l.UM = strings.Replace(l.UM, fmt.Sprintf("%s - ", um.Value), "", 1)
				i.InvoiceLineNew = append(i.InvoiceLineNew, l)
			}
		}
	}
}

func (i *Invoice) HasDiscount() {
	i.WithDiscount = "0"
	for _, l := range i.InvoiceLine {
		if l.AllowanceCharge.Amount.Text != "" {
			i.WithDiscount = "1"
			return
		}
	}
}

func openUM() UnitatiMasura {
	data, _ := ioutil.ReadFile("unitati.xml")

	ds := UnitatiMasura{}

	_ = xml.Unmarshal([]byte(data), &ds)
	return ds
}
