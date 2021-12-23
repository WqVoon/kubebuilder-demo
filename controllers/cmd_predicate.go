package controllers

import (
	distrunv1 "ghy-test/kubebuilder-demo/api/v1"

	"sigs.k8s.io/controller-runtime/pkg/event"
)

type CmdPredicate struct{}

// 接受 Create 和 Generic 事件
func (cp *CmdPredicate) Create(event.CreateEvent) bool   { return true }
func (cp *CmdPredicate) Generic(event.GenericEvent) bool { return true }

// 如果新旧 Cmd 对象待执行的命令相等，那么忽略这次更新
// 避免因为 Cmd.Status 的变化导致重复执行 Reconcile 逻辑
func (cp *CmdPredicate) Update(e event.UpdateEvent) bool {
	oldObj, ok := e.ObjectOld.(*distrunv1.Cmd)
	if !ok {
		return true
	}

	newObj, ok := e.ObjectNew.(*distrunv1.Cmd)
	if !ok {
		return true
	}

	if oldObj.CommandIsEqualTo(newObj) {
		return false
	}

	return true
}

// 忽略 Delete 事件
func (cp *CmdPredicate) Delete(event.DeleteEvent) bool {
	return false
}
