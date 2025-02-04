package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentPlatformRestrictionsConfiguration default Device Enrollment Platform Restrictions Configuration that restricts the types of devices a user can enroll
type DeviceEnrollmentPlatformRestrictionsConfiguration struct {
    DeviceEnrollmentConfiguration
}
// NewDeviceEnrollmentPlatformRestrictionsConfiguration instantiates a new DeviceEnrollmentPlatformRestrictionsConfiguration and sets the default values.
func NewDeviceEnrollmentPlatformRestrictionsConfiguration()(*DeviceEnrollmentPlatformRestrictionsConfiguration) {
    m := &DeviceEnrollmentPlatformRestrictionsConfiguration{
        DeviceEnrollmentConfiguration: *NewDeviceEnrollmentConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.deviceEnrollmentPlatformRestrictionsConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceEnrollmentPlatformRestrictionsConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceEnrollmentPlatformRestrictionsConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentPlatformRestrictionsConfiguration(), nil
}
// GetAndroidForWorkRestriction gets the androidForWorkRestriction property value. Android for work restrictions based on platform, platform operating system version, and device ownership
// returns a DeviceEnrollmentPlatformRestrictionable when successful
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetAndroidForWorkRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    val, err := m.GetBackingStore().Get("androidForWorkRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceEnrollmentPlatformRestrictionable)
    }
    return nil
}
// GetAndroidRestriction gets the androidRestriction property value. Android restrictions based on platform, platform operating system version, and device ownership
// returns a DeviceEnrollmentPlatformRestrictionable when successful
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetAndroidRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    val, err := m.GetBackingStore().Get("androidRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceEnrollmentPlatformRestrictionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceEnrollmentConfiguration.GetFieldDeserializers()
    res["androidForWorkRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidForWorkRestriction(val.(DeviceEnrollmentPlatformRestrictionable))
        }
        return nil
    }
    res["androidRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidRestriction(val.(DeviceEnrollmentPlatformRestrictionable))
        }
        return nil
    }
    res["iosRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosRestriction(val.(DeviceEnrollmentPlatformRestrictionable))
        }
        return nil
    }
    res["macOSRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOSRestriction(val.(DeviceEnrollmentPlatformRestrictionable))
        }
        return nil
    }
    res["macRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacRestriction(val.(DeviceEnrollmentPlatformRestrictionable))
        }
        return nil
    }
    res["tvosRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTvosRestriction(val.(DeviceEnrollmentPlatformRestrictionable))
        }
        return nil
    }
    res["visionOSRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisionOSRestriction(val.(DeviceEnrollmentPlatformRestrictionable))
        }
        return nil
    }
    res["windowsHomeSkuRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsHomeSkuRestriction(val.(DeviceEnrollmentPlatformRestrictionable))
        }
        return nil
    }
    res["windowsMobileRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsMobileRestriction(val.(DeviceEnrollmentPlatformRestrictionable))
        }
        return nil
    }
    res["windowsRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsRestriction(val.(DeviceEnrollmentPlatformRestrictionable))
        }
        return nil
    }
    return res
}
// GetIosRestriction gets the iosRestriction property value. Ios restrictions based on platform, platform operating system version, and device ownership
// returns a DeviceEnrollmentPlatformRestrictionable when successful
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetIosRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    val, err := m.GetBackingStore().Get("iosRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceEnrollmentPlatformRestrictionable)
    }
    return nil
}
// GetMacOSRestriction gets the macOSRestriction property value. Mac restrictions based on platform, platform operating system version, and device ownership
// returns a DeviceEnrollmentPlatformRestrictionable when successful
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetMacOSRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    val, err := m.GetBackingStore().Get("macOSRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceEnrollmentPlatformRestrictionable)
    }
    return nil
}
// GetMacRestriction gets the macRestriction property value. Mac restrictions based on platform, platform operating system version, and device ownership
// returns a DeviceEnrollmentPlatformRestrictionable when successful
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetMacRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    val, err := m.GetBackingStore().Get("macRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceEnrollmentPlatformRestrictionable)
    }
    return nil
}
// GetTvosRestriction gets the tvosRestriction property value. Indicates restrictions for TvOS platform.
// returns a DeviceEnrollmentPlatformRestrictionable when successful
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetTvosRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    val, err := m.GetBackingStore().Get("tvosRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceEnrollmentPlatformRestrictionable)
    }
    return nil
}
// GetVisionOSRestriction gets the visionOSRestriction property value. Indicates restrictions for VisionOS platform.
// returns a DeviceEnrollmentPlatformRestrictionable when successful
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetVisionOSRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    val, err := m.GetBackingStore().Get("visionOSRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceEnrollmentPlatformRestrictionable)
    }
    return nil
}
// GetWindowsHomeSkuRestriction gets the windowsHomeSkuRestriction property value. Windows Home Sku restrictions based on platform, platform operating system version, and device ownership
// returns a DeviceEnrollmentPlatformRestrictionable when successful
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetWindowsHomeSkuRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    val, err := m.GetBackingStore().Get("windowsHomeSkuRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceEnrollmentPlatformRestrictionable)
    }
    return nil
}
// GetWindowsMobileRestriction gets the windowsMobileRestriction property value. Windows mobile restrictions based on platform, platform operating system version, and device ownership
// returns a DeviceEnrollmentPlatformRestrictionable when successful
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetWindowsMobileRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    val, err := m.GetBackingStore().Get("windowsMobileRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceEnrollmentPlatformRestrictionable)
    }
    return nil
}
// GetWindowsRestriction gets the windowsRestriction property value. Windows restrictions based on platform, platform operating system version, and device ownership
// returns a DeviceEnrollmentPlatformRestrictionable when successful
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) GetWindowsRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    val, err := m.GetBackingStore().Get("windowsRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceEnrollmentPlatformRestrictionable)
    }
    return nil
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
        err = writer.WriteObjectValue("tvosRestriction", m.GetTvosRestriction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("visionOSRestriction", m.GetVisionOSRestriction())
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
    err := m.GetBackingStore().Set("androidForWorkRestriction", value)
    if err != nil {
        panic(err)
    }
}
// SetAndroidRestriction sets the androidRestriction property value. Android restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetAndroidRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    err := m.GetBackingStore().Set("androidRestriction", value)
    if err != nil {
        panic(err)
    }
}
// SetIosRestriction sets the iosRestriction property value. Ios restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetIosRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    err := m.GetBackingStore().Set("iosRestriction", value)
    if err != nil {
        panic(err)
    }
}
// SetMacOSRestriction sets the macOSRestriction property value. Mac restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetMacOSRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    err := m.GetBackingStore().Set("macOSRestriction", value)
    if err != nil {
        panic(err)
    }
}
// SetMacRestriction sets the macRestriction property value. Mac restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetMacRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    err := m.GetBackingStore().Set("macRestriction", value)
    if err != nil {
        panic(err)
    }
}
// SetTvosRestriction sets the tvosRestriction property value. Indicates restrictions for TvOS platform.
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetTvosRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    err := m.GetBackingStore().Set("tvosRestriction", value)
    if err != nil {
        panic(err)
    }
}
// SetVisionOSRestriction sets the visionOSRestriction property value. Indicates restrictions for VisionOS platform.
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetVisionOSRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    err := m.GetBackingStore().Set("visionOSRestriction", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsHomeSkuRestriction sets the windowsHomeSkuRestriction property value. Windows Home Sku restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetWindowsHomeSkuRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    err := m.GetBackingStore().Set("windowsHomeSkuRestriction", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsMobileRestriction sets the windowsMobileRestriction property value. Windows mobile restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetWindowsMobileRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    err := m.GetBackingStore().Set("windowsMobileRestriction", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsRestriction sets the windowsRestriction property value. Windows restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionsConfiguration) SetWindowsRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    err := m.GetBackingStore().Set("windowsRestriction", value)
    if err != nil {
        panic(err)
    }
}
type DeviceEnrollmentPlatformRestrictionsConfigurationable interface {
    DeviceEnrollmentConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAndroidForWorkRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetAndroidRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetIosRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetMacOSRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetMacRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetTvosRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetVisionOSRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetWindowsHomeSkuRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetWindowsMobileRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetWindowsRestriction()(DeviceEnrollmentPlatformRestrictionable)
    SetAndroidForWorkRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetAndroidRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetIosRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetMacOSRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetMacRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetTvosRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetVisionOSRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetWindowsHomeSkuRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetWindowsMobileRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetWindowsRestriction(value DeviceEnrollmentPlatformRestrictionable)()
}
