package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"text/template"

	"github.com/tidwall/gjson"
)

type Configuration struct {
	AndroidAdmobId  string
	Host            string
	AndroidBundleID string
}

func main() {
	fmt.Println("------- --------- ---------")
	fmt.Println("Reading the config file 📦")
	fmt.Println("------- --------- ---------")

	// Read the config file
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Create and parse the template
	templateManifestFile, err := template.ParseFiles("./src/android/AndroidManifest.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	// Parse config file
	config := string(configFile)
	androidAdmobId := gjson.Get(config, "androidAdmobId").String()
	host := gjson.Get(config, "host").String()
	androidBundleId := gjson.Get(config, "androidBundleId").String()

	configurationValues := Configuration{
		AndroidAdmobId:  androidAdmobId,
		Host:            host,
		AndroidBundleID: androidBundleId}

	// Generate Final final
	FinalFile, err := os.Create("./src/android/AndroidManifest.xml")
	if err != nil {
		log.Fatal(err)
	}

	templateManifestFile.Execute(FinalFile, configurationValues)

	fmt.Println("------- --------- ---------")
	fmt.Println("🔧🔧🔧🔧🔧 Configuration completed successfully 🔧🔧🔧🔧🔧")
	fmt.Println("------- --------- ---------")

}
