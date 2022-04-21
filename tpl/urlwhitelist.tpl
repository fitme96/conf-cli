{{define "urlwhitelist"}}
    {{index .Site 0}}_PathWhiteList:
      whiteList:
        active: true
        sourceService: 
          pathList:
      {{- range $par := .Urlwhite.Urllist }}
            - "{{$par}}"
      {{- end -}}
{{- end -}}
