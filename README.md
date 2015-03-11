# About

This is a simple little program for querying github from the command line. It is targeted to pull back code snipplets to provide examples of how to implement a specific library or concept.

# Usage

`./qgithub -search PHRASE_TO_SEARCH -oauth GITHUB_OAUTH_TOKEN -org GITHUB_USER_OR_ORG [ -lang PROGRAMMING_LANGUAGE ] [ -markdown ]`

An example to search for `fmt` in `go`

`./qgithub -search fmt -oauth 123456 -org golang -lang go`

By default, the output will be written to standard out. If -markdown is specified, the output will be written to `/tmp/searchResults.md`

# License
None, do whatever you like
