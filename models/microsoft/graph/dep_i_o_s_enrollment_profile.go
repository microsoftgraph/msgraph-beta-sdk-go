package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
    // Indicates timeout of temporary session
    userSessionTimeoutInSeconds *int32;
    // Indicates if the watch migration screen is disabled
    watchMigrationScreenDisabled *bool;
    // Indicates if Weclome screen is disabled
    welcomeScreenDisabled *bool;
    // Indicates if zoom setup pane is disabled
    zoomDisabled *bool;
}
// Instantiates a new depIOSEnrollmentProfile and sets the default values.
func NewDepIOSEnrollmentProfile()(*DepIOSEnrollmentProfile) {
    m := &DepIOSEnrollmentProfile{
        DepEnrollmentBaseProfile: *NewDepEnrollmentBaseProfile(),
    }
    return m
}
// Gets the appearanceScreenDisabled property value. Indicates if Apperance screen is disabled
func (m *DepIOSEnrollmentProfile) GetAppearanceScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appearanceScreenDisabled
    }
}
// Gets the awaitDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
func (m *DepIOSEnrollmentProfile) GetAwaitDeviceConfiguredConfirmation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.awaitDeviceConfiguredConfirmation
    }
}
// Gets the carrierActivationUrl property value. Carrier URL for activating device eSIM.
func (m *DepIOSEnrollmentProfile) GetCarrierActivationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.carrierActivationUrl
    }
}
// Gets the companyPortalVppTokenId property value. If set, indicates which Vpp token should be used to deploy the Company Portal w/ device licensing. 'enableAuthenticationViaCompanyPortal' must be set in order for this property to be set.
func (m *DepIOSEnrollmentProfile) GetCompanyPortalVppTokenId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.companyPortalVppTokenId
    }
}
// Gets the deviceToDeviceMigrationDisabled property value. Indicates if Device To Device Migration is disabled
func (m *DepIOSEnrollmentProfile) GetDeviceToDeviceMigrationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceToDeviceMigrationDisabled
    }
}
// Gets the enableSharedIPad property value. This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
func (m *DepIOSEnrollmentProfile) GetEnableSharedIPad()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableSharedIPad
    }
}
// Gets the enableSingleAppEnrollmentMode property value. Tells the device to enable single app mode and apply app-lock during enrollment. Default is false. 'enableAuthenticationViaCompanyPortal' and 'companyPortalVppTokenId' must be set for this property to be set.
func (m *DepIOSEnrollmentProfile) GetEnableSingleAppEnrollmentMode()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableSingleAppEnrollmentMode
    }
}
// Gets the expressLanguageScreenDisabled property value. Indicates if Express Language screen is disabled
func (m *DepIOSEnrollmentProfile) GetExpressLanguageScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.expressLanguageScreenDisabled
    }
}
// Gets the forceTemporarySession property value. Indicates if temporary sessions is enabled
func (m *DepIOSEnrollmentProfile) GetForceTemporarySession()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.forceTemporarySession
    }
}
// Gets the homeButtonScreenDisabled property value. Indicates if home button sensitivity screen is disabled
func (m *DepIOSEnrollmentProfile) GetHomeButtonScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.homeButtonScreenDisabled
    }
}
// Gets the iMessageAndFaceTimeScreenDisabled property value. Indicates if iMessage and FaceTime screen is disabled
func (m *DepIOSEnrollmentProfile) GetIMessageAndFaceTimeScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iMessageAndFaceTimeScreenDisabled
    }
}
// Gets the iTunesPairingMode property value. Indicates the iTunes pairing mode. Possible values are: disallow, allow, requiresCertificate.
func (m *DepIOSEnrollmentProfile) GetITunesPairingMode()(*ITunesPairingMode) {
    if m == nil {
        return nil
    } else {
        return m.iTunesPairingMode
    }
}
// Gets the managementCertificates property value. Management certificates for Apple Configurator
func (m *DepIOSEnrollmentProfile) GetManagementCertificates()([]ManagementCertificateWithThumbprint) {
    if m == nil {
        return nil
    } else {
        return m.managementCertificates
    }
}
// Gets the onBoardingScreenDisabled property value. Indicates if onboarding setup screen is disabled
func (m *DepIOSEnrollmentProfile) GetOnBoardingScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.onBoardingScreenDisabled
    }
}
// Gets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepIOSEnrollmentProfile) GetPassCodeDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passCodeDisabled
    }
}
// Gets the passcodeLockGracePeriodInSeconds property value. Indicates timeout before locked screen requires the user to enter the device passocde to unlock it
func (m *DepIOSEnrollmentProfile) GetPasscodeLockGracePeriodInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passcodeLockGracePeriodInSeconds
    }
}
// Gets the preferredLanguageScreenDisabled property value. Indicates if Preferred language screen is disabled
func (m *DepIOSEnrollmentProfile) GetPreferredLanguageScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.preferredLanguageScreenDisabled
    }
}
// Gets the restoreCompletedScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) GetRestoreCompletedScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.restoreCompletedScreenDisabled
    }
}
// Gets the restoreFromAndroidDisabled property value. Indicates if Restore from Android is disabled
func (m *DepIOSEnrollmentProfile) GetRestoreFromAndroidDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.restoreFromAndroidDisabled
    }
}
// Gets the sharedIPadMaximumUserCount property value. This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
func (m *DepIOSEnrollmentProfile) GetSharedIPadMaximumUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sharedIPadMaximumUserCount
    }
}
// Gets the simSetupScreenDisabled property value. Indicates if the SIMSetup screen is disabled
func (m *DepIOSEnrollmentProfile) GetSimSetupScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.simSetupScreenDisabled
    }
}
// Gets the softwareUpdateScreenDisabled property value. Indicates if the mandatory sofware update screen is disabled
func (m *DepIOSEnrollmentProfile) GetSoftwareUpdateScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.softwareUpdateScreenDisabled
    }
}
// Gets the temporarySessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) GetTemporarySessionTimeoutInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.temporarySessionTimeoutInSeconds
    }
}
// Gets the updateCompleteScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) GetUpdateCompleteScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.updateCompleteScreenDisabled
    }
}
// Gets the userSessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) GetUserSessionTimeoutInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.userSessionTimeoutInSeconds
    }
}
// Gets the watchMigrationScreenDisabled property value. Indicates if the watch migration screen is disabled
func (m *DepIOSEnrollmentProfile) GetWatchMigrationScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.watchMigrationScreenDisabled
    }
}
// Gets the welcomeScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) GetWelcomeScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.welcomeScreenDisabled
    }
}
// Gets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepIOSEnrollmentProfile) GetZoomDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.zoomDisabled
    }
}
// The deserialization information for the current model
func (m *DepIOSEnrollmentProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DepEnrollmentBaseProfile.GetFieldDeserializers()
    res["appearanceScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAppearanceScreenDisabled(val)
        return nil
    }
    res["awaitDeviceConfiguredConfirmation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAwaitDeviceConfiguredConfirmation(val)
        return nil
    }
    res["carrierActivationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCarrierActivationUrl(val)
        return nil
    }
    res["companyPortalVppTokenId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCompanyPortalVppTokenId(val)
        return nil
    }
    res["deviceToDeviceMigrationDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDeviceToDeviceMigrationDisabled(val)
        return nil
    }
    res["enableSharedIPad"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnableSharedIPad(val)
        return nil
    }
    res["enableSingleAppEnrollmentMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnableSingleAppEnrollmentMode(val)
        return nil
    }
    res["expressLanguageScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetExpressLanguageScreenDisabled(val)
        return nil
    }
    res["forceTemporarySession"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetForceTemporarySession(val)
        return nil
    }
    res["homeButtonScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHomeButtonScreenDisabled(val)
        return nil
    }
    res["iMessageAndFaceTimeScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIMessageAndFaceTimeScreenDisabled(val)
        return nil
    }
    res["iTunesPairingMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseITunesPairingMode)
        if err != nil {
            return err
        }
        cast := val.(ITunesPairingMode)
        m.SetITunesPairingMode(&cast)
        return nil
    }
    res["managementCertificates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementCertificateWithThumbprint() })
        if err != nil {
            return err
        }
        res := make([]ManagementCertificateWithThumbprint, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagementCertificateWithThumbprint))
        }
        m.SetManagementCertificates(res)
        return nil
    }
    res["onBoardingScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOnBoardingScreenDisabled(val)
        return nil
    }
    res["passCodeDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPassCodeDisabled(val)
        return nil
    }
    res["passcodeLockGracePeriodInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPasscodeLockGracePeriodInSeconds(val)
        return nil
    }
    res["preferredLanguageScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPreferredLanguageScreenDisabled(val)
        return nil
    }
    res["restoreCompletedScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRestoreCompletedScreenDisabled(val)
        return nil
    }
    res["restoreFromAndroidDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRestoreFromAndroidDisabled(val)
        return nil
    }
    res["sharedIPadMaximumUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSharedIPadMaximumUserCount(val)
        return nil
    }
    res["simSetupScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSimSetupScreenDisabled(val)
        return nil
    }
    res["softwareUpdateScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSoftwareUpdateScreenDisabled(val)
        return nil
    }
    res["temporarySessionTimeoutInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTemporarySessionTimeoutInSeconds(val)
        return nil
    }
    res["updateCompleteScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUpdateCompleteScreenDisabled(val)
        return nil
    }
    res["userSessionTimeoutInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUserSessionTimeoutInSeconds(val)
        return nil
    }
    res["watchMigrationScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetWatchMigrationScreenDisabled(val)
        return nil
    }
    res["welcomeScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetWelcomeScreenDisabled(val)
        return nil
    }
    res["zoomDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetZoomDisabled(val)
        return nil
    }
    return res
}
func (m *DepIOSEnrollmentProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        cast := m.GetITunesPairingMode().String()
        err = writer.WriteStringValue("iTunesPairingMode", &cast)
        if err != nil {
            return err
        }
    }
    {
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
// Sets the appearanceScreenDisabled property value. Indicates if Apperance screen is disabled
// Parameters:
//  - value : Value to set for the appearanceScreenDisabled property.
func (m *DepIOSEnrollmentProfile) SetAppearanceScreenDisabled(value *bool)() {
    m.appearanceScreenDisabled = value
}
// Sets the awaitDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
// Parameters:
//  - value : Value to set for the awaitDeviceConfiguredConfirmation property.
func (m *DepIOSEnrollmentProfile) SetAwaitDeviceConfiguredConfirmation(value *bool)() {
    m.awaitDeviceConfiguredConfirmation = value
}
// Sets the carrierActivationUrl property value. Carrier URL for activating device eSIM.
// Parameters:
//  - value : Value to set for the carrierActivationUrl property.
func (m *DepIOSEnrollmentProfile) SetCarrierActivationUrl(value *string)() {
    m.carrierActivationUrl = value
}
// Sets the companyPortalVppTokenId property value. If set, indicates which Vpp token should be used to deploy the Company Portal w/ device licensing. 'enableAuthenticationViaCompanyPortal' must be set in order for this property to be set.
// Parameters:
//  - value : Value to set for the companyPortalVppTokenId property.
func (m *DepIOSEnrollmentProfile) SetCompanyPortalVppTokenId(value *string)() {
    m.companyPortalVppTokenId = value
}
// Sets the deviceToDeviceMigrationDisabled property value. Indicates if Device To Device Migration is disabled
// Parameters:
//  - value : Value to set for the deviceToDeviceMigrationDisabled property.
func (m *DepIOSEnrollmentProfile) SetDeviceToDeviceMigrationDisabled(value *bool)() {
    m.deviceToDeviceMigrationDisabled = value
}
// Sets the enableSharedIPad property value. This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
// Parameters:
//  - value : Value to set for the enableSharedIPad property.
func (m *DepIOSEnrollmentProfile) SetEnableSharedIPad(value *bool)() {
    m.enableSharedIPad = value
}
// Sets the enableSingleAppEnrollmentMode property value. Tells the device to enable single app mode and apply app-lock during enrollment. Default is false. 'enableAuthenticationViaCompanyPortal' and 'companyPortalVppTokenId' must be set for this property to be set.
// Parameters:
//  - value : Value to set for the enableSingleAppEnrollmentMode property.
func (m *DepIOSEnrollmentProfile) SetEnableSingleAppEnrollmentMode(value *bool)() {
    m.enableSingleAppEnrollmentMode = value
}
// Sets the expressLanguageScreenDisabled property value. Indicates if Express Language screen is disabled
// Parameters:
//  - value : Value to set for the expressLanguageScreenDisabled property.
func (m *DepIOSEnrollmentProfile) SetExpressLanguageScreenDisabled(value *bool)() {
    m.expressLanguageScreenDisabled = value
}
// Sets the forceTemporarySession property value. Indicates if temporary sessions is enabled
// Parameters:
//  - value : Value to set for the forceTemporarySession property.
func (m *DepIOSEnrollmentProfile) SetForceTemporarySession(value *bool)() {
    m.forceTemporarySession = value
}
// Sets the homeButtonScreenDisabled property value. Indicates if home button sensitivity screen is disabled
// Parameters:
//  - value : Value to set for the homeButtonScreenDisabled property.
func (m *DepIOSEnrollmentProfile) SetHomeButtonScreenDisabled(value *bool)() {
    m.homeButtonScreenDisabled = value
}
// Sets the iMessageAndFaceTimeScreenDisabled property value. Indicates if iMessage and FaceTime screen is disabled
// Parameters:
//  - value : Value to set for the iMessageAndFaceTimeScreenDisabled property.
func (m *DepIOSEnrollmentProfile) SetIMessageAndFaceTimeScreenDisabled(value *bool)() {
    m.iMessageAndFaceTimeScreenDisabled = value
}
// Sets the iTunesPairingMode property value. Indicates the iTunes pairing mode. Possible values are: disallow, allow, requiresCertificate.
// Parameters:
//  - value : Value to set for the iTunesPairingMode property.
func (m *DepIOSEnrollmentProfile) SetITunesPairingMode(value *ITunesPairingMode)() {
    m.iTunesPairingMode = value
}
// Sets the managementCertificates property value. Management certificates for Apple Configurator
// Parameters:
//  - value : Value to set for the managementCertificates property.
func (m *DepIOSEnrollmentProfile) SetManagementCertificates(value []ManagementCertificateWithThumbprint)() {
    m.managementCertificates = value
}
// Sets the onBoardingScreenDisabled property value. Indicates if onboarding setup screen is disabled
// Parameters:
//  - value : Value to set for the onBoardingScreenDisabled property.
func (m *DepIOSEnrollmentProfile) SetOnBoardingScreenDisabled(value *bool)() {
    m.onBoardingScreenDisabled = value
}
// Sets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
// Parameters:
//  - value : Value to set for the passCodeDisabled property.
func (m *DepIOSEnrollmentProfile) SetPassCodeDisabled(value *bool)() {
    m.passCodeDisabled = value
}
// Sets the passcodeLockGracePeriodInSeconds property value. Indicates timeout before locked screen requires the user to enter the device passocde to unlock it
// Parameters:
//  - value : Value to set for the passcodeLockGracePeriodInSeconds property.
func (m *DepIOSEnrollmentProfile) SetPasscodeLockGracePeriodInSeconds(value *int32)() {
    m.passcodeLockGracePeriodInSeconds = value
}
// Sets the preferredLanguageScreenDisabled property value. Indicates if Preferred language screen is disabled
// Parameters:
//  - value : Value to set for the preferredLanguageScreenDisabled property.
func (m *DepIOSEnrollmentProfile) SetPreferredLanguageScreenDisabled(value *bool)() {
    m.preferredLanguageScreenDisabled = value
}
// Sets the restoreCompletedScreenDisabled property value. Indicates if Weclome screen is disabled
// Parameters:
//  - value : Value to set for the restoreCompletedScreenDisabled property.
func (m *DepIOSEnrollmentProfile) SetRestoreCompletedScreenDisabled(value *bool)() {
    m.restoreCompletedScreenDisabled = value
}
// Sets the restoreFromAndroidDisabled property value. Indicates if Restore from Android is disabled
// Parameters:
//  - value : Value to set for the restoreFromAndroidDisabled property.
func (m *DepIOSEnrollmentProfile) SetRestoreFromAndroidDisabled(value *bool)() {
    m.restoreFromAndroidDisabled = value
}
// Sets the sharedIPadMaximumUserCount property value. This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
// Parameters:
//  - value : Value to set for the sharedIPadMaximumUserCount property.
func (m *DepIOSEnrollmentProfile) SetSharedIPadMaximumUserCount(value *int32)() {
    m.sharedIPadMaximumUserCount = value
}
// Sets the simSetupScreenDisabled property value. Indicates if the SIMSetup screen is disabled
// Parameters:
//  - value : Value to set for the simSetupScreenDisabled property.
func (m *DepIOSEnrollmentProfile) SetSimSetupScreenDisabled(value *bool)() {
    m.simSetupScreenDisabled = value
}
// Sets the softwareUpdateScreenDisabled property value. Indicates if the mandatory sofware update screen is disabled
// Parameters:
//  - value : Value to set for the softwareUpdateScreenDisabled property.
func (m *DepIOSEnrollmentProfile) SetSoftwareUpdateScreenDisabled(value *bool)() {
    m.softwareUpdateScreenDisabled = value
}
// Sets the temporarySessionTimeoutInSeconds property value. Indicates timeout of temporary session
// Parameters:
//  - value : Value to set for the temporarySessionTimeoutInSeconds property.
func (m *DepIOSEnrollmentProfile) SetTemporarySessionTimeoutInSeconds(value *int32)() {
    m.temporarySessionTimeoutInSeconds = value
}
// Sets the updateCompleteScreenDisabled property value. Indicates if Weclome screen is disabled
// Parameters:
//  - value : Value to set for the updateCompleteScreenDisabled property.
func (m *DepIOSEnrollmentProfile) SetUpdateCompleteScreenDisabled(value *bool)() {
    m.updateCompleteScreenDisabled = value
}
// Sets the userSessionTimeoutInSeconds property value. Indicates timeout of temporary session
// Parameters:
//  - value : Value to set for the userSessionTimeoutInSeconds property.
func (m *DepIOSEnrollmentProfile) SetUserSessionTimeoutInSeconds(value *int32)() {
    m.userSessionTimeoutInSeconds = value
}
// Sets the watchMigrationScreenDisabled property value. Indicates if the watch migration screen is disabled
// Parameters:
//  - value : Value to set for the watchMigrationScreenDisabled property.
func (m *DepIOSEnrollmentProfile) SetWatchMigrationScreenDisabled(value *bool)() {
    m.watchMigrationScreenDisabled = value
}
// Sets the welcomeScreenDisabled property value. Indicates if Weclome screen is disabled
// Parameters:
//  - value : Value to set for the welcomeScreenDisabled property.
func (m *DepIOSEnrollmentProfile) SetWelcomeScreenDisabled(value *bool)() {
    m.welcomeScreenDisabled = value
}
// Sets the zoomDisabled property value. Indicates if zoom setup pane is disabled
// Parameters:
//  - value : Value to set for the zoomDisabled property.
func (m *DepIOSEnrollmentProfile) SetZoomDisabled(value *bool)() {
    m.zoomDisabled = value
}
