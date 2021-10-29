package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedDeviceComplianceTrend struct {
    Entity
    // The number of devices with a compliant status. Required. Read-only.
    compliantDeviceCount *int32;
    // The number of devices manged by Configuration Manager. Required. Read-only.
    configManagerDeviceCount *int32;
    // The date and time compliance snapshot was performed. Required. Read-only.
    countDateTime *string;
    // The number of devices with an error status. Required. Read-only.
    errorDeviceCount *int32;
    // The number of devices that are in a grace period status. Required. Read-only.
    inGracePeriodDeviceCount *int32;
    // The number of devices that are in a non-compliant status. Required. Read-only.
    noncompliantDeviceCount *int32;
    // The display name for the managed tenant. Optional. Read-only.
    tenantDisplayName *string;
    // The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
    tenantId *string;
    // The number of devices in an unknown status. Required. Read-only.
    unknownDeviceCount *int32;
}
// Instantiates a new managedDeviceComplianceTrend and sets the default values.
func NewManagedDeviceComplianceTrend()(*ManagedDeviceComplianceTrend) {
    m := &ManagedDeviceComplianceTrend{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the compliantDeviceCount property value. The number of devices with a compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
// Gets the configManagerDeviceCount property value. The number of devices manged by Configuration Manager. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetConfigManagerDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configManagerDeviceCount
    }
}
// Gets the countDateTime property value. The date and time compliance snapshot was performed. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetCountDateTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countDateTime
    }
}
// Gets the errorDeviceCount property value. The number of devices with an error status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// Gets the inGracePeriodDeviceCount property value. The number of devices that are in a grace period status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetInGracePeriodDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.inGracePeriodDeviceCount
    }
}
// Gets the noncompliantDeviceCount property value. The number of devices that are in a non-compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetNoncompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.noncompliantDeviceCount
    }
}
// Gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
// Gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Gets the unknownDeviceCount property value. The number of devices in an unknown status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
// The deserialization information for the current model
func (m *ManagedDeviceComplianceTrend) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCompliantDeviceCount(val)
        return nil
    }
    res["configManagerDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConfigManagerDeviceCount(val)
        return nil
    }
    res["countDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCountDateTime(val)
        return nil
    }
    res["errorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetErrorDeviceCount(val)
        return nil
    }
    res["inGracePeriodDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetInGracePeriodDeviceCount(val)
        return nil
    }
    res["noncompliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNoncompliantDeviceCount(val)
        return nil
    }
    res["tenantDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantDisplayName(val)
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
        return nil
    }
    res["unknownDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUnknownDeviceCount(val)
        return nil
    }
    return res
}
func (m *ManagedDeviceComplianceTrend) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagedDeviceComplianceTrend) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("compliantDeviceCount", m.GetCompliantDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("configManagerDeviceCount", m.GetConfigManagerDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countDateTime", m.GetCountDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorDeviceCount", m.GetErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("inGracePeriodDeviceCount", m.GetInGracePeriodDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("noncompliantDeviceCount", m.GetNoncompliantDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantDisplayName", m.GetTenantDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unknownDeviceCount", m.GetUnknownDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the compliantDeviceCount property value. The number of devices with a compliant status. Required. Read-only.
// Parameters:
//  - value : Value to set for the compliantDeviceCount property.
func (m *ManagedDeviceComplianceTrend) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
// Sets the configManagerDeviceCount property value. The number of devices manged by Configuration Manager. Required. Read-only.
// Parameters:
//  - value : Value to set for the configManagerDeviceCount property.
func (m *ManagedDeviceComplianceTrend) SetConfigManagerDeviceCount(value *int32)() {
    m.configManagerDeviceCount = value
}
// Sets the countDateTime property value. The date and time compliance snapshot was performed. Required. Read-only.
// Parameters:
//  - value : Value to set for the countDateTime property.
func (m *ManagedDeviceComplianceTrend) SetCountDateTime(value *string)() {
    m.countDateTime = value
}
// Sets the errorDeviceCount property value. The number of devices with an error status. Required. Read-only.
// Parameters:
//  - value : Value to set for the errorDeviceCount property.
func (m *ManagedDeviceComplianceTrend) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
// Sets the inGracePeriodDeviceCount property value. The number of devices that are in a grace period status. Required. Read-only.
// Parameters:
//  - value : Value to set for the inGracePeriodDeviceCount property.
func (m *ManagedDeviceComplianceTrend) SetInGracePeriodDeviceCount(value *int32)() {
    m.inGracePeriodDeviceCount = value
}
// Sets the noncompliantDeviceCount property value. The number of devices that are in a non-compliant status. Required. Read-only.
// Parameters:
//  - value : Value to set for the noncompliantDeviceCount property.
func (m *ManagedDeviceComplianceTrend) SetNoncompliantDeviceCount(value *int32)() {
    m.noncompliantDeviceCount = value
}
// Sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the tenantDisplayName property.
func (m *ManagedDeviceComplianceTrend) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
// Sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the tenantId property.
func (m *ManagedDeviceComplianceTrend) SetTenantId(value *string)() {
    m.tenantId = value
}
// Sets the unknownDeviceCount property value. The number of devices in an unknown status. Required. Read-only.
// Parameters:
//  - value : Value to set for the unknownDeviceCount property.
func (m *ManagedDeviceComplianceTrend) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
