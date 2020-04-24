package uri

import (
	"errors"
	"segaline/src/util"
	"strconv"
	"strings"
)

type Uri struct {
	form util.UriForm

	scheme util.UriScheme
	user   string
	host   string
	port   uint16

	path  []string
	query map[string]string
}

func Parse(method util.HttpMethod, raw string) (uri Uri, err error) {
	if len(raw) > util.RequestMaxURILength {
		err = errors.New(util.ErrorRequestURILengthExceeded)
		return
	}

	if raw == "*" && method == util.HttpMethodOptions {
		return Uri{
			form:   util.UriFormAsterisk,
			scheme: util.UriSchemeHttp,
		}, nil
	}

	if method == util.HttpMethodConnect {
		uri.user, uri.host, uri.port, err = parseAuthority(raw)
		if uri.user != "" {
			return uri, errors.New("authority with user info in connect request")
		}
		uri.form = util.UriFormAuthority
	} else if strings.HasPrefix(raw, "http:") || strings.HasPrefix(raw, "https:") {
		uri, err = parseAbsoluteUri(raw)
		if err != nil {
			return
		}
		uri.form = util.UriFormAbsolute
	} else if strings.HasPrefix(raw, "/") {
		uri.path, uri.query, err = parseAbsolutePathWithQuery(raw)
		if err != nil {
			return
		}
		uri.form = util.UriFormOrigin
	}
	return
}

func (uri *Uri) PathString() string {
	return "/" + strings.Join(uri.path, "/")
}

func (uri *Uri) String() string {
	var user, port, path, query string

	if uri.user != "" {
		user = util.PercentEncode(uri.user + "@")
	}
	if uri.port > 0 {
		port = ":" + strconv.Itoa(int(uri.port))
	}

	path = "/" + util.PercentEncode(strings.Join(uri.path, "/"))
	if len(uri.query) > 0 {
		query = "?"
		for name, value := range uri.query {
			query += name + "=" + value + "&"
		}
		query = query[:len(query)-1]
	}

	return string(uri.scheme) + user + util.PercentEncode(uri.host) + port + path + util.PercentEncode(query)
}
