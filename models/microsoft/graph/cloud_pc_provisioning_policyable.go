package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcProvisioningPolicyable 
type CloudPcProvisioningPolicyable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignments()([]CloudPcProvisioningPolicyAssignmentable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDomainJoinConfiguration()(CloudPcDomainJoinConfigurationable)
    GetImageDisplayName()(*string)
    GetImageId()(*string)
    GetImageType()(*CloudPcProvisioningPolicyImageType)
    GetMicrosoftManagedDesktop()(MicrosoftManagedDesktopable)
    GetOnPremisesConnectionId()(*string)
    SetAssignments(value []CloudPcProvisioningPolicyAssignmentable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDomainJoinConfiguration(value CloudPcDomainJoinConfigurationable)()
    SetImageDisplayName(value *string)()
    SetImageId(value *string)()
    SetImageType(value *CloudPcProvisioningPolicyImageType)()
    SetMicrosoftManagedDesktop(value MicrosoftManagedDesktopable)()
    SetOnPremisesConnectionId(value *string)()
}
