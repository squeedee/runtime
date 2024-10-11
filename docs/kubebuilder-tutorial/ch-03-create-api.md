# Chapter 3 - Create an API with kubebuilder

* [Table of Contents](README.md)

## Create the CronJob API.

When working with controller's in Go and with [Controller Runtime], we often refer to the public structures and code
that serialize your Custom Resources as "the API".

Let's create an API and see how [Controller Runtime] creates Custom Resources from the [marker comments]
and [struct tags] found in the source code. We'll create the `CronJob` API.

```shell
kubebuilder create api --version v1alpha1 --kind CronJob --controller --resource
```

`--controller` and `--resource` are specified to skip the prompts to create the controller stub, and the API stubs.

You will see what these create, and what `--version` and `--kind` influence later in this chapter.

## Update your repository (optional)

```shell
git add .
git ci -m "Add CronTab API"
```

## Looking at the new API

`./api/v1alpha1/groupversion_info.go` is created to reflect the first use of the `v1alpha1` resource version in your
API.
Subsequent calls to `kubebuilder create api` won't change this file.

`./api/v1alpha1/cronjob_types.go` is created to represent the new kind specified by `--kind CronJob`. This is a working
stub of your new CronJob Resource, with an empty status, and a `Spec` with a single field `Foo`.

## And the new Controller.

As well as the API specification, `kubebuilder` added a controller `./internal/controller/cronjob_controller.go` and
registered it in `./cmd/main.go`.

This newly created controller's job is to respond to changes to the CronJob resource's spec and metadata, and update the
status of the resource.
Right now, it does nothing, but it won't crash, so we can try it out in the next chapter.

## And the rest...

Before running your controller for the first time, let's take a look at one other thing that just happened.

When you ran `kubebuilder create api` it ran `make generate`.

`make generate` created the file `./api/v1alpha1/zz_generated.deepcopy.go`. Whenever you change the API, you should run
`make generate`, which ensures that any struct marked with `// +kubebuilder:object` complies with [Controller-Runtime's
`object` interface](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client@v0.19.0#Object).

You can learn more about the generator makers in
the [kubebuilder documentation](https://book.kubebuilder.io/reference/controller-gen#generators)

## Further reading

For more details about this step and the scaffolding that `kubebuilder` creates, you can read
the [kubebuilder new-api page](https://book.kubebuilder.io/cronjob-tutorial/new-api).

[//]: # (TODO)

## Next: [Chapter 3 - ](ch-03-.md)

[struct tags]: https://go.dev/ref/spec#Struct_types

[marker comments]: https://pkg.go.dev/sigs.k8s.io/controller-tools/pkg/markers