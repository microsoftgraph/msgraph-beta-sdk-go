package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IntuneBrandingProfile struct {
    Entity
    // The list of group assignments for the branding profile
    assignments []IntuneBrandingProfileAssignment;
    // Collection of blocked actions on the company portal as per platform and device ownership types.
    companyPortalBlockedActions []CompanyPortalBlockedAction;
    // E-mail address of the person/organization responsible for IT support
    contactITEmailAddress *string;
    // Name of the person/organization responsible for IT support
    contactITName *string;
    // Text comments regarding the person/organization responsible for IT support
    contactITNotes *string;
    // Phone number of the person/organization responsible for IT support
    contactITPhoneNumber *string;
    // Time when the BrandingProfile was created
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Text comments regarding what the admin has access to on the device
    customCanSeePrivacyMessage *string;
    // Text comments regarding what the admin doesn't have access to on the device
    customCantSeePrivacyMessage *string;
    // Text comments regarding what the admin doesn't have access to on the device
    customPrivacyMessage *string;
    // Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
    disableClientTelemetry *bool;
    // Company/organization name that is displayed to end users
    displayName *string;
    // Customized device enrollment flow displayed to the end user . Possible values are: availableWithPrompts, availableWithoutPrompts, unavailable.
    enrollmentAvailability *EnrollmentAvailabilityOptions;
    // Boolean that represents whether the profile is used as default or not
    isDefaultProfile *bool;
    // Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
    isFactoryResetDisabled *bool;
    // Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
    isRemoveDeviceDisabled *bool;
    // Customized image displayed in Company Portal apps landing page
    landingPageCustomizedImage *MimeContent;
    // Time when the BrandingProfile was last modified
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Logo image displayed in Company Portal apps which have a light background behind the logo
    lightBackgroundLogo *MimeContent;
    // Display name of the company/organization’s IT helpdesk site
    onlineSupportSiteName *string;
    // URL to the company/organization’s IT helpdesk site
    onlineSupportSiteUrl *string;
    // URL to the company/organization’s privacy policy
    privacyUrl *string;
    // Description of the profile
    profileDescription *string;
    // Name of the profile
    profileName *string;
    // List of scope tags assigned to the branding profile
    roleScopeTagIds []string;
    // Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
    sendDeviceOwnershipChangePushNotification *bool;
    // Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
    showAzureADEnterpriseApps *bool;
    // Boolean that represents whether the administrator-supplied display name will be shown next to the logo image or not
    showDisplayNameNextToLogo *bool;
    // Boolean that represents whether the administrator-supplied logo images are shown or not
    showLogo *bool;
    // Boolean that indicates if Office WebApps will be shown in Company Portal
    showOfficeWebApps *bool;
    // Primary theme color used in the Company Portal applications and web portal
    themeColor *RgbColor;
    // Logo image displayed in Company Portal apps which have a theme color background behind the logo
    themeColorLogo *MimeContent;
}
// Instantiates a new intuneBrandingProfile and sets the default values.
func NewIntuneBrandingProfile()(*IntuneBrandingProfile) {
    m := &IntuneBrandingProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of group assignments for the branding profile
func (m *IntuneBrandingProfile) GetAssignments()([]IntuneBrandingProfileAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrandingProfile) GetCompanyPortalBlockedActions()([]CompanyPortalBlockedAction) {
    if m == nil {
        return nil
    } else {
        return m.companyPortalBlockedActions
    }
}
// Gets the contactITEmailAddress property value. E-mail address of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITEmailAddress
    }
}
// Gets the contactITName property value. Name of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITName
    }
}
// Gets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITNotes
    }
}
// Gets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITPhoneNumber
    }
}
// Gets the createdDateTime property value. Time when the BrandingProfile was created
func (m *IntuneBrandingProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the customCanSeePrivacyMessage property value. Text comments regarding what the admin has access to on the device
func (m *IntuneBrandingProfile) GetCustomCanSeePrivacyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customCanSeePrivacyMessage
    }
}
// Gets the customCantSeePrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) GetCustomCantSeePrivacyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customCantSeePrivacyMessage
    }
}
// Gets the customPrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) GetCustomPrivacyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customPrivacyMessage
    }
}
// Gets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrandingProfile) GetDisableClientTelemetry()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableClientTelemetry
    }
}
// Gets the displayName property value. Company/organization name that is displayed to end users
func (m *IntuneBrandingProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the enrollmentAvailability property value. Customized device enrollment flow displayed to the end user . Possible values are: availableWithPrompts, availableWithoutPrompts, unavailable.
func (m *IntuneBrandingProfile) GetEnrollmentAvailability()(*EnrollmentAvailabilityOptions) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentAvailability
    }
}
// Gets the isDefaultProfile property value. Boolean that represents whether the profile is used as default or not
func (m *IntuneBrandingProfile) GetIsDefaultProfile()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefaultProfile
    }
}
// Gets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrandingProfile) GetIsFactoryResetDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFactoryResetDisabled
    }
}
// Gets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
func (m *IntuneBrandingProfile) GetIsRemoveDeviceDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRemoveDeviceDisabled
    }
}
// Gets the landingPageCustomizedImage property value. Customized image displayed in Company Portal apps landing page
func (m *IntuneBrandingProfile) GetLandingPageCustomizedImage()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.landingPageCustomizedImage
    }
}
// Gets the lastModifiedDateTime property value. Time when the BrandingProfile was last modified
func (m *IntuneBrandingProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo
func (m *IntuneBrandingProfile) GetLightBackgroundLogo()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.lightBackgroundLogo
    }
}
// Gets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) GetOnlineSupportSiteName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineSupportSiteName
    }
}
// Gets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) GetOnlineSupportSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineSupportSiteUrl
    }
}
// Gets the privacyUrl property value. URL to the company/organization’s privacy policy
func (m *IntuneBrandingProfile) GetPrivacyUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privacyUrl
    }
}
// Gets the profileDescription property value. Description of the profile
func (m *IntuneBrandingProfile) GetProfileDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profileDescription
    }
}
// Gets the profileName property value. Name of the profile
func (m *IntuneBrandingProfile) GetProfileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profileName
    }
}
// Gets the roleScopeTagIds property value. List of scope tags assigned to the branding profile
func (m *IntuneBrandingProfile) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
func (m *IntuneBrandingProfile) GetSendDeviceOwnershipChangePushNotification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sendDeviceOwnershipChangePushNotification
    }
}
// Gets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
func (m *IntuneBrandingProfile) GetShowAzureADEnterpriseApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showAzureADEnterpriseApps
    }
}
// Gets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image or not
func (m *IntuneBrandingProfile) GetShowDisplayNameNextToLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showDisplayNameNextToLogo
    }
}
// Gets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not
func (m *IntuneBrandingProfile) GetShowLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showLogo
    }
}
// Gets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrandingProfile) GetShowOfficeWebApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showOfficeWebApps
    }
}
// Gets the themeColor property value. Primary theme color used in the Company Portal applications and web portal
func (m *IntuneBrandingProfile) GetThemeColor()(*RgbColor) {
    if m == nil {
        return nil
    } else {
        return m.themeColor
    }
}
// Gets the themeColorLogo property value. Logo image displayed in Company Portal apps which have a theme color background behind the logo
func (m *IntuneBrandingProfile) GetThemeColorLogo()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.themeColorLogo
    }
}
// The deserialization information for the current model
func (m *IntuneBrandingProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntuneBrandingProfileAssignment() })
        if err != nil {
            return err
        }
        res := make([]IntuneBrandingProfileAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*IntuneBrandingProfileAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
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
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
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
    res["isDefaultProfile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefaultProfile(val)
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
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
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
    res["profileDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProfileDescription(val)
        return nil
    }
    res["profileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProfileName(val)
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
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
    res["themeColorLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        m.SetThemeColorLogo(val.(*MimeContent))
        return nil
    }
    return res
}
func (m *IntuneBrandingProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IntuneBrandingProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCompanyPortalBlockedActions()))
        for i, v := range m.GetCompanyPortalBlockedActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("companyPortalBlockedActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contactITEmailAddress", m.GetContactITEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contactITName", m.GetContactITName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contactITNotes", m.GetContactITNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contactITPhoneNumber", m.GetContactITPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customCanSeePrivacyMessage", m.GetCustomCanSeePrivacyMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customCantSeePrivacyMessage", m.GetCustomCantSeePrivacyMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customPrivacyMessage", m.GetCustomPrivacyMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableClientTelemetry", m.GetDisableClientTelemetry())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentAvailability() != nil {
        cast := m.GetEnrollmentAvailability().String()
        err = writer.WriteStringValue("enrollmentAvailability", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultProfile", m.GetIsDefaultProfile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFactoryResetDisabled", m.GetIsFactoryResetDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRemoveDeviceDisabled", m.GetIsRemoveDeviceDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("landingPageCustomizedImage", m.GetLandingPageCustomizedImage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lightBackgroundLogo", m.GetLightBackgroundLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onlineSupportSiteName", m.GetOnlineSupportSiteName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onlineSupportSiteUrl", m.GetOnlineSupportSiteUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("privacyUrl", m.GetPrivacyUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profileDescription", m.GetProfileDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profileName", m.GetProfileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("sendDeviceOwnershipChangePushNotification", m.GetSendDeviceOwnershipChangePushNotification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showAzureADEnterpriseApps", m.GetShowAzureADEnterpriseApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showDisplayNameNextToLogo", m.GetShowDisplayNameNextToLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showLogo", m.GetShowLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showOfficeWebApps", m.GetShowOfficeWebApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("themeColor", m.GetThemeColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("themeColorLogo", m.GetThemeColorLogo())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. The list of group assignments for the branding profile
// Parameters:
//  - value : Value to set for the assignments property.
func (m *IntuneBrandingProfile) SetAssignments(value []IntuneBrandingProfileAssignment)() {
    m.assignments = value
}
// Sets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
// Parameters:
//  - value : Value to set for the companyPortalBlockedActions property.
func (m *IntuneBrandingProfile) SetCompanyPortalBlockedActions(value []CompanyPortalBlockedAction)() {
    m.companyPortalBlockedActions = value
}
// Sets the contactITEmailAddress property value. E-mail address of the person/organization responsible for IT support
// Parameters:
//  - value : Value to set for the contactITEmailAddress property.
func (m *IntuneBrandingProfile) SetContactITEmailAddress(value *string)() {
    m.contactITEmailAddress = value
}
// Sets the contactITName property value. Name of the person/organization responsible for IT support
// Parameters:
//  - value : Value to set for the contactITName property.
func (m *IntuneBrandingProfile) SetContactITName(value *string)() {
    m.contactITName = value
}
// Sets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support
// Parameters:
//  - value : Value to set for the contactITNotes property.
func (m *IntuneBrandingProfile) SetContactITNotes(value *string)() {
    m.contactITNotes = value
}
// Sets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support
// Parameters:
//  - value : Value to set for the contactITPhoneNumber property.
func (m *IntuneBrandingProfile) SetContactITPhoneNumber(value *string)() {
    m.contactITPhoneNumber = value
}
// Sets the createdDateTime property value. Time when the BrandingProfile was created
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *IntuneBrandingProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the customCanSeePrivacyMessage property value. Text comments regarding what the admin has access to on the device
// Parameters:
//  - value : Value to set for the customCanSeePrivacyMessage property.
func (m *IntuneBrandingProfile) SetCustomCanSeePrivacyMessage(value *string)() {
    m.customCanSeePrivacyMessage = value
}
// Sets the customCantSeePrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
// Parameters:
//  - value : Value to set for the customCantSeePrivacyMessage property.
func (m *IntuneBrandingProfile) SetCustomCantSeePrivacyMessage(value *string)() {
    m.customCantSeePrivacyMessage = value
}
// Sets the customPrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
// Parameters:
//  - value : Value to set for the customPrivacyMessage property.
func (m *IntuneBrandingProfile) SetCustomPrivacyMessage(value *string)() {
    m.customPrivacyMessage = value
}
// Sets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
// Parameters:
//  - value : Value to set for the disableClientTelemetry property.
func (m *IntuneBrandingProfile) SetDisableClientTelemetry(value *bool)() {
    m.disableClientTelemetry = value
}
// Sets the displayName property value. Company/organization name that is displayed to end users
// Parameters:
//  - value : Value to set for the displayName property.
func (m *IntuneBrandingProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the enrollmentAvailability property value. Customized device enrollment flow displayed to the end user . Possible values are: availableWithPrompts, availableWithoutPrompts, unavailable.
// Parameters:
//  - value : Value to set for the enrollmentAvailability property.
func (m *IntuneBrandingProfile) SetEnrollmentAvailability(value *EnrollmentAvailabilityOptions)() {
    m.enrollmentAvailability = value
}
// Sets the isDefaultProfile property value. Boolean that represents whether the profile is used as default or not
// Parameters:
//  - value : Value to set for the isDefaultProfile property.
func (m *IntuneBrandingProfile) SetIsDefaultProfile(value *bool)() {
    m.isDefaultProfile = value
}
// Sets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
// Parameters:
//  - value : Value to set for the isFactoryResetDisabled property.
func (m *IntuneBrandingProfile) SetIsFactoryResetDisabled(value *bool)() {
    m.isFactoryResetDisabled = value
}
// Sets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
// Parameters:
//  - value : Value to set for the isRemoveDeviceDisabled property.
func (m *IntuneBrandingProfile) SetIsRemoveDeviceDisabled(value *bool)() {
    m.isRemoveDeviceDisabled = value
}
// Sets the landingPageCustomizedImage property value. Customized image displayed in Company Portal apps landing page
// Parameters:
//  - value : Value to set for the landingPageCustomizedImage property.
func (m *IntuneBrandingProfile) SetLandingPageCustomizedImage(value *MimeContent)() {
    m.landingPageCustomizedImage = value
}
// Sets the lastModifiedDateTime property value. Time when the BrandingProfile was last modified
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *IntuneBrandingProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo
// Parameters:
//  - value : Value to set for the lightBackgroundLogo property.
func (m *IntuneBrandingProfile) SetLightBackgroundLogo(value *MimeContent)() {
    m.lightBackgroundLogo = value
}
// Sets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site
// Parameters:
//  - value : Value to set for the onlineSupportSiteName property.
func (m *IntuneBrandingProfile) SetOnlineSupportSiteName(value *string)() {
    m.onlineSupportSiteName = value
}
// Sets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site
// Parameters:
//  - value : Value to set for the onlineSupportSiteUrl property.
func (m *IntuneBrandingProfile) SetOnlineSupportSiteUrl(value *string)() {
    m.onlineSupportSiteUrl = value
}
// Sets the privacyUrl property value. URL to the company/organization’s privacy policy
// Parameters:
//  - value : Value to set for the privacyUrl property.
func (m *IntuneBrandingProfile) SetPrivacyUrl(value *string)() {
    m.privacyUrl = value
}
// Sets the profileDescription property value. Description of the profile
// Parameters:
//  - value : Value to set for the profileDescription property.
func (m *IntuneBrandingProfile) SetProfileDescription(value *string)() {
    m.profileDescription = value
}
// Sets the profileName property value. Name of the profile
// Parameters:
//  - value : Value to set for the profileName property.
func (m *IntuneBrandingProfile) SetProfileName(value *string)() {
    m.profileName = value
}
// Sets the roleScopeTagIds property value. List of scope tags assigned to the branding profile
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *IntuneBrandingProfile) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
// Parameters:
//  - value : Value to set for the sendDeviceOwnershipChangePushNotification property.
func (m *IntuneBrandingProfile) SetSendDeviceOwnershipChangePushNotification(value *bool)() {
    m.sendDeviceOwnershipChangePushNotification = value
}
// Sets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
// Parameters:
//  - value : Value to set for the showAzureADEnterpriseApps property.
func (m *IntuneBrandingProfile) SetShowAzureADEnterpriseApps(value *bool)() {
    m.showAzureADEnterpriseApps = value
}
// Sets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image or not
// Parameters:
//  - value : Value to set for the showDisplayNameNextToLogo property.
func (m *IntuneBrandingProfile) SetShowDisplayNameNextToLogo(value *bool)() {
    m.showDisplayNameNextToLogo = value
}
// Sets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not
// Parameters:
//  - value : Value to set for the showLogo property.
func (m *IntuneBrandingProfile) SetShowLogo(value *bool)() {
    m.showLogo = value
}
// Sets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
// Parameters:
//  - value : Value to set for the showOfficeWebApps property.
func (m *IntuneBrandingProfile) SetShowOfficeWebApps(value *bool)() {
    m.showOfficeWebApps = value
}
// Sets the themeColor property value. Primary theme color used in the Company Portal applications and web portal
// Parameters:
//  - value : Value to set for the themeColor property.
func (m *IntuneBrandingProfile) SetThemeColor(value *RgbColor)() {
    m.themeColor = value
}
// Sets the themeColorLogo property value. Logo image displayed in Company Portal apps which have a theme color background behind the logo
// Parameters:
//  - value : Value to set for the themeColorLogo property.
func (m *IntuneBrandingProfile) SetThemeColorLogo(value *MimeContent)() {
    m.themeColorLogo = value
}
