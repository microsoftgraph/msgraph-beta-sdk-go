package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// VirtualEndpointable 
type VirtualEndpointable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
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
