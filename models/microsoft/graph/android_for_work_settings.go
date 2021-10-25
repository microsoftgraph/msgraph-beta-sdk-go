package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AndroidForWorkSettings struct {
    Entity
    bindStatus *AndroidForWorkBindStatus;
    deviceOwnerManagementEnabled *bool;
    enrollmentTarget *AndroidForWorkEnrollmentTarget;
    lastAppSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastAppSyncStatus *AndroidForWorkSyncStatus;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    ownerOrganizationName *string;
    ownerUserPrincipalName *string;
    targetGroupIds []string;
}
func NewAndroidForWorkSettings()(*AndroidForWorkSettings) {
    m := &AndroidForWorkSettings{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AndroidForWorkSettings) GetBindStatus()(*AndroidForWorkBindStatus) {
    if m == nil {
        return nil
    } else {
        return m.bindStatus
    }
}
func (m *AndroidForWorkSettings) GetDeviceOwnerManagementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceOwnerManagementEnabled
    }
}
func (m *AndroidForWorkSettings) GetEnrollmentTarget()(*AndroidForWorkEnrollmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentTarget
    }
}
func (m *AndroidForWorkSettings) GetLastAppSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastAppSyncDateTime
    }
}
func (m *AndroidForWorkSettings) GetLastAppSyncStatus()(*AndroidForWorkSyncStatus) {
    if m == nil {
        return nil
    } else {
        return m.lastAppSyncStatus
    }
}
func (m *AndroidForWorkSettings) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *AndroidForWorkSettings) GetOwnerOrganizationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerOrganizationName
    }
}
func (m *AndroidForWorkSettings) GetOwnerUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerUserPrincipalName
    }
}
func (m *AndroidForWorkSettings) GetTargetGroupIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.targetGroupIds
    }
}
func (m *AndroidForWorkSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bindStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidForWorkBindStatus)
        if err != nil {
            return err
        }
        cast := val.(AndroidForWorkBindStatus)
        m.SetBindStatus(&cast)
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
        val, err := n.GetEnumValue(ParseAndroidForWorkEnrollmentTarget)
        if err != nil {
            return err
        }
        cast := val.(AndroidForWorkEnrollmentTarget)
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
        val, err := n.GetEnumValue(ParseAndroidForWorkSyncStatus)
        if err != nil {
            return err
        }
        cast := val.(AndroidForWorkSyncStatus)
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
func (m *AndroidForWorkSettings) IsNil()(bool) {
    return m == nil
}
func (m *AndroidForWorkSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBindStatus() != nil {
        cast := m.GetBindStatus().String()
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
func (m *AndroidForWorkSettings) SetBindStatus(value *AndroidForWorkBindStatus)() {
    m.bindStatus = value
}
func (m *AndroidForWorkSettings) SetDeviceOwnerManagementEnabled(value *bool)() {
    m.deviceOwnerManagementEnabled = value
}
func (m *AndroidForWorkSettings) SetEnrollmentTarget(value *AndroidForWorkEnrollmentTarget)() {
    m.enrollmentTarget = value
}
func (m *AndroidForWorkSettings) SetLastAppSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastAppSyncDateTime = value
}
func (m *AndroidForWorkSettings) SetLastAppSyncStatus(value *AndroidForWorkSyncStatus)() {
    m.lastAppSyncStatus = value
}
func (m *AndroidForWorkSettings) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *AndroidForWorkSettings) SetOwnerOrganizationName(value *string)() {
    m.ownerOrganizationName = value
}
func (m *AndroidForWorkSettings) SetOwnerUserPrincipalName(value *string)() {
    m.ownerUserPrincipalName = value
}
func (m *AndroidForWorkSettings) SetTargetGroupIds(value []string)() {
    m.targetGroupIds = value
}
