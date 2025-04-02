package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const baconIpsumURL = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

type BeefSummaryResponse struct {
	Beef map[string]int `json:"beef"`
}

func (h *Handler) HandleBeefSummary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text, err := FetchBaconIpsumText()
	if err != nil {
		log.Printf("Error fetching text from Bacon Ipsum API: %v", err)
		http.Error(w, "Failed to fetch meat text", http.StatusInternalServerError)
		return
	}

	counts := h.counter.GetBeefCounts(text)

	response := BeefSummaryResponse{
		Beef: counts,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func FetchBaconIpsumText() (string, error) {
	res, err := http.Get(baconIpsumURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
