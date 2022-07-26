package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DeviceCompliancePolicySettingStateSummary provides operations to manage the collection of activityStatistics entities.
type DeviceCompliancePolicySettingStateSummary struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The number of devices in a conflict state. Optional. Read-only.
    conflictDeviceCount *int32
    // The number of devices in an error state. Optional. Read-only.
    errorDeviceCount *int32
    // The number of devices in a failed state. Optional. Read-only.
    failedDeviceCount *int32
    // The identifer for the Microsoft Intune account. Required. Read-only.
    intuneAccountId *string
    // The identifier for the Intune setting. Optional. Read-only.
    intuneSettingId *string
    // Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of devices in a not applicable state. Optional. Read-only.
    notApplicableDeviceCount *int32
    // The number of devices in a pending state. Optional. Read-only.
    pendingDeviceCount *int32
    // The type for the device compliance policy. Optional. Read-only.
    policyType *string
    // The name for the setting within the device compliance policy. Optional. Read-only.
    settingName *string
    // The number of devices in a succeeded state. Optional. Read-only.
    succeededDeviceCount *int32
    // The display name for the managed tenant. Required. Read-only.
    tenantDisplayName *string
    // The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
    tenantId *string
}
// NewDeviceCompliancePolicySettingStateSummary instantiates a new deviceCompliancePolicySettingStateSummary and sets the default values.
func NewDeviceCompliancePolicySettingStateSummary()(*DeviceCompliancePolicySettingStateSummary) {
    m := &DeviceCompliancePolicySettingStateSummary{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.deviceCompliancePolicySettingStateSummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceCompliancePolicySettingStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePolicySettingStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePolicySettingStateSummary(), nil
}
// GetConflictDeviceCount gets the conflictDeviceCount property value. The number of devices in a conflict state. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) GetConflictDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictDeviceCount
    }
}
// GetErrorDeviceCount gets the errorDeviceCount property value. The number of devices in an error state. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// GetFailedDeviceCount gets the failedDeviceCount property value. The number of devices in a failed state. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicySettingStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conflictDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictDeviceCount(val)
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
    res["failedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedDeviceCount(val)
        }
        return nil
    }
    res["intuneAccountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneAccountId(val)
        }
        return nil
    }
    res["intuneSettingId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneSettingId(val)
        }
        return nil
    }
    res["lastRefreshedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshedDateTime(val)
        }
        return nil
    }
    res["notApplicableDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableDeviceCount(val)
        }
        return nil
    }
    res["pendingDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingDeviceCount(val)
        }
        return nil
    }
    res["policyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyType(val)
        }
        return nil
    }
    res["settingName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    res["succeededDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSucceededDeviceCount(val)
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
    return res
}
// GetIntuneAccountId gets the intuneAccountId property value. The identifer for the Microsoft Intune account. Required. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) GetIntuneAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.intuneAccountId
    }
}
// GetIntuneSettingId gets the intuneSettingId property value. The identifier for the Intune setting. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) GetIntuneSettingId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.intuneSettingId
    }
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshedDateTime
    }
}
// GetNotApplicableDeviceCount gets the notApplicableDeviceCount property value. The number of devices in a not applicable state. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
// GetPendingDeviceCount gets the pendingDeviceCount property value. The number of devices in a pending state. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) GetPendingDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingDeviceCount
    }
}
// GetPolicyType gets the policyType property value. The type for the device compliance policy. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) GetPolicyType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyType
    }
}
// GetSettingName gets the settingName property value. The name for the setting within the device compliance policy. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// GetSucceededDeviceCount gets the succeededDeviceCount property value. The number of devices in a succeeded state. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) GetSucceededDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.succeededDeviceCount
    }
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Serialize serializes information the current object
func (m *DeviceCompliancePolicySettingStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("conflictDeviceCount", m.GetConflictDeviceCount())
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
        err = writer.WriteInt32Value("failedDeviceCount", m.GetFailedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("intuneAccountId", m.GetIntuneAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("intuneSettingId", m.GetIntuneSettingId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastRefreshedDateTime", m.GetLastRefreshedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableDeviceCount", m.GetNotApplicableDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pendingDeviceCount", m.GetPendingDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("policyType", m.GetPolicyType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("succeededDeviceCount", m.GetSucceededDeviceCount())
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
    return nil
}
// SetConflictDeviceCount sets the conflictDeviceCount property value. The number of devices in a conflict state. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) SetConflictDeviceCount(value *int32)() {
    if m != nil {
        m.conflictDeviceCount = value
    }
}
// SetErrorDeviceCount sets the errorDeviceCount property value. The number of devices in an error state. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) SetErrorDeviceCount(value *int32)() {
    if m != nil {
        m.errorDeviceCount = value
    }
}
// SetFailedDeviceCount sets the failedDeviceCount property value. The number of devices in a failed state. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) SetFailedDeviceCount(value *int32)() {
    if m != nil {
        m.failedDeviceCount = value
    }
}
// SetIntuneAccountId sets the intuneAccountId property value. The identifer for the Microsoft Intune account. Required. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) SetIntuneAccountId(value *string)() {
    if m != nil {
        m.intuneAccountId = value
    }
}
// SetIntuneSettingId sets the intuneSettingId property value. The identifier for the Intune setting. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) SetIntuneSettingId(value *string)() {
    if m != nil {
        m.intuneSettingId = value
    }
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastRefreshedDateTime = value
    }
}
// SetNotApplicableDeviceCount sets the notApplicableDeviceCount property value. The number of devices in a not applicable state. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) SetNotApplicableDeviceCount(value *int32)() {
    if m != nil {
        m.notApplicableDeviceCount = value
    }
}
// SetPendingDeviceCount sets the pendingDeviceCount property value. The number of devices in a pending state. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) SetPendingDeviceCount(value *int32)() {
    if m != nil {
        m.pendingDeviceCount = value
    }
}
// SetPolicyType sets the policyType property value. The type for the device compliance policy. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) SetPolicyType(value *string)() {
    if m != nil {
        m.policyType = value
    }
}
// SetSettingName sets the settingName property value. The name for the setting within the device compliance policy. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) SetSettingName(value *string)() {
    if m != nil {
        m.settingName = value
    }
}
// SetSucceededDeviceCount sets the succeededDeviceCount property value. The number of devices in a succeeded state. Optional. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) SetSucceededDeviceCount(value *int32)() {
    if m != nil {
        m.succeededDeviceCount = value
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) SetTenantDisplayName(value *string)() {
    if m != nil {
        m.tenantDisplayName = value
    }
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *DeviceCompliancePolicySettingStateSummary) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
