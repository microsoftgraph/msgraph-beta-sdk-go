package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AndroidForWorkSettings 
type AndroidForWorkSettings struct {
    Entity
    // Bind status of the tenant with the Google EMM API. Possible values are: notBound, bound, boundAndValidated, unbinding.
    bindStatus *AndroidForWorkBindStatus;
    // Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
    deviceOwnerManagementEnabled *bool;
    // Indicates which users can enroll devices in Android for Work device management. Possible values are: none, all, targeted, targetedAsEnrollmentRestrictions.
    enrollmentTarget *AndroidForWorkEnrollmentTarget;
    // Last completion time for app sync
    lastAppSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Last application sync result. Possible values are: success, credentialsNotValid, androidForWorkApiError, managementServiceError, unknownError, none.
    lastAppSyncStatus *AndroidForWorkSyncStatus;
    // Last modification time for Android for Work settings
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Organization name used when onboarding Android for Work
    ownerOrganizationName *string;
    // Owner UPN that created the enterprise
    ownerUserPrincipalName *string;
    // Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
    targetGroupIds []string;
}
// NewAndroidForWorkSettings instantiates a new androidForWorkSettings and sets the default values.
func NewAndroidForWorkSettings()(*AndroidForWorkSettings) {
    m := &AndroidForWorkSettings{
        Entity: *NewEntity(),
    }
    return m
}
// GetBindStatus gets the bindStatus property value. Bind status of the tenant with the Google EMM API. Possible values are: notBound, bound, boundAndValidated, unbinding.
func (m *AndroidForWorkSettings) GetBindStatus()(*AndroidForWorkBindStatus) {
    if m == nil {
        return nil
    } else {
        return m.bindStatus
    }
}
// GetDeviceOwnerManagementEnabled gets the deviceOwnerManagementEnabled property value. Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
func (m *AndroidForWorkSettings) GetDeviceOwnerManagementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceOwnerManagementEnabled
    }
}
// GetEnrollmentTarget gets the enrollmentTarget property value. Indicates which users can enroll devices in Android for Work device management. Possible values are: none, all, targeted, targetedAsEnrollmentRestrictions.
func (m *AndroidForWorkSettings) GetEnrollmentTarget()(*AndroidForWorkEnrollmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentTarget
    }
}
// GetLastAppSyncDateTime gets the lastAppSyncDateTime property value. Last completion time for app sync
func (m *AndroidForWorkSettings) GetLastAppSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastAppSyncDateTime
    }
}
// GetLastAppSyncStatus gets the lastAppSyncStatus property value. Last application sync result. Possible values are: success, credentialsNotValid, androidForWorkApiError, managementServiceError, unknownError, none.
func (m *AndroidForWorkSettings) GetLastAppSyncStatus()(*AndroidForWorkSyncStatus) {
    if m == nil {
        return nil
    } else {
        return m.lastAppSyncStatus
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modification time for Android for Work settings
func (m *AndroidForWorkSettings) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetOwnerOrganizationName gets the ownerOrganizationName property value. Organization name used when onboarding Android for Work
func (m *AndroidForWorkSettings) GetOwnerOrganizationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerOrganizationName
    }
}
// GetOwnerUserPrincipalName gets the ownerUserPrincipalName property value. Owner UPN that created the enterprise
func (m *AndroidForWorkSettings) GetOwnerUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerUserPrincipalName
    }
}
// GetTargetGroupIds gets the targetGroupIds property value. Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
func (m *AndroidForWorkSettings) GetTargetGroupIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bindStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidForWorkBindStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBindStatus(val.(*AndroidForWorkBindStatus))
        }
        return nil
    }
    res["deviceOwnerManagementEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOwnerManagementEnabled(val)
        }
        return nil
    }
    res["enrollmentTarget"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidForWorkEnrollmentTarget)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentTarget(val.(*AndroidForWorkEnrollmentTarget))
        }
        return nil
    }
    res["lastAppSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAppSyncDateTime(val)
        }
        return nil
    }
    res["lastAppSyncStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidForWorkSyncStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAppSyncStatus(val.(*AndroidForWorkSyncStatus))
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
    res["ownerOrganizationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerOrganizationName(val)
        }
        return nil
    }
    res["ownerUserPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerUserPrincipalName(val)
        }
        return nil
    }
    res["targetGroupIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTargetGroupIds(res)
        }
        return nil
    }
    return res
}
func (m *AndroidForWorkSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AndroidForWorkSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetBindStatus sets the bindStatus property value. Bind status of the tenant with the Google EMM API. Possible values are: notBound, bound, boundAndValidated, unbinding.
func (m *AndroidForWorkSettings) SetBindStatus(value *AndroidForWorkBindStatus)() {
    if m != nil {
        m.bindStatus = value
    }
}
// SetDeviceOwnerManagementEnabled sets the deviceOwnerManagementEnabled property value. Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
func (m *AndroidForWorkSettings) SetDeviceOwnerManagementEnabled(value *bool)() {
    if m != nil {
        m.deviceOwnerManagementEnabled = value
    }
}
// SetEnrollmentTarget sets the enrollmentTarget property value. Indicates which users can enroll devices in Android for Work device management. Possible values are: none, all, targeted, targetedAsEnrollmentRestrictions.
func (m *AndroidForWorkSettings) SetEnrollmentTarget(value *AndroidForWorkEnrollmentTarget)() {
    if m != nil {
        m.enrollmentTarget = value
    }
}
// SetLastAppSyncDateTime sets the lastAppSyncDateTime property value. Last completion time for app sync
func (m *AndroidForWorkSettings) SetLastAppSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastAppSyncDateTime = value
    }
}
// SetLastAppSyncStatus sets the lastAppSyncStatus property value. Last application sync result. Possible values are: success, credentialsNotValid, androidForWorkApiError, managementServiceError, unknownError, none.
func (m *AndroidForWorkSettings) SetLastAppSyncStatus(value *AndroidForWorkSyncStatus)() {
    if m != nil {
        m.lastAppSyncStatus = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modification time for Android for Work settings
func (m *AndroidForWorkSettings) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetOwnerOrganizationName sets the ownerOrganizationName property value. Organization name used when onboarding Android for Work
func (m *AndroidForWorkSettings) SetOwnerOrganizationName(value *string)() {
    if m != nil {
        m.ownerOrganizationName = value
    }
}
// SetOwnerUserPrincipalName sets the ownerUserPrincipalName property value. Owner UPN that created the enterprise
func (m *AndroidForWorkSettings) SetOwnerUserPrincipalName(value *string)() {
    if m != nil {
        m.ownerUserPrincipalName = value
    }
}
// SetTargetGroupIds sets the targetGroupIds property value. Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
func (m *AndroidForWorkSettings) SetTargetGroupIds(value []string)() {
    if m != nil {
        m.targetGroupIds = value
    }
}
