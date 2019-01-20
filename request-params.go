// Copyright © 2019 Max Shirokov
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package gotelegrambot

import (
	"net/url"
)

type RequestParams map[string]string

func Defaults() RequestParams {
	return make(RequestParams)
}

func (params RequestParams) ToURLValues() url.Values {
	v := url.Values{}

	for key, value := range params {
		v.Set(key, value)
	}

	return v
}
