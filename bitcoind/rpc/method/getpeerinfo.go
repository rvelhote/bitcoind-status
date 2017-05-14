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
	// Id is the Peer index
	Id int `json:"id"`

	// Addr is the ip address and port of the peer
	Addr string `json:"addr"`

	// AddrLocal is the local address
	AddrLocal string `json:"addrlocal"`

	// Services is an hexadecimal value that describes services offered by the peer
	Services string `json:"services"`

	// RelayTxes indicates whether peer has asked us to relay transactions to it
	RelayTxes bool `json:"relaytxes"`

	// LastSend is the time in seconds since epoch (Jan 1 1970 GMT) of the last send
	LastSend timestamp.Unix `json:"lastsend"`

	// LastRecv is the time in seconds since epoch (Jan 1 1970 GMT) of the last receive
	LastRecv timestamp.Unix `json:"lastrecv"`

	// BytesSent is the total bytes sent to this peer
	BytesSent uint64 `json:"bytessent"`

	// BytesRecv is the total bytes received from this peer
	BytesRecv uint64 `json:"bytesrecv"`

	// ConnTime is the connection time in seconds since epoch (Jan 1 1970 GMT)
	ConnTime timestamp.Unix `json:"conntime"`

	// TimeOffset is the time offset in seconds
	TimeOffset int `json:"timeoffset"`

	// PingTime is the ping to this peer (if available)
	PingTime float64 `json:"pingtime"`

	// MinPing is the minimum observed ping time (if any at all)
	MinPing float64 `json:"minping"`

	// Version is the peer version, such as 7001
	Version uint `json:"version"`

	// Subver is the string version of this peer
	Subver string `json:"subver"`

	// Inbound indicates if the connection is Inbound (true) or Outbound (false)
	Inbound bool `json:"inbound"`

	// AddNode indicates whether connection was due to addnode and is using an addnode slot
	AddNode bool `json:"addnode"`

	// StartHEight The starting height (block) of the peer
	StartHeight int `json:"startingheight"`

	// BanScore indicates the ban score for this peer
	BanScore uint `json:"banscore"`

	// SynchedHeaders is the last header we have in common with this peer
	SynchedHeaders int `json:"synced_headers"`

	// SynchedBlocks is the last block we have in common with this peer
	SynchedBlocks int `json:"synced_blocks"`

	// Whitelisted indicates whether the peer is whitelisted
	Whitelisted bool `json:"whitelisted"`

	// BytesSentPerMsg is the total bytes sent aggregated by message type
	BytesSentPerMsg PeerInfoBytesSent `json:"bytessent_per_msg"`

	// BytesReceivedPerMsg is the total bytes received aggregated by message type
	BytesReceivedPerMsg PeerInfoBytesSent `json:"bytesrecv_per_msg"`
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
func Hostname(ipaddress string) (string, error) {
	host, _, _ := net.SplitHostPort(ipaddress)
	names, err := net.LookupAddr(host)

	if err != nil {
		return "", err
	}

	if len(names) == 0 {
		log.Println("Names not found for " + ipaddress)
		return "", nil
	}

	return names[0], nil
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

	return result, nil
}
