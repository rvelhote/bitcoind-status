// Package bitcoind is a package
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
	"github.com/rvelhote/timestamp-marshal"
	"log"
	"net"
)

type PeerInfoArgs struct {
}

type PeerInfo struct {
	Id             int               `json:"id"`
	Addr           string            `json:"addr"`
	AddrLocal      string            `json:"addrlocal"`
	Services       string            `json:"services"`
	RelayTxes      bool              `json:"relaytxes"`
	LastSend       timestamp.Unix    `json:"lastsend"`
	LastRecv       timestamp.Unix    `json:"lastrecv"`
	BytesSent      int               `json:"bytessent"`
	BytesRecv      int               `json:"bytesrecv"`
	ConnTime       timestamp.Unix    `json:"conntime"`
	TimeOffset     int               `json:"timeoffset"`
	PingTime       float32           `json:"pingtime"`
	MinPing        float64           `json:"minping"`
	Version        uint              `json:"version"`
	Subver         string            `json:"subver"`
	Inbout         bool              `json:"inbound"`
	AddNode        bool              `json:"addnode"`
	StartHeight    int               `json:"startingheight"`
	BanScore       uint              `json:"banscore"`
	SynchedHeaders int               `json:"synced_headers"`
	SynchedBlocks  int               `json:"synced_blocks"`
	Whitelisted    bool              `json:"whitelisted"`
	BytesReceived  PeerInfoBytesSent `json:"bytessent_per_msg"`
	BytesSend      PeerInfoBytesSent `json:"bytesrecv_per_msg"`
	Hostname       string
}

type PeerInfoBytesSent struct {
	Addr        uint `json:"addr"`
	FeeFilter   uint `json:"feefilter"`
	Inv         uint `json:"inv"`
	Ping        uint `json:"ping"`
	Pont        uint `json:"pong"`
	SendCmpct   uint `json:"sendcmpct"`
	SendHeaders uint `json:"sendheaders"`
	VerAck      uint `json:"verack"`
	Version     uint `json:"version"`
}

type PeerInfoBytesReceived struct {
	Addr    uint `json:"addr"`
	Inv     uint `json:"inv"`
	Ping    uint `json:"ping"`
	Pong    uint `json:"pong"`
	VerAck  uint `json:"verack"`
	Version uint `json:"version"`
}

// hostname fetches the hostname associated to the peer's ip address and sets it automatically (just call the func).
func (p *PeerInfo) hostname() error {
	host, _, _ := net.SplitHostPort(p.Addr)
	names, err := net.LookupAddr(host)

	if err != nil {
		return err
	}

	if len(names) == 0 {
		log.Println("Names not found for " + p.Addr)
		return nil
	}

	p.Hostname = names[0]
	return nil
}

func GetPeerInfo(client *rpc.RPCClient) ([]PeerInfo, error) {
	response, err := client.Post("getpeerinfo", PeerInfoArgs{})

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var result []PeerInfo
	err = json2.DecodeClientResponse(response.Body, &result)

	if err != nil {
		return nil, err
	}

	for i, _ := range result {
		result[i].hostname()
	}

	return result, nil
}
