package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AppleUserInitiatedEnrollmentProfile struct {
    Entity
    assignments []AppleEnrollmentProfileAssignment;
    availableEnrollmentTypeOptions []AppleOwnerTypeEnrollmentType;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    defaultEnrollmentType *AppleUserInitiatedEnrollmentType;
    description *string;
    displayName *string;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    platform *DevicePlatformType;
    priority *int32;
}
func NewAppleUserInitiatedEnrollmentProfile()(*AppleUserInitiatedEnrollmentProfile) {
    m := &AppleUserInitiatedEnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AppleUserInitiatedEnrollmentProfile) GetAssignments()([]AppleEnrollmentProfileAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
func (m *AppleUserInitiatedEnrollmentProfile) GetAvailableEnrollmentTypeOptions()([]AppleOwnerTypeEnrollmentType) {
    if m == nil {
        return nil
    } else {
        return m.availableEnrollmentTypeOptions
    }
}
func (m *AppleUserInitiatedEnrollmentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *AppleUserInitiatedEnrollmentProfile) GetDefaultEnrollmentType()(*AppleUserInitiatedEnrollmentType) {
    if m == nil {
        return nil
    } else {
        return m.defaultEnrollmentType
    }
}
func (m *AppleUserInitiatedEnrollmentProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *AppleUserInitiatedEnrollmentProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *AppleUserInitiatedEnrollmentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *AppleUserInitiatedEnrollmentProfile) GetPlatform()(*DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
func (m *AppleUserInitiatedEnrollmentProfile) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
func (m *AppleUserInitiatedEnrollmentProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppleEnrollmentProfileAssignment() })
        if err != nil {
            return err
        }
        res := make([]AppleEnrollmentProfileAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*AppleEnrollmentProfileAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["availableEnrollmentTypeOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppleOwnerTypeEnrollmentType() })
        if err != nil {
            return err
        }
        res := make([]AppleOwnerTypeEnrollmentType, len(val))
        for i, v := range val {
            res[i] = *(v.(*AppleOwnerTypeEnrollmentType))
        }
        m.SetAvailableEnrollmentTypeOptions(res)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["defaultEnrollmentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppleUserInitiatedEnrollmentType)
        if err != nil {
            return err
        }
        cast := val.(AppleUserInitiatedEnrollmentType)
        m.SetDefaultEnrollmentType(&cast)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
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
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        cast := val.(DevicePlatformType)
        m.SetPlatform(&cast)
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPriority(val)
        return nil
    }
    return res
}
func (m *AppleUserInitiatedEnrollmentProfile) IsNil()(bool) {
    return m == nil
}
func (m *AppleUserInitiatedEnrollmentProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAvailableEnrollmentTypeOptions()))
        for i, v := range m.GetAvailableEnrollmentTypeOptions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("availableEnrollmentTypeOptions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultEnrollmentType() != nil {
        cast := m.GetDefaultEnrollmentType().String()
        err = writer.WriteStringValue("defaultEnrollmentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
    if m.GetPlatform() != nil {
        cast := m.GetPlatform().String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AppleUserInitiatedEnrollmentProfile) SetAssignments(value []AppleEnrollmentProfileAssignment)() {
    m.assignments = value
}
func (m *AppleUserInitiatedEnrollmentProfile) SetAvailableEnrollmentTypeOptions(value []AppleOwnerTypeEnrollmentType)() {
    m.availableEnrollmentTypeOptions = value
}
func (m *AppleUserInitiatedEnrollmentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *AppleUserInitiatedEnrollmentProfile) SetDefaultEnrollmentType(value *AppleUserInitiatedEnrollmentType)() {
    m.defaultEnrollmentType = value
}
func (m *AppleUserInitiatedEnrollmentProfile) SetDescription(value *string)() {
    m.description = value
}
func (m *AppleUserInitiatedEnrollmentProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *AppleUserInitiatedEnrollmentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *AppleUserInitiatedEnrollmentProfile) SetPlatform(value *DevicePlatformType)() {
    m.platform = value
}
func (m *AppleUserInitiatedEnrollmentProfile) SetPriority(value *int32)() {
    m.priority = value
}
