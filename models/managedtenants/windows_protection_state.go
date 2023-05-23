package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// WindowsProtectionState 
type WindowsProtectionState struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewWindowsProtectionState instantiates a new WindowsProtectionState and sets the default values.
func NewWindowsProtectionState()(*WindowsProtectionState) {
    m := &WindowsProtectionState{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateWindowsProtectionStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsProtectionStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsProtectionState(), nil
}
// GetAntiMalwareVersion gets the antiMalwareVersion property value. The anti-malware version for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetAntiMalwareVersion()(*string) {
    val, err := m.GetBackingStore().Get("antiMalwareVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAttentionRequired gets the attentionRequired property value. A flag indicating whether attention is required for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetAttentionRequired()(*bool) {
    val, err := m.GetBackingStore().Get("attentionRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDeviceDeleted gets the deviceDeleted property value. A flag indicating whether the managed device has been deleted. Optional. Read-only.
func (m *WindowsProtectionState) GetDeviceDeleted()(*bool) {
    val, err := m.GetBackingStore().Get("deviceDeleted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDevicePropertyRefreshDateTime gets the devicePropertyRefreshDateTime property value. The date and time the device property has been refreshed. Optional. Read-only.
func (m *WindowsProtectionState) GetDevicePropertyRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("devicePropertyRefreshDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEngineVersion gets the engineVersion property value. The anti-virus engine version for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetEngineVersion()(*string) {
    val, err := m.GetBackingStore().Get("engineVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsProtectionState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["antiMalwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAntiMalwareVersion(val)
        }
        return nil
    }
    res["attentionRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttentionRequired(val)
        }
        return nil
    }
    res["deviceDeleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDeleted(val)
        }
        return nil
    }
    res["devicePropertyRefreshDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePropertyRefreshDateTime(val)
        }
        return nil
    }
    res["engineVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEngineVersion(val)
        }
        return nil
    }
    res["fullScanOverdue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullScanOverdue(val)
        }
        return nil
    }
    res["fullScanRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullScanRequired(val)
        }
        return nil
    }
    res["lastFullScanDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastFullScanDateTime(val)
        }
        return nil
    }
    res["lastFullScanSignatureVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastFullScanSignatureVersion(val)
        }
        return nil
    }
    res["lastQuickScanDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastQuickScanDateTime(val)
        }
        return nil
    }
    res["lastQuickScanSignatureVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastQuickScanSignatureVersion(val)
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
    res["lastReportedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastReportedDateTime(val)
        }
        return nil
    }
    res["malwareProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMalwareProtectionEnabled(val)
        }
        return nil
    }
    res["managedDeviceHealthState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceHealthState(val)
        }
        return nil
    }
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["managedDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["networkInspectionSystemEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkInspectionSystemEnabled(val)
        }
        return nil
    }
    res["quickScanOverdue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuickScanOverdue(val)
        }
        return nil
    }
    res["realTimeProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRealTimeProtectionEnabled(val)
        }
        return nil
    }
    res["rebootRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRebootRequired(val)
        }
        return nil
    }
    res["signatureUpdateOverdue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignatureUpdateOverdue(val)
        }
        return nil
    }
    res["signatureVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignatureVersion(val)
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
// GetFullScanOverdue gets the fullScanOverdue property value. A flag indicating whether quick scan is overdue for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetFullScanOverdue()(*bool) {
    val, err := m.GetBackingStore().Get("fullScanOverdue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFullScanRequired gets the fullScanRequired property value. A flag indicating whether full scan is overdue for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetFullScanRequired()(*bool) {
    val, err := m.GetBackingStore().Get("fullScanRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastFullScanDateTime gets the lastFullScanDateTime property value. The date and time a full scan was completed. Optional. Read-only.
func (m *WindowsProtectionState) GetLastFullScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastFullScanDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastFullScanSignatureVersion gets the lastFullScanSignatureVersion property value. The version anti-malware version used to perform the last full scan. Optional. Read-only.
func (m *WindowsProtectionState) GetLastFullScanSignatureVersion()(*string) {
    val, err := m.GetBackingStore().Get("lastFullScanSignatureVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastQuickScanDateTime gets the lastQuickScanDateTime property value. The date and time a quick scan was completed. Optional. Read-only.
func (m *WindowsProtectionState) GetLastQuickScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastQuickScanDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastQuickScanSignatureVersion gets the lastQuickScanSignatureVersion property value. The version anti-malware version used to perform the last full scan. Optional. Read-only.
func (m *WindowsProtectionState) GetLastQuickScanSignatureVersion()(*string) {
    val, err := m.GetBackingStore().Get("lastQuickScanSignatureVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *WindowsProtectionState) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastRefreshedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastReportedDateTime gets the lastReportedDateTime property value. The date and time the protection state was last reported for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastReportedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMalwareProtectionEnabled gets the malwareProtectionEnabled property value. A flag indicating whether malware protection is enabled for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetMalwareProtectionEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("malwareProtectionEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetManagedDeviceHealthState gets the managedDeviceHealthState property value. The health state for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetManagedDeviceHealthState()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceHealthState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagedDeviceId gets the managedDeviceId property value. The unique identifier for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetManagedDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagedDeviceName gets the managedDeviceName property value. The display name for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetManagedDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkInspectionSystemEnabled gets the networkInspectionSystemEnabled property value. A flag indicating whether the network inspection system is enabled. Optional. Read-only.
func (m *WindowsProtectionState) GetNetworkInspectionSystemEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("networkInspectionSystemEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetQuickScanOverdue gets the quickScanOverdue property value. A flag indicating weather a quick scan is overdue. Optional. Read-only.
func (m *WindowsProtectionState) GetQuickScanOverdue()(*bool) {
    val, err := m.GetBackingStore().Get("quickScanOverdue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRealTimeProtectionEnabled gets the realTimeProtectionEnabled property value. A flag indicating whether real time protection is enabled. Optional. Read-only.
func (m *WindowsProtectionState) GetRealTimeProtectionEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("realTimeProtectionEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRebootRequired gets the rebootRequired property value. A flag indicating whether a reboot is required. Optional. Read-only.
func (m *WindowsProtectionState) GetRebootRequired()(*bool) {
    val, err := m.GetBackingStore().Get("rebootRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSignatureUpdateOverdue gets the signatureUpdateOverdue property value. A flag indicating whether an signature update is overdue. Optional. Read-only.
func (m *WindowsProtectionState) GetSignatureUpdateOverdue()(*bool) {
    val, err := m.GetBackingStore().Get("signatureUpdateOverdue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSignatureVersion gets the signatureVersion property value. The signature version for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetSignatureVersion()(*string) {
    val, err := m.GetBackingStore().Get("signatureVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *WindowsProtectionState) GetTenantDisplayName()(*string) {
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
func (m *WindowsProtectionState) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsProtectionState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("antiMalwareVersion", m.GetAntiMalwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("attentionRequired", m.GetAttentionRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceDeleted", m.GetDeviceDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("devicePropertyRefreshDateTime", m.GetDevicePropertyRefreshDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("engineVersion", m.GetEngineVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fullScanOverdue", m.GetFullScanOverdue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fullScanRequired", m.GetFullScanRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastFullScanDateTime", m.GetLastFullScanDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastFullScanSignatureVersion", m.GetLastFullScanSignatureVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastQuickScanDateTime", m.GetLastQuickScanDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastQuickScanSignatureVersion", m.GetLastQuickScanSignatureVersion())
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
        err = writer.WriteTimeValue("lastReportedDateTime", m.GetLastReportedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("malwareProtectionEnabled", m.GetMalwareProtectionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceHealthState", m.GetManagedDeviceHealthState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceName", m.GetManagedDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("networkInspectionSystemEnabled", m.GetNetworkInspectionSystemEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("quickScanOverdue", m.GetQuickScanOverdue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("realTimeProtectionEnabled", m.GetRealTimeProtectionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("rebootRequired", m.GetRebootRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("signatureUpdateOverdue", m.GetSignatureUpdateOverdue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signatureVersion", m.GetSignatureVersion())
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
// SetAntiMalwareVersion sets the antiMalwareVersion property value. The anti-malware version for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetAntiMalwareVersion(value *string)() {
    err := m.GetBackingStore().Set("antiMalwareVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetAttentionRequired sets the attentionRequired property value. A flag indicating whether attention is required for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetAttentionRequired(value *bool)() {
    err := m.GetBackingStore().Set("attentionRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceDeleted sets the deviceDeleted property value. A flag indicating whether the managed device has been deleted. Optional. Read-only.
func (m *WindowsProtectionState) SetDeviceDeleted(value *bool)() {
    err := m.GetBackingStore().Set("deviceDeleted", value)
    if err != nil {
        panic(err)
    }
}
// SetDevicePropertyRefreshDateTime sets the devicePropertyRefreshDateTime property value. The date and time the device property has been refreshed. Optional. Read-only.
func (m *WindowsProtectionState) SetDevicePropertyRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("devicePropertyRefreshDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEngineVersion sets the engineVersion property value. The anti-virus engine version for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetEngineVersion(value *string)() {
    err := m.GetBackingStore().Set("engineVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetFullScanOverdue sets the fullScanOverdue property value. A flag indicating whether quick scan is overdue for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetFullScanOverdue(value *bool)() {
    err := m.GetBackingStore().Set("fullScanOverdue", value)
    if err != nil {
        panic(err)
    }
}
// SetFullScanRequired sets the fullScanRequired property value. A flag indicating whether full scan is overdue for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetFullScanRequired(value *bool)() {
    err := m.GetBackingStore().Set("fullScanRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetLastFullScanDateTime sets the lastFullScanDateTime property value. The date and time a full scan was completed. Optional. Read-only.
func (m *WindowsProtectionState) SetLastFullScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastFullScanDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastFullScanSignatureVersion sets the lastFullScanSignatureVersion property value. The version anti-malware version used to perform the last full scan. Optional. Read-only.
func (m *WindowsProtectionState) SetLastFullScanSignatureVersion(value *string)() {
    err := m.GetBackingStore().Set("lastFullScanSignatureVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetLastQuickScanDateTime sets the lastQuickScanDateTime property value. The date and time a quick scan was completed. Optional. Read-only.
func (m *WindowsProtectionState) SetLastQuickScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastQuickScanDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastQuickScanSignatureVersion sets the lastQuickScanSignatureVersion property value. The version anti-malware version used to perform the last full scan. Optional. Read-only.
func (m *WindowsProtectionState) SetLastQuickScanSignatureVersion(value *string)() {
    err := m.GetBackingStore().Set("lastQuickScanSignatureVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *WindowsProtectionState) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastRefreshedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. The date and time the protection state was last reported for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastReportedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMalwareProtectionEnabled sets the malwareProtectionEnabled property value. A flag indicating whether malware protection is enabled for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetMalwareProtectionEnabled(value *bool)() {
    err := m.GetBackingStore().Set("malwareProtectionEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceHealthState sets the managedDeviceHealthState property value. The health state for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetManagedDeviceHealthState(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceHealthState", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. The unique identifier for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetManagedDeviceId(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceName sets the managedDeviceName property value. The display name for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetManagedDeviceName(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkInspectionSystemEnabled sets the networkInspectionSystemEnabled property value. A flag indicating whether the network inspection system is enabled. Optional. Read-only.
func (m *WindowsProtectionState) SetNetworkInspectionSystemEnabled(value *bool)() {
    err := m.GetBackingStore().Set("networkInspectionSystemEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetQuickScanOverdue sets the quickScanOverdue property value. A flag indicating weather a quick scan is overdue. Optional. Read-only.
func (m *WindowsProtectionState) SetQuickScanOverdue(value *bool)() {
    err := m.GetBackingStore().Set("quickScanOverdue", value)
    if err != nil {
        panic(err)
    }
}
// SetRealTimeProtectionEnabled sets the realTimeProtectionEnabled property value. A flag indicating whether real time protection is enabled. Optional. Read-only.
func (m *WindowsProtectionState) SetRealTimeProtectionEnabled(value *bool)() {
    err := m.GetBackingStore().Set("realTimeProtectionEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetRebootRequired sets the rebootRequired property value. A flag indicating whether a reboot is required. Optional. Read-only.
func (m *WindowsProtectionState) SetRebootRequired(value *bool)() {
    err := m.GetBackingStore().Set("rebootRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetSignatureUpdateOverdue sets the signatureUpdateOverdue property value. A flag indicating whether an signature update is overdue. Optional. Read-only.
func (m *WindowsProtectionState) SetSignatureUpdateOverdue(value *bool)() {
    err := m.GetBackingStore().Set("signatureUpdateOverdue", value)
    if err != nil {
        panic(err)
    }
}
// SetSignatureVersion sets the signatureVersion property value. The signature version for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetSignatureVersion(value *string)() {
    err := m.GetBackingStore().Set("signatureVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *WindowsProtectionState) SetTenantDisplayName(value *string)() {
    err := m.GetBackingStore().Set("tenantDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *WindowsProtectionState) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// WindowsProtectionStateable 
type WindowsProtectionStateable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAntiMalwareVersion()(*string)
    GetAttentionRequired()(*bool)
    GetDeviceDeleted()(*bool)
    GetDevicePropertyRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEngineVersion()(*string)
    GetFullScanOverdue()(*bool)
    GetFullScanRequired()(*bool)
    GetLastFullScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastFullScanSignatureVersion()(*string)
    GetLastQuickScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastQuickScanSignatureVersion()(*string)
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMalwareProtectionEnabled()(*bool)
    GetManagedDeviceHealthState()(*string)
    GetManagedDeviceId()(*string)
    GetManagedDeviceName()(*string)
    GetNetworkInspectionSystemEnabled()(*bool)
    GetQuickScanOverdue()(*bool)
    GetRealTimeProtectionEnabled()(*bool)
    GetRebootRequired()(*bool)
    GetSignatureUpdateOverdue()(*bool)
    GetSignatureVersion()(*string)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    SetAntiMalwareVersion(value *string)()
    SetAttentionRequired(value *bool)()
    SetDeviceDeleted(value *bool)()
    SetDevicePropertyRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEngineVersion(value *string)()
    SetFullScanOverdue(value *bool)()
    SetFullScanRequired(value *bool)()
    SetLastFullScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastFullScanSignatureVersion(value *string)()
    SetLastQuickScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastQuickScanSignatureVersion(value *string)()
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMalwareProtectionEnabled(value *bool)()
    SetManagedDeviceHealthState(value *string)()
    SetManagedDeviceId(value *string)()
    SetManagedDeviceName(value *string)()
    SetNetworkInspectionSystemEnabled(value *bool)()
    SetQuickScanOverdue(value *bool)()
    SetRealTimeProtectionEnabled(value *bool)()
    SetRebootRequired(value *bool)()
    SetSignatureUpdateOverdue(value *bool)()
    SetSignatureVersion(value *string)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
}
