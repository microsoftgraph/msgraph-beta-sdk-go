package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepEnrollmentBaseProfile the DepEnrollmentBaseProfile resource represents an Apple Device Enrollment Program (DEP) enrollment profile. This type of profile must be assigned to Apple DEP serial numbers before the corresponding devices can enroll via DEP.
type DepEnrollmentBaseProfile struct {
    EnrollmentProfile
}
// NewDepEnrollmentBaseProfile instantiates a new depEnrollmentBaseProfile and sets the default values.
func NewDepEnrollmentBaseProfile()(*DepEnrollmentBaseProfile) {
    m := &DepEnrollmentBaseProfile{
        EnrollmentProfile: *NewEnrollmentProfile(),
    }
    odataTypeValue := "#microsoft.graph.depEnrollmentBaseProfile"
    m.SetOdataType(&odataTypeValue)
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
    val, err := m.GetBackingStore().Get("appleIdDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApplePayDisabled gets the applePayDisabled property value. Indicates if Apple pay setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetApplePayDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("applePayDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetConfigurationWebUrl gets the configurationWebUrl property value. URL for setup assistant login
func (m *DepEnrollmentBaseProfile) GetConfigurationWebUrl()(*bool) {
    val, err := m.GetBackingStore().Get("configurationWebUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDeviceNameTemplate gets the deviceNameTemplate property value. Sets a literal or name pattern.
func (m *DepEnrollmentBaseProfile) GetDeviceNameTemplate()(*string) {
    val, err := m.GetBackingStore().Get("deviceNameTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDiagnosticsDisabled gets the diagnosticsDisabled property value. Indicates if diagnostics setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetDiagnosticsDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("diagnosticsDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisplayToneSetupDisabled gets the displayToneSetupDisabled property value. Indicates if displaytone setup screen is disabled
func (m *DepEnrollmentBaseProfile) GetDisplayToneSetupDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("displayToneSetupDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnabledSkipKeys gets the enabledSkipKeys property value. enabledSkipKeys contains all the enabled skip keys as strings
func (m *DepEnrollmentBaseProfile) GetEnabledSkipKeys()([]string) {
    val, err := m.GetBackingStore().Get("enabledSkipKeys")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetEnrollmentTimeAzureAdGroupIds gets the enrollmentTimeAzureAdGroupIds property value. EnrollmentTimeAzureAdGroupIds contains list of enrollment time Azure Group Ids to be associated with profile
func (m *DepEnrollmentBaseProfile) GetEnrollmentTimeAzureAdGroupIds()([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("enrollmentTimeAzureAdGroupIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DepEnrollmentBaseProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EnrollmentProfile.GetFieldDeserializers()
    res["appleIdDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleIdDisabled(val)
        }
        return nil
    }
    res["applePayDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplePayDisabled(val)
        }
        return nil
    }
    res["configurationWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationWebUrl(val)
        }
        return nil
    }
    res["deviceNameTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceNameTemplate(val)
        }
        return nil
    }
    res["diagnosticsDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiagnosticsDisabled(val)
        }
        return nil
    }
    res["displayToneSetupDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayToneSetupDisabled(val)
        }
        return nil
    }
    res["enabledSkipKeys"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetEnabledSkipKeys(res)
        }
        return nil
    }
    res["enrollmentTimeAzureAdGroupIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("uuid")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID))
                }
            }
            m.SetEnrollmentTimeAzureAdGroupIds(res)
        }
        return nil
    }
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isMandatory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMandatory(val)
        }
        return nil
    }
    res["locationDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationDisabled(val)
        }
        return nil
    }
    res["privacyPaneDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyPaneDisabled(val)
        }
        return nil
    }
    res["profileRemovalDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileRemovalDisabled(val)
        }
        return nil
    }
    res["restoreBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestoreBlocked(val)
        }
        return nil
    }
    res["screenTimeScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreenTimeScreenDisabled(val)
        }
        return nil
    }
    res["siriDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiriDisabled(val)
        }
        return nil
    }
    res["supervisedModeEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupervisedModeEnabled(val)
        }
        return nil
    }
    res["supportDepartment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportDepartment(val)
        }
        return nil
    }
    res["supportPhoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportPhoneNumber(val)
        }
        return nil
    }
    res["termsAndConditionsDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsAndConditionsDisabled(val)
        }
        return nil
    }
    res["touchIdDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTouchIdDisabled(val)
        }
        return nil
    }
    res["waitForDeviceConfiguredConfirmation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWaitForDeviceConfiguredConfirmation(val)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. Indicates if this is the default profile
func (m *DepEnrollmentBaseProfile) GetIsDefault()(*bool) {
    val, err := m.GetBackingStore().Get("isDefault")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsMandatory gets the isMandatory property value. Indicates if the profile is mandatory
func (m *DepEnrollmentBaseProfile) GetIsMandatory()(*bool) {
    val, err := m.GetBackingStore().Get("isMandatory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocationDisabled gets the locationDisabled property value. Indicates if Location service setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetLocationDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("locationDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPrivacyPaneDisabled gets the privacyPaneDisabled property value. Indicates if privacy screen is disabled
func (m *DepEnrollmentBaseProfile) GetPrivacyPaneDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("privacyPaneDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetProfileRemovalDisabled gets the profileRemovalDisabled property value. Indicates if the profile removal option is disabled
func (m *DepEnrollmentBaseProfile) GetProfileRemovalDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("profileRemovalDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRestoreBlocked gets the restoreBlocked property value. Indicates if Restore setup pane is blocked
func (m *DepEnrollmentBaseProfile) GetRestoreBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("restoreBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetScreenTimeScreenDisabled gets the screenTimeScreenDisabled property value. Indicates if screen timeout setup is disabled
func (m *DepEnrollmentBaseProfile) GetScreenTimeScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("screenTimeScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSiriDisabled gets the siriDisabled property value. Indicates if siri setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetSiriDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("siriDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSupervisedModeEnabled gets the supervisedModeEnabled property value. Supervised mode, True to enable, false otherwise. See https://learn.microsoft.com/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
func (m *DepEnrollmentBaseProfile) GetSupervisedModeEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("supervisedModeEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSupportDepartment gets the supportDepartment property value. Support department information
func (m *DepEnrollmentBaseProfile) GetSupportDepartment()(*string) {
    val, err := m.GetBackingStore().Get("supportDepartment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSupportPhoneNumber gets the supportPhoneNumber property value. Support phone number
func (m *DepEnrollmentBaseProfile) GetSupportPhoneNumber()(*string) {
    val, err := m.GetBackingStore().Get("supportPhoneNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTermsAndConditionsDisabled gets the termsAndConditionsDisabled property value. Indicates if 'Terms and Conditions' setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetTermsAndConditionsDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("termsAndConditionsDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTouchIdDisabled gets the touchIdDisabled property value. Indicates if touch id setup pane is disabled
func (m *DepEnrollmentBaseProfile) GetTouchIdDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("touchIdDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWaitForDeviceConfiguredConfirmation gets the waitForDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
func (m *DepEnrollmentBaseProfile) GetWaitForDeviceConfiguredConfirmation()(*bool) {
    val, err := m.GetBackingStore().Get("waitForDeviceConfiguredConfirmation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
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
    if m.GetEnabledSkipKeys() != nil {
        err = writer.WriteCollectionOfStringValues("enabledSkipKeys", m.GetEnabledSkipKeys())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentTimeAzureAdGroupIds() != nil {
        err = writer.WriteCollectionOfUUIDValues("enrollmentTimeAzureAdGroupIds", m.GetEnrollmentTimeAzureAdGroupIds())
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
    {
        err = writer.WriteBoolValue("waitForDeviceConfiguredConfirmation", m.GetWaitForDeviceConfiguredConfirmation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppleIdDisabled sets the appleIdDisabled property value. Indicates if Apple id setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetAppleIdDisabled(value *bool)() {
    err := m.GetBackingStore().Set("appleIdDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetApplePayDisabled sets the applePayDisabled property value. Indicates if Apple pay setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetApplePayDisabled(value *bool)() {
    err := m.GetBackingStore().Set("applePayDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationWebUrl sets the configurationWebUrl property value. URL for setup assistant login
func (m *DepEnrollmentBaseProfile) SetConfigurationWebUrl(value *bool)() {
    err := m.GetBackingStore().Set("configurationWebUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceNameTemplate sets the deviceNameTemplate property value. Sets a literal or name pattern.
func (m *DepEnrollmentBaseProfile) SetDeviceNameTemplate(value *string)() {
    err := m.GetBackingStore().Set("deviceNameTemplate", value)
    if err != nil {
        panic(err)
    }
}
// SetDiagnosticsDisabled sets the diagnosticsDisabled property value. Indicates if diagnostics setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetDiagnosticsDisabled(value *bool)() {
    err := m.GetBackingStore().Set("diagnosticsDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayToneSetupDisabled sets the displayToneSetupDisabled property value. Indicates if displaytone setup screen is disabled
func (m *DepEnrollmentBaseProfile) SetDisplayToneSetupDisabled(value *bool)() {
    err := m.GetBackingStore().Set("displayToneSetupDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetEnabledSkipKeys sets the enabledSkipKeys property value. enabledSkipKeys contains all the enabled skip keys as strings
func (m *DepEnrollmentBaseProfile) SetEnabledSkipKeys(value []string)() {
    err := m.GetBackingStore().Set("enabledSkipKeys", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentTimeAzureAdGroupIds sets the enrollmentTimeAzureAdGroupIds property value. EnrollmentTimeAzureAdGroupIds contains list of enrollment time Azure Group Ids to be associated with profile
func (m *DepEnrollmentBaseProfile) SetEnrollmentTimeAzureAdGroupIds(value []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("enrollmentTimeAzureAdGroupIds", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDefault sets the isDefault property value. Indicates if this is the default profile
func (m *DepEnrollmentBaseProfile) SetIsDefault(value *bool)() {
    err := m.GetBackingStore().Set("isDefault", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMandatory sets the isMandatory property value. Indicates if the profile is mandatory
func (m *DepEnrollmentBaseProfile) SetIsMandatory(value *bool)() {
    err := m.GetBackingStore().Set("isMandatory", value)
    if err != nil {
        panic(err)
    }
}
// SetLocationDisabled sets the locationDisabled property value. Indicates if Location service setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetLocationDisabled(value *bool)() {
    err := m.GetBackingStore().Set("locationDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetPrivacyPaneDisabled sets the privacyPaneDisabled property value. Indicates if privacy screen is disabled
func (m *DepEnrollmentBaseProfile) SetPrivacyPaneDisabled(value *bool)() {
    err := m.GetBackingStore().Set("privacyPaneDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileRemovalDisabled sets the profileRemovalDisabled property value. Indicates if the profile removal option is disabled
func (m *DepEnrollmentBaseProfile) SetProfileRemovalDisabled(value *bool)() {
    err := m.GetBackingStore().Set("profileRemovalDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetRestoreBlocked sets the restoreBlocked property value. Indicates if Restore setup pane is blocked
func (m *DepEnrollmentBaseProfile) SetRestoreBlocked(value *bool)() {
    err := m.GetBackingStore().Set("restoreBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetScreenTimeScreenDisabled sets the screenTimeScreenDisabled property value. Indicates if screen timeout setup is disabled
func (m *DepEnrollmentBaseProfile) SetScreenTimeScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("screenTimeScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSiriDisabled sets the siriDisabled property value. Indicates if siri setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetSiriDisabled(value *bool)() {
    err := m.GetBackingStore().Set("siriDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSupervisedModeEnabled sets the supervisedModeEnabled property value. Supervised mode, True to enable, false otherwise. See https://learn.microsoft.com/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
func (m *DepEnrollmentBaseProfile) SetSupervisedModeEnabled(value *bool)() {
    err := m.GetBackingStore().Set("supervisedModeEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportDepartment sets the supportDepartment property value. Support department information
func (m *DepEnrollmentBaseProfile) SetSupportDepartment(value *string)() {
    err := m.GetBackingStore().Set("supportDepartment", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportPhoneNumber sets the supportPhoneNumber property value. Support phone number
func (m *DepEnrollmentBaseProfile) SetSupportPhoneNumber(value *string)() {
    err := m.GetBackingStore().Set("supportPhoneNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetTermsAndConditionsDisabled sets the termsAndConditionsDisabled property value. Indicates if 'Terms and Conditions' setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetTermsAndConditionsDisabled(value *bool)() {
    err := m.GetBackingStore().Set("termsAndConditionsDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetTouchIdDisabled sets the touchIdDisabled property value. Indicates if touch id setup pane is disabled
func (m *DepEnrollmentBaseProfile) SetTouchIdDisabled(value *bool)() {
    err := m.GetBackingStore().Set("touchIdDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetWaitForDeviceConfiguredConfirmation sets the waitForDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
func (m *DepEnrollmentBaseProfile) SetWaitForDeviceConfiguredConfirmation(value *bool)() {
    err := m.GetBackingStore().Set("waitForDeviceConfiguredConfirmation", value)
    if err != nil {
        panic(err)
    }
}
// DepEnrollmentBaseProfileable 
type DepEnrollmentBaseProfileable interface {
    EnrollmentProfileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppleIdDisabled()(*bool)
    GetApplePayDisabled()(*bool)
    GetConfigurationWebUrl()(*bool)
    GetDeviceNameTemplate()(*string)
    GetDiagnosticsDisabled()(*bool)
    GetDisplayToneSetupDisabled()(*bool)
    GetEnabledSkipKeys()([]string)
    GetEnrollmentTimeAzureAdGroupIds()([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
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
    GetWaitForDeviceConfiguredConfirmation()(*bool)
    SetAppleIdDisabled(value *bool)()
    SetApplePayDisabled(value *bool)()
    SetConfigurationWebUrl(value *bool)()
    SetDeviceNameTemplate(value *string)()
    SetDiagnosticsDisabled(value *bool)()
    SetDisplayToneSetupDisabled(value *bool)()
    SetEnabledSkipKeys(value []string)()
    SetEnrollmentTimeAzureAdGroupIds(value []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
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
    SetWaitForDeviceConfiguredConfirmation(value *bool)()
}
