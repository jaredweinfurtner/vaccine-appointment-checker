package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type VaccinationCenter struct {
	Name    string `json:"Zentrumsname"`
	Zip     string `json:"PLZ"`
	City    string `json:"Ort"`
	State   string `json:"Bundesland"`
	BaseUrl string `json:"URL"`
	Address string `json:"Adresse"`
}

type LocationResponse struct {
	VaccinationAvailable bool `json:"termineVorhanden"`
}

const (
	vaccinationsUrl = "https://001-iz.impfterminservice.de/assets/static/its/vaccination-list.json"
	vaccinationCentersUrl = "https://www.impfterminservice.de/assets/static/impfzentren.json"
	appointmentCheckUrlPart  = "rest/suche/termincheck?plz=%v&leistungsmerkmale=%v"
	appointmentLinkPart   = "impftermine/service?plz=%v"
)

func main() {

	// the shared http client
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	// the CLI flags
	vaccineCode := flag.String("vaccineCode", "", "Vaccine Code (L920, L921, L922)")
	zipCodes := flag.String("zipCodes", "", "Zip Codes (comma separated, no-spaces)")
	listVaccines := flag.Bool("listVaccines", false, "List vaccine details & codes")
	flag.Parse()

	if *listVaccines {
		vaccines, err := getVaccines(&httpClient)
		if err != nil {
			panic(err)
		}

		fmt.Print(vaccines)
		os.Exit(1)
	}

	if *vaccineCode == "" || *zipCodes == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// we only care about zip codes, so convert the received json response to a map of locations using zipCode as key
	vaccinationCentersByZip, err := getVaccinationCentersByZip(&httpClient)
	if err != nil {
		panic(err)
	}

	// handle available locations (print it to the console)
	availableLocations := make(chan VaccinationCenter, 0)
	go func() {
		for availableLoc := range availableLocations {
			// print location details
			fmt.Printf("\n%v, %v %v - %v: \n", availableLoc.Zip, availableLoc.Address, availableLoc.City, availableLoc.Name)
			// print url to make an appointment
			fmt.Printf(availableLoc.BaseUrl+appointmentLinkPart+"\n\n", availableLoc.Zip)
		}
	}()

	// asynchronously query all locations
	var wg sync.WaitGroup
	for _, zipCode := range strings.Split(*zipCodes, ",") {

		// each zip code can have multiple vaccination centers, so we also check each of them
		for _, vaccinationCenter := range vaccinationCentersByZip[zipCode] {

			wg.Add(1)
			go func(vc string, loc *VaccinationCenter) {
				defer wg.Done()

				req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(loc.BaseUrl+appointmentCheckUrlPart, loc.Zip, vc), nil)
				if err != nil {
					panic(err)
				}

				addHeaders(req)

				// call the service
				resp, err := httpClient.Do(req)
				if err != nil {
					panic(err)
				}

				// extract the response
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					panic(err)
				}

				// unmarshal the json response to go structs
				var response LocationResponse
				err = json.Unmarshal(body, &response)
				if err != nil {
					panic(err)
				}

				// if there are appointments available, pass it to the channel for processing
				if response.VaccinationAvailable {
					availableLocations <- *loc
				}

			}(*vaccineCode, &vaccinationCenter)
		}
	}
	// wait until all locations are queried
	wg.Wait()

	time.Sleep(1 * time.Second)
	fmt.Println("Finished searching.")
}

func getVaccines(httpClient *http.Client) (string, error) {
	req, err := http.NewRequest(http.MethodGet, vaccinationsUrl, nil)
	if err != nil {
		return "", err
	}

	addHeaders(req)

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func getVaccinationCentersByZip(httpClient *http.Client) (map[string][]VaccinationCenter, error) {

	// create the map
	vaccinationCenterByZip := make(map[string][]VaccinationCenter)

	// get the vaccination centers
	req, err := http.NewRequest(http.MethodGet, vaccinationCentersUrl, nil)
	if err != nil {
		panic(err)
	}

	addHeaders(req)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rawResponse map[string]json.RawMessage
	err = json.Unmarshal(body, &rawResponse)
	if err != nil {
		return nil, err
	}

	vaccinationCentersByState := make(map[string][]VaccinationCenter)
	for k, v := range rawResponse {
		var vaccinationCenter []VaccinationCenter
		err = json.Unmarshal(v, &vaccinationCenter)
		if err != nil {
			return nil, err
		}

		vaccinationCentersByState[k] = vaccinationCenter
	}

	// convert it to a map based on zip code as key (multiple per zip possible (ex: 77656 Offenburg)
	for _, l := range vaccinationCentersByState {
		for _, v := range l {
			vaccinationCenterByZip[v.Zip] = append(vaccinationCenterByZip[v.Zip], v)
		}
	}

	return vaccinationCenterByZip, nil
}

// the REST service appears to need these headers or it just times out
func addHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux i686; rv:86.0) Gecko/20100101 Firefox/86.0")
	req.Header.Set("Accept", "application/json")
}
