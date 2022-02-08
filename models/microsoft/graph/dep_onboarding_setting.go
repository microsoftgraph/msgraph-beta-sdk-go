package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DepOnboardingSetting 
type DepOnboardingSetting struct {
    Entity
    // The Apple ID used to obtain the current token.
    appleIdentifier *string;
    // Consent granted for data sharing with Apple Dep Service
    dataSharingConsentGranted *bool;
    // Default iOS Enrollment Profile
    defaultIosEnrollmentProfile *DepIOSEnrollmentProfile;
    // Default MacOs Enrollment Profile
    defaultMacOsEnrollmentProfile *DepMacOSEnrollmentProfile;
    // The enrollment profiles.
    enrollmentProfiles []EnrollmentProfile;
    // The imported Apple device identities.
    importedAppleDeviceIdentities []ImportedAppleDeviceIdentity;
    // When the service was onboarded.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // When the service last syned with Intune
    lastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Error code reported by Apple during last dep sync.
    lastSyncErrorCode *int32;
    // When Intune last requested a sync.
    lastSyncTriggeredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string;
    // Whether or not the Dep token sharing is enabled with the School Data Sync service.
    shareTokenWithSchoolDataSyncService *bool;
    // Gets synced device count
    syncedDeviceCount *int32;
    // When the token will expire.
    tokenExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Friendly Name for Dep Token
    tokenName *string;
    // Gets or sets the Dep Token Type. Possible values are: none, dep, appleSchoolManager.
    tokenType *DepTokenType;
}
// NewDepOnboardingSetting instantiates a new depOnboardingSetting and sets the default values.
func NewDepOnboardingSetting()(*DepOnboardingSetting) {
    m := &DepOnboardingSetting{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppleIdentifier gets the appleIdentifier property value. The Apple ID used to obtain the current token.
func (m *DepOnboardingSetting) GetAppleIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appleIdentifier
    }
}
// GetDataSharingConsentGranted gets the dataSharingConsentGranted property value. Consent granted for data sharing with Apple Dep Service
func (m *DepOnboardingSetting) GetDataSharingConsentGranted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.dataSharingConsentGranted
    }
}
// GetDefaultIosEnrollmentProfile gets the defaultIosEnrollmentProfile property value. Default iOS Enrollment Profile
func (m *DepOnboardingSetting) GetDefaultIosEnrollmentProfile()(*DepIOSEnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.defaultIosEnrollmentProfile
    }
}
// GetDefaultMacOsEnrollmentProfile gets the defaultMacOsEnrollmentProfile property value. Default MacOs Enrollment Profile
func (m *DepOnboardingSetting) GetDefaultMacOsEnrollmentProfile()(*DepMacOSEnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.defaultMacOsEnrollmentProfile
    }
}
// GetEnrollmentProfiles gets the enrollmentProfiles property value. The enrollment profiles.
func (m *DepOnboardingSetting) GetEnrollmentProfiles()([]EnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentProfiles
    }
}
// GetImportedAppleDeviceIdentities gets the importedAppleDeviceIdentities property value. The imported Apple device identities.
func (m *DepOnboardingSetting) GetImportedAppleDeviceIdentities()([]ImportedAppleDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.importedAppleDeviceIdentities
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. When the service was onboarded.
func (m *DepOnboardingSetting) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLastSuccessfulSyncDateTime gets the lastSuccessfulSyncDateTime property value. When the service last syned with Intune
func (m *DepOnboardingSetting) GetLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSuccessfulSyncDateTime
    }
}
// GetLastSyncErrorCode gets the lastSyncErrorCode property value. Error code reported by Apple during last dep sync.
func (m *DepOnboardingSetting) GetLastSyncErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncErrorCode
    }
}
// GetLastSyncTriggeredDateTime gets the lastSyncTriggeredDateTime property value. When Intune last requested a sync.
func (m *DepOnboardingSetting) GetLastSyncTriggeredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncTriggeredDateTime
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DepOnboardingSetting) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetShareTokenWithSchoolDataSyncService gets the shareTokenWithSchoolDataSyncService property value. Whether or not the Dep token sharing is enabled with the School Data Sync service.
func (m *DepOnboardingSetting) GetShareTokenWithSchoolDataSyncService()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.shareTokenWithSchoolDataSyncService
    }
}
// GetSyncedDeviceCount gets the syncedDeviceCount property value. Gets synced device count
func (m *DepOnboardingSetting) GetSyncedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.syncedDeviceCount
    }
}
// GetTokenExpirationDateTime gets the tokenExpirationDateTime property value. When the token will expire.
func (m *DepOnboardingSetting) GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.tokenExpirationDateTime
    }
}
// GetTokenName gets the tokenName property value. Friendly Name for Dep Token
func (m *DepOnboardingSetting) GetTokenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenName
    }
}
// GetTokenType gets the tokenType property value. Gets or sets the Dep Token Type. Possible values are: none, dep, appleSchoolManager.
func (m *DepOnboardingSetting) GetTokenType()(*DepTokenType) {
    if m == nil {
        return nil
    } else {
        return m.tokenType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DepOnboardingSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appleIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleIdentifier(val)
        }
        return nil
    }
    res["dataSharingConsentGranted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSharingConsentGranted(val)
        }
        return nil
    }
    res["defaultIosEnrollmentProfile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDepIOSEnrollmentProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultIosEnrollmentProfile(val.(*DepIOSEnrollmentProfile))
        }
        return nil
    }
    res["defaultMacOsEnrollmentProfile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDepMacOSEnrollmentProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultMacOsEnrollmentProfile(val.(*DepMacOSEnrollmentProfile))
        }
        return nil
    }
    res["enrollmentProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEnrollmentProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EnrollmentProfile, len(val))
            for i, v := range val {
                res[i] = *(v.(*EnrollmentProfile))
            }
            m.SetEnrollmentProfiles(res)
        }
        return nil
    }
    res["importedAppleDeviceIdentities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImportedAppleDeviceIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImportedAppleDeviceIdentity, len(val))
            for i, v := range val {
                res[i] = *(v.(*ImportedAppleDeviceIdentity))
            }
            m.SetImportedAppleDeviceIdentities(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["lastSuccessfulSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["lastSyncErrorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncErrorCode(val)
        }
        return nil
    }
    res["lastSyncTriggeredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncTriggeredDateTime(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["shareTokenWithSchoolDataSyncService"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareTokenWithSchoolDataSyncService(val)
        }
        return nil
    }
    res["syncedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncedDeviceCount(val)
        }
        return nil
    }
    res["tokenExpirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenExpirationDateTime(val)
        }
        return nil
    }
    res["tokenName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenName(val)
        }
        return nil
    }
    res["tokenType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDepTokenType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenType(val.(*DepTokenType))
        }
        return nil
    }
    return res
}
func (m *DepOnboardingSetting) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DepOnboardingSetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appleIdentifier", m.GetAppleIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dataSharingConsentGranted", m.GetDataSharingConsentGranted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultIosEnrollmentProfile", m.GetDefaultIosEnrollmentProfile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultMacOsEnrollmentProfile", m.GetDefaultMacOsEnrollmentProfile())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentProfiles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEnrollmentProfiles()))
        for i, v := range m.GetEnrollmentProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("enrollmentProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportedAppleDeviceIdentities() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetImportedAppleDeviceIdentities()))
        for i, v := range m.GetImportedAppleDeviceIdentities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("importedAppleDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSuccessfulSyncDateTime", m.GetLastSuccessfulSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("lastSyncErrorCode", m.GetLastSyncErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncTriggeredDateTime", m.GetLastSyncTriggeredDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("shareTokenWithSchoolDataSyncService", m.GetShareTokenWithSchoolDataSyncService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("syncedDeviceCount", m.GetSyncedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("tokenExpirationDateTime", m.GetTokenExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tokenName", m.GetTokenName())
        if err != nil {
            return err
        }
    }
    if m.GetTokenType() != nil {
        cast := (*m.GetTokenType()).String()
        err = writer.WriteStringValue("tokenType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppleIdentifier sets the appleIdentifier property value. The Apple ID used to obtain the current token.
func (m *DepOnboardingSetting) SetAppleIdentifier(value *string)() {
    if m != nil {
        m.appleIdentifier = value
    }
}
// SetDataSharingConsentGranted sets the dataSharingConsentGranted property value. Consent granted for data sharing with Apple Dep Service
func (m *DepOnboardingSetting) SetDataSharingConsentGranted(value *bool)() {
    if m != nil {
        m.dataSharingConsentGranted = value
    }
}
// SetDefaultIosEnrollmentProfile sets the defaultIosEnrollmentProfile property value. Default iOS Enrollment Profile
func (m *DepOnboardingSetting) SetDefaultIosEnrollmentProfile(value *DepIOSEnrollmentProfile)() {
    if m != nil {
        m.defaultIosEnrollmentProfile = value
    }
}
// SetDefaultMacOsEnrollmentProfile sets the defaultMacOsEnrollmentProfile property value. Default MacOs Enrollment Profile
func (m *DepOnboardingSetting) SetDefaultMacOsEnrollmentProfile(value *DepMacOSEnrollmentProfile)() {
    if m != nil {
        m.defaultMacOsEnrollmentProfile = value
    }
}
// SetEnrollmentProfiles sets the enrollmentProfiles property value. The enrollment profiles.
func (m *DepOnboardingSetting) SetEnrollmentProfiles(value []EnrollmentProfile)() {
    if m != nil {
        m.enrollmentProfiles = value
    }
}
// SetImportedAppleDeviceIdentities sets the importedAppleDeviceIdentities property value. The imported Apple device identities.
func (m *DepOnboardingSetting) SetImportedAppleDeviceIdentities(value []ImportedAppleDeviceIdentity)() {
    if m != nil {
        m.importedAppleDeviceIdentities = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. When the service was onboarded.
func (m *DepOnboardingSetting) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetLastSuccessfulSyncDateTime sets the lastSuccessfulSyncDateTime property value. When the service last syned with Intune
func (m *DepOnboardingSetting) SetLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSuccessfulSyncDateTime = value
    }
}
// SetLastSyncErrorCode sets the lastSyncErrorCode property value. Error code reported by Apple during last dep sync.
func (m *DepOnboardingSetting) SetLastSyncErrorCode(value *int32)() {
    if m != nil {
        m.lastSyncErrorCode = value
    }
}
// SetLastSyncTriggeredDateTime sets the lastSyncTriggeredDateTime property value. When Intune last requested a sync.
func (m *DepOnboardingSetting) SetLastSyncTriggeredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncTriggeredDateTime = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DepOnboardingSetting) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetShareTokenWithSchoolDataSyncService sets the shareTokenWithSchoolDataSyncService property value. Whether or not the Dep token sharing is enabled with the School Data Sync service.
func (m *DepOnboardingSetting) SetShareTokenWithSchoolDataSyncService(value *bool)() {
    if m != nil {
        m.shareTokenWithSchoolDataSyncService = value
    }
}
// SetSyncedDeviceCount sets the syncedDeviceCount property value. Gets synced device count
func (m *DepOnboardingSetting) SetSyncedDeviceCount(value *int32)() {
    if m != nil {
        m.syncedDeviceCount = value
    }
}
// SetTokenExpirationDateTime sets the tokenExpirationDateTime property value. When the token will expire.
func (m *DepOnboardingSetting) SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.tokenExpirationDateTime = value
    }
}
// SetTokenName sets the tokenName property value. Friendly Name for Dep Token
func (m *DepOnboardingSetting) SetTokenName(value *string)() {
    if m != nil {
        m.tokenName = value
    }
}
// SetTokenType sets the tokenType property value. Gets or sets the Dep Token Type. Possible values are: none, dep, appleSchoolManager.
func (m *DepOnboardingSetting) SetTokenType(value *DepTokenType)() {
    if m != nil {
        m.tokenType = value
    }
}
