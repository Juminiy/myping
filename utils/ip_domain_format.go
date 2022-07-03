package utils

import (
	"regexp"
	"strings"
)

const (
	URL_PREFIX_HTTPS       = "https://"
	URL_PREFIX_HTTP        = "http://"
	URL_PREFIX_ALTERNATIVE = ""
	URL_VALID_DOMAIN_REGEX = `^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]\.[a-zA-Z]{2,}`
	URL_VALID_IP_REGEX     = `((?:(?:25[0-5]|2[0-4]\d|[01]?\d?\d)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d?\d))`
)

/*
  valid format :
	0. ip
	1. ip:port,
  	2. domain.com,
  	3. domain.com:port,
  	4. http://+(0,1,2,3)
  	5. https://+(0,1,2,3)
  invalid format:
  	else (1,2,3,4,5)
  finally format is :
	0. ip
	1. domain
*/
func URLFormatAssign(url string) (bool, string) {
	// first remove "https://" or "http://"
	if strings.HasPrefix(url, URL_PREFIX_HTTPS) {
		url = strings.ReplaceAll(url, URL_PREFIX_HTTPS, URL_PREFIX_ALTERNATIVE)
	}
	if strings.HasPrefix(url, URL_PREFIX_HTTP) {
		url = strings.ReplaceAll(url, URL_PREFIX_HTTP, URL_PREFIX_ALTERNATIVE)
	}
	// second remove suffix of ":port_number"
	if pindex := strings.Index(url, ":"); pindex != -1 {
		url = url[:pindex]
	}

	// Match either domain and ip regex
	valid_domain, _ := regexp.Compile(URL_VALID_DOMAIN_REGEX)
	valid_ip, _ := regexp.Compile(URL_VALID_IP_REGEX)

	return valid_domain.MatchString(url) || valid_ip.MatchString(url), url
}
