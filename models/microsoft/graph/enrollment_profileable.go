package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EnrollmentProfileable 
type EnrollmentProfileable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConfigurationEndpointUrl()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEnableAuthenticationViaCompanyPortal()(*bool)
    GetRequireCompanyPortalOnSetupAssistantEnrolledDevices()(*bool)
    GetRequiresUserAuthentication()(*bool)
    SetConfigurationEndpointUrl(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEnableAuthenticationViaCompanyPortal(value *bool)()
    SetRequireCompanyPortalOnSetupAssistantEnrolledDevices(value *bool)()
    SetRequiresUserAuthentication(value *bool)()
}
