package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSDmgApp 
type MacOSDmgApp struct {
    MobileLobApp
    // A value indicating whether the app's version will be used to detect the app after it is installed on a device. Set this to true for apps that use a self-update feature. Set this to false to install the app when it is not already installed on the device, or if the deploying app's version number does not match the version that's already installed on the device.
    ignoreVersionDetection *bool
    // The list of apps expected to be installed by the DMG.
    includedApps []MacOSIncludedAppable
    // The value for the minimum applicable operating system.
    minimumSupportedOperatingSystem MacOSMinimumOperatingSystemable
    // The primary CFBundleIdentifier of the DMG.
    primaryBundleId *string
    // The primary CFBundleVersion of the DMG.
    primaryBundleVersion *string
}
// NewMacOSDmgApp instantiates a new MacOSDmgApp and sets the default values.
func NewMacOSDmgApp()(*MacOSDmgApp) {
    m := &MacOSDmgApp{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.macOSDmgApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSDmgAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSDmgAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSDmgApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSDmgApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["ignoreVersionDetection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIgnoreVersionDetection)
    res["includedApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMacOSIncludedAppFromDiscriminatorValue , m.SetIncludedApps)
    res["minimumSupportedOperatingSystem"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMacOSMinimumOperatingSystemFromDiscriminatorValue , m.SetMinimumSupportedOperatingSystem)
    res["primaryBundleId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPrimaryBundleId)
    res["primaryBundleVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPrimaryBundleVersion)
    return res
}
// GetIgnoreVersionDetection gets the ignoreVersionDetection property value. A value indicating whether the app's version will be used to detect the app after it is installed on a device. Set this to true for apps that use a self-update feature. Set this to false to install the app when it is not already installed on the device, or if the deploying app's version number does not match the version that's already installed on the device.
func (m *MacOSDmgApp) GetIgnoreVersionDetection()(*bool) {
    return m.ignoreVersionDetection
}
// GetIncludedApps gets the includedApps property value. The list of apps expected to be installed by the DMG.
func (m *MacOSDmgApp) GetIncludedApps()([]MacOSIncludedAppable) {
    return m.includedApps
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *MacOSDmgApp) GetMinimumSupportedOperatingSystem()(MacOSMinimumOperatingSystemable) {
    return m.minimumSupportedOperatingSystem
}
// GetPrimaryBundleId gets the primaryBundleId property value. The primary CFBundleIdentifier of the DMG.
func (m *MacOSDmgApp) GetPrimaryBundleId()(*string) {
    return m.primaryBundleId
}
// GetPrimaryBundleVersion gets the primaryBundleVersion property value. The primary CFBundleVersion of the DMG.
func (m *MacOSDmgApp) GetPrimaryBundleVersion()(*string) {
    return m.primaryBundleVersion
}
// Serialize serializes information the current object
func (m *MacOSDmgApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("ignoreVersionDetection", m.GetIgnoreVersionDetection())
        if err != nil {
            return err
        }
    }
    if m.GetIncludedApps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetIncludedApps())
        err = writer.WriteCollectionOfObjectValues("includedApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryBundleId", m.GetPrimaryBundleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryBundleVersion", m.GetPrimaryBundleVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIgnoreVersionDetection sets the ignoreVersionDetection property value. A value indicating whether the app's version will be used to detect the app after it is installed on a device. Set this to true for apps that use a self-update feature. Set this to false to install the app when it is not already installed on the device, or if the deploying app's version number does not match the version that's already installed on the device.
func (m *MacOSDmgApp) SetIgnoreVersionDetection(value *bool)() {
    m.ignoreVersionDetection = value
}
// SetIncludedApps sets the includedApps property value. The list of apps expected to be installed by the DMG.
func (m *MacOSDmgApp) SetIncludedApps(value []MacOSIncludedAppable)() {
    m.includedApps = value
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *MacOSDmgApp) SetMinimumSupportedOperatingSystem(value MacOSMinimumOperatingSystemable)() {
    m.minimumSupportedOperatingSystem = value
}
// SetPrimaryBundleId sets the primaryBundleId property value. The primary CFBundleIdentifier of the DMG.
func (m *MacOSDmgApp) SetPrimaryBundleId(value *string)() {
    m.primaryBundleId = value
}
// SetPrimaryBundleVersion sets the primaryBundleVersion property value. The primary CFBundleVersion of the DMG.
func (m *MacOSDmgApp) SetPrimaryBundleVersion(value *string)() {
    m.primaryBundleVersion = value
}
