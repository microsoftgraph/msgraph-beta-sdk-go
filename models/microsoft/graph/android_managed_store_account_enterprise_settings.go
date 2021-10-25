package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AndroidManagedStoreAccountEnterpriseSettings struct {
    Entity
    androidDeviceOwnerFullyManagedEnrollmentEnabled *bool;
    bindStatus *AndroidManagedStoreAccountBindStatus;
    companyCodes []AndroidEnrollmentCompanyCode;
    deviceOwnerManagementEnabled *bool;
    enrollmentTarget *AndroidManagedStoreAccountEnrollmentTarget;
    lastAppSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastAppSyncStatus *AndroidManagedStoreAccountAppSyncStatus;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    ownerOrganizationName *string;
    ownerUserPrincipalName *string;
    targetGroupIds []string;
}
func NewAndroidManagedStoreAccountEnterpriseSettings()(*AndroidManagedStoreAccountEnterpriseSettings) {
    m := &AndroidManagedStoreAccountEnterpriseSettings{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetAndroidDeviceOwnerFullyManagedEnrollmentEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceOwnerFullyManagedEnrollmentEnabled
    }
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetBindStatus()(*AndroidManagedStoreAccountBindStatus) {
    if m == nil {
        return nil
    } else {
        return m.bindStatus
    }
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetCompanyCodes()([]AndroidEnrollmentCompanyCode) {
    if m == nil {
        return nil
    } else {
        return m.companyCodes
    }
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetDeviceOwnerManagementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceOwnerManagementEnabled
    }
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetEnrollmentTarget()(*AndroidManagedStoreAccountEnrollmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentTarget
    }
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetLastAppSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastAppSyncDateTime
    }
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetLastAppSyncStatus()(*AndroidManagedStoreAccountAppSyncStatus) {
    if m == nil {
        return nil
    } else {
        return m.lastAppSyncStatus
    }
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetOwnerOrganizationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerOrganizationName
    }
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetOwnerUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerUserPrincipalName
    }
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetTargetGroupIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupIds
    }
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidDeviceOwnerFullyManagedEnrollmentEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAndroidDeviceOwnerFullyManagedEnrollmentEnabled(val)
        return nil
    }
    res["bindStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAccountBindStatus)
        if err != nil {
            return err
        }
        cast := val.(AndroidManagedStoreAccountBindStatus)
        m.SetBindStatus(&cast)
        return nil
    }
    res["companyCodes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidEnrollmentCompanyCode() })
        if err != nil {
            return err
        }
        res := make([]AndroidEnrollmentCompanyCode, len(val))
        for i, v := range val {
            res[i] = *(v.(*AndroidEnrollmentCompanyCode))
        }
        m.SetCompanyCodes(res)
        return nil
    }
    res["deviceOwnerManagementEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDeviceOwnerManagementEnabled(val)
        return nil
    }
    res["enrollmentTarget"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAccountEnrollmentTarget)
        if err != nil {
            return err
        }
        cast := val.(AndroidManagedStoreAccountEnrollmentTarget)
        m.SetEnrollmentTarget(&cast)
        return nil
    }
    res["lastAppSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastAppSyncDateTime(val)
        return nil
    }
    res["lastAppSyncStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedStoreAccountAppSyncStatus)
        if err != nil {
            return err
        }
        cast := val.(AndroidManagedStoreAccountAppSyncStatus)
        m.SetLastAppSyncStatus(&cast)
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
    res["ownerOrganizationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOwnerOrganizationName(val)
        return nil
    }
    res["ownerUserPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOwnerUserPrincipalName(val)
        return nil
    }
    res["targetGroupIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetTargetGroupIds(res)
        return nil
    }
    return res
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) IsNil()(bool) {
    return m == nil
}
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
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetAndroidDeviceOwnerFullyManagedEnrollmentEnabled(value *bool)() {
    m.androidDeviceOwnerFullyManagedEnrollmentEnabled = value
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetBindStatus(value *AndroidManagedStoreAccountBindStatus)() {
    m.bindStatus = value
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetCompanyCodes(value []AndroidEnrollmentCompanyCode)() {
    m.companyCodes = value
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetDeviceOwnerManagementEnabled(value *bool)() {
    m.deviceOwnerManagementEnabled = value
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetEnrollmentTarget(value *AndroidManagedStoreAccountEnrollmentTarget)() {
    m.enrollmentTarget = value
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetLastAppSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastAppSyncDateTime = value
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetLastAppSyncStatus(value *AndroidManagedStoreAccountAppSyncStatus)() {
    m.lastAppSyncStatus = value
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetOwnerOrganizationName(value *string)() {
    m.ownerOrganizationName = value
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetOwnerUserPrincipalName(value *string)() {
    m.ownerUserPrincipalName = value
}
func (m *AndroidManagedStoreAccountEnterpriseSettings) SetTargetGroupIds(value []string)() {
    m.targetGroupIds = value
}
