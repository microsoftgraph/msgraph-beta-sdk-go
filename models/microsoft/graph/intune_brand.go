package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IntuneBrand struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Collection of blocked actions on the company portal as per platform and device ownership types.
    companyPortalBlockedActions []CompanyPortalBlockedAction;
    // Email address of the person/organization responsible for IT support.
    contactITEmailAddress *string;
    // Name of the person/organization responsible for IT support.
    contactITName *string;
    // Text comments regarding the person/organization responsible for IT support.
    contactITNotes *string;
    // Phone number of the person/organization responsible for IT support.
    contactITPhoneNumber *string;
    // The custom privacy message used to explain what the organization can see and do on managed devices.
    customCanSeePrivacyMessage *string;
    // The custom privacy message used to explain what the organization can’t see or do on managed devices.
    customCantSeePrivacyMessage *string;
    // The custom privacy message used to explain what the organization can’t see or do on managed devices.
    customPrivacyMessage *string;
    // Logo image displayed in Company Portal apps which have a dark background behind the logo.
    darkBackgroundLogo *MimeContent;
    // Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
    disableClientTelemetry *bool;
    // Company/organization name that is displayed to end users.
    displayName *string;
    // Customized device enrollment flow displayed to the end user . Possible values are: availableWithPrompts, availableWithoutPrompts, unavailable.
    enrollmentAvailability *EnrollmentAvailabilityOptions;
    // Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
    isFactoryResetDisabled *bool;
    // Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
    isRemoveDeviceDisabled *bool;
    // Customized image displayed in Company Portal app landing page
    landingPageCustomizedImage *MimeContent;
    // Logo image displayed in Company Portal apps which have a light background behind the logo.
    lightBackgroundLogo *MimeContent;
    // Display name of the company/organization’s IT helpdesk site.
    onlineSupportSiteName *string;
    // URL to the company/organization’s IT helpdesk site.
    onlineSupportSiteUrl *string;
    // URL to the company/organization’s privacy policy.
    privacyUrl *string;
    // List of scope tags assigned to the default branding profile
    roleScopeTagIds []string;
    // Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
    sendDeviceOwnershipChangePushNotification *bool;
    // Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
    showAzureADEnterpriseApps *bool;
    // Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
    showDisplayNameNextToLogo *bool;
    // Boolean that represents whether the administrator-supplied logo images are shown or not shown.
    showLogo *bool;
    // Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
    showNameNextToLogo *bool;
    // Boolean that indicates if Office WebApps will be shown in Company Portal
    showOfficeWebApps *bool;
    // Primary theme color used in the Company Portal applications and web portal.
    themeColor *RgbColor;
}
// Instantiates a new intuneBrand and sets the default values.
func NewIntuneBrand()(*IntuneBrand) {
    m := &IntuneBrand{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IntuneBrand) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrand) GetCompanyPortalBlockedActions()([]CompanyPortalBlockedAction) {
    if m == nil {
        return nil
    } else {
        return m.companyPortalBlockedActions
    }
}
// Gets the contactITEmailAddress property value. Email address of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITEmailAddress
    }
}
// Gets the contactITName property value. Name of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITName
    }
}
// Gets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITNotes
    }
}
// Gets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITPhoneNumber
    }
}
// Gets the customCanSeePrivacyMessage property value. The custom privacy message used to explain what the organization can see and do on managed devices.
func (m *IntuneBrand) GetCustomCanSeePrivacyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customCanSeePrivacyMessage
    }
}
// Gets the customCantSeePrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) GetCustomCantSeePrivacyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customCantSeePrivacyMessage
    }
}
// Gets the customPrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
func (m *IntuneBrand) GetCustomPrivacyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customPrivacyMessage
    }
}
// Gets the darkBackgroundLogo property value. Logo image displayed in Company Portal apps which have a dark background behind the logo.
func (m *IntuneBrand) GetDarkBackgroundLogo()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.darkBackgroundLogo
    }
}
// Gets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrand) GetDisableClientTelemetry()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableClientTelemetry
    }
}
// Gets the displayName property value. Company/organization name that is displayed to end users.
func (m *IntuneBrand) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the enrollmentAvailability property value. Customized device enrollment flow displayed to the end user . Possible values are: availableWithPrompts, availableWithoutPrompts, unavailable.
func (m *IntuneBrand) GetEnrollmentAvailability()(*EnrollmentAvailabilityOptions) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentAvailability
    }
}
// Gets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrand) GetIsFactoryResetDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFactoryResetDisabled
    }
}
// Gets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
func (m *IntuneBrand) GetIsRemoveDeviceDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRemoveDeviceDisabled
    }
}
// Gets the landingPageCustomizedImage property value. Customized image displayed in Company Portal app landing page
func (m *IntuneBrand) GetLandingPageCustomizedImage()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.landingPageCustomizedImage
    }
}
// Gets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo.
func (m *IntuneBrand) GetLightBackgroundLogo()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.lightBackgroundLogo
    }
}
// Gets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site.
func (m *IntuneBrand) GetOnlineSupportSiteName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineSupportSiteName
    }
}
// Gets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site.
func (m *IntuneBrand) GetOnlineSupportSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineSupportSiteUrl
    }
}
// Gets the privacyUrl property value. URL to the company/organization’s privacy policy.
func (m *IntuneBrand) GetPrivacyUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privacyUrl
    }
}
// Gets the roleScopeTagIds property value. List of scope tags assigned to the default branding profile
func (m *IntuneBrand) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
func (m *IntuneBrand) GetSendDeviceOwnershipChangePushNotification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sendDeviceOwnershipChangePushNotification
    }
}
// Gets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
func (m *IntuneBrand) GetShowAzureADEnterpriseApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showAzureADEnterpriseApps
    }
}
// Gets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) GetShowDisplayNameNextToLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showDisplayNameNextToLogo
    }
}
// Gets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not shown.
func (m *IntuneBrand) GetShowLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showLogo
    }
}
// Gets the showNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) GetShowNameNextToLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showNameNextToLogo
    }
}
// Gets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrand) GetShowOfficeWebApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showOfficeWebApps
    }
}
// Gets the themeColor property value. Primary theme color used in the Company Portal applications and web portal.
func (m *IntuneBrand) GetThemeColor()(*RgbColor) {
    if m == nil {
        return nil
    } else {
        return m.themeColor
    }
}
// The deserialization information for the current model
func (m *IntuneBrand) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["companyPortalBlockedActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCompanyPortalBlockedAction() })
        if err != nil {
            return err
        }
        res := make([]CompanyPortalBlockedAction, len(val))
        for i, v := range val {
            res[i] = *(v.(*CompanyPortalBlockedAction))
        }
        m.SetCompanyPortalBlockedActions(res)
        return nil
    }
    res["contactITEmailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContactITEmailAddress(val)
        return nil
    }
    res["contactITName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContactITName(val)
        return nil
    }
    res["contactITNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContactITNotes(val)
        return nil
    }
    res["contactITPhoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContactITPhoneNumber(val)
        return nil
    }
    res["customCanSeePrivacyMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomCanSeePrivacyMessage(val)
        return nil
    }
    res["customCantSeePrivacyMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomCantSeePrivacyMessage(val)
        return nil
    }
    res["customPrivacyMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomPrivacyMessage(val)
        return nil
    }
    res["darkBackgroundLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        m.SetDarkBackgroundLogo(val.(*MimeContent))
        return nil
    }
    res["disableClientTelemetry"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDisableClientTelemetry(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["enrollmentAvailability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentAvailabilityOptions)
        if err != nil {
            return err
        }
        cast := val.(EnrollmentAvailabilityOptions)
        m.SetEnrollmentAvailability(&cast)
        return nil
    }
    res["isFactoryResetDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsFactoryResetDisabled(val)
        return nil
    }
    res["isRemoveDeviceDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRemoveDeviceDisabled(val)
        return nil
    }
    res["landingPageCustomizedImage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        m.SetLandingPageCustomizedImage(val.(*MimeContent))
        return nil
    }
    res["lightBackgroundLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        m.SetLightBackgroundLogo(val.(*MimeContent))
        return nil
    }
    res["onlineSupportSiteName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnlineSupportSiteName(val)
        return nil
    }
    res["onlineSupportSiteUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnlineSupportSiteUrl(val)
        return nil
    }
    res["privacyUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrivacyUrl(val)
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetRoleScopeTagIds(res)
        return nil
    }
    res["sendDeviceOwnershipChangePushNotification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSendDeviceOwnershipChangePushNotification(val)
        return nil
    }
    res["showAzureADEnterpriseApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowAzureADEnterpriseApps(val)
        return nil
    }
    res["showDisplayNameNextToLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowDisplayNameNextToLogo(val)
        return nil
    }
    res["showLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowLogo(val)
        return nil
    }
    res["showNameNextToLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowNameNextToLogo(val)
        return nil
    }
    res["showOfficeWebApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowOfficeWebApps(val)
        return nil
    }
    res["themeColor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRgbColor() })
        if err != nil {
            return err
        }
        m.SetThemeColor(val.(*RgbColor))
        return nil
    }
    return res
}
func (m *IntuneBrand) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IntuneBrand) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCompanyPortalBlockedActions()))
        for i, v := range m.GetCompanyPortalBlockedActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentAvailability() != nil {
        cast := m.GetEnrollmentAvailability().String()
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
    {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *IntuneBrand) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
// Parameters:
//  - value : Value to set for the companyPortalBlockedActions property.
func (m *IntuneBrand) SetCompanyPortalBlockedActions(value []CompanyPortalBlockedAction)() {
    m.companyPortalBlockedActions = value
}
// Sets the contactITEmailAddress property value. Email address of the person/organization responsible for IT support.
// Parameters:
//  - value : Value to set for the contactITEmailAddress property.
func (m *IntuneBrand) SetContactITEmailAddress(value *string)() {
    m.contactITEmailAddress = value
}
// Sets the contactITName property value. Name of the person/organization responsible for IT support.
// Parameters:
//  - value : Value to set for the contactITName property.
func (m *IntuneBrand) SetContactITName(value *string)() {
    m.contactITName = value
}
// Sets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support.
// Parameters:
//  - value : Value to set for the contactITNotes property.
func (m *IntuneBrand) SetContactITNotes(value *string)() {
    m.contactITNotes = value
}
// Sets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support.
// Parameters:
//  - value : Value to set for the contactITPhoneNumber property.
func (m *IntuneBrand) SetContactITPhoneNumber(value *string)() {
    m.contactITPhoneNumber = value
}
// Sets the customCanSeePrivacyMessage property value. The custom privacy message used to explain what the organization can see and do on managed devices.
// Parameters:
//  - value : Value to set for the customCanSeePrivacyMessage property.
func (m *IntuneBrand) SetCustomCanSeePrivacyMessage(value *string)() {
    m.customCanSeePrivacyMessage = value
}
// Sets the customCantSeePrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
// Parameters:
//  - value : Value to set for the customCantSeePrivacyMessage property.
func (m *IntuneBrand) SetCustomCantSeePrivacyMessage(value *string)() {
    m.customCantSeePrivacyMessage = value
}
// Sets the customPrivacyMessage property value. The custom privacy message used to explain what the organization can’t see or do on managed devices.
// Parameters:
//  - value : Value to set for the customPrivacyMessage property.
func (m *IntuneBrand) SetCustomPrivacyMessage(value *string)() {
    m.customPrivacyMessage = value
}
// Sets the darkBackgroundLogo property value. Logo image displayed in Company Portal apps which have a dark background behind the logo.
// Parameters:
//  - value : Value to set for the darkBackgroundLogo property.
func (m *IntuneBrand) SetDarkBackgroundLogo(value *MimeContent)() {
    m.darkBackgroundLogo = value
}
// Sets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
// Parameters:
//  - value : Value to set for the disableClientTelemetry property.
func (m *IntuneBrand) SetDisableClientTelemetry(value *bool)() {
    m.disableClientTelemetry = value
}
// Sets the displayName property value. Company/organization name that is displayed to end users.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *IntuneBrand) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the enrollmentAvailability property value. Customized device enrollment flow displayed to the end user . Possible values are: availableWithPrompts, availableWithoutPrompts, unavailable.
// Parameters:
//  - value : Value to set for the enrollmentAvailability property.
func (m *IntuneBrand) SetEnrollmentAvailability(value *EnrollmentAvailabilityOptions)() {
    m.enrollmentAvailability = value
}
// Sets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
// Parameters:
//  - value : Value to set for the isFactoryResetDisabled property.
func (m *IntuneBrand) SetIsFactoryResetDisabled(value *bool)() {
    m.isFactoryResetDisabled = value
}
// Sets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
// Parameters:
//  - value : Value to set for the isRemoveDeviceDisabled property.
func (m *IntuneBrand) SetIsRemoveDeviceDisabled(value *bool)() {
    m.isRemoveDeviceDisabled = value
}
// Sets the landingPageCustomizedImage property value. Customized image displayed in Company Portal app landing page
// Parameters:
//  - value : Value to set for the landingPageCustomizedImage property.
func (m *IntuneBrand) SetLandingPageCustomizedImage(value *MimeContent)() {
    m.landingPageCustomizedImage = value
}
// Sets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo.
// Parameters:
//  - value : Value to set for the lightBackgroundLogo property.
func (m *IntuneBrand) SetLightBackgroundLogo(value *MimeContent)() {
    m.lightBackgroundLogo = value
}
// Sets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site.
// Parameters:
//  - value : Value to set for the onlineSupportSiteName property.
func (m *IntuneBrand) SetOnlineSupportSiteName(value *string)() {
    m.onlineSupportSiteName = value
}
// Sets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site.
// Parameters:
//  - value : Value to set for the onlineSupportSiteUrl property.
func (m *IntuneBrand) SetOnlineSupportSiteUrl(value *string)() {
    m.onlineSupportSiteUrl = value
}
// Sets the privacyUrl property value. URL to the company/organization’s privacy policy.
// Parameters:
//  - value : Value to set for the privacyUrl property.
func (m *IntuneBrand) SetPrivacyUrl(value *string)() {
    m.privacyUrl = value
}
// Sets the roleScopeTagIds property value. List of scope tags assigned to the default branding profile
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *IntuneBrand) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
// Parameters:
//  - value : Value to set for the sendDeviceOwnershipChangePushNotification property.
func (m *IntuneBrand) SetSendDeviceOwnershipChangePushNotification(value *bool)() {
    m.sendDeviceOwnershipChangePushNotification = value
}
// Sets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
// Parameters:
//  - value : Value to set for the showAzureADEnterpriseApps property.
func (m *IntuneBrand) SetShowAzureADEnterpriseApps(value *bool)() {
    m.showAzureADEnterpriseApps = value
}
// Sets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
// Parameters:
//  - value : Value to set for the showDisplayNameNextToLogo property.
func (m *IntuneBrand) SetShowDisplayNameNextToLogo(value *bool)() {
    m.showDisplayNameNextToLogo = value
}
// Sets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not shown.
// Parameters:
//  - value : Value to set for the showLogo property.
func (m *IntuneBrand) SetShowLogo(value *bool)() {
    m.showLogo = value
}
// Sets the showNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
// Parameters:
//  - value : Value to set for the showNameNextToLogo property.
func (m *IntuneBrand) SetShowNameNextToLogo(value *bool)() {
    m.showNameNextToLogo = value
}
// Sets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
// Parameters:
//  - value : Value to set for the showOfficeWebApps property.
func (m *IntuneBrand) SetShowOfficeWebApps(value *bool)() {
    m.showOfficeWebApps = value
}
// Sets the themeColor property value. Primary theme color used in the Company Portal applications and web portal.
// Parameters:
//  - value : Value to set for the themeColor property.
func (m *IntuneBrand) SetThemeColor(value *RgbColor)() {
    m.themeColor = value
}
