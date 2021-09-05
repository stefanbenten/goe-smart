package template

//go:generate $GOPATH/bin/go-bindata -pkg template -ignore template.go -ignore bindata.go .

import "html/template"

var (
	Index = template.Must(template.New("index").Parse(string(MustAsset("index.html.tmpl"))))
)
