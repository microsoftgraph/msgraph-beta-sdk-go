package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepOnboardingSetting the depOnboardingSetting represents an instance of the Apple DEP service being onboarded to Intune. The onboarded service instance manages an Apple Token used to synchronize data between Apple and Intune.
type DepOnboardingSetting struct {
    Entity
}
// NewDepOnboardingSetting instantiates a new depOnboardingSetting and sets the default values.
func NewDepOnboardingSetting()(*DepOnboardingSetting) {
    m := &DepOnboardingSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDepOnboardingSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDepOnboardingSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDepOnboardingSetting(), nil
}
// GetAppleIdentifier gets the appleIdentifier property value. The Apple ID used to obtain the current token.
func (m *DepOnboardingSetting) GetAppleIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("appleIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDataSharingConsentGranted gets the dataSharingConsentGranted property value. Consent granted for data sharing with Apple Dep Service
func (m *DepOnboardingSetting) GetDataSharingConsentGranted()(*bool) {
    val, err := m.GetBackingStore().Get("dataSharingConsentGranted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefaultIosEnrollmentProfile gets the defaultIosEnrollmentProfile property value. Default iOS Enrollment Profile
func (m *DepOnboardingSetting) GetDefaultIosEnrollmentProfile()(DepIOSEnrollmentProfileable) {
    val, err := m.GetBackingStore().Get("defaultIosEnrollmentProfile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DepIOSEnrollmentProfileable)
    }
    return nil
}
// GetDefaultMacOsEnrollmentProfile gets the defaultMacOsEnrollmentProfile property value. Default MacOs Enrollment Profile
func (m *DepOnboardingSetting) GetDefaultMacOsEnrollmentProfile()(DepMacOSEnrollmentProfileable) {
    val, err := m.GetBackingStore().Get("defaultMacOsEnrollmentProfile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DepMacOSEnrollmentProfileable)
    }
    return nil
}
// GetEnrollmentProfiles gets the enrollmentProfiles property value. The enrollment profiles.
func (m *DepOnboardingSetting) GetEnrollmentProfiles()([]EnrollmentProfileable) {
    val, err := m.GetBackingStore().Get("enrollmentProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EnrollmentProfileable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DepOnboardingSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appleIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleIdentifier(val)
        }
        return nil
    }
    res["dataSharingConsentGranted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSharingConsentGranted(val)
        }
        return nil
    }
    res["defaultIosEnrollmentProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDepIOSEnrollmentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultIosEnrollmentProfile(val.(DepIOSEnrollmentProfileable))
        }
        return nil
    }
    res["defaultMacOsEnrollmentProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDepMacOSEnrollmentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultMacOsEnrollmentProfile(val.(DepMacOSEnrollmentProfileable))
        }
        return nil
    }
    res["enrollmentProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEnrollmentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EnrollmentProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EnrollmentProfileable)
                }
            }
            m.SetEnrollmentProfiles(res)
        }
        return nil
    }
    res["importedAppleDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImportedAppleDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImportedAppleDeviceIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ImportedAppleDeviceIdentityable)
                }
            }
            m.SetImportedAppleDeviceIdentities(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["lastSuccessfulSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["lastSyncErrorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncErrorCode(val)
        }
        return nil
    }
    res["lastSyncTriggeredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncTriggeredDateTime(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["shareTokenWithSchoolDataSyncService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareTokenWithSchoolDataSyncService(val)
        }
        return nil
    }
    res["syncedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSyncedDeviceCount(val)
        }
        return nil
    }
    res["tokenExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenExpirationDateTime(val)
        }
        return nil
    }
    res["tokenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenName(val)
        }
        return nil
    }
    res["tokenType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetImportedAppleDeviceIdentities gets the importedAppleDeviceIdentities property value. The imported Apple device identities.
func (m *DepOnboardingSetting) GetImportedAppleDeviceIdentities()([]ImportedAppleDeviceIdentityable) {
    val, err := m.GetBackingStore().Get("importedAppleDeviceIdentities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ImportedAppleDeviceIdentityable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. When the service was onboarded.
func (m *DepOnboardingSetting) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastSuccessfulSyncDateTime gets the lastSuccessfulSyncDateTime property value. When the service last syned with Intune
func (m *DepOnboardingSetting) GetLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastSuccessfulSyncDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastSyncErrorCode gets the lastSyncErrorCode property value. Error code reported by Apple during last dep sync.
func (m *DepOnboardingSetting) GetLastSyncErrorCode()(*int32) {
    val, err := m.GetBackingStore().Get("lastSyncErrorCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetLastSyncTriggeredDateTime gets the lastSyncTriggeredDateTime property value. When Intune last requested a sync.
func (m *DepOnboardingSetting) GetLastSyncTriggeredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastSyncTriggeredDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DepOnboardingSetting) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DepOnboardingSetting) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetShareTokenWithSchoolDataSyncService gets the shareTokenWithSchoolDataSyncService property value. Whether or not the Dep token sharing is enabled with the School Data Sync service.
func (m *DepOnboardingSetting) GetShareTokenWithSchoolDataSyncService()(*bool) {
    val, err := m.GetBackingStore().Get("shareTokenWithSchoolDataSyncService")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSyncedDeviceCount gets the syncedDeviceCount property value. Gets synced device count
func (m *DepOnboardingSetting) GetSyncedDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("syncedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTokenExpirationDateTime gets the tokenExpirationDateTime property value. When the token will expire.
func (m *DepOnboardingSetting) GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("tokenExpirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetTokenName gets the tokenName property value. Friendly Name for Dep Token
func (m *DepOnboardingSetting) GetTokenName()(*string) {
    val, err := m.GetBackingStore().Get("tokenName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTokenType gets the tokenType property value. The tokenType property
func (m *DepOnboardingSetting) GetTokenType()(*DepTokenType) {
    val, err := m.GetBackingStore().Get("tokenType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DepTokenType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DepOnboardingSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnrollmentProfiles()))
        for i, v := range m.GetEnrollmentProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("enrollmentProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportedAppleDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImportedAppleDeviceIdentities()))
        for i, v := range m.GetImportedAppleDeviceIdentities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    err := m.GetBackingStore().Set("appleIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetDataSharingConsentGranted sets the dataSharingConsentGranted property value. Consent granted for data sharing with Apple Dep Service
func (m *DepOnboardingSetting) SetDataSharingConsentGranted(value *bool)() {
    err := m.GetBackingStore().Set("dataSharingConsentGranted", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultIosEnrollmentProfile sets the defaultIosEnrollmentProfile property value. Default iOS Enrollment Profile
func (m *DepOnboardingSetting) SetDefaultIosEnrollmentProfile(value DepIOSEnrollmentProfileable)() {
    err := m.GetBackingStore().Set("defaultIosEnrollmentProfile", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultMacOsEnrollmentProfile sets the defaultMacOsEnrollmentProfile property value. Default MacOs Enrollment Profile
func (m *DepOnboardingSetting) SetDefaultMacOsEnrollmentProfile(value DepMacOSEnrollmentProfileable)() {
    err := m.GetBackingStore().Set("defaultMacOsEnrollmentProfile", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentProfiles sets the enrollmentProfiles property value. The enrollment profiles.
func (m *DepOnboardingSetting) SetEnrollmentProfiles(value []EnrollmentProfileable)() {
    err := m.GetBackingStore().Set("enrollmentProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetImportedAppleDeviceIdentities sets the importedAppleDeviceIdentities property value. The imported Apple device identities.
func (m *DepOnboardingSetting) SetImportedAppleDeviceIdentities(value []ImportedAppleDeviceIdentityable)() {
    err := m.GetBackingStore().Set("importedAppleDeviceIdentities", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. When the service was onboarded.
func (m *DepOnboardingSetting) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSuccessfulSyncDateTime sets the lastSuccessfulSyncDateTime property value. When the service last syned with Intune
func (m *DepOnboardingSetting) SetLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastSuccessfulSyncDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSyncErrorCode sets the lastSyncErrorCode property value. Error code reported by Apple during last dep sync.
func (m *DepOnboardingSetting) SetLastSyncErrorCode(value *int32)() {
    err := m.GetBackingStore().Set("lastSyncErrorCode", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSyncTriggeredDateTime sets the lastSyncTriggeredDateTime property value. When Intune last requested a sync.
func (m *DepOnboardingSetting) SetLastSyncTriggeredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastSyncTriggeredDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DepOnboardingSetting) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DepOnboardingSetting) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetShareTokenWithSchoolDataSyncService sets the shareTokenWithSchoolDataSyncService property value. Whether or not the Dep token sharing is enabled with the School Data Sync service.
func (m *DepOnboardingSetting) SetShareTokenWithSchoolDataSyncService(value *bool)() {
    err := m.GetBackingStore().Set("shareTokenWithSchoolDataSyncService", value)
    if err != nil {
        panic(err)
    }
}
// SetSyncedDeviceCount sets the syncedDeviceCount property value. Gets synced device count
func (m *DepOnboardingSetting) SetSyncedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("syncedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTokenExpirationDateTime sets the tokenExpirationDateTime property value. When the token will expire.
func (m *DepOnboardingSetting) SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("tokenExpirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetTokenName sets the tokenName property value. Friendly Name for Dep Token
func (m *DepOnboardingSetting) SetTokenName(value *string)() {
    err := m.GetBackingStore().Set("tokenName", value)
    if err != nil {
        panic(err)
    }
}
// SetTokenType sets the tokenType property value. The tokenType property
func (m *DepOnboardingSetting) SetTokenType(value *DepTokenType)() {
    err := m.GetBackingStore().Set("tokenType", value)
    if err != nil {
        panic(err)
    }
}
// DepOnboardingSettingable 
type DepOnboardingSettingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppleIdentifier()(*string)
    GetDataSharingConsentGranted()(*bool)
    GetDefaultIosEnrollmentProfile()(DepIOSEnrollmentProfileable)
    GetDefaultMacOsEnrollmentProfile()(DepMacOSEnrollmentProfileable)
    GetEnrollmentProfiles()([]EnrollmentProfileable)
    GetImportedAppleDeviceIdentities()([]ImportedAppleDeviceIdentityable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSyncErrorCode()(*int32)
    GetLastSyncTriggeredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetRoleScopeTagIds()([]string)
    GetShareTokenWithSchoolDataSyncService()(*bool)
    GetSyncedDeviceCount()(*int32)
    GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTokenName()(*string)
    GetTokenType()(*DepTokenType)
    SetAppleIdentifier(value *string)()
    SetDataSharingConsentGranted(value *bool)()
    SetDefaultIosEnrollmentProfile(value DepIOSEnrollmentProfileable)()
    SetDefaultMacOsEnrollmentProfile(value DepMacOSEnrollmentProfileable)()
    SetEnrollmentProfiles(value []EnrollmentProfileable)()
    SetImportedAppleDeviceIdentities(value []ImportedAppleDeviceIdentityable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSyncErrorCode(value *int32)()
    SetLastSyncTriggeredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetShareTokenWithSchoolDataSyncService(value *bool)()
    SetSyncedDeviceCount(value *int32)()
    SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTokenName(value *string)()
    SetTokenType(value *DepTokenType)()
}
