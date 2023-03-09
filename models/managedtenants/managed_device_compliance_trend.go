package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagedDeviceComplianceTrend 
type ManagedDeviceComplianceTrend struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewManagedDeviceComplianceTrend instantiates a new managedDeviceComplianceTrend and sets the default values.
func NewManagedDeviceComplianceTrend()(*ManagedDeviceComplianceTrend) {
    m := &ManagedDeviceComplianceTrend{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateManagedDeviceComplianceTrendFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceComplianceTrendFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceComplianceTrend(), nil
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. The number of devices with a compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetCompliantDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("compliantDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetConfigManagerDeviceCount gets the configManagerDeviceCount property value. The number of devices manged by Configuration Manager. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetConfigManagerDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("configManagerDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCountDateTime gets the countDateTime property value. The date and time compliance snapshot was performed. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetCountDateTime()(*string) {
    val, err := m.GetBackingStore().Get("countDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetErrorDeviceCount gets the errorDeviceCount property value. The number of devices with an error status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetErrorDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("errorDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceComplianceTrend) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantDeviceCount(val)
        }
        return nil
    }
    res["configManagerDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigManagerDeviceCount(val)
        }
        return nil
    }
    res["countDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountDateTime(val)
        }
        return nil
    }
    res["errorDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDeviceCount(val)
        }
        return nil
    }
    res["inGracePeriodDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInGracePeriodDeviceCount(val)
        }
        return nil
    }
    res["noncompliantDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoncompliantDeviceCount(val)
        }
        return nil
    }
    res["tenantDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDisplayName(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["unknownDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetInGracePeriodDeviceCount gets the inGracePeriodDeviceCount property value. The number of devices that are in a grace period status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetInGracePeriodDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("inGracePeriodDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNoncompliantDeviceCount gets the noncompliantDeviceCount property value. The number of devices that are in a non-compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetNoncompliantDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("noncompliantDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) GetTenantDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("tenantDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. The number of devices in an unknown status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetUnknownDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("unknownDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceComplianceTrend) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    err := m.GetBackingStore().Set("compliantDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigManagerDeviceCount sets the configManagerDeviceCount property value. The number of devices manged by Configuration Manager. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetConfigManagerDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("configManagerDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCountDateTime sets the countDateTime property value. The date and time compliance snapshot was performed. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetCountDateTime(value *string)() {
    err := m.GetBackingStore().Set("countDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorDeviceCount sets the errorDeviceCount property value. The number of devices with an error status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetErrorDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("errorDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetInGracePeriodDeviceCount sets the inGracePeriodDeviceCount property value. The number of devices that are in a grace period status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetInGracePeriodDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("inGracePeriodDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNoncompliantDeviceCount sets the noncompliantDeviceCount property value. The number of devices that are in a non-compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetNoncompliantDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("noncompliantDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) SetTenantDisplayName(value *string)() {
    err := m.GetBackingStore().Set("tenantDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. The number of devices in an unknown status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetUnknownDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("unknownDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// ManagedDeviceComplianceTrendable 
type ManagedDeviceComplianceTrendable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompliantDeviceCount()(*int32)
    GetConfigManagerDeviceCount()(*int32)
    GetCountDateTime()(*string)
    GetErrorDeviceCount()(*int32)
    GetInGracePeriodDeviceCount()(*int32)
    GetNoncompliantDeviceCount()(*int32)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    GetUnknownDeviceCount()(*int32)
    SetCompliantDeviceCount(value *int32)()
    SetConfigManagerDeviceCount(value *int32)()
    SetCountDateTime(value *string)()
    SetErrorDeviceCount(value *int32)()
    SetInGracePeriodDeviceCount(value *int32)()
    SetNoncompliantDeviceCount(value *int32)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
    SetUnknownDeviceCount(value *int32)()
}
