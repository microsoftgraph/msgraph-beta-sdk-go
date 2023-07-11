package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreApp contains properties and inherited properties for Android Managed Store Apps.
type AndroidManagedStoreApp struct {
    MobileApp
    // The OdataType property
    OdataType *string
}
// NewAndroidManagedStoreApp instantiates a new androidManagedStoreApp and sets the default values.
func NewAndroidManagedStoreApp()(*AndroidManagedStoreApp) {
    m := &AndroidManagedStoreApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.androidManagedStoreApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidManagedStoreAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.androidManagedStoreWebApp":
                        return NewAndroidManagedStoreWebApp(), nil
                }
            }
        }
    }
    return NewAndroidManagedStoreApp(), nil
}
// GetAppIdentifier gets the appIdentifier property value. The Identity Name.
func (m *AndroidManagedStoreApp) GetAppIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("appIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppStoreUrl gets the appStoreUrl property value. The Play for Work Store app URL.
func (m *AndroidManagedStoreApp) GetAppStoreUrl()(*string) {
    val, err := m.GetBackingStore().Get("appStoreUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppTracks gets the appTracks property value. The tracks that are visible to this enterprise.
func (m *AndroidManagedStoreApp) GetAppTracks()([]AndroidManagedStoreAppTrackable) {
    val, err := m.GetBackingStore().Get("appTracks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidManagedStoreAppTrackable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppIdentifier(val)
        }
        return nil
    }
    res["appStoreUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreUrl(val)
        }
        return nil
    }
    res["appTracks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidManagedStoreAppTrackFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidManagedStoreAppTrackable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AndroidManagedStoreAppTrackable)
                }
            }
            m.SetAppTracks(res)
        }
        return nil
    }
    res["isPrivate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPrivate(val)
        }
        return nil
    }
    res["isSystemApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSystemApp(val)
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
    res["supportsOemConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportsOemConfig(val)
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
// GetIsPrivate gets the isPrivate property value. Indicates whether the app is only available to a given enterprise's users.
func (m *AndroidManagedStoreApp) GetIsPrivate()(*bool) {
    val, err := m.GetBackingStore().Get("isPrivate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSystemApp gets the isSystemApp property value. Indicates whether the app is a preinstalled system app.
func (m *AndroidManagedStoreApp) GetIsSystemApp()(*bool) {
    val, err := m.GetBackingStore().Get("isSystemApp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPackageId gets the packageId property value. The package identifier.
func (m *AndroidManagedStoreApp) GetPackageId()(*string) {
    val, err := m.GetBackingStore().Get("packageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSupportsOemConfig gets the supportsOemConfig property value. Whether this app supports OEMConfig policy.
func (m *AndroidManagedStoreApp) GetSupportsOemConfig()(*bool) {
    val, err := m.GetBackingStore().Get("supportsOemConfig")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTotalLicenseCount gets the totalLicenseCount property value. The total number of VPP licenses.
func (m *AndroidManagedStoreApp) GetTotalLicenseCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalLicenseCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUsedLicenseCount gets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *AndroidManagedStoreApp) GetUsedLicenseCount()(*int32) {
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
func (m *AndroidManagedStoreApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appIdentifier", m.GetAppIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appStoreUrl", m.GetAppStoreUrl())
        if err != nil {
            return err
        }
    }
    if m.GetAppTracks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppTracks()))
        for i, v := range m.GetAppTracks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appTracks", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPrivate", m.GetIsPrivate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSystemApp", m.GetIsSystemApp())
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
    {
        err = writer.WriteBoolValue("supportsOemConfig", m.GetSupportsOemConfig())
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
// SetAppIdentifier sets the appIdentifier property value. The Identity Name.
func (m *AndroidManagedStoreApp) SetAppIdentifier(value *string)() {
    err := m.GetBackingStore().Set("appIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetAppStoreUrl sets the appStoreUrl property value. The Play for Work Store app URL.
func (m *AndroidManagedStoreApp) SetAppStoreUrl(value *string)() {
    err := m.GetBackingStore().Set("appStoreUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetAppTracks sets the appTracks property value. The tracks that are visible to this enterprise.
func (m *AndroidManagedStoreApp) SetAppTracks(value []AndroidManagedStoreAppTrackable)() {
    err := m.GetBackingStore().Set("appTracks", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPrivate sets the isPrivate property value. Indicates whether the app is only available to a given enterprise's users.
func (m *AndroidManagedStoreApp) SetIsPrivate(value *bool)() {
    err := m.GetBackingStore().Set("isPrivate", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSystemApp sets the isSystemApp property value. Indicates whether the app is a preinstalled system app.
func (m *AndroidManagedStoreApp) SetIsSystemApp(value *bool)() {
    err := m.GetBackingStore().Set("isSystemApp", value)
    if err != nil {
        panic(err)
    }
}
// SetPackageId sets the packageId property value. The package identifier.
func (m *AndroidManagedStoreApp) SetPackageId(value *string)() {
    err := m.GetBackingStore().Set("packageId", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportsOemConfig sets the supportsOemConfig property value. Whether this app supports OEMConfig policy.
func (m *AndroidManagedStoreApp) SetSupportsOemConfig(value *bool)() {
    err := m.GetBackingStore().Set("supportsOemConfig", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalLicenseCount sets the totalLicenseCount property value. The total number of VPP licenses.
func (m *AndroidManagedStoreApp) SetTotalLicenseCount(value *int32)() {
    err := m.GetBackingStore().Set("totalLicenseCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUsedLicenseCount sets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *AndroidManagedStoreApp) SetUsedLicenseCount(value *int32)() {
    err := m.GetBackingStore().Set("usedLicenseCount", value)
    if err != nil {
        panic(err)
    }
}
// AndroidManagedStoreAppable 
type AndroidManagedStoreAppable interface {
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppIdentifier()(*string)
    GetAppStoreUrl()(*string)
    GetAppTracks()([]AndroidManagedStoreAppTrackable)
    GetIsPrivate()(*bool)
    GetIsSystemApp()(*bool)
    GetPackageId()(*string)
    GetSupportsOemConfig()(*bool)
    GetTotalLicenseCount()(*int32)
    GetUsedLicenseCount()(*int32)
    SetAppIdentifier(value *string)()
    SetAppStoreUrl(value *string)()
    SetAppTracks(value []AndroidManagedStoreAppTrackable)()
    SetIsPrivate(value *bool)()
    SetIsSystemApp(value *bool)()
    SetPackageId(value *string)()
    SetSupportsOemConfig(value *bool)()
    SetTotalLicenseCount(value *int32)()
    SetUsedLicenseCount(value *int32)()
}
