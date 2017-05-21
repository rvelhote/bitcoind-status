package humanize

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
	"github.com/patrickmn/go-cache"
	"net"
	"time"
)

// hostnameCache stores the resolved hostname of IP Addresses for faster loading.
var hostnameCache *cache.Cache = cache.New(5*time.Minute, 10*time.Minute)

// Hostname fetches the hostname associated to the peer's IP Address. If the IP Address has multiple hostnames, the
// first one will be returned. This method, because hostname resolving is a long operation, also caches the found
// hostname.
func Hostname(ipaddress string) string {
    if cached, found := hostnameCache.Get(ipaddress); found {
		return cached.(string)
	}

	host, _, _ := net.SplitHostPort(ipaddress)
	hostnames, err := net.LookupAddr(host)

	if err != nil {
		return ipaddress
	}

	if len(hostnames) == 0 {
		return ipaddress
	}

	hostnameCache.Set(ipaddress, hostnames[0], cache.NoExpiration)
	return hostnames[0]
}
