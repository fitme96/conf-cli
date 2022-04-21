{{ define "passwordreq" }}
    {{index .Site 0}}_PasswordRequire:
      passwordRequire:
        active: true
        password: "{{.Passwd}}"
        pathList:
      {{- range $par := .Passwdurl }}
          - {{$par}}
      {{- end }}
{{- end -}}
