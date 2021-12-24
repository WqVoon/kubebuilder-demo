/*
   这是来自 hack/boilerplate.go.txt 的内容
   执行 `make generate` 时会被 controller-gen 读取
*/

package controllers

import (
	"context"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	distrunv1 "ghy-test/kubebuilder-demo/api/v1"
)

const defaultNodeName = "test-node"

var nodeName string

func init() {
	nodeName = os.Getenv("KUBE_NODE_NAME")

	if nodeName == "" {
		nodeName = defaultNodeName
	}
}

// CmdReconciler reconciles a Cmd object
type CmdReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=distrun.demo.com,resources=cmds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=distrun.demo.com,resources=cmds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=distrun.demo.com,resources=cmds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Cmd object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *CmdReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	cmd := &distrunv1.Cmd{}
	if err := r.Get(ctx, req.NamespacedName, cmd); err != nil {
		logger.Error(err, "can not find resource")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	stdout, stderr, err := NewExecutor(cmd).Exec()
	logger.Info("result info",
		"nodeName", nodeName,
		"stdout", stdout,
		"stderr", stderr,
		"err", err,
	)
	cmd.Status.Results = map[string]distrunv1.CmdResult{nodeName: {
		Error:  err,
		Stdout: stdout,
		Stderr: stderr,
	}}

	if err := r.Status().Patch(ctx, cmd, client.Merge); err != nil {
		logger.Error(err, "can not patch resource")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CmdReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&distrunv1.Cmd{}).
		WithEventFilter(&CmdPredicate{}).
		Complete(r)
}
