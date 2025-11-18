package llmclient

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"
)


var client = &http.Client{
}

type embedReq struct {
    Text string `json:"text"`
}

type embedResp struct {
    Embedding []float64 `json:"embedding"`
}

func GetEmbedding(prompt string) ([]float64, error){
	llmLimit <- struct{}{}       
	defer func() { <-llmLimit }()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	body, _ := json.Marshal(embedReq{Text: prompt})
	embedURL := os.Getenv("PYTHON_URL") + "/embed"

	req, err := http.NewRequestWithContext(ctx, "POST", embedURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out embedResp
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil{
		return nil,err
	}

	return out.Embedding,nil
}