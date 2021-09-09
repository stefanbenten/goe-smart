package template

//go:generate $GOPATH/bin/go-bindata -pkg template -ignore template.go -ignore bindata.go .

import (
	"html/template"
	"time"
)

var (
	Index = template.Must(template.New("index").Funcs(
		template.FuncMap{
			"fromUnix": time.Unix,
			"Iterate": func(count int, step int) []int {
				var i int
				var Items []int
				if step < 0 {
					i = count
				} else {
					i = 0
				}
				for i <= count && i >= 0 {
					Items = append(Items, i)
					i = i + step
				}
				return Items
			},
		},
	).Parse(string(MustAsset("index.tmpl.html"))))
)
