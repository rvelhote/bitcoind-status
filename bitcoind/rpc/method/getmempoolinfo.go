// Package method is a package
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

type MempoolInfo struct {
	// Size is the current transaction count.
	Size uint `json:"size"`

	// Bytes is the sum of all virtual transaction sizes as defined in BIP 141. Differs from actual serialized size
	// because witness data is discounted
	Bytes uint64 `json:"bytes"`

	// Usage is the total memory usage for the mempool.
	Usage uint64 `json:"usage"`

	// MaxMempool is the maximum memory usage for the mempool.
	MaxMempool uint64 `json:"maxmempool"`

	// MempoolMinFee is the minimum fee for transaction to be accepted into the mempool.
	MempoolMinFee float64 `json:"mempoolminfee"`

	// Humanize humanizes the units (bytes, timestamps) that belong to this structure.
	Humanize HumanizedMempoolInfo
}

type HumanizedMempoolInfo struct {
	Bytes      string
	Usage      string
	MaxMempool string
}

// GetMempoolInfo returns details on the active state of the TX memory pool.
func GetMempoolInfo(client *rpc.RPCClient) (MempoolInfo, error) {
	response, err := client.Post("getmempoolinfo", PeerInfoArgs{})

	if err != nil {
		return MempoolInfo{}, err
	}

	defer response.Body.Close()

	var result MempoolInfo
	err = json2.DecodeClientResponse(response.Body, &result)

	if err != nil {
		return MempoolInfo{}, err
	}

	return result, nil
}
