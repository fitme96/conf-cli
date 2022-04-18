{{define "services"}}
  services:
    {{index .Site 0}}:
      loadBalancer:
        servers:
      {{- range $par := .Servers }}
          - url: "{{$par}}"
      {{- end }}
{{- end -}}