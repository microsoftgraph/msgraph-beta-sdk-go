package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedDeviceComplianceTrend 
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
// NewManagedDeviceComplianceTrend instantiates a new managedDeviceComplianceTrend and sets the default values.
func NewManagedDeviceComplianceTrend()(*ManagedDeviceComplianceTrend) {
    m := &ManagedDeviceComplianceTrend{
        Entity: *NewEntity(),
    }
    return m
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. The number of devices with a compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
// GetConfigManagerDeviceCount gets the configManagerDeviceCount property value. The number of devices manged by Configuration Manager. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetConfigManagerDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configManagerDeviceCount
    }
}
// GetCountDateTime gets the countDateTime property value. The date and time compliance snapshot was performed. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetCountDateTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countDateTime
    }
}
// GetErrorDeviceCount gets the errorDeviceCount property value. The number of devices with an error status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// GetInGracePeriodDeviceCount gets the inGracePeriodDeviceCount property value. The number of devices that are in a grace period status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetInGracePeriodDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.inGracePeriodDeviceCount
    }
}
// GetNoncompliantDeviceCount gets the noncompliantDeviceCount property value. The number of devices that are in a non-compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetNoncompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.noncompliantDeviceCount
    }
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. The number of devices in an unknown status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceComplianceTrend) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantDeviceCount(val)
        }
        return nil
    }
    res["configManagerDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigManagerDeviceCount(val)
        }
        return nil
    }
    res["countDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountDateTime(val)
        }
        return nil
    }
    res["errorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDeviceCount(val)
        }
        return nil
    }
    res["inGracePeriodDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInGracePeriodDeviceCount(val)
        }
        return nil
    }
    res["noncompliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoncompliantDeviceCount(val)
        }
        return nil
    }
    res["tenantDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDisplayName(val)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["unknownDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownDeviceCount(val)
        }
        return nil
    }
    return res
}
func (m *ManagedDeviceComplianceTrend) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCompliantDeviceCount sets the compliantDeviceCount property value. The number of devices with a compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetCompliantDeviceCount(value *int32)() {
    if m != nil {
        m.compliantDeviceCount = value
    }
}
// SetConfigManagerDeviceCount sets the configManagerDeviceCount property value. The number of devices manged by Configuration Manager. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetConfigManagerDeviceCount(value *int32)() {
    if m != nil {
        m.configManagerDeviceCount = value
    }
}
// SetCountDateTime sets the countDateTime property value. The date and time compliance snapshot was performed. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetCountDateTime(value *string)() {
    if m != nil {
        m.countDateTime = value
    }
}
// SetErrorDeviceCount sets the errorDeviceCount property value. The number of devices with an error status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetErrorDeviceCount(value *int32)() {
    if m != nil {
        m.errorDeviceCount = value
    }
}
// SetInGracePeriodDeviceCount sets the inGracePeriodDeviceCount property value. The number of devices that are in a grace period status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetInGracePeriodDeviceCount(value *int32)() {
    if m != nil {
        m.inGracePeriodDeviceCount = value
    }
}
// SetNoncompliantDeviceCount sets the noncompliantDeviceCount property value. The number of devices that are in a non-compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetNoncompliantDeviceCount(value *int32)() {
    if m != nil {
        m.noncompliantDeviceCount = value
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) SetTenantDisplayName(value *string)() {
    if m != nil {
        m.tenantDisplayName = value
    }
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. The number of devices in an unknown status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetUnknownDeviceCount(value *int32)() {
    if m != nil {
        m.unknownDeviceCount = value
    }
}
