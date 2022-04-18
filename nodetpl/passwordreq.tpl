{{ define "passwordreq" }}
    {{index .Site 0}}_PasswordRequire:
      passwordRequire:
        password: "{{.Passwd}}"
        pathList:
      {{- range $par := .Passwdurl }}
          - {{$par}}
      {{- end }}
{{- end -}}
