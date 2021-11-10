package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new depOnboardingSetting and sets the default values.
func NewDepOnboardingSetting()(*DepOnboardingSetting) {
    m := &DepOnboardingSetting{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appleIdentifier property value. The Apple ID used to obtain the current token.
func (m *DepOnboardingSetting) GetAppleIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appleIdentifier
    }
}
// Gets the dataSharingConsentGranted property value. Consent granted for data sharing with Apple Dep Service
func (m *DepOnboardingSetting) GetDataSharingConsentGranted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.dataSharingConsentGranted
    }
}
// Gets the defaultIosEnrollmentProfile property value. Default iOS Enrollment Profile
func (m *DepOnboardingSetting) GetDefaultIosEnrollmentProfile()(*DepIOSEnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.defaultIosEnrollmentProfile
    }
}
// Gets the defaultMacOsEnrollmentProfile property value. Default MacOs Enrollment Profile
func (m *DepOnboardingSetting) GetDefaultMacOsEnrollmentProfile()(*DepMacOSEnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.defaultMacOsEnrollmentProfile
    }
}
// Gets the enrollmentProfiles property value. The enrollment profiles.
func (m *DepOnboardingSetting) GetEnrollmentProfiles()([]EnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentProfiles
    }
}
// Gets the importedAppleDeviceIdentities property value. The imported Apple device identities.
func (m *DepOnboardingSetting) GetImportedAppleDeviceIdentities()([]ImportedAppleDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.importedAppleDeviceIdentities
    }
}
// Gets the lastModifiedDateTime property value. When the service was onboarded.
func (m *DepOnboardingSetting) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the lastSuccessfulSyncDateTime property value. When the service last syned with Intune
func (m *DepOnboardingSetting) GetLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSuccessfulSyncDateTime
    }
}
// Gets the lastSyncErrorCode property value. Error code reported by Apple during last dep sync.
func (m *DepOnboardingSetting) GetLastSyncErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncErrorCode
    }
}
// Gets the lastSyncTriggeredDateTime property value. When Intune last requested a sync.
func (m *DepOnboardingSetting) GetLastSyncTriggeredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncTriggeredDateTime
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DepOnboardingSetting) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the shareTokenWithSchoolDataSyncService property value. Whether or not the Dep token sharing is enabled with the School Data Sync service.
func (m *DepOnboardingSetting) GetShareTokenWithSchoolDataSyncService()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.shareTokenWithSchoolDataSyncService
    }
}
// Gets the syncedDeviceCount property value. Gets synced device count
func (m *DepOnboardingSetting) GetSyncedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.syncedDeviceCount
    }
}
// Gets the tokenExpirationDateTime property value. When the token will expire.
func (m *DepOnboardingSetting) GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.tokenExpirationDateTime
    }
}
// Gets the tokenName property value. Friendly Name for Dep Token
func (m *DepOnboardingSetting) GetTokenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenName
    }
}
// Gets the tokenType property value. Gets or sets the Dep Token Type. Possible values are: none, dep, appleSchoolManager.
func (m *DepOnboardingSetting) GetTokenType()(*DepTokenType) {
    if m == nil {
        return nil
    } else {
        return m.tokenType
    }
}
// The deserialization information for the current model
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
            cast := val.(DepTokenType)
            m.SetTokenType(&cast)
        }
        return nil
    }
    return res
}
func (m *DepOnboardingSetting) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
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
    {
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
    {
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
        cast := m.GetTokenType().String()
        err = writer.WriteStringValue("tokenType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appleIdentifier property value. The Apple ID used to obtain the current token.
// Parameters:
//  - value : Value to set for the appleIdentifier property.
func (m *DepOnboardingSetting) SetAppleIdentifier(value *string)() {
    m.appleIdentifier = value
}
// Sets the dataSharingConsentGranted property value. Consent granted for data sharing with Apple Dep Service
// Parameters:
//  - value : Value to set for the dataSharingConsentGranted property.
func (m *DepOnboardingSetting) SetDataSharingConsentGranted(value *bool)() {
    m.dataSharingConsentGranted = value
}
// Sets the defaultIosEnrollmentProfile property value. Default iOS Enrollment Profile
// Parameters:
//  - value : Value to set for the defaultIosEnrollmentProfile property.
func (m *DepOnboardingSetting) SetDefaultIosEnrollmentProfile(value *DepIOSEnrollmentProfile)() {
    m.defaultIosEnrollmentProfile = value
}
// Sets the defaultMacOsEnrollmentProfile property value. Default MacOs Enrollment Profile
// Parameters:
//  - value : Value to set for the defaultMacOsEnrollmentProfile property.
func (m *DepOnboardingSetting) SetDefaultMacOsEnrollmentProfile(value *DepMacOSEnrollmentProfile)() {
    m.defaultMacOsEnrollmentProfile = value
}
// Sets the enrollmentProfiles property value. The enrollment profiles.
// Parameters:
//  - value : Value to set for the enrollmentProfiles property.
func (m *DepOnboardingSetting) SetEnrollmentProfiles(value []EnrollmentProfile)() {
    m.enrollmentProfiles = value
}
// Sets the importedAppleDeviceIdentities property value. The imported Apple device identities.
// Parameters:
//  - value : Value to set for the importedAppleDeviceIdentities property.
func (m *DepOnboardingSetting) SetImportedAppleDeviceIdentities(value []ImportedAppleDeviceIdentity)() {
    m.importedAppleDeviceIdentities = value
}
// Sets the lastModifiedDateTime property value. When the service was onboarded.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *DepOnboardingSetting) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the lastSuccessfulSyncDateTime property value. When the service last syned with Intune
// Parameters:
//  - value : Value to set for the lastSuccessfulSyncDateTime property.
func (m *DepOnboardingSetting) SetLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSuccessfulSyncDateTime = value
}
// Sets the lastSyncErrorCode property value. Error code reported by Apple during last dep sync.
// Parameters:
//  - value : Value to set for the lastSyncErrorCode property.
func (m *DepOnboardingSetting) SetLastSyncErrorCode(value *int32)() {
    m.lastSyncErrorCode = value
}
// Sets the lastSyncTriggeredDateTime property value. When Intune last requested a sync.
// Parameters:
//  - value : Value to set for the lastSyncTriggeredDateTime property.
func (m *DepOnboardingSetting) SetLastSyncTriggeredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncTriggeredDateTime = value
}
// Sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *DepOnboardingSetting) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the shareTokenWithSchoolDataSyncService property value. Whether or not the Dep token sharing is enabled with the School Data Sync service.
// Parameters:
//  - value : Value to set for the shareTokenWithSchoolDataSyncService property.
func (m *DepOnboardingSetting) SetShareTokenWithSchoolDataSyncService(value *bool)() {
    m.shareTokenWithSchoolDataSyncService = value
}
// Sets the syncedDeviceCount property value. Gets synced device count
// Parameters:
//  - value : Value to set for the syncedDeviceCount property.
func (m *DepOnboardingSetting) SetSyncedDeviceCount(value *int32)() {
    m.syncedDeviceCount = value
}
// Sets the tokenExpirationDateTime property value. When the token will expire.
// Parameters:
//  - value : Value to set for the tokenExpirationDateTime property.
func (m *DepOnboardingSetting) SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.tokenExpirationDateTime = value
}
// Sets the tokenName property value. Friendly Name for Dep Token
// Parameters:
//  - value : Value to set for the tokenName property.
func (m *DepOnboardingSetting) SetTokenName(value *string)() {
    m.tokenName = value
}
// Sets the tokenType property value. Gets or sets the Dep Token Type. Possible values are: none, dep, appleSchoolManager.
// Parameters:
//  - value : Value to set for the tokenType property.
func (m *DepOnboardingSetting) SetTokenType(value *DepTokenType)() {
    m.tokenType = value
}
