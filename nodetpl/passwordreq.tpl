{{ define "passwordreq" }}
    {{index .Site 0}}_PasswordRequire:
      passwordRequire: "{{.Passwd}}"
        pathList:
      {{- range $par := .Passwdurl }}
          - {{$par}}
      {{- end }}
{{- end -}}