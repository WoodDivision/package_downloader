package nuget

import (
	"fmt"
	"log"
	"nuget-downloader/nexus"
	"nuget-downloader/service"
	"strings"
	"time"
)

func DownloadNuget(packageName string, packageVersion string, repository string) {
	name := strings.ToLower(packageName)
	pac := ToDo{name, packageVersion}
	packageToDownload, err := FindDependencies(pac)
	if err != nil {
		return
	}
	packageToDownload[pac] = true
	for pac, _ := range packageToDownload {
		nexusItems, err := nexus.CheckNexus(packageName, packageVersion, repository)
		if len(nexusItems) > 0 {
			log.Print("Package already in Nexus")
			return
		}
		if err != nil {
			return
		}
		fileName := fmt.Sprintf("%s.%s.nupkg", pac.Name, pac.Version)
		nuget, err := FindPackage(pac.Name, pac.Version)
		if err != nil {
			return
		}
		date := time.Date(2022, 02, 24, 00, 00, 00, 00, time.UTC)
		if nuget.Published.After(date) == false {
			log.Printf("Package: %s, version %s, published: %s", pac.Name, pac.Version, nuget.Published.Format("2006-01-02 15:04:05"))
			err = service.DownloadFile(nuget.PackageContent, fileName)
			if err != nil {
				return
			}
			log.Printf("Downloaded")
		} else {
			log.Printf("Package: %s, version %s, published: %s", pac.Name, pac.Version, nuget.Published.Format("2006-01-02 15:04:05"))
			log.Printf("Skip")
		}
	}
	log.Print("Done")
}