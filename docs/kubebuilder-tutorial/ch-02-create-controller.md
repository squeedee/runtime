# Chapter 2 - Create a controller with kubebuilder

* [Table of Contents](README.md)

## Create the project.

We will follow the steps shown in
the [kubebuilder book](https://book.kubebuilder.io/cronjob-tutorial/cronjob-tutorial#scaffolding-out-our-project),
however with some changes for clarity:

```shell
# Create a project working directory that matches what you're building
mkdir cronjob
cd cronjob

# create your project with kubebuilder
# you can leave the values provided, however when implementing your own controller:
# --domain should be a domain you own, if at all possible.
# --repo should point to the public, remote git repository you will use for this code.
# --project-name should reflect the name of your project, and must be a DNS 1123 label name
kubebuilder init \
  --domain tutorial.reconciler.io \
  --repo tutorial.reconciler.io/cronjob \
  --project-name=cronjob
```

Output:

```text
INFO Writing kustomize manifests for you to edit...
INFO Writing scaffold for you to edit...
INFO Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.19.0
go: downloading k8s.io/client-go v0.31.0
... snip ...
go: downloading golang.org/x/text v0.16.0
INFO Update dependencies:
$ go mod tidy
go: downloading github.com/onsi/ginkgo/v2 v2.19.0
... snip ...
go: downloading github.com/grpc-ecosystem/grpc-gateway v1.16.0
Next: define a resource with:
$ kubebuilder create api
```

**Note:
** [DNS 1123 Label Names defined](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-label-names)

## Initialize as a Git repository (optional)

Add the project to Git.

```shell
git init .
git add .
git ci -m "Initial Commit"
```

## Examining the structure (optional)

[//]: # (TODO)

**Note:** If you're accustomed to kubebuilder projects and want to skip the details, head over
to [Chapter 3 - ](ch-03-.md)

The resultant project structure is shown below, with notes for important sections

```text
.
├── Dockerfile (2)
├── Makefile (3)
├── PROJECT (1)
├── README.md
├── cmd
│   └── main.go (4)
├── config (2)
│   └── ... 
├── go.mod
├── go.sum
├── hack (5)
│   └── boilerplate.go.txt
└── test (6)
    ├── e2e
    │   ├── e2e_suite_test.go
    │   └── e2e_test.go
    └── utils
        └── utils.go
```

#### /PROJECT (1)

The `./PROJECT` file is kubebuilder's mechanism for keeping track of the projects layout and content.

If you use the command line options of `kubebuiler` correctly, you won't need to edit this file.

Eventually you may not want to continue to use kubebuilder to scaffold your project. If you reach this point, we
recommend that you delete the PROJECT file. You don't want to delete it yet, as the next chapter depends on it.

#### /config and /Dockerfile (2)

The `./config` directory holds generated [kustomize] templates for the deployment of your controller,
and `./Dockerfile` contains a generated docker image configuration to build your controller into a deployable image.

More information in
the [kubebuilder book](https://book.kubebuilder.io/cronjob-tutorial/basic-project#launch-configuration)

#### /Makefile (3)

The `./Makefile` has a lot of useful targets for building and running your code, along with generating CRDs for your
`./config`.

For more information, run:

```shell
make help
```

#### /cmd/main.go (4)

`./cmd/main.go` is the entry point for your controller. It contains quite a few kubebuilder scaffold markers (
`// +kubebuilder:scaffold:*`) that you should leave in place, unless you cease to use kubebuilder's scaffolding
features.

You can learn more about this initial `main.go` in
the [kubebuilder book](https://book.kubebuilder.io/cronjob-tutorial/empty-main).
In chapter 3, we will make significant changes here to initialize Reconciler.io Runtime controllers.

#### /hack/* (5)

`./hack` has become the idiomatic place for additional data and scripts. In kubebuilder projects, it minimally contains
`./hack/boilerplate.go.txt`, the "top matter" comment that prefixes your generated code. You should modify this in your
active projects to contain your own copyright notice and then run `make generate` to update generated code.

#### /test/* (6)

Kubebuilder also generates a test scaffold for e2e testing in `./test`. Reconciler.io Runtime [eschews the controller-runtime approach](https://github.com/reconcilerio/runtime?tab=readme-ov-file#testing) in favor of unit tests. We believe you'll find that
1. e2e testing with [envtest](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/envtest) is less desireable than using a few worthwhile smoke-tests against a fully configured cluster.
2. [Reconciler.io Runtime's Testing](https://github.com/reconcilerio/runtime?tab=readme-ov-file#testing) provides significantly improved confidence in the behavior of your controllers, whilst being fast, comprehensive and comprehensible.

[//]: # (TODO)

## Next: [Chapter 3 - ](ch-03-.md)

[kustomize]: https://kustomize.io/