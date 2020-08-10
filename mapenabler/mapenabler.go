/// Package mapenabler is a simple package that enables maps(location maps) based on the given parameters

package mapenabler

import (
	"assignment/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

// Init all required stuff here.
func Init() {
	defer glog.Flush()
}

// MapEnabler is to enable map based on the given data
type MapEnabler struct {
	APIKey     string   // ApiKey Key provided by the Here or any maps
	Categories []string // Category of the places example petrol stations , restaurants , parking-facility etc..
	BaseURL    string
	PlacesURI  string
	Size       int
}

// New creates a new Channel enabler
func New(apiKey, baseURL, placesURI string, size int, categories ...string) (*MapEnabler, error) {
	if apiKey == "" || baseURL == "" || placesURI == "" || size <= 0 {
		glog.Errorln("one or more input parameters are not provided")
		return nil, errors.New("one or more input parameters are not provided")
	}
	me := &MapEnabler{
		APIKey:     apiKey,     //"NQLeBf6xcolqAFhQyex0sHeAILpgHqSdTT45i1ahPdI",
		BaseURL:    baseURL,    //"https://places.ls.hereapi.com",
		Categories: categories, //"petrol-station", //[]string{"petrol-station", "parking-facility", "restaurant"},
		PlacesURI:  placesURI,  //"/places/v1/discover/explore?",
		Size:       size,       //1,
	}
	return me, nil
}

// FetchMapsDataWithChan is to fetch data provided by URI and parameters
func (m *MapEnabler) FetchMapsDataWithChan(location, category string, chanResult chan<- *models.Result) {
	result := &models.Result{}
	fullRequestURI := fmt.Sprintf("%sat=%s&cat=%s&apikey=%s&size=%d", m.BaseURL+m.PlacesURI, location, category, m.APIKey, m.Size)
	glog.Info("Fetching data from maps server. The get request is on the below URL\n", fullRequestURI)
	response, err := http.Get(fullRequestURI)
	if err != nil {
		chanResult <- nil
		glog.Errorln(err)
		return
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		chanResult <- nil
		glog.Errorln(err)
		return
	}
	if response.StatusCode != 200 {
		chanResult <- nil
		glog.Errorln(string(data))
		return
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		chanResult <- nil
		glog.Errorln(err)
		return
	}
	glog.Info("Result is fetched from maps server and assigned to the channel")
	chanResult <- result
}

// FetchMapsData is to fetch data provided by URI and parameters
func (m *MapEnabler) FetchMapsData(location, category string) (result *models.Result, err error) {
	result = &models.Result{}
	fullRequestURI := fmt.Sprintf("%sat=%s&cat=%s&apikey=%s&size=%d", m.BaseURL+m.PlacesURI, location, category, m.APIKey, m.Size)
	fmt.Println(fullRequestURI)
	response, err := http.Get(fullRequestURI)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
