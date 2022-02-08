package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IntuneBrandingProfile 
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
// NewIntuneBrandingProfile instantiates a new intuneBrandingProfile and sets the default values.
func NewIntuneBrandingProfile()(*IntuneBrandingProfile) {
    m := &IntuneBrandingProfile{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. The list of group assignments for the branding profile
func (m *IntuneBrandingProfile) GetAssignments()([]IntuneBrandingProfileAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCompanyPortalBlockedActions gets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrandingProfile) GetCompanyPortalBlockedActions()([]CompanyPortalBlockedAction) {
    if m == nil {
        return nil
    } else {
        return m.companyPortalBlockedActions
    }
}
// GetContactITEmailAddress gets the contactITEmailAddress property value. E-mail address of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITEmailAddress
    }
}
// GetContactITName gets the contactITName property value. Name of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITName
    }
}
// GetContactITNotes gets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITNotes
    }
}
// GetContactITPhoneNumber gets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITPhoneNumber
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Time when the BrandingProfile was created
func (m *IntuneBrandingProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCustomCanSeePrivacyMessage gets the customCanSeePrivacyMessage property value. Text comments regarding what the admin has access to on the device
func (m *IntuneBrandingProfile) GetCustomCanSeePrivacyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customCanSeePrivacyMessage
    }
}
// GetCustomCantSeePrivacyMessage gets the customCantSeePrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) GetCustomCantSeePrivacyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customCantSeePrivacyMessage
    }
}
// GetCustomPrivacyMessage gets the customPrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) GetCustomPrivacyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customPrivacyMessage
    }
}
// GetDisableClientTelemetry gets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrandingProfile) GetDisableClientTelemetry()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableClientTelemetry
    }
}
// GetDisplayName gets the displayName property value. Company/organization name that is displayed to end users
func (m *IntuneBrandingProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEnrollmentAvailability gets the enrollmentAvailability property value. Customized device enrollment flow displayed to the end user . Possible values are: availableWithPrompts, availableWithoutPrompts, unavailable.
func (m *IntuneBrandingProfile) GetEnrollmentAvailability()(*EnrollmentAvailabilityOptions) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentAvailability
    }
}
// GetIsDefaultProfile gets the isDefaultProfile property value. Boolean that represents whether the profile is used as default or not
func (m *IntuneBrandingProfile) GetIsDefaultProfile()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefaultProfile
    }
}
// GetIsFactoryResetDisabled gets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrandingProfile) GetIsFactoryResetDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFactoryResetDisabled
    }
}
// GetIsRemoveDeviceDisabled gets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
func (m *IntuneBrandingProfile) GetIsRemoveDeviceDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRemoveDeviceDisabled
    }
}
// GetLandingPageCustomizedImage gets the landingPageCustomizedImage property value. Customized image displayed in Company Portal apps landing page
func (m *IntuneBrandingProfile) GetLandingPageCustomizedImage()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.landingPageCustomizedImage
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Time when the BrandingProfile was last modified
func (m *IntuneBrandingProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLightBackgroundLogo gets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo
func (m *IntuneBrandingProfile) GetLightBackgroundLogo()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.lightBackgroundLogo
    }
}
// GetOnlineSupportSiteName gets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) GetOnlineSupportSiteName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineSupportSiteName
    }
}
// GetOnlineSupportSiteUrl gets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) GetOnlineSupportSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineSupportSiteUrl
    }
}
// GetPrivacyUrl gets the privacyUrl property value. URL to the company/organization’s privacy policy
func (m *IntuneBrandingProfile) GetPrivacyUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privacyUrl
    }
}
// GetProfileDescription gets the profileDescription property value. Description of the profile
func (m *IntuneBrandingProfile) GetProfileDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profileDescription
    }
}
// GetProfileName gets the profileName property value. Name of the profile
func (m *IntuneBrandingProfile) GetProfileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profileName
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of scope tags assigned to the branding profile
func (m *IntuneBrandingProfile) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetSendDeviceOwnershipChangePushNotification gets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
func (m *IntuneBrandingProfile) GetSendDeviceOwnershipChangePushNotification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sendDeviceOwnershipChangePushNotification
    }
}
// GetShowAzureADEnterpriseApps gets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
func (m *IntuneBrandingProfile) GetShowAzureADEnterpriseApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showAzureADEnterpriseApps
    }
}
// GetShowDisplayNameNextToLogo gets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image or not
func (m *IntuneBrandingProfile) GetShowDisplayNameNextToLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showDisplayNameNextToLogo
    }
}
// GetShowLogo gets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not
func (m *IntuneBrandingProfile) GetShowLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showLogo
    }
}
// GetShowOfficeWebApps gets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrandingProfile) GetShowOfficeWebApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showOfficeWebApps
    }
}
// GetThemeColor gets the themeColor property value. Primary theme color used in the Company Portal applications and web portal
func (m *IntuneBrandingProfile) GetThemeColor()(*RgbColor) {
    if m == nil {
        return nil
    } else {
        return m.themeColor
    }
}
// GetThemeColorLogo gets the themeColorLogo property value. Logo image displayed in Company Portal apps which have a theme color background behind the logo
func (m *IntuneBrandingProfile) GetThemeColorLogo()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.themeColorLogo
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IntuneBrandingProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntuneBrandingProfileAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntuneBrandingProfileAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*IntuneBrandingProfileAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["companyPortalBlockedActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCompanyPortalBlockedAction() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CompanyPortalBlockedAction, len(val))
            for i, v := range val {
                res[i] = *(v.(*CompanyPortalBlockedAction))
            }
            m.SetCompanyPortalBlockedActions(res)
        }
        return nil
    }
    res["contactITEmailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITEmailAddress(val)
        }
        return nil
    }
    res["contactITName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITName(val)
        }
        return nil
    }
    res["contactITNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITNotes(val)
        }
        return nil
    }
    res["contactITPhoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITPhoneNumber(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["customCanSeePrivacyMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCanSeePrivacyMessage(val)
        }
        return nil
    }
    res["customCantSeePrivacyMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCantSeePrivacyMessage(val)
        }
        return nil
    }
    res["customPrivacyMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomPrivacyMessage(val)
        }
        return nil
    }
    res["disableClientTelemetry"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableClientTelemetry(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enrollmentAvailability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentAvailabilityOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentAvailability(val.(*EnrollmentAvailabilityOptions))
        }
        return nil
    }
    res["isDefaultProfile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultProfile(val)
        }
        return nil
    }
    res["isFactoryResetDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFactoryResetDisabled(val)
        }
        return nil
    }
    res["isRemoveDeviceDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRemoveDeviceDisabled(val)
        }
        return nil
    }
    res["landingPageCustomizedImage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLandingPageCustomizedImage(val.(*MimeContent))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["lightBackgroundLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLightBackgroundLogo(val.(*MimeContent))
        }
        return nil
    }
    res["onlineSupportSiteName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineSupportSiteName(val)
        }
        return nil
    }
    res["onlineSupportSiteUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineSupportSiteUrl(val)
        }
        return nil
    }
    res["privacyUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyUrl(val)
        }
        return nil
    }
    res["profileDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileDescription(val)
        }
        return nil
    }
    res["profileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileName(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["sendDeviceOwnershipChangePushNotification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendDeviceOwnershipChangePushNotification(val)
        }
        return nil
    }
    res["showAzureADEnterpriseApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowAzureADEnterpriseApps(val)
        }
        return nil
    }
    res["showDisplayNameNextToLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowDisplayNameNextToLogo(val)
        }
        return nil
    }
    res["showLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowLogo(val)
        }
        return nil
    }
    res["showOfficeWebApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowOfficeWebApps(val)
        }
        return nil
    }
    res["themeColor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRgbColor() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThemeColor(val.(*RgbColor))
        }
        return nil
    }
    res["themeColorLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThemeColorLogo(val.(*MimeContent))
        }
        return nil
    }
    return res
}
func (m *IntuneBrandingProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *IntuneBrandingProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
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
    if m.GetCompanyPortalBlockedActions() != nil {
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
        cast := (*m.GetEnrollmentAvailability()).String()
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
    if m.GetRoleScopeTagIds() != nil {
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
// SetAssignments sets the assignments property value. The list of group assignments for the branding profile
func (m *IntuneBrandingProfile) SetAssignments(value []IntuneBrandingProfileAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCompanyPortalBlockedActions sets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrandingProfile) SetCompanyPortalBlockedActions(value []CompanyPortalBlockedAction)() {
    if m != nil {
        m.companyPortalBlockedActions = value
    }
}
// SetContactITEmailAddress sets the contactITEmailAddress property value. E-mail address of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) SetContactITEmailAddress(value *string)() {
    if m != nil {
        m.contactITEmailAddress = value
    }
}
// SetContactITName sets the contactITName property value. Name of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) SetContactITName(value *string)() {
    if m != nil {
        m.contactITName = value
    }
}
// SetContactITNotes sets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support
func (m *IntuneBrandingProfile) SetContactITNotes(value *string)() {
    if m != nil {
        m.contactITNotes = value
    }
}
// SetContactITPhoneNumber sets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) SetContactITPhoneNumber(value *string)() {
    if m != nil {
        m.contactITPhoneNumber = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Time when the BrandingProfile was created
func (m *IntuneBrandingProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCustomCanSeePrivacyMessage sets the customCanSeePrivacyMessage property value. Text comments regarding what the admin has access to on the device
func (m *IntuneBrandingProfile) SetCustomCanSeePrivacyMessage(value *string)() {
    if m != nil {
        m.customCanSeePrivacyMessage = value
    }
}
// SetCustomCantSeePrivacyMessage sets the customCantSeePrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) SetCustomCantSeePrivacyMessage(value *string)() {
    if m != nil {
        m.customCantSeePrivacyMessage = value
    }
}
// SetCustomPrivacyMessage sets the customPrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) SetCustomPrivacyMessage(value *string)() {
    if m != nil {
        m.customPrivacyMessage = value
    }
}
// SetDisableClientTelemetry sets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrandingProfile) SetDisableClientTelemetry(value *bool)() {
    if m != nil {
        m.disableClientTelemetry = value
    }
}
// SetDisplayName sets the displayName property value. Company/organization name that is displayed to end users
func (m *IntuneBrandingProfile) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEnrollmentAvailability sets the enrollmentAvailability property value. Customized device enrollment flow displayed to the end user . Possible values are: availableWithPrompts, availableWithoutPrompts, unavailable.
func (m *IntuneBrandingProfile) SetEnrollmentAvailability(value *EnrollmentAvailabilityOptions)() {
    if m != nil {
        m.enrollmentAvailability = value
    }
}
// SetIsDefaultProfile sets the isDefaultProfile property value. Boolean that represents whether the profile is used as default or not
func (m *IntuneBrandingProfile) SetIsDefaultProfile(value *bool)() {
    if m != nil {
        m.isDefaultProfile = value
    }
}
// SetIsFactoryResetDisabled sets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrandingProfile) SetIsFactoryResetDisabled(value *bool)() {
    if m != nil {
        m.isFactoryResetDisabled = value
    }
}
// SetIsRemoveDeviceDisabled sets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
func (m *IntuneBrandingProfile) SetIsRemoveDeviceDisabled(value *bool)() {
    if m != nil {
        m.isRemoveDeviceDisabled = value
    }
}
// SetLandingPageCustomizedImage sets the landingPageCustomizedImage property value. Customized image displayed in Company Portal apps landing page
func (m *IntuneBrandingProfile) SetLandingPageCustomizedImage(value *MimeContent)() {
    if m != nil {
        m.landingPageCustomizedImage = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Time when the BrandingProfile was last modified
func (m *IntuneBrandingProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetLightBackgroundLogo sets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo
func (m *IntuneBrandingProfile) SetLightBackgroundLogo(value *MimeContent)() {
    if m != nil {
        m.lightBackgroundLogo = value
    }
}
// SetOnlineSupportSiteName sets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) SetOnlineSupportSiteName(value *string)() {
    if m != nil {
        m.onlineSupportSiteName = value
    }
}
// SetOnlineSupportSiteUrl sets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) SetOnlineSupportSiteUrl(value *string)() {
    if m != nil {
        m.onlineSupportSiteUrl = value
    }
}
// SetPrivacyUrl sets the privacyUrl property value. URL to the company/organization’s privacy policy
func (m *IntuneBrandingProfile) SetPrivacyUrl(value *string)() {
    if m != nil {
        m.privacyUrl = value
    }
}
// SetProfileDescription sets the profileDescription property value. Description of the profile
func (m *IntuneBrandingProfile) SetProfileDescription(value *string)() {
    if m != nil {
        m.profileDescription = value
    }
}
// SetProfileName sets the profileName property value. Name of the profile
func (m *IntuneBrandingProfile) SetProfileName(value *string)() {
    if m != nil {
        m.profileName = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of scope tags assigned to the branding profile
func (m *IntuneBrandingProfile) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetSendDeviceOwnershipChangePushNotification sets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
func (m *IntuneBrandingProfile) SetSendDeviceOwnershipChangePushNotification(value *bool)() {
    if m != nil {
        m.sendDeviceOwnershipChangePushNotification = value
    }
}
// SetShowAzureADEnterpriseApps sets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
func (m *IntuneBrandingProfile) SetShowAzureADEnterpriseApps(value *bool)() {
    if m != nil {
        m.showAzureADEnterpriseApps = value
    }
}
// SetShowDisplayNameNextToLogo sets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image or not
func (m *IntuneBrandingProfile) SetShowDisplayNameNextToLogo(value *bool)() {
    if m != nil {
        m.showDisplayNameNextToLogo = value
    }
}
// SetShowLogo sets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not
func (m *IntuneBrandingProfile) SetShowLogo(value *bool)() {
    if m != nil {
        m.showLogo = value
    }
}
// SetShowOfficeWebApps sets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrandingProfile) SetShowOfficeWebApps(value *bool)() {
    if m != nil {
        m.showOfficeWebApps = value
    }
}
// SetThemeColor sets the themeColor property value. Primary theme color used in the Company Portal applications and web portal
func (m *IntuneBrandingProfile) SetThemeColor(value *RgbColor)() {
    if m != nil {
        m.themeColor = value
    }
}
// SetThemeColorLogo sets the themeColorLogo property value. Logo image displayed in Company Portal apps which have a theme color background behind the logo
func (m *IntuneBrandingProfile) SetThemeColorLogo(value *MimeContent)() {
    if m != nil {
        m.themeColorLogo = value
    }
}
