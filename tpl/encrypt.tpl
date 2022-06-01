{{- define "encrypt" }}
    {{index .Site 0}}_static:
      staticfiles: {}
    {{index .Site 0}}_encrypt:
      requestEncrypt:
        ShoudBlock: false
{{- end -}}