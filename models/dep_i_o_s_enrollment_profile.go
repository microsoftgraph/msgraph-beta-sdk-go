package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepIOSEnrollmentProfile 
type DepIOSEnrollmentProfile struct {
    DepEnrollmentBaseProfile
}
// NewDepIOSEnrollmentProfile instantiates a new DepIOSEnrollmentProfile and sets the default values.
func NewDepIOSEnrollmentProfile()(*DepIOSEnrollmentProfile) {
    m := &DepIOSEnrollmentProfile{
        DepEnrollmentBaseProfile: *NewDepEnrollmentBaseProfile(),
    }
    odataTypeValue := "#microsoft.graph.depIOSEnrollmentProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDepIOSEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDepIOSEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDepIOSEnrollmentProfile(), nil
}
// GetAppearanceScreenDisabled gets the appearanceScreenDisabled property value. Indicates if Apperance screen is disabled
func (m *DepIOSEnrollmentProfile) GetAppearanceScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("appearanceScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAwaitDeviceConfiguredConfirmation gets the awaitDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
func (m *DepIOSEnrollmentProfile) GetAwaitDeviceConfiguredConfirmation()(*bool) {
    val, err := m.GetBackingStore().Get("awaitDeviceConfiguredConfirmation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCarrierActivationUrl gets the carrierActivationUrl property value. Carrier URL for activating device eSIM.
func (m *DepIOSEnrollmentProfile) GetCarrierActivationUrl()(*string) {
    val, err := m.GetBackingStore().Get("carrierActivationUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCompanyPortalVppTokenId gets the companyPortalVppTokenId property value. If set, indicates which Vpp token should be used to deploy the Company Portal w/ device licensing. 'enableAuthenticationViaCompanyPortal' must be set in order for this property to be set.
func (m *DepIOSEnrollmentProfile) GetCompanyPortalVppTokenId()(*string) {
    val, err := m.GetBackingStore().Get("companyPortalVppTokenId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceToDeviceMigrationDisabled gets the deviceToDeviceMigrationDisabled property value. Indicates if Device To Device Migration is disabled
func (m *DepIOSEnrollmentProfile) GetDeviceToDeviceMigrationDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("deviceToDeviceMigrationDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableSharedIPad gets the enableSharedIPad property value. This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
func (m *DepIOSEnrollmentProfile) GetEnableSharedIPad()(*bool) {
    val, err := m.GetBackingStore().Get("enableSharedIPad")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableSingleAppEnrollmentMode gets the enableSingleAppEnrollmentMode property value. Tells the device to enable single app mode and apply app-lock during enrollment. Default is false. 'enableAuthenticationViaCompanyPortal' and 'companyPortalVppTokenId' must be set for this property to be set.
func (m *DepIOSEnrollmentProfile) GetEnableSingleAppEnrollmentMode()(*bool) {
    val, err := m.GetBackingStore().Get("enableSingleAppEnrollmentMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetExpressLanguageScreenDisabled gets the expressLanguageScreenDisabled property value. Indicates if Express Language screen is disabled
func (m *DepIOSEnrollmentProfile) GetExpressLanguageScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("expressLanguageScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DepIOSEnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DepEnrollmentBaseProfile.GetFieldDeserializers()
    res["appearanceScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppearanceScreenDisabled(val)
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
    res["carrierActivationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCarrierActivationUrl(val)
        }
        return nil
    }
    res["companyPortalVppTokenId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyPortalVppTokenId(val)
        }
        return nil
    }
    res["deviceToDeviceMigrationDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceToDeviceMigrationDisabled(val)
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
    res["enableSingleAppEnrollmentMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSingleAppEnrollmentMode(val)
        }
        return nil
    }
    res["expressLanguageScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpressLanguageScreenDisabled(val)
        }
        return nil
    }
    res["forceTemporarySession"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForceTemporarySession(val)
        }
        return nil
    }
    res["homeButtonScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeButtonScreenDisabled(val)
        }
        return nil
    }
    res["iMessageAndFaceTimeScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIMessageAndFaceTimeScreenDisabled(val)
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
    res["onBoardingScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnBoardingScreenDisabled(val)
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
    res["passcodeLockGracePeriodInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeLockGracePeriodInSeconds(val)
        }
        return nil
    }
    res["preferredLanguageScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguageScreenDisabled(val)
        }
        return nil
    }
    res["restoreCompletedScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestoreCompletedScreenDisabled(val)
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
    res["simSetupScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimSetupScreenDisabled(val)
        }
        return nil
    }
    res["softwareUpdateScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdateScreenDisabled(val)
        }
        return nil
    }
    res["temporarySessionTimeoutInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemporarySessionTimeoutInSeconds(val)
        }
        return nil
    }
    res["updateCompleteScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateCompleteScreenDisabled(val)
        }
        return nil
    }
    res["userlessSharedAadModeEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserlessSharedAadModeEnabled(val)
        }
        return nil
    }
    res["userSessionTimeoutInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserSessionTimeoutInSeconds(val)
        }
        return nil
    }
    res["watchMigrationScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWatchMigrationScreenDisabled(val)
        }
        return nil
    }
    res["welcomeScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWelcomeScreenDisabled(val)
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
// GetForceTemporarySession gets the forceTemporarySession property value. Indicates if temporary sessions is enabled
func (m *DepIOSEnrollmentProfile) GetForceTemporarySession()(*bool) {
    val, err := m.GetBackingStore().Get("forceTemporarySession")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetHomeButtonScreenDisabled gets the homeButtonScreenDisabled property value. Indicates if home button sensitivity screen is disabled
func (m *DepIOSEnrollmentProfile) GetHomeButtonScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("homeButtonScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIMessageAndFaceTimeScreenDisabled gets the iMessageAndFaceTimeScreenDisabled property value. Indicates if iMessage and FaceTime screen is disabled
func (m *DepIOSEnrollmentProfile) GetIMessageAndFaceTimeScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("iMessageAndFaceTimeScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetITunesPairingMode gets the iTunesPairingMode property value. The iTunesPairingMode property
func (m *DepIOSEnrollmentProfile) GetITunesPairingMode()(*ITunesPairingMode) {
    val, err := m.GetBackingStore().Get("iTunesPairingMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ITunesPairingMode)
    }
    return nil
}
// GetManagementCertificates gets the managementCertificates property value. Management certificates for Apple Configurator
func (m *DepIOSEnrollmentProfile) GetManagementCertificates()([]ManagementCertificateWithThumbprintable) {
    val, err := m.GetBackingStore().Get("managementCertificates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementCertificateWithThumbprintable)
    }
    return nil
}
// GetOnBoardingScreenDisabled gets the onBoardingScreenDisabled property value. Indicates if onboarding setup screen is disabled
func (m *DepIOSEnrollmentProfile) GetOnBoardingScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("onBoardingScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPassCodeDisabled gets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepIOSEnrollmentProfile) GetPassCodeDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("passCodeDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasscodeLockGracePeriodInSeconds gets the passcodeLockGracePeriodInSeconds property value. Indicates timeout before locked screen requires the user to enter the device passocde to unlock it
func (m *DepIOSEnrollmentProfile) GetPasscodeLockGracePeriodInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("passcodeLockGracePeriodInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPreferredLanguageScreenDisabled gets the preferredLanguageScreenDisabled property value. Indicates if Preferred language screen is disabled
func (m *DepIOSEnrollmentProfile) GetPreferredLanguageScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("preferredLanguageScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRestoreCompletedScreenDisabled gets the restoreCompletedScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) GetRestoreCompletedScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("restoreCompletedScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRestoreFromAndroidDisabled gets the restoreFromAndroidDisabled property value. Indicates if Restore from Android is disabled
func (m *DepIOSEnrollmentProfile) GetRestoreFromAndroidDisabled()(*bool) {
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
func (m *DepIOSEnrollmentProfile) GetSharedIPadMaximumUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("sharedIPadMaximumUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSimSetupScreenDisabled gets the simSetupScreenDisabled property value. Indicates if the SIMSetup screen is disabled
func (m *DepIOSEnrollmentProfile) GetSimSetupScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("simSetupScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSoftwareUpdateScreenDisabled gets the softwareUpdateScreenDisabled property value. Indicates if the mandatory sofware update screen is disabled
func (m *DepIOSEnrollmentProfile) GetSoftwareUpdateScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("softwareUpdateScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTemporarySessionTimeoutInSeconds gets the temporarySessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) GetTemporarySessionTimeoutInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("temporarySessionTimeoutInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUpdateCompleteScreenDisabled gets the updateCompleteScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) GetUpdateCompleteScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("updateCompleteScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUserlessSharedAadModeEnabled gets the userlessSharedAadModeEnabled property value. Indicates that this apple device is designated to support 'shared device mode' scenarios. This is distinct from the 'shared iPad' scenario. See https://learn.microsoft.com/mem/intune/enrollment/device-enrollment-shared-ios
func (m *DepIOSEnrollmentProfile) GetUserlessSharedAadModeEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("userlessSharedAadModeEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUserSessionTimeoutInSeconds gets the userSessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) GetUserSessionTimeoutInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("userSessionTimeoutInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWatchMigrationScreenDisabled gets the watchMigrationScreenDisabled property value. Indicates if the watch migration screen is disabled
func (m *DepIOSEnrollmentProfile) GetWatchMigrationScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("watchMigrationScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWelcomeScreenDisabled gets the welcomeScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) GetWelcomeScreenDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("welcomeScreenDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetZoomDisabled gets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepIOSEnrollmentProfile) GetZoomDisabled()(*bool) {
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
func (m *DepIOSEnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DepEnrollmentBaseProfile.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("appearanceScreenDisabled", m.GetAppearanceScreenDisabled())
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
        err = writer.WriteStringValue("carrierActivationUrl", m.GetCarrierActivationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyPortalVppTokenId", m.GetCompanyPortalVppTokenId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceToDeviceMigrationDisabled", m.GetDeviceToDeviceMigrationDisabled())
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
        err = writer.WriteBoolValue("enableSingleAppEnrollmentMode", m.GetEnableSingleAppEnrollmentMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("expressLanguageScreenDisabled", m.GetExpressLanguageScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("forceTemporarySession", m.GetForceTemporarySession())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("homeButtonScreenDisabled", m.GetHomeButtonScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iMessageAndFaceTimeScreenDisabled", m.GetIMessageAndFaceTimeScreenDisabled())
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
        err = writer.WriteBoolValue("onBoardingScreenDisabled", m.GetOnBoardingScreenDisabled())
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
        err = writer.WriteInt32Value("passcodeLockGracePeriodInSeconds", m.GetPasscodeLockGracePeriodInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("preferredLanguageScreenDisabled", m.GetPreferredLanguageScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("restoreCompletedScreenDisabled", m.GetRestoreCompletedScreenDisabled())
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
        err = writer.WriteBoolValue("simSetupScreenDisabled", m.GetSimSetupScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("softwareUpdateScreenDisabled", m.GetSoftwareUpdateScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("temporarySessionTimeoutInSeconds", m.GetTemporarySessionTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("updateCompleteScreenDisabled", m.GetUpdateCompleteScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("userlessSharedAadModeEnabled", m.GetUserlessSharedAadModeEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("userSessionTimeoutInSeconds", m.GetUserSessionTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("watchMigrationScreenDisabled", m.GetWatchMigrationScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("welcomeScreenDisabled", m.GetWelcomeScreenDisabled())
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
// SetAppearanceScreenDisabled sets the appearanceScreenDisabled property value. Indicates if Apperance screen is disabled
func (m *DepIOSEnrollmentProfile) SetAppearanceScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("appearanceScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetAwaitDeviceConfiguredConfirmation sets the awaitDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
func (m *DepIOSEnrollmentProfile) SetAwaitDeviceConfiguredConfirmation(value *bool)() {
    err := m.GetBackingStore().Set("awaitDeviceConfiguredConfirmation", value)
    if err != nil {
        panic(err)
    }
}
// SetCarrierActivationUrl sets the carrierActivationUrl property value. Carrier URL for activating device eSIM.
func (m *DepIOSEnrollmentProfile) SetCarrierActivationUrl(value *string)() {
    err := m.GetBackingStore().Set("carrierActivationUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCompanyPortalVppTokenId sets the companyPortalVppTokenId property value. If set, indicates which Vpp token should be used to deploy the Company Portal w/ device licensing. 'enableAuthenticationViaCompanyPortal' must be set in order for this property to be set.
func (m *DepIOSEnrollmentProfile) SetCompanyPortalVppTokenId(value *string)() {
    err := m.GetBackingStore().Set("companyPortalVppTokenId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceToDeviceMigrationDisabled sets the deviceToDeviceMigrationDisabled property value. Indicates if Device To Device Migration is disabled
func (m *DepIOSEnrollmentProfile) SetDeviceToDeviceMigrationDisabled(value *bool)() {
    err := m.GetBackingStore().Set("deviceToDeviceMigrationDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableSharedIPad sets the enableSharedIPad property value. This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
func (m *DepIOSEnrollmentProfile) SetEnableSharedIPad(value *bool)() {
    err := m.GetBackingStore().Set("enableSharedIPad", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableSingleAppEnrollmentMode sets the enableSingleAppEnrollmentMode property value. Tells the device to enable single app mode and apply app-lock during enrollment. Default is false. 'enableAuthenticationViaCompanyPortal' and 'companyPortalVppTokenId' must be set for this property to be set.
func (m *DepIOSEnrollmentProfile) SetEnableSingleAppEnrollmentMode(value *bool)() {
    err := m.GetBackingStore().Set("enableSingleAppEnrollmentMode", value)
    if err != nil {
        panic(err)
    }
}
// SetExpressLanguageScreenDisabled sets the expressLanguageScreenDisabled property value. Indicates if Express Language screen is disabled
func (m *DepIOSEnrollmentProfile) SetExpressLanguageScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("expressLanguageScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetForceTemporarySession sets the forceTemporarySession property value. Indicates if temporary sessions is enabled
func (m *DepIOSEnrollmentProfile) SetForceTemporarySession(value *bool)() {
    err := m.GetBackingStore().Set("forceTemporarySession", value)
    if err != nil {
        panic(err)
    }
}
// SetHomeButtonScreenDisabled sets the homeButtonScreenDisabled property value. Indicates if home button sensitivity screen is disabled
func (m *DepIOSEnrollmentProfile) SetHomeButtonScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("homeButtonScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIMessageAndFaceTimeScreenDisabled sets the iMessageAndFaceTimeScreenDisabled property value. Indicates if iMessage and FaceTime screen is disabled
func (m *DepIOSEnrollmentProfile) SetIMessageAndFaceTimeScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("iMessageAndFaceTimeScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetITunesPairingMode sets the iTunesPairingMode property value. The iTunesPairingMode property
func (m *DepIOSEnrollmentProfile) SetITunesPairingMode(value *ITunesPairingMode)() {
    err := m.GetBackingStore().Set("iTunesPairingMode", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementCertificates sets the managementCertificates property value. Management certificates for Apple Configurator
func (m *DepIOSEnrollmentProfile) SetManagementCertificates(value []ManagementCertificateWithThumbprintable)() {
    err := m.GetBackingStore().Set("managementCertificates", value)
    if err != nil {
        panic(err)
    }
}
// SetOnBoardingScreenDisabled sets the onBoardingScreenDisabled property value. Indicates if onboarding setup screen is disabled
func (m *DepIOSEnrollmentProfile) SetOnBoardingScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("onBoardingScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetPassCodeDisabled sets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepIOSEnrollmentProfile) SetPassCodeDisabled(value *bool)() {
    err := m.GetBackingStore().Set("passCodeDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetPasscodeLockGracePeriodInSeconds sets the passcodeLockGracePeriodInSeconds property value. Indicates timeout before locked screen requires the user to enter the device passocde to unlock it
func (m *DepIOSEnrollmentProfile) SetPasscodeLockGracePeriodInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("passcodeLockGracePeriodInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetPreferredLanguageScreenDisabled sets the preferredLanguageScreenDisabled property value. Indicates if Preferred language screen is disabled
func (m *DepIOSEnrollmentProfile) SetPreferredLanguageScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("preferredLanguageScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetRestoreCompletedScreenDisabled sets the restoreCompletedScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) SetRestoreCompletedScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("restoreCompletedScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetRestoreFromAndroidDisabled sets the restoreFromAndroidDisabled property value. Indicates if Restore from Android is disabled
func (m *DepIOSEnrollmentProfile) SetRestoreFromAndroidDisabled(value *bool)() {
    err := m.GetBackingStore().Set("restoreFromAndroidDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSharedIPadMaximumUserCount sets the sharedIPadMaximumUserCount property value. This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
func (m *DepIOSEnrollmentProfile) SetSharedIPadMaximumUserCount(value *int32)() {
    err := m.GetBackingStore().Set("sharedIPadMaximumUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSimSetupScreenDisabled sets the simSetupScreenDisabled property value. Indicates if the SIMSetup screen is disabled
func (m *DepIOSEnrollmentProfile) SetSimSetupScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("simSetupScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSoftwareUpdateScreenDisabled sets the softwareUpdateScreenDisabled property value. Indicates if the mandatory sofware update screen is disabled
func (m *DepIOSEnrollmentProfile) SetSoftwareUpdateScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("softwareUpdateScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetTemporarySessionTimeoutInSeconds sets the temporarySessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) SetTemporarySessionTimeoutInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("temporarySessionTimeoutInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateCompleteScreenDisabled sets the updateCompleteScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) SetUpdateCompleteScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("updateCompleteScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetUserlessSharedAadModeEnabled sets the userlessSharedAadModeEnabled property value. Indicates that this apple device is designated to support 'shared device mode' scenarios. This is distinct from the 'shared iPad' scenario. See https://learn.microsoft.com/mem/intune/enrollment/device-enrollment-shared-ios
func (m *DepIOSEnrollmentProfile) SetUserlessSharedAadModeEnabled(value *bool)() {
    err := m.GetBackingStore().Set("userlessSharedAadModeEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetUserSessionTimeoutInSeconds sets the userSessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) SetUserSessionTimeoutInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("userSessionTimeoutInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetWatchMigrationScreenDisabled sets the watchMigrationScreenDisabled property value. Indicates if the watch migration screen is disabled
func (m *DepIOSEnrollmentProfile) SetWatchMigrationScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("watchMigrationScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetWelcomeScreenDisabled sets the welcomeScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) SetWelcomeScreenDisabled(value *bool)() {
    err := m.GetBackingStore().Set("welcomeScreenDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetZoomDisabled sets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepIOSEnrollmentProfile) SetZoomDisabled(value *bool)() {
    err := m.GetBackingStore().Set("zoomDisabled", value)
    if err != nil {
        panic(err)
    }
}
// DepIOSEnrollmentProfileable 
type DepIOSEnrollmentProfileable interface {
    DepEnrollmentBaseProfileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppearanceScreenDisabled()(*bool)
    GetAwaitDeviceConfiguredConfirmation()(*bool)
    GetCarrierActivationUrl()(*string)
    GetCompanyPortalVppTokenId()(*string)
    GetDeviceToDeviceMigrationDisabled()(*bool)
    GetEnableSharedIPad()(*bool)
    GetEnableSingleAppEnrollmentMode()(*bool)
    GetExpressLanguageScreenDisabled()(*bool)
    GetForceTemporarySession()(*bool)
    GetHomeButtonScreenDisabled()(*bool)
    GetIMessageAndFaceTimeScreenDisabled()(*bool)
    GetITunesPairingMode()(*ITunesPairingMode)
    GetManagementCertificates()([]ManagementCertificateWithThumbprintable)
    GetOnBoardingScreenDisabled()(*bool)
    GetPassCodeDisabled()(*bool)
    GetPasscodeLockGracePeriodInSeconds()(*int32)
    GetPreferredLanguageScreenDisabled()(*bool)
    GetRestoreCompletedScreenDisabled()(*bool)
    GetRestoreFromAndroidDisabled()(*bool)
    GetSharedIPadMaximumUserCount()(*int32)
    GetSimSetupScreenDisabled()(*bool)
    GetSoftwareUpdateScreenDisabled()(*bool)
    GetTemporarySessionTimeoutInSeconds()(*int32)
    GetUpdateCompleteScreenDisabled()(*bool)
    GetUserlessSharedAadModeEnabled()(*bool)
    GetUserSessionTimeoutInSeconds()(*int32)
    GetWatchMigrationScreenDisabled()(*bool)
    GetWelcomeScreenDisabled()(*bool)
    GetZoomDisabled()(*bool)
    SetAppearanceScreenDisabled(value *bool)()
    SetAwaitDeviceConfiguredConfirmation(value *bool)()
    SetCarrierActivationUrl(value *string)()
    SetCompanyPortalVppTokenId(value *string)()
    SetDeviceToDeviceMigrationDisabled(value *bool)()
    SetEnableSharedIPad(value *bool)()
    SetEnableSingleAppEnrollmentMode(value *bool)()
    SetExpressLanguageScreenDisabled(value *bool)()
    SetForceTemporarySession(value *bool)()
    SetHomeButtonScreenDisabled(value *bool)()
    SetIMessageAndFaceTimeScreenDisabled(value *bool)()
    SetITunesPairingMode(value *ITunesPairingMode)()
    SetManagementCertificates(value []ManagementCertificateWithThumbprintable)()
    SetOnBoardingScreenDisabled(value *bool)()
    SetPassCodeDisabled(value *bool)()
    SetPasscodeLockGracePeriodInSeconds(value *int32)()
    SetPreferredLanguageScreenDisabled(value *bool)()
    SetRestoreCompletedScreenDisabled(value *bool)()
    SetRestoreFromAndroidDisabled(value *bool)()
    SetSharedIPadMaximumUserCount(value *int32)()
    SetSimSetupScreenDisabled(value *bool)()
    SetSoftwareUpdateScreenDisabled(value *bool)()
    SetTemporarySessionTimeoutInSeconds(value *int32)()
    SetUpdateCompleteScreenDisabled(value *bool)()
    SetUserlessSharedAadModeEnabled(value *bool)()
    SetUserSessionTimeoutInSeconds(value *int32)()
    SetWatchMigrationScreenDisabled(value *bool)()
    SetWelcomeScreenDisabled(value *bool)()
    SetZoomDisabled(value *bool)()
}
