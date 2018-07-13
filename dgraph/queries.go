/*
 * Copyright 2017 Ankur Yadav (ankurayadav@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 		http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dgraph

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"fmt"
)

var names = []string{"5fYfxWzxRh",
	"oNWp4vzHDe",
	"ZfUP2HO12v",
	"MAaJqdLoA0",
	"TtCyN46HxA",
	"vBa0lalrAc",
	"ofehApW3V6",
	"SOFIDC7KRX",
	"IlKoBGKL3x",
	"CK47Zbx5H9",
	"z7QZnIkVTi",
	"HEIQYGLE6D",
	"MAUaVKfPRj",
	"jWEdzfdbUj",
	"Wv7s0rM4au",
	"PPhTMScy7L",
	"5GwTwRwccU",
	"PnBEv6IRFx",
	"vTlcWJysLl",
	"DxsedOOKbJ",
	"MWUsTDfRJD",
	"s6bLQJqzCD",
	"SLDs80nDpv",
	"NcpIyP7e63",
	"V39HPj3an1",
	"jaS9FGznFy",
	"KnvGzZqVOP",
	"bgUy8O36xf",
	"2cwFy6xwvT",
	"ArtRLme8aN",
	"uGKJWYjApM",
	"ok6ERKM2ey",
	"azISzOCHET",
	"8CIbpeKbfd",
	"DmKB09fabT",
	"iiwKh6XQGV",
	"3Z9pis4xlF",
	"agvkMtXy2Q",
	"Q1hiOJS07H",
	"WoVPHM38Bm",
	"Ia4yBTXOVQ",
	"VX82z1ZXAn",
	"KiZbxybVlN",
	"JvO0MROmat",
	"hrydMdnIoI",
	"Y80jJ3juRR",
	"gtMqNbipQI",
	"ozgJxaVXva",
	"x7z5Ix7s16",
	"JrbIQobRCO",
	"yetdLEcB1w",
	"D5IiTi5aQ1",
	"m9olZ9AfJt",
	"cw4xX6r8Cv",
	"JxjKq46KZc",
	"DvnaE7EJS9",
	"vx5KC2yRvG",
	"94IlaaufW3",
	"ckF4IEBS0U",
	"Bxk9cl1m2F",
	"FRh6F2smIM",
	"8VW65MyURR",
	"TPJ26DhhYQ",
	"o3gKBOAq0R",
	"Imz3lUPDSx",
	"u82HVmrYyX",
	"Q89TN6t2q2",
	"uEaGvVOnPp",
	"y4oY7XqQbV",
	"OrB84aVWiV",
	"FqJGjMsyz8",
	"4nexwka15a",
	"5giKztRuB1",
	"ebjNedmLPJ",
	"2U3bRNHyV5",
	"FGQ8ZrnKYo",
	"2qlxlq9gjz",
	"QQZhhkLMJC",
	"3bUt3qmUL8",
	"wodzQkfIEr",
	"2nSx50D47p",
	"8PIm20d7hS",
	"l11UMe6bg1",
	"GjyCfN2aR3",
	"4zXFra0C3x",
	"AABZJ0xuus",
	"zJZrPIYMEY",
	"fccZy4KMEC",
	"3NAEL7GURB",
	"9WkoHcdMEB",
	"Z7OqeaI4gw",
	"1NhL8nkLH3",
	"SLCXjPCfKU",
	"wq0SIB8Ukg",
	"XWAU0X90O3",
	"vIBlx22uHj",
	"4ZMsJ97IAp",
	"qvZZt4Uf7I",
	"6IqnRlPe2t",
	"XhKFzfHMeT",
}

var benchmarkQueries = []struct {
	query string
}{
	// Finding movies and genre of movies directed by "Steven Spielberg"?
	{
		query: `
				{
				  var(func: eq(name, "%s")) {
						A AS follow
				  }

				  var(func: uid(A)) {
					P AS post
				  }
				  timeline(func: uid(P),orderdesc: post.post_time, offset: 0, first: 100) {
					~post{
					name
				  }
					post.content
					post.post_time
					post.media_type
				  }
				}
		`,
	},
}

func runBench(n int, b *testing.B) {

	// Http client for getting JSON response.
	hc := &http.Client{Transport: &http.Transport{
		MaxIdleConnsPerHost: 100,
	}}

	b.StopTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str := fmt.Sprintf(benchmarkQueries[n].query, names[i % len(names)])
		r, err := http.NewRequest("POST", "http://127.0.0.1:8080/query", bytes.NewBufferString(str))

		b.StartTimer()
		resp, err := hc.Do(r)
		b.StopTimer()

		if err != nil {
			log.Fatal("Error in query")
		} else {

			_, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Couldn't parse response body. %+v", err)
			}

			// log.Printf("Response body: %s", body)

		}
	}
}
