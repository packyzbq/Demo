package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	StatePending   string = "Pending"
	StateRunning   string = "Running"
	StateSucceeded string = "Succeeded"
	StateFailed    string = "Failed"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=example.com

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Demo struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata"`
	Spec          DemoSpec   `json:"spec"`
	Status        DemoStatus `json:"status"`
}

type DemoSpec struct {
	Foo string `json:"foo"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DemoStatus struct {
	State   string `json:"state"`
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DemoList struct {
	v1.TypeMeta `json:",inline"`
	v1.ListMeta `json:"metadata"`
	Items       []Demo `json:"items"`
}
