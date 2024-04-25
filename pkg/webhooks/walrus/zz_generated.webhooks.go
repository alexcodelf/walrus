//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

package walrus

import (
	v1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

func GetWebhookConfigurations(np string, c v1.WebhookClientConfig) (*v1.ValidatingWebhookConfiguration, *v1.MutatingWebhookConfiguration) {
	vwc := GetValidatingWebhookConfiguration(np+"-validation", c)
	mwc := GetMutatingWebhookConfiguration(np+"-mutation", c)
	return vwc, mwc
}

func GetValidatingWebhookConfiguration(n string, c v1.WebhookClientConfig) *v1.ValidatingWebhookConfiguration {
	return &v1.ValidatingWebhookConfiguration{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "admissionregistration.k8s.io/v1",
			Kind:       "ValidatingWebhookConfiguration",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: n,
		},
		Webhooks: []v1.ValidatingWebhook{
			vwh_walrus_pkg_webhooks_walrus_SettingWebhook(c),
			vwh_walrus_pkg_webhooks_walrus_VariableWebhook(c),
		},
	}
}

func GetMutatingWebhookConfiguration(n string, c v1.WebhookClientConfig) *v1.MutatingWebhookConfiguration {
	return nil
}

func (*SettingWebhook) ValidatePath() string {
	return "/validate-walrus-seal-io-v1-setting"
}

func vwh_walrus_pkg_webhooks_walrus_SettingWebhook(c v1.WebhookClientConfig) v1.ValidatingWebhook {
	path := "/validate-walrus-seal-io-v1-setting"

	cc := c.DeepCopy()
	if cc.Service != nil {
		cc.Service.Path = &path
	} else if c.URL != nil {
		cc.URL = ptr.To(*c.URL + path)
	}

	return v1.ValidatingWebhook{
		Name:         "validate.walrus.seal.io.v1.setting",
		ClientConfig: *cc,
		Rules: []v1.RuleWithOperations{
			{
				Rule: v1.Rule{
					APIGroups: []string{
						"walrus.seal.io",
					},
					APIVersions: []string{
						"v1",
					},
					Resources: []string{
						"settings",
					},
					Scope: ptr.To[v1.ScopeType]("Namespaced"),
				},
				Operations: []v1.OperationType{
					"CREATE",
					"UPDATE",
				},
			},
		},
		FailurePolicy:  ptr.To[v1.FailurePolicyType]("Fail"),
		MatchPolicy:    ptr.To[v1.MatchPolicyType]("Equivalent"),
		SideEffects:    ptr.To[v1.SideEffectClass]("None"),
		TimeoutSeconds: ptr.To[int32](10),
		AdmissionReviewVersions: []string{
			"v1",
		},
	}
}

func (*VariableWebhook) ValidatePath() string {
	return "/validate-walrus-seal-io-v1-variable"
}

func vwh_walrus_pkg_webhooks_walrus_VariableWebhook(c v1.WebhookClientConfig) v1.ValidatingWebhook {
	path := "/validate-walrus-seal-io-v1-variable"

	cc := c.DeepCopy()
	if cc.Service != nil {
		cc.Service.Path = &path
	} else if c.URL != nil {
		cc.URL = ptr.To(*c.URL + path)
	}

	return v1.ValidatingWebhook{
		Name:         "validate.walrus.seal.io.v1.variable",
		ClientConfig: *cc,
		Rules: []v1.RuleWithOperations{
			{
				Rule: v1.Rule{
					APIGroups: []string{
						"walrus.seal.io",
					},
					APIVersions: []string{
						"v1",
					},
					Resources: []string{
						"variables",
					},
					Scope: ptr.To[v1.ScopeType]("Namespaced"),
				},
				Operations: []v1.OperationType{
					"CREATE",
					"UPDATE",
				},
			},
		},
		FailurePolicy:  ptr.To[v1.FailurePolicyType]("Fail"),
		MatchPolicy:    ptr.To[v1.MatchPolicyType]("Equivalent"),
		SideEffects:    ptr.To[v1.SideEffectClass]("None"),
		TimeoutSeconds: ptr.To[int32](10),
		AdmissionReviewVersions: []string{
			"v1",
		},
	}
}