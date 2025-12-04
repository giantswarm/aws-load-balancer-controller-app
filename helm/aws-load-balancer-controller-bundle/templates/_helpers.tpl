{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "aws-load-balancer-controller-bundle.name" -}}
{{- default .Chart.Name .Values.bundleNameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "aws-load-balancer-controller-bundle.fullname" -}}
{{- if .Values.fullBundleNameOverride -}}
{{- .Values.fullBundleNameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.bundleNameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "aws-load-balancer-controller-bundle.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "aws-load-balancer-controller-bundle.labels" -}}
app.kubernetes.io/name: {{ include "aws-load-balancer-controller-bundle.name" . }}
helm.sh/chart: {{ include "aws-load-balancer-controller-bundle.chart" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
giantswarm.io/service-type: "managed"
application.giantswarm.io/team: {{ index .Chart.Annotations "application.giantswarm.io/team" | quote }}
giantswarm.io/cluster: {{ .Values.clusterID | quote }}
{{- end -}}

{{/*
Get trust policy statements for all provided OIDC domains
*/}}
{{- define "aws-load-balancer-controller-bundle.trustPolicyStatements" -}}
{{- $configmap := (lookup "v1" "ConfigMap" .Release.Namespace (printf "%s-crossplane-config" .Values.clusterID)) -}}
{{- $cmvalues := dict -}}
{{- if and $configmap $configmap.data $configmap.data.values -}}
  {{- $cmvalues = fromYaml $configmap.data.values -}}
{{- end -}}
{{- range $index, $oidcDomain := $cmvalues.oidcDomains -}}
{{- if not (eq $index 0) }}, {{ end }}{
  "Effect": "Allow",
  "Principal": {
    "Federated": "arn:{{ $cmvalues.awsPartition }}:iam::{{ $cmvalues.accountID }}:oidc-provider/{{ $oidcDomain }}"
  },
  "Action": "sts:AssumeRoleWithWebIdentity",
  "Condition": {
    "StringLike": {
      "{{ $oidcDomain }}:sub": "system:serviceaccount:aws-load-balancer-controller:aws-load-balancer-controller"
    }
  }
}
{{- end -}}
{{- end -}}

{{/*
Set Giant Swarm specific values.
*/}}
{{- define "giantswarm.setValues" -}}
{{- $_ := set .Values.serviceAccount.annotations "eks.amazonaws.com/role-arn" (tpl "{{ .Values.clusterID }}-aws-load-balancer-controller-role" .) }}

{{- if and (not .Values.clusterName) .Values.clusterID }}
{{- $_ := set .Values "clusterName" (tpl "{{ .Values.clusterID }}" .) }}
{{- end -}}

{{/*    We always need to pass the following tags so that resources created by this controller are removed by CAPA when removing a CAPA cluster */}}
{{/*    - "kubernetes.io/cluster/$clusterID=owned"*/}}
{{/*    - "kubernetes.io/service-name=aws-alb-controller"*/}}
{{- $_ := set .Values.defaultTags (printf "kubernetes.io/cluster/%s" .Values.clusterID) "owned" }}
{{- $_ := set .Values.defaultTags "kubernetes.io/service-name" "aws-alb-controller" }}
{{- end -}}
