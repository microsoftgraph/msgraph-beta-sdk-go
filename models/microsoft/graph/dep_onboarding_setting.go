package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DepOnboardingSetting struct {
    Entity
    appleIdentifier *string;
    dataSharingConsentGranted *bool;
    defaultIosEnrollmentProfile *DepIOSEnrollmentProfile;
    defaultMacOsEnrollmentProfile *DepMacOSEnrollmentProfile;
    enrollmentProfiles []EnrollmentProfile;
    importedAppleDeviceIdentities []ImportedAppleDeviceIdentity;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastSyncErrorCode *int32;
    lastSyncTriggeredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    roleScopeTagIds []string;
    shareTokenWithSchoolDataSyncService *bool;
    syncedDeviceCount *int32;
    tokenExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    tokenName *string;
    tokenType *DepTokenType;
}
func NewDepOnboardingSetting()(*DepOnboardingSetting) {
    m := &DepOnboardingSetting{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DepOnboardingSetting) GetAppleIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appleIdentifier
    }
}
func (m *DepOnboardingSetting) GetDataSharingConsentGranted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.dataSharingConsentGranted
    }
}
func (m *DepOnboardingSetting) GetDefaultIosEnrollmentProfile()(*DepIOSEnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.defaultIosEnrollmentProfile
    }
}
func (m *DepOnboardingSetting) GetDefaultMacOsEnrollmentProfile()(*DepMacOSEnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.defaultMacOsEnrollmentProfile
    }
}
func (m *DepOnboardingSetting) GetEnrollmentProfiles()([]EnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentProfiles
    }
}
func (m *DepOnboardingSetting) GetImportedAppleDeviceIdentities()([]ImportedAppleDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.importedAppleDeviceIdentities
    }
}
func (m *DepOnboardingSetting) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *DepOnboardingSetting) GetLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSuccessfulSyncDateTime
    }
}
func (m *DepOnboardingSetting) GetLastSyncErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncErrorCode
    }
}
func (m *DepOnboardingSetting) GetLastSyncTriggeredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncTriggeredDateTime
    }
}
func (m *DepOnboardingSetting) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
func (m *DepOnboardingSetting) GetShareTokenWithSchoolDataSyncService()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.shareTokenWithSchoolDataSyncService
    }
}
func (m *DepOnboardingSetting) GetSyncedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.syncedDeviceCount
    }
}
func (m *DepOnboardingSetting) GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.tokenExpirationDateTime
    }
}
func (m *DepOnboardingSetting) GetTokenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenName
    }
}
func (m *DepOnboardingSetting) GetTokenType()(*DepTokenType) {
    if m == nil {
        return nil
    } else {
        return m.tokenType
    }
}
func (m *DepOnboardingSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appleIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppleIdentifier(val)
        return nil
    }
    res["dataSharingConsentGranted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDataSharingConsentGranted(val)
        return nil
    }
    res["defaultIosEnrollmentProfile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDepIOSEnrollmentProfile() })
        if err != nil {
            return err
        }
        m.SetDefaultIosEnrollmentProfile(val.(*DepIOSEnrollmentProfile))
        return nil
    }
    res["defaultMacOsEnrollmentProfile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDepMacOSEnrollmentProfile() })
        if err != nil {
            return err
        }
        m.SetDefaultMacOsEnrollmentProfile(val.(*DepMacOSEnrollmentProfile))
        return nil
    }
    res["enrollmentProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEnrollmentProfile() })
        if err != nil {
            return err
        }
        res := make([]EnrollmentProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*EnrollmentProfile))
        }
        m.SetEnrollmentProfiles(res)
        return nil
    }
    res["importedAppleDeviceIdentities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImportedAppleDeviceIdentity() })
        if err != nil {
            return err
        }
        res := make([]ImportedAppleDeviceIdentity, len(val))
        for i, v := range val {
            res[i] = *(v.(*ImportedAppleDeviceIdentity))
        }
        m.SetImportedAppleDeviceIdentities(res)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["lastSuccessfulSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSuccessfulSyncDateTime(val)
        return nil
    }
    res["lastSyncErrorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetLastSyncErrorCode(val)
        return nil
    }
    res["lastSyncTriggeredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSyncTriggeredDateTime(val)
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRoleScopeTagIds(res)
        return nil
    }
    res["shareTokenWithSchoolDataSyncService"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShareTokenWithSchoolDataSyncService(val)
        return nil
    }
    res["syncedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSyncedDeviceCount(val)
        return nil
    }
    res["tokenExpirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetTokenExpirationDateTime(val)
        return nil
    }
    res["tokenName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTokenName(val)
        return nil
    }
    res["tokenType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDepTokenType)
        if err != nil {
            return err
        }
        cast := val.(DepTokenType)
        m.SetTokenType(&cast)
        return nil
    }
    return res
}
func (m *DepOnboardingSetting) IsNil()(bool) {
    return m == nil
}
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
func (m *DepOnboardingSetting) SetAppleIdentifier(value *string)() {
    m.appleIdentifier = value
}
func (m *DepOnboardingSetting) SetDataSharingConsentGranted(value *bool)() {
    m.dataSharingConsentGranted = value
}
func (m *DepOnboardingSetting) SetDefaultIosEnrollmentProfile(value *DepIOSEnrollmentProfile)() {
    m.defaultIosEnrollmentProfile = value
}
func (m *DepOnboardingSetting) SetDefaultMacOsEnrollmentProfile(value *DepMacOSEnrollmentProfile)() {
    m.defaultMacOsEnrollmentProfile = value
}
func (m *DepOnboardingSetting) SetEnrollmentProfiles(value []EnrollmentProfile)() {
    m.enrollmentProfiles = value
}
func (m *DepOnboardingSetting) SetImportedAppleDeviceIdentities(value []ImportedAppleDeviceIdentity)() {
    m.importedAppleDeviceIdentities = value
}
func (m *DepOnboardingSetting) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *DepOnboardingSetting) SetLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSuccessfulSyncDateTime = value
}
func (m *DepOnboardingSetting) SetLastSyncErrorCode(value *int32)() {
    m.lastSyncErrorCode = value
}
func (m *DepOnboardingSetting) SetLastSyncTriggeredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncTriggeredDateTime = value
}
func (m *DepOnboardingSetting) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
func (m *DepOnboardingSetting) SetShareTokenWithSchoolDataSyncService(value *bool)() {
    m.shareTokenWithSchoolDataSyncService = value
}
func (m *DepOnboardingSetting) SetSyncedDeviceCount(value *int32)() {
    m.syncedDeviceCount = value
}
func (m *DepOnboardingSetting) SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.tokenExpirationDateTime = value
}
func (m *DepOnboardingSetting) SetTokenName(value *string)() {
    m.tokenName = value
}
func (m *DepOnboardingSetting) SetTokenType(value *DepTokenType)() {
    m.tokenType = value
}
