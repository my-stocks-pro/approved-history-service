package history

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type BaseRespType struct {
	ResponseHeader struct {
		SearchServiceQTime int
		Params struct {
			Q          []string
			Media_type []string
			Namespace  []string
			Start      []string
			Fq         []string
			Source     []string
			Rows       []string
			Sort_order []string
		}
	}
	Total int
	Data []struct {
		Media_id string
	}
}

type ImageFormatType struct {
	Display_name  string
	DPI           int
	File_size     int
	Format        string
	Height        int
	Is_licensable bool
	Width         int
}

type ImageLinksType struct {
	Height int
	URL    string
	Width  int
}

type DataImageType struct {
	Data []struct {
		ID         string
		Added_date string
		Aspect     float64
		Assets struct {
			Small_jpg      ImageFormatType
			Medium_jpg     ImageFormatType
			Huge_jpg       ImageFormatType
			Supersize_jpg  ImageFormatType
			Huge_tiff      ImageFormatType
			Supersize_tiff ImageFormatType
			Preview        ImageLinksType
			Small_thumb    ImageLinksType
			Large_thumb    ImageLinksType
			Huge_thumb     ImageLinksType
		}
		Categories []struct {
			ID   string
			Name string
		}
		Contributor struct {
			ID string
		}
		Description          string
		Image_type           string
		Is_adult             bool
		Is_illustration      bool
		Has_property_release bool
		Keywords             []string
		media_type           string
	}
}


const (
	BaseURL ="https://submit.shutterstock.com/api/catalog_manager/media_types/all/items?filter_type=date_uploaded&filter_value=%s-%s-01%s%s-%s-%s&page_number=%s&per_page=%s&sort=popular"
	ApiURL = "https://api.shutterstock.com/v2/images"
	Session = "s%3AFLsDQ0KkRmbbHJSFijJz_5VxQPCQI7Ol.t5LQWhFeOPA9qV2S0fqa6JBsFB0Rq%2BrxMDPc1URXyHE"
	Token =  "v2/NjYxMWYtZjk4NjItNWI0ZDEtYjc3ODktNmIyYWEtZWU3NDUvMTEyNDM5ODQ1L2N1c3RvbWVyLzMveEZSaF85amREWHFaOFZuNjZfcjZVaUhWYzlYZmhVVl9kZURtU1EybHBSbDJ0eDFxUXh5czEzUXdRdkV6Wjl6dFBYcXpGbHk4RWhWZUpxU1U4TlFoWll3RjF1YkZKMFhsS0FndDQtSTY5Y0k0TV9nYy0yVFEzLXdzeC02TXdXMlloVFQyQ2kwZzBjZmtsVmNNVE5OUjZCSHdNY1kzSUQ4SW1CZHlwT1dYNFQ5enNIUkFPdUh3VElPRmxaZ214a003dDc1UHE0Tmpta0tIN3ZYM3g4V2xlUQ"
)

func (h *TypeApprovedHistory) BaseRequest(page string, current *TypeCurrDate) (*BaseRespType, error) {

	newURL := fmt.Sprintf(BaseURL, current.Year, current.Month, "%20", current.Year, current.Month, current.Day, page, "100")

	res, errRequest := h.NewRequest(newURL)
	if errRequest != nil {
		return nil, errRequest
	}

	baseResp := BaseRespType{}
	errUnmarshal := json.Unmarshal(res, &baseResp)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &baseResp, nil
}

func (h *TypeApprovedHistory) FullRequest(query string) (*DataImageType, error) {

	res, errRequest := h.NewRequest(fmt.Sprintf("%s?%s", ApiURL, query))
	if errRequest != nil {
		return nil, errRequest
	}

	resp := DataImageType{}
	errUnmarshal := json.Unmarshal(res, &resp)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &resp, nil

}

func (h *TypeApprovedHistory) NewRequest(url string) ([]byte, error) {

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", Token))
	cookie := http.Cookie{Name: "session", Value: Session}
	req.AddCookie(&cookie)

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, err
	}

	body, errReadBody := ioutil.ReadAll(resp.Body)
	if errReadBody != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return body, nil
}