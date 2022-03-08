package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileThreatDefenseConnectorable 
type MobileThreatDefenseConnectorable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowPartnerToCollectIOSApplicationMetadata()(*bool)
    GetAllowPartnerToCollectIOSPersonalApplicationMetadata()(*bool)
    GetAndroidDeviceBlockedOnMissingPartnerData()(*bool)
    GetAndroidEnabled()(*bool)
    GetAndroidMobileApplicationManagementEnabled()(*bool)
    GetIosDeviceBlockedOnMissingPartnerData()(*bool)
    GetIosEnabled()(*bool)
    GetIosMobileApplicationManagementEnabled()(*bool)
    GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMacDeviceBlockedOnMissingPartnerData()(*bool)
    GetMacEnabled()(*bool)
    GetMicrosoftDefenderForEndpointAttachEnabled()(*bool)
    GetPartnerState()(*MobileThreatPartnerTenantState)
    GetPartnerUnresponsivenessThresholdInDays()(*int32)
    GetPartnerUnsupportedOsVersionBlocked()(*bool)
    GetWindowsDeviceBlockedOnMissingPartnerData()(*bool)
    GetWindowsEnabled()(*bool)
    SetAllowPartnerToCollectIOSApplicationMetadata(value *bool)()
    SetAllowPartnerToCollectIOSPersonalApplicationMetadata(value *bool)()
    SetAndroidDeviceBlockedOnMissingPartnerData(value *bool)()
    SetAndroidEnabled(value *bool)()
    SetAndroidMobileApplicationManagementEnabled(value *bool)()
    SetIosDeviceBlockedOnMissingPartnerData(value *bool)()
    SetIosEnabled(value *bool)()
    SetIosMobileApplicationManagementEnabled(value *bool)()
    SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMacDeviceBlockedOnMissingPartnerData(value *bool)()
    SetMacEnabled(value *bool)()
    SetMicrosoftDefenderForEndpointAttachEnabled(value *bool)()
    SetPartnerState(value *MobileThreatPartnerTenantState)()
    SetPartnerUnresponsivenessThresholdInDays(value *int32)()
    SetPartnerUnsupportedOsVersionBlocked(value *bool)()
    SetWindowsDeviceBlockedOnMissingPartnerData(value *bool)()
    SetWindowsEnabled(value *bool)()
}
