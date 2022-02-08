package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileThreatDefenseConnector 
type MobileThreatDefenseConnector struct {
    Entity
    // For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about installed applications from Intune
    allowPartnerToCollectIOSApplicationMetadata *bool;
    // For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about personally installed applications from Intune
    allowPartnerToCollectIOSPersonalApplicationMetadata *bool;
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
// NewMobileThreatDefenseConnector instantiates a new mobileThreatDefenseConnector and sets the default values.
func NewMobileThreatDefenseConnector()(*MobileThreatDefenseConnector) {
    m := &MobileThreatDefenseConnector{
        Entity: *NewEntity(),
    }
    return m
}
// GetAllowPartnerToCollectIOSApplicationMetadata gets the allowPartnerToCollectIOSApplicationMetadata property value. For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about installed applications from Intune
func (m *MobileThreatDefenseConnector) GetAllowPartnerToCollectIOSApplicationMetadata()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowPartnerToCollectIOSApplicationMetadata
    }
}
// GetAllowPartnerToCollectIOSPersonalApplicationMetadata gets the allowPartnerToCollectIOSPersonalApplicationMetadata property value. For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about personally installed applications from Intune
func (m *MobileThreatDefenseConnector) GetAllowPartnerToCollectIOSPersonalApplicationMetadata()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowPartnerToCollectIOSPersonalApplicationMetadata
    }
}
// GetAndroidDeviceBlockedOnMissingPartnerData gets the androidDeviceBlockedOnMissingPartnerData property value. For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetAndroidDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceBlockedOnMissingPartnerData
    }
}
// GetAndroidEnabled gets the androidEnabled property value. For Android, set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetAndroidEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidEnabled
    }
}
// GetAndroidMobileApplicationManagementEnabled gets the androidMobileApplicationManagementEnabled property value. For Android, set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
func (m *MobileThreatDefenseConnector) GetAndroidMobileApplicationManagementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidMobileApplicationManagementEnabled
    }
}
// GetIosDeviceBlockedOnMissingPartnerData gets the iosDeviceBlockedOnMissingPartnerData property value. For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetIosDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosDeviceBlockedOnMissingPartnerData
    }
}
// GetIosEnabled gets the iosEnabled property value. For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetIosEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosEnabled
    }
}
// GetIosMobileApplicationManagementEnabled gets the iosMobileApplicationManagementEnabled property value. For IOS, get or set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
func (m *MobileThreatDefenseConnector) GetIosMobileApplicationManagementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosMobileApplicationManagementEnabled
    }
}
// GetLastHeartbeatDateTime gets the lastHeartbeatDateTime property value. DateTime of last Heartbeat recieved from the Data Sync Partner
func (m *MobileThreatDefenseConnector) GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastHeartbeatDateTime
    }
}
// GetMacDeviceBlockedOnMissingPartnerData gets the macDeviceBlockedOnMissingPartnerData property value. For Mac, get or set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetMacDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.macDeviceBlockedOnMissingPartnerData
    }
}
// GetMacEnabled gets the macEnabled property value. For Mac, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetMacEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.macEnabled
    }
}
// GetMicrosoftDefenderForEndpointAttachEnabled gets the microsoftDefenderForEndpointAttachEnabled property value. When TRUE, configuration profile management via Microsoft Defender for Endpoint is enabled. When FALSE, configuration profile management via Microsoft Defender for Endpoint is disabled.
func (m *MobileThreatDefenseConnector) GetMicrosoftDefenderForEndpointAttachEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.microsoftDefenderForEndpointAttachEnabled
    }
}
// GetPartnerState gets the partnerState property value. Data Sync Partner state for this account. Possible values are: unavailable, available, enabled, unresponsive.
func (m *MobileThreatDefenseConnector) GetPartnerState()(*MobileThreatPartnerTenantState) {
    if m == nil {
        return nil
    } else {
        return m.partnerState
    }
}
// GetPartnerUnresponsivenessThresholdInDays gets the partnerUnresponsivenessThresholdInDays property value. Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
func (m *MobileThreatDefenseConnector) GetPartnerUnresponsivenessThresholdInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.partnerUnresponsivenessThresholdInDays
    }
}
// GetPartnerUnsupportedOsVersionBlocked gets the partnerUnsupportedOsVersionBlocked property value. Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
func (m *MobileThreatDefenseConnector) GetPartnerUnsupportedOsVersionBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.partnerUnsupportedOsVersionBlocked
    }
}
// GetWindowsDeviceBlockedOnMissingPartnerData gets the windowsDeviceBlockedOnMissingPartnerData property value. For Windows, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetWindowsDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.windowsDeviceBlockedOnMissingPartnerData
    }
}
// GetWindowsEnabled gets the windowsEnabled property value. For Windows, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetWindowsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.windowsEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
    res["allowPartnerToCollectIOSPersonalApplicationMetadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowPartnerToCollectIOSPersonalApplicationMetadata(val)
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
            m.SetPartnerState(val.(*MobileThreatPartnerTenantState))
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
// Serialize serializes information the current object
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
        err = writer.WriteBoolValue("allowPartnerToCollectIOSPersonalApplicationMetadata", m.GetAllowPartnerToCollectIOSPersonalApplicationMetadata())
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
        cast := (*m.GetPartnerState()).String()
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
// SetAllowPartnerToCollectIOSApplicationMetadata sets the allowPartnerToCollectIOSApplicationMetadata property value. For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about installed applications from Intune
func (m *MobileThreatDefenseConnector) SetAllowPartnerToCollectIOSApplicationMetadata(value *bool)() {
    if m != nil {
        m.allowPartnerToCollectIOSApplicationMetadata = value
    }
}
// SetAllowPartnerToCollectIOSPersonalApplicationMetadata sets the allowPartnerToCollectIOSPersonalApplicationMetadata property value. For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about personally installed applications from Intune
func (m *MobileThreatDefenseConnector) SetAllowPartnerToCollectIOSPersonalApplicationMetadata(value *bool)() {
    if m != nil {
        m.allowPartnerToCollectIOSPersonalApplicationMetadata = value
    }
}
// SetAndroidDeviceBlockedOnMissingPartnerData sets the androidDeviceBlockedOnMissingPartnerData property value. For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) SetAndroidDeviceBlockedOnMissingPartnerData(value *bool)() {
    if m != nil {
        m.androidDeviceBlockedOnMissingPartnerData = value
    }
}
// SetAndroidEnabled sets the androidEnabled property value. For Android, set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) SetAndroidEnabled(value *bool)() {
    if m != nil {
        m.androidEnabled = value
    }
}
// SetAndroidMobileApplicationManagementEnabled sets the androidMobileApplicationManagementEnabled property value. For Android, set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
func (m *MobileThreatDefenseConnector) SetAndroidMobileApplicationManagementEnabled(value *bool)() {
    if m != nil {
        m.androidMobileApplicationManagementEnabled = value
    }
}
// SetIosDeviceBlockedOnMissingPartnerData sets the iosDeviceBlockedOnMissingPartnerData property value. For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) SetIosDeviceBlockedOnMissingPartnerData(value *bool)() {
    if m != nil {
        m.iosDeviceBlockedOnMissingPartnerData = value
    }
}
// SetIosEnabled sets the iosEnabled property value. For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) SetIosEnabled(value *bool)() {
    if m != nil {
        m.iosEnabled = value
    }
}
// SetIosMobileApplicationManagementEnabled sets the iosMobileApplicationManagementEnabled property value. For IOS, get or set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
func (m *MobileThreatDefenseConnector) SetIosMobileApplicationManagementEnabled(value *bool)() {
    if m != nil {
        m.iosMobileApplicationManagementEnabled = value
    }
}
// SetLastHeartbeatDateTime sets the lastHeartbeatDateTime property value. DateTime of last Heartbeat recieved from the Data Sync Partner
func (m *MobileThreatDefenseConnector) SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastHeartbeatDateTime = value
    }
}
// SetMacDeviceBlockedOnMissingPartnerData sets the macDeviceBlockedOnMissingPartnerData property value. For Mac, get or set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) SetMacDeviceBlockedOnMissingPartnerData(value *bool)() {
    if m != nil {
        m.macDeviceBlockedOnMissingPartnerData = value
    }
}
// SetMacEnabled sets the macEnabled property value. For Mac, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) SetMacEnabled(value *bool)() {
    if m != nil {
        m.macEnabled = value
    }
}
// SetMicrosoftDefenderForEndpointAttachEnabled sets the microsoftDefenderForEndpointAttachEnabled property value. When TRUE, configuration profile management via Microsoft Defender for Endpoint is enabled. When FALSE, configuration profile management via Microsoft Defender for Endpoint is disabled.
func (m *MobileThreatDefenseConnector) SetMicrosoftDefenderForEndpointAttachEnabled(value *bool)() {
    if m != nil {
        m.microsoftDefenderForEndpointAttachEnabled = value
    }
}
// SetPartnerState sets the partnerState property value. Data Sync Partner state for this account. Possible values are: unavailable, available, enabled, unresponsive.
func (m *MobileThreatDefenseConnector) SetPartnerState(value *MobileThreatPartnerTenantState)() {
    if m != nil {
        m.partnerState = value
    }
}
// SetPartnerUnresponsivenessThresholdInDays sets the partnerUnresponsivenessThresholdInDays property value. Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
func (m *MobileThreatDefenseConnector) SetPartnerUnresponsivenessThresholdInDays(value *int32)() {
    if m != nil {
        m.partnerUnresponsivenessThresholdInDays = value
    }
}
// SetPartnerUnsupportedOsVersionBlocked sets the partnerUnsupportedOsVersionBlocked property value. Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
func (m *MobileThreatDefenseConnector) SetPartnerUnsupportedOsVersionBlocked(value *bool)() {
    if m != nil {
        m.partnerUnsupportedOsVersionBlocked = value
    }
}
// SetWindowsDeviceBlockedOnMissingPartnerData sets the windowsDeviceBlockedOnMissingPartnerData property value. For Windows, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) SetWindowsDeviceBlockedOnMissingPartnerData(value *bool)() {
    if m != nil {
        m.windowsDeviceBlockedOnMissingPartnerData = value
    }
}
// SetWindowsEnabled sets the windowsEnabled property value. For Windows, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) SetWindowsEnabled(value *bool)() {
    if m != nil {
        m.windowsEnabled = value
    }
}
