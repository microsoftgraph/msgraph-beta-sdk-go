package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAccountEnterpriseSettings enterprise settings for an Android managed store account.
type AndroidManagedStoreAccountEnterpriseSettings struct {
    Entity
}
// NewAndroidManagedStoreAccountEnterpriseSettings instantiates a new androidManagedStoreAccountEnterpriseSettings and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettings()(*AndroidManagedStoreAccountEnterpriseSettings) {
    m := &AndroidManagedStoreAccountEnterpriseSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAccountEnterpriseSettings(), nil
}
// GetAndroidDeviceOwnerFullyManagedEnrollmentEnabled gets the androidDeviceOwnerFullyManagedEnrollmentEnabled property value. Company codes for AndroidManagedStoreAccountEnterpriseSettings
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetAndroidDeviceOwnerFullyManagedEnrollmentEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("androidDeviceOwnerFullyManagedEnrollmentEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBindStatus gets the bindStatus property value. Bind status of the tenant with the Google EMM API
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetBindStatus()(*AndroidManagedStoreAccountBindStatus) {
    val, err := m.GetBackingStore().Get("bindStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidManagedStoreAccountBindStatus)
    }
    return nil
}
// GetCompanyCodes gets the companyCodes property value. Company codes for AndroidManagedStoreAccountEnterpriseSettings
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetCompanyCodes()([]AndroidEnrollmentCompanyCodeable) {
    val, err := m.GetBackingStore().Get("companyCodes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidEnrollmentCompanyCodeable)
    }
    return nil
}
// GetDeviceOwnerManagementEnabled gets the deviceOwnerManagementEnabled property value. Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetDeviceOwnerManagementEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("deviceOwnerManagementEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnrollmentTarget gets the enrollmentTarget property value. Android for Work device management targeting type for the account
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetEnrollmentTarget()(*AndroidManagedStoreAccountEnrollmentTarget) {
    val, err := m.GetBackingStore().Get("enrollmentTarget")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidManagedStoreAccountEnrollmentTarget)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidDeviceOwnerFullyManagedEnrollmentEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDeviceOwnerFullyManagedEnrollmentEnabled(val)
        }
        return nil
    }
    res["bindStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAccountBindStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBindStatus(val.(*AndroidManagedStoreAccountBindStatus))
        }
        return nil
    }
    res["companyCodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidEnrollmentCompanyCodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidEnrollmentCompanyCodeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AndroidEnrollmentCompanyCodeable)
                }
            }
            m.SetCompanyCodes(res)
        }
        return nil
    }
    res["deviceOwnerManagementEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOwnerManagementEnabled(val)
        }
        return nil
    }
    res["enrollmentTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAccountEnrollmentTarget)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentTarget(val.(*AndroidManagedStoreAccountEnrollmentTarget))
        }
        return nil
    }
    res["lastAppSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAppSyncDateTime(val)
        }
        return nil
    }
    res["lastAppSyncStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAccountAppSyncStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAppSyncStatus(val.(*AndroidManagedStoreAccountAppSyncStatus))
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
    res["managedGooglePlayInitialScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetManagedGooglePlayInitialScopeTagIds(res)
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
    res["ownerOrganizationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerOrganizationName(val)
        }
        return nil
    }
    res["ownerUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerUserPrincipalName(val)
        }
        return nil
    }
    res["targetGroupIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTargetGroupIds(res)
        }
        return nil
    }
    return res
}
// GetLastAppSyncDateTime gets the lastAppSyncDateTime property value. Last completion time for app sync
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetLastAppSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastAppSyncDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastAppSyncStatus gets the lastAppSyncStatus property value. Sync status of the tenant with the Google EMM API
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetLastAppSyncStatus()(*AndroidManagedStoreAccountAppSyncStatus) {
    val, err := m.GetBackingStore().Get("lastAppSyncStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidManagedStoreAccountAppSyncStatus)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modification time for Android enterprise settings
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetManagedGooglePlayInitialScopeTagIds gets the managedGooglePlayInitialScopeTagIds property value. Initial scope tags for MGP apps
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetManagedGooglePlayInitialScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("managedGooglePlayInitialScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOwnerOrganizationName gets the ownerOrganizationName property value. Organization name used when onboarding Android Enterprise
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetOwnerOrganizationName()(*string) {
    val, err := m.GetBackingStore().Get("ownerOrganizationName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOwnerUserPrincipalName gets the ownerUserPrincipalName property value. Owner UPN that created the enterprise
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetOwnerUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("ownerUserPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetGroupIds gets the targetGroupIds property value. Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetTargetGroupIds()([]string) {
    val, err := m.GetBackingStore().Get("targetGroupIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAccountEnterpriseSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("androidDeviceOwnerFullyManagedEnrollmentEnabled", m.GetAndroidDeviceOwnerFullyManagedEnrollmentEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetBindStatus() != nil {
        cast := (*m.GetBindStatus()).String()
        err = writer.WriteStringValue("bindStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCompanyCodes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompanyCodes()))
        for i, v := range m.GetCompanyCodes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("companyCodes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceOwnerManagementEnabled", m.GetDeviceOwnerManagementEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentTarget() != nil {
        cast := (*m.GetEnrollmentTarget()).String()
        err = writer.WriteStringValue("enrollmentTarget", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastAppSyncDateTime", m.GetLastAppSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLastAppSyncStatus() != nil {
        cast := (*m.GetLastAppSyncStatus()).String()
        err = writer.WriteStringValue("lastAppSyncStatus", &cast)
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
    if m.GetManagedGooglePlayInitialScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("managedGooglePlayInitialScopeTagIds", m.GetManagedGooglePlayInitialScopeTagIds())
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
    {
        err = writer.WriteStringValue("ownerOrganizationName", m.GetOwnerOrganizationName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerUserPrincipalName", m.GetOwnerUserPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetTargetGroupIds() != nil {
        err = writer.WriteCollectionOfStringValues("targetGroupIds", m.GetTargetGroupIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAndroidDeviceOwnerFullyManagedEnrollmentEnabled sets the androidDeviceOwnerFullyManagedEnrollmentEnabled property value. Company codes for AndroidManagedStoreAccountEnterpriseSettings
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetAndroidDeviceOwnerFullyManagedEnrollmentEnabled(value *bool)() {
    err := m.GetBackingStore().Set("androidDeviceOwnerFullyManagedEnrollmentEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetBindStatus sets the bindStatus property value. Bind status of the tenant with the Google EMM API
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetBindStatus(value *AndroidManagedStoreAccountBindStatus)() {
    err := m.GetBackingStore().Set("bindStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetCompanyCodes sets the companyCodes property value. Company codes for AndroidManagedStoreAccountEnterpriseSettings
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetCompanyCodes(value []AndroidEnrollmentCompanyCodeable)() {
    err := m.GetBackingStore().Set("companyCodes", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceOwnerManagementEnabled sets the deviceOwnerManagementEnabled property value. Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetDeviceOwnerManagementEnabled(value *bool)() {
    err := m.GetBackingStore().Set("deviceOwnerManagementEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentTarget sets the enrollmentTarget property value. Android for Work device management targeting type for the account
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetEnrollmentTarget(value *AndroidManagedStoreAccountEnrollmentTarget)() {
    err := m.GetBackingStore().Set("enrollmentTarget", value)
    if err != nil {
        panic(err)
    }
}
// SetLastAppSyncDateTime sets the lastAppSyncDateTime property value. Last completion time for app sync
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetLastAppSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastAppSyncDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastAppSyncStatus sets the lastAppSyncStatus property value. Sync status of the tenant with the Google EMM API
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetLastAppSyncStatus(value *AndroidManagedStoreAccountAppSyncStatus)() {
    err := m.GetBackingStore().Set("lastAppSyncStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modification time for Android enterprise settings
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedGooglePlayInitialScopeTagIds sets the managedGooglePlayInitialScopeTagIds property value. Initial scope tags for MGP apps
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetManagedGooglePlayInitialScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("managedGooglePlayInitialScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnerOrganizationName sets the ownerOrganizationName property value. Organization name used when onboarding Android Enterprise
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetOwnerOrganizationName(value *string)() {
    err := m.GetBackingStore().Set("ownerOrganizationName", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnerUserPrincipalName sets the ownerUserPrincipalName property value. Owner UPN that created the enterprise
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetOwnerUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("ownerUserPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetGroupIds sets the targetGroupIds property value. Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetTargetGroupIds(value []string)() {
    err := m.GetBackingStore().Set("targetGroupIds", value)
    if err != nil {
        panic(err)
    }
}
// AndroidManagedStoreAccountEnterpriseSettingsable 
type AndroidManagedStoreAccountEnterpriseSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAndroidDeviceOwnerFullyManagedEnrollmentEnabled()(*bool)
    GetBindStatus()(*AndroidManagedStoreAccountBindStatus)
    GetCompanyCodes()([]AndroidEnrollmentCompanyCodeable)
    GetDeviceOwnerManagementEnabled()(*bool)
    GetEnrollmentTarget()(*AndroidManagedStoreAccountEnrollmentTarget)
    GetLastAppSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastAppSyncStatus()(*AndroidManagedStoreAccountAppSyncStatus)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedGooglePlayInitialScopeTagIds()([]string)
    GetOdataType()(*string)
    GetOwnerOrganizationName()(*string)
    GetOwnerUserPrincipalName()(*string)
    GetTargetGroupIds()([]string)
    SetAndroidDeviceOwnerFullyManagedEnrollmentEnabled(value *bool)()
    SetBindStatus(value *AndroidManagedStoreAccountBindStatus)()
    SetCompanyCodes(value []AndroidEnrollmentCompanyCodeable)()
    SetDeviceOwnerManagementEnabled(value *bool)()
    SetEnrollmentTarget(value *AndroidManagedStoreAccountEnrollmentTarget)()
    SetLastAppSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastAppSyncStatus(value *AndroidManagedStoreAccountAppSyncStatus)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedGooglePlayInitialScopeTagIds(value []string)()
    SetOdataType(value *string)()
    SetOwnerOrganizationName(value *string)()
    SetOwnerUserPrincipalName(value *string)()
    SetTargetGroupIds(value []string)()
}
