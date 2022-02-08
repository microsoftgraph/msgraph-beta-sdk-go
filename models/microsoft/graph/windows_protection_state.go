package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsProtectionState 
type WindowsProtectionState struct {
    Entity
    // Current anti malware version
    antiMalwareVersion *string;
    // Device malware list
    detectedMalwareState []WindowsDeviceMalwareState;
    // Computer's state (like clean or pending full scan or pending reboot etc). Possible values are: clean, fullScanPending, rebootPending, manualStepsPending, offlineScanPending, critical.
    deviceState *WindowsDeviceHealthState;
    // Current endpoint protection engine's version
    engineVersion *string;
    // Full scan overdue or not?
    fullScanOverdue *bool;
    // Full scan required or not?
    fullScanRequired *bool;
    // Indicates whether the device is a virtual machine.
    isVirtualMachine *bool;
    // Last quick scan datetime
    lastFullScanDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Last full scan signature version
    lastFullScanSignatureVersion *string;
    // Last quick scan datetime
    lastQuickScanDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Last quick scan signature version
    lastQuickScanSignatureVersion *string;
    // Last device health status reported time
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Anti malware is enabled or not
    malwareProtectionEnabled *bool;
    // Network inspection system enabled or not?
    networkInspectionSystemEnabled *bool;
    // Product Status of Windows Defender Antivirus. Possible values are: noStatus, serviceNotRunning, serviceStartedWithoutMalwareProtection, pendingFullScanDueToThreatAction, pendingRebootDueToThreatAction, pendingManualStepsDueToThreatAction, avSignaturesOutOfDate, asSignaturesOutOfDate, noQuickScanHappenedForSpecifiedPeriod, noFullScanHappenedForSpecifiedPeriod, systemInitiatedScanInProgress, systemInitiatedCleanInProgress, samplesPendingSubmission, productRunningInEvaluationMode, productRunningInNonGenuineMode, productExpired, offlineScanRequired, serviceShutdownAsPartOfSystemShutdown, threatRemediationFailedCritically, threatRemediationFailedNonCritically, noStatusFlagsSet, platformOutOfDate, platformUpdateInProgress, platformAboutToBeOutdated, signatureOrPlatformEndOfLifeIsPastOrIsImpending, windowsSModeSignaturesInUseOnNonWin10SInstall.
    productStatus *WindowsDefenderProductStatus;
    // Quick scan overdue or not?
    quickScanOverdue *bool;
    // Real time protection is enabled or not?
    realTimeProtectionEnabled *bool;
    // Reboot required or not?
    rebootRequired *bool;
    // Signature out of date or not?
    signatureUpdateOverdue *bool;
    // Current malware definitions version
    signatureVersion *string;
    // Indicates whether the Windows Defender tamper protection feature is enabled.
    tamperProtectionEnabled *bool;
}
// NewWindowsProtectionState instantiates a new windowsProtectionState and sets the default values.
func NewWindowsProtectionState()(*WindowsProtectionState) {
    m := &WindowsProtectionState{
        Entity: *NewEntity(),
    }
    return m
}
// GetAntiMalwareVersion gets the antiMalwareVersion property value. Current anti malware version
func (m *WindowsProtectionState) GetAntiMalwareVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.antiMalwareVersion
    }
}
// GetDetectedMalwareState gets the detectedMalwareState property value. Device malware list
func (m *WindowsProtectionState) GetDetectedMalwareState()([]WindowsDeviceMalwareState) {
    if m == nil {
        return nil
    } else {
        return m.detectedMalwareState
    }
}
// GetDeviceState gets the deviceState property value. Computer's state (like clean or pending full scan or pending reboot etc). Possible values are: clean, fullScanPending, rebootPending, manualStepsPending, offlineScanPending, critical.
func (m *WindowsProtectionState) GetDeviceState()(*WindowsDeviceHealthState) {
    if m == nil {
        return nil
    } else {
        return m.deviceState
    }
}
// GetEngineVersion gets the engineVersion property value. Current endpoint protection engine's version
func (m *WindowsProtectionState) GetEngineVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.engineVersion
    }
}
// GetFullScanOverdue gets the fullScanOverdue property value. Full scan overdue or not?
func (m *WindowsProtectionState) GetFullScanOverdue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fullScanOverdue
    }
}
// GetFullScanRequired gets the fullScanRequired property value. Full scan required or not?
func (m *WindowsProtectionState) GetFullScanRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fullScanRequired
    }
}
// GetIsVirtualMachine gets the isVirtualMachine property value. Indicates whether the device is a virtual machine.
func (m *WindowsProtectionState) GetIsVirtualMachine()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVirtualMachine
    }
}
// GetLastFullScanDateTime gets the lastFullScanDateTime property value. Last quick scan datetime
func (m *WindowsProtectionState) GetLastFullScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastFullScanDateTime
    }
}
// GetLastFullScanSignatureVersion gets the lastFullScanSignatureVersion property value. Last full scan signature version
func (m *WindowsProtectionState) GetLastFullScanSignatureVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastFullScanSignatureVersion
    }
}
// GetLastQuickScanDateTime gets the lastQuickScanDateTime property value. Last quick scan datetime
func (m *WindowsProtectionState) GetLastQuickScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastQuickScanDateTime
    }
}
// GetLastQuickScanSignatureVersion gets the lastQuickScanSignatureVersion property value. Last quick scan signature version
func (m *WindowsProtectionState) GetLastQuickScanSignatureVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastQuickScanSignatureVersion
    }
}
// GetLastReportedDateTime gets the lastReportedDateTime property value. Last device health status reported time
func (m *WindowsProtectionState) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastReportedDateTime
    }
}
// GetMalwareProtectionEnabled gets the malwareProtectionEnabled property value. Anti malware is enabled or not
func (m *WindowsProtectionState) GetMalwareProtectionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.malwareProtectionEnabled
    }
}
// GetNetworkInspectionSystemEnabled gets the networkInspectionSystemEnabled property value. Network inspection system enabled or not?
func (m *WindowsProtectionState) GetNetworkInspectionSystemEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.networkInspectionSystemEnabled
    }
}
// GetProductStatus gets the productStatus property value. Product Status of Windows Defender Antivirus. Possible values are: noStatus, serviceNotRunning, serviceStartedWithoutMalwareProtection, pendingFullScanDueToThreatAction, pendingRebootDueToThreatAction, pendingManualStepsDueToThreatAction, avSignaturesOutOfDate, asSignaturesOutOfDate, noQuickScanHappenedForSpecifiedPeriod, noFullScanHappenedForSpecifiedPeriod, systemInitiatedScanInProgress, systemInitiatedCleanInProgress, samplesPendingSubmission, productRunningInEvaluationMode, productRunningInNonGenuineMode, productExpired, offlineScanRequired, serviceShutdownAsPartOfSystemShutdown, threatRemediationFailedCritically, threatRemediationFailedNonCritically, noStatusFlagsSet, platformOutOfDate, platformUpdateInProgress, platformAboutToBeOutdated, signatureOrPlatformEndOfLifeIsPastOrIsImpending, windowsSModeSignaturesInUseOnNonWin10SInstall.
func (m *WindowsProtectionState) GetProductStatus()(*WindowsDefenderProductStatus) {
    if m == nil {
        return nil
    } else {
        return m.productStatus
    }
}
// GetQuickScanOverdue gets the quickScanOverdue property value. Quick scan overdue or not?
func (m *WindowsProtectionState) GetQuickScanOverdue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.quickScanOverdue
    }
}
// GetRealTimeProtectionEnabled gets the realTimeProtectionEnabled property value. Real time protection is enabled or not?
func (m *WindowsProtectionState) GetRealTimeProtectionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.realTimeProtectionEnabled
    }
}
// GetRebootRequired gets the rebootRequired property value. Reboot required or not?
func (m *WindowsProtectionState) GetRebootRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.rebootRequired
    }
}
// GetSignatureUpdateOverdue gets the signatureUpdateOverdue property value. Signature out of date or not?
func (m *WindowsProtectionState) GetSignatureUpdateOverdue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.signatureUpdateOverdue
    }
}
// GetSignatureVersion gets the signatureVersion property value. Current malware definitions version
func (m *WindowsProtectionState) GetSignatureVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signatureVersion
    }
}
// GetTamperProtectionEnabled gets the tamperProtectionEnabled property value. Indicates whether the Windows Defender tamper protection feature is enabled.
func (m *WindowsProtectionState) GetTamperProtectionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.tamperProtectionEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsProtectionState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["antiMalwareVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAntiMalwareVersion(val)
        }
        return nil
    }
    res["detectedMalwareState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDeviceMalwareState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDeviceMalwareState, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsDeviceMalwareState))
            }
            m.SetDetectedMalwareState(res)
        }
        return nil
    }
    res["deviceState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDeviceHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceState(val.(*WindowsDeviceHealthState))
        }
        return nil
    }
    res["engineVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEngineVersion(val)
        }
        return nil
    }
    res["fullScanOverdue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullScanOverdue(val)
        }
        return nil
    }
    res["fullScanRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullScanRequired(val)
        }
        return nil
    }
    res["isVirtualMachine"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVirtualMachine(val)
        }
        return nil
    }
    res["lastFullScanDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastFullScanDateTime(val)
        }
        return nil
    }
    res["lastFullScanSignatureVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastFullScanSignatureVersion(val)
        }
        return nil
    }
    res["lastQuickScanDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastQuickScanDateTime(val)
        }
        return nil
    }
    res["lastQuickScanSignatureVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastQuickScanSignatureVersion(val)
        }
        return nil
    }
    res["lastReportedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastReportedDateTime(val)
        }
        return nil
    }
    res["malwareProtectionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMalwareProtectionEnabled(val)
        }
        return nil
    }
    res["networkInspectionSystemEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkInspectionSystemEnabled(val)
        }
        return nil
    }
    res["productStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDefenderProductStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductStatus(val.(*WindowsDefenderProductStatus))
        }
        return nil
    }
    res["quickScanOverdue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuickScanOverdue(val)
        }
        return nil
    }
    res["realTimeProtectionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRealTimeProtectionEnabled(val)
        }
        return nil
    }
    res["rebootRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRebootRequired(val)
        }
        return nil
    }
    res["signatureUpdateOverdue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignatureUpdateOverdue(val)
        }
        return nil
    }
    res["signatureVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignatureVersion(val)
        }
        return nil
    }
    res["tamperProtectionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTamperProtectionEnabled(val)
        }
        return nil
    }
    return res
}
func (m *WindowsProtectionState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetDetectedMalwareState() != nil {
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
        cast := (*m.GetDeviceState()).String()
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
        cast := (*m.GetProductStatus()).String()
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
// SetAntiMalwareVersion sets the antiMalwareVersion property value. Current anti malware version
func (m *WindowsProtectionState) SetAntiMalwareVersion(value *string)() {
    if m != nil {
        m.antiMalwareVersion = value
    }
}
// SetDetectedMalwareState sets the detectedMalwareState property value. Device malware list
func (m *WindowsProtectionState) SetDetectedMalwareState(value []WindowsDeviceMalwareState)() {
    if m != nil {
        m.detectedMalwareState = value
    }
}
// SetDeviceState sets the deviceState property value. Computer's state (like clean or pending full scan or pending reboot etc). Possible values are: clean, fullScanPending, rebootPending, manualStepsPending, offlineScanPending, critical.
func (m *WindowsProtectionState) SetDeviceState(value *WindowsDeviceHealthState)() {
    if m != nil {
        m.deviceState = value
    }
}
// SetEngineVersion sets the engineVersion property value. Current endpoint protection engine's version
func (m *WindowsProtectionState) SetEngineVersion(value *string)() {
    if m != nil {
        m.engineVersion = value
    }
}
// SetFullScanOverdue sets the fullScanOverdue property value. Full scan overdue or not?
func (m *WindowsProtectionState) SetFullScanOverdue(value *bool)() {
    if m != nil {
        m.fullScanOverdue = value
    }
}
// SetFullScanRequired sets the fullScanRequired property value. Full scan required or not?
func (m *WindowsProtectionState) SetFullScanRequired(value *bool)() {
    if m != nil {
        m.fullScanRequired = value
    }
}
// SetIsVirtualMachine sets the isVirtualMachine property value. Indicates whether the device is a virtual machine.
func (m *WindowsProtectionState) SetIsVirtualMachine(value *bool)() {
    if m != nil {
        m.isVirtualMachine = value
    }
}
// SetLastFullScanDateTime sets the lastFullScanDateTime property value. Last quick scan datetime
func (m *WindowsProtectionState) SetLastFullScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastFullScanDateTime = value
    }
}
// SetLastFullScanSignatureVersion sets the lastFullScanSignatureVersion property value. Last full scan signature version
func (m *WindowsProtectionState) SetLastFullScanSignatureVersion(value *string)() {
    if m != nil {
        m.lastFullScanSignatureVersion = value
    }
}
// SetLastQuickScanDateTime sets the lastQuickScanDateTime property value. Last quick scan datetime
func (m *WindowsProtectionState) SetLastQuickScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastQuickScanDateTime = value
    }
}
// SetLastQuickScanSignatureVersion sets the lastQuickScanSignatureVersion property value. Last quick scan signature version
func (m *WindowsProtectionState) SetLastQuickScanSignatureVersion(value *string)() {
    if m != nil {
        m.lastQuickScanSignatureVersion = value
    }
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. Last device health status reported time
func (m *WindowsProtectionState) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastReportedDateTime = value
    }
}
// SetMalwareProtectionEnabled sets the malwareProtectionEnabled property value. Anti malware is enabled or not
func (m *WindowsProtectionState) SetMalwareProtectionEnabled(value *bool)() {
    if m != nil {
        m.malwareProtectionEnabled = value
    }
}
// SetNetworkInspectionSystemEnabled sets the networkInspectionSystemEnabled property value. Network inspection system enabled or not?
func (m *WindowsProtectionState) SetNetworkInspectionSystemEnabled(value *bool)() {
    if m != nil {
        m.networkInspectionSystemEnabled = value
    }
}
// SetProductStatus sets the productStatus property value. Product Status of Windows Defender Antivirus. Possible values are: noStatus, serviceNotRunning, serviceStartedWithoutMalwareProtection, pendingFullScanDueToThreatAction, pendingRebootDueToThreatAction, pendingManualStepsDueToThreatAction, avSignaturesOutOfDate, asSignaturesOutOfDate, noQuickScanHappenedForSpecifiedPeriod, noFullScanHappenedForSpecifiedPeriod, systemInitiatedScanInProgress, systemInitiatedCleanInProgress, samplesPendingSubmission, productRunningInEvaluationMode, productRunningInNonGenuineMode, productExpired, offlineScanRequired, serviceShutdownAsPartOfSystemShutdown, threatRemediationFailedCritically, threatRemediationFailedNonCritically, noStatusFlagsSet, platformOutOfDate, platformUpdateInProgress, platformAboutToBeOutdated, signatureOrPlatformEndOfLifeIsPastOrIsImpending, windowsSModeSignaturesInUseOnNonWin10SInstall.
func (m *WindowsProtectionState) SetProductStatus(value *WindowsDefenderProductStatus)() {
    if m != nil {
        m.productStatus = value
    }
}
// SetQuickScanOverdue sets the quickScanOverdue property value. Quick scan overdue or not?
func (m *WindowsProtectionState) SetQuickScanOverdue(value *bool)() {
    if m != nil {
        m.quickScanOverdue = value
    }
}
// SetRealTimeProtectionEnabled sets the realTimeProtectionEnabled property value. Real time protection is enabled or not?
func (m *WindowsProtectionState) SetRealTimeProtectionEnabled(value *bool)() {
    if m != nil {
        m.realTimeProtectionEnabled = value
    }
}
// SetRebootRequired sets the rebootRequired property value. Reboot required or not?
func (m *WindowsProtectionState) SetRebootRequired(value *bool)() {
    if m != nil {
        m.rebootRequired = value
    }
}
// SetSignatureUpdateOverdue sets the signatureUpdateOverdue property value. Signature out of date or not?
func (m *WindowsProtectionState) SetSignatureUpdateOverdue(value *bool)() {
    if m != nil {
        m.signatureUpdateOverdue = value
    }
}
// SetSignatureVersion sets the signatureVersion property value. Current malware definitions version
func (m *WindowsProtectionState) SetSignatureVersion(value *string)() {
    if m != nil {
        m.signatureVersion = value
    }
}
// SetTamperProtectionEnabled sets the tamperProtectionEnabled property value. Indicates whether the Windows Defender tamper protection feature is enabled.
func (m *WindowsProtectionState) SetTamperProtectionEnabled(value *bool)() {
    if m != nil {
        m.tamperProtectionEnabled = value
    }
}
