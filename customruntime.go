package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	apiclient "github.com/arcanericky/petstoreclient/client"
	runtimeclient "github.com/go-openapi/runtime/client"
)

type customRuntime struct {
	runtime             *runtimeclient.Runtime
	defaultRoundTripper http.RoundTripper
	body                string
}

func (cr *customRuntime) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := cr.defaultRoundTripper.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cr.body = string(responseBytes)

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	resp.Body = ioutil.NopCloser(bytes.NewReader(responseBytes))

	return resp, nil
}

func newCustomRuntime() *customRuntime {
	cr := new(customRuntime)

	cr.runtime = runtimeclient.New(apiclient.DefaultTransportConfig().Host,
		apiclient.DefaultTransportConfig().BasePath,
		apiclient.DefaultTransportConfig().Schemes)

	cr.defaultRoundTripper = cr.runtime.Transport
	cr.runtime.Transport = cr

	return cr
}
