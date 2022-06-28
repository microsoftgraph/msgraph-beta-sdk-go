package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalBrandingPropertiesable 
type OrganizationalBrandingPropertiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackgroundColor()(*string)
    GetBackgroundImage()([]byte)
    GetBackgroundImageRelativeUrl()(*string)
    GetBannerLogo()([]byte)
    GetBannerLogoRelativeUrl()(*string)
    GetCdnList()([]string)
    GetCustomAccountResetCredentialsUrl()(*string)
    GetCustomCannotAccessYourAccountText()(*string)
    GetCustomCannotAccessYourAccountUrl()(*string)
    GetCustomForgotMyPasswordText()(*string)
    GetCustomPrivacyAndCookiesText()(*string)
    GetCustomPrivacyAndCookiesUrl()(*string)
    GetCustomResetItNowText()(*string)
    GetCustomTermsOfUseText()(*string)
    GetCustomTermsOfUseUrl()(*string)
    GetFavicon()([]byte)
    GetFaviconRelativeUrl()(*string)
    GetHeaderBackgroundColor()(*string)
    GetLoginPageTextVisibilitySettings()(LoginPageTextVisibilitySettingsable)
    GetSignInPageText()(*string)
    GetSquareLogo()([]byte)
    GetSquareLogoRelativeUrl()(*string)
    GetUsernameHintText()(*string)
    SetBackgroundColor(value *string)()
    SetBackgroundImage(value []byte)()
    SetBackgroundImageRelativeUrl(value *string)()
    SetBannerLogo(value []byte)()
    SetBannerLogoRelativeUrl(value *string)()
    SetCdnList(value []string)()
    SetCustomAccountResetCredentialsUrl(value *string)()
    SetCustomCannotAccessYourAccountText(value *string)()
    SetCustomCannotAccessYourAccountUrl(value *string)()
    SetCustomForgotMyPasswordText(value *string)()
    SetCustomPrivacyAndCookiesText(value *string)()
    SetCustomPrivacyAndCookiesUrl(value *string)()
    SetCustomResetItNowText(value *string)()
    SetCustomTermsOfUseText(value *string)()
    SetCustomTermsOfUseUrl(value *string)()
    SetFavicon(value []byte)()
    SetFaviconRelativeUrl(value *string)()
    SetHeaderBackgroundColor(value *string)()
    SetLoginPageTextVisibilitySettings(value LoginPageTextVisibilitySettingsable)()
    SetSignInPageText(value *string)()
    SetSquareLogo(value []byte)()
    SetSquareLogoRelativeUrl(value *string)()
    SetUsernameHintText(value *string)()
}
