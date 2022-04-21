{{ define "blacklist" }}
    {{index .Site 0}}_IPBlacklist:
      ipBlackList:
        active: true
        sourceRange:
      {{- range $par := .Ipblacklist }}
          - "{{$par}}"
      {{- end }}       
        blackProvince: 
      {{- range $par := .Provincelist }}
          - "{{$par}}"
      {{- end }}
{{- end -}}
