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
	"time"
)

// Uptime consults the "uptime" RPC method to obtain how long (in seconds) has the Bitcoin server been running.
// Note that this method is only available in rvelhote:rpc-uptime branch and hopefully it will be merged for the next
// Bitcoin Core release.
func Uptime(client *rpc.RPCClient) (time.Duration, error) {
	response, err := client.Post("uptime", PeerInfoArgs{})

	if err != nil {
		return time.Duration(0), err
	}

	defer response.Body.Close()

	var result int64
	err = json2.DecodeClientResponse(response.Body, &result)

	if err != nil {
		return time.Duration(0), err
	}

	return time.Duration(result), nil
}
