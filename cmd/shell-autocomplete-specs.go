package cmd

var resourceSpecs = make(map[string][]string)

var metadataJsonPaths = []string{
	"$.metadata",
	"$.metadata.name",
	"$.metadata.namespace",
	"$.metadata.labels",
	"$.metadata.annotations",
	"$.metadata.creationTimestamp",
	"$.metadata.uid",
	"$.status",
}

func initResourceSpecs() {
	resourceSpecs["pods"] = []string{
		"$.spec",
		"$.spec.containers",
		"$.spec.containers.name",
		"$.spec.containers.image",
		"$.spec.initContainers",
		"$.spec.nodeSelector",
		"$.spec.nodeName",
		"$.spec.restartPolicy",
		"$.spec.serviceAccountName",
		"$.spec.affinity",
		"$.spec.tolerations",
		"$.spec.hostAliases",
		"$.spec.securityContext",
		"$.spec.imagePullSecrets",
		"$.status.phase",
		"$.status.podIP",
		"$.status.conditions",
		"$.status.containerStatuses",
		"$.status.containerStatuses.name",
		"$.status.containerStatuses.ready",
		"$.status.containerStatuses.restartCount",
		"$.status.containerStatuses.image",
		"$.status.containerStatuses.state",
		"$.status.hostIP",
		"$.status.initContainerStatuses",
	}
	resourceSpecs["deployments"] = []string{
		"$.spec",
		"$.spec.replicas",
		"$.spec.selector",
		"$.spec.template",
		"$.spec.template.metadata",
		"$.spec.template.metadata.labels",
		"$.spec.template.spec",
		"$.spec.template.spec.containers",
		"$.spec.template.spec.initContainers",
		"$.spec.strategy",
		"$.spec.minReadySeconds",
		"$.spec.revisionHistoryLimit",
		"$.spec.paused",
		"$.spec.progressDeadlineSeconds",
		"$.status.observedGeneration",
		"$.status.replicas",
		"$.status.updatedReplicas",
		"$.status.readyReplicas",
		"$.status.availableReplicas",
		"$.status.unavailableReplicas",
		"$.status.conditions",
	}
	resourceSpecs["services"] = []string{
		"$.spec",
		"$.spec.clusterIP",
		"$.spec.replicas",
		"$.spec.selector",
		"$.spec.template",
		"$.spec.template.metadata",
		"$.spec.template.metadata.labels",
		"$.spec.template.spec",
		"$.spec.template.spec.containers",
		"$.spec.template.spec.initContainers",
		"$.spec.strategy",
		"$.spec.minReadySeconds",
		"$.spec.revisionHistoryLimit",
		"$.spec.paused",
		"$.spec.progressDeadlineSeconds",
		"$.status.observedGeneration",
		"$.status.replicas",
		"$.status.updatedReplicas",
		"$.status.readyReplicas",
		"$.status.availableReplicas",
		"$.status.unavailableReplicas",
		"$.status.conditions",
		"$.status.LoadBalancer",
	}
	resourceSpecs["ingresses"] = []string{
		"$.spec",
		"$.spec.rules",
		"$.spec.rules.host",
		"$.spec.rules.http",
		"$.spec.rules.http.paths",
		"$.spec.rules.http.paths.path",
		"$.spec.rules.http.paths.backend",
		"$.spec.rules.http.paths.backend.serviceName",
		"$.spec.rules.http.paths.backend.servicePort",
		"$.spec.tls",
		"$.spec.tls.hosts",
		"$.spec.tls.secretName",
		"$.status.loadBalancer",
		"$.status.loadBalancer.ingress",
		"$.status.loadBalancer.ingress.ip",
		"$.status.loadBalancer.ingress.hostname",
	}
	resourceSpecs["replicasets"] = []string{
		"$.spec",
		"$.spec.replicas",
		"$.spec.minReadySeconds",
		"$.spec.selector",
		"$.spec.selector.matchLabels",
		"$.spec.selector.matchExpressions",
		"$.spec.template",
		"$.spec.template.metadata",
		"$.spec.template.metadata.labels",
		"$.spec.template.spec",
		"$.spec.template.spec.containers",
		"$.spec.template.spec.initContainers",
		"$.spec.template.spec.volumes",
		"$.status.replicas",
		"$.status.fullyLabeledReplicas",
		"$.status.readyReplicas",
		"$.status.availableReplicas",
		"$.status.observedGeneration",
	}
	resourceSpecs["daemonsets"] = []string{
		"$.spec",
		"$.spec.selector",
		"$.spec.selector.matchLabels",
		"$.spec.selector.matchExpressions",
		"$.spec.template",
		"$.spec.template.metadata",
		"$.spec.template.metadata.labels",
		"$.spec.template.spec",
		"$.spec.template.spec.containers",
		"$.spec.template.spec.initContainers",
		"$.spec.template.spec.volumes",
		"$.spec.updateStrategy",
		"$.spec.minReadySeconds",
		"$.status.currentNumberScheduled",
		"$.status.numberMisscheduled",
		"$.status.desiredNumberScheduled",
		"$.status.numberReady",
		"$.status.observedGeneration",
		"$.status.updatedNumberScheduled",
		"$.status.numberAvailable",
		"$.status.numberUnavailable",
		"$.status.collisionCount",
	}
	resourceSpecs["statefulsets"] = []string{
		"$.spec",
		"$.spec.selector",
		"$.spec.selector.matchLabels",
		"$.spec.selector.matchExpressions",
		"$.spec.serviceName",
		"$.spec.template",
		"$.spec.template.metadata",
		"$.spec.template.metadata.labels",
		"$.spec.template.spec",
		"$.spec.template.spec.containers",
		"$.spec.template.spec.initContainers",
		"$.spec.template.spec.volumes",
		"$.spec.volumeClaimTemplates",
		"$.spec.updateStrategy",
		"$.spec.updateStrategy.type",
		"$.spec.updateStrategy.rollingUpdate",
		"$.spec.podManagementPolicy",
		"$.spec.replicas",
		"$.spec.revisionHistoryLimit",
		"$.status",
		"$.status.observedGeneration",
		"$.status.replicas",
		"$.status.readyReplicas",
		"$.status.currentReplicas",
		"$.status.updatedReplicas",
		"$.status.currentRevision",
		"$.status.updateRevision",
		"$.status.collisionCount",
		"$.status.conditions",
	}
	resourceSpecs["jobs"] = []string{
		"$.spec",
		"$.spec.parallelism",
		"$.spec.completions",
		"$.spec.activeDeadlineSeconds",
		"$.spec.backoffLimit",
		"$.spec.selector",
		"$.spec.selector.matchLabels",
		"$.spec.template",
		"$.spec.template.metadata",
		"$.spec.template.metadata.labels",
		"$.spec.template.spec",
		"$.spec.template.spec.containers",
		"$.spec.template.spec.initContainers",
		"$.spec.template.spec.restartPolicy",
		"$.status",
		"$.status.conditions",
		"$.status.startTime",
		"$.status.completionTime",
		"$.status.active",
		"$.status.succeeded",
		"$.status.failed",
	}
	resourceSpecs["cronjobs"] = []string{
		"$.spec",
		"$.spec.schedule",
		"$.spec.startingDeadlineSeconds",
		"$.spec.concurrencyPolicy",
		"$.spec.suspend",
		"$.spec.jobTemplate",
		"$.spec.jobTemplate.spec",
		"$.spec.jobTemplate.spec.parallelism",
		"$.spec.jobTemplate.spec.completions",
		"$.spec.jobTemplate.spec.activeDeadlineSeconds",
		"$.spec.jobTemplate.spec.template",
		"$.spec.jobTemplate.spec.template.metadata",
		"$.spec.jobTemplate.spec.template.spec",
		"$.spec.jobTemplate.spec.template.spec.containers",
		"$.spec.jobTemplate.spec.template.spec.initContainers",
		"$.spec.jobTemplate.spec.template.spec.restartPolicy",
		"$.status",
		"$.status.lastScheduleTime",
		"$.status.active",
		"$.status.lastSuccessfulTime",
	}
}
