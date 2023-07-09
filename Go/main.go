package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type EmbedConfig struct {
	DashboardId    string `json:"DashboardId"`
	ServerUrl      string `json:"ServerUrl"`
	UserEmail      string `json:"UserEmail"`
	EmbedSecret    string `json:"EmbedSecret"`
	EmbedType      string `json:"EmbedType"`
	Environment    string `json:"Environment"`
	ExpirationTime string `json:"ExpirationTime"`
	SiteIdentifier string `json:"SiteIdentifier"`
}

// Create an instance of EmbedConfig struct
var config EmbedConfig

func main() {
	// Read the embedConfig.json file
	data, err := ioutil.ReadFile("embedConfig.json")
	if err != nil {
		log.Fatal("Error: embedConfig.json file not found.")
	}

	// Unmarshal the JSON data into the config struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	http.HandleFunc("/authorizationserver", AuthorizationServer)
	http.HandleFunc("/getServerDetails", getServerDetails)
	log.Fatal(http.ListenAndServe(":8086", nil))
}

func getServerDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	data, err := ioutil.ReadFile("embedConfig.json")
	if err != nil {
		log.Fatal("Error reading embedConfig.json:", err)
	}

	err = json.Unmarshal(data, &config)
	response, err := json.Marshal(config)
	w.Write(response)
}

func AuthorizationServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	if len(body) > 0 {
		if queryString, err := unmarshal(string(body)); err != nil {
			log.Println("error converting", err)
		} else {
			serverAPIUrl := queryString.(map[string]interface{})["dashboardServerApiUrl"].(string)
			embedQueryString := queryString.(map[string]interface{})["embedQuerString"].(string)
			embedQueryString += "&embed_user_email=" + config.UserEmail
			timeStamp := time.Now().Unix()
			embedQueryString += "&embed_server_timestamp=" + strconv.FormatInt(timeStamp, 10)
			signatureString, err := getSignatureUrl(embedQueryString)
			embedDetails := "/embed/authorize?" + embedQueryString + "&embed_signature=" + signatureString
			query := serverAPIUrl + embedDetails
			log.Println(query)
			result, err := http.Get(query)
			if err != nil {
				log.Println(err)
			}
			log.Println(result)
			response, err := ioutil.ReadAll(result.Body)
			if err != nil {
				log.Fatalln(err)
			}
			w.Write(response)
		}
	}
}

func getSignatureUrl(queryData string) (string, error) {
	encoding := ([]byte(config.EmbedSecret))
	messageBytes := ([]byte(queryData))
	hmacsha1 := hmac.New(sha256.New, encoding)
	hmacsha1.Write(messageBytes)
	sha := base64.StdEncoding.EncodeToString(hmacsha1.Sum(nil))
	return sha, nil
}

func unmarshal(data string) (interface{}, error) {
	var iface interface{}
	decoder := json.NewDecoder(strings.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(&iface); err != nil {
		return nil, err
	}
	return iface, nil
}
