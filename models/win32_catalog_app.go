package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32CatalogApp a mobileApp that is based on a referenced application in a Win32CatalogApp repository
type Win32CatalogApp struct {
    Win32LobApp
}
// NewWin32CatalogApp instantiates a new Win32CatalogApp and sets the default values.
func NewWin32CatalogApp()(*Win32CatalogApp) {
    m := &Win32CatalogApp{
        Win32LobApp: *NewWin32LobApp(),
    }
    odataTypeValue := "#microsoft.graph.win32CatalogApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32CatalogAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWin32CatalogAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32CatalogApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Win32CatalogApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Win32LobApp.GetFieldDeserializers()
    res["latestUpgradeCatalogPackage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppCatalogPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestUpgradeCatalogPackage(val.(MobileAppCatalogPackageable))
        }
        return nil
    }
    res["mobileAppCatalogPackageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileAppCatalogPackageId(val)
        }
        return nil
    }
    res["referencedCatalogPackage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppCatalogPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferencedCatalogPackage(val.(MobileAppCatalogPackageable))
        }
        return nil
    }
    return res
}
// GetLatestUpgradeCatalogPackage gets the latestUpgradeCatalogPackage property value. The latest available catalog package the app is upgradeable to. This property is read-only.
// returns a MobileAppCatalogPackageable when successful
func (m *Win32CatalogApp) GetLatestUpgradeCatalogPackage()(MobileAppCatalogPackageable) {
    val, err := m.GetBackingStore().Get("latestUpgradeCatalogPackage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MobileAppCatalogPackageable)
    }
    return nil
}
// GetMobileAppCatalogPackageId gets the mobileAppCatalogPackageId property value. The mobileAppCatalogPackageId property references the mobileAppCatalogPackage entity which contains information about an application catalog package that can be deployed to Intune-managed devices
// returns a *string when successful
func (m *Win32CatalogApp) GetMobileAppCatalogPackageId()(*string) {
    val, err := m.GetBackingStore().Get("mobileAppCatalogPackageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReferencedCatalogPackage gets the referencedCatalogPackage property value. The current catalog package the app is synced from. This property is read-only.
// returns a MobileAppCatalogPackageable when successful
func (m *Win32CatalogApp) GetReferencedCatalogPackage()(MobileAppCatalogPackageable) {
    val, err := m.GetBackingStore().Get("referencedCatalogPackage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MobileAppCatalogPackageable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Win32CatalogApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Win32LobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("latestUpgradeCatalogPackage", m.GetLatestUpgradeCatalogPackage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mobileAppCatalogPackageId", m.GetMobileAppCatalogPackageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("referencedCatalogPackage", m.GetReferencedCatalogPackage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLatestUpgradeCatalogPackage sets the latestUpgradeCatalogPackage property value. The latest available catalog package the app is upgradeable to. This property is read-only.
func (m *Win32CatalogApp) SetLatestUpgradeCatalogPackage(value MobileAppCatalogPackageable)() {
    err := m.GetBackingStore().Set("latestUpgradeCatalogPackage", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileAppCatalogPackageId sets the mobileAppCatalogPackageId property value. The mobileAppCatalogPackageId property references the mobileAppCatalogPackage entity which contains information about an application catalog package that can be deployed to Intune-managed devices
func (m *Win32CatalogApp) SetMobileAppCatalogPackageId(value *string)() {
    err := m.GetBackingStore().Set("mobileAppCatalogPackageId", value)
    if err != nil {
        panic(err)
    }
}
// SetReferencedCatalogPackage sets the referencedCatalogPackage property value. The current catalog package the app is synced from. This property is read-only.
func (m *Win32CatalogApp) SetReferencedCatalogPackage(value MobileAppCatalogPackageable)() {
    err := m.GetBackingStore().Set("referencedCatalogPackage", value)
    if err != nil {
        panic(err)
    }
}
type Win32CatalogAppable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppable
    GetLatestUpgradeCatalogPackage()(MobileAppCatalogPackageable)
    GetMobileAppCatalogPackageId()(*string)
    GetReferencedCatalogPackage()(MobileAppCatalogPackageable)
    SetLatestUpgradeCatalogPackage(value MobileAppCatalogPackageable)()
    SetMobileAppCatalogPackageId(value *string)()
    SetReferencedCatalogPackage(value MobileAppCatalogPackageable)()
}
