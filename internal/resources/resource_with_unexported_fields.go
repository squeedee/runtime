/*
Copyright 2022 the original author or authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resources

import (
	"context"
	"encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"reconciler.io/runtime/apis"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var (
	_ webhook.CustomDefaulter = &TestResourceUnexportedFields{}
	_ webhook.CustomValidator = &TestResourceUnexportedFields{}
	_ client.Object           = &TestResourceUnexportedFields{}
)

// +kubebuilder:object:root=true
// +genclient

type TestResourceUnexportedFields struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestResourceUnexportedFieldsSpec   `json:"spec"`
	Status TestResourceUnexportedFieldsStatus `json:"status"`
}

func (*TestResourceUnexportedFields) Default(ctx context.Context, obj runtime.Object) error {
	r, ok := obj.(*TestResourceUnexportedFields)
	if !ok {
		return fmt.Errorf("expected obj to be TestResourceUnexportedFields")
	}
	if r.Spec.Fields == nil {
		r.Spec.Fields = map[string]string{}
	}
	r.Spec.Fields["Defaulter"] = "ran"

	return nil
}

func (*TestResourceUnexportedFields) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	r, ok := obj.(*TestResourceUnexportedFields)
	if !ok {
		return nil, fmt.Errorf("expected obj to be TestResourceUnexportedFields")
	}

	return nil, r.validate(ctx).ToAggregate()
}

func (*TestResourceUnexportedFields) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	_, ok := oldObj.(*TestResourceUnexportedFields)
	if !ok {
		return nil, fmt.Errorf("expected oldObj to be TestResourceUnexportedFields")
	}
	r, ok := newObj.(*TestResourceUnexportedFields)
	if !ok {
		return nil, fmt.Errorf("expected newObj to be TestResourceUnexportedFields")
	}

	return nil, r.validate(ctx).ToAggregate()
}

func (*TestResourceUnexportedFields) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	_, ok := obj.(*TestResourceUnexportedFields)
	if !ok {
		return nil, fmt.Errorf("expected obj to be TestResourceUnexportedFields")
	}

	return nil, nil
}

func (r *TestResourceUnexportedFields) validate(ctx context.Context) field.ErrorList {
	errs := field.ErrorList{}

	if r.Spec.Fields != nil {
		if _, ok := r.Spec.Fields["invalid"]; ok {
			field.Invalid(field.NewPath("spec", "fields", "invalid"), r.Spec.Fields["invalid"], "")
		}
	}

	return errs
}

func (r *TestResourceUnexportedFields) ReflectUnexportedFieldsToStatus() {
	r.Status.unexportedFields = r.Spec.unexportedFields
}

// +kubebuilder:object:generate=true
type TestResourceUnexportedFieldsSpec struct {
	Fields           map[string]string `json:"fields,omitempty"`
	unexportedFields map[string]string
	Template         corev1.PodTemplateSpec `json:"template,omitempty"`

	ErrOnMarshal   bool `json:"errOnMarhsal,omitempty"`
	ErrOnUnmarshal bool `json:"errOnUnmarhsal,omitempty"`
}

func (r *TestResourceUnexportedFieldsSpec) GetUnexportedFields() map[string]string {
	return r.unexportedFields
}

func (r *TestResourceUnexportedFieldsSpec) SetUnexportedFields(f map[string]string) {
	r.unexportedFields = f
}

func (r *TestResourceUnexportedFieldsSpec) AddUnexportedField(key, value string) {
	if r.unexportedFields == nil {
		r.unexportedFields = map[string]string{}
	}
	r.unexportedFields[key] = value
}

func (r *TestResourceUnexportedFieldsSpec) MarshalJSON() ([]byte, error) {
	if r.ErrOnMarshal {
		return nil, fmt.Errorf("ErrOnMarshal true")
	}
	return json.Marshal(&struct {
		Fields         map[string]string      `json:"fields,omitempty"`
		Template       corev1.PodTemplateSpec `json:"template,omitempty"`
		ErrOnMarshal   bool                   `json:"errOnMarshal,omitempty"`
		ErrOnUnmarshal bool                   `json:"errOnUnmarshal,omitempty"`
	}{
		Fields:         r.Fields,
		Template:       r.Template,
		ErrOnMarshal:   r.ErrOnMarshal,
		ErrOnUnmarshal: r.ErrOnUnmarshal,
	})
}

func (r *TestResourceUnexportedFieldsSpec) UnmarshalJSON(data []byte) error {
	type alias struct {
		Fields         map[string]string      `json:"fields,omitempty"`
		Template       corev1.PodTemplateSpec `json:"template,omitempty"`
		ErrOnMarshal   bool                   `json:"errOnMarshal,omitempty"`
		ErrOnUnmarshal bool                   `json:"errOnUnmarshal,omitempty"`
	}
	a := &alias{}
	if err := json.Unmarshal(data, a); err != nil {
		return err
	}
	r.Fields = a.Fields
	r.Template = a.Template
	r.ErrOnMarshal = a.ErrOnMarshal
	r.ErrOnUnmarshal = a.ErrOnUnmarshal
	if r.ErrOnUnmarshal {
		return fmt.Errorf("ErrOnUnmarshal true")
	}
	return nil
}

// +kubebuilder:object:generate=true
type TestResourceUnexportedFieldsStatus struct {
	apis.Status      `json:",inline"`
	Fields           map[string]string `json:"fields,omitempty"`
	unexportedFields map[string]string
}

func (r *TestResourceUnexportedFieldsStatus) GetUnexportedFields() map[string]string {
	return r.unexportedFields
}

func (r *TestResourceUnexportedFieldsStatus) SetUnexportedFields(f map[string]string) {
	r.unexportedFields = f
}

func (r *TestResourceUnexportedFieldsStatus) AddUnexportedField(key, value string) {
	if r.unexportedFields == nil {
		r.unexportedFields = map[string]string{}
	}
	r.unexportedFields[key] = value
}

// +kubebuilder:object:root=true

type TestResourceUnexportedFieldsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []TestResourceUnexportedFields `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TestResourceUnexportedFields{}, &TestResourceUnexportedFieldsList{})

	if err := equality.Semantic.AddFuncs(
		func(a, b TestResourceUnexportedFieldsSpec) bool {
			return equality.Semantic.DeepEqual(a.Fields, b.Fields) &&
				equality.Semantic.DeepEqual(a.Template, b.Template) &&
				equality.Semantic.DeepEqual(a.ErrOnMarshal, b.ErrOnMarshal) &&
				equality.Semantic.DeepEqual(a.ErrOnUnmarshal, b.ErrOnUnmarshal)
		},
		func(a, b TestResourceUnexportedFieldsStatus) bool {
			return equality.Semantic.DeepEqual(a.Status, b.Status) &&
				equality.Semantic.DeepEqual(a.Fields, b.Fields)
		},
	); err != nil {
		panic(err)
	}
}
