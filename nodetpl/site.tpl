{{- define "site" -}}
{{ $service := index .Site 0 -}}
http:
  routers:
    {{$service}}:
      entryPoints:
        - web
      service: {{$service}}
      rule: Host(`{{index .Site 1}}`)
      middlewares: 
      {{- range $middl := .Middlewares }}
        - {{$service}}_{{$middl}}
      {{- end }}
  middlewares:
{{- end -}}