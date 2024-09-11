package model

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"
)

// NetworkInfo is the network of vnode, will be set into node addresses
type NetworkInfo struct {
	NodeIP   string
	HostName string
}

// NodeMetadata is the base data of a vnode, will be transfer to default labels of a vnode
type NodeMetadata struct {
	// Name is the name of vnode
	Name string
	// Version is the version of vnode
	Version string
	// Status is the curr status of vnode
	Status NodeStatus
}

// NodeInfo is the data of node info.
type NodeInfo struct {
	Metadata     NodeMetadata
	NetworkInfo  NetworkInfo
	CustomTaints []v1.Taint
}

// NodeResource is the data of node resource
type NodeResource struct {
	Capacity    resource.Quantity
	Allocatable resource.Quantity
}

// NodeStatusData is the status of a node, you can set some custom attributes in this data structure
type NodeStatusData struct {
	Resources         map[v1.ResourceName]NodeResource
	CustomLabels      map[string]string
	CustomAnnotations map[string]string
	CustomConditions  []v1.NodeCondition
}

// ContainerStatusData is the status data of a container
type ContainerStatusData struct {
	// Key generated by tunnel, need to be the same as Tunnel GetContainerUniqueKey of same container
	Key string
	// Name container name
	Name string
	// PodKey is the key of pod which contains this container ,you can set it to PodKeyAll to present a shared container
	PodKey     string
	State      ContainerState
	ChangeTime time.Time
	Reason     string
	Message    string
}

type BuildVNodeProviderConfig struct {
	// NodeIP is the ip of the node
	NodeIP string

	// NodeHostname is the hostname of the node
	NodeHostname string

	// Name is the node name, will be sent to utils.FormatNodeName to construct vnode name, and Name will be set to label
	Name string

	// Name is the node version, will be set to label
	Version string

	// ENV is the env of node, will be set to label
	Env string

	// CustomTaints is the taints set by tunnel
	CustomTaints []v1.Taint
}

type BuildVNodeConfig struct {
	// Client is the runtime client instance
	Client client.Client

	// KubeCache is the cache of kube resources
	KubeCache cache.Cache

	// NodeID is the unique id of node, should be unique key of system
	NodeID string

	// NodeIP is the ip of node
	NodeIP string

	// NodeHostname is the hostname of node
	NodeHostname string

	// NodeName is the name of node, will set to node's label
	NodeName string

	// NodeVersion is the version of node, will set to node's label
	NodeVersion string

	// Env is the runtime env of virtual-kubelet, will set to node created by virtual kubelet
	Env string

	// CustomTaints is the taint set by tunnel
	CustomTaints []v1.Taint
}

type BuildVNodeControllerConfig struct {
	ClientID string

	Env string

	// VPodIdentity is the vpod special value of model.LabelKeyOfComponent
	VPodIdentity string
}