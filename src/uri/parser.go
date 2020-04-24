package uri

import (
	"errors"
	"math"
	"regexp"
	"segaline/src/util"
	"strconv"
	"strings"
)

var hostNameChars = regexp.MustCompile(`^[a-zA-Z\-._~%!$&'()*+,;=]*$`)
var ipAddress = regexp.MustCompile("^" + strings.Repeat(`\.(\d|\d\d|1\d\d|2[0-4]\d|25[0-6])`, 4)[1:] + "$")
var userInfoChars = regexp.MustCompile(`^[a-zA-Z\-._~%!$&'()*+,;=:]*$`)
var pathChars = regexp.MustCompile(`^[a-zA-Z\-._~%!$&'()*+,;=:@]*$`)
var queryChars = regexp.MustCompile(`^[a-zA-Z\-._~%!$&'()*+,;=:@/?]*$`)

func parseAbsoluteUri(raw string) (uri Uri, err error) {
	if !strings.HasPrefix(raw, util.UriSchemeHttp+":") && !strings.HasPrefix(raw, util.UriSchemeHttps+":") {
		err = errors.New("unsupported scheme")
		return
	}

	var scheme util.UriScheme
	switch raw[4] {
	case ':':
		scheme = util.UriSchemeHttp
	case 's':
		scheme = util.UriSchemeHttps
	}

	raw = raw[strings.Index(raw, ":")+1:]
	var user, host string
	var port uint16
	var path []string
	var query map[string]string

	if raw[0:2] == "//" && len(raw) > 2 {
		raw = raw[2:]
		pathStart := strings.Index(raw, "/")
		if pathStart > 0 {
			user, host, port, err = parseAuthority(raw[:pathStart])
			if err != nil {
				return
			}
			path, query, err = parseAbsolutePathWithQuery(raw[pathStart:])
		} else {
			user, host, port, err = parseAuthority(raw)
		}
	} else if raw[0] == '/' {
		path, query, err = parseAbsolutePathWithQuery(raw)
	} else {
		path, query, err = parseAbsolutePathWithQuery("/" + raw)
	}

	uri = Uri{util.UriFormAbsolute, scheme, user, host, port, path, query}
	return
}

func parseAuthority(raw string) (user string, host string, port uint16, err error) {
	userAndRest := strings.SplitN(raw, "@", 2)
	var rest string
	if len(userAndRest) < 2 {
		rest = userAndRest[0]
	} else {
		if !isUserInfo(userAndRest[0]) {
			err = errors.New("invalid user info")
			return
		}
		user = util.PercentDecode(userAndRest[0])
		rest = userAndRest[1]
	}

	hostAndPort := strings.Split(rest, ":")
	host = hostAndPort[0]
	if len(hostAndPort) > 1 {
		var maybePort int

		host = strings.Join(hostAndPort[0:len(hostAndPort)-1], "")
		maybePort, err = strconv.Atoi(hostAndPort[len(hostAndPort)-1])

		if err != nil || maybePort < 0 || maybePort > math.MaxUint16 {
			err = errors.New("invalid port")
			return
		}
		port = uint16(maybePort)
	}

	if !isHostName(host) && !isIPAddress(host) {
		err = errors.New("invalid host")
	}
	host = util.PercentDecode(host)
	return
}

func parseAbsolutePathWithQuery(raw string) (path []string, query map[string]string, err error) {
	query = map[string]string{}

	pathAndQuery := strings.SplitN(raw, "?", 2)
	stringPath, stringQuery := pathAndQuery[0], ""
	if len(pathAndQuery) > 1 {
		stringQuery = pathAndQuery[1]
	}

	if strings.Contains(stringPath, "//") {
		err = errors.New("empty intermediate path segment")
		return
	}

	pathParts := strings.Split(strings.TrimPrefix(strings.TrimSuffix(stringPath, "/"), "/"), "/")
	for _, part := range pathParts {
		if !isPath(part) || part == ".." || strings.EqualFold(part, "%2e%2e") {
			err = errors.New("invalid or unsupported path segment")
			return
		}
		path = append(path, util.PercentDecode(part))
	}

	if len(stringQuery) == 0 {
		return
	}

	queryParts := strings.Split(stringQuery, "&")
	for _, param := range queryParts {
		nameAndValue := strings.SplitN(param, "=", 2)
		if len(nameAndValue) < 2 || !isQuery(param) {
			err = errors.New("invalid query parameter")
			return
		}
		query[util.PercentDecode(nameAndValue[0])] = util.PercentDecode(nameAndValue[1])
	}
	return
}

func isQuery(str string) bool {
	return queryChars.MatchString(str)
}

func isPath(str string) bool {
	return pathChars.MatchString(str)
}

func isUserInfo(str string) bool {
	return userInfoChars.MatchString(str)
}

func isIPAddress(str string) bool {
	return ipAddress.MatchString(str)
}

func isHostName(str string) bool {
	return hostNameChars.MatchString(str)
}
