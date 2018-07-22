package history

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type BaseRespType struct {
	ResponseHeader struct {
		SearchServiceQTime int `json:"search_service_q_time"`
		Params struct {
			Q         []string
			MediaType []string `json:"media_type"`
			Namespace []string `json:"namespace"`
			Start     []string `json:"start"`
			Fq        []string `json:"fq"`
			Source    []string `json:"source"`
			Rows      []string `json:"params"`
			SortOrder []string `json:"sort_order"`
		}
	}
	Total int
	Data []struct {
		MediaID string `json:"media_id"`
	}
}

type ImageFormatType struct {
	DisplayName  string `json:"display_name"`
	DPI          int    `json:"dpi"`
	FileSize     int    `json:"file_size"`
	Format       string `json:"format"`
	Height       int    `json:"height"`
	IsLicensable bool   `json:"is_licensable"`
	Width        int    `json:"width"`
}

type ImageLinksType struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type DataImageType struct {
	Data []struct {
		ID        string  `json:"id"`
		AddedDate string  `json:"added_date"`
		Aspect    float64 `json:"aspect"`
		Assets struct {
			SmallJpg      ImageFormatType `json:"small_jpg"`
			MediumJpg     ImageFormatType `json:"medium_jpg"`
			HugeJpg       ImageFormatType `json:"huge_jpg"`
			SupersizeJpg  ImageFormatType `json:"supersize_jpg"`
			HugeTiff      ImageFormatType `json:"huge_tiff"`
			SupersizeTiff ImageFormatType `json:"supersize_tiff"`
			Preview       ImageLinksType  `json:"preview"`
			SmallThumb    ImageLinksType  `json:"small_thumb"`
			LargeThumb    ImageLinksType  `json:"large_thumb"`
			HugeThumb     ImageLinksType  `json:"huge_thumb"`
		}
		Categories []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}
		Contributor struct {
			ID string `json:"id"`
		}
		Description        string   `json:"description"`
		ImageType          string   `json:"image_type"`
		IsAdult            bool     `json:"is_adult"`
		IsIllustration     bool     `json:"is_illustration"`
		HasPropertyRelease bool     `json:"has_property_release"`
		Keywords           []string `json:"keywords"`
		MediaType          string   `json:"media_type"`
	}
}

//const (
//	BaseURL = "https://submit.shutterstock.com/api/catalog_manager/media_types/all/items?filter_type=date_uploaded&filter_value=%s-%s-01%s%s-%s-%s&page_number=%s&per_page=%s&sort=popular"
//	ApiURL  = "https://api.shutterstock.com/v2/images"
//	Session = "s%3AFLsDQ0KkRmbbHJSFijJz_5VxQPCQI7Ol.t5LQWhFeOPA9qV2S0fqa6JBsFB0Rq%2BrxMDPc1URXyHE"
//	Token   = "v2/NjYxMWYtZjk4NjItNWI0ZDEtYjc3ODktNmIyYWEtZWU3NDUvMTEyNDM5ODQ1L2N1c3RvbWVyLzMveEZSaF85amREWHFaOFZuNjZfcjZVaUhWYzlYZmhVVl9kZURtU1EybHBSbDJ0eDFxUXh5czEzUXdRdkV6Wjl6dFBYcXpGbHk4RWhWZUpxU1U4TlFoWll3RjF1YkZKMFhsS0FndDQtSTY5Y0k0TV9nYy0yVFEzLXdzeC02TXdXMlloVFQyQ2kwZzBjZmtsVmNNVE5OUjZCSHdNY1kzSUQ4SW1CZHlwT1dYNFQ5enNIUkFPdUh3VElPRmxaZ214a003dDc1UHE0Tmpta0tIN3ZYM3g4V2xlUQ"
//)

func (h *TypeApprovedHistory) BaseRequest(page string, current *TypeCurrDate) (*BaseRespType, error) {

	newURL := fmt.Sprintf(h.Config.Baseurl, current.Year, current.Month, "%20", current.Year, current.Month, current.Day, page, "100")

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

	res, errRequest := h.NewRequest(fmt.Sprintf("%s?%s", h.Config.Apiurl, query))
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

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", h.Config.Token))
	cookie := http.Cookie{Name: "session", Value: h.Config.Session}
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
