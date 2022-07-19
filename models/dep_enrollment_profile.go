package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepEnrollmentProfile 
type DepEnrollmentProfile struct {
    EnrollmentProfile
    // Indicates if Apple id setup pane is disabled
    appleIdDisabled *bool
    // Indicates if Apple pay setup pane is disabled
    applePayDisabled *bool
    // Indicates if the device will need to wait for configured confirmation
    awaitDeviceConfiguredConfirmation *bool
    // Indicates if diagnostics setup pane is disabled
    diagnosticsDisabled *bool
    // This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
    enableSharedIPad *bool
    // Indicates if this is the default profile
    isDefault *bool
    // Indicates if the profile is mandatory
    isMandatory *bool
    // The iTunesPairingMode property
    iTunesPairingMode *ITunesPairingMode
    // Indicates if Location service setup pane is disabled
    locationDisabled *bool
    // Indicates if Mac OS file vault is disabled
    macOSFileVaultDisabled *bool
    // Indicates if Mac OS registration is disabled
    macOSRegistrationDisabled *bool
    // Management certificates for Apple Configurator
    managementCertificates []ManagementCertificateWithThumbprintable
    // Indicates if Passcode setup pane is disabled
    passCodeDisabled *bool
    // Indicates if the profile removal option is disabled
    profileRemovalDisabled *bool
    // Indicates if Restore setup pane is blocked
    restoreBlocked *bool
    // Indicates if Restore from Android is disabled
    restoreFromAndroidDisabled *bool
    // This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
    sharedIPadMaximumUserCount *int32
    // Indicates if siri setup pane is disabled
    siriDisabled *bool
    // Supervised mode, True to enable, false otherwise. See https://docs.microsoft.com/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
    supervisedModeEnabled *bool
    // Support department information
    supportDepartment *string
    // Support phone number
    supportPhoneNumber *string
    // Indicates if 'Terms and Conditions' setup pane is disabled
    termsAndConditionsDisabled *bool
    // Indicates if touch id setup pane is disabled
    touchIdDisabled *bool
    // Indicates if zoom setup pane is disabled
    zoomDisabled *bool
}
// NewDepEnrollmentProfile instantiates a new DepEnrollmentProfile and sets the default values.
func NewDepEnrollmentProfile()(*DepEnrollmentProfile) {
    m := &DepEnrollmentProfile{
        EnrollmentProfile: *NewEnrollmentProfile(),
    }
    odataTypeValue := "#microsoft.graph.depEnrollmentProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDepEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDepEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDepEnrollmentProfile(), nil
}
// GetAppleIdDisabled gets the appleIdDisabled property value. Indicates if Apple id setup pane is disabled
func (m *DepEnrollmentProfile) GetAppleIdDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appleIdDisabled
    }
}
// GetApplePayDisabled gets the applePayDisabled property value. Indicates if Apple pay setup pane is disabled
func (m *DepEnrollmentProfile) GetApplePayDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applePayDisabled
    }
}
// GetAwaitDeviceConfiguredConfirmation gets the awaitDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
func (m *DepEnrollmentProfile) GetAwaitDeviceConfiguredConfirmation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.awaitDeviceConfiguredConfirmation
    }
}
// GetDiagnosticsDisabled gets the diagnosticsDisabled property value. Indicates if diagnostics setup pane is disabled
func (m *DepEnrollmentProfile) GetDiagnosticsDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.diagnosticsDisabled
    }
}
// GetEnableSharedIPad gets the enableSharedIPad property value. This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
func (m *DepEnrollmentProfile) GetEnableSharedIPad()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableSharedIPad
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
                res[i] = v.(ManagementCertificateWithThumbprintable)
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
func (m *DepEnrollmentProfile) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetIsMandatory gets the isMandatory property value. Indicates if the profile is mandatory
func (m *DepEnrollmentProfile) GetIsMandatory()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMandatory
    }
}
// GetITunesPairingMode gets the iTunesPairingMode property value. The iTunesPairingMode property
func (m *DepEnrollmentProfile) GetITunesPairingMode()(*ITunesPairingMode) {
    if m == nil {
        return nil
    } else {
        return m.iTunesPairingMode
    }
}
// GetLocationDisabled gets the locationDisabled property value. Indicates if Location service setup pane is disabled
func (m *DepEnrollmentProfile) GetLocationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.locationDisabled
    }
}
// GetMacOSFileVaultDisabled gets the macOSFileVaultDisabled property value. Indicates if Mac OS file vault is disabled
func (m *DepEnrollmentProfile) GetMacOSFileVaultDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.macOSFileVaultDisabled
    }
}
// GetMacOSRegistrationDisabled gets the macOSRegistrationDisabled property value. Indicates if Mac OS registration is disabled
func (m *DepEnrollmentProfile) GetMacOSRegistrationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.macOSRegistrationDisabled
    }
}
// GetManagementCertificates gets the managementCertificates property value. Management certificates for Apple Configurator
func (m *DepEnrollmentProfile) GetManagementCertificates()([]ManagementCertificateWithThumbprintable) {
    if m == nil {
        return nil
    } else {
        return m.managementCertificates
    }
}
// GetPassCodeDisabled gets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepEnrollmentProfile) GetPassCodeDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passCodeDisabled
    }
}
// GetProfileRemovalDisabled gets the profileRemovalDisabled property value. Indicates if the profile removal option is disabled
func (m *DepEnrollmentProfile) GetProfileRemovalDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.profileRemovalDisabled
    }
}
// GetRestoreBlocked gets the restoreBlocked property value. Indicates if Restore setup pane is blocked
func (m *DepEnrollmentProfile) GetRestoreBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.restoreBlocked
    }
}
// GetRestoreFromAndroidDisabled gets the restoreFromAndroidDisabled property value. Indicates if Restore from Android is disabled
func (m *DepEnrollmentProfile) GetRestoreFromAndroidDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.restoreFromAndroidDisabled
    }
}
// GetSharedIPadMaximumUserCount gets the sharedIPadMaximumUserCount property value. This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
func (m *DepEnrollmentProfile) GetSharedIPadMaximumUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sharedIPadMaximumUserCount
    }
}
// GetSiriDisabled gets the siriDisabled property value. Indicates if siri setup pane is disabled
func (m *DepEnrollmentProfile) GetSiriDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.siriDisabled
    }
}
// GetSupervisedModeEnabled gets the supervisedModeEnabled property value. Supervised mode, True to enable, false otherwise. See https://docs.microsoft.com/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
func (m *DepEnrollmentProfile) GetSupervisedModeEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.supervisedModeEnabled
    }
}
// GetSupportDepartment gets the supportDepartment property value. Support department information
func (m *DepEnrollmentProfile) GetSupportDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportDepartment
    }
}
// GetSupportPhoneNumber gets the supportPhoneNumber property value. Support phone number
func (m *DepEnrollmentProfile) GetSupportPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportPhoneNumber
    }
}
// GetTermsAndConditionsDisabled gets the termsAndConditionsDisabled property value. Indicates if 'Terms and Conditions' setup pane is disabled
func (m *DepEnrollmentProfile) GetTermsAndConditionsDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.termsAndConditionsDisabled
    }
}
// GetTouchIdDisabled gets the touchIdDisabled property value. Indicates if touch id setup pane is disabled
func (m *DepEnrollmentProfile) GetTouchIdDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.touchIdDisabled
    }
}
// GetZoomDisabled gets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepEnrollmentProfile) GetZoomDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.zoomDisabled
    }
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m != nil {
        m.appleIdDisabled = value
    }
}
// SetApplePayDisabled sets the applePayDisabled property value. Indicates if Apple pay setup pane is disabled
func (m *DepEnrollmentProfile) SetApplePayDisabled(value *bool)() {
    if m != nil {
        m.applePayDisabled = value
    }
}
// SetAwaitDeviceConfiguredConfirmation sets the awaitDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
func (m *DepEnrollmentProfile) SetAwaitDeviceConfiguredConfirmation(value *bool)() {
    if m != nil {
        m.awaitDeviceConfiguredConfirmation = value
    }
}
// SetDiagnosticsDisabled sets the diagnosticsDisabled property value. Indicates if diagnostics setup pane is disabled
func (m *DepEnrollmentProfile) SetDiagnosticsDisabled(value *bool)() {
    if m != nil {
        m.diagnosticsDisabled = value
    }
}
// SetEnableSharedIPad sets the enableSharedIPad property value. This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
func (m *DepEnrollmentProfile) SetEnableSharedIPad(value *bool)() {
    if m != nil {
        m.enableSharedIPad = value
    }
}
// SetIsDefault sets the isDefault property value. Indicates if this is the default profile
func (m *DepEnrollmentProfile) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetIsMandatory sets the isMandatory property value. Indicates if the profile is mandatory
func (m *DepEnrollmentProfile) SetIsMandatory(value *bool)() {
    if m != nil {
        m.isMandatory = value
    }
}
// SetITunesPairingMode sets the iTunesPairingMode property value. The iTunesPairingMode property
func (m *DepEnrollmentProfile) SetITunesPairingMode(value *ITunesPairingMode)() {
    if m != nil {
        m.iTunesPairingMode = value
    }
}
// SetLocationDisabled sets the locationDisabled property value. Indicates if Location service setup pane is disabled
func (m *DepEnrollmentProfile) SetLocationDisabled(value *bool)() {
    if m != nil {
        m.locationDisabled = value
    }
}
// SetMacOSFileVaultDisabled sets the macOSFileVaultDisabled property value. Indicates if Mac OS file vault is disabled
func (m *DepEnrollmentProfile) SetMacOSFileVaultDisabled(value *bool)() {
    if m != nil {
        m.macOSFileVaultDisabled = value
    }
}
// SetMacOSRegistrationDisabled sets the macOSRegistrationDisabled property value. Indicates if Mac OS registration is disabled
func (m *DepEnrollmentProfile) SetMacOSRegistrationDisabled(value *bool)() {
    if m != nil {
        m.macOSRegistrationDisabled = value
    }
}
// SetManagementCertificates sets the managementCertificates property value. Management certificates for Apple Configurator
func (m *DepEnrollmentProfile) SetManagementCertificates(value []ManagementCertificateWithThumbprintable)() {
    if m != nil {
        m.managementCertificates = value
    }
}
// SetPassCodeDisabled sets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepEnrollmentProfile) SetPassCodeDisabled(value *bool)() {
    if m != nil {
        m.passCodeDisabled = value
    }
}
// SetProfileRemovalDisabled sets the profileRemovalDisabled property value. Indicates if the profile removal option is disabled
func (m *DepEnrollmentProfile) SetProfileRemovalDisabled(value *bool)() {
    if m != nil {
        m.profileRemovalDisabled = value
    }
}
// SetRestoreBlocked sets the restoreBlocked property value. Indicates if Restore setup pane is blocked
func (m *DepEnrollmentProfile) SetRestoreBlocked(value *bool)() {
    if m != nil {
        m.restoreBlocked = value
    }
}
// SetRestoreFromAndroidDisabled sets the restoreFromAndroidDisabled property value. Indicates if Restore from Android is disabled
func (m *DepEnrollmentProfile) SetRestoreFromAndroidDisabled(value *bool)() {
    if m != nil {
        m.restoreFromAndroidDisabled = value
    }
}
// SetSharedIPadMaximumUserCount sets the sharedIPadMaximumUserCount property value. This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
func (m *DepEnrollmentProfile) SetSharedIPadMaximumUserCount(value *int32)() {
    if m != nil {
        m.sharedIPadMaximumUserCount = value
    }
}
// SetSiriDisabled sets the siriDisabled property value. Indicates if siri setup pane is disabled
func (m *DepEnrollmentProfile) SetSiriDisabled(value *bool)() {
    if m != nil {
        m.siriDisabled = value
    }
}
// SetSupervisedModeEnabled sets the supervisedModeEnabled property value. Supervised mode, True to enable, false otherwise. See https://docs.microsoft.com/intune/deploy-use/enroll-devices-in-microsoft-intune for additional information.
func (m *DepEnrollmentProfile) SetSupervisedModeEnabled(value *bool)() {
    if m != nil {
        m.supervisedModeEnabled = value
    }
}
// SetSupportDepartment sets the supportDepartment property value. Support department information
func (m *DepEnrollmentProfile) SetSupportDepartment(value *string)() {
    if m != nil {
        m.supportDepartment = value
    }
}
// SetSupportPhoneNumber sets the supportPhoneNumber property value. Support phone number
func (m *DepEnrollmentProfile) SetSupportPhoneNumber(value *string)() {
    if m != nil {
        m.supportPhoneNumber = value
    }
}
// SetTermsAndConditionsDisabled sets the termsAndConditionsDisabled property value. Indicates if 'Terms and Conditions' setup pane is disabled
func (m *DepEnrollmentProfile) SetTermsAndConditionsDisabled(value *bool)() {
    if m != nil {
        m.termsAndConditionsDisabled = value
    }
}
// SetTouchIdDisabled sets the touchIdDisabled property value. Indicates if touch id setup pane is disabled
func (m *DepEnrollmentProfile) SetTouchIdDisabled(value *bool)() {
    if m != nil {
        m.touchIdDisabled = value
    }
}
// SetZoomDisabled sets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepEnrollmentProfile) SetZoomDisabled(value *bool)() {
    if m != nil {
        m.zoomDisabled = value
    }
}
