{{- define "aws-load-balancer-controller.vpcIdArg" -}}
{{- if .Values.vpcId }}
- --aws-vpc-id={{ .Values.vpcId }}
{{- else if and .Values.aws .Values.aws.vpcID }}
- --aws-vpc-id={{ .Values.aws.vpcID }}
{{- end }}
{{- end -}}

{{- define "aws-load-balancer-controller.regionArg" -}}
{{- if .Values.region }}
- --aws-region={{ .Values.region }}
{{- else if and .Values.aws .Values.aws.region }}
- --aws-region={{ .Values.aws.region }}
{{- end }}
{{- end -}}

{{- define "aws-load-balancer-controller.iamPodAnnotation" -}}
{{- if .Values.clusterID }}
iam.amazonaws.com/role: {{ printf "gs-%s-ALBController-Role" .Values.clusterID | quote }}
{{- end }}
{{- end -}}

{{/*
Set Giant Swarm serviceAccountAnnotations.
*/}}
{{- define "giantswarm.serviceAccountAnnotations" -}}
{{- if and (eq .Values.provider "aws") (eq .Values.aws.irsa "true") (not (hasKey .Values.serviceAccount.annotations "eks.amazonaws.com/role-arn")) }}
{{- $_ := set .Values.serviceAccount.annotations "eks.amazonaws.com/role-arn" (tpl "arn:aws:iam::{{ .Values.aws.accountID }}:role/gs-{{ .Values.clusterID }}-ALBController-Role" .) }}
{{- else if and (eq .Values.provider "capa") (not (hasKey .Values.serviceAccount.annotations "eks.amazonaws.com/role-arn")) }}
{{- $_ := set .Values.serviceAccount.annotations "eks.amazonaws.com/role-arn" (tpl "gs-{{ .Values.clusterID }}-ALBController-Role" .) }}
{{- end }}
{{- end -}}

{{- define "resource.vpa.enabled" -}}
{{- if and (.Capabilities.APIVersions.Has "autoscaling.k8s.io/v1") (.Values.verticalPodAutoscaler.enabled) }}true{{ else }}false{{ end }}
{{- end -}}

{{- define "resource.ingressClassParams.enabled" -}}
{{- if and (.Capabilities.APIVersions.Has "elbv2.k8s.aws/v1beta1") (.Values.ingressClassParams.create) }}true{{ else }}false{{ end }}
{{- end -}}