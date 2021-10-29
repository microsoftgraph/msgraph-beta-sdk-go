package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPcUserSettingAssignment struct {
    Entity
    // The date and time this assignment was created. The Timestamp type represents the date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this: '2014-01-01T00:00:00Z'.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The assignment target for the user setting. Currently, the only target supported for this user setting is a user group. For details, see cloudPcManagementGroupAssignmentTarget.
    target *CloudPcManagementAssignmentTarget;
}
// Instantiates a new cloudPcUserSettingAssignment and sets the default values.
func NewCloudPcUserSettingAssignment()(*CloudPcUserSettingAssignment) {
    m := &CloudPcUserSettingAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdDateTime property value. The date and time this assignment was created. The Timestamp type represents the date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this: '2014-01-01T00:00:00Z'.
func (m *CloudPcUserSettingAssignment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the target property value. The assignment target for the user setting. Currently, the only target supported for this user setting is a user group. For details, see cloudPcManagementGroupAssignmentTarget.
func (m *CloudPcUserSettingAssignment) GetTarget()(*CloudPcManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// The deserialization information for the current model
func (m *CloudPcUserSettingAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcManagementAssignmentTarget() })
        if err != nil {
            return err
        }
        m.SetTarget(val.(*CloudPcManagementAssignmentTarget))
        return nil
    }
    return res
}
func (m *CloudPcUserSettingAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CloudPcUserSettingAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the createdDateTime property value. The date and time this assignment was created. The Timestamp type represents the date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this: '2014-01-01T00:00:00Z'.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *CloudPcUserSettingAssignment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the target property value. The assignment target for the user setting. Currently, the only target supported for this user setting is a user group. For details, see cloudPcManagementGroupAssignmentTarget.
// Parameters:
//  - value : Value to set for the target property.
func (m *CloudPcUserSettingAssignment) SetTarget(value *CloudPcManagementAssignmentTarget)() {
    m.target = value
}
