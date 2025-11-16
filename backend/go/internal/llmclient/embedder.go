package llmclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)


var client = &http.Client{
	Timeout: 10 * time.Second,
}

type embedReq struct {
    Text string `json:"text"`
}

type embedResp struct {
    Embedding []float64 `json:"embedding"`
}

func GetEmbedding(prompt string) ([]float64, error){
	body, _ := json.Marshal(embedReq{Text: prompt})
	embedURL := os.Getenv("PYTHON_URL") + "/embed"

	resp, err := client.Post(embedURL, "application/json", bytes.NewBuffer(body))
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200{
		return nil, errors.New("embedder error: non-200 response")
	}

	var out embedResp
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil{
		return nil,err
	}

	return out.Embedding,nil
}