package apiHitter

import (
	"io/ioutil"
)

//Hit will sends Http request and return HitResponse
func (p *Payload) Hit() (response HitResponse, err error) {
	if err := p.buildHttpRequest(); err != nil {
		return response, err
	}
	p.buildClient()
	if response.resp, err = p.client.Do(p.req); err != nil {
		return response, err
	}
	response.Data, err = ioutil.ReadAll(response.resp.Body)
	response.Code = response.resp.StatusCode
	return response, err
}
