package utils

import "net/url"

// 拼接url中的path
func EncodeUrl(m map[string]string) string {
	p := url.Values{}

	for k := range m {
		p.Add(k, m[k])
	}

	return p.Encode()
}
