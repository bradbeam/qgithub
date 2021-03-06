package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var searchPhrase = flag.String("search", "", "search string")
var oauthToken = flag.String("oauth", "", "oauth token")
var org = flag.String("org", "", "github user or org to search")
var file = flag.String("file", "/tmp/searchResults.md", "Output file to use with -markdown, defaults to /tmp/searchResults.md")
var markdown = flag.Bool("markdown", false, "render a markdown output file")
var language = flag.String("lang", "", "language to search for")
var perPage = flag.Int("pp", 30, "number of results to grab, default 30, max 100")

func main() {
	searchargs := []string{"+in:file"}

	flag.Parse()

	if *oauthToken == "" {
		log.Fatal("Need to specify an oAuth token with -oauth")
	}

	if *searchPhrase == "" {
		log.Fatal("Need to specify a search phrase with -search")
	}

	if *org != "" {
		searchargs = append(searchargs, "user:"+*org)
	}

	if *language != "" {
		searchargs = append(searchargs, "language:"+*language)
	}

	// Doing this weird thing to encode the query string to be safe
	// Might be overkill
	u, err := url.Parse("https://api.github.com/search/code?")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("per_page", strconv.Itoa(*perPage))
	q.Set("q", *searchPhrase)
	u.RawQuery = q.Encode()
	u, err = url.Parse(u.String() + strings.Join(searchargs, "+"))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(u.String())
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Add("Authorization", "token "+*oauthToken)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var results Results
	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		fmt.Println("error:", err)
	}

	if *markdown == true {
		output, err := os.Create(*file)
		defer output.Close()
		if err != nil {
			log.Fatal(err)
		}
		writer := bufio.NewWriter(output)
		for _, v := range results.Items {
			writer.WriteString(fmt.Sprintf("%s %s\n", "#", v.Name))
			writer.WriteString(fmt.Sprintf("[%s](%s)\n", v.Path, v.HtmlURL))
			for k, v2 := range v.TextMatches {
				writer.WriteString(fmt.Sprintf("### %s %d\n", "Match", k+1))
				writer.WriteString(fmt.Sprintln("```" + strings.Split(*language, ":")[1]))
				writer.WriteString(fmt.Sprintln(v2.Fragment))
				writer.WriteString(fmt.Sprintln("```"))
			}
		}
		writer.Flush()
	} else {
		for _, v := range results.Items {
			fmt.Println("=============================================")
			fmt.Printf("%s | %s\n", v.Name, v.Path)
			fmt.Println(v.HtmlURL)
			for _, v2 := range v.TextMatches {
				fmt.Println("--------------------------------------------")
				fmt.Println(v2.Fragment)
			}
		}
	}
}
