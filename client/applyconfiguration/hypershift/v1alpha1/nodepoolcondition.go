/*


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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NodePoolConditionApplyConfiguration represents an declarative configuration of the NodePoolCondition type for use
// with apply.
type NodePoolConditionApplyConfiguration struct {
	Type               *string             `json:"type,omitempty"`
	Status             *v1.ConditionStatus `json:"status,omitempty"`
	Severity           *string             `json:"severity,omitempty"`
	LastTransitionTime *metav1.Time        `json:"lastTransitionTime,omitempty"`
	Reason             *string             `json:"reason,omitempty"`
	Message            *string             `json:"message,omitempty"`
	ObservedGeneration *int64              `json:"observedGeneration,omitempty"`
}

// NodePoolConditionApplyConfiguration constructs an declarative configuration of the NodePoolCondition type for use with
// apply.
func NodePoolCondition() *NodePoolConditionApplyConfiguration {
	return &NodePoolConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *NodePoolConditionApplyConfiguration) WithType(value string) *NodePoolConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *NodePoolConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *NodePoolConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithSeverity sets the Severity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Severity field is set to the value of the last call.
func (b *NodePoolConditionApplyConfiguration) WithSeverity(value string) *NodePoolConditionApplyConfiguration {
	b.Severity = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *NodePoolConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *NodePoolConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *NodePoolConditionApplyConfiguration) WithReason(value string) *NodePoolConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *NodePoolConditionApplyConfiguration) WithMessage(value string) *NodePoolConditionApplyConfiguration {
	b.Message = &value
	return b
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *NodePoolConditionApplyConfiguration) WithObservedGeneration(value int64) *NodePoolConditionApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}