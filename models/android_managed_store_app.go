package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreApp 
type AndroidManagedStoreApp struct {
    MobileApp
    // The Identity Name.
    appIdentifier *string
    // The Play for Work Store app URL.
    appStoreUrl *string
    // The tracks that are visible to this enterprise.
    appTracks []AndroidManagedStoreAppTrackable
    // Indicates whether the app is only available to a given enterprise's users.
    isPrivate *bool
    // Indicates whether the app is a preinstalled system app.
    isSystemApp *bool
    // The package identifier.
    packageId *string
    // Whether this app supports OEMConfig policy.
    supportsOemConfig *bool
    // The total number of VPP licenses.
    totalLicenseCount *int32
    // The number of VPP licenses in use.
    usedLicenseCount *int32
}
// NewAndroidManagedStoreApp instantiates a new AndroidManagedStoreApp and sets the default values.
func NewAndroidManagedStoreApp()(*AndroidManagedStoreApp) {
    m := &AndroidManagedStoreApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.androidManagedStoreApp";
    m.SetOdataType(&odataTypeValue);
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
    return m.appIdentifier
}
// GetAppStoreUrl gets the appStoreUrl property value. The Play for Work Store app URL.
func (m *AndroidManagedStoreApp) GetAppStoreUrl()(*string) {
    return m.appStoreUrl
}
// GetAppTracks gets the appTracks property value. The tracks that are visible to this enterprise.
func (m *AndroidManagedStoreApp) GetAppTracks()([]AndroidManagedStoreAppTrackable) {
    return m.appTracks
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appIdentifier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppIdentifier)
    res["appStoreUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAppStoreUrl)
    res["appTracks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAndroidManagedStoreAppTrackFromDiscriminatorValue , m.SetAppTracks)
    res["isPrivate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsPrivate)
    res["isSystemApp"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSystemApp)
    res["packageId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPackageId)
    res["supportsOemConfig"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSupportsOemConfig)
    res["totalLicenseCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalLicenseCount)
    res["usedLicenseCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUsedLicenseCount)
    return res
}
// GetIsPrivate gets the isPrivate property value. Indicates whether the app is only available to a given enterprise's users.
func (m *AndroidManagedStoreApp) GetIsPrivate()(*bool) {
    return m.isPrivate
}
// GetIsSystemApp gets the isSystemApp property value. Indicates whether the app is a preinstalled system app.
func (m *AndroidManagedStoreApp) GetIsSystemApp()(*bool) {
    return m.isSystemApp
}
// GetPackageId gets the packageId property value. The package identifier.
func (m *AndroidManagedStoreApp) GetPackageId()(*string) {
    return m.packageId
}
// GetSupportsOemConfig gets the supportsOemConfig property value. Whether this app supports OEMConfig policy.
func (m *AndroidManagedStoreApp) GetSupportsOemConfig()(*bool) {
    return m.supportsOemConfig
}
// GetTotalLicenseCount gets the totalLicenseCount property value. The total number of VPP licenses.
func (m *AndroidManagedStoreApp) GetTotalLicenseCount()(*int32) {
    return m.totalLicenseCount
}
// GetUsedLicenseCount gets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *AndroidManagedStoreApp) GetUsedLicenseCount()(*int32) {
    return m.usedLicenseCount
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppTracks())
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
    m.appIdentifier = value
}
// SetAppStoreUrl sets the appStoreUrl property value. The Play for Work Store app URL.
func (m *AndroidManagedStoreApp) SetAppStoreUrl(value *string)() {
    m.appStoreUrl = value
}
// SetAppTracks sets the appTracks property value. The tracks that are visible to this enterprise.
func (m *AndroidManagedStoreApp) SetAppTracks(value []AndroidManagedStoreAppTrackable)() {
    m.appTracks = value
}
// SetIsPrivate sets the isPrivate property value. Indicates whether the app is only available to a given enterprise's users.
func (m *AndroidManagedStoreApp) SetIsPrivate(value *bool)() {
    m.isPrivate = value
}
// SetIsSystemApp sets the isSystemApp property value. Indicates whether the app is a preinstalled system app.
func (m *AndroidManagedStoreApp) SetIsSystemApp(value *bool)() {
    m.isSystemApp = value
}
// SetPackageId sets the packageId property value. The package identifier.
func (m *AndroidManagedStoreApp) SetPackageId(value *string)() {
    m.packageId = value
}
// SetSupportsOemConfig sets the supportsOemConfig property value. Whether this app supports OEMConfig policy.
func (m *AndroidManagedStoreApp) SetSupportsOemConfig(value *bool)() {
    m.supportsOemConfig = value
}
// SetTotalLicenseCount sets the totalLicenseCount property value. The total number of VPP licenses.
func (m *AndroidManagedStoreApp) SetTotalLicenseCount(value *int32)() {
    m.totalLicenseCount = value
}
// SetUsedLicenseCount sets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *AndroidManagedStoreApp) SetUsedLicenseCount(value *int32)() {
    m.usedLicenseCount = value
}
