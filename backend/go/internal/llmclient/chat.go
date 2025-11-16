package llmclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
)


type chatReq struct{
	Text string `json:"text"`
}

type chatResp struct{
	Response string `json:"response"`
}

func GetResponse(prompt string) (string, error){
	body, _ := json.Marshal(chatReq{Text: prompt})
	chatUrl := os.Getenv("PYTHON_URL") + "/chat"

	resp, err := client.Post(chatUrl, "application/json", bytes.NewBuffer(body))
	if err != nil{
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("llm error: non-200 response")
	}

	var out chatResp
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil{
		return "",err
	}

	return out.Response, nil
}