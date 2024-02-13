package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDefenderAdvancedThreatProtectionConfiguration windows Defender AdvancedThreatProtection Configuration.
type WindowsDefenderAdvancedThreatProtectionConfiguration struct {
    DeviceConfiguration
}
// NewWindowsDefenderAdvancedThreatProtectionConfiguration instantiates a new WindowsDefenderAdvancedThreatProtectionConfiguration and sets the default values.
func NewWindowsDefenderAdvancedThreatProtectionConfiguration()(*WindowsDefenderAdvancedThreatProtectionConfiguration) {
    m := &WindowsDefenderAdvancedThreatProtectionConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsDefenderAdvancedThreatProtectionConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsDefenderAdvancedThreatProtectionConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsDefenderAdvancedThreatProtectionConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDefenderAdvancedThreatProtectionConfiguration(), nil
}
// GetAdvancedThreatProtectionAutoPopulateOnboardingBlob gets the advancedThreatProtectionAutoPopulateOnboardingBlob property value. Auto populate onboarding blob programmatically from Advanced Threat protection service
// returns a *bool when successful
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetAdvancedThreatProtectionAutoPopulateOnboardingBlob()(*bool) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionAutoPopulateOnboardingBlob")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAdvancedThreatProtectionOffboardingBlob gets the advancedThreatProtectionOffboardingBlob property value. Windows Defender AdvancedThreatProtection Offboarding Blob.
// returns a *string when successful
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetAdvancedThreatProtectionOffboardingBlob()(*string) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionOffboardingBlob")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdvancedThreatProtectionOffboardingFilename gets the advancedThreatProtectionOffboardingFilename property value. Name of the file from which AdvancedThreatProtectionOffboardingBlob was obtained.
// returns a *string when successful
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetAdvancedThreatProtectionOffboardingFilename()(*string) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionOffboardingFilename")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdvancedThreatProtectionOnboardingBlob gets the advancedThreatProtectionOnboardingBlob property value. Windows Defender AdvancedThreatProtection Onboarding Blob.
// returns a *string when successful
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetAdvancedThreatProtectionOnboardingBlob()(*string) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionOnboardingBlob")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdvancedThreatProtectionOnboardingFilename gets the advancedThreatProtectionOnboardingFilename property value. Name of the file from which AdvancedThreatProtectionOnboardingBlob was obtained.
// returns a *string when successful
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetAdvancedThreatProtectionOnboardingFilename()(*string) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionOnboardingFilename")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAllowSampleSharing gets the allowSampleSharing property value. Windows Defender AdvancedThreatProtection 'Allow Sample Sharing' Rule
// returns a *bool when successful
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetAllowSampleSharing()(*bool) {
    val, err := m.GetBackingStore().Get("allowSampleSharing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableExpeditedTelemetryReporting gets the enableExpeditedTelemetryReporting property value. Expedite Windows Defender Advanced Threat Protection telemetry reporting frequency.
// returns a *bool when successful
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetEnableExpeditedTelemetryReporting()(*bool) {
    val, err := m.GetBackingStore().Get("enableExpeditedTelemetryReporting")
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
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["advancedThreatProtectionAutoPopulateOnboardingBlob"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionAutoPopulateOnboardingBlob(val)
        }
        return nil
    }
    res["advancedThreatProtectionOffboardingBlob"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionOffboardingBlob(val)
        }
        return nil
    }
    res["advancedThreatProtectionOffboardingFilename"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionOffboardingFilename(val)
        }
        return nil
    }
    res["advancedThreatProtectionOnboardingBlob"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionOnboardingBlob(val)
        }
        return nil
    }
    res["advancedThreatProtectionOnboardingFilename"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionOnboardingFilename(val)
        }
        return nil
    }
    res["allowSampleSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowSampleSharing(val)
        }
        return nil
    }
    res["enableExpeditedTelemetryReporting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableExpeditedTelemetryReporting(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("advancedThreatProtectionAutoPopulateOnboardingBlob", m.GetAdvancedThreatProtectionAutoPopulateOnboardingBlob())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("advancedThreatProtectionOffboardingBlob", m.GetAdvancedThreatProtectionOffboardingBlob())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("advancedThreatProtectionOffboardingFilename", m.GetAdvancedThreatProtectionOffboardingFilename())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("advancedThreatProtectionOnboardingBlob", m.GetAdvancedThreatProtectionOnboardingBlob())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("advancedThreatProtectionOnboardingFilename", m.GetAdvancedThreatProtectionOnboardingFilename())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowSampleSharing", m.GetAllowSampleSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableExpeditedTelemetryReporting", m.GetEnableExpeditedTelemetryReporting())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdvancedThreatProtectionAutoPopulateOnboardingBlob sets the advancedThreatProtectionAutoPopulateOnboardingBlob property value. Auto populate onboarding blob programmatically from Advanced Threat protection service
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetAdvancedThreatProtectionAutoPopulateOnboardingBlob(value *bool)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionAutoPopulateOnboardingBlob", value)
    if err != nil {
        panic(err)
    }
}
// SetAdvancedThreatProtectionOffboardingBlob sets the advancedThreatProtectionOffboardingBlob property value. Windows Defender AdvancedThreatProtection Offboarding Blob.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetAdvancedThreatProtectionOffboardingBlob(value *string)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionOffboardingBlob", value)
    if err != nil {
        panic(err)
    }
}
// SetAdvancedThreatProtectionOffboardingFilename sets the advancedThreatProtectionOffboardingFilename property value. Name of the file from which AdvancedThreatProtectionOffboardingBlob was obtained.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetAdvancedThreatProtectionOffboardingFilename(value *string)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionOffboardingFilename", value)
    if err != nil {
        panic(err)
    }
}
// SetAdvancedThreatProtectionOnboardingBlob sets the advancedThreatProtectionOnboardingBlob property value. Windows Defender AdvancedThreatProtection Onboarding Blob.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetAdvancedThreatProtectionOnboardingBlob(value *string)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionOnboardingBlob", value)
    if err != nil {
        panic(err)
    }
}
// SetAdvancedThreatProtectionOnboardingFilename sets the advancedThreatProtectionOnboardingFilename property value. Name of the file from which AdvancedThreatProtectionOnboardingBlob was obtained.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetAdvancedThreatProtectionOnboardingFilename(value *string)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionOnboardingFilename", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowSampleSharing sets the allowSampleSharing property value. Windows Defender AdvancedThreatProtection 'Allow Sample Sharing' Rule
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetAllowSampleSharing(value *bool)() {
    err := m.GetBackingStore().Set("allowSampleSharing", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableExpeditedTelemetryReporting sets the enableExpeditedTelemetryReporting property value. Expedite Windows Defender Advanced Threat Protection telemetry reporting frequency.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetEnableExpeditedTelemetryReporting(value *bool)() {
    err := m.GetBackingStore().Set("enableExpeditedTelemetryReporting", value)
    if err != nil {
        panic(err)
    }
}
type WindowsDefenderAdvancedThreatProtectionConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedThreatProtectionAutoPopulateOnboardingBlob()(*bool)
    GetAdvancedThreatProtectionOffboardingBlob()(*string)
    GetAdvancedThreatProtectionOffboardingFilename()(*string)
    GetAdvancedThreatProtectionOnboardingBlob()(*string)
    GetAdvancedThreatProtectionOnboardingFilename()(*string)
    GetAllowSampleSharing()(*bool)
    GetEnableExpeditedTelemetryReporting()(*bool)
    SetAdvancedThreatProtectionAutoPopulateOnboardingBlob(value *bool)()
    SetAdvancedThreatProtectionOffboardingBlob(value *string)()
    SetAdvancedThreatProtectionOffboardingFilename(value *string)()
    SetAdvancedThreatProtectionOnboardingBlob(value *string)()
    SetAdvancedThreatProtectionOnboardingFilename(value *string)()
    SetAllowSampleSharing(value *bool)()
    SetEnableExpeditedTelemetryReporting(value *bool)()
}
