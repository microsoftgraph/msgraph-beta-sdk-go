package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CustomCalloutExtensionable 
type CustomCalloutExtensionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAuthenticationConfiguration()(CustomExtensionAuthenticationConfigurationable)
    GetClientConfiguration()(CustomExtensionClientConfigurationable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEndpointConfiguration()(CustomExtensionEndpointConfigurationable)
    SetAuthenticationConfiguration(value CustomExtensionAuthenticationConfigurationable)()
    SetClientConfiguration(value CustomExtensionClientConfigurationable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEndpointConfiguration(value CustomExtensionEndpointConfigurationable)()
}
