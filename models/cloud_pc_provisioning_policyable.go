package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcProvisioningPolicyable 
type CloudPcProvisioningPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]CloudPcProvisioningPolicyAssignmentable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDomainJoinConfiguration()(CloudPcDomainJoinConfigurationable)
    GetImageDisplayName()(*string)
    GetImageId()(*string)
    GetImageType()(*CloudPcProvisioningPolicyImageType)
    GetMicrosoftManagedDesktop()(MicrosoftManagedDesktopable)
    GetOnPremisesConnectionId()(*string)
    GetWindowsSettings()(CloudPcWindowsSettingsable)
    SetAssignments(value []CloudPcProvisioningPolicyAssignmentable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDomainJoinConfiguration(value CloudPcDomainJoinConfigurationable)()
    SetImageDisplayName(value *string)()
    SetImageId(value *string)()
    SetImageType(value *CloudPcProvisioningPolicyImageType)()
    SetMicrosoftManagedDesktop(value MicrosoftManagedDesktopable)()
    SetOnPremisesConnectionId(value *string)()
    SetWindowsSettings(value CloudPcWindowsSettingsable)()
}
