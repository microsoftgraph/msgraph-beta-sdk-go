package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftStoreForBusinessApp microsoft Store for Business Apps. This class does not support Create, Delete, or Update.
type MicrosoftStoreForBusinessApp struct {
    MobileApp
    // The OdataType property
    OdataType *string
}
// NewMicrosoftStoreForBusinessApp instantiates a new microsoftStoreForBusinessApp and sets the default values.
func NewMicrosoftStoreForBusinessApp()(*MicrosoftStoreForBusinessApp) {
    m := &MicrosoftStoreForBusinessApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.microsoftStoreForBusinessApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMicrosoftStoreForBusinessAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftStoreForBusinessAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftStoreForBusinessApp(), nil
}
// GetContainedApps gets the containedApps property value. The collection of contained apps in a mobileApp acting as a package.
func (m *MicrosoftStoreForBusinessApp) GetContainedApps()([]MobileContainedAppable) {
    val, err := m.GetBackingStore().Get("containedApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MobileContainedAppable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftStoreForBusinessApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["containedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileContainedAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileContainedAppable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MobileContainedAppable)
                }
            }
            m.SetContainedApps(res)
        }
        return nil
    }
    res["licenseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftStoreForBusinessLicenseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseType(val.(*MicrosoftStoreForBusinessLicenseType))
        }
        return nil
    }
    res["licensingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVppLicensingTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicensingType(val.(VppLicensingTypeable))
        }
        return nil
    }
    res["packageIdentityName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageIdentityName(val)
        }
        return nil
    }
    res["productKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductKey(val)
        }
        return nil
    }
    res["totalLicenseCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLicenseCount(val)
        }
        return nil
    }
    res["usedLicenseCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedLicenseCount(val)
        }
        return nil
    }
    return res
}
// GetLicenseType gets the licenseType property value. The licenseType property
func (m *MicrosoftStoreForBusinessApp) GetLicenseType()(*MicrosoftStoreForBusinessLicenseType) {
    val, err := m.GetBackingStore().Get("licenseType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MicrosoftStoreForBusinessLicenseType)
    }
    return nil
}
// GetLicensingType gets the licensingType property value. The supported License Type.
func (m *MicrosoftStoreForBusinessApp) GetLicensingType()(VppLicensingTypeable) {
    val, err := m.GetBackingStore().Get("licensingType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VppLicensingTypeable)
    }
    return nil
}
// GetPackageIdentityName gets the packageIdentityName property value. The app package identifier
func (m *MicrosoftStoreForBusinessApp) GetPackageIdentityName()(*string) {
    val, err := m.GetBackingStore().Get("packageIdentityName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProductKey gets the productKey property value. The app product key
func (m *MicrosoftStoreForBusinessApp) GetProductKey()(*string) {
    val, err := m.GetBackingStore().Get("productKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotalLicenseCount gets the totalLicenseCount property value. The total number of Microsoft Store for Business licenses.
func (m *MicrosoftStoreForBusinessApp) GetTotalLicenseCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalLicenseCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUsedLicenseCount gets the usedLicenseCount property value. The number of Microsoft Store for Business licenses in use.
func (m *MicrosoftStoreForBusinessApp) GetUsedLicenseCount()(*int32) {
    val, err := m.GetBackingStore().Get("usedLicenseCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MicrosoftStoreForBusinessApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetContainedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContainedApps()))
        for i, v := range m.GetContainedApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("containedApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLicenseType() != nil {
        cast := (*m.GetLicenseType()).String()
        err = writer.WriteStringValue("licenseType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("licensingType", m.GetLicensingType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("packageIdentityName", m.GetPackageIdentityName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productKey", m.GetProductKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalLicenseCount", m.GetTotalLicenseCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("usedLicenseCount", m.GetUsedLicenseCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContainedApps sets the containedApps property value. The collection of contained apps in a mobileApp acting as a package.
func (m *MicrosoftStoreForBusinessApp) SetContainedApps(value []MobileContainedAppable)() {
    err := m.GetBackingStore().Set("containedApps", value)
    if err != nil {
        panic(err)
    }
}
// SetLicenseType sets the licenseType property value. The licenseType property
func (m *MicrosoftStoreForBusinessApp) SetLicenseType(value *MicrosoftStoreForBusinessLicenseType)() {
    err := m.GetBackingStore().Set("licenseType", value)
    if err != nil {
        panic(err)
    }
}
// SetLicensingType sets the licensingType property value. The supported License Type.
func (m *MicrosoftStoreForBusinessApp) SetLicensingType(value VppLicensingTypeable)() {
    err := m.GetBackingStore().Set("licensingType", value)
    if err != nil {
        panic(err)
    }
}
// SetPackageIdentityName sets the packageIdentityName property value. The app package identifier
func (m *MicrosoftStoreForBusinessApp) SetPackageIdentityName(value *string)() {
    err := m.GetBackingStore().Set("packageIdentityName", value)
    if err != nil {
        panic(err)
    }
}
// SetProductKey sets the productKey property value. The app product key
func (m *MicrosoftStoreForBusinessApp) SetProductKey(value *string)() {
    err := m.GetBackingStore().Set("productKey", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalLicenseCount sets the totalLicenseCount property value. The total number of Microsoft Store for Business licenses.
func (m *MicrosoftStoreForBusinessApp) SetTotalLicenseCount(value *int32)() {
    err := m.GetBackingStore().Set("totalLicenseCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUsedLicenseCount sets the usedLicenseCount property value. The number of Microsoft Store for Business licenses in use.
func (m *MicrosoftStoreForBusinessApp) SetUsedLicenseCount(value *int32)() {
    err := m.GetBackingStore().Set("usedLicenseCount", value)
    if err != nil {
        panic(err)
    }
}
// MicrosoftStoreForBusinessAppable 
type MicrosoftStoreForBusinessAppable interface {
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContainedApps()([]MobileContainedAppable)
    GetLicenseType()(*MicrosoftStoreForBusinessLicenseType)
    GetLicensingType()(VppLicensingTypeable)
    GetPackageIdentityName()(*string)
    GetProductKey()(*string)
    GetTotalLicenseCount()(*int32)
    GetUsedLicenseCount()(*int32)
    SetContainedApps(value []MobileContainedAppable)()
    SetLicenseType(value *MicrosoftStoreForBusinessLicenseType)()
    SetLicensingType(value VppLicensingTypeable)()
    SetPackageIdentityName(value *string)()
    SetProductKey(value *string)()
    SetTotalLicenseCount(value *int32)()
    SetUsedLicenseCount(value *int32)()
}
