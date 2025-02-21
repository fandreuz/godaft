package godaft

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/valyala/fastjson"
)

const endpoint string = "https://gateway.daft.ie/old/v1/listings"

func DoRequest(payload PayloadType) (*fastjson.Value, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, endpoint, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("brand", "daft")
	req.Header.Set("platform", "web")

	httpResult, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if httpResult.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP status code: %d", httpResult.StatusCode)
	}

	resultBody, err := io.ReadAll(httpResult.Body)
	if err != nil {
		return nil, err
	}

	var p fastjson.Parser
	return p.ParseBytes(resultBody)
}
