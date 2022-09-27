package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSLobApp 
type MacOSLobApp struct {
    MobileLobApp
    // The build number of MacOS Line of Business (LoB) app.
    buildNumber *string
    // The bundle id.
    bundleId *string
    // The app list in this bundle package
    childApps []MacOSLobChildAppable
    // The identity version.
    identityVersion *string
    // A boolean to control whether the app's version will be used to detect the app after it is installed on a device. Set this to true for macOS Line of Business (LoB) apps that use a self update feature.
    ignoreVersionDetection *bool
    // A boolean to control whether the app will be installed as managed (requires macOS 11.0 and other PKG restrictions).
    installAsManaged *bool
    // The MD5 hash codes
    md5Hash []string
    // The chunk size for MD5 hash
    md5HashChunkSize *int32
    // The value for the minimum applicable operating system.
    minimumSupportedOperatingSystem MacOSMinimumOperatingSystemable
    // The version number of MacOS Line of Business (LoB) app.
    versionNumber *string
}
// NewMacOSLobApp instantiates a new MacOSLobApp and sets the default values.
func NewMacOSLobApp()(*MacOSLobApp) {
    m := &MacOSLobApp{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.macOSLobApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSLobAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSLobAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSLobApp(), nil
}
// GetBuildNumber gets the buildNumber property value. The build number of MacOS Line of Business (LoB) app.
func (m *MacOSLobApp) GetBuildNumber()(*string) {
    return m.buildNumber
}
// GetBundleId gets the bundleId property value. The bundle id.
func (m *MacOSLobApp) GetBundleId()(*string) {
    return m.bundleId
}
// GetChildApps gets the childApps property value. The app list in this bundle package
func (m *MacOSLobApp) GetChildApps()([]MacOSLobChildAppable) {
    return m.childApps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSLobApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["buildNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBuildNumber)
    res["bundleId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBundleId)
    res["childApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMacOSLobChildAppFromDiscriminatorValue , m.SetChildApps)
    res["identityVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIdentityVersion)
    res["ignoreVersionDetection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIgnoreVersionDetection)
    res["installAsManaged"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetInstallAsManaged)
    res["md5Hash"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetMd5Hash)
    res["md5HashChunkSize"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMd5HashChunkSize)
    res["minimumSupportedOperatingSystem"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMacOSMinimumOperatingSystemFromDiscriminatorValue , m.SetMinimumSupportedOperatingSystem)
    res["versionNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVersionNumber)
    return res
}
// GetIdentityVersion gets the identityVersion property value. The identity version.
func (m *MacOSLobApp) GetIdentityVersion()(*string) {
    return m.identityVersion
}
// GetIgnoreVersionDetection gets the ignoreVersionDetection property value. A boolean to control whether the app's version will be used to detect the app after it is installed on a device. Set this to true for macOS Line of Business (LoB) apps that use a self update feature.
func (m *MacOSLobApp) GetIgnoreVersionDetection()(*bool) {
    return m.ignoreVersionDetection
}
// GetInstallAsManaged gets the installAsManaged property value. A boolean to control whether the app will be installed as managed (requires macOS 11.0 and other PKG restrictions).
func (m *MacOSLobApp) GetInstallAsManaged()(*bool) {
    return m.installAsManaged
}
// GetMd5Hash gets the md5Hash property value. The MD5 hash codes
func (m *MacOSLobApp) GetMd5Hash()([]string) {
    return m.md5Hash
}
// GetMd5HashChunkSize gets the md5HashChunkSize property value. The chunk size for MD5 hash
func (m *MacOSLobApp) GetMd5HashChunkSize()(*int32) {
    return m.md5HashChunkSize
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *MacOSLobApp) GetMinimumSupportedOperatingSystem()(MacOSMinimumOperatingSystemable) {
    return m.minimumSupportedOperatingSystem
}
// GetVersionNumber gets the versionNumber property value. The version number of MacOS Line of Business (LoB) app.
func (m *MacOSLobApp) GetVersionNumber()(*string) {
    return m.versionNumber
}
// Serialize serializes information the current object
func (m *MacOSLobApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("buildNumber", m.GetBuildNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bundleId", m.GetBundleId())
        if err != nil {
            return err
        }
    }
    if m.GetChildApps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetChildApps())
        err = writer.WriteCollectionOfObjectValues("childApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityVersion", m.GetIdentityVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("ignoreVersionDetection", m.GetIgnoreVersionDetection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("installAsManaged", m.GetInstallAsManaged())
        if err != nil {
            return err
        }
    }
    if m.GetMd5Hash() != nil {
        err = writer.WriteCollectionOfStringValues("md5Hash", m.GetMd5Hash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("md5HashChunkSize", m.GetMd5HashChunkSize())
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
        err = writer.WriteStringValue("versionNumber", m.GetVersionNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBuildNumber sets the buildNumber property value. The build number of MacOS Line of Business (LoB) app.
func (m *MacOSLobApp) SetBuildNumber(value *string)() {
    m.buildNumber = value
}
// SetBundleId sets the bundleId property value. The bundle id.
func (m *MacOSLobApp) SetBundleId(value *string)() {
    m.bundleId = value
}
// SetChildApps sets the childApps property value. The app list in this bundle package
func (m *MacOSLobApp) SetChildApps(value []MacOSLobChildAppable)() {
    m.childApps = value
}
// SetIdentityVersion sets the identityVersion property value. The identity version.
func (m *MacOSLobApp) SetIdentityVersion(value *string)() {
    m.identityVersion = value
}
// SetIgnoreVersionDetection sets the ignoreVersionDetection property value. A boolean to control whether the app's version will be used to detect the app after it is installed on a device. Set this to true for macOS Line of Business (LoB) apps that use a self update feature.
func (m *MacOSLobApp) SetIgnoreVersionDetection(value *bool)() {
    m.ignoreVersionDetection = value
}
// SetInstallAsManaged sets the installAsManaged property value. A boolean to control whether the app will be installed as managed (requires macOS 11.0 and other PKG restrictions).
func (m *MacOSLobApp) SetInstallAsManaged(value *bool)() {
    m.installAsManaged = value
}
// SetMd5Hash sets the md5Hash property value. The MD5 hash codes
func (m *MacOSLobApp) SetMd5Hash(value []string)() {
    m.md5Hash = value
}
// SetMd5HashChunkSize sets the md5HashChunkSize property value. The chunk size for MD5 hash
func (m *MacOSLobApp) SetMd5HashChunkSize(value *int32)() {
    m.md5HashChunkSize = value
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *MacOSLobApp) SetMinimumSupportedOperatingSystem(value MacOSMinimumOperatingSystemable)() {
    m.minimumSupportedOperatingSystem = value
}
// SetVersionNumber sets the versionNumber property value. The version number of MacOS Line of Business (LoB) app.
func (m *MacOSLobApp) SetVersionNumber(value *string)() {
    m.versionNumber = value
}
