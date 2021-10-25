package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type VirtualEndpoint struct {
    Entity
    auditEvents []CloudPcAuditEvent;
    cloudPCs []CloudPC;
    deviceImages []CloudPcDeviceImage;
    onPremisesConnections []CloudPcOnPremisesConnection;
    provisioningPolicies []CloudPcProvisioningPolicy;
    userSettings []CloudPcUserSetting;
}
func NewVirtualEndpoint()(*VirtualEndpoint) {
    m := &VirtualEndpoint{
        Entity: *NewEntity(),
    }
    return m
}
func (m *VirtualEndpoint) GetAuditEvents()([]CloudPcAuditEvent) {
    if m == nil {
        return nil
    } else {
        return m.auditEvents
    }
}
func (m *VirtualEndpoint) GetCloudPCs()([]CloudPC) {
    if m == nil {
        return nil
    } else {
        return m.cloudPCs
    }
}
func (m *VirtualEndpoint) GetDeviceImages()([]CloudPcDeviceImage) {
    if m == nil {
        return nil
    } else {
        return m.deviceImages
    }
}
func (m *VirtualEndpoint) GetOnPremisesConnections()([]CloudPcOnPremisesConnection) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesConnections
    }
}
func (m *VirtualEndpoint) GetProvisioningPolicies()([]CloudPcProvisioningPolicy) {
    if m == nil {
        return nil
    } else {
        return m.provisioningPolicies
    }
}
func (m *VirtualEndpoint) GetUserSettings()([]CloudPcUserSetting) {
    if m == nil {
        return nil
    } else {
        return m.userSettings
    }
}
func (m *VirtualEndpoint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["auditEvents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcAuditEvent() })
        if err != nil {
            return err
        }
        res := make([]CloudPcAuditEvent, len(val))
        for i, v := range val {
            res[i] = *(v.(*CloudPcAuditEvent))
        }
        m.SetAuditEvents(res)
        return nil
    }
    res["cloudPCs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPC() })
        if err != nil {
            return err
        }
        res := make([]CloudPC, len(val))
        for i, v := range val {
            res[i] = *(v.(*CloudPC))
        }
        m.SetCloudPCs(res)
        return nil
    }
    res["deviceImages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcDeviceImage() })
        if err != nil {
            return err
        }
        res := make([]CloudPcDeviceImage, len(val))
        for i, v := range val {
            res[i] = *(v.(*CloudPcDeviceImage))
        }
        m.SetDeviceImages(res)
        return nil
    }
    res["onPremisesConnections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcOnPremisesConnection() })
        if err != nil {
            return err
        }
        res := make([]CloudPcOnPremisesConnection, len(val))
        for i, v := range val {
            res[i] = *(v.(*CloudPcOnPremisesConnection))
        }
        m.SetOnPremisesConnections(res)
        return nil
    }
    res["provisioningPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcProvisioningPolicy() })
        if err != nil {
            return err
        }
        res := make([]CloudPcProvisioningPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*CloudPcProvisioningPolicy))
        }
        m.SetProvisioningPolicies(res)
        return nil
    }
    res["userSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcUserSetting() })
        if err != nil {
            return err
        }
        res := make([]CloudPcUserSetting, len(val))
        for i, v := range val {
            res[i] = *(v.(*CloudPcUserSetting))
        }
        m.SetUserSettings(res)
        return nil
    }
    return res
}
func (m *VirtualEndpoint) IsNil()(bool) {
    return m == nil
}
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
func (m *VirtualEndpoint) SetAuditEvents(value []CloudPcAuditEvent)() {
    m.auditEvents = value
}
func (m *VirtualEndpoint) SetCloudPCs(value []CloudPC)() {
    m.cloudPCs = value
}
func (m *VirtualEndpoint) SetDeviceImages(value []CloudPcDeviceImage)() {
    m.deviceImages = value
}
func (m *VirtualEndpoint) SetOnPremisesConnections(value []CloudPcOnPremisesConnection)() {
    m.onPremisesConnections = value
}
func (m *VirtualEndpoint) SetProvisioningPolicies(value []CloudPcProvisioningPolicy)() {
    m.provisioningPolicies = value
}
func (m *VirtualEndpoint) SetUserSettings(value []CloudPcUserSetting)() {
    m.userSettings = value
}
