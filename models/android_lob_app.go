package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidLobApp contains properties and inherited properties for Android Line Of Business apps.
type AndroidLobApp struct {
    MobileLobApp
}
// NewAndroidLobApp instantiates a new AndroidLobApp and sets the default values.
func NewAndroidLobApp()(*AndroidLobApp) {
    m := &AndroidLobApp{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.androidLobApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidLobAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAndroidLobAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidLobApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AndroidLobApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(AndroidMinimumOperatingSystemable))
        }
        return nil
    }
    res["packageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageId(val)
        }
        return nil
    }
    res["targetedPlatforms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidTargetedPlatforms)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedPlatforms(val.(*AndroidTargetedPlatforms))
        }
        return nil
    }
    res["versionCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionCode(val)
        }
        return nil
    }
    res["versionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionName(val)
        }
        return nil
    }
    return res
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
// returns a AndroidMinimumOperatingSystemable when successful
func (m *AndroidLobApp) GetMinimumSupportedOperatingSystem()(AndroidMinimumOperatingSystemable) {
    val, err := m.GetBackingStore().Get("minimumSupportedOperatingSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidMinimumOperatingSystemable)
    }
    return nil
}
// GetPackageId gets the packageId property value. The package identifier.
// returns a *string when successful
func (m *AndroidLobApp) GetPackageId()(*string) {
    val, err := m.GetBackingStore().Get("packageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetedPlatforms gets the targetedPlatforms property value. Specifies which platform(s) can be targeted for a given Android LOB application or Managed Android LOB application.
// returns a *AndroidTargetedPlatforms when successful
func (m *AndroidLobApp) GetTargetedPlatforms()(*AndroidTargetedPlatforms) {
    val, err := m.GetBackingStore().Get("targetedPlatforms")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidTargetedPlatforms)
    }
    return nil
}
// GetVersionCode gets the versionCode property value. The version code of Android Line of Business (LoB) app.
// returns a *string when successful
func (m *AndroidLobApp) GetVersionCode()(*string) {
    val, err := m.GetBackingStore().Get("versionCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVersionName gets the versionName property value. The version name of Android Line of Business (LoB) app.
// returns a *string when successful
func (m *AndroidLobApp) GetVersionName()(*string) {
    val, err := m.GetBackingStore().Get("versionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidLobApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("packageId", m.GetPackageId())
        if err != nil {
            return err
        }
    }
    if m.GetTargetedPlatforms() != nil {
        cast := (*m.GetTargetedPlatforms()).String()
        err = writer.WriteStringValue("targetedPlatforms", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("versionCode", m.GetVersionCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("versionName", m.GetVersionName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *AndroidLobApp) SetMinimumSupportedOperatingSystem(value AndroidMinimumOperatingSystemable)() {
    err := m.GetBackingStore().Set("minimumSupportedOperatingSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetPackageId sets the packageId property value. The package identifier.
func (m *AndroidLobApp) SetPackageId(value *string)() {
    err := m.GetBackingStore().Set("packageId", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedPlatforms sets the targetedPlatforms property value. Specifies which platform(s) can be targeted for a given Android LOB application or Managed Android LOB application.
func (m *AndroidLobApp) SetTargetedPlatforms(value *AndroidTargetedPlatforms)() {
    err := m.GetBackingStore().Set("targetedPlatforms", value)
    if err != nil {
        panic(err)
    }
}
// SetVersionCode sets the versionCode property value. The version code of Android Line of Business (LoB) app.
func (m *AndroidLobApp) SetVersionCode(value *string)() {
    err := m.GetBackingStore().Set("versionCode", value)
    if err != nil {
        panic(err)
    }
}
// SetVersionName sets the versionName property value. The version name of Android Line of Business (LoB) app.
func (m *AndroidLobApp) SetVersionName(value *string)() {
    err := m.GetBackingStore().Set("versionName", value)
    if err != nil {
        panic(err)
    }
}
type AndroidLobAppable interface {
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMinimumSupportedOperatingSystem()(AndroidMinimumOperatingSystemable)
    GetPackageId()(*string)
    GetTargetedPlatforms()(*AndroidTargetedPlatforms)
    GetVersionCode()(*string)
    GetVersionName()(*string)
    SetMinimumSupportedOperatingSystem(value AndroidMinimumOperatingSystemable)()
    SetPackageId(value *string)()
    SetTargetedPlatforms(value *AndroidTargetedPlatforms)()
    SetVersionCode(value *string)()
    SetVersionName(value *string)()
}
