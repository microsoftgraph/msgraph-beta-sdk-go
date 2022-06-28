package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSLobApp 
type MacOSLobApp struct {
    MobileLobApp
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
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
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacOSLobAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSLobAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSLobApp(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSLobApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBuildNumber gets the buildNumber property value. The build number of MacOS Line of Business (LoB) app.
func (m *MacOSLobApp) GetBuildNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.buildNumber
    }
}
// GetBundleId gets the bundleId property value. The bundle id.
func (m *MacOSLobApp) GetBundleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bundleId
    }
}
// GetChildApps gets the childApps property value. The app list in this bundle package
func (m *MacOSLobApp) GetChildApps()([]MacOSLobChildAppable) {
    if m == nil {
        return nil
    } else {
        return m.childApps
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSLobApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["buildNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuildNumber(val)
        }
        return nil
    }
    res["bundleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleId(val)
        }
        return nil
    }
    res["childApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSLobChildAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSLobChildAppable, len(val))
            for i, v := range val {
                res[i] = v.(MacOSLobChildAppable)
            }
            m.SetChildApps(res)
        }
        return nil
    }
    res["identityVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityVersion(val)
        }
        return nil
    }
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
    res["installAsManaged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallAsManaged(val)
        }
        return nil
    }
    res["md5Hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMd5Hash(res)
        }
        return nil
    }
    res["md5HashChunkSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMd5HashChunkSize(val)
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
    res["versionNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionNumber(val)
        }
        return nil
    }
    return res
}
// GetIdentityVersion gets the identityVersion property value. The identity version.
func (m *MacOSLobApp) GetIdentityVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityVersion
    }
}
// GetIgnoreVersionDetection gets the ignoreVersionDetection property value. A boolean to control whether the app's version will be used to detect the app after it is installed on a device. Set this to true for macOS Line of Business (LoB) apps that use a self update feature.
func (m *MacOSLobApp) GetIgnoreVersionDetection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ignoreVersionDetection
    }
}
// GetInstallAsManaged gets the installAsManaged property value. A boolean to control whether the app will be installed as managed (requires macOS 11.0 and other PKG restrictions).
func (m *MacOSLobApp) GetInstallAsManaged()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.installAsManaged
    }
}
// GetMd5Hash gets the md5Hash property value. The MD5 hash codes
func (m *MacOSLobApp) GetMd5Hash()([]string) {
    if m == nil {
        return nil
    } else {
        return m.md5Hash
    }
}
// GetMd5HashChunkSize gets the md5HashChunkSize property value. The chunk size for MD5 hash
func (m *MacOSLobApp) GetMd5HashChunkSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.md5HashChunkSize
    }
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *MacOSLobApp) GetMinimumSupportedOperatingSystem()(MacOSMinimumOperatingSystemable) {
    if m == nil {
        return nil
    } else {
        return m.minimumSupportedOperatingSystem
    }
}
// GetVersionNumber gets the versionNumber property value. The version number of MacOS Line of Business (LoB) app.
func (m *MacOSLobApp) GetVersionNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.versionNumber
    }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildApps()))
        for i, v := range m.GetChildApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSLobApp) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBuildNumber sets the buildNumber property value. The build number of MacOS Line of Business (LoB) app.
func (m *MacOSLobApp) SetBuildNumber(value *string)() {
    if m != nil {
        m.buildNumber = value
    }
}
// SetBundleId sets the bundleId property value. The bundle id.
func (m *MacOSLobApp) SetBundleId(value *string)() {
    if m != nil {
        m.bundleId = value
    }
}
// SetChildApps sets the childApps property value. The app list in this bundle package
func (m *MacOSLobApp) SetChildApps(value []MacOSLobChildAppable)() {
    if m != nil {
        m.childApps = value
    }
}
// SetIdentityVersion sets the identityVersion property value. The identity version.
func (m *MacOSLobApp) SetIdentityVersion(value *string)() {
    if m != nil {
        m.identityVersion = value
    }
}
// SetIgnoreVersionDetection sets the ignoreVersionDetection property value. A boolean to control whether the app's version will be used to detect the app after it is installed on a device. Set this to true for macOS Line of Business (LoB) apps that use a self update feature.
func (m *MacOSLobApp) SetIgnoreVersionDetection(value *bool)() {
    if m != nil {
        m.ignoreVersionDetection = value
    }
}
// SetInstallAsManaged sets the installAsManaged property value. A boolean to control whether the app will be installed as managed (requires macOS 11.0 and other PKG restrictions).
func (m *MacOSLobApp) SetInstallAsManaged(value *bool)() {
    if m != nil {
        m.installAsManaged = value
    }
}
// SetMd5Hash sets the md5Hash property value. The MD5 hash codes
func (m *MacOSLobApp) SetMd5Hash(value []string)() {
    if m != nil {
        m.md5Hash = value
    }
}
// SetMd5HashChunkSize sets the md5HashChunkSize property value. The chunk size for MD5 hash
func (m *MacOSLobApp) SetMd5HashChunkSize(value *int32)() {
    if m != nil {
        m.md5HashChunkSize = value
    }
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *MacOSLobApp) SetMinimumSupportedOperatingSystem(value MacOSMinimumOperatingSystemable)() {
    if m != nil {
        m.minimumSupportedOperatingSystem = value
    }
}
// SetVersionNumber sets the versionNumber property value. The version number of MacOS Line of Business (LoB) app.
func (m *MacOSLobApp) SetVersionNumber(value *string)() {
    if m != nil {
        m.versionNumber = value
    }
}
