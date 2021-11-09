package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DepEnrollmentBaseProfile struct {
    EnrollmentProfile
    // Indicates if Apple id setup pane is disabled
    appleIdDisabled *bool;
    // Indicates if Apple pay setup pane is disabled
    applePayDisabled *bool;
    // URL for setup assistant login
    configurationWebUrl *bool;
    // Sets a literal or name pattern.
    deviceNameTemplate *string;
    // Indicates if diagnostics setup pane is disabled
    diagnosticsDisabled *bool;
    // Indicates if displaytone setup screen is disabled
    displayToneSetupDisabled *bool;
    // Indicates if this is the default profile
    isDefault *bool;
    // Indicates if the profile is mandatory
    isMandatory *bool;
    // Indicates if Location service setup pane is disabled
    locationDisabled *bool;
    // Indicates if privacy screen is disabled
    privacyPaneDisabled *bool;
    // Indicates if the profile removal option is disabled
    profileRemovalDisabled *bool;
    // Indicates if Restore setup pane is blocked
    restoreBlocked *bool;
    // Indicates if screen timeout setup is disabled
    screenTimeScreenDisabled *bool;
    // Indicates if siri setup pane is disabled
    siriDisabled *bool;
    // Supervised mode, True to enable, false otherwise. See https://docs.microsoft.com/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
    supervisedModeEnabled *bool;
    // Support department information
    supportDepartment *string;
    // Support phone number
    supportPhoneNumber *string;
    // Indicates if 'Terms and Conditions' setup pane is disabled
    termsAndConditionsDisabled *bool;
    // Indicates if touch id setup pane is disabled
    touchIdDisabled *bool;
}
// Instantiates a new depEnrollmentBaseProfile and sets the default values.
func NewDepEnrollmentBaseProfile()(*DepEnrollmentBaseProfile) {
    m := &DepEnrollmentBaseProfile{
        EnrollmentProfile: *NewEnrollmentProfile(),
    }
    return m
}
// Gets the appleIdDisabled property value. Indicates if Apple id setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetAppleIdDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appleIdDisabled
    }
}
// Gets the applePayDisabled property value. Indicates if Apple pay setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetApplePayDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applePayDisabled
    }
}
// Gets the configurationWebUrl property value. URL for setup assistant login
func (m *DepEnrollmentBaseProfile) GetConfigurationWebUrl()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.configurationWebUrl
    }
}
// Gets the deviceNameTemplate property value. Sets a literal or name pattern.
func (m *DepEnrollmentBaseProfile) GetDeviceNameTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceNameTemplate
    }
}
// Gets the diagnosticsDisabled property value. Indicates if diagnostics setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetDiagnosticsDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.diagnosticsDisabled
    }
}
// Gets the displayToneSetupDisabled property value. Indicates if displaytone setup screen is disabled
func (m *DepEnrollmentBaseProfile) GetDisplayToneSetupDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.displayToneSetupDisabled
    }
}
// Gets the isDefault property value. Indicates if this is the default profile
func (m *DepEnrollmentBaseProfile) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// Gets the isMandatory property value. Indicates if the profile is mandatory
func (m *DepEnrollmentBaseProfile) GetIsMandatory()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMandatory
    }
}
// Gets the locationDisabled property value. Indicates if Location service setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetLocationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.locationDisabled
    }
}
// Gets the privacyPaneDisabled property value. Indicates if privacy screen is disabled
func (m *DepEnrollmentBaseProfile) GetPrivacyPaneDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.privacyPaneDisabled
    }
}
// Gets the profileRemovalDisabled property value. Indicates if the profile removal option is disabled
func (m *DepEnrollmentBaseProfile) GetProfileRemovalDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.profileRemovalDisabled
    }
}
// Gets the restoreBlocked property value. Indicates if Restore setup pane is blocked
func (m *DepEnrollmentBaseProfile) GetRestoreBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.restoreBlocked
    }
}
// Gets the screenTimeScreenDisabled property value. Indicates if screen timeout setup is disabled
func (m *DepEnrollmentBaseProfile) GetScreenTimeScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.screenTimeScreenDisabled
    }
}
// Gets the siriDisabled property value. Indicates if siri setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetSiriDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.siriDisabled
    }
}
// Gets the supervisedModeEnabled property value. Supervised mode, True to enable, false otherwise. See https://docs.microsoft.com/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
func (m *DepEnrollmentBaseProfile) GetSupervisedModeEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.supervisedModeEnabled
    }
}
// Gets the supportDepartment property value. Support department information
func (m *DepEnrollmentBaseProfile) GetSupportDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportDepartment
    }
}
// Gets the supportPhoneNumber property value. Support phone number
func (m *DepEnrollmentBaseProfile) GetSupportPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportPhoneNumber
    }
}
// Gets the termsAndConditionsDisabled property value. Indicates if 'Terms and Conditions' setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetTermsAndConditionsDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.termsAndConditionsDisabled
    }
}
// Gets the touchIdDisabled property value. Indicates if touch id setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetTouchIdDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.touchIdDisabled
    }
}
// The deserialization information for the current model
func (m *DepEnrollmentBaseProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.EnrollmentProfile.GetFieldDeserializers()
    res["appleIdDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleIdDisabled(val)
        }
        return nil
    }
    res["applePayDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplePayDisabled(val)
        }
        return nil
    }
    res["configurationWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationWebUrl(val)
        }
        return nil
    }
    res["deviceNameTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceNameTemplate(val)
        }
        return nil
    }
    res["diagnosticsDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiagnosticsDisabled(val)
        }
        return nil
    }
    res["displayToneSetupDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayToneSetupDisabled(val)
        }
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isMandatory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMandatory(val)
        }
        return nil
    }
    res["locationDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationDisabled(val)
        }
        return nil
    }
    res["privacyPaneDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyPaneDisabled(val)
        }
        return nil
    }
    res["profileRemovalDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileRemovalDisabled(val)
        }
        return nil
    }
    res["restoreBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestoreBlocked(val)
        }
        return nil
    }
    res["screenTimeScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreenTimeScreenDisabled(val)
        }
        return nil
    }
    res["siriDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiriDisabled(val)
        }
        return nil
    }
    res["supervisedModeEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupervisedModeEnabled(val)
        }
        return nil
    }
    res["supportDepartment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportDepartment(val)
        }
        return nil
    }
    res["supportPhoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportPhoneNumber(val)
        }
        return nil
    }
    res["termsAndConditionsDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsAndConditionsDisabled(val)
        }
        return nil
    }
    res["touchIdDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTouchIdDisabled(val)
        }
        return nil
    }
    return res
}
func (m *DepEnrollmentBaseProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DepEnrollmentBaseProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.EnrollmentProfile.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("appleIdDisabled", m.GetAppleIdDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applePayDisabled", m.GetApplePayDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("configurationWebUrl", m.GetConfigurationWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceNameTemplate", m.GetDeviceNameTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("diagnosticsDisabled", m.GetDiagnosticsDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("displayToneSetupDisabled", m.GetDisplayToneSetupDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMandatory", m.GetIsMandatory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("locationDisabled", m.GetLocationDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("privacyPaneDisabled", m.GetPrivacyPaneDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("profileRemovalDisabled", m.GetProfileRemovalDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("restoreBlocked", m.GetRestoreBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("screenTimeScreenDisabled", m.GetScreenTimeScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("siriDisabled", m.GetSiriDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("supervisedModeEnabled", m.GetSupervisedModeEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("supportDepartment", m.GetSupportDepartment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("supportPhoneNumber", m.GetSupportPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("termsAndConditionsDisabled", m.GetTermsAndConditionsDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("touchIdDisabled", m.GetTouchIdDisabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appleIdDisabled property value. Indicates if Apple id setup pane is disabled
// Parameters:
//  - value : Value to set for the appleIdDisabled property.
func (m *DepEnrollmentBaseProfile) SetAppleIdDisabled(value *bool)() {
    m.appleIdDisabled = value
}
// Sets the applePayDisabled property value. Indicates if Apple pay setup pane is disabled
// Parameters:
//  - value : Value to set for the applePayDisabled property.
func (m *DepEnrollmentBaseProfile) SetApplePayDisabled(value *bool)() {
    m.applePayDisabled = value
}
// Sets the configurationWebUrl property value. URL for setup assistant login
// Parameters:
//  - value : Value to set for the configurationWebUrl property.
func (m *DepEnrollmentBaseProfile) SetConfigurationWebUrl(value *bool)() {
    m.configurationWebUrl = value
}
// Sets the deviceNameTemplate property value. Sets a literal or name pattern.
// Parameters:
//  - value : Value to set for the deviceNameTemplate property.
func (m *DepEnrollmentBaseProfile) SetDeviceNameTemplate(value *string)() {
    m.deviceNameTemplate = value
}
// Sets the diagnosticsDisabled property value. Indicates if diagnostics setup pane is disabled
// Parameters:
//  - value : Value to set for the diagnosticsDisabled property.
func (m *DepEnrollmentBaseProfile) SetDiagnosticsDisabled(value *bool)() {
    m.diagnosticsDisabled = value
}
// Sets the displayToneSetupDisabled property value. Indicates if displaytone setup screen is disabled
// Parameters:
//  - value : Value to set for the displayToneSetupDisabled property.
func (m *DepEnrollmentBaseProfile) SetDisplayToneSetupDisabled(value *bool)() {
    m.displayToneSetupDisabled = value
}
// Sets the isDefault property value. Indicates if this is the default profile
// Parameters:
//  - value : Value to set for the isDefault property.
func (m *DepEnrollmentBaseProfile) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// Sets the isMandatory property value. Indicates if the profile is mandatory
// Parameters:
//  - value : Value to set for the isMandatory property.
func (m *DepEnrollmentBaseProfile) SetIsMandatory(value *bool)() {
    m.isMandatory = value
}
// Sets the locationDisabled property value. Indicates if Location service setup pane is disabled
// Parameters:
//  - value : Value to set for the locationDisabled property.
func (m *DepEnrollmentBaseProfile) SetLocationDisabled(value *bool)() {
    m.locationDisabled = value
}
// Sets the privacyPaneDisabled property value. Indicates if privacy screen is disabled
// Parameters:
//  - value : Value to set for the privacyPaneDisabled property.
func (m *DepEnrollmentBaseProfile) SetPrivacyPaneDisabled(value *bool)() {
    m.privacyPaneDisabled = value
}
// Sets the profileRemovalDisabled property value. Indicates if the profile removal option is disabled
// Parameters:
//  - value : Value to set for the profileRemovalDisabled property.
func (m *DepEnrollmentBaseProfile) SetProfileRemovalDisabled(value *bool)() {
    m.profileRemovalDisabled = value
}
// Sets the restoreBlocked property value. Indicates if Restore setup pane is blocked
// Parameters:
//  - value : Value to set for the restoreBlocked property.
func (m *DepEnrollmentBaseProfile) SetRestoreBlocked(value *bool)() {
    m.restoreBlocked = value
}
// Sets the screenTimeScreenDisabled property value. Indicates if screen timeout setup is disabled
// Parameters:
//  - value : Value to set for the screenTimeScreenDisabled property.
func (m *DepEnrollmentBaseProfile) SetScreenTimeScreenDisabled(value *bool)() {
    m.screenTimeScreenDisabled = value
}
// Sets the siriDisabled property value. Indicates if siri setup pane is disabled
// Parameters:
//  - value : Value to set for the siriDisabled property.
func (m *DepEnrollmentBaseProfile) SetSiriDisabled(value *bool)() {
    m.siriDisabled = value
}
// Sets the supervisedModeEnabled property value. Supervised mode, True to enable, false otherwise. See https://docs.microsoft.com/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
// Parameters:
//  - value : Value to set for the supervisedModeEnabled property.
func (m *DepEnrollmentBaseProfile) SetSupervisedModeEnabled(value *bool)() {
    m.supervisedModeEnabled = value
}
// Sets the supportDepartment property value. Support department information
// Parameters:
//  - value : Value to set for the supportDepartment property.
func (m *DepEnrollmentBaseProfile) SetSupportDepartment(value *string)() {
    m.supportDepartment = value
}
// Sets the supportPhoneNumber property value. Support phone number
// Parameters:
//  - value : Value to set for the supportPhoneNumber property.
func (m *DepEnrollmentBaseProfile) SetSupportPhoneNumber(value *string)() {
    m.supportPhoneNumber = value
}
// Sets the termsAndConditionsDisabled property value. Indicates if 'Terms and Conditions' setup pane is disabled
// Parameters:
//  - value : Value to set for the termsAndConditionsDisabled property.
func (m *DepEnrollmentBaseProfile) SetTermsAndConditionsDisabled(value *bool)() {
    m.termsAndConditionsDisabled = value
}
// Sets the touchIdDisabled property value. Indicates if touch id setup pane is disabled
// Parameters:
//  - value : Value to set for the touchIdDisabled property.
func (m *DepEnrollmentBaseProfile) SetTouchIdDisabled(value *bool)() {
    m.touchIdDisabled = value
}
