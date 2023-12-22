package utils

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
)

type ExtractDataFromRequest struct {
	dataFormat string
}

type ExtractDataFromRequestOptions func(*ExtractDataFromRequest)

func WithJSONFormat() ExtractDataFromRequestOptions {
	return func(e *ExtractDataFromRequest) {
		e.dataFormat = "json"
	}
}

func WithXMLFormat() ExtractDataFromRequestOptions {
	return func(e *ExtractDataFromRequest) {
		e.dataFormat = "xml"
	}
}

func NewExtractDataFromRequest(options ...ExtractDataFromRequestOptions) *ExtractDataFromRequest {

	var extraData ExtractDataFromRequest = ExtractDataFromRequest{}

	for _, option := range options {
		option(&extraData)
	}

	return &extraData
}

func (e ExtractDataFromRequest) Extract(r *http.Request) (ExtractDataFromRequest, error) {
	var contentType string = r.Header.Get("Content-Type")

	switch contentType {
	case "application/json":
		{
			return *NewExtractDataFromRequest(WithJSONFormat()), nil
		}
	case "application/xml":
		{
			return *NewExtractDataFromRequest(WithXMLFormat()), nil
		}
	}

	return e, errors.New("error format")
}

type DataProcessor interface {
	Process() error
}

func (e *ExtractDataFromRequest) UnmarshalAndProcess(r *http.Request, processor DataProcessor) error {

	if e.dataFormat == "json" {
		if err := json.NewDecoder(r.Body).Decode(&processor); err != nil {
			return err
		}
	}

	if e.dataFormat == "xml" {
		if err := xml.NewDecoder(r.Body).Decode(&processor); err != nil {
			return err
		}
	}
	defer r.Body.Close()

	return processor.Process()
}
