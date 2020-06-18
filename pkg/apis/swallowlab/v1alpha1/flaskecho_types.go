package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FlaskEchoSpec defines the desired state of FlaskEcho
type FlaskEchoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
    // Size is the size of the FlaskEcho deployment
    Size int32 `json:"size"`
}

// FlaskEchoStatus defines the observed state of FlaskEcho
type FlaskEchoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
    // Nodes are the names of the FlaskEcho pods
    Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FlaskEcho is the Schema for the flaskechos API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=flaskechos,scope=Namespaced
type FlaskEcho struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FlaskEchoSpec   `json:"spec,omitempty"`
	Status FlaskEchoStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FlaskEchoList contains a list of FlaskEcho
type FlaskEchoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlaskEcho `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FlaskEcho{}, &FlaskEchoList{})
}
