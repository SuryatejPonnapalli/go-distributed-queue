package llmclient

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"
)


type chatReq struct{
	Text string `json:"text"`
}

type chatResp struct{
	Response string `json:"response"`
}

func GetResponse(prompt string) (string, error){
	llmLimit <- struct{}{}       
	defer func() { <-llmLimit }()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()


	body, _ := json.Marshal(chatReq{Text: prompt})
	chatUrl := os.Getenv("PYTHON_URL") + "/chat"

	req, err := http.NewRequestWithContext(ctx, "POST", chatUrl, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var out chatResp
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil{
		return "",err
	}

	return out.Response, nil
}