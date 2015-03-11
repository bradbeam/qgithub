# About

This is a simple little program for querying github from the command line. It is targeted to pull back code snipplets to provide examples of how to implement a specific library or concept.

# Usage

`./qgithub -search PHRASE_TO_SEARCH -oauth GITHUB_OAUTH_TOKEN -org GITHUB_USER_OR_ORG [ -lang PROGRAMMING_LANGUAGE ] [ -markdown [ -file FILENAME ] ] [ -pp MAX_NUMBER_OF_RESULTS ]`

An example to search for `fmt` in `go`

`./qgithub -search fmt -oauth 123456 -org golang -lang go`

By default, the output will be written to standard out. If -markdown is specified, the output will be written to `/tmp/searchResults.md` or a filename specified by `-file`

Raw usage:
```
  -file="/tmp/searchResults.md": Output file to use with -markdown, defaults to /tmp/searchResults.md
  -h=false: Print all flags
  -help=false: Print all flags
  -lang="": language to search for
  -markdown=false: render a markdown output file
  -oauth="": oauth token
  -org="": github user or org to search
  -pp=30: number of results to grab, default 30, max 100
  -search="": search string
```

# License
None, do whatever you like
