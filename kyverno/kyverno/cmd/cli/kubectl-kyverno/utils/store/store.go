package store

import (
	"github.com/kyverno/kyverno/pkg/registryclient"
	rbacv1 "k8s.io/api/rbac/v1"
)

var (
	Mock           bool
	registryClient registryclient.Client
	AllowApiCalls  bool
	ContextVar     Context
	ForeachElement int
	Subjects       Subject
)

func SetMock(mock bool) {
	Mock = mock
}

func GetMock() bool {
	return Mock
}

func SetForEachElement(foreachElement int) {
	ForeachElement = foreachElement
}

func GetForeachElement() int {
	return ForeachElement
}

func SetRegistryAccess(access bool) {
}

func GetRegistryAccess() bool {
	return false
}

func GetRegistryClient() registryclient.Client {
	return registryClient
}

func SetContext(context Context) {
	ContextVar = context
}

func GetContext() Context {
	return ContextVar
}

func GetPolicyFromContext(policyName string) *Policy {
	for _, policy := range ContextVar.Policies {
		if policy.Name == policyName {
			return &policy
		}
	}
	return nil
}

func GetPolicyRuleFromContext(policyName string, ruleName string) *Rule {
	for _, policy := range ContextVar.Policies {
		if policy.Name == policyName {
			for _, rule := range policy.Rules {
				if rule.Name == ruleName {
					return &rule
				}
			}
		}
	}
	return nil
}

type Context struct {
	Policies []Policy `json:"policies"`
}

type Policy struct {
	Name  string `json:"name"`
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Name          string                   `json:"name"`
	Values        map[string]interface{}   `json:"values"`
	ForEachValues map[string][]interface{} `json:"foreachValues"`
}

func SetSubjects(subjects Subject) {
	Subjects = subjects
}

func GetSubjects() Subject {
	return Subjects
}

type Subject struct {
	Subject rbacv1.Subject `json:"subject,omitempty" yaml:"subject,omitempty"`
}

func AllowApiCall(allow bool) {
	AllowApiCalls = allow
}

func IsAllowApiCall() bool {
	return AllowApiCalls
}