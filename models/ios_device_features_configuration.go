package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosDeviceFeaturesConfiguration iOS Device Features Configuration Profile.
type IosDeviceFeaturesConfiguration struct {
    AppleDeviceFeaturesConfigurationBase
}
// NewIosDeviceFeaturesConfiguration instantiates a new IosDeviceFeaturesConfiguration and sets the default values.
func NewIosDeviceFeaturesConfiguration()(*IosDeviceFeaturesConfiguration) {
    m := &IosDeviceFeaturesConfiguration{
        AppleDeviceFeaturesConfigurationBase: *NewAppleDeviceFeaturesConfigurationBase(),
    }
    odataTypeValue := "#microsoft.graph.iosDeviceFeaturesConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosDeviceFeaturesConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIosDeviceFeaturesConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosDeviceFeaturesConfiguration(), nil
}
// GetAssetTagTemplate gets the assetTagTemplate property value. Asset tag information for the device, displayed on the login window and lock screen.
// returns a *string when successful
func (m *IosDeviceFeaturesConfiguration) GetAssetTagTemplate()(*string) {
    val, err := m.GetBackingStore().Get("assetTagTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContentFilterSettings gets the contentFilterSettings property value. Gets or sets iOS Web Content Filter settings, supervised mode only
// returns a IosWebContentFilterBaseable when successful
func (m *IosDeviceFeaturesConfiguration) GetContentFilterSettings()(IosWebContentFilterBaseable) {
    val, err := m.GetBackingStore().Get("contentFilterSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosWebContentFilterBaseable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IosDeviceFeaturesConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AppleDeviceFeaturesConfigurationBase.GetFieldDeserializers()
    res["assetTagTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssetTagTemplate(val)
        }
        return nil
    }
    res["contentFilterSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosWebContentFilterBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentFilterSettings(val.(IosWebContentFilterBaseable))
        }
        return nil
    }
    res["homeScreenDockIcons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosHomeScreenItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosHomeScreenItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IosHomeScreenItemable)
                }
            }
            m.SetHomeScreenDockIcons(res)
        }
        return nil
    }
    res["homeScreenGridHeight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeScreenGridHeight(val)
        }
        return nil
    }
    res["homeScreenGridWidth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeScreenGridWidth(val)
        }
        return nil
    }
    res["homeScreenPages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosHomeScreenPageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosHomeScreenPageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IosHomeScreenPageable)
                }
            }
            m.SetHomeScreenPages(res)
        }
        return nil
    }
    res["identityCertificateForClientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityCertificateForClientAuthentication(val.(IosCertificateProfileBaseable))
        }
        return nil
    }
    res["iosSingleSignOnExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosSingleSignOnExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosSingleSignOnExtension(val.(IosSingleSignOnExtensionable))
        }
        return nil
    }
    res["lockScreenFootnote"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenFootnote(val)
        }
        return nil
    }
    res["notificationSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosNotificationSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosNotificationSettingsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IosNotificationSettingsable)
                }
            }
            m.SetNotificationSettings(res)
        }
        return nil
    }
    res["singleSignOnExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSingleSignOnExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnExtension(val.(SingleSignOnExtensionable))
        }
        return nil
    }
    res["singleSignOnExtensionPkinitCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnExtensionPkinitCertificate(val.(IosCertificateProfileBaseable))
        }
        return nil
    }
    res["singleSignOnSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosSingleSignOnSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnSettings(val.(IosSingleSignOnSettingsable))
        }
        return nil
    }
    res["wallpaperDisplayLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIosWallpaperDisplayLocation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWallpaperDisplayLocation(val.(*IosWallpaperDisplayLocation))
        }
        return nil
    }
    res["wallpaperImage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWallpaperImage(val.(MimeContentable))
        }
        return nil
    }
    return res
}
// GetHomeScreenDockIcons gets the homeScreenDockIcons property value. A list of app and folders to appear on the Home Screen Dock. This collection can contain a maximum of 500 elements.
// returns a []IosHomeScreenItemable when successful
func (m *IosDeviceFeaturesConfiguration) GetHomeScreenDockIcons()([]IosHomeScreenItemable) {
    val, err := m.GetBackingStore().Get("homeScreenDockIcons")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IosHomeScreenItemable)
    }
    return nil
}
// GetHomeScreenGridHeight gets the homeScreenGridHeight property value. Gets or sets the number of rows to render when configuring iOS home screen layout settings. If this value is configured, homeScreenGridWidth must be configured as well.
// returns a *int32 when successful
func (m *IosDeviceFeaturesConfiguration) GetHomeScreenGridHeight()(*int32) {
    val, err := m.GetBackingStore().Get("homeScreenGridHeight")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetHomeScreenGridWidth gets the homeScreenGridWidth property value. Gets or sets the number of columns to render when configuring iOS home screen layout settings. If this value is configured, homeScreenGridHeight must be configured as well.
// returns a *int32 when successful
func (m *IosDeviceFeaturesConfiguration) GetHomeScreenGridWidth()(*int32) {
    val, err := m.GetBackingStore().Get("homeScreenGridWidth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetHomeScreenPages gets the homeScreenPages property value. A list of pages on the Home Screen. This collection can contain a maximum of 500 elements.
// returns a []IosHomeScreenPageable when successful
func (m *IosDeviceFeaturesConfiguration) GetHomeScreenPages()([]IosHomeScreenPageable) {
    val, err := m.GetBackingStore().Get("homeScreenPages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IosHomeScreenPageable)
    }
    return nil
}
// GetIdentityCertificateForClientAuthentication gets the identityCertificateForClientAuthentication property value. Identity Certificate for the renewal of Kerberos ticket used in single sign-on settings.
// returns a IosCertificateProfileBaseable when successful
func (m *IosDeviceFeaturesConfiguration) GetIdentityCertificateForClientAuthentication()(IosCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("identityCertificateForClientAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosCertificateProfileBaseable)
    }
    return nil
}
// GetIosSingleSignOnExtension gets the iosSingleSignOnExtension property value. Gets or sets a single sign-on extension profile.
// returns a IosSingleSignOnExtensionable when successful
func (m *IosDeviceFeaturesConfiguration) GetIosSingleSignOnExtension()(IosSingleSignOnExtensionable) {
    val, err := m.GetBackingStore().Get("iosSingleSignOnExtension")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosSingleSignOnExtensionable)
    }
    return nil
}
// GetLockScreenFootnote gets the lockScreenFootnote property value. A footnote displayed on the login window and lock screen. Available in iOS 9.3.1 and later.
// returns a *string when successful
func (m *IosDeviceFeaturesConfiguration) GetLockScreenFootnote()(*string) {
    val, err := m.GetBackingStore().Get("lockScreenFootnote")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNotificationSettings gets the notificationSettings property value. Notification settings for each bundle id. Applicable to devices in supervised mode only (iOS 9.3 and later). This collection can contain a maximum of 500 elements.
// returns a []IosNotificationSettingsable when successful
func (m *IosDeviceFeaturesConfiguration) GetNotificationSettings()([]IosNotificationSettingsable) {
    val, err := m.GetBackingStore().Get("notificationSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IosNotificationSettingsable)
    }
    return nil
}
// GetSingleSignOnExtension gets the singleSignOnExtension property value. Gets or sets a single sign-on extension profile. Deprecated: use IOSSingleSignOnExtension instead.
// returns a SingleSignOnExtensionable when successful
func (m *IosDeviceFeaturesConfiguration) GetSingleSignOnExtension()(SingleSignOnExtensionable) {
    val, err := m.GetBackingStore().Get("singleSignOnExtension")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SingleSignOnExtensionable)
    }
    return nil
}
// GetSingleSignOnExtensionPkinitCertificate gets the singleSignOnExtensionPkinitCertificate property value. PKINIT Certificate for the authentication with single sign-on extension settings.
// returns a IosCertificateProfileBaseable when successful
func (m *IosDeviceFeaturesConfiguration) GetSingleSignOnExtensionPkinitCertificate()(IosCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("singleSignOnExtensionPkinitCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosCertificateProfileBaseable)
    }
    return nil
}
// GetSingleSignOnSettings gets the singleSignOnSettings property value. The Kerberos login settings that enable apps on receiving devices to authenticate smoothly.
// returns a IosSingleSignOnSettingsable when successful
func (m *IosDeviceFeaturesConfiguration) GetSingleSignOnSettings()(IosSingleSignOnSettingsable) {
    val, err := m.GetBackingStore().Get("singleSignOnSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosSingleSignOnSettingsable)
    }
    return nil
}
// GetWallpaperDisplayLocation gets the wallpaperDisplayLocation property value. An enum type for wallpaper display location specifier.
// returns a *IosWallpaperDisplayLocation when successful
func (m *IosDeviceFeaturesConfiguration) GetWallpaperDisplayLocation()(*IosWallpaperDisplayLocation) {
    val, err := m.GetBackingStore().Get("wallpaperDisplayLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IosWallpaperDisplayLocation)
    }
    return nil
}
// GetWallpaperImage gets the wallpaperImage property value. A wallpaper image must be in either PNG or JPEG format. It requires a supervised device with iOS 8 or later version.
// returns a MimeContentable when successful
func (m *IosDeviceFeaturesConfiguration) GetWallpaperImage()(MimeContentable) {
    val, err := m.GetBackingStore().Get("wallpaperImage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MimeContentable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosDeviceFeaturesConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AppleDeviceFeaturesConfigurationBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assetTagTemplate", m.GetAssetTagTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("contentFilterSettings", m.GetContentFilterSettings())
        if err != nil {
            return err
        }
    }
    if m.GetHomeScreenDockIcons() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHomeScreenDockIcons()))
        for i, v := range m.GetHomeScreenDockIcons() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("homeScreenDockIcons", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("homeScreenGridHeight", m.GetHomeScreenGridHeight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("homeScreenGridWidth", m.GetHomeScreenGridWidth())
        if err != nil {
            return err
        }
    }
    if m.GetHomeScreenPages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHomeScreenPages()))
        for i, v := range m.GetHomeScreenPages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("homeScreenPages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identityCertificateForClientAuthentication", m.GetIdentityCertificateForClientAuthentication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("iosSingleSignOnExtension", m.GetIosSingleSignOnExtension())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lockScreenFootnote", m.GetLockScreenFootnote())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotificationSettings()))
        for i, v := range m.GetNotificationSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("notificationSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("singleSignOnExtension", m.GetSingleSignOnExtension())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("singleSignOnExtensionPkinitCertificate", m.GetSingleSignOnExtensionPkinitCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("singleSignOnSettings", m.GetSingleSignOnSettings())
        if err != nil {
            return err
        }
    }
    if m.GetWallpaperDisplayLocation() != nil {
        cast := (*m.GetWallpaperDisplayLocation()).String()
        err = writer.WriteStringValue("wallpaperDisplayLocation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("wallpaperImage", m.GetWallpaperImage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssetTagTemplate sets the assetTagTemplate property value. Asset tag information for the device, displayed on the login window and lock screen.
func (m *IosDeviceFeaturesConfiguration) SetAssetTagTemplate(value *string)() {
    err := m.GetBackingStore().Set("assetTagTemplate", value)
    if err != nil {
        panic(err)
    }
}
// SetContentFilterSettings sets the contentFilterSettings property value. Gets or sets iOS Web Content Filter settings, supervised mode only
func (m *IosDeviceFeaturesConfiguration) SetContentFilterSettings(value IosWebContentFilterBaseable)() {
    err := m.GetBackingStore().Set("contentFilterSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetHomeScreenDockIcons sets the homeScreenDockIcons property value. A list of app and folders to appear on the Home Screen Dock. This collection can contain a maximum of 500 elements.
func (m *IosDeviceFeaturesConfiguration) SetHomeScreenDockIcons(value []IosHomeScreenItemable)() {
    err := m.GetBackingStore().Set("homeScreenDockIcons", value)
    if err != nil {
        panic(err)
    }
}
// SetHomeScreenGridHeight sets the homeScreenGridHeight property value. Gets or sets the number of rows to render when configuring iOS home screen layout settings. If this value is configured, homeScreenGridWidth must be configured as well.
func (m *IosDeviceFeaturesConfiguration) SetHomeScreenGridHeight(value *int32)() {
    err := m.GetBackingStore().Set("homeScreenGridHeight", value)
    if err != nil {
        panic(err)
    }
}
// SetHomeScreenGridWidth sets the homeScreenGridWidth property value. Gets or sets the number of columns to render when configuring iOS home screen layout settings. If this value is configured, homeScreenGridHeight must be configured as well.
func (m *IosDeviceFeaturesConfiguration) SetHomeScreenGridWidth(value *int32)() {
    err := m.GetBackingStore().Set("homeScreenGridWidth", value)
    if err != nil {
        panic(err)
    }
}
// SetHomeScreenPages sets the homeScreenPages property value. A list of pages on the Home Screen. This collection can contain a maximum of 500 elements.
func (m *IosDeviceFeaturesConfiguration) SetHomeScreenPages(value []IosHomeScreenPageable)() {
    err := m.GetBackingStore().Set("homeScreenPages", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityCertificateForClientAuthentication sets the identityCertificateForClientAuthentication property value. Identity Certificate for the renewal of Kerberos ticket used in single sign-on settings.
func (m *IosDeviceFeaturesConfiguration) SetIdentityCertificateForClientAuthentication(value IosCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("identityCertificateForClientAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetIosSingleSignOnExtension sets the iosSingleSignOnExtension property value. Gets or sets a single sign-on extension profile.
func (m *IosDeviceFeaturesConfiguration) SetIosSingleSignOnExtension(value IosSingleSignOnExtensionable)() {
    err := m.GetBackingStore().Set("iosSingleSignOnExtension", value)
    if err != nil {
        panic(err)
    }
}
// SetLockScreenFootnote sets the lockScreenFootnote property value. A footnote displayed on the login window and lock screen. Available in iOS 9.3.1 and later.
func (m *IosDeviceFeaturesConfiguration) SetLockScreenFootnote(value *string)() {
    err := m.GetBackingStore().Set("lockScreenFootnote", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationSettings sets the notificationSettings property value. Notification settings for each bundle id. Applicable to devices in supervised mode only (iOS 9.3 and later). This collection can contain a maximum of 500 elements.
func (m *IosDeviceFeaturesConfiguration) SetNotificationSettings(value []IosNotificationSettingsable)() {
    err := m.GetBackingStore().Set("notificationSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetSingleSignOnExtension sets the singleSignOnExtension property value. Gets or sets a single sign-on extension profile. Deprecated: use IOSSingleSignOnExtension instead.
func (m *IosDeviceFeaturesConfiguration) SetSingleSignOnExtension(value SingleSignOnExtensionable)() {
    err := m.GetBackingStore().Set("singleSignOnExtension", value)
    if err != nil {
        panic(err)
    }
}
// SetSingleSignOnExtensionPkinitCertificate sets the singleSignOnExtensionPkinitCertificate property value. PKINIT Certificate for the authentication with single sign-on extension settings.
func (m *IosDeviceFeaturesConfiguration) SetSingleSignOnExtensionPkinitCertificate(value IosCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("singleSignOnExtensionPkinitCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetSingleSignOnSettings sets the singleSignOnSettings property value. The Kerberos login settings that enable apps on receiving devices to authenticate smoothly.
func (m *IosDeviceFeaturesConfiguration) SetSingleSignOnSettings(value IosSingleSignOnSettingsable)() {
    err := m.GetBackingStore().Set("singleSignOnSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetWallpaperDisplayLocation sets the wallpaperDisplayLocation property value. An enum type for wallpaper display location specifier.
func (m *IosDeviceFeaturesConfiguration) SetWallpaperDisplayLocation(value *IosWallpaperDisplayLocation)() {
    err := m.GetBackingStore().Set("wallpaperDisplayLocation", value)
    if err != nil {
        panic(err)
    }
}
// SetWallpaperImage sets the wallpaperImage property value. A wallpaper image must be in either PNG or JPEG format. It requires a supervised device with iOS 8 or later version.
func (m *IosDeviceFeaturesConfiguration) SetWallpaperImage(value MimeContentable)() {
    err := m.GetBackingStore().Set("wallpaperImage", value)
    if err != nil {
        panic(err)
    }
}
type IosDeviceFeaturesConfigurationable interface {
    AppleDeviceFeaturesConfigurationBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssetTagTemplate()(*string)
    GetContentFilterSettings()(IosWebContentFilterBaseable)
    GetHomeScreenDockIcons()([]IosHomeScreenItemable)
    GetHomeScreenGridHeight()(*int32)
    GetHomeScreenGridWidth()(*int32)
    GetHomeScreenPages()([]IosHomeScreenPageable)
    GetIdentityCertificateForClientAuthentication()(IosCertificateProfileBaseable)
    GetIosSingleSignOnExtension()(IosSingleSignOnExtensionable)
    GetLockScreenFootnote()(*string)
    GetNotificationSettings()([]IosNotificationSettingsable)
    GetSingleSignOnExtension()(SingleSignOnExtensionable)
    GetSingleSignOnExtensionPkinitCertificate()(IosCertificateProfileBaseable)
    GetSingleSignOnSettings()(IosSingleSignOnSettingsable)
    GetWallpaperDisplayLocation()(*IosWallpaperDisplayLocation)
    GetWallpaperImage()(MimeContentable)
    SetAssetTagTemplate(value *string)()
    SetContentFilterSettings(value IosWebContentFilterBaseable)()
    SetHomeScreenDockIcons(value []IosHomeScreenItemable)()
    SetHomeScreenGridHeight(value *int32)()
    SetHomeScreenGridWidth(value *int32)()
    SetHomeScreenPages(value []IosHomeScreenPageable)()
    SetIdentityCertificateForClientAuthentication(value IosCertificateProfileBaseable)()
    SetIosSingleSignOnExtension(value IosSingleSignOnExtensionable)()
    SetLockScreenFootnote(value *string)()
    SetNotificationSettings(value []IosNotificationSettingsable)()
    SetSingleSignOnExtension(value SingleSignOnExtensionable)()
    SetSingleSignOnExtensionPkinitCertificate(value IosCertificateProfileBaseable)()
    SetSingleSignOnSettings(value IosSingleSignOnSettingsable)()
    SetWallpaperDisplayLocation(value *IosWallpaperDisplayLocation)()
    SetWallpaperImage(value MimeContentable)()
}
