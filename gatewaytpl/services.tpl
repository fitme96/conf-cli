{{define "services"}}
    {{index .Site 0}}:
      loadBalancer:
        servers:Servers
      {{- range $par := .Servers }}
          - url: "{{$par}}"
      {{- end }}
{{- end -}}