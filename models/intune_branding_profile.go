package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntuneBrandingProfile 
type IntuneBrandingProfile struct {
    Entity
}
// NewIntuneBrandingProfile instantiates a new IntuneBrandingProfile and sets the default values.
func NewIntuneBrandingProfile()(*IntuneBrandingProfile) {
    m := &IntuneBrandingProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIntuneBrandingProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIntuneBrandingProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntuneBrandingProfile(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments for the branding profile
func (m *IntuneBrandingProfile) GetAssignments()([]IntuneBrandingProfileAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IntuneBrandingProfileAssignmentable)
    }
    return nil
}
// GetCompanyPortalBlockedActions gets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrandingProfile) GetCompanyPortalBlockedActions()([]CompanyPortalBlockedActionable) {
    val, err := m.GetBackingStore().Get("companyPortalBlockedActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CompanyPortalBlockedActionable)
    }
    return nil
}
// GetContactITEmailAddress gets the contactITEmailAddress property value. E-mail address of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITEmailAddress()(*string) {
    val, err := m.GetBackingStore().Get("contactITEmailAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContactITName gets the contactITName property value. Name of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITName()(*string) {
    val, err := m.GetBackingStore().Get("contactITName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContactITNotes gets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITNotes()(*string) {
    val, err := m.GetBackingStore().Get("contactITNotes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContactITPhoneNumber gets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) GetContactITPhoneNumber()(*string) {
    val, err := m.GetBackingStore().Get("contactITPhoneNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Time when the BrandingProfile was created
func (m *IntuneBrandingProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCustomCanSeePrivacyMessage gets the customCanSeePrivacyMessage property value. Text comments regarding what the admin has access to on the device
func (m *IntuneBrandingProfile) GetCustomCanSeePrivacyMessage()(*string) {
    val, err := m.GetBackingStore().Get("customCanSeePrivacyMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomCantSeePrivacyMessage gets the customCantSeePrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) GetCustomCantSeePrivacyMessage()(*string) {
    val, err := m.GetBackingStore().Get("customCantSeePrivacyMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomPrivacyMessage gets the customPrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) GetCustomPrivacyMessage()(*string) {
    val, err := m.GetBackingStore().Get("customPrivacyMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisableClientTelemetry gets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrandingProfile) GetDisableClientTelemetry()(*bool) {
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
func (m *IntuneBrandingProfile) GetDisableDeviceCategorySelection()(*bool) {
    val, err := m.GetBackingStore().Get("disableDeviceCategorySelection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Company/organization name that is displayed to end users
func (m *IntuneBrandingProfile) GetDisplayName()(*string) {
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
func (m *IntuneBrandingProfile) GetEnrollmentAvailability()(*EnrollmentAvailabilityOptions) {
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
func (m *IntuneBrandingProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIntuneBrandingProfileAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntuneBrandingProfileAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IntuneBrandingProfileAssignmentable)
                }
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["companyPortalBlockedActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCompanyPortalBlockedActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CompanyPortalBlockedActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CompanyPortalBlockedActionable)
                }
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
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["isDefaultProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultProfile(val)
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
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
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
    res["profileDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileDescription(val)
        }
        return nil
    }
    res["profileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileName(val)
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
    res["themeColorLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThemeColorLogo(val.(MimeContentable))
        }
        return nil
    }
    return res
}
// GetIsDefaultProfile gets the isDefaultProfile property value. Boolean that represents whether the profile is used as default or not
func (m *IntuneBrandingProfile) GetIsDefaultProfile()(*bool) {
    val, err := m.GetBackingStore().Get("isDefaultProfile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsFactoryResetDisabled gets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrandingProfile) GetIsFactoryResetDisabled()(*bool) {
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
func (m *IntuneBrandingProfile) GetIsRemoveDeviceDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("isRemoveDeviceDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLandingPageCustomizedImage gets the landingPageCustomizedImage property value. Customized image displayed in Company Portal apps landing page
func (m *IntuneBrandingProfile) GetLandingPageCustomizedImage()(MimeContentable) {
    val, err := m.GetBackingStore().Get("landingPageCustomizedImage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MimeContentable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Time when the BrandingProfile was last modified
func (m *IntuneBrandingProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLightBackgroundLogo gets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo
func (m *IntuneBrandingProfile) GetLightBackgroundLogo()(MimeContentable) {
    val, err := m.GetBackingStore().Get("lightBackgroundLogo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MimeContentable)
    }
    return nil
}
// GetOnlineSupportSiteName gets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) GetOnlineSupportSiteName()(*string) {
    val, err := m.GetBackingStore().Get("onlineSupportSiteName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOnlineSupportSiteUrl gets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) GetOnlineSupportSiteUrl()(*string) {
    val, err := m.GetBackingStore().Get("onlineSupportSiteUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrivacyUrl gets the privacyUrl property value. URL to the company/organization’s privacy policy
func (m *IntuneBrandingProfile) GetPrivacyUrl()(*string) {
    val, err := m.GetBackingStore().Get("privacyUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProfileDescription gets the profileDescription property value. Description of the profile
func (m *IntuneBrandingProfile) GetProfileDescription()(*string) {
    val, err := m.GetBackingStore().Get("profileDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProfileName gets the profileName property value. Name of the profile
func (m *IntuneBrandingProfile) GetProfileName()(*string) {
    val, err := m.GetBackingStore().Get("profileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of scope tags assigned to the branding profile
func (m *IntuneBrandingProfile) GetRoleScopeTagIds()([]string) {
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
func (m *IntuneBrandingProfile) GetSendDeviceOwnershipChangePushNotification()(*bool) {
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
func (m *IntuneBrandingProfile) GetShowAzureADEnterpriseApps()(*bool) {
    val, err := m.GetBackingStore().Get("showAzureADEnterpriseApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShowConfigurationManagerApps gets the showConfigurationManagerApps property value. Boolean that indicates if Configuration Manager Apps will be shown in Company Portal
func (m *IntuneBrandingProfile) GetShowConfigurationManagerApps()(*bool) {
    val, err := m.GetBackingStore().Get("showConfigurationManagerApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShowDisplayNameNextToLogo gets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image or not
func (m *IntuneBrandingProfile) GetShowDisplayNameNextToLogo()(*bool) {
    val, err := m.GetBackingStore().Get("showDisplayNameNextToLogo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShowLogo gets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not
func (m *IntuneBrandingProfile) GetShowLogo()(*bool) {
    val, err := m.GetBackingStore().Get("showLogo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShowOfficeWebApps gets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrandingProfile) GetShowOfficeWebApps()(*bool) {
    val, err := m.GetBackingStore().Get("showOfficeWebApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetThemeColor gets the themeColor property value. Primary theme color used in the Company Portal applications and web portal
func (m *IntuneBrandingProfile) GetThemeColor()(RgbColorable) {
    val, err := m.GetBackingStore().Get("themeColor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RgbColorable)
    }
    return nil
}
// GetThemeColorLogo gets the themeColorLogo property value. Logo image displayed in Company Portal apps which have a theme color background behind the logo
func (m *IntuneBrandingProfile) GetThemeColorLogo()(MimeContentable) {
    val, err := m.GetBackingStore().Get("themeColorLogo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MimeContentable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IntuneBrandingProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCompanyPortalBlockedActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompanyPortalBlockedActions()))
        for i, v := range m.GetCompanyPortalBlockedActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        err = writer.WriteBoolValue("disableDeviceCategorySelection", m.GetDisableDeviceCategorySelection())
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
        err = writer.WriteBoolValue("showConfigurationManagerApps", m.GetShowConfigurationManagerApps())
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
func (m *IntuneBrandingProfile) SetAssignments(value []IntuneBrandingProfileAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetCompanyPortalBlockedActions sets the companyPortalBlockedActions property value. Collection of blocked actions on the company portal as per platform and device ownership types.
func (m *IntuneBrandingProfile) SetCompanyPortalBlockedActions(value []CompanyPortalBlockedActionable)() {
    err := m.GetBackingStore().Set("companyPortalBlockedActions", value)
    if err != nil {
        panic(err)
    }
}
// SetContactITEmailAddress sets the contactITEmailAddress property value. E-mail address of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) SetContactITEmailAddress(value *string)() {
    err := m.GetBackingStore().Set("contactITEmailAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetContactITName sets the contactITName property value. Name of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) SetContactITName(value *string)() {
    err := m.GetBackingStore().Set("contactITName", value)
    if err != nil {
        panic(err)
    }
}
// SetContactITNotes sets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support
func (m *IntuneBrandingProfile) SetContactITNotes(value *string)() {
    err := m.GetBackingStore().Set("contactITNotes", value)
    if err != nil {
        panic(err)
    }
}
// SetContactITPhoneNumber sets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support
func (m *IntuneBrandingProfile) SetContactITPhoneNumber(value *string)() {
    err := m.GetBackingStore().Set("contactITPhoneNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Time when the BrandingProfile was created
func (m *IntuneBrandingProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomCanSeePrivacyMessage sets the customCanSeePrivacyMessage property value. Text comments regarding what the admin has access to on the device
func (m *IntuneBrandingProfile) SetCustomCanSeePrivacyMessage(value *string)() {
    err := m.GetBackingStore().Set("customCanSeePrivacyMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomCantSeePrivacyMessage sets the customCantSeePrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) SetCustomCantSeePrivacyMessage(value *string)() {
    err := m.GetBackingStore().Set("customCantSeePrivacyMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomPrivacyMessage sets the customPrivacyMessage property value. Text comments regarding what the admin doesn't have access to on the device
func (m *IntuneBrandingProfile) SetCustomPrivacyMessage(value *string)() {
    err := m.GetBackingStore().Set("customPrivacyMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableClientTelemetry sets the disableClientTelemetry property value. Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
func (m *IntuneBrandingProfile) SetDisableClientTelemetry(value *bool)() {
    err := m.GetBackingStore().Set("disableClientTelemetry", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableDeviceCategorySelection sets the disableDeviceCategorySelection property value. Boolean that indicates if Device Category Selection will be shown in Company Portal
func (m *IntuneBrandingProfile) SetDisableDeviceCategorySelection(value *bool)() {
    err := m.GetBackingStore().Set("disableDeviceCategorySelection", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Company/organization name that is displayed to end users
func (m *IntuneBrandingProfile) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentAvailability sets the enrollmentAvailability property value. Options available for enrollment flow customization
func (m *IntuneBrandingProfile) SetEnrollmentAvailability(value *EnrollmentAvailabilityOptions)() {
    err := m.GetBackingStore().Set("enrollmentAvailability", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDefaultProfile sets the isDefaultProfile property value. Boolean that represents whether the profile is used as default or not
func (m *IntuneBrandingProfile) SetIsDefaultProfile(value *bool)() {
    err := m.GetBackingStore().Set("isDefaultProfile", value)
    if err != nil {
        panic(err)
    }
}
// SetIsFactoryResetDisabled sets the isFactoryResetDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned devices.
func (m *IntuneBrandingProfile) SetIsFactoryResetDisabled(value *bool)() {
    err := m.GetBackingStore().Set("isFactoryResetDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRemoveDeviceDisabled sets the isRemoveDeviceDisabled property value. Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned devices.
func (m *IntuneBrandingProfile) SetIsRemoveDeviceDisabled(value *bool)() {
    err := m.GetBackingStore().Set("isRemoveDeviceDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetLandingPageCustomizedImage sets the landingPageCustomizedImage property value. Customized image displayed in Company Portal apps landing page
func (m *IntuneBrandingProfile) SetLandingPageCustomizedImage(value MimeContentable)() {
    err := m.GetBackingStore().Set("landingPageCustomizedImage", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Time when the BrandingProfile was last modified
func (m *IntuneBrandingProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLightBackgroundLogo sets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo
func (m *IntuneBrandingProfile) SetLightBackgroundLogo(value MimeContentable)() {
    err := m.GetBackingStore().Set("lightBackgroundLogo", value)
    if err != nil {
        panic(err)
    }
}
// SetOnlineSupportSiteName sets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) SetOnlineSupportSiteName(value *string)() {
    err := m.GetBackingStore().Set("onlineSupportSiteName", value)
    if err != nil {
        panic(err)
    }
}
// SetOnlineSupportSiteUrl sets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site
func (m *IntuneBrandingProfile) SetOnlineSupportSiteUrl(value *string)() {
    err := m.GetBackingStore().Set("onlineSupportSiteUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetPrivacyUrl sets the privacyUrl property value. URL to the company/organization’s privacy policy
func (m *IntuneBrandingProfile) SetPrivacyUrl(value *string)() {
    err := m.GetBackingStore().Set("privacyUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileDescription sets the profileDescription property value. Description of the profile
func (m *IntuneBrandingProfile) SetProfileDescription(value *string)() {
    err := m.GetBackingStore().Set("profileDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileName sets the profileName property value. Name of the profile
func (m *IntuneBrandingProfile) SetProfileName(value *string)() {
    err := m.GetBackingStore().Set("profileName", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of scope tags assigned to the branding profile
func (m *IntuneBrandingProfile) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetSendDeviceOwnershipChangePushNotification sets the sendDeviceOwnershipChangePushNotification property value. Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal to corporate
func (m *IntuneBrandingProfile) SetSendDeviceOwnershipChangePushNotification(value *bool)() {
    err := m.GetBackingStore().Set("sendDeviceOwnershipChangePushNotification", value)
    if err != nil {
        panic(err)
    }
}
// SetShowAzureADEnterpriseApps sets the showAzureADEnterpriseApps property value. Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
func (m *IntuneBrandingProfile) SetShowAzureADEnterpriseApps(value *bool)() {
    err := m.GetBackingStore().Set("showAzureADEnterpriseApps", value)
    if err != nil {
        panic(err)
    }
}
// SetShowConfigurationManagerApps sets the showConfigurationManagerApps property value. Boolean that indicates if Configuration Manager Apps will be shown in Company Portal
func (m *IntuneBrandingProfile) SetShowConfigurationManagerApps(value *bool)() {
    err := m.GetBackingStore().Set("showConfigurationManagerApps", value)
    if err != nil {
        panic(err)
    }
}
// SetShowDisplayNameNextToLogo sets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image or not
func (m *IntuneBrandingProfile) SetShowDisplayNameNextToLogo(value *bool)() {
    err := m.GetBackingStore().Set("showDisplayNameNextToLogo", value)
    if err != nil {
        panic(err)
    }
}
// SetShowLogo sets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not
func (m *IntuneBrandingProfile) SetShowLogo(value *bool)() {
    err := m.GetBackingStore().Set("showLogo", value)
    if err != nil {
        panic(err)
    }
}
// SetShowOfficeWebApps sets the showOfficeWebApps property value. Boolean that indicates if Office WebApps will be shown in Company Portal
func (m *IntuneBrandingProfile) SetShowOfficeWebApps(value *bool)() {
    err := m.GetBackingStore().Set("showOfficeWebApps", value)
    if err != nil {
        panic(err)
    }
}
// SetThemeColor sets the themeColor property value. Primary theme color used in the Company Portal applications and web portal
func (m *IntuneBrandingProfile) SetThemeColor(value RgbColorable)() {
    err := m.GetBackingStore().Set("themeColor", value)
    if err != nil {
        panic(err)
    }
}
// SetThemeColorLogo sets the themeColorLogo property value. Logo image displayed in Company Portal apps which have a theme color background behind the logo
func (m *IntuneBrandingProfile) SetThemeColorLogo(value MimeContentable)() {
    err := m.GetBackingStore().Set("themeColorLogo", value)
    if err != nil {
        panic(err)
    }
}
// IntuneBrandingProfileable 
type IntuneBrandingProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]IntuneBrandingProfileAssignmentable)
    GetCompanyPortalBlockedActions()([]CompanyPortalBlockedActionable)
    GetContactITEmailAddress()(*string)
    GetContactITName()(*string)
    GetContactITNotes()(*string)
    GetContactITPhoneNumber()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomCanSeePrivacyMessage()(*string)
    GetCustomCantSeePrivacyMessage()(*string)
    GetCustomPrivacyMessage()(*string)
    GetDisableClientTelemetry()(*bool)
    GetDisableDeviceCategorySelection()(*bool)
    GetDisplayName()(*string)
    GetEnrollmentAvailability()(*EnrollmentAvailabilityOptions)
    GetIsDefaultProfile()(*bool)
    GetIsFactoryResetDisabled()(*bool)
    GetIsRemoveDeviceDisabled()(*bool)
    GetLandingPageCustomizedImage()(MimeContentable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLightBackgroundLogo()(MimeContentable)
    GetOnlineSupportSiteName()(*string)
    GetOnlineSupportSiteUrl()(*string)
    GetPrivacyUrl()(*string)
    GetProfileDescription()(*string)
    GetProfileName()(*string)
    GetRoleScopeTagIds()([]string)
    GetSendDeviceOwnershipChangePushNotification()(*bool)
    GetShowAzureADEnterpriseApps()(*bool)
    GetShowConfigurationManagerApps()(*bool)
    GetShowDisplayNameNextToLogo()(*bool)
    GetShowLogo()(*bool)
    GetShowOfficeWebApps()(*bool)
    GetThemeColor()(RgbColorable)
    GetThemeColorLogo()(MimeContentable)
    SetAssignments(value []IntuneBrandingProfileAssignmentable)()
    SetCompanyPortalBlockedActions(value []CompanyPortalBlockedActionable)()
    SetContactITEmailAddress(value *string)()
    SetContactITName(value *string)()
    SetContactITNotes(value *string)()
    SetContactITPhoneNumber(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomCanSeePrivacyMessage(value *string)()
    SetCustomCantSeePrivacyMessage(value *string)()
    SetCustomPrivacyMessage(value *string)()
    SetDisableClientTelemetry(value *bool)()
    SetDisableDeviceCategorySelection(value *bool)()
    SetDisplayName(value *string)()
    SetEnrollmentAvailability(value *EnrollmentAvailabilityOptions)()
    SetIsDefaultProfile(value *bool)()
    SetIsFactoryResetDisabled(value *bool)()
    SetIsRemoveDeviceDisabled(value *bool)()
    SetLandingPageCustomizedImage(value MimeContentable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLightBackgroundLogo(value MimeContentable)()
    SetOnlineSupportSiteName(value *string)()
    SetOnlineSupportSiteUrl(value *string)()
    SetPrivacyUrl(value *string)()
    SetProfileDescription(value *string)()
    SetProfileName(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetSendDeviceOwnershipChangePushNotification(value *bool)()
    SetShowAzureADEnterpriseApps(value *bool)()
    SetShowConfigurationManagerApps(value *bool)()
    SetShowDisplayNameNextToLogo(value *bool)()
    SetShowLogo(value *bool)()
    SetShowOfficeWebApps(value *bool)()
    SetThemeColor(value RgbColorable)()
    SetThemeColorLogo(value MimeContentable)()
}
