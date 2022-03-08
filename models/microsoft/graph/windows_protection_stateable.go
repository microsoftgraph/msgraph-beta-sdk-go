package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsProtectionStateable 
type WindowsProtectionStateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAntiMalwareVersion()(*string)
    GetDetectedMalwareState()([]WindowsDeviceMalwareStateable)
    GetDeviceState()(*WindowsDeviceHealthState)
    GetEngineVersion()(*string)
    GetFullScanOverdue()(*bool)
    GetFullScanRequired()(*bool)
    GetIsVirtualMachine()(*bool)
    GetLastFullScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastFullScanSignatureVersion()(*string)
    GetLastQuickScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastQuickScanSignatureVersion()(*string)
    GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMalwareProtectionEnabled()(*bool)
    GetNetworkInspectionSystemEnabled()(*bool)
    GetProductStatus()(*WindowsDefenderProductStatus)
    GetQuickScanOverdue()(*bool)
    GetRealTimeProtectionEnabled()(*bool)
    GetRebootRequired()(*bool)
    GetSignatureUpdateOverdue()(*bool)
    GetSignatureVersion()(*string)
    GetTamperProtectionEnabled()(*bool)
    SetAntiMalwareVersion(value *string)()
    SetDetectedMalwareState(value []WindowsDeviceMalwareStateable)()
    SetDeviceState(value *WindowsDeviceHealthState)()
    SetEngineVersion(value *string)()
    SetFullScanOverdue(value *bool)()
    SetFullScanRequired(value *bool)()
    SetIsVirtualMachine(value *bool)()
    SetLastFullScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastFullScanSignatureVersion(value *string)()
    SetLastQuickScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastQuickScanSignatureVersion(value *string)()
    SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMalwareProtectionEnabled(value *bool)()
    SetNetworkInspectionSystemEnabled(value *bool)()
    SetProductStatus(value *WindowsDefenderProductStatus)()
    SetQuickScanOverdue(value *bool)()
    SetRealTimeProtectionEnabled(value *bool)()
    SetRebootRequired(value *bool)()
    SetSignatureUpdateOverdue(value *bool)()
    SetSignatureVersion(value *string)()
    SetTamperProtectionEnabled(value *bool)()
}
