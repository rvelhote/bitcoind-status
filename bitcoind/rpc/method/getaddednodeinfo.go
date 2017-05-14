// Package method is a package
package method

import (
	"github.com/gorilla/rpc/v2/json2"
	"github.com/rvelhote/bitcoind-status/bitcoind/rpc"
)

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
type AddedNodeInfo struct {
	// AddedNode is the node ip address or name (as provided to addnode)
	AddedNode string `json:"addednode"`
	// Connected is a boolean that specifies if we are currently connected to the node
	Connected bool `json:"connected"`
	// Addresses is a list of addresses of the added node. Only present if connected == true
	Addresses []AddedNodeAddress `json:"addresses"`
}

type AddedNodeAddress struct {
	// Address is the bitcoin server IP and port we're connected to
	Address string `json:"address"`
	// Connected specified wether the connection is inbound or outbound
	Connected string `json:"connected"`
}

// GetAddedNodeInfo Returns information about the given added node, or all added nodes.
func GetAddedNodeInfo(client *rpc.RPCClient) ([]AddedNodeInfo, error) {
	response, err := client.Post("getaddednodeinfo", PeerInfoArgs{})

	if err != nil {
		return []AddedNodeInfo{}, err
	}

	defer response.Body.Close()

	var result []AddedNodeInfo
	err = json2.DecodeClientResponse(response.Body, &result)

	if err != nil {
		return []AddedNodeInfo{}, err
	}

	return result, nil
}
