{{- define "ovra-zero.labels" -}}
app.kubernetes.io/name: ovra-zero
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version | replace "+" "_" }}
{{- end }}

{{- define "ovra-zero.selectorLabels" -}}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "ovra-zero.etcdInitialCluster" -}}
{{- $root := . -}}
{{- $items := list -}}
{{- range $i := until (int .Values.etcd.replicas) -}}
{{- $host := printf "%s-%d=http://%s-%d.%s.%s.svc.cluster.local:%d" $root.Values.etcd.name $i $root.Values.etcd.name $i $root.Values.etcd.headlessServiceName $root.Release.Namespace (int $root.Values.etcd.peerPort) -}}
{{- $items = append $items $host -}}
{{- end -}}
{{ join "," $items }}
{{- end }}

{{- define "ovra-zero.etcdClientHostsYaml" -}}
{{- $root := . -}}
{{- range $i := until (int .Values.etcd.replicas) }}
      - {{ printf "%s-%d.%s.%s.svc.cluster.local:%d" $root.Values.etcd.name $i $root.Values.etcd.headlessServiceName $root.Release.Namespace (int $root.Values.etcd.port) }}
{{- end }}
{{- end }}

{{- define "ovra-zero.redisClusterHosts" -}}
{{- $root := . -}}
{{- $items := list -}}
{{- range $i := until (int .Values.redis.cluster.replicas) -}}
{{- $items = append $items (printf "%s-%d.%s.%s.svc.cluster.local:%d" $root.Values.redis.name $i $root.Values.redis.headlessServiceName $root.Release.Namespace (int $root.Values.redis.port)) -}}
{{- end -}}
{{ join "," $items }}
{{- end }}
