package http

import (
	"errors"
	"segaline/src/util"
	"strconv"
	"strings"
)

type Uri struct {
	form Form

	scheme Scheme
	user   string
	host   string
	port   uint16

	path  []string
	query map[string]string
}

func ParseUri(method Method, raw string) (uri Uri, err error) {
	if len(raw) > util.RequestMaxURILength {
		err = errors.New(util.ErrorRequestURILengthExceeded)
		return
	}

	if raw == "*" && method == MethodOptions {
		return Uri{
			form:   FormAsterisk,
			scheme: SchemeHttp,
		}, nil
	}

	if method == MethodConnect {
		uri.user, uri.host, uri.port, err = parseAuthority(raw)
		if uri.user != "" {
			return uri, errors.New("authority with user info in connect request")
		}
		uri.form = FormAuthority
	} else if strings.HasPrefix(raw, "http:") || strings.HasPrefix(raw, "https:") {
		uri, err = parseAbsoluteUri(raw)
		if err != nil {
			return
		}
		uri.form = FormAbsolute
	} else if strings.HasPrefix(raw, "/") {
		uri.path, uri.query, err = parseAbsolutePathWithQuery(raw)
		if err != nil {
			return
		}
		uri.form = FormOrigin
	}
	return
}

func (uri *Uri) PathString() string {
	return "/" + strings.Join(uri.path, "/")
}

func (uri *Uri) String() string {
	var user, port, path, query string

	if uri.user != "" {
		user = encodePercent(uri.user + "@")
	}
	if uri.port > 0 {
		port = ":" + strconv.Itoa(int(uri.port))
	}

	path = "/" + encodePercent(strings.Join(uri.path, "/"))
	if len(uri.query) > 0 {
		query = "?"
		for name, value := range uri.query {
			query += name + "=" + value + "&"
		}
		query = query[:len(query)-1]
	}

	return string(uri.scheme) + user + encodePercent(uri.host) + port + path + encodePercent(query)
}
