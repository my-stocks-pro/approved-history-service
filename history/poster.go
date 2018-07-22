package history

import (
	"net/http"
	"bytes"
	"fmt"
)

//const (
//	postURL = "http://127.0.0.1:8001/data/psql/approved"
//)

func (h *TypeApprovedHistory) Post(data []byte) (*http.Response, error) {

	url := fmt.Sprintf("http://%s:%s/data/psql/approved", h.Config.Host, h.Config.Port)

	req, errReq := http.NewRequest("POST", url, bytes.NewReader(data))
	if errReq != nil {
		return nil, errReq
	}

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}

	return resp, nil
	//fmt.Println(string(data))
	//return nil, nil
}
