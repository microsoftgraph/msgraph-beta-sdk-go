package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MobileThreatDefenseConnector struct {
    Entity
    // For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about installed applications from Intune
    allowPartnerToCollectIOSApplicationMetadata *bool;
    // For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
    androidDeviceBlockedOnMissingPartnerData *bool;
    // For Android, set whether data from the data sync partner should be used during compliance evaluations
    androidEnabled *bool;
    // For Android, set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
    androidMobileApplicationManagementEnabled *bool;
    // For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
    iosDeviceBlockedOnMissingPartnerData *bool;
    // For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
    iosEnabled *bool;
    // For IOS, get or set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
    iosMobileApplicationManagementEnabled *bool;
    // DateTime of last Heartbeat recieved from the Data Sync Partner
    lastHeartbeatDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // For Mac, get or set whether Intune must receive data from the data sync partner prior to marking a device compliant
    macDeviceBlockedOnMissingPartnerData *bool;
    // For Mac, get or set whether data from the data sync partner should be used during compliance evaluations
    macEnabled *bool;
    // When TRUE, configuration profile management via Microsoft Defender for Endpoint is enabled. When FALSE, configuration profile management via Microsoft Defender for Endpoint is disabled.
    microsoftDefenderForEndpointAttachEnabled *bool;
    // Data Sync Partner state for this account. Possible values are: unavailable, available, enabled, unresponsive.
    partnerState *MobileThreatPartnerTenantState;
    // Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
    partnerUnresponsivenessThresholdInDays *int32;
    // Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
    partnerUnsupportedOsVersionBlocked *bool;
    // For Windows, set whether Intune must receive data from the data sync partner prior to marking a device compliant
    windowsDeviceBlockedOnMissingPartnerData *bool;
    // For Windows, get or set whether data from the data sync partner should be used during compliance evaluations
    windowsEnabled *bool;
}
// Instantiates a new mobileThreatDefenseConnector and sets the default values.
func NewMobileThreatDefenseConnector()(*MobileThreatDefenseConnector) {
    m := &MobileThreatDefenseConnector{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the allowPartnerToCollectIOSApplicationMetadata property value. For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about installed applications from Intune
func (m *MobileThreatDefenseConnector) GetAllowPartnerToCollectIOSApplicationMetadata()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowPartnerToCollectIOSApplicationMetadata
    }
}
// Gets the androidDeviceBlockedOnMissingPartnerData property value. For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetAndroidDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceBlockedOnMissingPartnerData
    }
}
// Gets the androidEnabled property value. For Android, set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetAndroidEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidEnabled
    }
}
// Gets the androidMobileApplicationManagementEnabled property value. For Android, set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
func (m *MobileThreatDefenseConnector) GetAndroidMobileApplicationManagementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidMobileApplicationManagementEnabled
    }
}
// Gets the iosDeviceBlockedOnMissingPartnerData property value. For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetIosDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosDeviceBlockedOnMissingPartnerData
    }
}
// Gets the iosEnabled property value. For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetIosEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosEnabled
    }
}
// Gets the iosMobileApplicationManagementEnabled property value. For IOS, get or set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
func (m *MobileThreatDefenseConnector) GetIosMobileApplicationManagementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosMobileApplicationManagementEnabled
    }
}
// Gets the lastHeartbeatDateTime property value. DateTime of last Heartbeat recieved from the Data Sync Partner
func (m *MobileThreatDefenseConnector) GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastHeartbeatDateTime
    }
}
// Gets the macDeviceBlockedOnMissingPartnerData property value. For Mac, get or set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetMacDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.macDeviceBlockedOnMissingPartnerData
    }
}
// Gets the macEnabled property value. For Mac, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetMacEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.macEnabled
    }
}
// Gets the microsoftDefenderForEndpointAttachEnabled property value. When TRUE, configuration profile management via Microsoft Defender for Endpoint is enabled. When FALSE, configuration profile management via Microsoft Defender for Endpoint is disabled.
func (m *MobileThreatDefenseConnector) GetMicrosoftDefenderForEndpointAttachEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.microsoftDefenderForEndpointAttachEnabled
    }
}
// Gets the partnerState property value. Data Sync Partner state for this account. Possible values are: unavailable, available, enabled, unresponsive.
func (m *MobileThreatDefenseConnector) GetPartnerState()(*MobileThreatPartnerTenantState) {
    if m == nil {
        return nil
    } else {
        return m.partnerState
    }
}
// Gets the partnerUnresponsivenessThresholdInDays property value. Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
func (m *MobileThreatDefenseConnector) GetPartnerUnresponsivenessThresholdInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.partnerUnresponsivenessThresholdInDays
    }
}
// Gets the partnerUnsupportedOsVersionBlocked property value. Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
func (m *MobileThreatDefenseConnector) GetPartnerUnsupportedOsVersionBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.partnerUnsupportedOsVersionBlocked
    }
}
// Gets the windowsDeviceBlockedOnMissingPartnerData property value. For Windows, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetWindowsDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.windowsDeviceBlockedOnMissingPartnerData
    }
}
// Gets the windowsEnabled property value. For Windows, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetWindowsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.windowsEnabled
    }
}
// The deserialization information for the current model
func (m *MobileThreatDefenseConnector) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowPartnerToCollectIOSApplicationMetadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowPartnerToCollectIOSApplicationMetadata(val)
        }
        return nil
    }
    res["androidDeviceBlockedOnMissingPartnerData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDeviceBlockedOnMissingPartnerData(val)
        }
        return nil
    }
    res["androidEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidEnabled(val)
        }
        return nil
    }
    res["androidMobileApplicationManagementEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidMobileApplicationManagementEnabled(val)
        }
        return nil
    }
    res["iosDeviceBlockedOnMissingPartnerData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosDeviceBlockedOnMissingPartnerData(val)
        }
        return nil
    }
    res["iosEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosEnabled(val)
        }
        return nil
    }
    res["iosMobileApplicationManagementEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosMobileApplicationManagementEnabled(val)
        }
        return nil
    }
    res["lastHeartbeatDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastHeartbeatDateTime(val)
        }
        return nil
    }
    res["macDeviceBlockedOnMissingPartnerData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacDeviceBlockedOnMissingPartnerData(val)
        }
        return nil
    }
    res["macEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacEnabled(val)
        }
        return nil
    }
    res["microsoftDefenderForEndpointAttachEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftDefenderForEndpointAttachEnabled(val)
        }
        return nil
    }
    res["partnerState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileThreatPartnerTenantState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MobileThreatPartnerTenantState)
            m.SetPartnerState(&cast)
        }
        return nil
    }
    res["partnerUnresponsivenessThresholdInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerUnresponsivenessThresholdInDays(val)
        }
        return nil
    }
    res["partnerUnsupportedOsVersionBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerUnsupportedOsVersionBlocked(val)
        }
        return nil
    }
    res["windowsDeviceBlockedOnMissingPartnerData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsDeviceBlockedOnMissingPartnerData(val)
        }
        return nil
    }
    res["windowsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsEnabled(val)
        }
        return nil
    }
    return res
}
func (m *MobileThreatDefenseConnector) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MobileThreatDefenseConnector) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowPartnerToCollectIOSApplicationMetadata", m.GetAllowPartnerToCollectIOSApplicationMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("androidDeviceBlockedOnMissingPartnerData", m.GetAndroidDeviceBlockedOnMissingPartnerData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("androidEnabled", m.GetAndroidEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("androidMobileApplicationManagementEnabled", m.GetAndroidMobileApplicationManagementEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iosDeviceBlockedOnMissingPartnerData", m.GetIosDeviceBlockedOnMissingPartnerData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iosEnabled", m.GetIosEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iosMobileApplicationManagementEnabled", m.GetIosMobileApplicationManagementEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastHeartbeatDateTime", m.GetLastHeartbeatDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("macDeviceBlockedOnMissingPartnerData", m.GetMacDeviceBlockedOnMissingPartnerData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("macEnabled", m.GetMacEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftDefenderForEndpointAttachEnabled", m.GetMicrosoftDefenderForEndpointAttachEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetPartnerState() != nil {
        cast := m.GetPartnerState().String()
        err = writer.WriteStringValue("partnerState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("partnerUnresponsivenessThresholdInDays", m.GetPartnerUnresponsivenessThresholdInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("partnerUnsupportedOsVersionBlocked", m.GetPartnerUnsupportedOsVersionBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsDeviceBlockedOnMissingPartnerData", m.GetWindowsDeviceBlockedOnMissingPartnerData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsEnabled", m.GetWindowsEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the allowPartnerToCollectIOSApplicationMetadata property value. For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about installed applications from Intune
// Parameters:
//  - value : Value to set for the allowPartnerToCollectIOSApplicationMetadata property.
func (m *MobileThreatDefenseConnector) SetAllowPartnerToCollectIOSApplicationMetadata(value *bool)() {
    m.allowPartnerToCollectIOSApplicationMetadata = value
}
// Sets the androidDeviceBlockedOnMissingPartnerData property value. For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
// Parameters:
//  - value : Value to set for the androidDeviceBlockedOnMissingPartnerData property.
func (m *MobileThreatDefenseConnector) SetAndroidDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.androidDeviceBlockedOnMissingPartnerData = value
}
// Sets the androidEnabled property value. For Android, set whether data from the data sync partner should be used during compliance evaluations
// Parameters:
//  - value : Value to set for the androidEnabled property.
func (m *MobileThreatDefenseConnector) SetAndroidEnabled(value *bool)() {
    m.androidEnabled = value
}
// Sets the androidMobileApplicationManagementEnabled property value. For Android, set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
// Parameters:
//  - value : Value to set for the androidMobileApplicationManagementEnabled property.
func (m *MobileThreatDefenseConnector) SetAndroidMobileApplicationManagementEnabled(value *bool)() {
    m.androidMobileApplicationManagementEnabled = value
}
// Sets the iosDeviceBlockedOnMissingPartnerData property value. For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
// Parameters:
//  - value : Value to set for the iosDeviceBlockedOnMissingPartnerData property.
func (m *MobileThreatDefenseConnector) SetIosDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.iosDeviceBlockedOnMissingPartnerData = value
}
// Sets the iosEnabled property value. For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
// Parameters:
//  - value : Value to set for the iosEnabled property.
func (m *MobileThreatDefenseConnector) SetIosEnabled(value *bool)() {
    m.iosEnabled = value
}
// Sets the iosMobileApplicationManagementEnabled property value. For IOS, get or set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
// Parameters:
//  - value : Value to set for the iosMobileApplicationManagementEnabled property.
func (m *MobileThreatDefenseConnector) SetIosMobileApplicationManagementEnabled(value *bool)() {
    m.iosMobileApplicationManagementEnabled = value
}
// Sets the lastHeartbeatDateTime property value. DateTime of last Heartbeat recieved from the Data Sync Partner
// Parameters:
//  - value : Value to set for the lastHeartbeatDateTime property.
func (m *MobileThreatDefenseConnector) SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastHeartbeatDateTime = value
}
// Sets the macDeviceBlockedOnMissingPartnerData property value. For Mac, get or set whether Intune must receive data from the data sync partner prior to marking a device compliant
// Parameters:
//  - value : Value to set for the macDeviceBlockedOnMissingPartnerData property.
func (m *MobileThreatDefenseConnector) SetMacDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.macDeviceBlockedOnMissingPartnerData = value
}
// Sets the macEnabled property value. For Mac, get or set whether data from the data sync partner should be used during compliance evaluations
// Parameters:
//  - value : Value to set for the macEnabled property.
func (m *MobileThreatDefenseConnector) SetMacEnabled(value *bool)() {
    m.macEnabled = value
}
// Sets the microsoftDefenderForEndpointAttachEnabled property value. When TRUE, configuration profile management via Microsoft Defender for Endpoint is enabled. When FALSE, configuration profile management via Microsoft Defender for Endpoint is disabled.
// Parameters:
//  - value : Value to set for the microsoftDefenderForEndpointAttachEnabled property.
func (m *MobileThreatDefenseConnector) SetMicrosoftDefenderForEndpointAttachEnabled(value *bool)() {
    m.microsoftDefenderForEndpointAttachEnabled = value
}
// Sets the partnerState property value. Data Sync Partner state for this account. Possible values are: unavailable, available, enabled, unresponsive.
// Parameters:
//  - value : Value to set for the partnerState property.
func (m *MobileThreatDefenseConnector) SetPartnerState(value *MobileThreatPartnerTenantState)() {
    m.partnerState = value
}
// Sets the partnerUnresponsivenessThresholdInDays property value. Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
// Parameters:
//  - value : Value to set for the partnerUnresponsivenessThresholdInDays property.
func (m *MobileThreatDefenseConnector) SetPartnerUnresponsivenessThresholdInDays(value *int32)() {
    m.partnerUnresponsivenessThresholdInDays = value
}
// Sets the partnerUnsupportedOsVersionBlocked property value. Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
// Parameters:
//  - value : Value to set for the partnerUnsupportedOsVersionBlocked property.
func (m *MobileThreatDefenseConnector) SetPartnerUnsupportedOsVersionBlocked(value *bool)() {
    m.partnerUnsupportedOsVersionBlocked = value
}
// Sets the windowsDeviceBlockedOnMissingPartnerData property value. For Windows, set whether Intune must receive data from the data sync partner prior to marking a device compliant
// Parameters:
//  - value : Value to set for the windowsDeviceBlockedOnMissingPartnerData property.
func (m *MobileThreatDefenseConnector) SetWindowsDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.windowsDeviceBlockedOnMissingPartnerData = value
}
// Sets the windowsEnabled property value. For Windows, get or set whether data from the data sync partner should be used during compliance evaluations
// Parameters:
//  - value : Value to set for the windowsEnabled property.
func (m *MobileThreatDefenseConnector) SetWindowsEnabled(value *bool)() {
    m.windowsEnabled = value
}
