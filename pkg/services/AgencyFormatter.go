package services

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
)

type AgencyFormatter interface {
	Format(agencies []Agency) (string, error)
}

type csvAgencyFormatter struct{}

func NewCsvAgencyFormatter(separator rune) AgencyFormatter {
	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = separator
		return gocsv.NewSafeCSVWriter(writer)
	})
	return &csvAgencyFormatter{}
}

type csvContent struct {
	Siret string `csv:"siret"`
	Naf   string `csv:"naf"`
}

func (c *csvAgencyFormatter) Format(agencies []Agency) (string, error) {
	var rows []csvContent
	for _, agency := range agencies {
		csvRow := csvContent{
			Siret: agency.Code,
			Naf:   agency.Naf,
		}
		rows = append(rows, csvRow)
	}

	csvContent, err := gocsv.MarshalString(&rows)
	if err != nil {
		return "", err
	}
	return csvContent, nil
}
