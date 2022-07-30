package genome

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetGenomeData(variants []string) []Genome {
	postBody, _ := json.Marshal(variants)
	url := "https://www.genomenexus.org/annotation/summary?projection=ALL"

	responseBody := bytes.NewBuffer(postBody)

	response, err := http.Post(url, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	data := []Genome{}
	json.Unmarshal(body, &data)
	fmt.Println(string(body))

	return data
}
