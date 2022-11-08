package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepEnrollmentBaseProfile 
type DepEnrollmentBaseProfile struct {
    EnrollmentProfile
    // Indicates if Apple id setup pane is disabled
    appleIdDisabled *bool
    // Indicates if Apple pay setup pane is disabled
    applePayDisabled *bool
    // URL for setup assistant login
    configurationWebUrl *bool
    // Sets a literal or name pattern.
    deviceNameTemplate *string
    // Indicates if diagnostics setup pane is disabled
    diagnosticsDisabled *bool
    // Indicates if displaytone setup screen is disabled
    displayToneSetupDisabled *bool
    // Indicates if this is the default profile
    isDefault *bool
    // Indicates if the profile is mandatory
    isMandatory *bool
    // Indicates if Location service setup pane is disabled
    locationDisabled *bool
    // Indicates if privacy screen is disabled
    privacyPaneDisabled *bool
    // Indicates if the profile removal option is disabled
    profileRemovalDisabled *bool
    // Indicates if Restore setup pane is blocked
    restoreBlocked *bool
    // Indicates if screen timeout setup is disabled
    screenTimeScreenDisabled *bool
    // Indicates if siri setup pane is disabled
    siriDisabled *bool
    // Supervised mode, True to enable, false otherwise. See https://learn.microsoft.com/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
    supervisedModeEnabled *bool
    // Support department information
    supportDepartment *string
    // Support phone number
    supportPhoneNumber *string
    // Indicates if 'Terms and Conditions' setup pane is disabled
    termsAndConditionsDisabled *bool
    // Indicates if touch id setup pane is disabled
    touchIdDisabled *bool
}
// NewDepEnrollmentBaseProfile instantiates a new DepEnrollmentBaseProfile and sets the default values.
func NewDepEnrollmentBaseProfile()(*DepEnrollmentBaseProfile) {
    m := &DepEnrollmentBaseProfile{
        EnrollmentProfile: *NewEnrollmentProfile(),
    }
    odataTypeValue := "#microsoft.graph.depEnrollmentBaseProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDepEnrollmentBaseProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDepEnrollmentBaseProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.depIOSEnrollmentProfile":
                        return NewDepIOSEnrollmentProfile(), nil
                    case "#microsoft.graph.depMacOSEnrollmentProfile":
                        return NewDepMacOSEnrollmentProfile(), nil
                }
            }
        }
    }
    return NewDepEnrollmentBaseProfile(), nil
}
// GetAppleIdDisabled gets the appleIdDisabled property value. Indicates if Apple id setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetAppleIdDisabled()(*bool) {
    return m.appleIdDisabled
}
// GetApplePayDisabled gets the applePayDisabled property value. Indicates if Apple pay setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetApplePayDisabled()(*bool) {
    return m.applePayDisabled
}
// GetConfigurationWebUrl gets the configurationWebUrl property value. URL for setup assistant login
func (m *DepEnrollmentBaseProfile) GetConfigurationWebUrl()(*bool) {
    return m.configurationWebUrl
}
// GetDeviceNameTemplate gets the deviceNameTemplate property value. Sets a literal or name pattern.
func (m *DepEnrollmentBaseProfile) GetDeviceNameTemplate()(*string) {
    return m.deviceNameTemplate
}
// GetDiagnosticsDisabled gets the diagnosticsDisabled property value. Indicates if diagnostics setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetDiagnosticsDisabled()(*bool) {
    return m.diagnosticsDisabled
}
// GetDisplayToneSetupDisabled gets the displayToneSetupDisabled property value. Indicates if displaytone setup screen is disabled
func (m *DepEnrollmentBaseProfile) GetDisplayToneSetupDisabled()(*bool) {
    return m.displayToneSetupDisabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DepEnrollmentBaseProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EnrollmentProfile.GetFieldDeserializers()
    res["appleIdDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAppleIdDisabled)
    res["applePayDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetApplePayDisabled)
    res["configurationWebUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetConfigurationWebUrl)
    res["deviceNameTemplate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceNameTemplate)
    res["diagnosticsDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDiagnosticsDisabled)
    res["displayToneSetupDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDisplayToneSetupDisabled)
    res["isDefault"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsDefault)
    res["isMandatory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsMandatory)
    res["locationDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetLocationDisabled)
    res["privacyPaneDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPrivacyPaneDisabled)
    res["profileRemovalDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetProfileRemovalDisabled)
    res["restoreBlocked"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRestoreBlocked)
    res["screenTimeScreenDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetScreenTimeScreenDisabled)
    res["siriDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSiriDisabled)
    res["supervisedModeEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSupervisedModeEnabled)
    res["supportDepartment"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSupportDepartment)
    res["supportPhoneNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSupportPhoneNumber)
    res["termsAndConditionsDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetTermsAndConditionsDisabled)
    res["touchIdDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetTouchIdDisabled)
    return res
}
// GetIsDefault gets the isDefault property value. Indicates if this is the default profile
func (m *DepEnrollmentBaseProfile) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetIsMandatory gets the isMandatory property value. Indicates if the profile is mandatory
func (m *DepEnrollmentBaseProfile) GetIsMandatory()(*bool) {
    return m.isMandatory
}
// GetLocationDisabled gets the locationDisabled property value. Indicates if Location service setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetLocationDisabled()(*bool) {
    return m.locationDisabled
}
// GetPrivacyPaneDisabled gets the privacyPaneDisabled property value. Indicates if privacy screen is disabled
func (m *DepEnrollmentBaseProfile) GetPrivacyPaneDisabled()(*bool) {
    return m.privacyPaneDisabled
}
// GetProfileRemovalDisabled gets the profileRemovalDisabled property value. Indicates if the profile removal option is disabled
func (m *DepEnrollmentBaseProfile) GetProfileRemovalDisabled()(*bool) {
    return m.profileRemovalDisabled
}
// GetRestoreBlocked gets the restoreBlocked property value. Indicates if Restore setup pane is blocked
func (m *DepEnrollmentBaseProfile) GetRestoreBlocked()(*bool) {
    return m.restoreBlocked
}
// GetScreenTimeScreenDisabled gets the screenTimeScreenDisabled property value. Indicates if screen timeout setup is disabled
func (m *DepEnrollmentBaseProfile) GetScreenTimeScreenDisabled()(*bool) {
    return m.screenTimeScreenDisabled
}
// GetSiriDisabled gets the siriDisabled property value. Indicates if siri setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetSiriDisabled()(*bool) {
    return m.siriDisabled
}
// GetSupervisedModeEnabled gets the supervisedModeEnabled property value. Supervised mode, True to enable, false otherwise. See https://learn.microsoft.com/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
func (m *DepEnrollmentBaseProfile) GetSupervisedModeEnabled()(*bool) {
    return m.supervisedModeEnabled
}
// GetSupportDepartment gets the supportDepartment property value. Support department information
func (m *DepEnrollmentBaseProfile) GetSupportDepartment()(*string) {
    return m.supportDepartment
}
// GetSupportPhoneNumber gets the supportPhoneNumber property value. Support phone number
func (m *DepEnrollmentBaseProfile) GetSupportPhoneNumber()(*string) {
    return m.supportPhoneNumber
}
// GetTermsAndConditionsDisabled gets the termsAndConditionsDisabled property value. Indicates if 'Terms and Conditions' setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetTermsAndConditionsDisabled()(*bool) {
    return m.termsAndConditionsDisabled
}
// GetTouchIdDisabled gets the touchIdDisabled property value. Indicates if touch id setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetTouchIdDisabled()(*bool) {
    return m.touchIdDisabled
}
// Serialize serializes information the current object
func (m *DepEnrollmentBaseProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAppleIdDisabled sets the appleIdDisabled property value. Indicates if Apple id setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetAppleIdDisabled(value *bool)() {
    m.appleIdDisabled = value
}
// SetApplePayDisabled sets the applePayDisabled property value. Indicates if Apple pay setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetApplePayDisabled(value *bool)() {
    m.applePayDisabled = value
}
// SetConfigurationWebUrl sets the configurationWebUrl property value. URL for setup assistant login
func (m *DepEnrollmentBaseProfile) SetConfigurationWebUrl(value *bool)() {
    m.configurationWebUrl = value
}
// SetDeviceNameTemplate sets the deviceNameTemplate property value. Sets a literal or name pattern.
func (m *DepEnrollmentBaseProfile) SetDeviceNameTemplate(value *string)() {
    m.deviceNameTemplate = value
}
// SetDiagnosticsDisabled sets the diagnosticsDisabled property value. Indicates if diagnostics setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetDiagnosticsDisabled(value *bool)() {
    m.diagnosticsDisabled = value
}
// SetDisplayToneSetupDisabled sets the displayToneSetupDisabled property value. Indicates if displaytone setup screen is disabled
func (m *DepEnrollmentBaseProfile) SetDisplayToneSetupDisabled(value *bool)() {
    m.displayToneSetupDisabled = value
}
// SetIsDefault sets the isDefault property value. Indicates if this is the default profile
func (m *DepEnrollmentBaseProfile) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetIsMandatory sets the isMandatory property value. Indicates if the profile is mandatory
func (m *DepEnrollmentBaseProfile) SetIsMandatory(value *bool)() {
    m.isMandatory = value
}
// SetLocationDisabled sets the locationDisabled property value. Indicates if Location service setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetLocationDisabled(value *bool)() {
    m.locationDisabled = value
}
// SetPrivacyPaneDisabled sets the privacyPaneDisabled property value. Indicates if privacy screen is disabled
func (m *DepEnrollmentBaseProfile) SetPrivacyPaneDisabled(value *bool)() {
    m.privacyPaneDisabled = value
}
// SetProfileRemovalDisabled sets the profileRemovalDisabled property value. Indicates if the profile removal option is disabled
func (m *DepEnrollmentBaseProfile) SetProfileRemovalDisabled(value *bool)() {
    m.profileRemovalDisabled = value
}
// SetRestoreBlocked sets the restoreBlocked property value. Indicates if Restore setup pane is blocked
func (m *DepEnrollmentBaseProfile) SetRestoreBlocked(value *bool)() {
    m.restoreBlocked = value
}
// SetScreenTimeScreenDisabled sets the screenTimeScreenDisabled property value. Indicates if screen timeout setup is disabled
func (m *DepEnrollmentBaseProfile) SetScreenTimeScreenDisabled(value *bool)() {
    m.screenTimeScreenDisabled = value
}
// SetSiriDisabled sets the siriDisabled property value. Indicates if siri setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetSiriDisabled(value *bool)() {
    m.siriDisabled = value
}
// SetSupervisedModeEnabled sets the supervisedModeEnabled property value. Supervised mode, True to enable, false otherwise. See https://learn.microsoft.com/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
func (m *DepEnrollmentBaseProfile) SetSupervisedModeEnabled(value *bool)() {
    m.supervisedModeEnabled = value
}
// SetSupportDepartment sets the supportDepartment property value. Support department information
func (m *DepEnrollmentBaseProfile) SetSupportDepartment(value *string)() {
    m.supportDepartment = value
}
// SetSupportPhoneNumber sets the supportPhoneNumber property value. Support phone number
func (m *DepEnrollmentBaseProfile) SetSupportPhoneNumber(value *string)() {
    m.supportPhoneNumber = value
}
// SetTermsAndConditionsDisabled sets the termsAndConditionsDisabled property value. Indicates if 'Terms and Conditions' setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetTermsAndConditionsDisabled(value *bool)() {
    m.termsAndConditionsDisabled = value
}
// SetTouchIdDisabled sets the touchIdDisabled property value. Indicates if touch id setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetTouchIdDisabled(value *bool)() {
    m.touchIdDisabled = value
}
