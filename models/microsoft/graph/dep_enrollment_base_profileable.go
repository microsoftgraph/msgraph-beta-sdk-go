package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DepEnrollmentBaseProfileable 
type DepEnrollmentBaseProfileable interface {
    EnrollmentProfileable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppleIdDisabled()(*bool)
    GetApplePayDisabled()(*bool)
    GetConfigurationWebUrl()(*bool)
    GetDeviceNameTemplate()(*string)
    GetDiagnosticsDisabled()(*bool)
    GetDisplayToneSetupDisabled()(*bool)
    GetIsDefault()(*bool)
    GetIsMandatory()(*bool)
    GetLocationDisabled()(*bool)
    GetPrivacyPaneDisabled()(*bool)
    GetProfileRemovalDisabled()(*bool)
    GetRestoreBlocked()(*bool)
    GetScreenTimeScreenDisabled()(*bool)
    GetSiriDisabled()(*bool)
    GetSupervisedModeEnabled()(*bool)
    GetSupportDepartment()(*string)
    GetSupportPhoneNumber()(*string)
    GetTermsAndConditionsDisabled()(*bool)
    GetTouchIdDisabled()(*bool)
    SetAppleIdDisabled(value *bool)()
    SetApplePayDisabled(value *bool)()
    SetConfigurationWebUrl(value *bool)()
    SetDeviceNameTemplate(value *string)()
    SetDiagnosticsDisabled(value *bool)()
    SetDisplayToneSetupDisabled(value *bool)()
    SetIsDefault(value *bool)()
    SetIsMandatory(value *bool)()
    SetLocationDisabled(value *bool)()
    SetPrivacyPaneDisabled(value *bool)()
    SetProfileRemovalDisabled(value *bool)()
    SetRestoreBlocked(value *bool)()
    SetScreenTimeScreenDisabled(value *bool)()
    SetSiriDisabled(value *bool)()
    SetSupervisedModeEnabled(value *bool)()
    SetSupportDepartment(value *string)()
    SetSupportPhoneNumber(value *string)()
    SetTermsAndConditionsDisabled(value *bool)()
    SetTouchIdDisabled(value *bool)()
}
