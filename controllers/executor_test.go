package controllers

import (
	distrunv1 "ghy-test/kubebuilder-demo/api/v1"
	"testing"
)

// 一个不存在的命令
const notExistsCmd = "1011928"

func TestExecutor(t *testing.T) {
	// 结果应该是 `nothing to do`
	t.Log(NewExecutor(nil).Exec())

	cmd := &distrunv1.Cmd{
		Spec: distrunv1.CmdSpec{
			Command: "ls",
		},
	}
	// 结果应该是 `[SUCCESS] 目录中的内容`
	t.Log(NewExecutor(cmd).Exec())

	cmd.Spec.Command = notExistsCmd
	// 结果应该是 `[ERROR] "exec: \"1011928\": executable file not found in $PATH"`
	t.Log(NewExecutor(cmd).Exec())

	cmd.Spec.Command = "sleep"
	cmd.Spec.Args = append(cmd.Spec.Args, "10")
	// 结果应该是 `[ERROR] "context deadline exceeded"`
	t.Log(NewExecutor(cmd).Exec())
}
