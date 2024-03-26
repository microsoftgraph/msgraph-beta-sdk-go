package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32MobileAppCatalogPackage win32MobileAppCatalogPackage extends mobileAppCatalogPackage by providing information necessary for the creation of a win32CatalogApp instance.
type Win32MobileAppCatalogPackage struct {
    MobileAppCatalogPackage
}
// NewWin32MobileAppCatalogPackage instantiates a new Win32MobileAppCatalogPackage and sets the default values.
func NewWin32MobileAppCatalogPackage()(*Win32MobileAppCatalogPackage) {
    m := &Win32MobileAppCatalogPackage{
        MobileAppCatalogPackage: *NewMobileAppCatalogPackage(),
    }
    odataTypeValue := "#microsoft.graph.win32MobileAppCatalogPackage"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32MobileAppCatalogPackageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWin32MobileAppCatalogPackageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32MobileAppCatalogPackage(), nil
}
// GetApplicableArchitectures gets the applicableArchitectures property value. Contains properties for Windows architecture.
// returns a *WindowsArchitecture when successful
func (m *Win32MobileAppCatalogPackage) GetApplicableArchitectures()(*WindowsArchitecture) {
    val, err := m.GetBackingStore().Get("applicableArchitectures")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsArchitecture)
    }
    return nil
}
// GetBranchDisplayName gets the branchDisplayName property value. The product branch name, which is a specific subset of product functionality as defined by the publisher (example: "Fabrikam for Business (x64)"). A specific product will have one or more branchDisplayNames. Read-only. Supports $filter, $search, $select. This property is read-only.
// returns a *string when successful
func (m *Win32MobileAppCatalogPackage) GetBranchDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("branchDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Win32MobileAppCatalogPackage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppCatalogPackage.GetFieldDeserializers()
    res["applicableArchitectures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableArchitectures(val.(*WindowsArchitecture))
        }
        return nil
    }
    res["branchDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranchDisplayName(val)
        }
        return nil
    }
    res["locales"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetLocales(res)
        }
        return nil
    }
    res["packageAutoUpdateCapable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageAutoUpdateCapable(val)
        }
        return nil
    }
    return res
}
// GetLocales gets the locales property value. One or more locale(s) supported by the branch. Value is a two-letter ISO 639 language tags with optional two-letter subtags (example: en-US, ko, de, de-DE), or mul to indicate multi-language. Read-only. This property is read-only.
// returns a []string when successful
func (m *Win32MobileAppCatalogPackage) GetLocales()([]string) {
    val, err := m.GetBackingStore().Get("locales")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetPackageAutoUpdateCapable gets the packageAutoUpdateCapable property value. Indicates whether the package is capable to auto-update to latest when software/application updates are available. When TRUE, it indicates it is an auto-updating application. When FALSE, it indicates that it is not an auto-updating application. This property is read-only.
// returns a *bool when successful
func (m *Win32MobileAppCatalogPackage) GetPackageAutoUpdateCapable()(*bool) {
    val, err := m.GetBackingStore().Get("packageAutoUpdateCapable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Win32MobileAppCatalogPackage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppCatalogPackage.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicableArchitectures() != nil {
        cast := (*m.GetApplicableArchitectures()).String()
        err = writer.WriteStringValue("applicableArchitectures", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicableArchitectures sets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *Win32MobileAppCatalogPackage) SetApplicableArchitectures(value *WindowsArchitecture)() {
    err := m.GetBackingStore().Set("applicableArchitectures", value)
    if err != nil {
        panic(err)
    }
}
// SetBranchDisplayName sets the branchDisplayName property value. The product branch name, which is a specific subset of product functionality as defined by the publisher (example: "Fabrikam for Business (x64)"). A specific product will have one or more branchDisplayNames. Read-only. Supports $filter, $search, $select. This property is read-only.
func (m *Win32MobileAppCatalogPackage) SetBranchDisplayName(value *string)() {
    err := m.GetBackingStore().Set("branchDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLocales sets the locales property value. One or more locale(s) supported by the branch. Value is a two-letter ISO 639 language tags with optional two-letter subtags (example: en-US, ko, de, de-DE), or mul to indicate multi-language. Read-only. This property is read-only.
func (m *Win32MobileAppCatalogPackage) SetLocales(value []string)() {
    err := m.GetBackingStore().Set("locales", value)
    if err != nil {
        panic(err)
    }
}
// SetPackageAutoUpdateCapable sets the packageAutoUpdateCapable property value. Indicates whether the package is capable to auto-update to latest when software/application updates are available. When TRUE, it indicates it is an auto-updating application. When FALSE, it indicates that it is not an auto-updating application. This property is read-only.
func (m *Win32MobileAppCatalogPackage) SetPackageAutoUpdateCapable(value *bool)() {
    err := m.GetBackingStore().Set("packageAutoUpdateCapable", value)
    if err != nil {
        panic(err)
    }
}
type Win32MobileAppCatalogPackageable interface {
    MobileAppCatalogPackageable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableArchitectures()(*WindowsArchitecture)
    GetBranchDisplayName()(*string)
    GetLocales()([]string)
    GetPackageAutoUpdateCapable()(*bool)
    SetApplicableArchitectures(value *WindowsArchitecture)()
    SetBranchDisplayName(value *string)()
    SetLocales(value []string)()
    SetPackageAutoUpdateCapable(value *bool)()
}
