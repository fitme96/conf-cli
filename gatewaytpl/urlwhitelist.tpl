{{define "urlwhitelist"}}
    {{index .Site 0}}_PathWhiteList:
      whiteList:
        sourceService: 
          pathList:
      {{- range $par := .Urlwhite.Urllist }}
            - "{{$par}}"
      {{- end -}}
{{- end -}}