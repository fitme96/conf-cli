{{ define "blacklist" }}
    {{index .Site 0}}_IPBlacklist:
      regionList:
        active: true
        WhiteProvince:
      {{- range $par := .Provincelist }}
          - "{{$par}}"
      {{- end }}
{{- end -}}