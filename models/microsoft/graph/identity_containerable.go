package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityContainerable 
type IdentityContainerable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApiConnectors()([]IdentityApiConnectorable)
    GetB2cUserFlows()([]B2cIdentityUserFlowable)
    GetB2xUserFlows()([]B2xIdentityUserFlowable)
    GetConditionalAccess()(ConditionalAccessRootable)
    GetContinuousAccessEvaluationPolicy()(ContinuousAccessEvaluationPolicyable)
    GetIdentityProviders()([]IdentityProviderBaseable)
    GetUserFlowAttributes()([]IdentityUserFlowAttributeable)
    GetUserFlows()([]IdentityUserFlowable)
    SetApiConnectors(value []IdentityApiConnectorable)()
    SetB2cUserFlows(value []B2cIdentityUserFlowable)()
    SetB2xUserFlows(value []B2xIdentityUserFlowable)()
    SetConditionalAccess(value ConditionalAccessRootable)()
    SetContinuousAccessEvaluationPolicy(value ContinuousAccessEvaluationPolicyable)()
    SetIdentityProviders(value []IdentityProviderBaseable)()
    SetUserFlowAttributes(value []IdentityUserFlowAttributeable)()
    SetUserFlows(value []IdentityUserFlowable)()
}
