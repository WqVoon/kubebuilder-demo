# kubebuilder-demo

> 每个节点上一个 agent，组成 DaemonSet，用一个 CR 来下发 shell 命令，agent 在本地执行后将结果写回到 CR 的 status 中

## CR 定义
```yaml
apiVersion: distrun.demo.com/v1
kind: Cmd
metadata:
  name: cmd-sample
spec:
# 命令主体，作为 `exec.CommandContext()` 的 name 参数
  command: echo
# 命令参数，作为 `exec.CommandContext()` 的 args 参数
  args:
  - Hello
  - World
# 最大执行时间，如果超过该时间命令还未执行完成，返回 `[ERROR] "context deadline exceeded"`
  timeout: 1 # 单位秒，默认为 1
status:
# 是一个 map，键是节点名（利用 fieldRef 获得），值是命令执行的结果
  results:
    minikube-m02: '[SUCCESS] "Hello World\n"'
    minikube-m03: '[SUCCESS] "Hello World\n"'
```


## 部署方式
### 本地部署
```shell
# 部署 CRD
make install

# 本地运行 operator
make run

# 部署 CR，后面可通过修改 `./config/samples` 中的文件并重新执行该命令来更新 CR
kubectl apply -f ./config/samples
```

### minikube 部署
```shell
# `minikube node add` 来加入新节点

# 部署 CRD
make install

# 会在本地生成 distrun:v1 镜像，IMG 的值可修改
make docker-build-new IMG=distrun:v1

# 将本地镜像载入到 minikube 中
minikube image load distrun:v1

# 部署相关组件到 minikube 中
make deploy IMG=distrun:v1

# 部署 CR，后面可通过修改 `./config/samples` 中的文件并重新执行该命令来更新 CR
kubectl apply -f ./config/samples
```


## 说明
- 仅在 `spec.command` 或 `spec.args` 变化时，agent 才会响应

- 应使用 `make docker-build-new` 命令来基于 busybox 构建镜像，kubebuilder 提供的原 Dockerfile 无法执行 shell 命令

- 因为场景原因，去掉了 manager 的 `--leader-elect` 参数