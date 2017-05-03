// Package rpc is a package
package rpc

/*
 * The MIT License (MIT)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
import (
	"bytes"
	"github.com/gorilla/rpc/v2/json2"
	"net/http"
	"time"
)

type RPCClient struct {
	url      string
	username string
	password string
	client   *http.Client
}

func NewRPCClient(url, user string, password string) *RPCClient {
	return &RPCClient{
		url,
		user,
		password,
		&http.Client{Timeout: 10 * time.Second},
	}
}

func (r *RPCClient) Post(method string, args interface{}) (*http.Response, error) {
	encoded, err := json2.EncodeClientRequest(method, args)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", r.url, bytes.NewBuffer(encoded))

	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(r.username, r.password)
	response, err := r.client.Do(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}
