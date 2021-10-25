package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MobileThreatDefenseConnector struct {
    Entity
    allowPartnerToCollectIOSApplicationMetadata *bool;
    androidDeviceBlockedOnMissingPartnerData *bool;
    androidEnabled *bool;
    androidMobileApplicationManagementEnabled *bool;
    iosDeviceBlockedOnMissingPartnerData *bool;
    iosEnabled *bool;
    iosMobileApplicationManagementEnabled *bool;
    lastHeartbeatDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    macDeviceBlockedOnMissingPartnerData *bool;
    macEnabled *bool;
    partnerState *MobileThreatPartnerTenantState;
    partnerUnresponsivenessThresholdInDays *int32;
    partnerUnsupportedOsVersionBlocked *bool;
    windowsDeviceBlockedOnMissingPartnerData *bool;
    windowsEnabled *bool;
}
func NewMobileThreatDefenseConnector()(*MobileThreatDefenseConnector) {
    m := &MobileThreatDefenseConnector{
        Entity: *NewEntity(),
    }
    return m
}
func (m *MobileThreatDefenseConnector) GetAllowPartnerToCollectIOSApplicationMetadata()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowPartnerToCollectIOSApplicationMetadata
    }
}
func (m *MobileThreatDefenseConnector) GetAndroidDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceBlockedOnMissingPartnerData
    }
}
func (m *MobileThreatDefenseConnector) GetAndroidEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidEnabled
    }
}
func (m *MobileThreatDefenseConnector) GetAndroidMobileApplicationManagementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidMobileApplicationManagementEnabled
    }
}
func (m *MobileThreatDefenseConnector) GetIosDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosDeviceBlockedOnMissingPartnerData
    }
}
func (m *MobileThreatDefenseConnector) GetIosEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosEnabled
    }
}
func (m *MobileThreatDefenseConnector) GetIosMobileApplicationManagementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iosMobileApplicationManagementEnabled
    }
}
func (m *MobileThreatDefenseConnector) GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastHeartbeatDateTime
    }
}
func (m *MobileThreatDefenseConnector) GetMacDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.macDeviceBlockedOnMissingPartnerData
    }
}
func (m *MobileThreatDefenseConnector) GetMacEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.macEnabled
    }
}
func (m *MobileThreatDefenseConnector) GetPartnerState()(*MobileThreatPartnerTenantState) {
    if m == nil {
        return nil
    } else {
        return m.partnerState
    }
}
func (m *MobileThreatDefenseConnector) GetPartnerUnresponsivenessThresholdInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.partnerUnresponsivenessThresholdInDays
    }
}
func (m *MobileThreatDefenseConnector) GetPartnerUnsupportedOsVersionBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.partnerUnsupportedOsVersionBlocked
    }
}
func (m *MobileThreatDefenseConnector) GetWindowsDeviceBlockedOnMissingPartnerData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.windowsDeviceBlockedOnMissingPartnerData
    }
}
func (m *MobileThreatDefenseConnector) GetWindowsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.windowsEnabled
    }
}
func (m *MobileThreatDefenseConnector) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowPartnerToCollectIOSApplicationMetadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowPartnerToCollectIOSApplicationMetadata(val)
        return nil
    }
    res["androidDeviceBlockedOnMissingPartnerData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAndroidDeviceBlockedOnMissingPartnerData(val)
        return nil
    }
    res["androidEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAndroidEnabled(val)
        return nil
    }
    res["androidMobileApplicationManagementEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAndroidMobileApplicationManagementEnabled(val)
        return nil
    }
    res["iosDeviceBlockedOnMissingPartnerData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIosDeviceBlockedOnMissingPartnerData(val)
        return nil
    }
    res["iosEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIosEnabled(val)
        return nil
    }
    res["iosMobileApplicationManagementEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIosMobileApplicationManagementEnabled(val)
        return nil
    }
    res["lastHeartbeatDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastHeartbeatDateTime(val)
        return nil
    }
    res["macDeviceBlockedOnMissingPartnerData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMacDeviceBlockedOnMissingPartnerData(val)
        return nil
    }
    res["macEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMacEnabled(val)
        return nil
    }
    res["partnerState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileThreatPartnerTenantState)
        if err != nil {
            return err
        }
        cast := val.(MobileThreatPartnerTenantState)
        m.SetPartnerState(&cast)
        return nil
    }
    res["partnerUnresponsivenessThresholdInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPartnerUnresponsivenessThresholdInDays(val)
        return nil
    }
    res["partnerUnsupportedOsVersionBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPartnerUnsupportedOsVersionBlocked(val)
        return nil
    }
    res["windowsDeviceBlockedOnMissingPartnerData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetWindowsDeviceBlockedOnMissingPartnerData(val)
        return nil
    }
    res["windowsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetWindowsEnabled(val)
        return nil
    }
    return res
}
func (m *MobileThreatDefenseConnector) IsNil()(bool) {
    return m == nil
}
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
func (m *MobileThreatDefenseConnector) SetAllowPartnerToCollectIOSApplicationMetadata(value *bool)() {
    m.allowPartnerToCollectIOSApplicationMetadata = value
}
func (m *MobileThreatDefenseConnector) SetAndroidDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.androidDeviceBlockedOnMissingPartnerData = value
}
func (m *MobileThreatDefenseConnector) SetAndroidEnabled(value *bool)() {
    m.androidEnabled = value
}
func (m *MobileThreatDefenseConnector) SetAndroidMobileApplicationManagementEnabled(value *bool)() {
    m.androidMobileApplicationManagementEnabled = value
}
func (m *MobileThreatDefenseConnector) SetIosDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.iosDeviceBlockedOnMissingPartnerData = value
}
func (m *MobileThreatDefenseConnector) SetIosEnabled(value *bool)() {
    m.iosEnabled = value
}
func (m *MobileThreatDefenseConnector) SetIosMobileApplicationManagementEnabled(value *bool)() {
    m.iosMobileApplicationManagementEnabled = value
}
func (m *MobileThreatDefenseConnector) SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastHeartbeatDateTime = value
}
func (m *MobileThreatDefenseConnector) SetMacDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.macDeviceBlockedOnMissingPartnerData = value
}
func (m *MobileThreatDefenseConnector) SetMacEnabled(value *bool)() {
    m.macEnabled = value
}
func (m *MobileThreatDefenseConnector) SetPartnerState(value *MobileThreatPartnerTenantState)() {
    m.partnerState = value
}
func (m *MobileThreatDefenseConnector) SetPartnerUnresponsivenessThresholdInDays(value *int32)() {
    m.partnerUnresponsivenessThresholdInDays = value
}
func (m *MobileThreatDefenseConnector) SetPartnerUnsupportedOsVersionBlocked(value *bool)() {
    m.partnerUnsupportedOsVersionBlocked = value
}
func (m *MobileThreatDefenseConnector) SetWindowsDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.windowsDeviceBlockedOnMissingPartnerData = value
}
func (m *MobileThreatDefenseConnector) SetWindowsEnabled(value *bool)() {
    m.windowsEnabled = value
}
