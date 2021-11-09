package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AndroidManagedStoreAccountEnterpriseSettings struct {
    Entity
    // Company codes for AndroidManagedStoreAccountEnterpriseSettings
    androidDeviceOwnerFullyManagedEnrollmentEnabled *bool;
    // Bind status of the tenant with the Google EMM API. Possible values are: notBound, bound, boundAndValidated, unbinding.
    bindStatus *AndroidManagedStoreAccountBindStatus;
    // Company codes for AndroidManagedStoreAccountEnterpriseSettings
    companyCodes []AndroidEnrollmentCompanyCode;
    // Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
    deviceOwnerManagementEnabled *bool;
    // Indicates which users can enroll devices in Android Enterprise device management. Possible values are: none, all, targeted, targetedAsEnrollmentRestrictions.
    enrollmentTarget *AndroidManagedStoreAccountEnrollmentTarget;
    // Last completion time for app sync
    lastAppSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Last application sync result. Possible values are: success, credentialsNotValid, androidForWorkApiError, managementServiceError, unknownError, none.
    lastAppSyncStatus *AndroidManagedStoreAccountAppSyncStatus;
    // Last modification time for Android enterprise settings
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Initial scope tags for MGP apps
    managedGooglePlayInitialScopeTagIds []string;
    // Organization name used when onboarding Android Enterprise
    ownerOrganizationName *string;
    // Owner UPN that created the enterprise
    ownerUserPrincipalName *string;
    // Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
    targetGroupIds []string;
}
// Instantiates a new androidManagedStoreAccountEnterpriseSettings and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettings()(*AndroidManagedStoreAccountEnterpriseSettings) {
    m := &AndroidManagedStoreAccountEnterpriseSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the androidDeviceOwnerFullyManagedEnrollmentEnabled property value. Company codes for AndroidManagedStoreAccountEnterpriseSettings
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetAndroidDeviceOwnerFullyManagedEnrollmentEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceOwnerFullyManagedEnrollmentEnabled
    }
}
// Gets the bindStatus property value. Bind status of the tenant with the Google EMM API. Possible values are: notBound, bound, boundAndValidated, unbinding.
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetBindStatus()(*AndroidManagedStoreAccountBindStatus) {
    if m == nil {
        return nil
    } else {
        return m.bindStatus
    }
}
// Gets the companyCodes property value. Company codes for AndroidManagedStoreAccountEnterpriseSettings
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetCompanyCodes()([]AndroidEnrollmentCompanyCode) {
    if m == nil {
        return nil
    } else {
        return m.companyCodes
    }
}
// Gets the deviceOwnerManagementEnabled property value. Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetDeviceOwnerManagementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceOwnerManagementEnabled
    }
}
// Gets the enrollmentTarget property value. Indicates which users can enroll devices in Android Enterprise device management. Possible values are: none, all, targeted, targetedAsEnrollmentRestrictions.
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetEnrollmentTarget()(*AndroidManagedStoreAccountEnrollmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentTarget
    }
}
// Gets the lastAppSyncDateTime property value. Last completion time for app sync
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetLastAppSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastAppSyncDateTime
    }
}
// Gets the lastAppSyncStatus property value. Last application sync result. Possible values are: success, credentialsNotValid, androidForWorkApiError, managementServiceError, unknownError, none.
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetLastAppSyncStatus()(*AndroidManagedStoreAccountAppSyncStatus) {
    if m == nil {
        return nil
    } else {
        return m.lastAppSyncStatus
    }
}
// Gets the lastModifiedDateTime property value. Last modification time for Android enterprise settings
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the managedGooglePlayInitialScopeTagIds property value. Initial scope tags for MGP apps
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetManagedGooglePlayInitialScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.managedGooglePlayInitialScopeTagIds
    }
}
// Gets the ownerOrganizationName property value. Organization name used when onboarding Android Enterprise
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetOwnerOrganizationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerOrganizationName
    }
}
// Gets the ownerUserPrincipalName property value. Owner UPN that created the enterprise
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetOwnerUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerUserPrincipalName
    }
}
// Gets the targetGroupIds property value. Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetTargetGroupIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupIds
    }
}
// The deserialization information for the current model
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidDeviceOwnerFullyManagedEnrollmentEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDeviceOwnerFullyManagedEnrollmentEnabled(val)
        }
        return nil
    }
    res["bindStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAccountBindStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AndroidManagedStoreAccountBindStatus)
            m.SetBindStatus(&cast)
        }
        return nil
    }
    res["companyCodes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidEnrollmentCompanyCode() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidEnrollmentCompanyCode, len(val))
            for i, v := range val {
                res[i] = *(v.(*AndroidEnrollmentCompanyCode))
            }
            m.SetCompanyCodes(res)
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
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAccountEnrollmentTarget)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AndroidManagedStoreAccountEnrollmentTarget)
            m.SetEnrollmentTarget(&cast)
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
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAccountAppSyncStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AndroidManagedStoreAccountAppSyncStatus)
            m.SetLastAppSyncStatus(&cast)
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
    res["managedGooglePlayInitialScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetManagedGooglePlayInitialScopeTagIds(res)
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
func (m *AndroidManagedStoreAccountEnterpriseSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AndroidManagedStoreAccountEnterpriseSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := m.GetBindStatus().String()
        err = writer.WriteStringValue("bindStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCompanyCodes()))
        for i, v := range m.GetCompanyCodes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := m.GetEnrollmentTarget().String()
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
        cast := m.GetLastAppSyncStatus().String()
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
        err = writer.WriteCollectionOfStringValues("managedGooglePlayInitialScopeTagIds", m.GetManagedGooglePlayInitialScopeTagIds())
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
    {
        err = writer.WriteCollectionOfStringValues("targetGroupIds", m.GetTargetGroupIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the androidDeviceOwnerFullyManagedEnrollmentEnabled property value. Company codes for AndroidManagedStoreAccountEnterpriseSettings
// Parameters:
//  - value : Value to set for the androidDeviceOwnerFullyManagedEnrollmentEnabled property.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetAndroidDeviceOwnerFullyManagedEnrollmentEnabled(value *bool)() {
    m.androidDeviceOwnerFullyManagedEnrollmentEnabled = value
}
// Sets the bindStatus property value. Bind status of the tenant with the Google EMM API. Possible values are: notBound, bound, boundAndValidated, unbinding.
// Parameters:
//  - value : Value to set for the bindStatus property.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetBindStatus(value *AndroidManagedStoreAccountBindStatus)() {
    m.bindStatus = value
}
// Sets the companyCodes property value. Company codes for AndroidManagedStoreAccountEnterpriseSettings
// Parameters:
//  - value : Value to set for the companyCodes property.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetCompanyCodes(value []AndroidEnrollmentCompanyCode)() {
    m.companyCodes = value
}
// Sets the deviceOwnerManagementEnabled property value. Indicates if this account is flighting for Android Device Owner Management with CloudDPC.
// Parameters:
//  - value : Value to set for the deviceOwnerManagementEnabled property.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetDeviceOwnerManagementEnabled(value *bool)() {
    m.deviceOwnerManagementEnabled = value
}
// Sets the enrollmentTarget property value. Indicates which users can enroll devices in Android Enterprise device management. Possible values are: none, all, targeted, targetedAsEnrollmentRestrictions.
// Parameters:
//  - value : Value to set for the enrollmentTarget property.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetEnrollmentTarget(value *AndroidManagedStoreAccountEnrollmentTarget)() {
    m.enrollmentTarget = value
}
// Sets the lastAppSyncDateTime property value. Last completion time for app sync
// Parameters:
//  - value : Value to set for the lastAppSyncDateTime property.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetLastAppSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastAppSyncDateTime = value
}
// Sets the lastAppSyncStatus property value. Last application sync result. Possible values are: success, credentialsNotValid, androidForWorkApiError, managementServiceError, unknownError, none.
// Parameters:
//  - value : Value to set for the lastAppSyncStatus property.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetLastAppSyncStatus(value *AndroidManagedStoreAccountAppSyncStatus)() {
    m.lastAppSyncStatus = value
}
// Sets the lastModifiedDateTime property value. Last modification time for Android enterprise settings
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the managedGooglePlayInitialScopeTagIds property value. Initial scope tags for MGP apps
// Parameters:
//  - value : Value to set for the managedGooglePlayInitialScopeTagIds property.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetManagedGooglePlayInitialScopeTagIds(value []string)() {
    m.managedGooglePlayInitialScopeTagIds = value
}
// Sets the ownerOrganizationName property value. Organization name used when onboarding Android Enterprise
// Parameters:
//  - value : Value to set for the ownerOrganizationName property.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetOwnerOrganizationName(value *string)() {
    m.ownerOrganizationName = value
}
// Sets the ownerUserPrincipalName property value. Owner UPN that created the enterprise
// Parameters:
//  - value : Value to set for the ownerUserPrincipalName property.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetOwnerUserPrincipalName(value *string)() {
    m.ownerUserPrincipalName = value
}
// Sets the targetGroupIds property value. Specifies which AAD groups can enroll devices in Android for Work device management if enrollmentTarget is set to 'Targeted'
// Parameters:
//  - value : Value to set for the targetGroupIds property.
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetTargetGroupIds(value []string)() {
    m.targetGroupIds = value
}
