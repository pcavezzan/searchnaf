package services

import (
	"encoding/csv"
	"errors"
	"github.com/gocarina/gocsv"
	"io"
	"os"
	"strings"
)

type SearchNafCodeParser interface {
	Parse() ([]string, error)
}

type argumentSiretParser struct {
	siret string
}

func NewArgumentSiretParser(siren string) SearchNafCodeParser {
	return &argumentSiretParser{siret: siren}
}

func (a *argumentSiretParser) Parse() ([]string, error) {
	if a.siret == "" {
		return nil, errors.New("siret is required")
	}

	return []string{a.siret}, nil
}

type csvFileSiretParser struct {
	filePath string
}

type inputCsvContent struct {
	Siret string `csv:"siret"`
}

func NewCsvFileSiretParser(filePath string, separator rune) SearchNafCodeParser {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = separator
		return r // Allows use pipe as delimiter
	})

	return &csvFileSiretParser{filePath: filePath}
}

func (c *csvFileSiretParser) Parse() ([]string, error) {
	if c.filePath == "" {
		return nil, errors.New("file path is required")
	}

	searchNafCode, err := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer searchNafCode.Close()

	var inputCsvContent []*inputCsvContent
	if err = gocsv.UnmarshalFile(searchNafCode, &inputCsvContent); err != nil {
		return nil, err
	}

	var searchNafCodeApi []string
	for _, content := range inputCsvContent {
		if hasSirenNumber(content.Siret) {
			searchNafCodeApi = append(searchNafCodeApi, content.Siret)
		}
	}
	return searchNafCodeApi, nil
}

func hasSirenNumber(siret string) bool {
	return len(siret) > 0 && !strings.EqualFold(siret, "?")
}
