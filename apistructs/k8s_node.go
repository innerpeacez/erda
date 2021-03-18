package apistructs

import "time"

type DrainNodeRequest struct {
	NodeName string `json:"nodeName"`

	// Continue even if there are pods not managed by a ReplicationController, ReplicaSet, Job, DaemonSet or StatefulSet
	Force bool `json:"force"`
	// Ignore DaemonSet-managed pods
	IgnoreAllDaemonSets bool `json:"ignoreAllDaemonSets"`
	// Continue even if there are pods using emptyDir (local data that will be deleted when the node is drained)
	DeleteLocalData bool `json:"deleteLocalData"`
	// The length of time to wait before giving up, zero means infinite
	Timeout time.Duration `json:"timeout"`
	// Period of time in seconds given to each pod to terminate gracefully. If negative, the default value specified in the pod will be use
	GracePeriodSeconds int `json:"gracePeriodSeconds"`
	// Label selector to filter pods on the node
	PodSelector string `json:"podSelector"`

	Selector string `json:"selector"`
	// DisableEviction forces drain to use delete rather than evict
	DisableEviction bool `json:"disableEviction"`
	// SkipWaitForDeleteTimeoutSeconds ignores pods that have a
	// DeletionTimeStamp > N seconds. It's up to the user to decide when this
	// option is appropriate; examples include the Node is unready and the pods
	// won't drain otherwise
	SkipWaitForDeleteTimeoutSeconds int
}
