package cep

import (
	"crypto/tls"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/andresprogra/pagos/models"
)

type Transferencia models.Transferencia

func (t Transferencia) Validar() (*models.CEP, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Transport: tr,
	}

	payload := url.Values{}

	payload.Set("tipoCriterio", "T")
	payload.Set("criterio", t.Criterio)
	payload.Set("fecha", t.Fecha)
	payload.Set("emisor", strconv.Itoa(t.Emisor))
	payload.Set("receptor", strconv.Itoa(t.Receptor))
	payload.Set("cuenta", strconv.Itoa(t.Cuenta))
	payload.Set("receptorParticipante", "0")
	payload.Set("monto", strconv.FormatFloat(float64(t.Monto), 'f', 2, 32))
	payload.Set("captcha", "c")
	payload.Set("tipoConsulta", strconv.Itoa(1))

	data := strings.NewReader(payload.Encode())

	req, err := http.NewRequest(http.MethodPost, "https://www.banxico.org.mx/cep/valida.do", data)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		err := errors.New("hubo un error")
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		err := errors.New("hubo un error")
		return nil, err
	}

	setCookie := res.Header.Values("Set-Cookie")

	defer res.Body.Close()

	req, err = http.NewRequest(http.MethodGet, "https://www.banxico.org.mx/cep/descarga.do?formato=XML", nil)

	for _, v := range setCookie {
		req.Header.Add("Cookie", v)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		err := errors.New("hubo un error")
		return nil, err
	}

	res, err = client.Do(req)
	if err != nil {
		err := errors.New("hubo un error")
		return nil, err
	}

	response, _ := ioutil.ReadAll(res.Body)

	cep := &models.CEP{}

	err = xml.Unmarshal(response, cep)

	if err != nil {
		err := errors.New("hubo un error")
		return nil, err
	}

	return cep, nil

}
