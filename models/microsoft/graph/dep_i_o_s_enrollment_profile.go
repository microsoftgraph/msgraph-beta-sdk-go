package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DepIOSEnrollmentProfile 
type DepIOSEnrollmentProfile struct {
    DepEnrollmentBaseProfile
    // Indicates if Apperance screen is disabled
    appearanceScreenDisabled *bool;
    // Indicates if the device will need to wait for configured confirmation
    awaitDeviceConfiguredConfirmation *bool;
    // Carrier URL for activating device eSIM.
    carrierActivationUrl *string;
    // If set, indicates which Vpp token should be used to deploy the Company Portal w/ device licensing. 'enableAuthenticationViaCompanyPortal' must be set in order for this property to be set.
    companyPortalVppTokenId *string;
    // Indicates if Device To Device Migration is disabled
    deviceToDeviceMigrationDisabled *bool;
    // This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
    enableSharedIPad *bool;
    // Tells the device to enable single app mode and apply app-lock during enrollment. Default is false. 'enableAuthenticationViaCompanyPortal' and 'companyPortalVppTokenId' must be set for this property to be set.
    enableSingleAppEnrollmentMode *bool;
    // Indicates if Express Language screen is disabled
    expressLanguageScreenDisabled *bool;
    // Indicates if temporary sessions is enabled
    forceTemporarySession *bool;
    // Indicates if home button sensitivity screen is disabled
    homeButtonScreenDisabled *bool;
    // Indicates if iMessage and FaceTime screen is disabled
    iMessageAndFaceTimeScreenDisabled *bool;
    // Indicates the iTunes pairing mode. Possible values are: disallow, allow, requiresCertificate.
    iTunesPairingMode *ITunesPairingMode;
    // Management certificates for Apple Configurator
    managementCertificates []ManagementCertificateWithThumbprint;
    // Indicates if onboarding setup screen is disabled
    onBoardingScreenDisabled *bool;
    // Indicates if Passcode setup pane is disabled
    passCodeDisabled *bool;
    // Indicates timeout before locked screen requires the user to enter the device passocde to unlock it
    passcodeLockGracePeriodInSeconds *int32;
    // Indicates if Preferred language screen is disabled
    preferredLanguageScreenDisabled *bool;
    // Indicates if Weclome screen is disabled
    restoreCompletedScreenDisabled *bool;
    // Indicates if Restore from Android is disabled
    restoreFromAndroidDisabled *bool;
    // This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
    sharedIPadMaximumUserCount *int32;
    // Indicates if the SIMSetup screen is disabled
    simSetupScreenDisabled *bool;
    // Indicates if the mandatory sofware update screen is disabled
    softwareUpdateScreenDisabled *bool;
    // Indicates timeout of temporary session
    temporarySessionTimeoutInSeconds *int32;
    // Indicates if Weclome screen is disabled
    updateCompleteScreenDisabled *bool;
    // Indicates that this apple device is designated to support 'shared device mode' scenarios. This is distinct from the 'shared iPad' scenario. See Shared iOS and iPadOS devices.
    userlessSharedAadModeEnabled *bool;
    // Indicates timeout of temporary session
    userSessionTimeoutInSeconds *int32;
    // Indicates if the watch migration screen is disabled
    watchMigrationScreenDisabled *bool;
    // Indicates if Weclome screen is disabled
    welcomeScreenDisabled *bool;
    // Indicates if zoom setup pane is disabled
    zoomDisabled *bool;
}
// NewDepIOSEnrollmentProfile instantiates a new depIOSEnrollmentProfile and sets the default values.
func NewDepIOSEnrollmentProfile()(*DepIOSEnrollmentProfile) {
    m := &DepIOSEnrollmentProfile{
        DepEnrollmentBaseProfile: *NewDepEnrollmentBaseProfile(),
    }
    return m
}
// GetAppearanceScreenDisabled gets the appearanceScreenDisabled property value. Indicates if Apperance screen is disabled
func (m *DepIOSEnrollmentProfile) GetAppearanceScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appearanceScreenDisabled
    }
}
// GetAwaitDeviceConfiguredConfirmation gets the awaitDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
func (m *DepIOSEnrollmentProfile) GetAwaitDeviceConfiguredConfirmation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.awaitDeviceConfiguredConfirmation
    }
}
// GetCarrierActivationUrl gets the carrierActivationUrl property value. Carrier URL for activating device eSIM.
func (m *DepIOSEnrollmentProfile) GetCarrierActivationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.carrierActivationUrl
    }
}
// GetCompanyPortalVppTokenId gets the companyPortalVppTokenId property value. If set, indicates which Vpp token should be used to deploy the Company Portal w/ device licensing. 'enableAuthenticationViaCompanyPortal' must be set in order for this property to be set.
func (m *DepIOSEnrollmentProfile) GetCompanyPortalVppTokenId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.companyPortalVppTokenId
    }
}
// GetDeviceToDeviceMigrationDisabled gets the deviceToDeviceMigrationDisabled property value. Indicates if Device To Device Migration is disabled
func (m *DepIOSEnrollmentProfile) GetDeviceToDeviceMigrationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceToDeviceMigrationDisabled
    }
}
// GetEnableSharedIPad gets the enableSharedIPad property value. This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
func (m *DepIOSEnrollmentProfile) GetEnableSharedIPad()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableSharedIPad
    }
}
// GetEnableSingleAppEnrollmentMode gets the enableSingleAppEnrollmentMode property value. Tells the device to enable single app mode and apply app-lock during enrollment. Default is false. 'enableAuthenticationViaCompanyPortal' and 'companyPortalVppTokenId' must be set for this property to be set.
func (m *DepIOSEnrollmentProfile) GetEnableSingleAppEnrollmentMode()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableSingleAppEnrollmentMode
    }
}
// GetExpressLanguageScreenDisabled gets the expressLanguageScreenDisabled property value. Indicates if Express Language screen is disabled
func (m *DepIOSEnrollmentProfile) GetExpressLanguageScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.expressLanguageScreenDisabled
    }
}
// GetForceTemporarySession gets the forceTemporarySession property value. Indicates if temporary sessions is enabled
func (m *DepIOSEnrollmentProfile) GetForceTemporarySession()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.forceTemporarySession
    }
}
// GetHomeButtonScreenDisabled gets the homeButtonScreenDisabled property value. Indicates if home button sensitivity screen is disabled
func (m *DepIOSEnrollmentProfile) GetHomeButtonScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.homeButtonScreenDisabled
    }
}
// GetIMessageAndFaceTimeScreenDisabled gets the iMessageAndFaceTimeScreenDisabled property value. Indicates if iMessage and FaceTime screen is disabled
func (m *DepIOSEnrollmentProfile) GetIMessageAndFaceTimeScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iMessageAndFaceTimeScreenDisabled
    }
}
// GetITunesPairingMode gets the iTunesPairingMode property value. Indicates the iTunes pairing mode. Possible values are: disallow, allow, requiresCertificate.
func (m *DepIOSEnrollmentProfile) GetITunesPairingMode()(*ITunesPairingMode) {
    if m == nil {
        return nil
    } else {
        return m.iTunesPairingMode
    }
}
// GetManagementCertificates gets the managementCertificates property value. Management certificates for Apple Configurator
func (m *DepIOSEnrollmentProfile) GetManagementCertificates()([]ManagementCertificateWithThumbprint) {
    if m == nil {
        return nil
    } else {
        return m.managementCertificates
    }
}
// GetOnBoardingScreenDisabled gets the onBoardingScreenDisabled property value. Indicates if onboarding setup screen is disabled
func (m *DepIOSEnrollmentProfile) GetOnBoardingScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.onBoardingScreenDisabled
    }
}
// GetPassCodeDisabled gets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepIOSEnrollmentProfile) GetPassCodeDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passCodeDisabled
    }
}
// GetPasscodeLockGracePeriodInSeconds gets the passcodeLockGracePeriodInSeconds property value. Indicates timeout before locked screen requires the user to enter the device passocde to unlock it
func (m *DepIOSEnrollmentProfile) GetPasscodeLockGracePeriodInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodeLockGracePeriodInSeconds
    }
}
// GetPreferredLanguageScreenDisabled gets the preferredLanguageScreenDisabled property value. Indicates if Preferred language screen is disabled
func (m *DepIOSEnrollmentProfile) GetPreferredLanguageScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.preferredLanguageScreenDisabled
    }
}
// GetRestoreCompletedScreenDisabled gets the restoreCompletedScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) GetRestoreCompletedScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.restoreCompletedScreenDisabled
    }
}
// GetRestoreFromAndroidDisabled gets the restoreFromAndroidDisabled property value. Indicates if Restore from Android is disabled
func (m *DepIOSEnrollmentProfile) GetRestoreFromAndroidDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.restoreFromAndroidDisabled
    }
}
// GetSharedIPadMaximumUserCount gets the sharedIPadMaximumUserCount property value. This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
func (m *DepIOSEnrollmentProfile) GetSharedIPadMaximumUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sharedIPadMaximumUserCount
    }
}
// GetSimSetupScreenDisabled gets the simSetupScreenDisabled property value. Indicates if the SIMSetup screen is disabled
func (m *DepIOSEnrollmentProfile) GetSimSetupScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.simSetupScreenDisabled
    }
}
// GetSoftwareUpdateScreenDisabled gets the softwareUpdateScreenDisabled property value. Indicates if the mandatory sofware update screen is disabled
func (m *DepIOSEnrollmentProfile) GetSoftwareUpdateScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.softwareUpdateScreenDisabled
    }
}
// GetTemporarySessionTimeoutInSeconds gets the temporarySessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) GetTemporarySessionTimeoutInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.temporarySessionTimeoutInSeconds
    }
}
// GetUpdateCompleteScreenDisabled gets the updateCompleteScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) GetUpdateCompleteScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.updateCompleteScreenDisabled
    }
}
// GetUserlessSharedAadModeEnabled gets the userlessSharedAadModeEnabled property value. Indicates that this apple device is designated to support 'shared device mode' scenarios. This is distinct from the 'shared iPad' scenario. See Shared iOS and iPadOS devices.
func (m *DepIOSEnrollmentProfile) GetUserlessSharedAadModeEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.userlessSharedAadModeEnabled
    }
}
// GetUserSessionTimeoutInSeconds gets the userSessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) GetUserSessionTimeoutInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.userSessionTimeoutInSeconds
    }
}
// GetWatchMigrationScreenDisabled gets the watchMigrationScreenDisabled property value. Indicates if the watch migration screen is disabled
func (m *DepIOSEnrollmentProfile) GetWatchMigrationScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.watchMigrationScreenDisabled
    }
}
// GetWelcomeScreenDisabled gets the welcomeScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) GetWelcomeScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.welcomeScreenDisabled
    }
}
// GetZoomDisabled gets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepIOSEnrollmentProfile) GetZoomDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.zoomDisabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DepIOSEnrollmentProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DepEnrollmentBaseProfile.GetFieldDeserializers()
    res["appearanceScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppearanceScreenDisabled(val)
        }
        return nil
    }
    res["awaitDeviceConfiguredConfirmation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAwaitDeviceConfiguredConfirmation(val)
        }
        return nil
    }
    res["carrierActivationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCarrierActivationUrl(val)
        }
        return nil
    }
    res["companyPortalVppTokenId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyPortalVppTokenId(val)
        }
        return nil
    }
    res["deviceToDeviceMigrationDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceToDeviceMigrationDisabled(val)
        }
        return nil
    }
    res["enableSharedIPad"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSharedIPad(val)
        }
        return nil
    }
    res["enableSingleAppEnrollmentMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSingleAppEnrollmentMode(val)
        }
        return nil
    }
    res["expressLanguageScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpressLanguageScreenDisabled(val)
        }
        return nil
    }
    res["forceTemporarySession"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForceTemporarySession(val)
        }
        return nil
    }
    res["homeButtonScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeButtonScreenDisabled(val)
        }
        return nil
    }
    res["iMessageAndFaceTimeScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIMessageAndFaceTimeScreenDisabled(val)
        }
        return nil
    }
    res["iTunesPairingMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseITunesPairingMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetITunesPairingMode(val.(*ITunesPairingMode))
        }
        return nil
    }
    res["managementCertificates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementCertificateWithThumbprint() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementCertificateWithThumbprint, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementCertificateWithThumbprint))
            }
            m.SetManagementCertificates(res)
        }
        return nil
    }
    res["onBoardingScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnBoardingScreenDisabled(val)
        }
        return nil
    }
    res["passCodeDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassCodeDisabled(val)
        }
        return nil
    }
    res["passcodeLockGracePeriodInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeLockGracePeriodInSeconds(val)
        }
        return nil
    }
    res["preferredLanguageScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguageScreenDisabled(val)
        }
        return nil
    }
    res["restoreCompletedScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestoreCompletedScreenDisabled(val)
        }
        return nil
    }
    res["restoreFromAndroidDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestoreFromAndroidDisabled(val)
        }
        return nil
    }
    res["sharedIPadMaximumUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedIPadMaximumUserCount(val)
        }
        return nil
    }
    res["simSetupScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimSetupScreenDisabled(val)
        }
        return nil
    }
    res["softwareUpdateScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdateScreenDisabled(val)
        }
        return nil
    }
    res["temporarySessionTimeoutInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemporarySessionTimeoutInSeconds(val)
        }
        return nil
    }
    res["updateCompleteScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateCompleteScreenDisabled(val)
        }
        return nil
    }
    res["userlessSharedAadModeEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserlessSharedAadModeEnabled(val)
        }
        return nil
    }
    res["userSessionTimeoutInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserSessionTimeoutInSeconds(val)
        }
        return nil
    }
    res["watchMigrationScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWatchMigrationScreenDisabled(val)
        }
        return nil
    }
    res["welcomeScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWelcomeScreenDisabled(val)
        }
        return nil
    }
    res["zoomDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *DepIOSEnrollmentProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DepIOSEnrollmentProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementCertificates()))
        for i, v := range m.GetManagementCertificates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    if m != nil {
        m.appearanceScreenDisabled = value
    }
}
// SetAwaitDeviceConfiguredConfirmation sets the awaitDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
func (m *DepIOSEnrollmentProfile) SetAwaitDeviceConfiguredConfirmation(value *bool)() {
    if m != nil {
        m.awaitDeviceConfiguredConfirmation = value
    }
}
// SetCarrierActivationUrl sets the carrierActivationUrl property value. Carrier URL for activating device eSIM.
func (m *DepIOSEnrollmentProfile) SetCarrierActivationUrl(value *string)() {
    if m != nil {
        m.carrierActivationUrl = value
    }
}
// SetCompanyPortalVppTokenId sets the companyPortalVppTokenId property value. If set, indicates which Vpp token should be used to deploy the Company Portal w/ device licensing. 'enableAuthenticationViaCompanyPortal' must be set in order for this property to be set.
func (m *DepIOSEnrollmentProfile) SetCompanyPortalVppTokenId(value *string)() {
    if m != nil {
        m.companyPortalVppTokenId = value
    }
}
// SetDeviceToDeviceMigrationDisabled sets the deviceToDeviceMigrationDisabled property value. Indicates if Device To Device Migration is disabled
func (m *DepIOSEnrollmentProfile) SetDeviceToDeviceMigrationDisabled(value *bool)() {
    if m != nil {
        m.deviceToDeviceMigrationDisabled = value
    }
}
// SetEnableSharedIPad sets the enableSharedIPad property value. This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
func (m *DepIOSEnrollmentProfile) SetEnableSharedIPad(value *bool)() {
    if m != nil {
        m.enableSharedIPad = value
    }
}
// SetEnableSingleAppEnrollmentMode sets the enableSingleAppEnrollmentMode property value. Tells the device to enable single app mode and apply app-lock during enrollment. Default is false. 'enableAuthenticationViaCompanyPortal' and 'companyPortalVppTokenId' must be set for this property to be set.
func (m *DepIOSEnrollmentProfile) SetEnableSingleAppEnrollmentMode(value *bool)() {
    if m != nil {
        m.enableSingleAppEnrollmentMode = value
    }
}
// SetExpressLanguageScreenDisabled sets the expressLanguageScreenDisabled property value. Indicates if Express Language screen is disabled
func (m *DepIOSEnrollmentProfile) SetExpressLanguageScreenDisabled(value *bool)() {
    if m != nil {
        m.expressLanguageScreenDisabled = value
    }
}
// SetForceTemporarySession sets the forceTemporarySession property value. Indicates if temporary sessions is enabled
func (m *DepIOSEnrollmentProfile) SetForceTemporarySession(value *bool)() {
    if m != nil {
        m.forceTemporarySession = value
    }
}
// SetHomeButtonScreenDisabled sets the homeButtonScreenDisabled property value. Indicates if home button sensitivity screen is disabled
func (m *DepIOSEnrollmentProfile) SetHomeButtonScreenDisabled(value *bool)() {
    if m != nil {
        m.homeButtonScreenDisabled = value
    }
}
// SetIMessageAndFaceTimeScreenDisabled sets the iMessageAndFaceTimeScreenDisabled property value. Indicates if iMessage and FaceTime screen is disabled
func (m *DepIOSEnrollmentProfile) SetIMessageAndFaceTimeScreenDisabled(value *bool)() {
    if m != nil {
        m.iMessageAndFaceTimeScreenDisabled = value
    }
}
// SetITunesPairingMode sets the iTunesPairingMode property value. Indicates the iTunes pairing mode. Possible values are: disallow, allow, requiresCertificate.
func (m *DepIOSEnrollmentProfile) SetITunesPairingMode(value *ITunesPairingMode)() {
    if m != nil {
        m.iTunesPairingMode = value
    }
}
// SetManagementCertificates sets the managementCertificates property value. Management certificates for Apple Configurator
func (m *DepIOSEnrollmentProfile) SetManagementCertificates(value []ManagementCertificateWithThumbprint)() {
    if m != nil {
        m.managementCertificates = value
    }
}
// SetOnBoardingScreenDisabled sets the onBoardingScreenDisabled property value. Indicates if onboarding setup screen is disabled
func (m *DepIOSEnrollmentProfile) SetOnBoardingScreenDisabled(value *bool)() {
    if m != nil {
        m.onBoardingScreenDisabled = value
    }
}
// SetPassCodeDisabled sets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepIOSEnrollmentProfile) SetPassCodeDisabled(value *bool)() {
    if m != nil {
        m.passCodeDisabled = value
    }
}
// SetPasscodeLockGracePeriodInSeconds sets the passcodeLockGracePeriodInSeconds property value. Indicates timeout before locked screen requires the user to enter the device passocde to unlock it
func (m *DepIOSEnrollmentProfile) SetPasscodeLockGracePeriodInSeconds(value *int32)() {
    if m != nil {
        m.passcodeLockGracePeriodInSeconds = value
    }
}
// SetPreferredLanguageScreenDisabled sets the preferredLanguageScreenDisabled property value. Indicates if Preferred language screen is disabled
func (m *DepIOSEnrollmentProfile) SetPreferredLanguageScreenDisabled(value *bool)() {
    if m != nil {
        m.preferredLanguageScreenDisabled = value
    }
}
// SetRestoreCompletedScreenDisabled sets the restoreCompletedScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) SetRestoreCompletedScreenDisabled(value *bool)() {
    if m != nil {
        m.restoreCompletedScreenDisabled = value
    }
}
// SetRestoreFromAndroidDisabled sets the restoreFromAndroidDisabled property value. Indicates if Restore from Android is disabled
func (m *DepIOSEnrollmentProfile) SetRestoreFromAndroidDisabled(value *bool)() {
    if m != nil {
        m.restoreFromAndroidDisabled = value
    }
}
// SetSharedIPadMaximumUserCount sets the sharedIPadMaximumUserCount property value. This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
func (m *DepIOSEnrollmentProfile) SetSharedIPadMaximumUserCount(value *int32)() {
    if m != nil {
        m.sharedIPadMaximumUserCount = value
    }
}
// SetSimSetupScreenDisabled sets the simSetupScreenDisabled property value. Indicates if the SIMSetup screen is disabled
func (m *DepIOSEnrollmentProfile) SetSimSetupScreenDisabled(value *bool)() {
    if m != nil {
        m.simSetupScreenDisabled = value
    }
}
// SetSoftwareUpdateScreenDisabled sets the softwareUpdateScreenDisabled property value. Indicates if the mandatory sofware update screen is disabled
func (m *DepIOSEnrollmentProfile) SetSoftwareUpdateScreenDisabled(value *bool)() {
    if m != nil {
        m.softwareUpdateScreenDisabled = value
    }
}
// SetTemporarySessionTimeoutInSeconds sets the temporarySessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) SetTemporarySessionTimeoutInSeconds(value *int32)() {
    if m != nil {
        m.temporarySessionTimeoutInSeconds = value
    }
}
// SetUpdateCompleteScreenDisabled sets the updateCompleteScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) SetUpdateCompleteScreenDisabled(value *bool)() {
    if m != nil {
        m.updateCompleteScreenDisabled = value
    }
}
// SetUserlessSharedAadModeEnabled sets the userlessSharedAadModeEnabled property value. Indicates that this apple device is designated to support 'shared device mode' scenarios. This is distinct from the 'shared iPad' scenario. See Shared iOS and iPadOS devices.
func (m *DepIOSEnrollmentProfile) SetUserlessSharedAadModeEnabled(value *bool)() {
    if m != nil {
        m.userlessSharedAadModeEnabled = value
    }
}
// SetUserSessionTimeoutInSeconds sets the userSessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) SetUserSessionTimeoutInSeconds(value *int32)() {
    if m != nil {
        m.userSessionTimeoutInSeconds = value
    }
}
// SetWatchMigrationScreenDisabled sets the watchMigrationScreenDisabled property value. Indicates if the watch migration screen is disabled
func (m *DepIOSEnrollmentProfile) SetWatchMigrationScreenDisabled(value *bool)() {
    if m != nil {
        m.watchMigrationScreenDisabled = value
    }
}
// SetWelcomeScreenDisabled sets the welcomeScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) SetWelcomeScreenDisabled(value *bool)() {
    if m != nil {
        m.welcomeScreenDisabled = value
    }
}
// SetZoomDisabled sets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepIOSEnrollmentProfile) SetZoomDisabled(value *bool)() {
    if m != nil {
        m.zoomDisabled = value
    }
}
