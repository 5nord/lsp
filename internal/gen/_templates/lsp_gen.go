package lsp

{{range .TypeAliases}}
	{{ $base := type .Type}}
	{{ $typ := title .Name }}
	{{with .Documentation}}{{comment .}}{{end}}
	type {{$typ}} {{$base}}
{{end}}

{{range .Enumerations}}
	{{ $base := type .Type }}
	{{ $typ := title .Name }}
	{{with .Documentation}}{{comment .}}{{end}}
	type {{$typ}} {{$base}}
	const (
		{{range .Values -}}{{with .Documentation}}

		{{comment .}}{{end}}
		{{ title .Name}} {{ $typ }} = {{if eq $base "string"}}{{printf "%q" .Value}}{{else}}{{.Value}}{{end}}
		{{- end}}
	)
{{end}}

