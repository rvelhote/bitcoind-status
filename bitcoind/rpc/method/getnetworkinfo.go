// Package method contains an implemenetation of the RPC methods made available by the Bitcoin server
package method

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
	"github.com/gorilla/rpc/v2/json2"
	"github.com/rvelhote/bitcoind-status/bitcoind/rpc"
)

// NetworkInfo is the structure returned by the Bitcoin Daemon with information about the node's network.
type NetworkInfo struct {
	Version         uint      `json:"version"`
	Subversion      string    `json:"subversion"`
	ProtocolVersion uint      `json:"protocolversion"`
	Services        string    `json:"localservices"`
	LocalRelay      bool      `json:"localrelay"`
	TimeOffset      int       `json:"timeoffset"`
	NetworkActive   bool      `json:"networkactive"`
	Connections     uint      `json:"connections"`
	Warnings        string    `json:"warnings"`
	RelayFee        float64   `json:"relayfee"`
	IncrementalFee  float64   `json:"incrementalfee"`
	LocalAddresses  []string  `json:"localaddresses"`
	Networks        []Network `json:"networks"`
}

// Network is the list of networks that the node is able to connected from (IPv4, IPv6, Onion)
type Network struct {
	Name                      string `json:"name"`
	Limited                   bool   `json:"limited"`
	Reachable                 bool   `json:"reachable"`
	Proxy                     string `json:"proxy"`
	ProxyRandomizeCredentials bool   `json:"proxy_randomize_credentials"`
}

// GetNetworkInfo obtains information about the network conditions of the local node.
func GetNetworkInfo(client *rpc.RPCClient) (NetworkInfo, error) {
	response, err := client.Post("getnetworkinfo", PeerInfoArgs{})

	if err != nil {
		return NetworkInfo{}, err
	}

	defer response.Body.Close()

	var result NetworkInfo
	err = json2.DecodeClientResponse(response.Body, &result)

	if err != nil {
		return NetworkInfo{}, err
	}

	return result, nil
}
