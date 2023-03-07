package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// IntuneBrand intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
type IntuneBrand struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewIntuneBrand instantiates a new intuneBrand and sets the default values.
func NewIntuneBrand()(*IntuneBrand) {
    m := &IntuneBrand{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIntuneBrandFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIntuneBrandFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntuneBrand(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IntuneBrand) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *IntuneBrand) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCompanyPortalBlockedActions gets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrand) GetCompanyPortalBlockedActions()([]CompanyPortalBlockedActionable) {
    val, err := m.GetBackingStore().Get("companyPortalBlockedActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CompanyPortalBlockedActionable)
    }
    return nil
}
// GetContactITEmailAddress gets the contactITEmailAddress property value. Email address of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITEmailAddress()(*string) {
    val, err := m.GetBackingStore().Get("contactITEmailAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContactITName gets the contactITName property value. Name of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITName()(*string) {
    val, err := m.GetBackingStore().Get("contactITName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContactITNotes gets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITNotes()(*string) {
    val, err := m.GetBackingStore().Get("contactITNotes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContactITPhoneNumber gets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITPhoneNumber()(*string) {
    val, err := m.GetBackingStore().Get("contactITPhoneNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomCanSeePrivacyMessage gets the customCanSeePrivacyMessage property value. The custom privacy message used to explain what the organization can see and do on managed devices.
func (m *IntuneBrand) GetCustomCanSeePrivacyMessage()(*string) {
    val, err := m.GetBackingStore().Get("customCanSeePrivacyMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomCantSeePrivacyMessage gets the customCantSeePrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) GetCustomCantSeePrivacyMessage()(*string) {
    val, err := m.GetBackingStore().Get("customCantSeePrivacyMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomPrivacyMessage gets the customPrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) GetCustomPrivacyMessage()(*string) {
    val, err := m.GetBackingStore().Get("customPrivacyMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDarkBackgroundLogo gets the darkBackgroundLogo property value. Logo image displayed in Company Portal apps which have a dark background behind the logo.
func (m *IntuneBrand) GetDarkBackgroundLogo()(MimeContentable) {
    val, err := m.GetBackingStore().Get("darkBackgroundLogo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MimeContentable)
    }
    return nil
}
// GetDisableClientTelemetry gets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrand) GetDisableClientTelemetry()(*bool) {
    val, err := m.GetBackingStore().Get("disableClientTelemetry")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisableDeviceCategorySelection gets the disableDeviceCategorySelection property value. Boolean that indicates if Device Category Selection will be shown in Company Portal
func (m *IntuneBrand) GetDisableDeviceCategorySelection()(*bool) {
    val, err := m.GetBackingStore().Get("disableDeviceCategorySelection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Company/organization name that is displayed to end users.
func (m *IntuneBrand) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnrollmentAvailability gets the enrollmentAvailability property value. Options available for enrollment flow customization
func (m *IntuneBrand) GetEnrollmentAvailability()(*EnrollmentAvailabilityOptions) {
    val, err := m.GetBackingStore().Get("enrollmentAvailability")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EnrollmentAvailabilityOptions)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IntuneBrand) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["companyPortalBlockedActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCompanyPortalBlockedActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CompanyPortalBlockedActionable, len(val))
            for i, v := range val {
                res[i] = v.(CompanyPortalBlockedActionable)
            }
            m.SetCompanyPortalBlockedActions(res)
        }
        return nil
    }
    res["contactITEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITEmailAddress(val)
        }
        return nil
    }
    res["contactITName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITName(val)
        }
        return nil
    }
    res["contactITNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITNotes(val)
        }
        return nil
    }
    res["contactITPhoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITPhoneNumber(val)
        }
        return nil
    }
    res["customCanSeePrivacyMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCanSeePrivacyMessage(val)
        }
        return nil
    }
    res["customCantSeePrivacyMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCantSeePrivacyMessage(val)
        }
        return nil
    }
    res["customPrivacyMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomPrivacyMessage(val)
        }
        return nil
    }
    res["darkBackgroundLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDarkBackgroundLogo(val.(MimeContentable))
        }
        return nil
    }
    res["disableClientTelemetry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableClientTelemetry(val)
        }
        return nil
    }
    res["disableDeviceCategorySelection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableDeviceCategorySelection(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enrollmentAvailability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentAvailabilityOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentAvailability(val.(*EnrollmentAvailabilityOptions))
        }
        return nil
    }
    res["isFactoryResetDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFactoryResetDisabled(val)
        }
        return nil
    }
    res["isRemoveDeviceDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRemoveDeviceDisabled(val)
        }
        return nil
    }
    res["landingPageCustomizedImage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLandingPageCustomizedImage(val.(MimeContentable))
        }
        return nil
    }
    res["lightBackgroundLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLightBackgroundLogo(val.(MimeContentable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["onlineSupportSiteName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineSupportSiteName(val)
        }
        return nil
    }
    res["onlineSupportSiteUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineSupportSiteUrl(val)
        }
        return nil
    }
    res["privacyUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyUrl(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["sendDeviceOwnershipChangePushNotification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendDeviceOwnershipChangePushNotification(val)
        }
        return nil
    }
    res["showAzureADEnterpriseApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowAzureADEnterpriseApps(val)
        }
        return nil
    }
    res["showConfigurationManagerApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowConfigurationManagerApps(val)
        }
        return nil
    }
    res["showDisplayNameNextToLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowDisplayNameNextToLogo(val)
        }
        return nil
    }
    res["showLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowLogo(val)
        }
        return nil
    }
    res["showNameNextToLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowNameNextToLogo(val)
        }
        return nil
    }
    res["showOfficeWebApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowOfficeWebApps(val)
        }
        return nil
    }
    res["themeColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRgbColorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThemeColor(val.(RgbColorable))
        }
        return nil
    }
    return res
}
// GetIsFactoryResetDisabled gets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrand) GetIsFactoryResetDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("isFactoryResetDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsRemoveDeviceDisabled gets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
func (m *IntuneBrand) GetIsRemoveDeviceDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("isRemoveDeviceDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLandingPageCustomizedImage gets the landingPageCustomizedImage property value. Customized image displayed in Company Portal app landing page
func (m *IntuneBrand) GetLandingPageCustomizedImage()(MimeContentable) {
    val, err := m.GetBackingStore().Get("landingPageCustomizedImage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MimeContentable)
    }
    return nil
}
// GetLightBackgroundLogo gets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo.
func (m *IntuneBrand) GetLightBackgroundLogo()(MimeContentable) {
    val, err := m.GetBackingStore().Get("lightBackgroundLogo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MimeContentable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IntuneBrand) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOnlineSupportSiteName gets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site.
func (m *IntuneBrand) GetOnlineSupportSiteName()(*string) {
    val, err := m.GetBackingStore().Get("onlineSupportSiteName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOnlineSupportSiteUrl gets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site.
func (m *IntuneBrand) GetOnlineSupportSiteUrl()(*string) {
    val, err := m.GetBackingStore().Get("onlineSupportSiteUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrivacyUrl gets the privacyUrl property value. URL to the company/organization’s privacy policy.
func (m *IntuneBrand) GetPrivacyUrl()(*string) {
    val, err := m.GetBackingStore().Get("privacyUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of scope tags assigned to the default branding profile
func (m *IntuneBrand) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSendDeviceOwnershipChangePushNotification gets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
func (m *IntuneBrand) GetSendDeviceOwnershipChangePushNotification()(*bool) {
    val, err := m.GetBackingStore().Get("sendDeviceOwnershipChangePushNotification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShowAzureADEnterpriseApps gets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
func (m *IntuneBrand) GetShowAzureADEnterpriseApps()(*bool) {
    val, err := m.GetBackingStore().Get("showAzureADEnterpriseApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShowConfigurationManagerApps gets the showConfigurationManagerApps property value. Boolean that indicates if ConfigurationManagerApps will be shown in Company Portal
func (m *IntuneBrand) GetShowConfigurationManagerApps()(*bool) {
    val, err := m.GetBackingStore().Get("showConfigurationManagerApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShowDisplayNameNextToLogo gets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) GetShowDisplayNameNextToLogo()(*bool) {
    val, err := m.GetBackingStore().Get("showDisplayNameNextToLogo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShowLogo gets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not shown.
func (m *IntuneBrand) GetShowLogo()(*bool) {
    val, err := m.GetBackingStore().Get("showLogo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShowNameNextToLogo gets the showNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) GetShowNameNextToLogo()(*bool) {
    val, err := m.GetBackingStore().Get("showNameNextToLogo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShowOfficeWebApps gets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrand) GetShowOfficeWebApps()(*bool) {
    val, err := m.GetBackingStore().Get("showOfficeWebApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetThemeColor gets the themeColor property value. Primary theme color used in the Company Portal applications and web portal.
func (m *IntuneBrand) GetThemeColor()(RgbColorable) {
    val, err := m.GetBackingStore().Get("themeColor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RgbColorable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IntuneBrand) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCompanyPortalBlockedActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompanyPortalBlockedActions()))
        for i, v := range m.GetCompanyPortalBlockedActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("companyPortalBlockedActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contactITEmailAddress", m.GetContactITEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contactITName", m.GetContactITName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contactITNotes", m.GetContactITNotes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contactITPhoneNumber", m.GetContactITPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customCanSeePrivacyMessage", m.GetCustomCanSeePrivacyMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customCantSeePrivacyMessage", m.GetCustomCantSeePrivacyMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customPrivacyMessage", m.GetCustomPrivacyMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("darkBackgroundLogo", m.GetDarkBackgroundLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("disableClientTelemetry", m.GetDisableClientTelemetry())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("disableDeviceCategorySelection", m.GetDisableDeviceCategorySelection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentAvailability() != nil {
        cast := (*m.GetEnrollmentAvailability()).String()
        err := writer.WriteStringValue("enrollmentAvailability", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isFactoryResetDisabled", m.GetIsFactoryResetDisabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRemoveDeviceDisabled", m.GetIsRemoveDeviceDisabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("landingPageCustomizedImage", m.GetLandingPageCustomizedImage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lightBackgroundLogo", m.GetLightBackgroundLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("onlineSupportSiteName", m.GetOnlineSupportSiteName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("onlineSupportSiteUrl", m.GetOnlineSupportSiteUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("privacyUrl", m.GetPrivacyUrl())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err := writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sendDeviceOwnershipChangePushNotification", m.GetSendDeviceOwnershipChangePushNotification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showAzureADEnterpriseApps", m.GetShowAzureADEnterpriseApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showConfigurationManagerApps", m.GetShowConfigurationManagerApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showDisplayNameNextToLogo", m.GetShowDisplayNameNextToLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showLogo", m.GetShowLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showNameNextToLogo", m.GetShowNameNextToLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showOfficeWebApps", m.GetShowOfficeWebApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("themeColor", m.GetThemeColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IntuneBrand) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *IntuneBrand) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCompanyPortalBlockedActions sets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrand) SetCompanyPortalBlockedActions(value []CompanyPortalBlockedActionable)() {
    err := m.GetBackingStore().Set("companyPortalBlockedActions", value)
    if err != nil {
        panic(err)
    }
}
// SetContactITEmailAddress sets the contactITEmailAddress property value. Email address of the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITEmailAddress(value *string)() {
    err := m.GetBackingStore().Set("contactITEmailAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetContactITName sets the contactITName property value. Name of the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITName(value *string)() {
    err := m.GetBackingStore().Set("contactITName", value)
    if err != nil {
        panic(err)
    }
}
// SetContactITNotes sets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITNotes(value *string)() {
    err := m.GetBackingStore().Set("contactITNotes", value)
    if err != nil {
        panic(err)
    }
}
// SetContactITPhoneNumber sets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITPhoneNumber(value *string)() {
    err := m.GetBackingStore().Set("contactITPhoneNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomCanSeePrivacyMessage sets the customCanSeePrivacyMessage property value. The custom privacy message used to explain what the organization can see and do on managed devices.
func (m *IntuneBrand) SetCustomCanSeePrivacyMessage(value *string)() {
    err := m.GetBackingStore().Set("customCanSeePrivacyMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomCantSeePrivacyMessage sets the customCantSeePrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) SetCustomCantSeePrivacyMessage(value *string)() {
    err := m.GetBackingStore().Set("customCantSeePrivacyMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomPrivacyMessage sets the customPrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) SetCustomPrivacyMessage(value *string)() {
    err := m.GetBackingStore().Set("customPrivacyMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetDarkBackgroundLogo sets the darkBackgroundLogo property value. Logo image displayed in Company Portal apps which have a dark background behind the logo.
func (m *IntuneBrand) SetDarkBackgroundLogo(value MimeContentable)() {
    err := m.GetBackingStore().Set("darkBackgroundLogo", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableClientTelemetry sets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrand) SetDisableClientTelemetry(value *bool)() {
    err := m.GetBackingStore().Set("disableClientTelemetry", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableDeviceCategorySelection sets the disableDeviceCategorySelection property value. Boolean that indicates if Device Category Selection will be shown in Company Portal
func (m *IntuneBrand) SetDisableDeviceCategorySelection(value *bool)() {
    err := m.GetBackingStore().Set("disableDeviceCategorySelection", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Company/organization name that is displayed to end users.
func (m *IntuneBrand) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentAvailability sets the enrollmentAvailability property value. Options available for enrollment flow customization
func (m *IntuneBrand) SetEnrollmentAvailability(value *EnrollmentAvailabilityOptions)() {
    err := m.GetBackingStore().Set("enrollmentAvailability", value)
    if err != nil {
        panic(err)
    }
}
// SetIsFactoryResetDisabled sets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrand) SetIsFactoryResetDisabled(value *bool)() {
    err := m.GetBackingStore().Set("isFactoryResetDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRemoveDeviceDisabled sets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
func (m *IntuneBrand) SetIsRemoveDeviceDisabled(value *bool)() {
    err := m.GetBackingStore().Set("isRemoveDeviceDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetLandingPageCustomizedImage sets the landingPageCustomizedImage property value. Customized image displayed in Company Portal app landing page
func (m *IntuneBrand) SetLandingPageCustomizedImage(value MimeContentable)() {
    err := m.GetBackingStore().Set("landingPageCustomizedImage", value)
    if err != nil {
        panic(err)
    }
}
// SetLightBackgroundLogo sets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo.
func (m *IntuneBrand) SetLightBackgroundLogo(value MimeContentable)() {
    err := m.GetBackingStore().Set("lightBackgroundLogo", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IntuneBrand) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOnlineSupportSiteName sets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site.
func (m *IntuneBrand) SetOnlineSupportSiteName(value *string)() {
    err := m.GetBackingStore().Set("onlineSupportSiteName", value)
    if err != nil {
        panic(err)
    }
}
// SetOnlineSupportSiteUrl sets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site.
func (m *IntuneBrand) SetOnlineSupportSiteUrl(value *string)() {
    err := m.GetBackingStore().Set("onlineSupportSiteUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetPrivacyUrl sets the privacyUrl property value. URL to the company/organization’s privacy policy.
func (m *IntuneBrand) SetPrivacyUrl(value *string)() {
    err := m.GetBackingStore().Set("privacyUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of scope tags assigned to the default branding profile
func (m *IntuneBrand) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetSendDeviceOwnershipChangePushNotification sets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
func (m *IntuneBrand) SetSendDeviceOwnershipChangePushNotification(value *bool)() {
    err := m.GetBackingStore().Set("sendDeviceOwnershipChangePushNotification", value)
    if err != nil {
        panic(err)
    }
}
// SetShowAzureADEnterpriseApps sets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
func (m *IntuneBrand) SetShowAzureADEnterpriseApps(value *bool)() {
    err := m.GetBackingStore().Set("showAzureADEnterpriseApps", value)
    if err != nil {
        panic(err)
    }
}
// SetShowConfigurationManagerApps sets the showConfigurationManagerApps property value. Boolean that indicates if ConfigurationManagerApps will be shown in Company Portal
func (m *IntuneBrand) SetShowConfigurationManagerApps(value *bool)() {
    err := m.GetBackingStore().Set("showConfigurationManagerApps", value)
    if err != nil {
        panic(err)
    }
}
// SetShowDisplayNameNextToLogo sets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) SetShowDisplayNameNextToLogo(value *bool)() {
    err := m.GetBackingStore().Set("showDisplayNameNextToLogo", value)
    if err != nil {
        panic(err)
    }
}
// SetShowLogo sets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not shown.
func (m *IntuneBrand) SetShowLogo(value *bool)() {
    err := m.GetBackingStore().Set("showLogo", value)
    if err != nil {
        panic(err)
    }
}
// SetShowNameNextToLogo sets the showNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) SetShowNameNextToLogo(value *bool)() {
    err := m.GetBackingStore().Set("showNameNextToLogo", value)
    if err != nil {
        panic(err)
    }
}
// SetShowOfficeWebApps sets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrand) SetShowOfficeWebApps(value *bool)() {
    err := m.GetBackingStore().Set("showOfficeWebApps", value)
    if err != nil {
        panic(err)
    }
}
// SetThemeColor sets the themeColor property value. Primary theme color used in the Company Portal applications and web portal.
func (m *IntuneBrand) SetThemeColor(value RgbColorable)() {
    err := m.GetBackingStore().Set("themeColor", value)
    if err != nil {
        panic(err)
    }
}
// IntuneBrandable 
type IntuneBrandable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCompanyPortalBlockedActions()([]CompanyPortalBlockedActionable)
    GetContactITEmailAddress()(*string)
    GetContactITName()(*string)
    GetContactITNotes()(*string)
    GetContactITPhoneNumber()(*string)
    GetCustomCanSeePrivacyMessage()(*string)
    GetCustomCantSeePrivacyMessage()(*string)
    GetCustomPrivacyMessage()(*string)
    GetDarkBackgroundLogo()(MimeContentable)
    GetDisableClientTelemetry()(*bool)
    GetDisableDeviceCategorySelection()(*bool)
    GetDisplayName()(*string)
    GetEnrollmentAvailability()(*EnrollmentAvailabilityOptions)
    GetIsFactoryResetDisabled()(*bool)
    GetIsRemoveDeviceDisabled()(*bool)
    GetLandingPageCustomizedImage()(MimeContentable)
    GetLightBackgroundLogo()(MimeContentable)
    GetOdataType()(*string)
    GetOnlineSupportSiteName()(*string)
    GetOnlineSupportSiteUrl()(*string)
    GetPrivacyUrl()(*string)
    GetRoleScopeTagIds()([]string)
    GetSendDeviceOwnershipChangePushNotification()(*bool)
    GetShowAzureADEnterpriseApps()(*bool)
    GetShowConfigurationManagerApps()(*bool)
    GetShowDisplayNameNextToLogo()(*bool)
    GetShowLogo()(*bool)
    GetShowNameNextToLogo()(*bool)
    GetShowOfficeWebApps()(*bool)
    GetThemeColor()(RgbColorable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCompanyPortalBlockedActions(value []CompanyPortalBlockedActionable)()
    SetContactITEmailAddress(value *string)()
    SetContactITName(value *string)()
    SetContactITNotes(value *string)()
    SetContactITPhoneNumber(value *string)()
    SetCustomCanSeePrivacyMessage(value *string)()
    SetCustomCantSeePrivacyMessage(value *string)()
    SetCustomPrivacyMessage(value *string)()
    SetDarkBackgroundLogo(value MimeContentable)()
    SetDisableClientTelemetry(value *bool)()
    SetDisableDeviceCategorySelection(value *bool)()
    SetDisplayName(value *string)()
    SetEnrollmentAvailability(value *EnrollmentAvailabilityOptions)()
    SetIsFactoryResetDisabled(value *bool)()
    SetIsRemoveDeviceDisabled(value *bool)()
    SetLandingPageCustomizedImage(value MimeContentable)()
    SetLightBackgroundLogo(value MimeContentable)()
    SetOdataType(value *string)()
    SetOnlineSupportSiteName(value *string)()
    SetOnlineSupportSiteUrl(value *string)()
    SetPrivacyUrl(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetSendDeviceOwnershipChangePushNotification(value *bool)()
    SetShowAzureADEnterpriseApps(value *bool)()
    SetShowConfigurationManagerApps(value *bool)()
    SetShowDisplayNameNextToLogo(value *bool)()
    SetShowLogo(value *bool)()
    SetShowNameNextToLogo(value *bool)()
    SetShowOfficeWebApps(value *bool)()
    SetThemeColor(value RgbColorable)()
}
