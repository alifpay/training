package gurl

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

//Req model of http client request
type Req struct {
	PostData []byte
	URL      string
	Method   string
	RespData []byte
	Headers  map[string]string
	InSecure bool
}

//ToByte json marshal to PostData
func (r *Req) ToByte(val interface{}) (err error) {
	r.PostData, err = json.Marshal(val)
	return
}

//Send http client request
func (r *Req) Send() error {

	req, err := http.NewRequest(r.Method, r.URL, bytes.NewReader(r.PostData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	var netClient = &http.Client{
		Timeout: time.Second * 30,
	}

	if r.InSecure {
		netClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		}
	}

	resp, err := netClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.Body == nil {
		return errors.New("empty_http_response")
	}

	r.RespData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

//SendCtx http client request
func (r *Req) SendCtx(ctx context.Context) error {

	req, err := http.NewRequest(r.Method, r.URL, bytes.NewReader(r.PostData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	var netClient = &http.Client{
		Timeout: time.Second * 30,
	}

	if r.InSecure {
		netClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		}
	}
	reqCtx := req.WithContext(ctx)
	resp, err := netClient.Do(reqCtx)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.Body == nil {
		return errors.New("empty_http_response")
	}

	r.RespData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

//ToStruct unmarshal json to struct
func (r *Req) ToStruct(val interface{}) (err error) {
	err = json.Unmarshal(r.RespData, &val)
	return
}
