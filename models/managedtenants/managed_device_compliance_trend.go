package managedtenants

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagedDeviceComplianceTrend provides operations to manage the collection of activityStatistics entities.
type ManagedDeviceComplianceTrend struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The number of devices with a compliant status. Required. Read-only.
    compliantDeviceCount *int32
    // The number of devices manged by Configuration Manager. Required. Read-only.
    configManagerDeviceCount *int32
    // The date and time compliance snapshot was performed. Required. Read-only.
    countDateTime *string
    // The number of devices with an error status. Required. Read-only.
    errorDeviceCount *int32
    // The number of devices that are in a grace period status. Required. Read-only.
    inGracePeriodDeviceCount *int32
    // The number of devices that are in a non-compliant status. Required. Read-only.
    noncompliantDeviceCount *int32
    // The display name for the managed tenant. Optional. Read-only.
    tenantDisplayName *string
    // The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
    tenantId *string
    // The number of devices in an unknown status. Required. Read-only.
    unknownDeviceCount *int32
}
// NewManagedDeviceComplianceTrend instantiates a new managedDeviceComplianceTrend and sets the default values.
func NewManagedDeviceComplianceTrend()(*ManagedDeviceComplianceTrend) {
    m := &ManagedDeviceComplianceTrend{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.managedDeviceComplianceTrend";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagedDeviceComplianceTrendFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceComplianceTrendFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceComplianceTrend(), nil
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. The number of devices with a compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetCompliantDeviceCount()(*int32) {
    return m.compliantDeviceCount
}
// GetConfigManagerDeviceCount gets the configManagerDeviceCount property value. The number of devices manged by Configuration Manager. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetConfigManagerDeviceCount()(*int32) {
    return m.configManagerDeviceCount
}
// GetCountDateTime gets the countDateTime property value. The date and time compliance snapshot was performed. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetCountDateTime()(*string) {
    return m.countDateTime
}
// GetErrorDeviceCount gets the errorDeviceCount property value. The number of devices with an error status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetErrorDeviceCount()(*int32) {
    return m.errorDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceComplianceTrend) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetCompliantDeviceCount)
    res["configManagerDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetConfigManagerDeviceCount)
    res["countDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCountDateTime)
    res["errorDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetErrorDeviceCount)
    res["inGracePeriodDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetInGracePeriodDeviceCount)
    res["noncompliantDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNoncompliantDeviceCount)
    res["tenantDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantDisplayName)
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    res["unknownDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUnknownDeviceCount)
    return res
}
// GetInGracePeriodDeviceCount gets the inGracePeriodDeviceCount property value. The number of devices that are in a grace period status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetInGracePeriodDeviceCount()(*int32) {
    return m.inGracePeriodDeviceCount
}
// GetNoncompliantDeviceCount gets the noncompliantDeviceCount property value. The number of devices that are in a non-compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetNoncompliantDeviceCount()(*int32) {
    return m.noncompliantDeviceCount
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) GetTenantDisplayName()(*string) {
    return m.tenantDisplayName
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) GetTenantId()(*string) {
    return m.tenantId
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. The number of devices in an unknown status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetUnknownDeviceCount()(*int32) {
    return m.unknownDeviceCount
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
    m.compliantDeviceCount = value
}
// SetConfigManagerDeviceCount sets the configManagerDeviceCount property value. The number of devices manged by Configuration Manager. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetConfigManagerDeviceCount(value *int32)() {
    m.configManagerDeviceCount = value
}
// SetCountDateTime sets the countDateTime property value. The date and time compliance snapshot was performed. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetCountDateTime(value *string)() {
    m.countDateTime = value
}
// SetErrorDeviceCount sets the errorDeviceCount property value. The number of devices with an error status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
// SetInGracePeriodDeviceCount sets the inGracePeriodDeviceCount property value. The number of devices that are in a grace period status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetInGracePeriodDeviceCount(value *int32)() {
    m.inGracePeriodDeviceCount = value
}
// SetNoncompliantDeviceCount sets the noncompliantDeviceCount property value. The number of devices that are in a non-compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetNoncompliantDeviceCount(value *int32)() {
    m.noncompliantDeviceCount = value
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. The number of devices in an unknown status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
