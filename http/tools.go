package http

import (
	"bytes"
	"context"
	"errors"
	"gitlab.tools.vipshop.com/terry.zeng/golang-mq/g"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

/*
func FalconToken() (string) {

	// crate falcon api header token access

	token, err := json.Marshal(map[string]string{"name": g.FalconAuthName, "sig": g.FalconAuthSig})

	if err != nil {
		log.Println(err)
	}

	return  string(token)
}

*/

func HttpApiPut(fullApiUrl string, jsonData []byte) (status bool, err error) {

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, fullApiUrl, bytes.NewBuffer(jsonData))

	if err != nil {
		return false, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")


	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return false, err
	}

	if  ( resp.StatusCode  == 200 ) {
		return true, nil
	} else {
		return false, errors.New("[ERROR] HttpApiPut() response not 200")
	}

}

func HttpApiGet(fullApiUrl string, params string) (io.ReadCloser, error) {

	client := &http.Client{}
	httpUrl := fullApiUrl + params

	req, err := http.NewRequest("GET", httpUrl, nil)

	if err != nil {
		log.Println(err)
		return nil, errors.New("HttpApiGet() http get error with NewRequest")
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println("[ERROR] HttpApiGet() ", fullApiUrl, err)
		return nil, errors.New("HttpApiGet() http get error")
	}

	if ( resp.StatusCode  == 200 ) {
		return resp.Body, nil
	} else {
		return nil, errors.New("HttpApiGet() resp status code not 200.")
	}

}

func HttpApiPost(fullApiUrl string, params []byte) (io.ReadCloser, error) {
	// use to access http post
	// params = post params  [must be []byte format]
	// return http response


	tr := &http.Transport{
		MaxIdleConns: 10,
		IdleConnTimeout: 10 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", fullApiUrl, bytes.NewBuffer(params))
	ctx, cancelFunc := context.WithCancel(context.Background())
	request := req.WithContext(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New("HttpApiPost() http post error with NewRequest")
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(request)

	if err != nil {
		return nil, errors.New("HttpApiPost()  client access error.")
	}
	defer cancelFunc()
	if resp.Body != nil {
		defer resp.Body.Close()
	} else {
		return nil, err
	}

	if ( resp.StatusCode  == 200 ) {
		return resp.Body, nil
	} else {
		if g.Debug == true {

			b, err := httputil.DumpResponse(resp, true)
			if err != nil {
				log.Fatalln("[ERROR] HttpApiPost() dump with ", err)
			}

			log.Println("[ERROR] HttpApiPost() ", string(b))
		}
		return nil, errors.New("HttpApiPost() resp status code not 200.")
	}
}

func HttpApiDelete(fullApiUrl string, params string) (io.ReadCloser, error) {
	// use to do http Delete request
	// METHOD: DELETE

	client := &http.Client{}
	httpUrl := fullApiUrl + params
	req, err := http.NewRequest("DELETE", httpUrl, nil)

	if err != nil {
		log.Println(err)
		return nil, errors.New("HttpApiDelete() http delete error with NewRequest")
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Println("[ERROR] HttpApiDelete() ", fullApiUrl, err)
		return nil, errors.New("HttpApiDelete() http delete error")
	}

	if ( resp.StatusCode  == 200 ) {
		return resp.Body, nil
	} else {
		return nil, errors.New("HttpApiDelete() resp status code not 200.")
	}
}