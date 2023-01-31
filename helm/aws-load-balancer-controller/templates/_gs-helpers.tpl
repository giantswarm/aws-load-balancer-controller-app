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
