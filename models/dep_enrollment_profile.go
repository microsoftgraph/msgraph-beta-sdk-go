package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepEnrollmentProfile the depEnrollmentProfile resource represents an Apple Device Enrollment Program (DEP) enrollment profile. This type of profile must be assigned to Apple DEP serial numbers before the corresponding devices can enroll via DEP.
type DepEnrollmentProfile struct {
    EnrollmentProfile
}
// NewDepEnrollmentProfile instantiates a new DepEnrollmentProfile and sets the default values.
func NewDepEnrollmentProfile()(*DepEnrollmentProfile) {
    m := &DepEnrollmentProfile{
        EnrollmentProfile: *NewEnrollmentProfile(),
    }
    odataTypeValue := "#microsoft.graph.depEnrollmentProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDepEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDepEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDepEnrollmentProfile(), nil
}
// GetAppleIdDisabled gets the appleIdDisabled property value. Indicates if Apple id setup pane is disabled
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetAppleIdDisabled()(*bool) {
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
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetApplePayDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("applePayDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAwaitDeviceConfiguredConfirmation gets the awaitDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetAwaitDeviceConfiguredConfirmation()(*bool) {
    val, err := m.GetBackingStore().Get("awaitDeviceConfiguredConfirmation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDiagnosticsDisabled gets the diagnosticsDisabled property value. Indicates if diagnostics setup pane is disabled
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetDiagnosticsDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("diagnosticsDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableSharedIPad gets the enableSharedIPad property value. This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetEnableSharedIPad()(*bool) {
    val, err := m.GetBackingStore().Get("enableSharedIPad")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DepEnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["awaitDeviceConfiguredConfirmation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAwaitDeviceConfiguredConfirmation(val)
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
    res["enableSharedIPad"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSharedIPad(val)
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
    res["iTunesPairingMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseITunesPairingMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetITunesPairingMode(val.(*ITunesPairingMode))
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
    res["macOSFileVaultDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOSFileVaultDisabled(val)
        }
        return nil
    }
    res["macOSRegistrationDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOSRegistrationDisabled(val)
        }
        return nil
    }
    res["managementCertificates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementCertificateWithThumbprintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementCertificateWithThumbprintable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagementCertificateWithThumbprintable)
                }
            }
            m.SetManagementCertificates(res)
        }
        return nil
    }
    res["passCodeDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassCodeDisabled(val)
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
    res["restoreFromAndroidDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestoreFromAndroidDisabled(val)
        }
        return nil
    }
    res["sharedIPadMaximumUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedIPadMaximumUserCount(val)
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
    res["zoomDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZoomDisabled(val)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. Indicates if this is the default profile
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetIsDefault()(*bool) {
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
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetIsMandatory()(*bool) {
    val, err := m.GetBackingStore().Get("isMandatory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetITunesPairingMode gets the iTunesPairingMode property value. The iTunesPairingMode property
// returns a *ITunesPairingMode when successful
func (m *DepEnrollmentProfile) GetITunesPairingMode()(*ITunesPairingMode) {
    val, err := m.GetBackingStore().Get("iTunesPairingMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ITunesPairingMode)
    }
    return nil
}
// GetLocationDisabled gets the locationDisabled property value. Indicates if Location service setup pane is disabled
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetLocationDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("locationDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMacOSFileVaultDisabled gets the macOSFileVaultDisabled property value. Indicates if Mac OS file vault is disabled
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetMacOSFileVaultDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("macOSFileVaultDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMacOSRegistrationDisabled gets the macOSRegistrationDisabled property value. Indicates if Mac OS registration is disabled
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetMacOSRegistrationDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("macOSRegistrationDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetManagementCertificates gets the managementCertificates property value. Management certificates for Apple Configurator
// returns a []ManagementCertificateWithThumbprintable when successful
func (m *DepEnrollmentProfile) GetManagementCertificates()([]ManagementCertificateWithThumbprintable) {
    val, err := m.GetBackingStore().Get("managementCertificates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementCertificateWithThumbprintable)
    }
    return nil
}
// GetPassCodeDisabled gets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetPassCodeDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("passCodeDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetProfileRemovalDisabled gets the profileRemovalDisabled property value. Indicates if the profile removal option is disabled
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetProfileRemovalDisabled()(*bool) {
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
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetRestoreBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("restoreBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRestoreFromAndroidDisabled gets the restoreFromAndroidDisabled property value. Indicates if Restore from Android is disabled
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetRestoreFromAndroidDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("restoreFromAndroidDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSharedIPadMaximumUserCount gets the sharedIPadMaximumUserCount property value. This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
// returns a *int32 when successful
func (m *DepEnrollmentProfile) GetSharedIPadMaximumUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("sharedIPadMaximumUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSiriDisabled gets the siriDisabled property value. Indicates if siri setup pane is disabled
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetSiriDisabled()(*bool) {
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
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetSupervisedModeEnabled()(*bool) {
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
// returns a *string when successful
func (m *DepEnrollmentProfile) GetSupportDepartment()(*string) {
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
// returns a *string when successful
func (m *DepEnrollmentProfile) GetSupportPhoneNumber()(*string) {
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
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetTermsAndConditionsDisabled()(*bool) {
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
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetTouchIdDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("touchIdDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetZoomDisabled gets the zoomDisabled property value. Indicates if zoom setup pane is disabled
// returns a *bool when successful
func (m *DepEnrollmentProfile) GetZoomDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("zoomDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DepEnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("awaitDeviceConfiguredConfirmation", m.GetAwaitDeviceConfiguredConfirmation())
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
        err = writer.WriteBoolValue("enableSharedIPad", m.GetEnableSharedIPad())
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
    if m.GetITunesPairingMode() != nil {
        cast := (*m.GetITunesPairingMode()).String()
        err = writer.WriteStringValue("iTunesPairingMode", &cast)
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
        err = writer.WriteBoolValue("macOSFileVaultDisabled", m.GetMacOSFileVaultDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("macOSRegistrationDisabled", m.GetMacOSRegistrationDisabled())
        if err != nil {
            return err
        }
    }
    if m.GetManagementCertificates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementCertificates()))
        for i, v := range m.GetManagementCertificates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("managementCertificates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passCodeDisabled", m.GetPassCodeDisabled())
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
        err = writer.WriteBoolValue("restoreFromAndroidDisabled", m.GetRestoreFromAndroidDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("sharedIPadMaximumUserCount", m.GetSharedIPadMaximumUserCount())
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
        err = writer.WriteBoolValue("zoomDisabled", m.GetZoomDisabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppleIdDisabled sets the appleIdDisabled property value. Indicates if Apple id setup pane is disabled
func (m *DepEnrollmentProfile) SetAppleIdDisabled(value *bool)() {
    err := m.GetBackingStore().Set("appleIdDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetApplePayDisabled sets the applePayDisabled property value. Indicates if Apple pay setup pane is disabled
func (m *DepEnrollmentProfile) SetApplePayDisabled(value *bool)() {
    err := m.GetBackingStore().Set("applePayDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetAwaitDeviceConfiguredConfirmation sets the awaitDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
func (m *DepEnrollmentProfile) SetAwaitDeviceConfiguredConfirmation(value *bool)() {
    err := m.GetBackingStore().Set("awaitDeviceConfiguredConfirmation", value)
    if err != nil {
        panic(err)
    }
}
// SetDiagnosticsDisabled sets the diagnosticsDisabled property value. Indicates if diagnostics setup pane is disabled
func (m *DepEnrollmentProfile) SetDiagnosticsDisabled(value *bool)() {
    err := m.GetBackingStore().Set("diagnosticsDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableSharedIPad sets the enableSharedIPad property value. This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
func (m *DepEnrollmentProfile) SetEnableSharedIPad(value *bool)() {
    err := m.GetBackingStore().Set("enableSharedIPad", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDefault sets the isDefault property value. Indicates if this is the default profile
func (m *DepEnrollmentProfile) SetIsDefault(value *bool)() {
    err := m.GetBackingStore().Set("isDefault", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMandatory sets the isMandatory property value. Indicates if the profile is mandatory
func (m *DepEnrollmentProfile) SetIsMandatory(value *bool)() {
    err := m.GetBackingStore().Set("isMandatory", value)
    if err != nil {
        panic(err)
    }
}
// SetITunesPairingMode sets the iTunesPairingMode property value. The iTunesPairingMode property
func (m *DepEnrollmentProfile) SetITunesPairingMode(value *ITunesPairingMode)() {
    err := m.GetBackingStore().Set("iTunesPairingMode", value)
    if err != nil {
        panic(err)
    }
}
// SetLocationDisabled sets the locationDisabled property value. Indicates if Location service setup pane is disabled
func (m *DepEnrollmentProfile) SetLocationDisabled(value *bool)() {
    err := m.GetBackingStore().Set("locationDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetMacOSFileVaultDisabled sets the macOSFileVaultDisabled property value. Indicates if Mac OS file vault is disabled
func (m *DepEnrollmentProfile) SetMacOSFileVaultDisabled(value *bool)() {
    err := m.GetBackingStore().Set("macOSFileVaultDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetMacOSRegistrationDisabled sets the macOSRegistrationDisabled property value. Indicates if Mac OS registration is disabled
func (m *DepEnrollmentProfile) SetMacOSRegistrationDisabled(value *bool)() {
    err := m.GetBackingStore().Set("macOSRegistrationDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementCertificates sets the managementCertificates property value. Management certificates for Apple Configurator
func (m *DepEnrollmentProfile) SetManagementCertificates(value []ManagementCertificateWithThumbprintable)() {
    err := m.GetBackingStore().Set("managementCertificates", value)
    if err != nil {
        panic(err)
    }
}
// SetPassCodeDisabled sets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepEnrollmentProfile) SetPassCodeDisabled(value *bool)() {
    err := m.GetBackingStore().Set("passCodeDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileRemovalDisabled sets the profileRemovalDisabled property value. Indicates if the profile removal option is disabled
func (m *DepEnrollmentProfile) SetProfileRemovalDisabled(value *bool)() {
    err := m.GetBackingStore().Set("profileRemovalDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetRestoreBlocked sets the restoreBlocked property value. Indicates if Restore setup pane is blocked
func (m *DepEnrollmentProfile) SetRestoreBlocked(value *bool)() {
    err := m.GetBackingStore().Set("restoreBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetRestoreFromAndroidDisabled sets the restoreFromAndroidDisabled property value. Indicates if Restore from Android is disabled
func (m *DepEnrollmentProfile) SetRestoreFromAndroidDisabled(value *bool)() {
    err := m.GetBackingStore().Set("restoreFromAndroidDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSharedIPadMaximumUserCount sets the sharedIPadMaximumUserCount property value. This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
func (m *DepEnrollmentProfile) SetSharedIPadMaximumUserCount(value *int32)() {
    err := m.GetBackingStore().Set("sharedIPadMaximumUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSiriDisabled sets the siriDisabled property value. Indicates if siri setup pane is disabled
func (m *DepEnrollmentProfile) SetSiriDisabled(value *bool)() {
    err := m.GetBackingStore().Set("siriDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSupervisedModeEnabled sets the supervisedModeEnabled property value. Supervised mode, True to enable, false otherwise. See https://learn.microsoft.com/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
func (m *DepEnrollmentProfile) SetSupervisedModeEnabled(value *bool)() {
    err := m.GetBackingStore().Set("supervisedModeEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportDepartment sets the supportDepartment property value. Support department information
func (m *DepEnrollmentProfile) SetSupportDepartment(value *string)() {
    err := m.GetBackingStore().Set("supportDepartment", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportPhoneNumber sets the supportPhoneNumber property value. Support phone number
func (m *DepEnrollmentProfile) SetSupportPhoneNumber(value *string)() {
    err := m.GetBackingStore().Set("supportPhoneNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetTermsAndConditionsDisabled sets the termsAndConditionsDisabled property value. Indicates if 'Terms and Conditions' setup pane is disabled
func (m *DepEnrollmentProfile) SetTermsAndConditionsDisabled(value *bool)() {
    err := m.GetBackingStore().Set("termsAndConditionsDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetTouchIdDisabled sets the touchIdDisabled property value. Indicates if touch id setup pane is disabled
func (m *DepEnrollmentProfile) SetTouchIdDisabled(value *bool)() {
    err := m.GetBackingStore().Set("touchIdDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetZoomDisabled sets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepEnrollmentProfile) SetZoomDisabled(value *bool)() {
    err := m.GetBackingStore().Set("zoomDisabled", value)
    if err != nil {
        panic(err)
    }
}
type DepEnrollmentProfileable interface {
    EnrollmentProfileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppleIdDisabled()(*bool)
    GetApplePayDisabled()(*bool)
    GetAwaitDeviceConfiguredConfirmation()(*bool)
    GetDiagnosticsDisabled()(*bool)
    GetEnableSharedIPad()(*bool)
    GetIsDefault()(*bool)
    GetIsMandatory()(*bool)
    GetITunesPairingMode()(*ITunesPairingMode)
    GetLocationDisabled()(*bool)
    GetMacOSFileVaultDisabled()(*bool)
    GetMacOSRegistrationDisabled()(*bool)
    GetManagementCertificates()([]ManagementCertificateWithThumbprintable)
    GetPassCodeDisabled()(*bool)
    GetProfileRemovalDisabled()(*bool)
    GetRestoreBlocked()(*bool)
    GetRestoreFromAndroidDisabled()(*bool)
    GetSharedIPadMaximumUserCount()(*int32)
    GetSiriDisabled()(*bool)
    GetSupervisedModeEnabled()(*bool)
    GetSupportDepartment()(*string)
    GetSupportPhoneNumber()(*string)
    GetTermsAndConditionsDisabled()(*bool)
    GetTouchIdDisabled()(*bool)
    GetZoomDisabled()(*bool)
    SetAppleIdDisabled(value *bool)()
    SetApplePayDisabled(value *bool)()
    SetAwaitDeviceConfiguredConfirmation(value *bool)()
    SetDiagnosticsDisabled(value *bool)()
    SetEnableSharedIPad(value *bool)()
    SetIsDefault(value *bool)()
    SetIsMandatory(value *bool)()
    SetITunesPairingMode(value *ITunesPairingMode)()
    SetLocationDisabled(value *bool)()
    SetMacOSFileVaultDisabled(value *bool)()
    SetMacOSRegistrationDisabled(value *bool)()
    SetManagementCertificates(value []ManagementCertificateWithThumbprintable)()
    SetPassCodeDisabled(value *bool)()
    SetProfileRemovalDisabled(value *bool)()
    SetRestoreBlocked(value *bool)()
    SetRestoreFromAndroidDisabled(value *bool)()
    SetSharedIPadMaximumUserCount(value *int32)()
    SetSiriDisabled(value *bool)()
    SetSupervisedModeEnabled(value *bool)()
    SetSupportDepartment(value *string)()
    SetSupportPhoneNumber(value *string)()
    SetTermsAndConditionsDisabled(value *bool)()
    SetTouchIdDisabled(value *bool)()
    SetZoomDisabled(value *bool)()
}
