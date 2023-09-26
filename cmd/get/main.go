package get

import (
	"flag"
	"log"
	"os"
	"package-downloader/rest"
)

var (
	name       string
	version    string
	repository string
)

func main() {
	log.SetOutput(os.Stdout)
	flag.StringVar(&name, "n", "", "Package name")
	flag.StringVar(&version, "v", "", "Package version")
	flag.StringVar(&repository, "r", "nuget-freeze", "Package repository name in Nexus.Action")
	flag.Parse()
	c := rest.NewClient("khristolyubov", "qwB6jr_nm", "https://nexus.action-media.ru/")
	c.NpmClient.DownloadDependency(name, version, repository)
}