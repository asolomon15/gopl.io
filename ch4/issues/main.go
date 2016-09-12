// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

//!+
func main() {

	today := time.Now()
	monthAgo := today.AddDate(0, -1, 0)
	yearAgo := today.AddDate(-1, 0, 0)

	fmt.Println(monthAgo)
	fmt.Println(yearAgo)
	fmt.Println(os.Args[1:])
	os.Args = append(os.Args, "created:>2016-07-05") // Not sure I like this

	monthlyResult, err := github.SearchIssues("something string")
	yearlyResult, err := github.SearchIussues("something lese ")
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	//var monthReport []github.Issue
	// This should actually work

	report := make(map[string]*github.Issue)

	for _, item := range result.Items {
		//fmt.Println(item)
		if monthAgo.Before(item.CreatedAt) {
			report["lmonth"] = item
		} else if yearAgo.Before(item.CreatedAt) {
			//fmt.Println(item)
			report["lyear"] = item
		} else {
			//fmt.Println(item)
			report["myear"] = item
		}

		//fmt.Printf("#%-5d %10.9s %10s %.55s\n",
		//	item.Number, item.CreatedAt, item.User.Login, item.Title)
	}

	/*
		proper url
		https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+is%3Aopen+json+decoder+created:%3E2016-07-05+created:%3E2015-07-05&order=asc
		https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+is%3Aopen+json+decoder+created:%3E2016-07-05

		go run main.go repo:golang/go is:open json decoder
		https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+is%3Aopen+json+decoder
		#11046 2015-06-0      kurin encoding/json: Decoder internally buffers full input
		#15314 2016-04-1     okdave proposal: some way to reject unknown fields in encoding
		#5680  2013-06-1    eaigner encoding/json: set key converter on en/decoder
		#16212 2016-06-2  josharian encoding/json: do all reflect work before decoding
		#8658  2014-09-0  gopherbot encoding/json: use bufio
		#12001 2015-08-0  lukescott encoding/json: Marshaler/Unmarshaler not stream friendl
		#15146 2016-04-0 chanxuehong encoding/json: ",string" bool field accepts "terrible"
		#5901  2013-07-1        rsc encoding/json: allow override type marshaling
		#6716  2013-11-0  gopherbot encoding/json: include field name in unmarshal error me
		#7872  2014-04-2 extemporalgenome encoding/json: Encoder internally buffers full output
		#14750 2016-03-1 cyberphone encoding/json: parser ignores the case of member names
		#6384  2013-09-1    joeshaw encoding/json: encode precise floating point integers u
		#6647  2013-10-2    btracey x/tools/cmd/godoc: display type kind of each named type
		#8717  2014-09-1    dvyukov cmd/compile: random performance fluctuations after unre
		#4237  2012-10-1  gjemiller encoding/base64: URLEncoding padding is optional
		#14811 2016-03-1 hvnsweeting cmd/go: go test -cover failed to import, without -cover
	*/

}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
