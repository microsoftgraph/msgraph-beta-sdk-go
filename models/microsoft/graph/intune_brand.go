package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type IntuneBrand struct {
    additionalData map[string]interface{};
    companyPortalBlockedActions []CompanyPortalBlockedAction;
    contactITEmailAddress *string;
    contactITName *string;
    contactITNotes *string;
    contactITPhoneNumber *string;
    customCanSeePrivacyMessage *string;
    customCantSeePrivacyMessage *string;
    customPrivacyMessage *string;
    darkBackgroundLogo *MimeContent;
    disableClientTelemetry *bool;
    displayName *string;
    enrollmentAvailability *EnrollmentAvailabilityOptions;
    isFactoryResetDisabled *bool;
    isRemoveDeviceDisabled *bool;
    landingPageCustomizedImage *MimeContent;
    lightBackgroundLogo *MimeContent;
    onlineSupportSiteName *string;
    onlineSupportSiteUrl *string;
    privacyUrl *string;
    roleScopeTagIds []string;
    sendDeviceOwnershipChangePushNotification *bool;
    showAzureADEnterpriseApps *bool;
    showDisplayNameNextToLogo *bool;
    showLogo *bool;
    showNameNextToLogo *bool;
    showOfficeWebApps *bool;
    themeColor *RgbColor;
}
func NewIntuneBrand()(*IntuneBrand) {
    m := &IntuneBrand{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *IntuneBrand) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *IntuneBrand) GetCompanyPortalBlockedActions()([]CompanyPortalBlockedAction) {
    if m == nil {
        return nil
    } else {
        return m.companyPortalBlockedActions
    }
}
func (m *IntuneBrand) GetContactITEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITEmailAddress
    }
}
func (m *IntuneBrand) GetContactITName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITName
    }
}
func (m *IntuneBrand) GetContactITNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITNotes
    }
}
func (m *IntuneBrand) GetContactITPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITPhoneNumber
    }
}
func (m *IntuneBrand) GetCustomCanSeePrivacyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customCanSeePrivacyMessage
    }
}
func (m *IntuneBrand) GetCustomCantSeePrivacyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customCantSeePrivacyMessage
    }
}
func (m *IntuneBrand) GetCustomPrivacyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customPrivacyMessage
    }
}
func (m *IntuneBrand) GetDarkBackgroundLogo()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.darkBackgroundLogo
    }
}
func (m *IntuneBrand) GetDisableClientTelemetry()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableClientTelemetry
    }
}
func (m *IntuneBrand) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *IntuneBrand) GetEnrollmentAvailability()(*EnrollmentAvailabilityOptions) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentAvailability
    }
}
func (m *IntuneBrand) GetIsFactoryResetDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFactoryResetDisabled
    }
}
func (m *IntuneBrand) GetIsRemoveDeviceDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRemoveDeviceDisabled
    }
}
func (m *IntuneBrand) GetLandingPageCustomizedImage()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.landingPageCustomizedImage
    }
}
func (m *IntuneBrand) GetLightBackgroundLogo()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.lightBackgroundLogo
    }
}
func (m *IntuneBrand) GetOnlineSupportSiteName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineSupportSiteName
    }
}
func (m *IntuneBrand) GetOnlineSupportSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineSupportSiteUrl
    }
}
func (m *IntuneBrand) GetPrivacyUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privacyUrl
    }
}
func (m *IntuneBrand) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
func (m *IntuneBrand) GetSendDeviceOwnershipChangePushNotification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sendDeviceOwnershipChangePushNotification
    }
}
func (m *IntuneBrand) GetShowAzureADEnterpriseApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showAzureADEnterpriseApps
    }
}
func (m *IntuneBrand) GetShowDisplayNameNextToLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showDisplayNameNextToLogo
    }
}
func (m *IntuneBrand) GetShowLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showLogo
    }
}
func (m *IntuneBrand) GetShowNameNextToLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showNameNextToLogo
    }
}
func (m *IntuneBrand) GetShowOfficeWebApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showOfficeWebApps
    }
}
func (m *IntuneBrand) GetThemeColor()(*RgbColor) {
    if m == nil {
        return nil
    } else {
        return m.themeColor
    }
}
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
func (m *IntuneBrand) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *IntuneBrand) SetCompanyPortalBlockedActions(value []CompanyPortalBlockedAction)() {
    m.companyPortalBlockedActions = value
}
func (m *IntuneBrand) SetContactITEmailAddress(value *string)() {
    m.contactITEmailAddress = value
}
func (m *IntuneBrand) SetContactITName(value *string)() {
    m.contactITName = value
}
func (m *IntuneBrand) SetContactITNotes(value *string)() {
    m.contactITNotes = value
}
func (m *IntuneBrand) SetContactITPhoneNumber(value *string)() {
    m.contactITPhoneNumber = value
}
func (m *IntuneBrand) SetCustomCanSeePrivacyMessage(value *string)() {
    m.customCanSeePrivacyMessage = value
}
func (m *IntuneBrand) SetCustomCantSeePrivacyMessage(value *string)() {
    m.customCantSeePrivacyMessage = value
}
func (m *IntuneBrand) SetCustomPrivacyMessage(value *string)() {
    m.customPrivacyMessage = value
}
func (m *IntuneBrand) SetDarkBackgroundLogo(value *MimeContent)() {
    m.darkBackgroundLogo = value
}
func (m *IntuneBrand) SetDisableClientTelemetry(value *bool)() {
    m.disableClientTelemetry = value
}
func (m *IntuneBrand) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *IntuneBrand) SetEnrollmentAvailability(value *EnrollmentAvailabilityOptions)() {
    m.enrollmentAvailability = value
}
func (m *IntuneBrand) SetIsFactoryResetDisabled(value *bool)() {
    m.isFactoryResetDisabled = value
}
func (m *IntuneBrand) SetIsRemoveDeviceDisabled(value *bool)() {
    m.isRemoveDeviceDisabled = value
}
func (m *IntuneBrand) SetLandingPageCustomizedImage(value *MimeContent)() {
    m.landingPageCustomizedImage = value
}
func (m *IntuneBrand) SetLightBackgroundLogo(value *MimeContent)() {
    m.lightBackgroundLogo = value
}
func (m *IntuneBrand) SetOnlineSupportSiteName(value *string)() {
    m.onlineSupportSiteName = value
}
func (m *IntuneBrand) SetOnlineSupportSiteUrl(value *string)() {
    m.onlineSupportSiteUrl = value
}
func (m *IntuneBrand) SetPrivacyUrl(value *string)() {
    m.privacyUrl = value
}
func (m *IntuneBrand) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
func (m *IntuneBrand) SetSendDeviceOwnershipChangePushNotification(value *bool)() {
    m.sendDeviceOwnershipChangePushNotification = value
}
func (m *IntuneBrand) SetShowAzureADEnterpriseApps(value *bool)() {
    m.showAzureADEnterpriseApps = value
}
func (m *IntuneBrand) SetShowDisplayNameNextToLogo(value *bool)() {
    m.showDisplayNameNextToLogo = value
}
func (m *IntuneBrand) SetShowLogo(value *bool)() {
    m.showLogo = value
}
func (m *IntuneBrand) SetShowNameNextToLogo(value *bool)() {
    m.showNameNextToLogo = value
}
func (m *IntuneBrand) SetShowOfficeWebApps(value *bool)() {
    m.showOfficeWebApps = value
}
func (m *IntuneBrand) SetThemeColor(value *RgbColor)() {
    m.themeColor = value
}
