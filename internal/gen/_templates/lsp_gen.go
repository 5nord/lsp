package lsp

{{range .Enumerations}}
	{{ $base := print .Type.Name | type }}
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

