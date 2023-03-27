package nuget

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Package struct {
	Id             string    `json:"@id"`
	CatalogEntry   string    `json:"catalogEntry"`
	PackageContent string    `json:"packageContent"`
	Published      time.Time `json:"published"`
}

var nuget *Package

func FindPackage(packageName string, packageVersion string) (*Package, error) {
	name := strings.ToLower(packageName)
	version := packageVersion
	endpoint := fmt.Sprintf("/%s/%s.json", name, version)
	apiURL := NUGET_URL + endpoint
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Print("Wrong request.Check that you enter correct package name or version ")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print("There is error in body pars")
		return nil, err
	}

	err = json.Unmarshal(body, &nuget)
	if err != nil {
		log.Print("Can't unmarsh json for Nuget.Org")
		return nil, err
	}

	return nuget, nil
}
