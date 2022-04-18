{{- define "gateway" -}}
http:
  routers:
    whoami:
      entryPoints:
      - web
      service: {{index .Site 0}}_{{ .White }}
      rule: "PathPrefix(`/`)"
      middlewares:
        - defaultReject
        - defaultRealIp
  middlewares:
    defaultReject:
      reject:
        statusCode: "403"
        showMiddleware: false
        contactPerson: ""
        contactNumber: ""
    defaultRealIp:
      realIp:
        xforwarfordepth: 0
        headername: "X-Forwarded-For"
        useheader: true
  services:
    GatewayNodeBreaker:
      nodeBreaker:
        loadLimit:
          - 2000
        nodes: 
          - traefik_node
    traefik_node:
      loadBalancer:
        servers:
          - url: http://traefik_node
{{- end -}}