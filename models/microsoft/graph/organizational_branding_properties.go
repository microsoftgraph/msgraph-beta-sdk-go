package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OrganizationalBrandingProperties struct {
    Entity
    // Color that will appear in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
    backgroundColor *string;
    // Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
    backgroundImage []byte;
    // A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
    backgroundImageRelativeUrl *string;
    // A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
    bannerLogo []byte;
    // A relative url for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
    bannerLogoRelativeUrl *string;
    // A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
    cdnList []string;
    // 
    customAccountResetCredentialsUrl *string;
    // 
    customCannotAccessYourAccountText *string;
    // 
    customCannotAccessYourAccountUrl *string;
    // 
    customForgotMyPasswordText *string;
    // 
    customPrivacyAndCookiesText *string;
    // 
    customPrivacyAndCookiesUrl *string;
    // 
    customResetItNowText *string;
    // 
    customTermsOfUseText *string;
    // 
    customTermsOfUseUrl *string;
    // 
    favicon []byte;
    // 
    faviconRelativeUrl *string;
    // 
    headerBackgroundColor *string;
    // 
    loginPageTextVisibilitySettings *LoginPageTextVisibilitySettings;
    // Text that appears at the bottom of the sign-in box. You can use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be Unicode and not exceed 1024 characters.
    signInPageText *string;
    // A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG no larger than 240 x 240 pixels and no more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
    squareLogo []byte;
    // A relative url for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
    squareLogoRelativeUrl *string;
    // String that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
    usernameHintText *string;
}
// Instantiates a new organizationalBrandingProperties and sets the default values.
func NewOrganizationalBrandingProperties()(*OrganizationalBrandingProperties) {
    m := &OrganizationalBrandingProperties{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the backgroundColor property value. Color that will appear in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
func (m *OrganizationalBrandingProperties) GetBackgroundColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.backgroundColor
    }
}
// Gets the backgroundImage property value. Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
func (m *OrganizationalBrandingProperties) GetBackgroundImage()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.backgroundImage
    }
}
// Gets the backgroundImageRelativeUrl property value. A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetBackgroundImageRelativeUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.backgroundImageRelativeUrl
    }
}
// Gets the bannerLogo property value. A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) GetBannerLogo()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.bannerLogo
    }
}
// Gets the bannerLogoRelativeUrl property value. A relative url for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetBannerLogoRelativeUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bannerLogoRelativeUrl
    }
}
// Gets the cdnList property value. A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
func (m *OrganizationalBrandingProperties) GetCdnList()([]string) {
    if m == nil {
        return nil
    } else {
        return m.cdnList
    }
}
// Gets the customAccountResetCredentialsUrl property value. 
func (m *OrganizationalBrandingProperties) GetCustomAccountResetCredentialsUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customAccountResetCredentialsUrl
    }
}
// Gets the customCannotAccessYourAccountText property value. 
func (m *OrganizationalBrandingProperties) GetCustomCannotAccessYourAccountText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customCannotAccessYourAccountText
    }
}
// Gets the customCannotAccessYourAccountUrl property value. 
func (m *OrganizationalBrandingProperties) GetCustomCannotAccessYourAccountUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customCannotAccessYourAccountUrl
    }
}
// Gets the customForgotMyPasswordText property value. 
func (m *OrganizationalBrandingProperties) GetCustomForgotMyPasswordText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customForgotMyPasswordText
    }
}
// Gets the customPrivacyAndCookiesText property value. 
func (m *OrganizationalBrandingProperties) GetCustomPrivacyAndCookiesText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customPrivacyAndCookiesText
    }
}
// Gets the customPrivacyAndCookiesUrl property value. 
func (m *OrganizationalBrandingProperties) GetCustomPrivacyAndCookiesUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customPrivacyAndCookiesUrl
    }
}
// Gets the customResetItNowText property value. 
func (m *OrganizationalBrandingProperties) GetCustomResetItNowText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customResetItNowText
    }
}
// Gets the customTermsOfUseText property value. 
func (m *OrganizationalBrandingProperties) GetCustomTermsOfUseText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customTermsOfUseText
    }
}
// Gets the customTermsOfUseUrl property value. 
func (m *OrganizationalBrandingProperties) GetCustomTermsOfUseUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customTermsOfUseUrl
    }
}
// Gets the favicon property value. 
func (m *OrganizationalBrandingProperties) GetFavicon()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.favicon
    }
}
// Gets the faviconRelativeUrl property value. 
func (m *OrganizationalBrandingProperties) GetFaviconRelativeUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.faviconRelativeUrl
    }
}
// Gets the headerBackgroundColor property value. 
func (m *OrganizationalBrandingProperties) GetHeaderBackgroundColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.headerBackgroundColor
    }
}
// Gets the loginPageTextVisibilitySettings property value. 
func (m *OrganizationalBrandingProperties) GetLoginPageTextVisibilitySettings()(*LoginPageTextVisibilitySettings) {
    if m == nil {
        return nil
    } else {
        return m.loginPageTextVisibilitySettings
    }
}
// Gets the signInPageText property value. Text that appears at the bottom of the sign-in box. You can use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be Unicode and not exceed 1024 characters.
func (m *OrganizationalBrandingProperties) GetSignInPageText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signInPageText
    }
}
// Gets the squareLogo property value. A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG no larger than 240 x 240 pixels and no more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) GetSquareLogo()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.squareLogo
    }
}
// Gets the squareLogoRelativeUrl property value. A relative url for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetSquareLogoRelativeUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.squareLogoRelativeUrl
    }
}
// Gets the usernameHintText property value. String that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
func (m *OrganizationalBrandingProperties) GetUsernameHintText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.usernameHintText
    }
}
// The deserialization information for the current model
func (m *OrganizationalBrandingProperties) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["backgroundColor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBackgroundColor(val)
        return nil
    }
    res["backgroundImage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetBackgroundImage(val)
        return nil
    }
    res["backgroundImageRelativeUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBackgroundImageRelativeUrl(val)
        return nil
    }
    res["bannerLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetBannerLogo(val)
        return nil
    }
    res["bannerLogoRelativeUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBannerLogoRelativeUrl(val)
        return nil
    }
    res["cdnList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetCdnList(res)
        return nil
    }
    res["customAccountResetCredentialsUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomAccountResetCredentialsUrl(val)
        return nil
    }
    res["customCannotAccessYourAccountText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomCannotAccessYourAccountText(val)
        return nil
    }
    res["customCannotAccessYourAccountUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomCannotAccessYourAccountUrl(val)
        return nil
    }
    res["customForgotMyPasswordText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomForgotMyPasswordText(val)
        return nil
    }
    res["customPrivacyAndCookiesText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomPrivacyAndCookiesText(val)
        return nil
    }
    res["customPrivacyAndCookiesUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomPrivacyAndCookiesUrl(val)
        return nil
    }
    res["customResetItNowText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomResetItNowText(val)
        return nil
    }
    res["customTermsOfUseText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomTermsOfUseText(val)
        return nil
    }
    res["customTermsOfUseUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomTermsOfUseUrl(val)
        return nil
    }
    res["favicon"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetFavicon(val)
        return nil
    }
    res["faviconRelativeUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFaviconRelativeUrl(val)
        return nil
    }
    res["headerBackgroundColor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHeaderBackgroundColor(val)
        return nil
    }
    res["loginPageTextVisibilitySettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLoginPageTextVisibilitySettings() })
        if err != nil {
            return err
        }
        m.SetLoginPageTextVisibilitySettings(val.(*LoginPageTextVisibilitySettings))
        return nil
    }
    res["signInPageText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSignInPageText(val)
        return nil
    }
    res["squareLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetSquareLogo(val)
        return nil
    }
    res["squareLogoRelativeUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSquareLogoRelativeUrl(val)
        return nil
    }
    res["usernameHintText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUsernameHintText(val)
        return nil
    }
    return res
}
func (m *OrganizationalBrandingProperties) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OrganizationalBrandingProperties) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("backgroundColor", m.GetBackgroundColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("backgroundImage", m.GetBackgroundImage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("backgroundImageRelativeUrl", m.GetBackgroundImageRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("bannerLogo", m.GetBannerLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bannerLogoRelativeUrl", m.GetBannerLogoRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("cdnList", m.GetCdnList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customAccountResetCredentialsUrl", m.GetCustomAccountResetCredentialsUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customCannotAccessYourAccountText", m.GetCustomCannotAccessYourAccountText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customCannotAccessYourAccountUrl", m.GetCustomCannotAccessYourAccountUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customForgotMyPasswordText", m.GetCustomForgotMyPasswordText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customPrivacyAndCookiesText", m.GetCustomPrivacyAndCookiesText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customPrivacyAndCookiesUrl", m.GetCustomPrivacyAndCookiesUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customResetItNowText", m.GetCustomResetItNowText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customTermsOfUseText", m.GetCustomTermsOfUseText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customTermsOfUseUrl", m.GetCustomTermsOfUseUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("favicon", m.GetFavicon())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("faviconRelativeUrl", m.GetFaviconRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("headerBackgroundColor", m.GetHeaderBackgroundColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("loginPageTextVisibilitySettings", m.GetLoginPageTextVisibilitySettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInPageText", m.GetSignInPageText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("squareLogo", m.GetSquareLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("squareLogoRelativeUrl", m.GetSquareLogoRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usernameHintText", m.GetUsernameHintText())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the backgroundColor property value. Color that will appear in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
// Parameters:
//  - value : Value to set for the backgroundColor property.
func (m *OrganizationalBrandingProperties) SetBackgroundColor(value *string)() {
    m.backgroundColor = value
}
// Sets the backgroundImage property value. Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
// Parameters:
//  - value : Value to set for the backgroundImage property.
func (m *OrganizationalBrandingProperties) SetBackgroundImage(value []byte)() {
    m.backgroundImage = value
}
// Sets the backgroundImageRelativeUrl property value. A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
// Parameters:
//  - value : Value to set for the backgroundImageRelativeUrl property.
func (m *OrganizationalBrandingProperties) SetBackgroundImageRelativeUrl(value *string)() {
    m.backgroundImageRelativeUrl = value
}
// Sets the bannerLogo property value. A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
// Parameters:
//  - value : Value to set for the bannerLogo property.
func (m *OrganizationalBrandingProperties) SetBannerLogo(value []byte)() {
    m.bannerLogo = value
}
// Sets the bannerLogoRelativeUrl property value. A relative url for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
// Parameters:
//  - value : Value to set for the bannerLogoRelativeUrl property.
func (m *OrganizationalBrandingProperties) SetBannerLogoRelativeUrl(value *string)() {
    m.bannerLogoRelativeUrl = value
}
// Sets the cdnList property value. A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
// Parameters:
//  - value : Value to set for the cdnList property.
func (m *OrganizationalBrandingProperties) SetCdnList(value []string)() {
    m.cdnList = value
}
// Sets the customAccountResetCredentialsUrl property value. 
// Parameters:
//  - value : Value to set for the customAccountResetCredentialsUrl property.
func (m *OrganizationalBrandingProperties) SetCustomAccountResetCredentialsUrl(value *string)() {
    m.customAccountResetCredentialsUrl = value
}
// Sets the customCannotAccessYourAccountText property value. 
// Parameters:
//  - value : Value to set for the customCannotAccessYourAccountText property.
func (m *OrganizationalBrandingProperties) SetCustomCannotAccessYourAccountText(value *string)() {
    m.customCannotAccessYourAccountText = value
}
// Sets the customCannotAccessYourAccountUrl property value. 
// Parameters:
//  - value : Value to set for the customCannotAccessYourAccountUrl property.
func (m *OrganizationalBrandingProperties) SetCustomCannotAccessYourAccountUrl(value *string)() {
    m.customCannotAccessYourAccountUrl = value
}
// Sets the customForgotMyPasswordText property value. 
// Parameters:
//  - value : Value to set for the customForgotMyPasswordText property.
func (m *OrganizationalBrandingProperties) SetCustomForgotMyPasswordText(value *string)() {
    m.customForgotMyPasswordText = value
}
// Sets the customPrivacyAndCookiesText property value. 
// Parameters:
//  - value : Value to set for the customPrivacyAndCookiesText property.
func (m *OrganizationalBrandingProperties) SetCustomPrivacyAndCookiesText(value *string)() {
    m.customPrivacyAndCookiesText = value
}
// Sets the customPrivacyAndCookiesUrl property value. 
// Parameters:
//  - value : Value to set for the customPrivacyAndCookiesUrl property.
func (m *OrganizationalBrandingProperties) SetCustomPrivacyAndCookiesUrl(value *string)() {
    m.customPrivacyAndCookiesUrl = value
}
// Sets the customResetItNowText property value. 
// Parameters:
//  - value : Value to set for the customResetItNowText property.
func (m *OrganizationalBrandingProperties) SetCustomResetItNowText(value *string)() {
    m.customResetItNowText = value
}
// Sets the customTermsOfUseText property value. 
// Parameters:
//  - value : Value to set for the customTermsOfUseText property.
func (m *OrganizationalBrandingProperties) SetCustomTermsOfUseText(value *string)() {
    m.customTermsOfUseText = value
}
// Sets the customTermsOfUseUrl property value. 
// Parameters:
//  - value : Value to set for the customTermsOfUseUrl property.
func (m *OrganizationalBrandingProperties) SetCustomTermsOfUseUrl(value *string)() {
    m.customTermsOfUseUrl = value
}
// Sets the favicon property value. 
// Parameters:
//  - value : Value to set for the favicon property.
func (m *OrganizationalBrandingProperties) SetFavicon(value []byte)() {
    m.favicon = value
}
// Sets the faviconRelativeUrl property value. 
// Parameters:
//  - value : Value to set for the faviconRelativeUrl property.
func (m *OrganizationalBrandingProperties) SetFaviconRelativeUrl(value *string)() {
    m.faviconRelativeUrl = value
}
// Sets the headerBackgroundColor property value. 
// Parameters:
//  - value : Value to set for the headerBackgroundColor property.
func (m *OrganizationalBrandingProperties) SetHeaderBackgroundColor(value *string)() {
    m.headerBackgroundColor = value
}
// Sets the loginPageTextVisibilitySettings property value. 
// Parameters:
//  - value : Value to set for the loginPageTextVisibilitySettings property.
func (m *OrganizationalBrandingProperties) SetLoginPageTextVisibilitySettings(value *LoginPageTextVisibilitySettings)() {
    m.loginPageTextVisibilitySettings = value
}
// Sets the signInPageText property value. Text that appears at the bottom of the sign-in box. You can use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be Unicode and not exceed 1024 characters.
// Parameters:
//  - value : Value to set for the signInPageText property.
func (m *OrganizationalBrandingProperties) SetSignInPageText(value *string)() {
    m.signInPageText = value
}
// Sets the squareLogo property value. A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG no larger than 240 x 240 pixels and no more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
// Parameters:
//  - value : Value to set for the squareLogo property.
func (m *OrganizationalBrandingProperties) SetSquareLogo(value []byte)() {
    m.squareLogo = value
}
// Sets the squareLogoRelativeUrl property value. A relative url for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
// Parameters:
//  - value : Value to set for the squareLogoRelativeUrl property.
func (m *OrganizationalBrandingProperties) SetSquareLogoRelativeUrl(value *string)() {
    m.squareLogoRelativeUrl = value
}
// Sets the usernameHintText property value. String that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
// Parameters:
//  - value : Value to set for the usernameHintText property.
func (m *OrganizationalBrandingProperties) SetUsernameHintText(value *string)() {
    m.usernameHintText = value
}
