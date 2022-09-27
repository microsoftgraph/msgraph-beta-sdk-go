package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentPlatformRestrictionsConfiguration 
type DeviceEnrollmentPlatformRestrictionsConfiguration struct {
    DeviceEnrollmentConfiguration
    // Android for work restrictions based on platform, platform operating system version, and device ownership
    androidForWorkRestriction DeviceEnrollmentPlatformRestrictionable
    // Android restrictions based on platform, platform operating system version, and device ownership
    androidRestriction DeviceEnrollmentPlatformRestrictionable
    // Ios restrictions based on platform, platform operating system version, and device ownership
    iosRestriction DeviceEnrollmentPlatformRestrictionable
    // Mac restrictions based on platform, platform operating system version, and device ownership
    macOSRestriction DeviceEnrollmentPlatformRestrictionable
    // Mac restrictions based on platform, platform operating system version, and device ownership
    macRestriction DeviceEnrollmentPlatformRestrictionable
    // Windows Home Sku restrictions based on platform, platform operating system version, and device ownership
    windowsHomeSkuRestriction DeviceEnrollmentPlatformRestrictionable
    // Windows mobile restrictions based on platform, platform operating system version, and device ownership
    windowsMobileRestriction DeviceEnrollmentPlatformRestrictionable
    // Windows restrictions based on platform, platform operating system version, and device ownership
    windowsRestriction DeviceEnrollmentPlatformRestrictionable
}
// NewDeviceEnrollmentPlatformRestrictionsConfiguration instantiates a new DeviceEnrollmentPlatformRestrictionsConfiguration and sets the default values.
func NewDeviceEnrollmentPlatformRestrictionsConfiguration()(*DeviceEnrollmentPlatformRestrictionsConfiguration) {
    m := &DeviceEnrollmentPlatformRestrictionsConfiguration{
        DeviceEnrollmentConfiguration: *NewDeviceEnrollmentConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.deviceEnrollmentPlatformRestrictionsConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceEnrollmentPlatformRestrictionsConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentPlatformRestrictionsConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentPlatformRestrictionsConfiguration(), nil
}
// GetAndroidForWorkRestriction gets the androidForWorkRestriction property value. Android for work restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetAndroidForWorkRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    return m.androidForWorkRestriction
}
// GetAndroidRestriction gets the androidRestriction property value. Android restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetAndroidRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    return m.androidRestriction
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceEnrollmentConfiguration.GetFieldDeserializers()
    res["androidForWorkRestriction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue , m.SetAndroidForWorkRestriction)
    res["androidRestriction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue , m.SetAndroidRestriction)
    res["iosRestriction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue , m.SetIosRestriction)
    res["macOSRestriction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue , m.SetMacOSRestriction)
    res["macRestriction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue , m.SetMacRestriction)
    res["windowsHomeSkuRestriction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue , m.SetWindowsHomeSkuRestriction)
    res["windowsMobileRestriction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue , m.SetWindowsMobileRestriction)
    res["windowsRestriction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue , m.SetWindowsRestriction)
    return res
}
// GetIosRestriction gets the iosRestriction property value. Ios restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetIosRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    return m.iosRestriction
}
// GetMacOSRestriction gets the macOSRestriction property value. Mac restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetMacOSRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    return m.macOSRestriction
}
// GetMacRestriction gets the macRestriction property value. Mac restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetMacRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    return m.macRestriction
}
// GetWindowsHomeSkuRestriction gets the windowsHomeSkuRestriction property value. Windows Home Sku restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetWindowsHomeSkuRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    return m.windowsHomeSkuRestriction
}
// GetWindowsMobileRestriction gets the windowsMobileRestriction property value. Windows mobile restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetWindowsMobileRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    return m.windowsMobileRestriction
}
// GetWindowsRestriction gets the windowsRestriction property value. Windows restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetWindowsRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    return m.windowsRestriction
}
// Serialize serializes information the current object
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceEnrollmentConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("androidForWorkRestriction", m.GetAndroidForWorkRestriction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("androidRestriction", m.GetAndroidRestriction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("iosRestriction", m.GetIosRestriction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("macOSRestriction", m.GetMacOSRestriction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("macRestriction", m.GetMacRestriction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsHomeSkuRestriction", m.GetWindowsHomeSkuRestriction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsMobileRestriction", m.GetWindowsMobileRestriction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsRestriction", m.GetWindowsRestriction())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAndroidForWorkRestriction sets the androidForWorkRestriction property value. Android for work restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetAndroidForWorkRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    m.androidForWorkRestriction = value
}
// SetAndroidRestriction sets the androidRestriction property value. Android restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetAndroidRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    m.androidRestriction = value
}
// SetIosRestriction sets the iosRestriction property value. Ios restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetIosRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    m.iosRestriction = value
}
// SetMacOSRestriction sets the macOSRestriction property value. Mac restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetMacOSRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    m.macOSRestriction = value
}
// SetMacRestriction sets the macRestriction property value. Mac restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetMacRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    m.macRestriction = value
}
// SetWindowsHomeSkuRestriction sets the windowsHomeSkuRestriction property value. Windows Home Sku restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetWindowsHomeSkuRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    m.windowsHomeSkuRestriction = value
}
// SetWindowsMobileRestriction sets the windowsMobileRestriction property value. Windows mobile restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetWindowsMobileRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    m.windowsMobileRestriction = value
}
// SetWindowsRestriction sets the windowsRestriction property value. Windows restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetWindowsRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    m.windowsRestriction = value
}
