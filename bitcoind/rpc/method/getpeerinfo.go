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
	"github.com/dustin/go-humanize"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/rvelhote/bitcoind-status/bitcoind/rpc"
	"github.com/rvelhote/timestamp-marshal"
	"log"
	"math"
	"net"
	"strconv"
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

	// Humanize contains humanized data from the raw data in this
	Humanize PeerInfoHumanize
}

type PeerInfoHumanize struct {
	Hostname  string
	BytesSent string
	BytesRecv string
	ConnTime  string
	Services  []string
	PingTime  string
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

const (
	// Nothing
	NODE_NONE uint64 = 0

	// NODE_NETWORK means that the node is capable of serving the block chain. It is currently
	// set by all Bitcoin Core nodes, and is unset by SPV clients or other peers that just want
	// network services but don't provide them.
	NODE_NETWORK uint64 = 1 << 0

	// NODE_GETUTXO means the node is capable of responding to the getutxo protocol request.
	// Bitcoin Core does not support this but a patch set called Bitcoin XT does.
	// See BIP 64 for details on how this is implemented.
	NODE_GETUTXO uint64 = 1 << 1

	// NODE_BLOOM means the node is capable and willing to handle bloom-filtered connections.
	// Bitcoin Core nodes used to support this by default, without advertising this bit,
	// but no longer do as of protocol version 70011 (= NO_BLOOM_VERSION)
	NODE_BLOOM uint64 = 1 << 2

	// NODE_WITNESS indicates that a node can be asked for blocks and transactions including
	// witness data.
	NODE_WITNESS uint64 = 1 << 3

	// NODE_XTHIN means the node supports Xtreme Thinblocks
	// If this is turned off then the node will not service nor make xthin requests
	NODE_XTHIN uint64 = 1 << 4
)

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

// HasServiceFlag checks the services flag sent by a peer and determines if a certain flag is set.
func HasServiceFlag(services string, flag uint64) bool {
	dec, _ := strconv.ParseUint(services, 16, 64)
	return (dec & flag) != 0
}

func Services(flag string) []string {
	available := []string{}

	if HasServiceFlag(flag, NODE_NONE) {
		available = append(available, "NODE_NONE")
	}

	if HasServiceFlag(flag, NODE_NETWORK) {
		available = append(available, "NODE_NETWORK")
	}

	if HasServiceFlag(flag, NODE_GETUTXO) {
		available = append(available, "NODE_GETUTXO")
	}

	if HasServiceFlag(flag, NODE_BLOOM) {
		available = append(available, "NODE_BLOOM")
	}

	if HasServiceFlag(flag, NODE_WITNESS) {
		available = append(available, "NODE_WITNESS")
	}

	if HasServiceFlag(flag, NODE_XTHIN) {
		available = append(available, "NODE_XTHIN")
	}

	return available
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

	for i, peer := range result {
		result[i].Humanize.Hostname = peer.Addr // , _ = Hostname(peer.Addr)
		result[i].Humanize.BytesRecv = humanize.Bytes(peer.BytesRecv)
		result[i].Humanize.BytesSent = humanize.Bytes(peer.BytesSent)
		result[i].Humanize.ConnTime = humanize.Time(peer.ConnTime.Time)
		result[i].Humanize.Services = Services(peer.Services)
		result[i].Humanize.PingTime = strconv.FormatInt(int64(math.Ceil(peer.PingTime*1000)), 10)
	}

	return result, nil
}
