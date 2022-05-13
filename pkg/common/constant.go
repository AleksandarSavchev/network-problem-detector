// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package common

const (
	// 	MDNSServiceNodeNetAgent is the mDNS service name of the agent running on the host network.
	MDNSServiceNodeNetAgent = "network-problem-detector-host-node._tcp"
	// NamespaceKubeSystem is the kube-system namespace
	NamespaceKubeSystem = "kube-system"
	// AgentConfigFilename is the name of the config file
	AgentConfigFilename = "agent-config.yaml"
	// EnvNodeName is the env variable to get the node name in an agent pod
	EnvNodeName = "NODE_NAME"
	// EnvNodeIP is the env variable to get the node ip in an agent pod
	EnvNodeIP = "NODE_IP"
	// EnvPodIP is the env variable to get the pod ip in an agent pod
	EnvPodIP = "POD_IP"
	// LabelKeyK8sApp is the label key used to mark the pods
	LabelKeyK8sApp = "k8s-app"
	// ApplicationName is the application name
	ApplicationName = "network-problem-detector"
	// NameAgentConfigMap name of the config map for the agents
	NameAgentConfigMap = ApplicationName + "-config"
	// NameDaemonSetAgentNodeNet name of the daemon set running in the node network
	NameDaemonSetAgentNodeNet = ApplicationName + "-host"
	// NameDaemonSetAgentPodNet name of the daemon set running in the pod network
	NameDaemonSetAgentPodNet = ApplicationName + "-pod"
	// NameDeploymentAgentController name of the deployment running the agent controller
	NameDeploymentAgentController = ApplicationName + "-controller"
	// PathOutputBaseDir parente directory path of output directory with observations in pods
	PathOutputBaseDir = "/var/lib/gardener"
	// PathOutputDir path of output directory with observations in pods
	PathOutputDir = PathOutputBaseDir + "/" + ApplicationName
	// PodNetPodGRPCPort is the port used for the GRPC server of the pods running in the pod network
	PodNetPodGRPCPort = 8880
	// PodNetPodMetricsPort is the port used for the metrics http server of the pods running in the pod network
	PodNetPodMetricsPort = 8881
	// NodeNetPodGRPCPort is the port used for the GRPC server of the pods running in the node network
	NodeNetPodGRPCPort = 1011
	// NodeNetPodMetricsPort is the port used for the metrics http server of the pods running in the node network
	NodeNetPodMetricsPort = 1012
)
