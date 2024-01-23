package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32CatalogApp a mobileApp that is based on a referenced application in a Win32CatalogApp repository
type Win32CatalogApp struct {
    Win32LobApp
}
// NewWin32CatalogApp instantiates a new win32CatalogApp and sets the default values.
func NewWin32CatalogApp()(*Win32CatalogApp) {
    m := &Win32CatalogApp{
        Win32LobApp: *NewWin32LobApp(),
    }
    odataTypeValue := "#microsoft.graph.win32CatalogApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32CatalogAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32CatalogAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32CatalogApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32CatalogApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Win32LobApp.GetFieldDeserializers()
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
    return res
}
// GetMobileAppCatalogPackageId gets the mobileAppCatalogPackageId property value. The mobileAppCatalogPackageId property references the mobileAppCatalogPackage entity which contains information about an application catalog package that can be deployed to Intune-managed devices
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
// Serialize serializes information the current object
func (m *Win32CatalogApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Win32LobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("mobileAppCatalogPackageId", m.GetMobileAppCatalogPackageId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMobileAppCatalogPackageId sets the mobileAppCatalogPackageId property value. The mobileAppCatalogPackageId property references the mobileAppCatalogPackage entity which contains information about an application catalog package that can be deployed to Intune-managed devices
func (m *Win32CatalogApp) SetMobileAppCatalogPackageId(value *string)() {
    err := m.GetBackingStore().Set("mobileAppCatalogPackageId", value)
    if err != nil {
        panic(err)
    }
}
// Win32CatalogAppable 
type Win32CatalogAppable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppable
    GetMobileAppCatalogPackageId()(*string)
    SetMobileAppCatalogPackageId(value *string)()
}
