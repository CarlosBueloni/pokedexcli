package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListEncounters(location string) (EncountersResp, error) {
	endpoint := "/location-area/" + location
	fullUrl := baseURL + endpoint

	dat, ok := c.cache.Get(fullUrl)
	if ok {
		//cache hit
		encountersResp := EncountersResp{}
		err := json.Unmarshal(dat, &encountersResp)
		if err != nil {
			return EncountersResp{}, err
		}
		return encountersResp, nil
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return EncountersResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return EncountersResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return EncountersResp{}, fmt.Errorf("bad status error: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return EncountersResp{}, err
	}

	encountersResp := EncountersResp{}
	err = json.Unmarshal(dat, &encountersResp)
	if err != nil {
		return EncountersResp{}, err
	}

	c.cache.Add(fullUrl, dat)

	return encountersResp, nil

}
