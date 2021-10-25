package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsProtectionState struct {
    Entity
    antiMalwareVersion *string;
    attentionRequired *bool;
    deviceDeleted *bool;
    devicePropertyRefreshDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    engineVersion *string;
    fullScanOverdue *bool;
    fullScanRequired *bool;
    lastFullScanDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastFullScanSignatureVersion *string;
    lastQuickScanDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastQuickScanSignatureVersion *string;
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    malwareProtectionEnabled *bool;
    managedDeviceHealthState *string;
    managedDeviceId *string;
    managedDeviceName *string;
    networkInspectionSystemEnabled *bool;
    quickScanOverdue *bool;
    realTimeProtectionEnabled *bool;
    rebootRequired *bool;
    signatureUpdateOverdue *bool;
    signatureVersion *string;
    tenantDisplayName *string;
    tenantId *string;
}
func NewWindowsProtectionState()(*WindowsProtectionState) {
    m := &WindowsProtectionState{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WindowsProtectionState) GetAntiMalwareVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.antiMalwareVersion
    }
}
func (m *WindowsProtectionState) GetAttentionRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.attentionRequired
    }
}
func (m *WindowsProtectionState) GetDeviceDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceDeleted
    }
}
func (m *WindowsProtectionState) GetDevicePropertyRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.devicePropertyRefreshDateTime
    }
}
func (m *WindowsProtectionState) GetEngineVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.engineVersion
    }
}
func (m *WindowsProtectionState) GetFullScanOverdue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fullScanOverdue
    }
}
func (m *WindowsProtectionState) GetFullScanRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fullScanRequired
    }
}
func (m *WindowsProtectionState) GetLastFullScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastFullScanDateTime
    }
}
func (m *WindowsProtectionState) GetLastFullScanSignatureVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastFullScanSignatureVersion
    }
}
func (m *WindowsProtectionState) GetLastQuickScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastQuickScanDateTime
    }
}
func (m *WindowsProtectionState) GetLastQuickScanSignatureVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastQuickScanSignatureVersion
    }
}
func (m *WindowsProtectionState) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshedDateTime
    }
}
func (m *WindowsProtectionState) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastReportedDateTime
    }
}
func (m *WindowsProtectionState) GetMalwareProtectionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.malwareProtectionEnabled
    }
}
func (m *WindowsProtectionState) GetManagedDeviceHealthState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceHealthState
    }
}
func (m *WindowsProtectionState) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
func (m *WindowsProtectionState) GetManagedDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceName
    }
}
func (m *WindowsProtectionState) GetNetworkInspectionSystemEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.networkInspectionSystemEnabled
    }
}
func (m *WindowsProtectionState) GetQuickScanOverdue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.quickScanOverdue
    }
}
func (m *WindowsProtectionState) GetRealTimeProtectionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.realTimeProtectionEnabled
    }
}
func (m *WindowsProtectionState) GetRebootRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.rebootRequired
    }
}
func (m *WindowsProtectionState) GetSignatureUpdateOverdue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.signatureUpdateOverdue
    }
}
func (m *WindowsProtectionState) GetSignatureVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signatureVersion
    }
}
func (m *WindowsProtectionState) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
func (m *WindowsProtectionState) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
func (m *WindowsProtectionState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["antiMalwareVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAntiMalwareVersion(val)
        return nil
    }
    res["attentionRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAttentionRequired(val)
        return nil
    }
    res["deviceDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDeviceDeleted(val)
        return nil
    }
    res["devicePropertyRefreshDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDevicePropertyRefreshDateTime(val)
        return nil
    }
    res["engineVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEngineVersion(val)
        return nil
    }
    res["fullScanOverdue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFullScanOverdue(val)
        return nil
    }
    res["fullScanRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFullScanRequired(val)
        return nil
    }
    res["lastFullScanDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastFullScanDateTime(val)
        return nil
    }
    res["lastFullScanSignatureVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastFullScanSignatureVersion(val)
        return nil
    }
    res["lastQuickScanDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastQuickScanDateTime(val)
        return nil
    }
    res["lastQuickScanSignatureVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastQuickScanSignatureVersion(val)
        return nil
    }
    res["lastRefreshedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastRefreshedDateTime(val)
        return nil
    }
    res["lastReportedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastReportedDateTime(val)
        return nil
    }
    res["malwareProtectionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMalwareProtectionEnabled(val)
        return nil
    }
    res["managedDeviceHealthState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceHealthState(val)
        return nil
    }
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceId(val)
        return nil
    }
    res["managedDeviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceName(val)
        return nil
    }
    res["networkInspectionSystemEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetNetworkInspectionSystemEnabled(val)
        return nil
    }
    res["quickScanOverdue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetQuickScanOverdue(val)
        return nil
    }
    res["realTimeProtectionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRealTimeProtectionEnabled(val)
        return nil
    }
    res["rebootRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRebootRequired(val)
        return nil
    }
    res["signatureUpdateOverdue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSignatureUpdateOverdue(val)
        return nil
    }
    res["signatureVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSignatureVersion(val)
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
    return res
}
func (m *WindowsProtectionState) IsNil()(bool) {
    return m == nil
}
func (m *WindowsProtectionState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *WindowsProtectionState) SetAntiMalwareVersion(value *string)() {
    m.antiMalwareVersion = value
}
func (m *WindowsProtectionState) SetAttentionRequired(value *bool)() {
    m.attentionRequired = value
}
func (m *WindowsProtectionState) SetDeviceDeleted(value *bool)() {
    m.deviceDeleted = value
}
func (m *WindowsProtectionState) SetDevicePropertyRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.devicePropertyRefreshDateTime = value
}
func (m *WindowsProtectionState) SetEngineVersion(value *string)() {
    m.engineVersion = value
}
func (m *WindowsProtectionState) SetFullScanOverdue(value *bool)() {
    m.fullScanOverdue = value
}
func (m *WindowsProtectionState) SetFullScanRequired(value *bool)() {
    m.fullScanRequired = value
}
func (m *WindowsProtectionState) SetLastFullScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastFullScanDateTime = value
}
func (m *WindowsProtectionState) SetLastFullScanSignatureVersion(value *string)() {
    m.lastFullScanSignatureVersion = value
}
func (m *WindowsProtectionState) SetLastQuickScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastQuickScanDateTime = value
}
func (m *WindowsProtectionState) SetLastQuickScanSignatureVersion(value *string)() {
    m.lastQuickScanSignatureVersion = value
}
func (m *WindowsProtectionState) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
func (m *WindowsProtectionState) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
func (m *WindowsProtectionState) SetMalwareProtectionEnabled(value *bool)() {
    m.malwareProtectionEnabled = value
}
func (m *WindowsProtectionState) SetManagedDeviceHealthState(value *string)() {
    m.managedDeviceHealthState = value
}
func (m *WindowsProtectionState) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
func (m *WindowsProtectionState) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
func (m *WindowsProtectionState) SetNetworkInspectionSystemEnabled(value *bool)() {
    m.networkInspectionSystemEnabled = value
}
func (m *WindowsProtectionState) SetQuickScanOverdue(value *bool)() {
    m.quickScanOverdue = value
}
func (m *WindowsProtectionState) SetRealTimeProtectionEnabled(value *bool)() {
    m.realTimeProtectionEnabled = value
}
func (m *WindowsProtectionState) SetRebootRequired(value *bool)() {
    m.rebootRequired = value
}
func (m *WindowsProtectionState) SetSignatureUpdateOverdue(value *bool)() {
    m.signatureUpdateOverdue = value
}
func (m *WindowsProtectionState) SetSignatureVersion(value *string)() {
    m.signatureVersion = value
}
func (m *WindowsProtectionState) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
func (m *WindowsProtectionState) SetTenantId(value *string)() {
    m.tenantId = value
}
