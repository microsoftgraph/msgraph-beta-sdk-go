package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DepIOSEnrollmentProfile struct {
    DepEnrollmentBaseProfile
    appearanceScreenDisabled *bool;
    awaitDeviceConfiguredConfirmation *bool;
    companyPortalVppTokenId *string;
    deviceToDeviceMigrationDisabled *bool;
    enableSharedIPad *bool;
    enableSingleAppEnrollmentMode *bool;
    expressLanguageScreenDisabled *bool;
    homeButtonScreenDisabled *bool;
    iMessageAndFaceTimeScreenDisabled *bool;
    iTunesPairingMode *ITunesPairingMode;
    managementCertificates []ManagementCertificateWithThumbprint;
    onBoardingScreenDisabled *bool;
    passCodeDisabled *bool;
    preferredLanguageScreenDisabled *bool;
    restoreCompletedScreenDisabled *bool;
    restoreFromAndroidDisabled *bool;
    sharedIPadMaximumUserCount *int32;
    simSetupScreenDisabled *bool;
    softwareUpdateScreenDisabled *bool;
    updateCompleteScreenDisabled *bool;
    watchMigrationScreenDisabled *bool;
    welcomeScreenDisabled *bool;
    zoomDisabled *bool;
}
func NewDepIOSEnrollmentProfile()(*DepIOSEnrollmentProfile) {
    m := &DepIOSEnrollmentProfile{
        DepEnrollmentBaseProfile: *NewDepEnrollmentBaseProfile(),
    }
    return m
}
func (m *DepIOSEnrollmentProfile) GetAppearanceScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appearanceScreenDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetAwaitDeviceConfiguredConfirmation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.awaitDeviceConfiguredConfirmation
    }
}
func (m *DepIOSEnrollmentProfile) GetCompanyPortalVppTokenId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.companyPortalVppTokenId
    }
}
func (m *DepIOSEnrollmentProfile) GetDeviceToDeviceMigrationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceToDeviceMigrationDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetEnableSharedIPad()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableSharedIPad
    }
}
func (m *DepIOSEnrollmentProfile) GetEnableSingleAppEnrollmentMode()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableSingleAppEnrollmentMode
    }
}
func (m *DepIOSEnrollmentProfile) GetExpressLanguageScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.expressLanguageScreenDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetHomeButtonScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.homeButtonScreenDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetIMessageAndFaceTimeScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iMessageAndFaceTimeScreenDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetITunesPairingMode()(*ITunesPairingMode) {
    if m == nil {
        return nil
    } else {
        return m.iTunesPairingMode
    }
}
func (m *DepIOSEnrollmentProfile) GetManagementCertificates()([]ManagementCertificateWithThumbprint) {
    if m == nil {
        return nil
    } else {
        return m.managementCertificates
    }
}
func (m *DepIOSEnrollmentProfile) GetOnBoardingScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.onBoardingScreenDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetPassCodeDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passCodeDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetPreferredLanguageScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.preferredLanguageScreenDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetRestoreCompletedScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.restoreCompletedScreenDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetRestoreFromAndroidDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.restoreFromAndroidDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetSharedIPadMaximumUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sharedIPadMaximumUserCount
    }
}
func (m *DepIOSEnrollmentProfile) GetSimSetupScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.simSetupScreenDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetSoftwareUpdateScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.softwareUpdateScreenDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetUpdateCompleteScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.updateCompleteScreenDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetWatchMigrationScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.watchMigrationScreenDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetWelcomeScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.welcomeScreenDisabled
    }
}
func (m *DepIOSEnrollmentProfile) GetZoomDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.zoomDisabled
    }
}
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
    res["updateCompleteScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUpdateCompleteScreenDisabled(val)
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
        err = writer.WriteBoolValue("updateCompleteScreenDisabled", m.GetUpdateCompleteScreenDisabled())
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
func (m *DepIOSEnrollmentProfile) SetAppearanceScreenDisabled(value *bool)() {
    m.appearanceScreenDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetAwaitDeviceConfiguredConfirmation(value *bool)() {
    m.awaitDeviceConfiguredConfirmation = value
}
func (m *DepIOSEnrollmentProfile) SetCompanyPortalVppTokenId(value *string)() {
    m.companyPortalVppTokenId = value
}
func (m *DepIOSEnrollmentProfile) SetDeviceToDeviceMigrationDisabled(value *bool)() {
    m.deviceToDeviceMigrationDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetEnableSharedIPad(value *bool)() {
    m.enableSharedIPad = value
}
func (m *DepIOSEnrollmentProfile) SetEnableSingleAppEnrollmentMode(value *bool)() {
    m.enableSingleAppEnrollmentMode = value
}
func (m *DepIOSEnrollmentProfile) SetExpressLanguageScreenDisabled(value *bool)() {
    m.expressLanguageScreenDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetHomeButtonScreenDisabled(value *bool)() {
    m.homeButtonScreenDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetIMessageAndFaceTimeScreenDisabled(value *bool)() {
    m.iMessageAndFaceTimeScreenDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetITunesPairingMode(value *ITunesPairingMode)() {
    m.iTunesPairingMode = value
}
func (m *DepIOSEnrollmentProfile) SetManagementCertificates(value []ManagementCertificateWithThumbprint)() {
    m.managementCertificates = value
}
func (m *DepIOSEnrollmentProfile) SetOnBoardingScreenDisabled(value *bool)() {
    m.onBoardingScreenDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetPassCodeDisabled(value *bool)() {
    m.passCodeDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetPreferredLanguageScreenDisabled(value *bool)() {
    m.preferredLanguageScreenDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetRestoreCompletedScreenDisabled(value *bool)() {
    m.restoreCompletedScreenDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetRestoreFromAndroidDisabled(value *bool)() {
    m.restoreFromAndroidDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetSharedIPadMaximumUserCount(value *int32)() {
    m.sharedIPadMaximumUserCount = value
}
func (m *DepIOSEnrollmentProfile) SetSimSetupScreenDisabled(value *bool)() {
    m.simSetupScreenDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetSoftwareUpdateScreenDisabled(value *bool)() {
    m.softwareUpdateScreenDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetUpdateCompleteScreenDisabled(value *bool)() {
    m.updateCompleteScreenDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetWatchMigrationScreenDisabled(value *bool)() {
    m.watchMigrationScreenDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetWelcomeScreenDisabled(value *bool)() {
    m.welcomeScreenDisabled = value
}
func (m *DepIOSEnrollmentProfile) SetZoomDisabled(value *bool)() {
    m.zoomDisabled = value
}
