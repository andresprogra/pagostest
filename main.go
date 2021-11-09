package main

import (
	"fmt"

	"github.com/andresprogra/pagos/cep"
	"github.com/andresprogra/pagos/models"
)

func main() {

	cep1, _ := GetCEP("03-11-2021", "MBAN01002111040097466893", 40012, 40002, 5256783407957290, 14000)
	cep2, _ := GetCEP("25-10-2021", "2021102640014BMOV0030430923640", 40014, 40012, 6361171735, 5000)

	fmt.Printf("cep1.Sello: %v\n", cep1.Sello)
	fmt.Printf("cep2.Sello: %v\n", cep2.Sello)
}

func GetCEP(fecha, criterio string, emisor, receptor, cuenta int, monto float32) (*models.CEP, error) {

	t := cep.Transferencia{
		TipoCriterio:         "T",
		Fecha:                fecha,
		Criterio:             criterio,
		Emisor:               emisor,
		Receptor:             receptor,
		Cuenta:               cuenta,
		ReceptorParticipante: 0,
		Monto:                monto,
		TipoConsulta:         1,
		Captcha:              "c",
	}

	return t.Validar()

}
