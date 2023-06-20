package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkSettings 
type AndroidForWorkSettings struct {
    Entity
}
// NewAndroidForWorkSettings instantiates a new androidForWorkSettings and sets the default values.
func NewAndroidForWorkSettings()(*AndroidForWorkSettings) {
    m := &AndroidForWorkSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAndroidForWorkSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkSettings(), nil
}
// GetBindStatus gets the bindStatus property value. Bind status of the tenant with the Google EMM API
func (m *AndroidForWorkSettings) GetBindStatus()(*AndroidForWorkBindStatus) {
    val, err := m.GetBackingStore().Get("bindStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidForWorkBindStatus)
    }
    return nil
}
// GetDeviceOwnerManagementEnabled gets the deviceOwnerManagementEnabled property value. Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
func (m *AndroidForWorkSettings) GetDeviceOwnerManagementEnabled()(*bool) {
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
func (m *AndroidForWorkSettings) GetEnrollmentTarget()(*AndroidForWorkEnrollmentTarget) {
    val, err := m.GetBackingStore().Get("enrollmentTarget")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidForWorkEnrollmentTarget)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bindStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidForWorkBindStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBindStatus(val.(*AndroidForWorkBindStatus))
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
        val, err := n.GetEnumValue(ParseAndroidForWorkEnrollmentTarget)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentTarget(val.(*AndroidForWorkEnrollmentTarget))
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
        val, err := n.GetEnumValue(ParseAndroidForWorkSyncStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAppSyncStatus(val.(*AndroidForWorkSyncStatus))
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
func (m *AndroidForWorkSettings) GetLastAppSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *AndroidForWorkSettings) GetLastAppSyncStatus()(*AndroidForWorkSyncStatus) {
    val, err := m.GetBackingStore().Get("lastAppSyncStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidForWorkSyncStatus)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modification time for Android for Work settings
func (m *AndroidForWorkSettings) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOwnerOrganizationName gets the ownerOrganizationName property value. Organization name used when onboarding Android for Work
func (m *AndroidForWorkSettings) GetOwnerOrganizationName()(*string) {
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
func (m *AndroidForWorkSettings) GetOwnerUserPrincipalName()(*string) {
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
func (m *AndroidForWorkSettings) GetTargetGroupIds()([]string) {
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
func (m *AndroidForWorkSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBindStatus() != nil {
        cast := (*m.GetBindStatus()).String()
        err = writer.WriteStringValue("bindStatus", &cast)
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
// SetBindStatus sets the bindStatus property value. Bind status of the tenant with the Google EMM API
func (m *AndroidForWorkSettings) SetBindStatus(value *AndroidForWorkBindStatus)() {
    err := m.GetBackingStore().Set("bindStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceOwnerManagementEnabled sets the deviceOwnerManagementEnabled property value. Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
func (m *AndroidForWorkSettings) SetDeviceOwnerManagementEnabled(value *bool)() {
    err := m.GetBackingStore().Set("deviceOwnerManagementEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentTarget sets the enrollmentTarget property value. Android for Work device management targeting type for the account
func (m *AndroidForWorkSettings) SetEnrollmentTarget(value *AndroidForWorkEnrollmentTarget)() {
    err := m.GetBackingStore().Set("enrollmentTarget", value)
    if err != nil {
        panic(err)
    }
}
// SetLastAppSyncDateTime sets the lastAppSyncDateTime property value. Last completion time for app sync
func (m *AndroidForWorkSettings) SetLastAppSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastAppSyncDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastAppSyncStatus sets the lastAppSyncStatus property value. Sync status of the tenant with the Google EMM API
func (m *AndroidForWorkSettings) SetLastAppSyncStatus(value *AndroidForWorkSyncStatus)() {
    err := m.GetBackingStore().Set("lastAppSyncStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modification time for Android for Work settings
func (m *AndroidForWorkSettings) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnerOrganizationName sets the ownerOrganizationName property value. Organization name used when onboarding Android for Work
func (m *AndroidForWorkSettings) SetOwnerOrganizationName(value *string)() {
    err := m.GetBackingStore().Set("ownerOrganizationName", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnerUserPrincipalName sets the ownerUserPrincipalName property value. Owner UPN that created the enterprise
func (m *AndroidForWorkSettings) SetOwnerUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("ownerUserPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetGroupIds sets the targetGroupIds property value. Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
func (m *AndroidForWorkSettings) SetTargetGroupIds(value []string)() {
    err := m.GetBackingStore().Set("targetGroupIds", value)
    if err != nil {
        panic(err)
    }
}
// AndroidForWorkSettingsable 
type AndroidForWorkSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBindStatus()(*AndroidForWorkBindStatus)
    GetDeviceOwnerManagementEnabled()(*bool)
    GetEnrollmentTarget()(*AndroidForWorkEnrollmentTarget)
    GetLastAppSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastAppSyncStatus()(*AndroidForWorkSyncStatus)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOwnerOrganizationName()(*string)
    GetOwnerUserPrincipalName()(*string)
    GetTargetGroupIds()([]string)
    SetBindStatus(value *AndroidForWorkBindStatus)()
    SetDeviceOwnerManagementEnabled(value *bool)()
    SetEnrollmentTarget(value *AndroidForWorkEnrollmentTarget)()
    SetLastAppSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastAppSyncStatus(value *AndroidForWorkSyncStatus)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOwnerOrganizationName(value *string)()
    SetOwnerUserPrincipalName(value *string)()
    SetTargetGroupIds(value []string)()
}
