package service

import (
	"encoding/csv"
	"errors"
	"strings"

	repository "a21hc3NpZ25tZW50/repository/fileRepository"
)

type FileService struct {
	Repo *repository.FileRepository
}

func (s *FileService) ProcessFile(content string) (map[string][]string, error) {
	if content == "" {
		return nil, errors.New("CSV file is empty")
	}

	reader := csv.NewReader(strings.NewReader(content))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 1 {
		return nil, errors.New("CSV file does not have a header row")
	}

	headers := records[0]
	data := make(map[string][]string)
	for _, header := range headers {
		data[header] = []string{}
	}

	for _, record := range records[1:] {
		if len(record) != len(headers) {
			return nil, errors.New("CSV file is invalid")
		}
		for i, value := range record {
			data[headers[i]] = append(data[headers[i]], value)
		}
	}

	return data, nil
}
