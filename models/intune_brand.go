package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntuneBrand intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
type IntuneBrand struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Collection of blocked actions on the company portal as per platform and device ownership types.
    companyPortalBlockedActions []CompanyPortalBlockedActionable
    // Email address of the person/organization responsible for IT support.
    contactITEmailAddress *string
    // Name of the person/organization responsible for IT support.
    contactITName *string
    // Text comments regarding the person/organization responsible for IT support.
    contactITNotes *string
    // Phone number of the person/organization responsible for IT support.
    contactITPhoneNumber *string
    // The custom privacy message used to explain what the organization can see and do on managed devices.
    customCanSeePrivacyMessage *string
    // The custom privacy message used to explain what the organization can’t see or do on managed devices.
    customCantSeePrivacyMessage *string
    // The custom privacy message used to explain what the organization can’t see or do on managed devices.
    customPrivacyMessage *string
    // Logo image displayed in Company Portal apps which have a dark background behind the logo.
    darkBackgroundLogo MimeContentable
    // Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
    disableClientTelemetry *bool
    // Company/organization name that is displayed to end users.
    displayName *string
    // Options available for enrollment flow customization
    enrollmentAvailability *EnrollmentAvailabilityOptions
    // Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
    isFactoryResetDisabled *bool
    // Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
    isRemoveDeviceDisabled *bool
    // Customized image displayed in Company Portal app landing page
    landingPageCustomizedImage MimeContentable
    // Logo image displayed in Company Portal apps which have a light background behind the logo.
    lightBackgroundLogo MimeContentable
    // The OdataType property
    odataType *string
    // Display name of the company/organization’s IT helpdesk site.
    onlineSupportSiteName *string
    // URL to the company/organization’s IT helpdesk site.
    onlineSupportSiteUrl *string
    // URL to the company/organization’s privacy policy.
    privacyUrl *string
    // List of scope tags assigned to the default branding profile
    roleScopeTagIds []string
    // Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
    sendDeviceOwnershipChangePushNotification *bool
    // Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
    showAzureADEnterpriseApps *bool
    // Boolean that indicates if ConfigurationManagerApps will be shown in Company Portal
    showConfigurationManagerApps *bool
    // Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
    showDisplayNameNextToLogo *bool
    // Boolean that represents whether the administrator-supplied logo images are shown or not shown.
    showLogo *bool
    // Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
    showNameNextToLogo *bool
    // Boolean that indicates if Office WebApps will be shown in Company Portal
    showOfficeWebApps *bool
    // Primary theme color used in the Company Portal applications and web portal.
    themeColor RgbColorable
}
// NewIntuneBrand instantiates a new intuneBrand and sets the default values.
func NewIntuneBrand()(*IntuneBrand) {
    m := &IntuneBrand{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIntuneBrandFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIntuneBrandFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntuneBrand(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IntuneBrand) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCompanyPortalBlockedActions gets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrand) GetCompanyPortalBlockedActions()([]CompanyPortalBlockedActionable) {
    return m.companyPortalBlockedActions
}
// GetContactITEmailAddress gets the contactITEmailAddress property value. Email address of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITEmailAddress()(*string) {
    return m.contactITEmailAddress
}
// GetContactITName gets the contactITName property value. Name of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITName()(*string) {
    return m.contactITName
}
// GetContactITNotes gets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITNotes()(*string) {
    return m.contactITNotes
}
// GetContactITPhoneNumber gets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITPhoneNumber()(*string) {
    return m.contactITPhoneNumber
}
// GetCustomCanSeePrivacyMessage gets the customCanSeePrivacyMessage property value. The custom privacy message used to explain what the organization can see and do on managed devices.
func (m *IntuneBrand) GetCustomCanSeePrivacyMessage()(*string) {
    return m.customCanSeePrivacyMessage
}
// GetCustomCantSeePrivacyMessage gets the customCantSeePrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) GetCustomCantSeePrivacyMessage()(*string) {
    return m.customCantSeePrivacyMessage
}
// GetCustomPrivacyMessage gets the customPrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) GetCustomPrivacyMessage()(*string) {
    return m.customPrivacyMessage
}
// GetDarkBackgroundLogo gets the darkBackgroundLogo property value. Logo image displayed in Company Portal apps which have a dark background behind the logo.
func (m *IntuneBrand) GetDarkBackgroundLogo()(MimeContentable) {
    return m.darkBackgroundLogo
}
// GetDisableClientTelemetry gets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrand) GetDisableClientTelemetry()(*bool) {
    return m.disableClientTelemetry
}
// GetDisplayName gets the displayName property value. Company/organization name that is displayed to end users.
func (m *IntuneBrand) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnrollmentAvailability gets the enrollmentAvailability property value. Options available for enrollment flow customization
func (m *IntuneBrand) GetEnrollmentAvailability()(*EnrollmentAvailabilityOptions) {
    return m.enrollmentAvailability
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IntuneBrand) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["companyPortalBlockedActions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCompanyPortalBlockedActionFromDiscriminatorValue , m.SetCompanyPortalBlockedActions)
    res["contactITEmailAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContactITEmailAddress)
    res["contactITName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContactITName)
    res["contactITNotes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContactITNotes)
    res["contactITPhoneNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContactITPhoneNumber)
    res["customCanSeePrivacyMessage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomCanSeePrivacyMessage)
    res["customCantSeePrivacyMessage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomCantSeePrivacyMessage)
    res["customPrivacyMessage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomPrivacyMessage)
    res["darkBackgroundLogo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMimeContentFromDiscriminatorValue , m.SetDarkBackgroundLogo)
    res["disableClientTelemetry"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDisableClientTelemetry)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["enrollmentAvailability"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEnrollmentAvailabilityOptions , m.SetEnrollmentAvailability)
    res["isFactoryResetDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsFactoryResetDisabled)
    res["isRemoveDeviceDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsRemoveDeviceDisabled)
    res["landingPageCustomizedImage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMimeContentFromDiscriminatorValue , m.SetLandingPageCustomizedImage)
    res["lightBackgroundLogo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMimeContentFromDiscriminatorValue , m.SetLightBackgroundLogo)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["onlineSupportSiteName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOnlineSupportSiteName)
    res["onlineSupportSiteUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOnlineSupportSiteUrl)
    res["privacyUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPrivacyUrl)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["sendDeviceOwnershipChangePushNotification"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSendDeviceOwnershipChangePushNotification)
    res["showAzureADEnterpriseApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetShowAzureADEnterpriseApps)
    res["showConfigurationManagerApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetShowConfigurationManagerApps)
    res["showDisplayNameNextToLogo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetShowDisplayNameNextToLogo)
    res["showLogo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetShowLogo)
    res["showNameNextToLogo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetShowNameNextToLogo)
    res["showOfficeWebApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetShowOfficeWebApps)
    res["themeColor"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateRgbColorFromDiscriminatorValue , m.SetThemeColor)
    return res
}
// GetIsFactoryResetDisabled gets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrand) GetIsFactoryResetDisabled()(*bool) {
    return m.isFactoryResetDisabled
}
// GetIsRemoveDeviceDisabled gets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
func (m *IntuneBrand) GetIsRemoveDeviceDisabled()(*bool) {
    return m.isRemoveDeviceDisabled
}
// GetLandingPageCustomizedImage gets the landingPageCustomizedImage property value. Customized image displayed in Company Portal app landing page
func (m *IntuneBrand) GetLandingPageCustomizedImage()(MimeContentable) {
    return m.landingPageCustomizedImage
}
// GetLightBackgroundLogo gets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo.
func (m *IntuneBrand) GetLightBackgroundLogo()(MimeContentable) {
    return m.lightBackgroundLogo
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IntuneBrand) GetOdataType()(*string) {
    return m.odataType
}
// GetOnlineSupportSiteName gets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site.
func (m *IntuneBrand) GetOnlineSupportSiteName()(*string) {
    return m.onlineSupportSiteName
}
// GetOnlineSupportSiteUrl gets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site.
func (m *IntuneBrand) GetOnlineSupportSiteUrl()(*string) {
    return m.onlineSupportSiteUrl
}
// GetPrivacyUrl gets the privacyUrl property value. URL to the company/organization’s privacy policy.
func (m *IntuneBrand) GetPrivacyUrl()(*string) {
    return m.privacyUrl
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of scope tags assigned to the default branding profile
func (m *IntuneBrand) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetSendDeviceOwnershipChangePushNotification gets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
func (m *IntuneBrand) GetSendDeviceOwnershipChangePushNotification()(*bool) {
    return m.sendDeviceOwnershipChangePushNotification
}
// GetShowAzureADEnterpriseApps gets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
func (m *IntuneBrand) GetShowAzureADEnterpriseApps()(*bool) {
    return m.showAzureADEnterpriseApps
}
// GetShowConfigurationManagerApps gets the showConfigurationManagerApps property value. Boolean that indicates if ConfigurationManagerApps will be shown in Company Portal
func (m *IntuneBrand) GetShowConfigurationManagerApps()(*bool) {
    return m.showConfigurationManagerApps
}
// GetShowDisplayNameNextToLogo gets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) GetShowDisplayNameNextToLogo()(*bool) {
    return m.showDisplayNameNextToLogo
}
// GetShowLogo gets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not shown.
func (m *IntuneBrand) GetShowLogo()(*bool) {
    return m.showLogo
}
// GetShowNameNextToLogo gets the showNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) GetShowNameNextToLogo()(*bool) {
    return m.showNameNextToLogo
}
// GetShowOfficeWebApps gets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrand) GetShowOfficeWebApps()(*bool) {
    return m.showOfficeWebApps
}
// GetThemeColor gets the themeColor property value. Primary theme color used in the Company Portal applications and web portal.
func (m *IntuneBrand) GetThemeColor()(RgbColorable) {
    return m.themeColor
}
// Serialize serializes information the current object
func (m *IntuneBrand) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCompanyPortalBlockedActions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCompanyPortalBlockedActions())
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
func (m *IntuneBrand) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCompanyPortalBlockedActions sets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrand) SetCompanyPortalBlockedActions(value []CompanyPortalBlockedActionable)() {
    m.companyPortalBlockedActions = value
}
// SetContactITEmailAddress sets the contactITEmailAddress property value. Email address of the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITEmailAddress(value *string)() {
    m.contactITEmailAddress = value
}
// SetContactITName sets the contactITName property value. Name of the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITName(value *string)() {
    m.contactITName = value
}
// SetContactITNotes sets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITNotes(value *string)() {
    m.contactITNotes = value
}
// SetContactITPhoneNumber sets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITPhoneNumber(value *string)() {
    m.contactITPhoneNumber = value
}
// SetCustomCanSeePrivacyMessage sets the customCanSeePrivacyMessage property value. The custom privacy message used to explain what the organization can see and do on managed devices.
func (m *IntuneBrand) SetCustomCanSeePrivacyMessage(value *string)() {
    m.customCanSeePrivacyMessage = value
}
// SetCustomCantSeePrivacyMessage sets the customCantSeePrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) SetCustomCantSeePrivacyMessage(value *string)() {
    m.customCantSeePrivacyMessage = value
}
// SetCustomPrivacyMessage sets the customPrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) SetCustomPrivacyMessage(value *string)() {
    m.customPrivacyMessage = value
}
// SetDarkBackgroundLogo sets the darkBackgroundLogo property value. Logo image displayed in Company Portal apps which have a dark background behind the logo.
func (m *IntuneBrand) SetDarkBackgroundLogo(value MimeContentable)() {
    m.darkBackgroundLogo = value
}
// SetDisableClientTelemetry sets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrand) SetDisableClientTelemetry(value *bool)() {
    m.disableClientTelemetry = value
}
// SetDisplayName sets the displayName property value. Company/organization name that is displayed to end users.
func (m *IntuneBrand) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnrollmentAvailability sets the enrollmentAvailability property value. Options available for enrollment flow customization
func (m *IntuneBrand) SetEnrollmentAvailability(value *EnrollmentAvailabilityOptions)() {
    m.enrollmentAvailability = value
}
// SetIsFactoryResetDisabled sets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrand) SetIsFactoryResetDisabled(value *bool)() {
    m.isFactoryResetDisabled = value
}
// SetIsRemoveDeviceDisabled sets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
func (m *IntuneBrand) SetIsRemoveDeviceDisabled(value *bool)() {
    m.isRemoveDeviceDisabled = value
}
// SetLandingPageCustomizedImage sets the landingPageCustomizedImage property value. Customized image displayed in Company Portal app landing page
func (m *IntuneBrand) SetLandingPageCustomizedImage(value MimeContentable)() {
    m.landingPageCustomizedImage = value
}
// SetLightBackgroundLogo sets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo.
func (m *IntuneBrand) SetLightBackgroundLogo(value MimeContentable)() {
    m.lightBackgroundLogo = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IntuneBrand) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOnlineSupportSiteName sets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site.
func (m *IntuneBrand) SetOnlineSupportSiteName(value *string)() {
    m.onlineSupportSiteName = value
}
// SetOnlineSupportSiteUrl sets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site.
func (m *IntuneBrand) SetOnlineSupportSiteUrl(value *string)() {
    m.onlineSupportSiteUrl = value
}
// SetPrivacyUrl sets the privacyUrl property value. URL to the company/organization’s privacy policy.
func (m *IntuneBrand) SetPrivacyUrl(value *string)() {
    m.privacyUrl = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of scope tags assigned to the default branding profile
func (m *IntuneBrand) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetSendDeviceOwnershipChangePushNotification sets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
func (m *IntuneBrand) SetSendDeviceOwnershipChangePushNotification(value *bool)() {
    m.sendDeviceOwnershipChangePushNotification = value
}
// SetShowAzureADEnterpriseApps sets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
func (m *IntuneBrand) SetShowAzureADEnterpriseApps(value *bool)() {
    m.showAzureADEnterpriseApps = value
}
// SetShowConfigurationManagerApps sets the showConfigurationManagerApps property value. Boolean that indicates if ConfigurationManagerApps will be shown in Company Portal
func (m *IntuneBrand) SetShowConfigurationManagerApps(value *bool)() {
    m.showConfigurationManagerApps = value
}
// SetShowDisplayNameNextToLogo sets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) SetShowDisplayNameNextToLogo(value *bool)() {
    m.showDisplayNameNextToLogo = value
}
// SetShowLogo sets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not shown.
func (m *IntuneBrand) SetShowLogo(value *bool)() {
    m.showLogo = value
}
// SetShowNameNextToLogo sets the showNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) SetShowNameNextToLogo(value *bool)() {
    m.showNameNextToLogo = value
}
// SetShowOfficeWebApps sets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrand) SetShowOfficeWebApps(value *bool)() {
    m.showOfficeWebApps = value
}
// SetThemeColor sets the themeColor property value. Primary theme color used in the Company Portal applications and web portal.
func (m *IntuneBrand) SetThemeColor(value RgbColorable)() {
    m.themeColor = value
}
