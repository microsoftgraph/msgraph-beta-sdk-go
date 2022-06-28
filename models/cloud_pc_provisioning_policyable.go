package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcProvisioningPolicyable 
type CloudPcProvisioningPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlternateResourceUrl()(*string)
    GetAssignments()([]CloudPcProvisioningPolicyAssignmentable)
    GetCloudPcGroupDisplayName()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDomainJoinConfiguration()(CloudPcDomainJoinConfigurationable)
    GetGracePeriodInHours()(*int32)
    GetImageDisplayName()(*string)
    GetImageId()(*string)
    GetImageType()(*CloudPcProvisioningPolicyImageType)
    GetLocalAdminEnabled()(*bool)
    GetManagedBy()(*CloudPcManagementService)
    GetMicrosoftManagedDesktop()(MicrosoftManagedDesktopable)
    GetOnPremisesConnectionId()(*string)
    GetWindowsSettings()(CloudPcWindowsSettingsable)
    SetAlternateResourceUrl(value *string)()
    SetAssignments(value []CloudPcProvisioningPolicyAssignmentable)()
    SetCloudPcGroupDisplayName(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDomainJoinConfiguration(value CloudPcDomainJoinConfigurationable)()
    SetGracePeriodInHours(value *int32)()
    SetImageDisplayName(value *string)()
    SetImageId(value *string)()
    SetImageType(value *CloudPcProvisioningPolicyImageType)()
    SetLocalAdminEnabled(value *bool)()
    SetManagedBy(value *CloudPcManagementService)()
    SetMicrosoftManagedDesktop(value MicrosoftManagedDesktopable)()
    SetOnPremisesConnectionId(value *string)()
    SetWindowsSettings(value CloudPcWindowsSettingsable)()
}
