package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsProtectionState struct {
    Entity
    antiMalwareVersion *string;
    detectedMalwareState []WindowsDeviceMalwareState;
    deviceState *WindowsDeviceHealthState;
    engineVersion *string;
    fullScanOverdue *bool;
    fullScanRequired *bool;
    isVirtualMachine *bool;
    lastFullScanDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastFullScanSignatureVersion *string;
    lastQuickScanDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastQuickScanSignatureVersion *string;
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    malwareProtectionEnabled *bool;
    networkInspectionSystemEnabled *bool;
    productStatus *WindowsDefenderProductStatus;
    quickScanOverdue *bool;
    realTimeProtectionEnabled *bool;
    rebootRequired *bool;
    signatureUpdateOverdue *bool;
    signatureVersion *string;
    tamperProtectionEnabled *bool;
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
func (m *WindowsProtectionState) GetDetectedMalwareState()([]WindowsDeviceMalwareState) {
    if m == nil {
        return nil
    } else {
        return m.detectedMalwareState
    }
}
func (m *WindowsProtectionState) GetDeviceState()(*WindowsDeviceHealthState) {
    if m == nil {
        return nil
    } else {
        return m.deviceState
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
func (m *WindowsProtectionState) GetIsVirtualMachine()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVirtualMachine
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
func (m *WindowsProtectionState) GetNetworkInspectionSystemEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.networkInspectionSystemEnabled
    }
}
func (m *WindowsProtectionState) GetProductStatus()(*WindowsDefenderProductStatus) {
    if m == nil {
        return nil
    } else {
        return m.productStatus
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
func (m *WindowsProtectionState) GetTamperProtectionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.tamperProtectionEnabled
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
    res["detectedMalwareState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDeviceMalwareState() })
        if err != nil {
            return err
        }
        res := make([]WindowsDeviceMalwareState, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsDeviceMalwareState))
        }
        m.SetDetectedMalwareState(res)
        return nil
    }
    res["deviceState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDeviceHealthState)
        if err != nil {
            return err
        }
        cast := val.(WindowsDeviceHealthState)
        m.SetDeviceState(&cast)
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
    res["isVirtualMachine"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsVirtualMachine(val)
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
    res["networkInspectionSystemEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetNetworkInspectionSystemEnabled(val)
        return nil
    }
    res["productStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDefenderProductStatus)
        if err != nil {
            return err
        }
        cast := val.(WindowsDefenderProductStatus)
        m.SetProductStatus(&cast)
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
    res["tamperProtectionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetTamperProtectionEnabled(val)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDetectedMalwareState()))
        for i, v := range m.GetDetectedMalwareState() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("detectedMalwareState", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceState() != nil {
        cast := m.GetDeviceState().String()
        err = writer.WriteStringValue("deviceState", &cast)
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
        err = writer.WriteBoolValue("isVirtualMachine", m.GetIsVirtualMachine())
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
        err = writer.WriteBoolValue("networkInspectionSystemEnabled", m.GetNetworkInspectionSystemEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetProductStatus() != nil {
        cast := m.GetProductStatus().String()
        err = writer.WriteStringValue("productStatus", &cast)
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
        err = writer.WriteBoolValue("tamperProtectionEnabled", m.GetTamperProtectionEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WindowsProtectionState) SetAntiMalwareVersion(value *string)() {
    m.antiMalwareVersion = value
}
func (m *WindowsProtectionState) SetDetectedMalwareState(value []WindowsDeviceMalwareState)() {
    m.detectedMalwareState = value
}
func (m *WindowsProtectionState) SetDeviceState(value *WindowsDeviceHealthState)() {
    m.deviceState = value
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
func (m *WindowsProtectionState) SetIsVirtualMachine(value *bool)() {
    m.isVirtualMachine = value
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
func (m *WindowsProtectionState) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
func (m *WindowsProtectionState) SetMalwareProtectionEnabled(value *bool)() {
    m.malwareProtectionEnabled = value
}
func (m *WindowsProtectionState) SetNetworkInspectionSystemEnabled(value *bool)() {
    m.networkInspectionSystemEnabled = value
}
func (m *WindowsProtectionState) SetProductStatus(value *WindowsDefenderProductStatus)() {
    m.productStatus = value
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
func (m *WindowsProtectionState) SetTamperProtectionEnabled(value *bool)() {
    m.tamperProtectionEnabled = value
}
