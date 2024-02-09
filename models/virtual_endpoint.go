package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VirtualEndpoint struct {
    Entity
}
// NewVirtualEndpoint instantiates a new VirtualEndpoint and sets the default values.
func NewVirtualEndpoint()(*VirtualEndpoint) {
    m := &VirtualEndpoint{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpoint(), nil
}
// GetAuditEvents gets the auditEvents property value. Cloud PC audit event.
// returns a []CloudPcAuditEventable when successful
func (m *VirtualEndpoint) GetAuditEvents()([]CloudPcAuditEventable) {
    val, err := m.GetBackingStore().Get("auditEvents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcAuditEventable)
    }
    return nil
}
// GetBulkActions gets the bulkActions property value. The bulkActions property
// returns a []CloudPcBulkActionable when successful
func (m *VirtualEndpoint) GetBulkActions()([]CloudPcBulkActionable) {
    val, err := m.GetBackingStore().Get("bulkActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcBulkActionable)
    }
    return nil
}
// GetCloudPCs gets the cloudPCs property value. Cloud managed virtual desktops.
// returns a []CloudPCable when successful
func (m *VirtualEndpoint) GetCloudPCs()([]CloudPCable) {
    val, err := m.GetBackingStore().Get("cloudPCs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPCable)
    }
    return nil
}
// GetCrossCloudGovernmentOrganizationMapping gets the crossCloudGovernmentOrganizationMapping property value. Cloud PC organization mapping between public and US Government Community Cloud (GCC) organizations.
// returns a CloudPcCrossCloudGovernmentOrganizationMappingable when successful
func (m *VirtualEndpoint) GetCrossCloudGovernmentOrganizationMapping()(CloudPcCrossCloudGovernmentOrganizationMappingable) {
    val, err := m.GetBackingStore().Get("crossCloudGovernmentOrganizationMapping")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcCrossCloudGovernmentOrganizationMappingable)
    }
    return nil
}
// GetDeviceImages gets the deviceImages property value. The image resource on Cloud PC.
// returns a []CloudPcDeviceImageable when successful
func (m *VirtualEndpoint) GetDeviceImages()([]CloudPcDeviceImageable) {
    val, err := m.GetBackingStore().Get("deviceImages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcDeviceImageable)
    }
    return nil
}
// GetExternalPartnerSettings gets the externalPartnerSettings property value. The external partner settings on a Cloud PC.
// returns a []CloudPcExternalPartnerSettingable when successful
func (m *VirtualEndpoint) GetExternalPartnerSettings()([]CloudPcExternalPartnerSettingable) {
    val, err := m.GetBackingStore().Get("externalPartnerSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcExternalPartnerSettingable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualEndpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["auditEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcAuditEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcAuditEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcAuditEventable)
                }
            }
            m.SetAuditEvents(res)
        }
        return nil
    }
    res["bulkActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcBulkActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcBulkActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcBulkActionable)
                }
            }
            m.SetBulkActions(res)
        }
        return nil
    }
    res["cloudPCs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPCFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPCable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPCable)
                }
            }
            m.SetCloudPCs(res)
        }
        return nil
    }
    res["crossCloudGovernmentOrganizationMapping"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcCrossCloudGovernmentOrganizationMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossCloudGovernmentOrganizationMapping(val.(CloudPcCrossCloudGovernmentOrganizationMappingable))
        }
        return nil
    }
    res["deviceImages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcDeviceImageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcDeviceImageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcDeviceImageable)
                }
            }
            m.SetDeviceImages(res)
        }
        return nil
    }
    res["externalPartnerSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcExternalPartnerSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcExternalPartnerSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcExternalPartnerSettingable)
                }
            }
            m.SetExternalPartnerSettings(res)
        }
        return nil
    }
    res["frontLineServicePlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcFrontLineServicePlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcFrontLineServicePlanable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcFrontLineServicePlanable)
                }
            }
            m.SetFrontLineServicePlans(res)
        }
        return nil
    }
    res["galleryImages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcGalleryImageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcGalleryImageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcGalleryImageable)
                }
            }
            m.SetGalleryImages(res)
        }
        return nil
    }
    res["onPremisesConnections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcOnPremisesConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcOnPremisesConnectionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcOnPremisesConnectionable)
                }
            }
            m.SetOnPremisesConnections(res)
        }
        return nil
    }
    res["organizationSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcOrganizationSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationSettings(val.(CloudPcOrganizationSettingsable))
        }
        return nil
    }
    res["provisioningPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcProvisioningPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcProvisioningPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcProvisioningPolicyable)
                }
            }
            m.SetProvisioningPolicies(res)
        }
        return nil
    }
    res["reports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcReportsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReports(val.(CloudPcReportsable))
        }
        return nil
    }
    res["servicePlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcServicePlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcServicePlanable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcServicePlanable)
                }
            }
            m.SetServicePlans(res)
        }
        return nil
    }
    res["sharedUseServicePlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcSharedUseServicePlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcSharedUseServicePlanable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcSharedUseServicePlanable)
                }
            }
            m.SetSharedUseServicePlans(res)
        }
        return nil
    }
    res["snapshots"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcSnapshotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcSnapshotable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcSnapshotable)
                }
            }
            m.SetSnapshots(res)
        }
        return nil
    }
    res["supportedRegions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcSupportedRegionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcSupportedRegionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcSupportedRegionable)
                }
            }
            m.SetSupportedRegions(res)
        }
        return nil
    }
    res["userSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcUserSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcUserSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcUserSettingable)
                }
            }
            m.SetUserSettings(res)
        }
        return nil
    }
    return res
}
// GetFrontLineServicePlans gets the frontLineServicePlans property value. Front-line service plans for a Cloud PC.
// returns a []CloudPcFrontLineServicePlanable when successful
func (m *VirtualEndpoint) GetFrontLineServicePlans()([]CloudPcFrontLineServicePlanable) {
    val, err := m.GetBackingStore().Get("frontLineServicePlans")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcFrontLineServicePlanable)
    }
    return nil
}
// GetGalleryImages gets the galleryImages property value. The gallery image resource on Cloud PC.
// returns a []CloudPcGalleryImageable when successful
func (m *VirtualEndpoint) GetGalleryImages()([]CloudPcGalleryImageable) {
    val, err := m.GetBackingStore().Get("galleryImages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcGalleryImageable)
    }
    return nil
}
// GetOnPremisesConnections gets the onPremisesConnections property value. A defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
// returns a []CloudPcOnPremisesConnectionable when successful
func (m *VirtualEndpoint) GetOnPremisesConnections()([]CloudPcOnPremisesConnectionable) {
    val, err := m.GetBackingStore().Get("onPremisesConnections")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcOnPremisesConnectionable)
    }
    return nil
}
// GetOrganizationSettings gets the organizationSettings property value. The Cloud PC organization settings for a tenant.
// returns a CloudPcOrganizationSettingsable when successful
func (m *VirtualEndpoint) GetOrganizationSettings()(CloudPcOrganizationSettingsable) {
    val, err := m.GetBackingStore().Get("organizationSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcOrganizationSettingsable)
    }
    return nil
}
// GetProvisioningPolicies gets the provisioningPolicies property value. Cloud PC provisioning policy.
// returns a []CloudPcProvisioningPolicyable when successful
func (m *VirtualEndpoint) GetProvisioningPolicies()([]CloudPcProvisioningPolicyable) {
    val, err := m.GetBackingStore().Get("provisioningPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcProvisioningPolicyable)
    }
    return nil
}
// GetReports gets the reports property value. Cloud PC related reports.
// returns a CloudPcReportsable when successful
func (m *VirtualEndpoint) GetReports()(CloudPcReportsable) {
    val, err := m.GetBackingStore().Get("reports")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcReportsable)
    }
    return nil
}
// GetServicePlans gets the servicePlans property value. Cloud PC service plans.
// returns a []CloudPcServicePlanable when successful
func (m *VirtualEndpoint) GetServicePlans()([]CloudPcServicePlanable) {
    val, err := m.GetBackingStore().Get("servicePlans")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcServicePlanable)
    }
    return nil
}
// GetSharedUseServicePlans gets the sharedUseServicePlans property value. The sharedUseServicePlans property
// returns a []CloudPcSharedUseServicePlanable when successful
func (m *VirtualEndpoint) GetSharedUseServicePlans()([]CloudPcSharedUseServicePlanable) {
    val, err := m.GetBackingStore().Get("sharedUseServicePlans")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcSharedUseServicePlanable)
    }
    return nil
}
// GetSnapshots gets the snapshots property value. Cloud PC snapshots.
// returns a []CloudPcSnapshotable when successful
func (m *VirtualEndpoint) GetSnapshots()([]CloudPcSnapshotable) {
    val, err := m.GetBackingStore().Get("snapshots")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcSnapshotable)
    }
    return nil
}
// GetSupportedRegions gets the supportedRegions property value. Cloud PC supported regions.
// returns a []CloudPcSupportedRegionable when successful
func (m *VirtualEndpoint) GetSupportedRegions()([]CloudPcSupportedRegionable) {
    val, err := m.GetBackingStore().Get("supportedRegions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcSupportedRegionable)
    }
    return nil
}
// GetUserSettings gets the userSettings property value. Cloud PC user settings.
// returns a []CloudPcUserSettingable when successful
func (m *VirtualEndpoint) GetUserSettings()([]CloudPcUserSettingable) {
    val, err := m.GetBackingStore().Get("userSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcUserSettingable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualEndpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuditEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuditEvents()))
        for i, v := range m.GetAuditEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("auditEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBulkActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBulkActions()))
        for i, v := range m.GetBulkActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("bulkActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPCs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudPCs()))
        for i, v := range m.GetCloudPCs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("cloudPCs", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("crossCloudGovernmentOrganizationMapping", m.GetCrossCloudGovernmentOrganizationMapping())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceImages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceImages()))
        for i, v := range m.GetDeviceImages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceImages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExternalPartnerSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExternalPartnerSettings()))
        for i, v := range m.GetExternalPartnerSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("externalPartnerSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFrontLineServicePlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFrontLineServicePlans()))
        for i, v := range m.GetFrontLineServicePlans() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("frontLineServicePlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGalleryImages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGalleryImages()))
        for i, v := range m.GetGalleryImages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("galleryImages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOnPremisesConnections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnPremisesConnections()))
        for i, v := range m.GetOnPremisesConnections() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProvisioningPolicies()))
        for i, v := range m.GetProvisioningPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServicePlans()))
        for i, v := range m.GetServicePlans() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("servicePlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSharedUseServicePlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSharedUseServicePlans()))
        for i, v := range m.GetSharedUseServicePlans() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sharedUseServicePlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSnapshots() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSnapshots()))
        for i, v := range m.GetSnapshots() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("snapshots", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSupportedRegions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSupportedRegions()))
        for i, v := range m.GetSupportedRegions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("supportedRegions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserSettings()))
        for i, v := range m.GetUserSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuditEvents sets the auditEvents property value. Cloud PC audit event.
func (m *VirtualEndpoint) SetAuditEvents(value []CloudPcAuditEventable)() {
    err := m.GetBackingStore().Set("auditEvents", value)
    if err != nil {
        panic(err)
    }
}
// SetBulkActions sets the bulkActions property value. The bulkActions property
func (m *VirtualEndpoint) SetBulkActions(value []CloudPcBulkActionable)() {
    err := m.GetBackingStore().Set("bulkActions", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPCs sets the cloudPCs property value. Cloud managed virtual desktops.
func (m *VirtualEndpoint) SetCloudPCs(value []CloudPCable)() {
    err := m.GetBackingStore().Set("cloudPCs", value)
    if err != nil {
        panic(err)
    }
}
// SetCrossCloudGovernmentOrganizationMapping sets the crossCloudGovernmentOrganizationMapping property value. Cloud PC organization mapping between public and US Government Community Cloud (GCC) organizations.
func (m *VirtualEndpoint) SetCrossCloudGovernmentOrganizationMapping(value CloudPcCrossCloudGovernmentOrganizationMappingable)() {
    err := m.GetBackingStore().Set("crossCloudGovernmentOrganizationMapping", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceImages sets the deviceImages property value. The image resource on Cloud PC.
func (m *VirtualEndpoint) SetDeviceImages(value []CloudPcDeviceImageable)() {
    err := m.GetBackingStore().Set("deviceImages", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalPartnerSettings sets the externalPartnerSettings property value. The external partner settings on a Cloud PC.
func (m *VirtualEndpoint) SetExternalPartnerSettings(value []CloudPcExternalPartnerSettingable)() {
    err := m.GetBackingStore().Set("externalPartnerSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetFrontLineServicePlans sets the frontLineServicePlans property value. Front-line service plans for a Cloud PC.
func (m *VirtualEndpoint) SetFrontLineServicePlans(value []CloudPcFrontLineServicePlanable)() {
    err := m.GetBackingStore().Set("frontLineServicePlans", value)
    if err != nil {
        panic(err)
    }
}
// SetGalleryImages sets the galleryImages property value. The gallery image resource on Cloud PC.
func (m *VirtualEndpoint) SetGalleryImages(value []CloudPcGalleryImageable)() {
    err := m.GetBackingStore().Set("galleryImages", value)
    if err != nil {
        panic(err)
    }
}
// SetOnPremisesConnections sets the onPremisesConnections property value. A defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *VirtualEndpoint) SetOnPremisesConnections(value []CloudPcOnPremisesConnectionable)() {
    err := m.GetBackingStore().Set("onPremisesConnections", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationSettings sets the organizationSettings property value. The Cloud PC organization settings for a tenant.
func (m *VirtualEndpoint) SetOrganizationSettings(value CloudPcOrganizationSettingsable)() {
    err := m.GetBackingStore().Set("organizationSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetProvisioningPolicies sets the provisioningPolicies property value. Cloud PC provisioning policy.
func (m *VirtualEndpoint) SetProvisioningPolicies(value []CloudPcProvisioningPolicyable)() {
    err := m.GetBackingStore().Set("provisioningPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetReports sets the reports property value. Cloud PC related reports.
func (m *VirtualEndpoint) SetReports(value CloudPcReportsable)() {
    err := m.GetBackingStore().Set("reports", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePlans sets the servicePlans property value. Cloud PC service plans.
func (m *VirtualEndpoint) SetServicePlans(value []CloudPcServicePlanable)() {
    err := m.GetBackingStore().Set("servicePlans", value)
    if err != nil {
        panic(err)
    }
}
// SetSharedUseServicePlans sets the sharedUseServicePlans property value. The sharedUseServicePlans property
func (m *VirtualEndpoint) SetSharedUseServicePlans(value []CloudPcSharedUseServicePlanable)() {
    err := m.GetBackingStore().Set("sharedUseServicePlans", value)
    if err != nil {
        panic(err)
    }
}
// SetSnapshots sets the snapshots property value. Cloud PC snapshots.
func (m *VirtualEndpoint) SetSnapshots(value []CloudPcSnapshotable)() {
    err := m.GetBackingStore().Set("snapshots", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedRegions sets the supportedRegions property value. Cloud PC supported regions.
func (m *VirtualEndpoint) SetSupportedRegions(value []CloudPcSupportedRegionable)() {
    err := m.GetBackingStore().Set("supportedRegions", value)
    if err != nil {
        panic(err)
    }
}
// SetUserSettings sets the userSettings property value. Cloud PC user settings.
func (m *VirtualEndpoint) SetUserSettings(value []CloudPcUserSettingable)() {
    err := m.GetBackingStore().Set("userSettings", value)
    if err != nil {
        panic(err)
    }
}
type VirtualEndpointable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuditEvents()([]CloudPcAuditEventable)
    GetBulkActions()([]CloudPcBulkActionable)
    GetCloudPCs()([]CloudPCable)
    GetCrossCloudGovernmentOrganizationMapping()(CloudPcCrossCloudGovernmentOrganizationMappingable)
    GetDeviceImages()([]CloudPcDeviceImageable)
    GetExternalPartnerSettings()([]CloudPcExternalPartnerSettingable)
    GetFrontLineServicePlans()([]CloudPcFrontLineServicePlanable)
    GetGalleryImages()([]CloudPcGalleryImageable)
    GetOnPremisesConnections()([]CloudPcOnPremisesConnectionable)
    GetOrganizationSettings()(CloudPcOrganizationSettingsable)
    GetProvisioningPolicies()([]CloudPcProvisioningPolicyable)
    GetReports()(CloudPcReportsable)
    GetServicePlans()([]CloudPcServicePlanable)
    GetSharedUseServicePlans()([]CloudPcSharedUseServicePlanable)
    GetSnapshots()([]CloudPcSnapshotable)
    GetSupportedRegions()([]CloudPcSupportedRegionable)
    GetUserSettings()([]CloudPcUserSettingable)
    SetAuditEvents(value []CloudPcAuditEventable)()
    SetBulkActions(value []CloudPcBulkActionable)()
    SetCloudPCs(value []CloudPCable)()
    SetCrossCloudGovernmentOrganizationMapping(value CloudPcCrossCloudGovernmentOrganizationMappingable)()
    SetDeviceImages(value []CloudPcDeviceImageable)()
    SetExternalPartnerSettings(value []CloudPcExternalPartnerSettingable)()
    SetFrontLineServicePlans(value []CloudPcFrontLineServicePlanable)()
    SetGalleryImages(value []CloudPcGalleryImageable)()
    SetOnPremisesConnections(value []CloudPcOnPremisesConnectionable)()
    SetOrganizationSettings(value CloudPcOrganizationSettingsable)()
    SetProvisioningPolicies(value []CloudPcProvisioningPolicyable)()
    SetReports(value CloudPcReportsable)()
    SetServicePlans(value []CloudPcServicePlanable)()
    SetSharedUseServicePlans(value []CloudPcSharedUseServicePlanable)()
    SetSnapshots(value []CloudPcSnapshotable)()
    SetSupportedRegions(value []CloudPcSupportedRegionable)()
    SetUserSettings(value []CloudPcUserSettingable)()
}
