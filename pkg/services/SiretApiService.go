package services

import (
	"github.com/imroc/req/v3"
)

type HttpSiretService struct{}

type DataGouvAgency struct {
	Etablissement struct {
		Id                 int    `json:"id"`
		Siren              string `json:"siren"`
		Nic                string `json:"nic"`
		Siret              string `json:"siret"`
		CodePostal         string `json:"code_postal"`
		ActivitePrincipale string `json:"activite_principale"`
		Longitude          string `json:"longitude"`
		Latitude           string `json:"latitude"`
		GeoAdresse         string `json:"geo_adresse"`
	} `json:"etablissement"`
}

type Agency struct {
	CodeCompany string `json:"siren" csv:"siren"`
	Code        string `json:"siret" csv:"siret"`
	ZipCode     string `json:"code_postal" csv:"code_postal"`
	Address     string `json:"geo_adresse" csv:"geo_adresse"`
	Naf         string `json:"activite_principale" csv:"code_naf"`
}

func (httpSirenService *HttpSiretService) GetAgency(siret string) (Agency, error) {
	var result DataGouvAgency
	resp, err := req.C().R(). // Use R() to create a request
					SetHeader("Accept", "application/json"). // Chainable request settings
					SetPathParam("siret", siret).
					SetResult(&result).
					Get("https://entreprise.data.gouv.fr/api/sirene/v3/etablissements/{siret}")
	var agency Agency
	if resp.IsSuccess() {
		agency = Agency{
			Code:        result.Etablissement.Siret,
			CodeCompany: result.Etablissement.Siren,
			Address:     result.Etablissement.GeoAdresse,
			ZipCode:     result.Etablissement.CodePostal,
			Naf:         result.Etablissement.ActivitePrincipale,
		}
	}

	return agency, err
}
