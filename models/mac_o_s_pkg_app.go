package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSPkgApp contains properties and inherited properties for the MacOSPkgApp.
type MacOSPkgApp struct {
    MobileLobApp
}
// NewMacOSPkgApp instantiates a new macOSPkgApp and sets the default values.
func NewMacOSPkgApp()(*MacOSPkgApp) {
    m := &MacOSPkgApp{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.macOSPkgApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSPkgAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSPkgAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSPkgApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSPkgApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["ignoreVersionDetection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreVersionDetection(val)
        }
        return nil
    }
    res["includedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSIncludedAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSIncludedAppable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MacOSIncludedAppable)
                }
            }
            m.SetIncludedApps(res)
        }
        return nil
    }
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMacOSMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(MacOSMinimumOperatingSystemable))
        }
        return nil
    }
    res["postInstallScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMacOSAppScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostInstallScript(val.(MacOSAppScriptable))
        }
        return nil
    }
    res["preInstallScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMacOSAppScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreInstallScript(val.(MacOSAppScriptable))
        }
        return nil
    }
    res["primaryBundleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryBundleId(val)
        }
        return nil
    }
    res["primaryBundleVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryBundleVersion(val)
        }
        return nil
    }
    return res
}
// GetIgnoreVersionDetection gets the ignoreVersionDetection property value. When TRUE, indicates that the app's version will NOT be used to detect if the app is installed on a device. When FALSE, indicates that the app's version will be used to detect if the app is installed on a device. Set this to true for apps that use a self update feature. The default value is FALSE.
func (m *MacOSPkgApp) GetIgnoreVersionDetection()(*bool) {
    val, err := m.GetBackingStore().Get("ignoreVersionDetection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIncludedApps gets the includedApps property value. The list of apps expected to be installed by the PKG.
func (m *MacOSPkgApp) GetIncludedApps()([]MacOSIncludedAppable) {
    val, err := m.GetBackingStore().Get("includedApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOSIncludedAppable)
    }
    return nil
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. ComplexType macOSMinimumOperatingSystem that indicates the minimum operating system applicable for the application.
func (m *MacOSPkgApp) GetMinimumSupportedOperatingSystem()(MacOSMinimumOperatingSystemable) {
    val, err := m.GetBackingStore().Get("minimumSupportedOperatingSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MacOSMinimumOperatingSystemable)
    }
    return nil
}
// GetPostInstallScript gets the postInstallScript property value. ComplexType macOSAppScript the contains the post-install script for the app. This will execute on the macOS device after the app is installed.
func (m *MacOSPkgApp) GetPostInstallScript()(MacOSAppScriptable) {
    val, err := m.GetBackingStore().Get("postInstallScript")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MacOSAppScriptable)
    }
    return nil
}
// GetPreInstallScript gets the preInstallScript property value. ComplexType macOSAppScript the contains the post-install script for the app. This will execute on the macOS device after the app is installed.
func (m *MacOSPkgApp) GetPreInstallScript()(MacOSAppScriptable) {
    val, err := m.GetBackingStore().Get("preInstallScript")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MacOSAppScriptable)
    }
    return nil
}
// GetPrimaryBundleId gets the primaryBundleId property value. The bundleId of the primary app in the PKG. This maps to the CFBundleIdentifier in the app's bundle configuration.
func (m *MacOSPkgApp) GetPrimaryBundleId()(*string) {
    val, err := m.GetBackingStore().Get("primaryBundleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrimaryBundleVersion gets the primaryBundleVersion property value. The version of the primary app in the PKG. This maps to the CFBundleShortVersion in the app's bundle configuration.
func (m *MacOSPkgApp) GetPrimaryBundleVersion()(*string) {
    val, err := m.GetBackingStore().Get("primaryBundleVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSPkgApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludedApps()))
        for i, v := range m.GetIncludedApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        err = writer.WriteObjectValue("postInstallScript", m.GetPostInstallScript())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("preInstallScript", m.GetPreInstallScript())
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
// SetIgnoreVersionDetection sets the ignoreVersionDetection property value. When TRUE, indicates that the app's version will NOT be used to detect if the app is installed on a device. When FALSE, indicates that the app's version will be used to detect if the app is installed on a device. Set this to true for apps that use a self update feature. The default value is FALSE.
func (m *MacOSPkgApp) SetIgnoreVersionDetection(value *bool)() {
    err := m.GetBackingStore().Set("ignoreVersionDetection", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludedApps sets the includedApps property value. The list of apps expected to be installed by the PKG.
func (m *MacOSPkgApp) SetIncludedApps(value []MacOSIncludedAppable)() {
    err := m.GetBackingStore().Set("includedApps", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. ComplexType macOSMinimumOperatingSystem that indicates the minimum operating system applicable for the application.
func (m *MacOSPkgApp) SetMinimumSupportedOperatingSystem(value MacOSMinimumOperatingSystemable)() {
    err := m.GetBackingStore().Set("minimumSupportedOperatingSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetPostInstallScript sets the postInstallScript property value. ComplexType macOSAppScript the contains the post-install script for the app. This will execute on the macOS device after the app is installed.
func (m *MacOSPkgApp) SetPostInstallScript(value MacOSAppScriptable)() {
    err := m.GetBackingStore().Set("postInstallScript", value)
    if err != nil {
        panic(err)
    }
}
// SetPreInstallScript sets the preInstallScript property value. ComplexType macOSAppScript the contains the post-install script for the app. This will execute on the macOS device after the app is installed.
func (m *MacOSPkgApp) SetPreInstallScript(value MacOSAppScriptable)() {
    err := m.GetBackingStore().Set("preInstallScript", value)
    if err != nil {
        panic(err)
    }
}
// SetPrimaryBundleId sets the primaryBundleId property value. The bundleId of the primary app in the PKG. This maps to the CFBundleIdentifier in the app's bundle configuration.
func (m *MacOSPkgApp) SetPrimaryBundleId(value *string)() {
    err := m.GetBackingStore().Set("primaryBundleId", value)
    if err != nil {
        panic(err)
    }
}
// SetPrimaryBundleVersion sets the primaryBundleVersion property value. The version of the primary app in the PKG. This maps to the CFBundleShortVersion in the app's bundle configuration.
func (m *MacOSPkgApp) SetPrimaryBundleVersion(value *string)() {
    err := m.GetBackingStore().Set("primaryBundleVersion", value)
    if err != nil {
        panic(err)
    }
}
// MacOSPkgAppable 
type MacOSPkgAppable interface {
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIgnoreVersionDetection()(*bool)
    GetIncludedApps()([]MacOSIncludedAppable)
    GetMinimumSupportedOperatingSystem()(MacOSMinimumOperatingSystemable)
    GetPostInstallScript()(MacOSAppScriptable)
    GetPreInstallScript()(MacOSAppScriptable)
    GetPrimaryBundleId()(*string)
    GetPrimaryBundleVersion()(*string)
    SetIgnoreVersionDetection(value *bool)()
    SetIncludedApps(value []MacOSIncludedAppable)()
    SetMinimumSupportedOperatingSystem(value MacOSMinimumOperatingSystemable)()
    SetPostInstallScript(value MacOSAppScriptable)()
    SetPreInstallScript(value MacOSAppScriptable)()
    SetPrimaryBundleId(value *string)()
    SetPrimaryBundleVersion(value *string)()
}
