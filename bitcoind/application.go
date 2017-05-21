// Package bitcoind is a package
package bitcoind

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
	humanize2 "github.com/rvelhote/bitcoind-status/bitcoind/humanize"
	"github.com/rvelhote/bitcoind-status/bitcoind/rpc"
	"github.com/rvelhote/bitcoind-status/bitcoind/rpc/method"
	"github.com/rvelhote/bitcoind-status/configuration"
	"html/template"
	"log"
	"net/http"
	"time"
)

// IndexTemplateParams holds various values to be passed to the main template
type IndexTemplateParams struct {
	Title         string
	Peers         []method.PeerInfo
	Network       method.NetworkInfo
	Banned        []method.Banned
	Mempool       method.MempoolInfo
	AddedNodeInfo []method.AddedNodeInfo
	NetTotals     method.NetTotals
	Uptime        time.Duration
}

// IndexRequestHandler handles the requests to present the main url of the application
type IndexRequestHandler struct {
	// Configuration contains the app configuration. In this context only the server list is used.
	Configuration configuration.Configuration
}

// ServeHTTP handles the request made to the homepage of the app. It will only serve the required files to start
// the RectJS app as well as some important configuration.
func (i IndexRequestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	funcs := template.FuncMap{
		"bytes":           humanize.Bytes,
		"time":            humanize.Time,
		"serviceflag":     humanize2.ServiceFlag,
		"serviceflagjoin": humanize2.ServiceFlagJoin,
		"ping":            humanize2.Ping,
		"uptime":          humanize2.Uptime,
		"hostname":        humanize2.Hostname,
	}

	t, err := template.New("index.html").Funcs(funcs).ParseFiles("templates/index.html")
	if err != nil {
		t, _ = template.New("index.html").Funcs(funcs).ParseFiles("../templates/index.html")
	}

	client := rpc.NewRPCClient(i.Configuration.Url, i.Configuration.Username, i.Configuration.Password)

	peerinfo, err := method.GetPeerInfo(client)
	if err != nil {
		log.Fatal(err)
	}

	networkinfo, err := method.GetNetworkInfo(client)
	if err != nil {
		log.Fatal(err)
	}

	banned, err := method.ListBanned(client)
	if err != nil {
		log.Fatal(err)
	}

	mempool, err := method.GetMempoolInfo(client)
	if err != nil {
		log.Fatal(err)
	}

	addednodeinfo, err := method.GetAddedNodeInfo(client)
	if err != nil {
		log.Fatal(err)
	}

	nettotals, err := method.GetNetTotals(client)
	if err != nil {
		log.Fatal(err)
	}

	//uptime, err := method.Uptime(client)
	//if err != nil {
	//	log.Fatal(err)
	//}

	params := IndexTemplateParams{
		Title:         "Bitcoin Daemon Status",
		Peers:         peerinfo,
		Network:       networkinfo,
		Banned:        banned,
		Mempool:       mempool,
		AddedNodeInfo: addednodeinfo,
		NetTotals:     nettotals,
		Uptime:        0,
	}

	err = t.Execute(w, params)
	if err != nil {
		log.Println(err)
	}
}

type AddNodeHandler struct {
	// Configuration contains the app configuration. In this context only the server list is used.
	Configuration configuration.Configuration
}

func (i AddNodeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	client := rpc.NewRPCClient(i.Configuration.Url, i.Configuration.Username, i.Configuration.Password)
	method.AddNode(client, req.PostFormValue("node"))

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("true"))
}

type RemoveNodeHandler struct {
	// Configuration contains the app configuration. In this context only the server list is used.
	Configuration configuration.Configuration
}

func (i RemoveNodeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	client := rpc.NewRPCClient(i.Configuration.Url, i.Configuration.Username, i.Configuration.Password)
	method.RemoveNode(client, req.PostFormValue("node"))

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("true"))
}

func Init(mux *http.ServeMux, configuration configuration.Configuration) {
	indexHandler := IndexRequestHandler{Configuration: configuration}
	addNodeHandler := AddNodeHandler{Configuration: configuration}
	removeNodeHandler := RemoveNodeHandler{Configuration: configuration}

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	mux.Handle("/api/v1/addnode", addNodeHandler)
	mux.Handle("/api/v1/removenode", removeNodeHandler)
	mux.Handle("/", indexHandler)
}
