# A test template

ssl {{.SSLPath}}

{{range $server := .Servers -}}
server {{$server.Name}} {{$server.Url}}
{{end}}