package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEndpoint 
type VirtualEndpoint struct {
    Entity
    // Cloud PC audit event.
    auditEvents []CloudPcAuditEventable
    // Cloud managed virtual desktops.
    cloudPCs []CloudPCable
    // The image resource on Cloud PC.
    deviceImages []CloudPcDeviceImageable
    // The external partner settings on a Cloud PC.
    externalPartnerSettings []CloudPcExternalPartnerSettingable
    // The gallery image resource on Cloud PC.
    galleryImages []CloudPcGalleryImageable
    // A defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
    onPremisesConnections []CloudPcOnPremisesConnectionable
    // The Cloud PC organization settings for a tenant.
    organizationSettings CloudPcOrganizationSettingsable
    // Cloud PC provisioning policy.
    provisioningPolicies []CloudPcProvisioningPolicyable
    // The reports property
    reports CloudPcReportsable
    // Cloud PC service plans.
    servicePlans []CloudPcServicePlanable
    // Cloud PC snapshots.
    snapshots []CloudPcSnapshotable
    // Cloud PC supported regions.
    supportedRegions []CloudPcSupportedRegionable
    // Cloud PC user settings.
    userSettings []CloudPcUserSettingable
}
// NewVirtualEndpoint instantiates a new virtualEndpoint and sets the default values.
func NewVirtualEndpoint()(*VirtualEndpoint) {
    m := &VirtualEndpoint{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.virtualEndpoint";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateVirtualEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpoint(), nil
}
// GetAuditEvents gets the auditEvents property value. Cloud PC audit event.
func (m *VirtualEndpoint) GetAuditEvents()([]CloudPcAuditEventable) {
    return m.auditEvents
}
// GetCloudPCs gets the cloudPCs property value. Cloud managed virtual desktops.
func (m *VirtualEndpoint) GetCloudPCs()([]CloudPCable) {
    return m.cloudPCs
}
// GetDeviceImages gets the deviceImages property value. The image resource on Cloud PC.
func (m *VirtualEndpoint) GetDeviceImages()([]CloudPcDeviceImageable) {
    return m.deviceImages
}
// GetExternalPartnerSettings gets the externalPartnerSettings property value. The external partner settings on a Cloud PC.
func (m *VirtualEndpoint) GetExternalPartnerSettings()([]CloudPcExternalPartnerSettingable) {
    return m.externalPartnerSettings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEndpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["auditEvents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcAuditEventFromDiscriminatorValue , m.SetAuditEvents)
    res["cloudPCs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPCFromDiscriminatorValue , m.SetCloudPCs)
    res["deviceImages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcDeviceImageFromDiscriminatorValue , m.SetDeviceImages)
    res["externalPartnerSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcExternalPartnerSettingFromDiscriminatorValue , m.SetExternalPartnerSettings)
    res["galleryImages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcGalleryImageFromDiscriminatorValue , m.SetGalleryImages)
    res["onPremisesConnections"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcOnPremisesConnectionFromDiscriminatorValue , m.SetOnPremisesConnections)
    res["organizationSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCloudPcOrganizationSettingsFromDiscriminatorValue , m.SetOrganizationSettings)
    res["provisioningPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcProvisioningPolicyFromDiscriminatorValue , m.SetProvisioningPolicies)
    res["reports"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCloudPcReportsFromDiscriminatorValue , m.SetReports)
    res["servicePlans"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcServicePlanFromDiscriminatorValue , m.SetServicePlans)
    res["snapshots"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcSnapshotFromDiscriminatorValue , m.SetSnapshots)
    res["supportedRegions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcSupportedRegionFromDiscriminatorValue , m.SetSupportedRegions)
    res["userSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcUserSettingFromDiscriminatorValue , m.SetUserSettings)
    return res
}
// GetGalleryImages gets the galleryImages property value. The gallery image resource on Cloud PC.
func (m *VirtualEndpoint) GetGalleryImages()([]CloudPcGalleryImageable) {
    return m.galleryImages
}
// GetOnPremisesConnections gets the onPremisesConnections property value. A defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *VirtualEndpoint) GetOnPremisesConnections()([]CloudPcOnPremisesConnectionable) {
    return m.onPremisesConnections
}
// GetOrganizationSettings gets the organizationSettings property value. The Cloud PC organization settings for a tenant.
func (m *VirtualEndpoint) GetOrganizationSettings()(CloudPcOrganizationSettingsable) {
    return m.organizationSettings
}
// GetProvisioningPolicies gets the provisioningPolicies property value. Cloud PC provisioning policy.
func (m *VirtualEndpoint) GetProvisioningPolicies()([]CloudPcProvisioningPolicyable) {
    return m.provisioningPolicies
}
// GetReports gets the reports property value. The reports property
func (m *VirtualEndpoint) GetReports()(CloudPcReportsable) {
    return m.reports
}
// GetServicePlans gets the servicePlans property value. Cloud PC service plans.
func (m *VirtualEndpoint) GetServicePlans()([]CloudPcServicePlanable) {
    return m.servicePlans
}
// GetSnapshots gets the snapshots property value. Cloud PC snapshots.
func (m *VirtualEndpoint) GetSnapshots()([]CloudPcSnapshotable) {
    return m.snapshots
}
// GetSupportedRegions gets the supportedRegions property value. Cloud PC supported regions.
func (m *VirtualEndpoint) GetSupportedRegions()([]CloudPcSupportedRegionable) {
    return m.supportedRegions
}
// GetUserSettings gets the userSettings property value. Cloud PC user settings.
func (m *VirtualEndpoint) GetUserSettings()([]CloudPcUserSettingable) {
    return m.userSettings
}
// Serialize serializes information the current object
func (m *VirtualEndpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuditEvents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAuditEvents())
        err = writer.WriteCollectionOfObjectValues("auditEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPCs() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCloudPCs())
        err = writer.WriteCollectionOfObjectValues("cloudPCs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceImages() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceImages())
        err = writer.WriteCollectionOfObjectValues("deviceImages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExternalPartnerSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExternalPartnerSettings())
        err = writer.WriteCollectionOfObjectValues("externalPartnerSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGalleryImages() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetGalleryImages())
        err = writer.WriteCollectionOfObjectValues("galleryImages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOnPremisesConnections() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOnPremisesConnections())
        err = writer.WriteCollectionOfObjectValues("onPremisesConnections", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("organizationSettings", m.GetOrganizationSettings())
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetProvisioningPolicies())
        err = writer.WriteCollectionOfObjectValues("provisioningPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reports", m.GetReports())
        if err != nil {
            return err
        }
    }
    if m.GetServicePlans() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetServicePlans())
        err = writer.WriteCollectionOfObjectValues("servicePlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSnapshots() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSnapshots())
        err = writer.WriteCollectionOfObjectValues("snapshots", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSupportedRegions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSupportedRegions())
        err = writer.WriteCollectionOfObjectValues("supportedRegions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserSettings())
        err = writer.WriteCollectionOfObjectValues("userSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuditEvents sets the auditEvents property value. Cloud PC audit event.
func (m *VirtualEndpoint) SetAuditEvents(value []CloudPcAuditEventable)() {
    m.auditEvents = value
}
// SetCloudPCs sets the cloudPCs property value. Cloud managed virtual desktops.
func (m *VirtualEndpoint) SetCloudPCs(value []CloudPCable)() {
    m.cloudPCs = value
}
// SetDeviceImages sets the deviceImages property value. The image resource on Cloud PC.
func (m *VirtualEndpoint) SetDeviceImages(value []CloudPcDeviceImageable)() {
    m.deviceImages = value
}
// SetExternalPartnerSettings sets the externalPartnerSettings property value. The external partner settings on a Cloud PC.
func (m *VirtualEndpoint) SetExternalPartnerSettings(value []CloudPcExternalPartnerSettingable)() {
    m.externalPartnerSettings = value
}
// SetGalleryImages sets the galleryImages property value. The gallery image resource on Cloud PC.
func (m *VirtualEndpoint) SetGalleryImages(value []CloudPcGalleryImageable)() {
    m.galleryImages = value
}
// SetOnPremisesConnections sets the onPremisesConnections property value. A defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *VirtualEndpoint) SetOnPremisesConnections(value []CloudPcOnPremisesConnectionable)() {
    m.onPremisesConnections = value
}
// SetOrganizationSettings sets the organizationSettings property value. The Cloud PC organization settings for a tenant.
func (m *VirtualEndpoint) SetOrganizationSettings(value CloudPcOrganizationSettingsable)() {
    m.organizationSettings = value
}
// SetProvisioningPolicies sets the provisioningPolicies property value. Cloud PC provisioning policy.
func (m *VirtualEndpoint) SetProvisioningPolicies(value []CloudPcProvisioningPolicyable)() {
    m.provisioningPolicies = value
}
// SetReports sets the reports property value. The reports property
func (m *VirtualEndpoint) SetReports(value CloudPcReportsable)() {
    m.reports = value
}
// SetServicePlans sets the servicePlans property value. Cloud PC service plans.
func (m *VirtualEndpoint) SetServicePlans(value []CloudPcServicePlanable)() {
    m.servicePlans = value
}
// SetSnapshots sets the snapshots property value. Cloud PC snapshots.
func (m *VirtualEndpoint) SetSnapshots(value []CloudPcSnapshotable)() {
    m.snapshots = value
}
// SetSupportedRegions sets the supportedRegions property value. Cloud PC supported regions.
func (m *VirtualEndpoint) SetSupportedRegions(value []CloudPcSupportedRegionable)() {
    m.supportedRegions = value
}
// SetUserSettings sets the userSettings property value. Cloud PC user settings.
func (m *VirtualEndpoint) SetUserSettings(value []CloudPcUserSettingable)() {
    m.userSettings = value
}
