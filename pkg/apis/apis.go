// Copyright (c) 2020 Red Hat, Inc.
package apis

import (
	policiesv1 "github.com/open-cluster-management/governance-policy-propagator/pkg/apis/policies/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	// add policy scheme
	schemeBuilder := runtime.NewSchemeBuilder(addKnownTypes)
	if err := schemeBuilder.AddToScheme(s); err != nil {
		return err
	}
	return AddToSchemes.AddToScheme(s)
}

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(policiesv1.SchemeGroupVersion,
		&policiesv1.Policy{},
		&policiesv1.PolicyList{},
	)
	metav1.AddToGroupVersion(scheme, policiesv1.SchemeGroupVersion)

	scheme.AddKnownTypes(corev1.SchemeGroupVersion,
		&corev1.EventList{},
	)
	metav1.AddToGroupVersion(scheme, corev1.SchemeGroupVersion)
	return nil
}
