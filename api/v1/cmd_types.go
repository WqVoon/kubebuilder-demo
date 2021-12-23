/*
   这是来自 hack/boilerplate.go.txt 的内容
   执行 `make generate` 时会被 controller-gen 读取
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CmdSpec defines the desired state of Cmd
type CmdSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Command 保存要被执行的 shell 命令，作为 `exec.CommandContext()` 的 name 参数
	Command string `json:"command,omitempty"`
	// Args 保存要 shell 命令的参数，作为 `exec.CommandContext()` 的 args 参数
	Args []string `json:"args,omitempty"`
	// Timeout 保存命令执行的最大等待时间（单位秒，默认为 1），超时则取消
	Timeout int64 `json:"timeout,omitempty"`
}

// CmdStatus defines the observed state of Cmd
type CmdStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Results 保存 worker 执行的结果，其中的 key 是主机名，value 是命令的输出（stdout & stderr）
	Results map[string]string `json:"results,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Cmd is the Schema for the cmds API
type Cmd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CmdSpec   `json:"spec,omitempty"`
	Status CmdStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CmdList contains a list of Cmd
type CmdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cmd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cmd{}, &CmdList{})
}

// 判断两个 Cmd 需要执行的命令是否相等
func (self *Cmd) CommandIsEqualTo(other *Cmd) bool {
	if self == nil && other == nil {
		return true
	}
	if self == nil || other == nil {
		return false
	}

	if self.Spec.Command != other.Spec.Command {
		return false
	}

	if len(self.Spec.Args) != len(other.Spec.Args) {
		return false
	}

	for idx, selfArg := range self.Spec.Args {
		otherArg := other.Spec.Args[idx]

		if selfArg != otherArg {
			return false
		}
	}

	return true
}
