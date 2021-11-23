package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// virtualEndpoint 
type VirtualEndpoint struct {
    Entity
    // Cloud PC audit event.
    auditEvents []CloudPcAuditEvent;
    // Cloud managed virtual desktops.
    cloudPCs []CloudPC;
    // The image resource on Cloud PC.
    deviceImages []CloudPcDeviceImage;
    // The gallery image resource on Cloud PC.
    galleryImages []CloudPcGalleryImage;
    // A defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
    onPremisesConnections []CloudPcOnPremisesConnection;
    // Cloud PC provisioning policy.
    provisioningPolicies []CloudPcProvisioningPolicy;
    // Cloud PC service plans.
    servicePlans []CloudPcServicePlan;
    // Cloud PC supported regions.
    supportedRegions []CloudPcSupportedRegion;
    // Cloud PC user settings.
    userSettings []CloudPcUserSetting;
}
// NewVirtualEndpoint instantiates a new virtualEndpoint and sets the default values.
func NewVirtualEndpoint()(*VirtualEndpoint) {
    m := &VirtualEndpoint{
        Entity: *NewEntity(),
    }
    return m
}
// GetAuditEvents gets the auditEvents property value. Cloud PC audit event.
func (m *VirtualEndpoint) GetAuditEvents()([]CloudPcAuditEvent) {
    if m == nil {
        return nil
    } else {
        return m.auditEvents
    }
}
// GetCloudPCs gets the cloudPCs property value. Cloud managed virtual desktops.
func (m *VirtualEndpoint) GetCloudPCs()([]CloudPC) {
    if m == nil {
        return nil
    } else {
        return m.cloudPCs
    }
}
// GetDeviceImages gets the deviceImages property value. The image resource on Cloud PC.
func (m *VirtualEndpoint) GetDeviceImages()([]CloudPcDeviceImage) {
    if m == nil {
        return nil
    } else {
        return m.deviceImages
    }
}
// GetGalleryImages gets the galleryImages property value. The gallery image resource on Cloud PC.
func (m *VirtualEndpoint) GetGalleryImages()([]CloudPcGalleryImage) {
    if m == nil {
        return nil
    } else {
        return m.galleryImages
    }
}
// GetOnPremisesConnections gets the onPremisesConnections property value. A defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *VirtualEndpoint) GetOnPremisesConnections()([]CloudPcOnPremisesConnection) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesConnections
    }
}
// GetProvisioningPolicies gets the provisioningPolicies property value. Cloud PC provisioning policy.
func (m *VirtualEndpoint) GetProvisioningPolicies()([]CloudPcProvisioningPolicy) {
    if m == nil {
        return nil
    } else {
        return m.provisioningPolicies
    }
}
// GetServicePlans gets the servicePlans property value. Cloud PC service plans.
func (m *VirtualEndpoint) GetServicePlans()([]CloudPcServicePlan) {
    if m == nil {
        return nil
    } else {
        return m.servicePlans
    }
}
// GetSupportedRegions gets the supportedRegions property value. Cloud PC supported regions.
func (m *VirtualEndpoint) GetSupportedRegions()([]CloudPcSupportedRegion) {
    if m == nil {
        return nil
    } else {
        return m.supportedRegions
    }
}
// GetUserSettings gets the userSettings property value. Cloud PC user settings.
func (m *VirtualEndpoint) GetUserSettings()([]CloudPcUserSetting) {
    if m == nil {
        return nil
    } else {
        return m.userSettings
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEndpoint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["auditEvents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcAuditEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcAuditEvent, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcAuditEvent))
            }
            m.SetAuditEvents(res)
        }
        return nil
    }
    res["cloudPCs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPC() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPC, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPC))
            }
            m.SetCloudPCs(res)
        }
        return nil
    }
    res["deviceImages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcDeviceImage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcDeviceImage, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcDeviceImage))
            }
            m.SetDeviceImages(res)
        }
        return nil
    }
    res["galleryImages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcGalleryImage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcGalleryImage, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcGalleryImage))
            }
            m.SetGalleryImages(res)
        }
        return nil
    }
    res["onPremisesConnections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcOnPremisesConnection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcOnPremisesConnection, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcOnPremisesConnection))
            }
            m.SetOnPremisesConnections(res)
        }
        return nil
    }
    res["provisioningPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcProvisioningPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcProvisioningPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcProvisioningPolicy))
            }
            m.SetProvisioningPolicies(res)
        }
        return nil
    }
    res["servicePlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcServicePlan() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcServicePlan, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcServicePlan))
            }
            m.SetServicePlans(res)
        }
        return nil
    }
    res["supportedRegions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcSupportedRegion() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcSupportedRegion, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcSupportedRegion))
            }
            m.SetSupportedRegions(res)
        }
        return nil
    }
    res["userSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcUserSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcUserSetting, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcUserSetting))
            }
            m.SetUserSettings(res)
        }
        return nil
    }
    return res
}
func (m *VirtualEndpoint) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *VirtualEndpoint) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAuditEvents()))
        for i, v := range m.GetAuditEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("auditEvents", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCloudPCs()))
        for i, v := range m.GetCloudPCs() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPCs", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceImages()))
        for i, v := range m.GetDeviceImages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceImages", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGalleryImages()))
        for i, v := range m.GetGalleryImages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("galleryImages", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOnPremisesConnections()))
        for i, v := range m.GetOnPremisesConnections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("onPremisesConnections", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProvisioningPolicies()))
        for i, v := range m.GetProvisioningPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("provisioningPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetServicePlans()))
        for i, v := range m.GetServicePlans() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("servicePlans", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSupportedRegions()))
        for i, v := range m.GetSupportedRegions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("supportedRegions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserSettings()))
        for i, v := range m.GetUserSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuditEvents sets the auditEvents property value. Cloud PC audit event.
func (m *VirtualEndpoint) SetAuditEvents(value []CloudPcAuditEvent)() {
    m.auditEvents = value
}
// SetCloudPCs sets the cloudPCs property value. Cloud managed virtual desktops.
func (m *VirtualEndpoint) SetCloudPCs(value []CloudPC)() {
    m.cloudPCs = value
}
// SetDeviceImages sets the deviceImages property value. The image resource on Cloud PC.
func (m *VirtualEndpoint) SetDeviceImages(value []CloudPcDeviceImage)() {
    m.deviceImages = value
}
// SetGalleryImages sets the galleryImages property value. The gallery image resource on Cloud PC.
func (m *VirtualEndpoint) SetGalleryImages(value []CloudPcGalleryImage)() {
    m.galleryImages = value
}
// SetOnPremisesConnections sets the onPremisesConnections property value. A defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *VirtualEndpoint) SetOnPremisesConnections(value []CloudPcOnPremisesConnection)() {
    m.onPremisesConnections = value
}
// SetProvisioningPolicies sets the provisioningPolicies property value. Cloud PC provisioning policy.
func (m *VirtualEndpoint) SetProvisioningPolicies(value []CloudPcProvisioningPolicy)() {
    m.provisioningPolicies = value
}
// SetServicePlans sets the servicePlans property value. Cloud PC service plans.
func (m *VirtualEndpoint) SetServicePlans(value []CloudPcServicePlan)() {
    m.servicePlans = value
}
// SetSupportedRegions sets the supportedRegions property value. Cloud PC supported regions.
func (m *VirtualEndpoint) SetSupportedRegions(value []CloudPcSupportedRegion)() {
    m.supportedRegions = value
}
// SetUserSettings sets the userSettings property value. Cloud PC user settings.
func (m *VirtualEndpoint) SetUserSettings(value []CloudPcUserSetting)() {
    m.userSettings = value
}
