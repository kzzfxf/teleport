// Copyright 2023 kzzfxf
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

type Config struct {
	Routes  []Route `json:"routes"`
	Groups  []Group `json:"groups"`
	Proxies []Proxy `json:"proxies"`
}

type Route struct {
	Server   string `json:"server"`
	IP       string `json:"ip"`
	Selector string `json:"selector"`
}

type Group struct {
	Name    string   `json:"name"`
	Servers []string `json:"servers"`
}

type Proxy struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	URL    string   `json:"url"`
	Labels []string `json:"labels"`
}
