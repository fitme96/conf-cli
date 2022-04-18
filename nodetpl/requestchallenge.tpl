{{ define "requestchallenge" }}
    {{ index .Site 0 }}_RequestChallenge:
      requestChallenge:
        jschallengeLimit: {{index .Requestchall 0}}
        captchaChallengeLimit: {{index .Requestchall 1}}
        timeLimit: "{{index .Requestchall 2}}"
        shoudBlock: {{index .Requestchall 3}}
{{- end -}}