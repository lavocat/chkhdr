# ChkHdr

Simple application written in Go that fetches the headers for a given URL and indicates if security headers are present.  

### Build

In the root of the project `go build`

### Usage

`./chkhdr --url https://your-site.com`

### Output

```
Checking URL: https://ytmnd.com (200 OK)
	Access-Control-Allow-Origin: Header Missing!
	Content-Type: [text/html; charset=UTF-8]
	Content-Security-Policy: Header Missing!
	Cross-Origin-Embedder-Policy: Header Missing!
	Cross-Origin-Resource-Policy: Header Missing!
	Cross-Origin-Opener-Policy: Header Missing!
	Strict-Transport-Security: Header Missing!
	Referrer-Policy: Header Missing!
	X-Content-Type-Options: Header Missing!
	X-Frame-Options: Header Missing!
	X-XSS-Protection: Header Missing!
```

### Planned Improvements

- Print error if URL is malformed
- Add --header flag to specify a single or array of headers to check
- Add capability to check multiple URLs either with --urls flag or handling a list
- Sometimes it is not enough for headers to be present, they should also be a certain value
