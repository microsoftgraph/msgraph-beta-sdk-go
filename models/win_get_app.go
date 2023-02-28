package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WinGetApp 
type WinGetApp struct {
    MobileApp
}
// NewWinGetApp instantiates a new WinGetApp and sets the default values.
func NewWinGetApp()(*WinGetApp) {
    m := &WinGetApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.winGetApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWinGetAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWinGetAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWinGetApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WinGetApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["installExperience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWinGetAppInstallExperienceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallExperience(val.(WinGetAppInstallExperienceable))
        }
        return nil
    }
    res["manifestHash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManifestHash(val)
        }
        return nil
    }
    res["packageIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageIdentifier(val)
        }
        return nil
    }
    return res
}
// GetInstallExperience gets the installExperience property value. The install experience settings associated with this application, which are used to ensure the desired install experiences on the target device are taken into account. This includes the account type (System or User) that actions should be run as on target devices. Required at creation time.
func (m *WinGetApp) GetInstallExperience()(WinGetAppInstallExperienceable) {
    val, err := m.GetBackingStore().Get("installExperience")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WinGetAppInstallExperienceable)
    }
    return nil
}
// GetManifestHash gets the manifestHash property value. Hash of package metadata properties used to validate that the application matches the metadata in the source repository.
func (m *WinGetApp) GetManifestHash()(*string) {
    val, err := m.GetBackingStore().Get("manifestHash")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPackageIdentifier gets the packageIdentifier property value. The PackageIdentifier from the WinGet source repository REST API. This also maps to the Id when using the WinGet client command line application. Required at creation time, cannot be modified on existing objects.
func (m *WinGetApp) GetPackageIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("packageIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WinGetApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("installExperience", m.GetInstallExperience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manifestHash", m.GetManifestHash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("packageIdentifier", m.GetPackageIdentifier())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInstallExperience sets the installExperience property value. The install experience settings associated with this application, which are used to ensure the desired install experiences on the target device are taken into account. This includes the account type (System or User) that actions should be run as on target devices. Required at creation time.
func (m *WinGetApp) SetInstallExperience(value WinGetAppInstallExperienceable)() {
    err := m.GetBackingStore().Set("installExperience", value)
    if err != nil {
        panic(err)
    }
}
// SetManifestHash sets the manifestHash property value. Hash of package metadata properties used to validate that the application matches the metadata in the source repository.
func (m *WinGetApp) SetManifestHash(value *string)() {
    err := m.GetBackingStore().Set("manifestHash", value)
    if err != nil {
        panic(err)
    }
}
// SetPackageIdentifier sets the packageIdentifier property value. The PackageIdentifier from the WinGet source repository REST API. This also maps to the Id when using the WinGet client command line application. Required at creation time, cannot be modified on existing objects.
func (m *WinGetApp) SetPackageIdentifier(value *string)() {
    err := m.GetBackingStore().Set("packageIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// WinGetAppable 
type WinGetAppable interface {
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInstallExperience()(WinGetAppInstallExperienceable)
    GetManifestHash()(*string)
    GetPackageIdentifier()(*string)
    SetInstallExperience(value WinGetAppInstallExperienceable)()
    SetManifestHash(value *string)()
    SetPackageIdentifier(value *string)()
}
