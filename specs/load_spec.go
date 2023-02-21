package specs

import (
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/go-openapi/loads"
)

// this file is to parse a single Spec file

var swagBase string

const SwagBaseKey = "SWAG_BASE"

func SwagBase() string {
	if swagBase == "" {
		swagBase = os.Getenv(SwagBaseKey)
	}
	return swagBase
}

func LoadSpec(swagPath string) (*loads.Document, error) {
	if !(strings.HasPrefix(swagPath, "/") || strings.HasPrefix(swagPath, ".")) {
		if urlIns, err := url.Parse(swagPath); err != nil || urlIns.Host == "" {
			swagPath = path.Join(SwagBase(), swagPath)
		}
	}
	doc, err := loads.Spec(swagPath)
	return doc, err
}

func LoadExpanded(path string) (*loads.Document, error) {
	doc, err := LoadSpec(path)
	if err != nil {
		return nil, err
	}
	doc, err = doc.Expanded()
	return doc, err
}
