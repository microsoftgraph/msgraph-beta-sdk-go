package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// B2cIdentityUserFlowable 
type B2cIdentityUserFlowable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    IdentityUserFlowable
    GetApiConnectorConfiguration()(UserFlowApiConnectorConfigurationable)
    GetDefaultLanguageTag()(*string)
    GetIdentityProviders()([]IdentityProviderable)
    GetIsLanguageCustomizationEnabled()(*bool)
    GetLanguages()([]UserFlowLanguageConfigurationable)
    GetUserAttributeAssignments()([]IdentityUserFlowAttributeAssignmentable)
    GetUserFlowIdentityProviders()([]IdentityProviderBaseable)
    SetApiConnectorConfiguration(value UserFlowApiConnectorConfigurationable)()
    SetDefaultLanguageTag(value *string)()
    SetIdentityProviders(value []IdentityProviderable)()
    SetIsLanguageCustomizationEnabled(value *bool)()
    SetLanguages(value []UserFlowLanguageConfigurationable)()
    SetUserAttributeAssignments(value []IdentityUserFlowAttributeAssignmentable)()
    SetUserFlowIdentityProviders(value []IdentityProviderBaseable)()
}
