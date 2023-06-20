package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/valyala/fasthttp"
)

var headerContentTypeJson = []byte("application/json")

var client *fasthttp.Client


func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	client = &fasthttp.Client{}
	return nil, diags
}


func get(url string) (map[string]string, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)

	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)

	code := uint32(resp.StatusCode())
	if code != fasthttp.StatusOK {
		return nil, fmt.Errorf("response code %d", code)
	}

	respBody := resp.Body()
	var resData map[string]string
	json.Unmarshal(respBody, &resData)

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)
	return resData, err
}

func Post(url string, data map[string]string) (map[string]string, error) {
	var x = schema.ResourceData{}
	var c = context.TODO()

	providerConfigure(c, &x)
	return post(url, data)
}


func post(url string, data map[string]string) (map[string]string, error) {
	var dataBytes, _ = json.Marshal(data)
	
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes(headerContentTypeJson)
	req.SetBodyRaw(dataBytes)

	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)

	code := uint32(resp.StatusCode())
	if code != fasthttp.StatusOK {
		return nil, fmt.Errorf("response code %d", code)
	}

	respBody := resp.Body()
	var resData map[string]string
	json.Unmarshal(respBody, &resData)

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)
	return resData, err
}
