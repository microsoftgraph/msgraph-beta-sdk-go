package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LoginPageTextVisibilitySettingsable 
type LoginPageTextVisibilitySettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetHideCannotAccessYourAccount()(*bool)
    GetHideForgotMyPassword()(*bool)
    GetHidePrivacyAndCookies()(*bool)
    GetHideResetItNow()(*bool)
    GetHideTermsOfUse()(*bool)
    SetHideCannotAccessYourAccount(value *bool)()
    SetHideForgotMyPassword(value *bool)()
    SetHidePrivacyAndCookies(value *bool)()
    SetHideResetItNow(value *bool)()
    SetHideTermsOfUse(value *bool)()
}
