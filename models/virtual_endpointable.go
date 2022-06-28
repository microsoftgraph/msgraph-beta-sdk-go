package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEndpointable 
type VirtualEndpointable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuditEvents()([]CloudPcAuditEventable)
    GetCloudPCs()([]CloudPCable)
    GetDeviceImages()([]CloudPcDeviceImageable)
    GetGalleryImages()([]CloudPcGalleryImageable)
    GetOnPremisesConnections()([]CloudPcOnPremisesConnectionable)
    GetOrganizationSettings()(CloudPcOrganizationSettingsable)
    GetProvisioningPolicies()([]CloudPcProvisioningPolicyable)
    GetServicePlans()([]CloudPcServicePlanable)
    GetSnapshots()([]CloudPcSnapshotable)
    GetSupportedRegions()([]CloudPcSupportedRegionable)
    GetUserSettings()([]CloudPcUserSettingable)
    SetAuditEvents(value []CloudPcAuditEventable)()
    SetCloudPCs(value []CloudPCable)()
    SetDeviceImages(value []CloudPcDeviceImageable)()
    SetGalleryImages(value []CloudPcGalleryImageable)()
    SetOnPremisesConnections(value []CloudPcOnPremisesConnectionable)()
    SetOrganizationSettings(value CloudPcOrganizationSettingsable)()
    SetProvisioningPolicies(value []CloudPcProvisioningPolicyable)()
    SetServicePlans(value []CloudPcServicePlanable)()
    SetSnapshots(value []CloudPcSnapshotable)()
    SetSupportedRegions(value []CloudPcSupportedRegionable)()
    SetUserSettings(value []CloudPcUserSettingable)()
}
