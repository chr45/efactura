package models

import (
	"encoding/xml"
)

type OrderReference struct {
	Text string `xml:",chardata"`
	ID   string `xml:"ID"`
}

type ContractDocumentReference struct {
	Text string `xml:",chardata"`
	ID   string `xml:"ID"`
}

type PostalAddress struct {
	Text             string  `xml:",chardata"`
	StreetName       string  `xml:"StreetName"`
	CityName         string  `xml:"CityName"`
	CountrySubentity string  `xml:"CountrySubentity"`
	Country          Country `xml:"Country"`
}

type Country struct {
	Text               string `xml:",chardata"`
	IdentificationCode string `xml:"IdentificationCode"`
}

type PartyTaxScheme struct {
	Text      string    `xml:",chardata"`
	CompanyID string    `xml:"CompanyID"`
	TaxScheme TaxScheme `xml:"TaxScheme"`
}

type TaxScheme struct {
	Text string `xml:",chardata"`
	ID   string `xml:"ID"`
}

type PartyLegalEntity struct {
	Text             string `xml:",chardata"`
	RegistrationName string `xml:"RegistrationName"`
	CompanyID        string `xml:"CompanyID"`
}

type Contact struct {
	Text string `xml:",chardata"`
	Name string `xml:"Name"`
}

type Party struct {
	Text             string           `xml:",chardata"`
	PostalAddress    PostalAddress    `xml:"PostalAddress"`
	PartyTaxScheme   PartyTaxScheme   `xml:"PartyTaxScheme"`
	PartyLegalEntity PartyLegalEntity `xml:"PartyLegalEntity"`
	Contact          Contact          `xml:"Contact"`
}

type AccountingParty struct {
	Text  string `xml:",chardata"`
	Party Party  `xml:"Party"`
}

type DeliveryLocation struct {
	Text    string        `xml:",chardata"`
	Address PostalAddress `xml:"Address"`
}

type Delivery struct {
	Text               string           `xml:",chardata"`
	ActualDeliveryDate string           `xml:"ActualDeliveryDate"`
	DeliveryLocation   DeliveryLocation `xml:"DeliveryLocation"`
}

type PayeeFinancialAccount struct {
	Text string `xml:",chardata"`
	ID   string `xml:"ID"`
}

type PaymentMeans struct {
	Text                  string                `xml:",chardata"`
	PaymentMeansCode      string                `xml:"PaymentMeansCode"`
	PayeeFinancialAccount PayeeFinancialAccount `xml:"PayeeFinancialAccount"`
}

type TaxAmount struct {
	Text       string `xml:",chardata"`
	CurrencyID string `xml:"currencyID,attr"`
}

type TaxableAmount struct {
	Text       string `xml:",chardata"`
	CurrencyID string `xml:"currencyID,attr"`
}

type TaxCategory struct {
	Text      string    `xml:",chardata"`
	ID        string    `xml:"ID"`
	Percent   string    `xml:"Percent"`
	TaxScheme TaxScheme `xml:"TaxScheme"`
}

type TaxSubtotal struct {
	Text          string        `xml:",chardata"`
	TaxableAmount TaxableAmount `xml:"TaxableAmount"`
	TaxAmount     TaxAmount     `xml:"TaxAmount"`
	TaxCategory   TaxCategory   `xml:"TaxCategory"`
}

type TaxTotal struct {
	Text        string      `xml:",chardata"`
	TaxAmount   TaxAmount   `xml:"TaxAmount"`
	TaxSubtotal TaxSubtotal `xml:"TaxSubtotal"`
}

type TaxCurrency struct {
	Text       string `xml:",chardata"`
	CurrencyID string `xml:"currencyID,attr"`
}

type LegalMonetaryTotal struct {
	Text                string      `xml:",chardata"`
	LineExtensionAmount TaxCurrency `xml:"LineExtensionAmount"`
	TaxExclusiveAmount  TaxCurrency `xml:"TaxExclusiveAmount"`
	TaxInclusiveAmount  TaxCurrency `xml:"TaxInclusiveAmount"`
	PayableAmount       TaxCurrency `xml:"PayableAmount"`
}

type InvoicedQuantity struct {
	Text     string `xml:",chardata"`
	UnitCode string `xml:"unitCode,attr"`
}

type ClassifiedTaxCategory struct {
	Text      string    `xml:",chardata"`
	ID        string    `xml:"ID"`
	Percent   string    `xml:"Percent"`
	TaxScheme TaxScheme `xml:"TaxScheme"`
}

type Item struct {
	Text                  string                `xml:",chardata"`
	Description           string                `xml:"Description"`
	Name                  string                `xml:"Name"`
	ClassifiedTaxCategory ClassifiedTaxCategory `xml:"ClassifiedTaxCategory"`
}

type Price struct {
	Text        string      `xml:",chardata"`
	PriceAmount TaxCurrency `xml:"PriceAmount"`
}

type InvoiceLine struct {
	Text                string           `xml:",chardata"`
	ID                  string           `xml:"ID"`
	Note                string           `xml:"Note"`
	InvoicedQuantity    InvoicedQuantity `xml:"InvoicedQuantity"`
	LineExtensionAmount TaxCurrency      `xml:"LineExtensionAmount"`
	AllowanceCharge     AllowanceCharge  `xml:"AllowanceCharge"`
	Item                Item             `xml:"Item"`
	Price               Price            `xml:"Price"`
	UM                  string           `xml:"-"`
}

type AllowanceCharge struct {
	ChargeIndicator           string      `xml:"ChargeIndicator"`
	AllowanceChargeReasonCode string      `xml:"AllowanceChargeReasonCode"`
	AllowanceChargeReason     string      `xml:"AllowanceChargeReason"`
	Amount                    TaxCurrency `xml:"Amount"`
	BaseAmount                TaxCurrency `xml:"BaseAmount"`
}

type UnitatiMasura struct {
	XMLName xml.Name `xml:"select"`
	Text    string   `xml:",chardata"`
	Option  []struct {
		Text      string `xml:",chardata"`
		DataLabel string `xml:"data-label,attr"`
		Value     string `xml:"value,attr"`
	} `xml:"option"`
}
