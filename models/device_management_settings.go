package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DeviceManagementSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceManagementSettings instantiates a new DeviceManagementSettings and sets the default values.
func NewDeviceManagementSettings()(*DeviceManagementSettings) {
    m := &DeviceManagementSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceManagementSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeviceManagementSettings) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAndroidDeviceAdministratorEnrollmentEnabled gets the androidDeviceAdministratorEnrollmentEnabled property value. The property to determine if Android device administrator enrollment is enabled for this account.
// returns a *bool when successful
func (m *DeviceManagementSettings) GetAndroidDeviceAdministratorEnrollmentEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("androidDeviceAdministratorEnrollmentEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *DeviceManagementSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDerivedCredentialProvider gets the derivedCredentialProvider property value. Provider type for Derived Credentials.
// returns a *DerivedCredentialProviderType when successful
func (m *DeviceManagementSettings) GetDerivedCredentialProvider()(*DerivedCredentialProviderType) {
    val, err := m.GetBackingStore().Get("derivedCredentialProvider")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DerivedCredentialProviderType)
    }
    return nil
}
// GetDerivedCredentialUrl gets the derivedCredentialUrl property value. The Derived Credential Provider self-service URI.
// returns a *string when successful
func (m *DeviceManagementSettings) GetDerivedCredentialUrl()(*string) {
    val, err := m.GetBackingStore().Get("derivedCredentialUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceComplianceCheckinThresholdDays gets the deviceComplianceCheckinThresholdDays property value. The number of days a device is allowed to go without checking in to remain compliant.
// returns a *int32 when successful
func (m *DeviceManagementSettings) GetDeviceComplianceCheckinThresholdDays()(*int32) {
    val, err := m.GetBackingStore().Get("deviceComplianceCheckinThresholdDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDeviceInactivityBeforeRetirementInDay gets the deviceInactivityBeforeRetirementInDay property value. When the device does not check in for specified number of days, the company data might be removed and the device will not be under management. Valid values 30 to 270
// returns a *int32 when successful
func (m *DeviceManagementSettings) GetDeviceInactivityBeforeRetirementInDay()(*int32) {
    val, err := m.GetBackingStore().Get("deviceInactivityBeforeRetirementInDay")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEnableAutopilotDiagnostics gets the enableAutopilotDiagnostics property value. Determines whether the autopilot diagnostic feature is enabled or not.
// returns a *bool when successful
func (m *DeviceManagementSettings) GetEnableAutopilotDiagnostics()(*bool) {
    val, err := m.GetBackingStore().Get("enableAutopilotDiagnostics")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableDeviceGroupMembershipReport gets the enableDeviceGroupMembershipReport property value. Determines whether the device group membership report feature is enabled or not.
// returns a *bool when successful
func (m *DeviceManagementSettings) GetEnableDeviceGroupMembershipReport()(*bool) {
    val, err := m.GetBackingStore().Get("enableDeviceGroupMembershipReport")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableEnhancedTroubleshootingExperience gets the enableEnhancedTroubleshootingExperience property value. Determines whether the enhanced troubleshooting UX is enabled or not.
// returns a *bool when successful
func (m *DeviceManagementSettings) GetEnableEnhancedTroubleshootingExperience()(*bool) {
    val, err := m.GetBackingStore().Get("enableEnhancedTroubleshootingExperience")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableLogCollection gets the enableLogCollection property value. Determines whether the log collection feature should be available for use.
// returns a *bool when successful
func (m *DeviceManagementSettings) GetEnableLogCollection()(*bool) {
    val, err := m.GetBackingStore().Get("enableLogCollection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnhancedJailBreak gets the enhancedJailBreak property value. Is feature enabled or not for enhanced jailbreak detection.
// returns a *bool when successful
func (m *DeviceManagementSettings) GetEnhancedJailBreak()(*bool) {
    val, err := m.GetBackingStore().Get("enhancedJailBreak")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagementSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["androidDeviceAdministratorEnrollmentEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDeviceAdministratorEnrollmentEnabled(val)
        }
        return nil
    }
    res["derivedCredentialProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDerivedCredentialProviderType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDerivedCredentialProvider(val.(*DerivedCredentialProviderType))
        }
        return nil
    }
    res["derivedCredentialUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDerivedCredentialUrl(val)
        }
        return nil
    }
    res["deviceComplianceCheckinThresholdDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceCheckinThresholdDays(val)
        }
        return nil
    }
    res["deviceInactivityBeforeRetirementInDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceInactivityBeforeRetirementInDay(val)
        }
        return nil
    }
    res["enableAutopilotDiagnostics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableAutopilotDiagnostics(val)
        }
        return nil
    }
    res["enableDeviceGroupMembershipReport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableDeviceGroupMembershipReport(val)
        }
        return nil
    }
    res["enableEnhancedTroubleshootingExperience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableEnhancedTroubleshootingExperience(val)
        }
        return nil
    }
    res["enableLogCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableLogCollection(val)
        }
        return nil
    }
    res["enhancedJailBreak"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnhancedJailBreak(val)
        }
        return nil
    }
    res["ignoreDevicesForUnsupportedSettingsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreDevicesForUnsupportedSettingsEnabled(val)
        }
        return nil
    }
    res["isScheduledActionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsScheduledActionEnabled(val)
        }
        return nil
    }
    res["m365AppDiagnosticsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetM365AppDiagnosticsEnabled(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["secureByDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureByDefault(val)
        }
        return nil
    }
    return res
}
// GetIgnoreDevicesForUnsupportedSettingsEnabled gets the ignoreDevicesForUnsupportedSettingsEnabled property value. The property to determine whether to ignore unsupported compliance settings on certian models of devices.
// returns a *bool when successful
func (m *DeviceManagementSettings) GetIgnoreDevicesForUnsupportedSettingsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("ignoreDevicesForUnsupportedSettingsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsScheduledActionEnabled gets the isScheduledActionEnabled property value. Is feature enabled or not for scheduled action for rule.
// returns a *bool when successful
func (m *DeviceManagementSettings) GetIsScheduledActionEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isScheduledActionEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetM365AppDiagnosticsEnabled gets the m365AppDiagnosticsEnabled property value. The property to determine if M365 App log collection is enabled for account. When TRUE it indicates that M365 app log collection is enabled for account.  When FALSE it indicates that M365 app log collection is disabled for account. Default value is FALSE
// returns a *bool when successful
func (m *DeviceManagementSettings) GetM365AppDiagnosticsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("m365AppDiagnosticsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DeviceManagementSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSecureByDefault gets the secureByDefault property value. Device should be noncompliant when there is no compliance policy targeted when this is true
// returns a *bool when successful
func (m *DeviceManagementSettings) GetSecureByDefault()(*bool) {
    val, err := m.GetBackingStore().Get("secureByDefault")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("androidDeviceAdministratorEnrollmentEnabled", m.GetAndroidDeviceAdministratorEnrollmentEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetDerivedCredentialProvider() != nil {
        cast := (*m.GetDerivedCredentialProvider()).String()
        err := writer.WriteStringValue("derivedCredentialProvider", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("derivedCredentialUrl", m.GetDerivedCredentialUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("deviceComplianceCheckinThresholdDays", m.GetDeviceComplianceCheckinThresholdDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("deviceInactivityBeforeRetirementInDay", m.GetDeviceInactivityBeforeRetirementInDay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableAutopilotDiagnostics", m.GetEnableAutopilotDiagnostics())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableDeviceGroupMembershipReport", m.GetEnableDeviceGroupMembershipReport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableEnhancedTroubleshootingExperience", m.GetEnableEnhancedTroubleshootingExperience())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableLogCollection", m.GetEnableLogCollection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enhancedJailBreak", m.GetEnhancedJailBreak())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ignoreDevicesForUnsupportedSettingsEnabled", m.GetIgnoreDevicesForUnsupportedSettingsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isScheduledActionEnabled", m.GetIsScheduledActionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("m365AppDiagnosticsEnabled", m.GetM365AppDiagnosticsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("secureByDefault", m.GetSecureByDefault())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAndroidDeviceAdministratorEnrollmentEnabled sets the androidDeviceAdministratorEnrollmentEnabled property value. The property to determine if Android device administrator enrollment is enabled for this account.
func (m *DeviceManagementSettings) SetAndroidDeviceAdministratorEnrollmentEnabled(value *bool)() {
    err := m.GetBackingStore().Set("androidDeviceAdministratorEnrollmentEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeviceManagementSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDerivedCredentialProvider sets the derivedCredentialProvider property value. Provider type for Derived Credentials.
func (m *DeviceManagementSettings) SetDerivedCredentialProvider(value *DerivedCredentialProviderType)() {
    err := m.GetBackingStore().Set("derivedCredentialProvider", value)
    if err != nil {
        panic(err)
    }
}
// SetDerivedCredentialUrl sets the derivedCredentialUrl property value. The Derived Credential Provider self-service URI.
func (m *DeviceManagementSettings) SetDerivedCredentialUrl(value *string)() {
    err := m.GetBackingStore().Set("derivedCredentialUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceComplianceCheckinThresholdDays sets the deviceComplianceCheckinThresholdDays property value. The number of days a device is allowed to go without checking in to remain compliant.
func (m *DeviceManagementSettings) SetDeviceComplianceCheckinThresholdDays(value *int32)() {
    err := m.GetBackingStore().Set("deviceComplianceCheckinThresholdDays", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceInactivityBeforeRetirementInDay sets the deviceInactivityBeforeRetirementInDay property value. When the device does not check in for specified number of days, the company data might be removed and the device will not be under management. Valid values 30 to 270
func (m *DeviceManagementSettings) SetDeviceInactivityBeforeRetirementInDay(value *int32)() {
    err := m.GetBackingStore().Set("deviceInactivityBeforeRetirementInDay", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableAutopilotDiagnostics sets the enableAutopilotDiagnostics property value. Determines whether the autopilot diagnostic feature is enabled or not.
func (m *DeviceManagementSettings) SetEnableAutopilotDiagnostics(value *bool)() {
    err := m.GetBackingStore().Set("enableAutopilotDiagnostics", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableDeviceGroupMembershipReport sets the enableDeviceGroupMembershipReport property value. Determines whether the device group membership report feature is enabled or not.
func (m *DeviceManagementSettings) SetEnableDeviceGroupMembershipReport(value *bool)() {
    err := m.GetBackingStore().Set("enableDeviceGroupMembershipReport", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableEnhancedTroubleshootingExperience sets the enableEnhancedTroubleshootingExperience property value. Determines whether the enhanced troubleshooting UX is enabled or not.
func (m *DeviceManagementSettings) SetEnableEnhancedTroubleshootingExperience(value *bool)() {
    err := m.GetBackingStore().Set("enableEnhancedTroubleshootingExperience", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableLogCollection sets the enableLogCollection property value. Determines whether the log collection feature should be available for use.
func (m *DeviceManagementSettings) SetEnableLogCollection(value *bool)() {
    err := m.GetBackingStore().Set("enableLogCollection", value)
    if err != nil {
        panic(err)
    }
}
// SetEnhancedJailBreak sets the enhancedJailBreak property value. Is feature enabled or not for enhanced jailbreak detection.
func (m *DeviceManagementSettings) SetEnhancedJailBreak(value *bool)() {
    err := m.GetBackingStore().Set("enhancedJailBreak", value)
    if err != nil {
        panic(err)
    }
}
// SetIgnoreDevicesForUnsupportedSettingsEnabled sets the ignoreDevicesForUnsupportedSettingsEnabled property value. The property to determine whether to ignore unsupported compliance settings on certian models of devices.
func (m *DeviceManagementSettings) SetIgnoreDevicesForUnsupportedSettingsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("ignoreDevicesForUnsupportedSettingsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsScheduledActionEnabled sets the isScheduledActionEnabled property value. Is feature enabled or not for scheduled action for rule.
func (m *DeviceManagementSettings) SetIsScheduledActionEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isScheduledActionEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetM365AppDiagnosticsEnabled sets the m365AppDiagnosticsEnabled property value. The property to determine if M365 App log collection is enabled for account. When TRUE it indicates that M365 app log collection is enabled for account.  When FALSE it indicates that M365 app log collection is disabled for account. Default value is FALSE
func (m *DeviceManagementSettings) SetM365AppDiagnosticsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("m365AppDiagnosticsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSecureByDefault sets the secureByDefault property value. Device should be noncompliant when there is no compliance policy targeted when this is true
func (m *DeviceManagementSettings) SetSecureByDefault(value *bool)() {
    err := m.GetBackingStore().Set("secureByDefault", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAndroidDeviceAdministratorEnrollmentEnabled()(*bool)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDerivedCredentialProvider()(*DerivedCredentialProviderType)
    GetDerivedCredentialUrl()(*string)
    GetDeviceComplianceCheckinThresholdDays()(*int32)
    GetDeviceInactivityBeforeRetirementInDay()(*int32)
    GetEnableAutopilotDiagnostics()(*bool)
    GetEnableDeviceGroupMembershipReport()(*bool)
    GetEnableEnhancedTroubleshootingExperience()(*bool)
    GetEnableLogCollection()(*bool)
    GetEnhancedJailBreak()(*bool)
    GetIgnoreDevicesForUnsupportedSettingsEnabled()(*bool)
    GetIsScheduledActionEnabled()(*bool)
    GetM365AppDiagnosticsEnabled()(*bool)
    GetOdataType()(*string)
    GetSecureByDefault()(*bool)
    SetAndroidDeviceAdministratorEnrollmentEnabled(value *bool)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDerivedCredentialProvider(value *DerivedCredentialProviderType)()
    SetDerivedCredentialUrl(value *string)()
    SetDeviceComplianceCheckinThresholdDays(value *int32)()
    SetDeviceInactivityBeforeRetirementInDay(value *int32)()
    SetEnableAutopilotDiagnostics(value *bool)()
    SetEnableDeviceGroupMembershipReport(value *bool)()
    SetEnableEnhancedTroubleshootingExperience(value *bool)()
    SetEnableLogCollection(value *bool)()
    SetEnhancedJailBreak(value *bool)()
    SetIgnoreDevicesForUnsupportedSettingsEnabled(value *bool)()
    SetIsScheduledActionEnabled(value *bool)()
    SetM365AppDiagnosticsEnabled(value *bool)()
    SetOdataType(value *string)()
    SetSecureByDefault(value *bool)()
}
