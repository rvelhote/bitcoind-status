// Package humanize contains funcs that will turn raw data from bitcoind into human readable data
package humanize

import (
	"strconv"
	"strings"
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
const (
	// NODE_NONE mens that the node does not offer any services.
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

// HasServiceFlag checks the services flag sent by a peer and determines if a certain flag is set.
func HasServiceFlag(services string, flag uint64) bool {
	dec, _ := strconv.ParseUint(services, 16, 64)
	return (dec & flag) != 0
}

// ServiceFlag will take the flag returned by the daemon and check which bits are set. Each bit that is set represents
// a service offered by the peer.
func ServiceFlag(flag string) []string {
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

// ServiceFlagJoin is a helper func to display the service flags in a single line.
func ServiceFlagJoin(flag string) string {
	return strings.Join(ServiceFlag(flag), ", ")
}
