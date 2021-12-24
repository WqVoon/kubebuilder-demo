package controllers

import (
	"bytes"
	"context"
	distrunv1 "ghy-test/kubebuilder-demo/api/v1"
	"os/exec"
	"strings"
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
func (ce *CmdExecutor) Exec() (stdout string, stderr string, errorMsg string) {
	if ce == nil || strings.TrimSpace(ce.Command) == "" {
		return defaultResult, "", ""
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), ce.Timeout)
	defer cancelFunc()

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.CommandContext(ctx, ce.Command, ce.Args...)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	if err := cmd.Run(); err != nil {
		errorMsg = err.Error()
	}

	// 优先返回 ctx 的错误信息（超时错误）
	if ctx.Err() != nil {
		errorMsg = ctx.Err().Error()
	}

	stdout, stderr = stdoutBuf.String(), stderrBuf.String()
	return
}
