package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DepEnrollmentBaseProfile struct {
    EnrollmentProfile
    appleIdDisabled *bool;
    applePayDisabled *bool;
    configurationWebUrl *bool;
    deviceNameTemplate *string;
    diagnosticsDisabled *bool;
    displayToneSetupDisabled *bool;
    isDefault *bool;
    isMandatory *bool;
    locationDisabled *bool;
    privacyPaneDisabled *bool;
    profileRemovalDisabled *bool;
    restoreBlocked *bool;
    screenTimeScreenDisabled *bool;
    siriDisabled *bool;
    supervisedModeEnabled *bool;
    supportDepartment *string;
    supportPhoneNumber *string;
    termsAndConditionsDisabled *bool;
    touchIdDisabled *bool;
}
func NewDepEnrollmentBaseProfile()(*DepEnrollmentBaseProfile) {
    m := &DepEnrollmentBaseProfile{
        EnrollmentProfile: *NewEnrollmentProfile(),
    }
    return m
}
func (m *DepEnrollmentBaseProfile) GetAppleIdDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appleIdDisabled
    }
}
func (m *DepEnrollmentBaseProfile) GetApplePayDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applePayDisabled
    }
}
func (m *DepEnrollmentBaseProfile) GetConfigurationWebUrl()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.configurationWebUrl
    }
}
func (m *DepEnrollmentBaseProfile) GetDeviceNameTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceNameTemplate
    }
}
func (m *DepEnrollmentBaseProfile) GetDiagnosticsDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.diagnosticsDisabled
    }
}
func (m *DepEnrollmentBaseProfile) GetDisplayToneSetupDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.displayToneSetupDisabled
    }
}
func (m *DepEnrollmentBaseProfile) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
func (m *DepEnrollmentBaseProfile) GetIsMandatory()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMandatory
    }
}
func (m *DepEnrollmentBaseProfile) GetLocationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.locationDisabled
    }
}
func (m *DepEnrollmentBaseProfile) GetPrivacyPaneDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.privacyPaneDisabled
    }
}
func (m *DepEnrollmentBaseProfile) GetProfileRemovalDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.profileRemovalDisabled
    }
}
func (m *DepEnrollmentBaseProfile) GetRestoreBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.restoreBlocked
    }
}
func (m *DepEnrollmentBaseProfile) GetScreenTimeScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.screenTimeScreenDisabled
    }
}
func (m *DepEnrollmentBaseProfile) GetSiriDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.siriDisabled
    }
}
func (m *DepEnrollmentBaseProfile) GetSupervisedModeEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.supervisedModeEnabled
    }
}
func (m *DepEnrollmentBaseProfile) GetSupportDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportDepartment
    }
}
func (m *DepEnrollmentBaseProfile) GetSupportPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportPhoneNumber
    }
}
func (m *DepEnrollmentBaseProfile) GetTermsAndConditionsDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.termsAndConditionsDisabled
    }
}
func (m *DepEnrollmentBaseProfile) GetTouchIdDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.touchIdDisabled
    }
}
func (m *DepEnrollmentBaseProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.EnrollmentProfile.GetFieldDeserializers()
    res["appleIdDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAppleIdDisabled(val)
        return nil
    }
    res["applePayDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetApplePayDisabled(val)
        return nil
    }
    res["configurationWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetConfigurationWebUrl(val)
        return nil
    }
    res["deviceNameTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceNameTemplate(val)
        return nil
    }
    res["diagnosticsDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDiagnosticsDisabled(val)
        return nil
    }
    res["displayToneSetupDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDisplayToneSetupDisabled(val)
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
        return nil
    }
    res["isMandatory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMandatory(val)
        return nil
    }
    res["locationDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetLocationDisabled(val)
        return nil
    }
    res["privacyPaneDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPrivacyPaneDisabled(val)
        return nil
    }
    res["profileRemovalDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetProfileRemovalDisabled(val)
        return nil
    }
    res["restoreBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRestoreBlocked(val)
        return nil
    }
    res["screenTimeScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetScreenTimeScreenDisabled(val)
        return nil
    }
    res["siriDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSiriDisabled(val)
        return nil
    }
    res["supervisedModeEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSupervisedModeEnabled(val)
        return nil
    }
    res["supportDepartment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSupportDepartment(val)
        return nil
    }
    res["supportPhoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSupportPhoneNumber(val)
        return nil
    }
    res["termsAndConditionsDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetTermsAndConditionsDisabled(val)
        return nil
    }
    res["touchIdDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetTouchIdDisabled(val)
        return nil
    }
    return res
}
func (m *DepEnrollmentBaseProfile) IsNil()(bool) {
    return m == nil
}
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
func (m *DepEnrollmentBaseProfile) SetAppleIdDisabled(value *bool)() {
    m.appleIdDisabled = value
}
func (m *DepEnrollmentBaseProfile) SetApplePayDisabled(value *bool)() {
    m.applePayDisabled = value
}
func (m *DepEnrollmentBaseProfile) SetConfigurationWebUrl(value *bool)() {
    m.configurationWebUrl = value
}
func (m *DepEnrollmentBaseProfile) SetDeviceNameTemplate(value *string)() {
    m.deviceNameTemplate = value
}
func (m *DepEnrollmentBaseProfile) SetDiagnosticsDisabled(value *bool)() {
    m.diagnosticsDisabled = value
}
func (m *DepEnrollmentBaseProfile) SetDisplayToneSetupDisabled(value *bool)() {
    m.displayToneSetupDisabled = value
}
func (m *DepEnrollmentBaseProfile) SetIsDefault(value *bool)() {
    m.isDefault = value
}
func (m *DepEnrollmentBaseProfile) SetIsMandatory(value *bool)() {
    m.isMandatory = value
}
func (m *DepEnrollmentBaseProfile) SetLocationDisabled(value *bool)() {
    m.locationDisabled = value
}
func (m *DepEnrollmentBaseProfile) SetPrivacyPaneDisabled(value *bool)() {
    m.privacyPaneDisabled = value
}
func (m *DepEnrollmentBaseProfile) SetProfileRemovalDisabled(value *bool)() {
    m.profileRemovalDisabled = value
}
func (m *DepEnrollmentBaseProfile) SetRestoreBlocked(value *bool)() {
    m.restoreBlocked = value
}
func (m *DepEnrollmentBaseProfile) SetScreenTimeScreenDisabled(value *bool)() {
    m.screenTimeScreenDisabled = value
}
func (m *DepEnrollmentBaseProfile) SetSiriDisabled(value *bool)() {
    m.siriDisabled = value
}
func (m *DepEnrollmentBaseProfile) SetSupervisedModeEnabled(value *bool)() {
    m.supervisedModeEnabled = value
}
func (m *DepEnrollmentBaseProfile) SetSupportDepartment(value *string)() {
    m.supportDepartment = value
}
func (m *DepEnrollmentBaseProfile) SetSupportPhoneNumber(value *string)() {
    m.supportPhoneNumber = value
}
func (m *DepEnrollmentBaseProfile) SetTermsAndConditionsDisabled(value *bool)() {
    m.termsAndConditionsDisabled = value
}
func (m *DepEnrollmentBaseProfile) SetTouchIdDisabled(value *bool)() {
    m.touchIdDisabled = value
}
