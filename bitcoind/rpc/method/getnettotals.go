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
    "github.com/rvelhote/timestamp-marshal"
    "time"
    "github.com/gorilla/rpc/v2/json2"
    "github.com/rvelhote/bitcoind-status/bitcoind/rpc"
)

// NetTotals is the struct that contains information about the amount of data transfered by our node as well as
// possible restrictions/limitations on the amount of transferred data.
type NetTotals struct {
    TotalBytesRecv uint64 `json:"totalbytesrecv"`
    TotalBytesSent uint64 `json:"totalbytessent"`
    TimeMillis timestamp.Unix `json:"timemillis"`
    UploadTarget UploadTarget `json:"uploadtarget"`
}

type UploadTarget struct {
    Timeframe time.Duration `json:"timeframe"`
    Target uint `json:"target"`
    TargetReached bool `json:"target_reached"`
    ServeHistoricalBlocks bool `json:"serve_historical_blocks"`
    BytesLeftInCycle uint64 `json:"bytes_left_in_cycle"`
    TimeLeftInCycle uint64 `json:"time_left_in_cycle"`
}

func GetNetTotals(client *rpc.RPCClient) (NetTotals, error) {
    response, err := client.Post("getnettotals", PeerInfoArgs{})

    if err != nil {
        return NetTotals{}, err
    }

    defer response.Body.Close()

    var result NetTotals
    err = json2.DecodeClientResponse(response.Body, &result)

    if err != nil {
        return NetTotals{}, err
    }

    return result, nil
}
