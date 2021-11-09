package models

import "encoding/xml"

// Transferencia struct
type Transferencia struct {
	TipoCriterio         string  `json:"tipoCriterio"`
	Fecha                string  `json:"fecha"`
	Criterio             string  `json:"criterio"`
	Emisor               int     `json:"emisor"`
	Receptor             int     `json:"receptor"`
	Cuenta               int     `json:"cuenta"`
	ReceptorParticipante int     `json:"receptorParticipante"`
	Monto                float32 `json:"monto"`
	Captcha              string  `json:"captcha"`
	TipoConsulta         int     `json:"tipoConsulta"`
}

type CEP struct {
	XMLName           xml.Name `xml:"SPEI_Tercero"`
	FechaOperacion    string   `json:"FechaOperacion" xml:"FechaOperacion,attr"`
	Hora              string   `xml:"Hora,attr"`
	ClaveSPEI         string   `xml:"ClaveSPEI,attr"`
	Sello             string   `xml:"sello,attr"`
	NumeroCertificado string   `xml:"numeroCertificado,attr"`
	CadenaCDA         string   `xml:"cadenaCDA,attr"`
	ClaveRastreo      string   `xml:"claveRastreo,attr"`
	Beneficiario      struct {
		Text          string `xml:",chardata"`
		BancoReceptor string `xml:"BancoReceptor,attr"`
		Nombre        string `xml:"Nombre,attr"`
		TipoCuenta    string `xml:"TipoCuenta,attr"`
		Cuenta        string `xml:"Cuenta,attr"`
		RFC           string `xml:"RFC,attr"`
		Concepto      string `xml:"Concepto,attr"`
		IVA           string `xml:"IVA,attr"`
		MontoPago     string `xml:"MontoPago,attr"`
	} `xml:"Beneficiario"`
	Ordenante struct {
		Text        string `xml:",chardata"`
		BancoEmisor string `xml:"BancoEmisor,attr"`
		Nombre      string `xml:"Nombre,attr"`
		TipoCuenta  string `xml:"TipoCuenta,attr"`
		Cuenta      string `xml:"Cuenta,attr"`
		RFC         string `xml:"RFC,attr"`
	} `xml:"Ordenante"`
}
