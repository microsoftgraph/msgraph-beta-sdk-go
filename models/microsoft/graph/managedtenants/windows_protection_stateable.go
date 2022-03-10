package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// WindowsProtectionStateable 
type WindowsProtectionStateable interface {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
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
