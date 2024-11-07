package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type ResponseXe struct {
	Timestamp int64              `json:"timestamp"`
	Rates     map[string]float64 `json:"rates"`
}

type Api struct {
	Client *http.Client
}

func (a *Api) Rates() (bool, interface{}) {
	req, _ := http.NewRequest("GET", "https://www.xe.com/api/protected/midmarket-converter/", nil)
	req.Header.Set("Authorization", "Basic bG9kZXN0YXI6cHVnc25heA==")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")

	resp, err := a.Client.Do(req)

	if err != nil {
		return false, map[string]any{"status": false, "error": err.Error()}
	}

	defer resp.Body.Close()

	data_resp, err := io.ReadAll(resp.Body)

	if err != nil {
		return false, map[string]any{"status": false, "error": err.Error()}
	} else if resp.StatusCode != 200 {
		return false, map[string]any{"status": false, "error": string(data_resp)}
	}

	var data ResponseXe

	json.Unmarshal(data_resp, &data)

	return true, data
}
