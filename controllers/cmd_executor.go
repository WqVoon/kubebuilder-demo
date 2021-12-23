package controllers

import (
	"context"
	"fmt"
	distrunv1 "ghy-test/kubebuilder-demo/api/v1"
	"os/exec"
	"time"
)

const (
	defaultTimeout = time.Second
	defaultResult  = "nothing to do"
)

type CmdExecutor struct {
	Command string
	Args    []string
	Timeout time.Duration
}

func NewExecutor(cmd *distrunv1.Cmd) *CmdExecutor {
	if cmd == nil {
		return nil
	}

	timeout := time.Duration(cmd.Spec.Timeout) * time.Second
	if timeout == 0 {
		timeout = defaultTimeout
	}

	return &CmdExecutor{
		Command: cmd.Spec.Command,
		Args:    cmd.Spec.Args,
		Timeout: timeout,
	}
}

// 执行命令，返回命令执行结果的字符串形式（错误信息或正常结果）
func (ce *CmdExecutor) Exec() string {
	if ce == nil {
		return defaultResult
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), ce.Timeout)
	defer cancelFunc()

	cmd := exec.CommandContext(ctx, ce.Command, ce.Args...)
	out, err := cmd.CombinedOutput()

	// 优先返回 ctx 的错误信息（超时错误）
	if ctx.Err() != nil {
		return fmt.Sprintf("[ERROR] %q", ctx.Err().Error())
	}
	if err != nil {
		return fmt.Sprintf("[ERROR] %q", err.Error())
	}

	return fmt.Sprintf("[SUCCESS] %q", string(out))
}
