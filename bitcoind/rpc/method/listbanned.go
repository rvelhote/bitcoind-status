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
	"github.com/dustin/go-humanize"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/rvelhote/bitcoind-status/bitcoind/rpc"
	"github.com/rvelhote/timestamp-marshal"
)

type Banned struct {
	Address     string         `json:"address"`
	BannedUntil timestamp.Unix `json:"banned_until"`
	BanCreated  timestamp.Unix `json:"ban_created"`
	BanReason   string         `json:"ban_reason"`
	Humanized   HumanizeBanned
}

type HumanizeBanned struct {
	Address     string
	BannedUntil string
	BanCreated  string
}

func ListBanned(client *rpc.RPCClient) ([]Banned, error) {
	response, err := client.Post("listbanned", PeerInfoArgs{})

	if err != nil {
		return []Banned{}, err
	}

	defer response.Body.Close()

	var result []Banned
	err = json2.DecodeClientResponse(response.Body, &result)

	if err != nil {
		return []Banned{}, err
	}

	for i, peer := range result {
		result[i].Humanized.Address, _ = Hostname(peer.Address)
		result[i].Humanized.BanCreated = humanize.Time(peer.BanCreated.Time)
		result[i].Humanized.BannedUntil = humanize.Time(peer.BannedUntil.Time)
	}

	return result, nil
}
