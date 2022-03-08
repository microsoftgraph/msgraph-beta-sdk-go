package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TenantAppManagementPolicyable 
type TenantAppManagementPolicyable interface {
    PolicyBaseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicationRestrictions()(AppManagementConfigurationable)
    GetIsEnabled()(*bool)
    GetServicePrincipalRestrictions()(AppManagementConfigurationable)
    SetApplicationRestrictions(value AppManagementConfigurationable)()
    SetIsEnabled(value *bool)()
    SetServicePrincipalRestrictions(value AppManagementConfigurationable)()
}
