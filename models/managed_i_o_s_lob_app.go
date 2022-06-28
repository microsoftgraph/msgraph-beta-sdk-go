package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedIOSLobApp 
type ManagedIOSLobApp struct {
    ManagedMobileLobApp
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Contains properties of the possible iOS device types the mobile app can run on.
    applicableDeviceType IosDeviceTypeable
    // The build number of managed iOS Line of Business (LoB) app.
    buildNumber *string
    // The Identity Name.
    bundleId *string
    // The expiration time.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The identity version.
    identityVersion *string
    // The value for the minimum applicable operating system.
    minimumSupportedOperatingSystem IosMinimumOperatingSystemable
    // The version number of managed iOS Line of Business (LoB) app.
    versionNumber *string
}
// NewManagedIOSLobApp instantiates a new ManagedIOSLobApp and sets the default values.
func NewManagedIOSLobApp()(*ManagedIOSLobApp) {
    m := &ManagedIOSLobApp{
        ManagedMobileLobApp: *NewManagedMobileLobApp(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagedIOSLobAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedIOSLobAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedIOSLobApp(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedIOSLobApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicableDeviceType gets the applicableDeviceType property value. Contains properties of the possible iOS device types the mobile app can run on.
func (m *ManagedIOSLobApp) GetApplicableDeviceType()(IosDeviceTypeable) {
    if m == nil {
        return nil
    } else {
        return m.applicableDeviceType
    }
}
// GetBuildNumber gets the buildNumber property value. The build number of managed iOS Line of Business (LoB) app.
func (m *ManagedIOSLobApp) GetBuildNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.buildNumber
    }
}
// GetBundleId gets the bundleId property value. The Identity Name.
func (m *ManagedIOSLobApp) GetBundleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bundleId
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. The expiration time.
func (m *ManagedIOSLobApp) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedIOSLobApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedMobileLobApp.GetFieldDeserializers()
    res["applicableDeviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosDeviceTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableDeviceType(val.(IosDeviceTypeable))
        }
        return nil
    }
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
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
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
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(IosMinimumOperatingSystemable))
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
func (m *ManagedIOSLobApp) GetIdentityVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityVersion
    }
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *ManagedIOSLobApp) GetMinimumSupportedOperatingSystem()(IosMinimumOperatingSystemable) {
    if m == nil {
        return nil
    } else {
        return m.minimumSupportedOperatingSystem
    }
}
// GetVersionNumber gets the versionNumber property value. The version number of managed iOS Line of Business (LoB) app.
func (m *ManagedIOSLobApp) GetVersionNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.versionNumber
    }
}
// Serialize serializes information the current object
func (m *ManagedIOSLobApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedMobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("applicableDeviceType", m.GetApplicableDeviceType())
        if err != nil {
            return err
        }
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
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
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
func (m *ManagedIOSLobApp) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicableDeviceType sets the applicableDeviceType property value. Contains properties of the possible iOS device types the mobile app can run on.
func (m *ManagedIOSLobApp) SetApplicableDeviceType(value IosDeviceTypeable)() {
    if m != nil {
        m.applicableDeviceType = value
    }
}
// SetBuildNumber sets the buildNumber property value. The build number of managed iOS Line of Business (LoB) app.
func (m *ManagedIOSLobApp) SetBuildNumber(value *string)() {
    if m != nil {
        m.buildNumber = value
    }
}
// SetBundleId sets the bundleId property value. The Identity Name.
func (m *ManagedIOSLobApp) SetBundleId(value *string)() {
    if m != nil {
        m.bundleId = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The expiration time.
func (m *ManagedIOSLobApp) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetIdentityVersion sets the identityVersion property value. The identity version.
func (m *ManagedIOSLobApp) SetIdentityVersion(value *string)() {
    if m != nil {
        m.identityVersion = value
    }
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *ManagedIOSLobApp) SetMinimumSupportedOperatingSystem(value IosMinimumOperatingSystemable)() {
    if m != nil {
        m.minimumSupportedOperatingSystem = value
    }
}
// SetVersionNumber sets the versionNumber property value. The version number of managed iOS Line of Business (LoB) app.
func (m *ManagedIOSLobApp) SetVersionNumber(value *string)() {
    if m != nil {
        m.versionNumber = value
    }
}
