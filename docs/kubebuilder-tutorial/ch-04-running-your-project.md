# Chapter 4 - Running your project

* [Table of Contents](README.md)

## Generate your CRDs

Along with running `make generate` when you change the API, you should also update your CRD definitions. Let's do that
now:

```shell
make manifests
```

This creates/updates:

* `./config/crd/bases/tutorial.reconciler.io_cronjobs.yaml`
* `./config/rbac/role.yaml`

The `role.yaml` changes are derived from the `// +kubebuilder:rbac` marker comments in
`./internal/controller/cronjob_controller.go`, which define the roles the controller must have to operate on your CRDs,
and any other resources you interact with. We will see more of this later, when we interact with a child resource.

`tutorial.reconciler.io_cronjobs.yaml` is the
complete [CustomResourceDefinition](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/)
derived from your API code.

The [marker documentation](https://book.kubebuilder.io/reference/controller-gen#generators) goes into detail about how
`controller-gen` (the tool behind `make generate` and `make manifests`) produces the generated files. Don't let them
overwhelm you, we'll visit the important ones in this tutorial, and you wont need to worry about the rest for some time.

## Add your CRD to git

```shell
git add .
git ci -m "generated first crd"
```

## Create a `kind` cluster

As we mentioned in [chapter 1](ch-01-preparation.md), you should use `kind` for local testing. If you haven't already,
make sure docker is running on your machine, and then create a cluster:

```shell
kind create cluster
```

You now have kubernetes running on your local machine! Test it with:

```shell
$ kubectl get ns
NAME                 STATUS   AGE
default              Active   34s
kube-node-lease      Active   34s
kube-public          Active   34s
kube-system          Active   34s
local-path-storage   Active   31s
```

## Install your CRD on the cluster

Install your CRDs with:

```shell
make install
```

And see the CRD is installed:

```shell
$ kubectl get crd
NAME                              CREATED AT
cronjobs.tutorial.reconciler.io   2024-10-11T00:39:59Z
```

`make install` did not build or run (or deploy) your controller. At the moment, we've only told the Kubernetes cluster what your CronJob Resource looks like.

## Run your controller

For development, we recommend running locally. You can even debug your code locally.

In a separate terminal window:
```shell
make run
```

You will see the controller begin logging:

```text
2024-10-11T08:41:03-04:00 INFO setup starting manager
2024-10-11T08:41:03-04:00 INFO starting server {"name": "health probe", "addr": "[::]:8081"}
2024-10-11T08:41:03-04:00 INFO Starting EventSource {"controller": "cronjob", "controllerGroup": "tutorial.reconciler.io", "controllerKind": "CronJob", "source": "kind source: *v1alpha1.CronJob"}
2024-10-11T08:41:03-04:00 INFO Starting Controller {"controller": "cronjob", "controllerGroup": "tutorial.reconciler.io", "controllerKind": "CronJob"}
2024-10-11T08:41:03-04:00 INFO Starting workers {"controller": "cronjob", "controllerGroup": "tutorial.reconciler.io", "controllerKind": "CronJob", "worker count": 1}
```

**Note**: Running and debugging webhooks locally is quite difficult, so most folks do not bother. This tutorial does not
go into webhooks, so if you do add them, remember to use `export ENABLE_WEBHOOKS=false` before you `make run`.

You can learn more about running your controller in the [kubebuilder documentation](https://book.kubebuilder.io/cronjob-tutorial/running)
