{{define "ipwhitelist"}}
    {{index .Site 0}}_IPWhiteList:
      whiteList:
        active: true
        sourceService: {{index .Site 0}}
        nextService: GatewayNodeBreaker
        ipList:
      {{- range $par := .Ipwhite }}
          - "{{$par}}"
      {{- end }}
        pathList:
      {{- range $par := .Urlwhite }}
          - "{{$par}}"
      {{- end -}}
{{- end -}}
