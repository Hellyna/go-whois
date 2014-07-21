package whois

import (
	`fmt`
	`net`
	`io`
	`regexp`
)

var iana_regexp = regexp.MustCompile(
	`whois:\s+([a-zA-Z0-9.-]+)`)
//var iana_regexp = regexp.MustCompile(
//	`domain:\s+([A-Z]+)\s*`)

func whois(domain, server string) (string, error) {
	conn, err := net.Dial(`tcp`, server+`:43`)
	defer conn.Close()

	if err != nil {
		return ``, err
	}

	conn.Write([]byte(domain + "\r\n"))

	buf := make([]byte, 1024)
	res := []byte{}

	for {
		numbytes, err := conn.Read(buf)
		res = append(res, buf[0:numbytes]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return ``, err
		}
	}

	return string(res), nil
}

func get_tld(domain string) string {
	strlen := len(domain)
	var tld string
	for i := strlen - 1; i >= 0; i-- {
		if domain[i] == '.' {
			tld = domain[i+1:strlen]
			break
		}
	}
	return tld
}

func Whois(domain string) (string, error) {
	tld := get_tld(domain)

	res, err := whois(tld, `whois.iana.org`)

	if err != nil {
		return ``, err
	}

	matches := iana_regexp.FindStringSubmatch(res)

	if matches == nil || len(matches) < 2 {
		return ``, fmt.Errorf(
			`iana_regexp cannot find any matches of domain or whois`)
	}

	res, err = whois(domain, matches[1])

	if err != nil {
		return ``, err
	}

	return res, nil
}

// vim:ts=4 sw=4 noet:
