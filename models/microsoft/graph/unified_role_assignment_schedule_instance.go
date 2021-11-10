package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UnifiedRoleAssignmentScheduleInstance struct {
    UnifiedRoleScheduleInstanceBase
    // If the roleAssignmentScheduleInstance is activated by a roleEligibilityScheduleRequest, this is the link to the related schedule instance.
    activatedUsing *UnifiedRoleEligibilityScheduleInstance;
    // Type of the assignment. It can either be Assigned or Activated.
    assignmentType *string;
    // Time that the roleAssignmentInstance will expire
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Membership type of the assignment. It can either be Inherited, Direct, or Group.
    memberType *string;
    // ID of the roleAssignment in the directory
    roleAssignmentOriginId *string;
    // ID of the parent roleAssignmentSchedule for this instance
    roleAssignmentScheduleId *string;
    // Time that the roleAssignmentInstance will start
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new unifiedRoleAssignmentScheduleInstance and sets the default values.
func NewUnifiedRoleAssignmentScheduleInstance()(*UnifiedRoleAssignmentScheduleInstance) {
    m := &UnifiedRoleAssignmentScheduleInstance{
        UnifiedRoleScheduleInstanceBase: *NewUnifiedRoleScheduleInstanceBase(),
    }
    return m
}
// Gets the activatedUsing property value. If the roleAssignmentScheduleInstance is activated by a roleEligibilityScheduleRequest, this is the link to the related schedule instance.
func (m *UnifiedRoleAssignmentScheduleInstance) GetActivatedUsing()(*UnifiedRoleEligibilityScheduleInstance) {
    if m == nil {
        return nil
    } else {
        return m.activatedUsing
    }
}
// Gets the assignmentType property value. Type of the assignment. It can either be Assigned or Activated.
func (m *UnifiedRoleAssignmentScheduleInstance) GetAssignmentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentType
    }
}
// Gets the endDateTime property value. Time that the roleAssignmentInstance will expire
func (m *UnifiedRoleAssignmentScheduleInstance) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// Gets the memberType property value. Membership type of the assignment. It can either be Inherited, Direct, or Group.
func (m *UnifiedRoleAssignmentScheduleInstance) GetMemberType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberType
    }
}
// Gets the roleAssignmentOriginId property value. ID of the roleAssignment in the directory
func (m *UnifiedRoleAssignmentScheduleInstance) GetRoleAssignmentOriginId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentOriginId
    }
}
// Gets the roleAssignmentScheduleId property value. ID of the parent roleAssignmentSchedule for this instance
func (m *UnifiedRoleAssignmentScheduleInstance) GetRoleAssignmentScheduleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentScheduleId
    }
}
// Gets the startDateTime property value. Time that the roleAssignmentInstance will start
func (m *UnifiedRoleAssignmentScheduleInstance) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// The deserialization information for the current model
func (m *UnifiedRoleAssignmentScheduleInstance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.UnifiedRoleScheduleInstanceBase.GetFieldDeserializers()
    res["activatedUsing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleEligibilityScheduleInstance() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivatedUsing(val.(*UnifiedRoleEligibilityScheduleInstance))
        }
        return nil
    }
    res["assignmentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentType(val)
        }
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["memberType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberType(val)
        }
        return nil
    }
    res["roleAssignmentOriginId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleAssignmentOriginId(val)
        }
        return nil
    }
    res["roleAssignmentScheduleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleAssignmentScheduleId(val)
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    return res
}
func (m *UnifiedRoleAssignmentScheduleInstance) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UnifiedRoleAssignmentScheduleInstance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.UnifiedRoleScheduleInstanceBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("activatedUsing", m.GetActivatedUsing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assignmentType", m.GetAssignmentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("memberType", m.GetMemberType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleAssignmentOriginId", m.GetRoleAssignmentOriginId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleAssignmentScheduleId", m.GetRoleAssignmentScheduleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the activatedUsing property value. If the roleAssignmentScheduleInstance is activated by a roleEligibilityScheduleRequest, this is the link to the related schedule instance.
// Parameters:
//  - value : Value to set for the activatedUsing property.
func (m *UnifiedRoleAssignmentScheduleInstance) SetActivatedUsing(value *UnifiedRoleEligibilityScheduleInstance)() {
    m.activatedUsing = value
}
// Sets the assignmentType property value. Type of the assignment. It can either be Assigned or Activated.
// Parameters:
//  - value : Value to set for the assignmentType property.
func (m *UnifiedRoleAssignmentScheduleInstance) SetAssignmentType(value *string)() {
    m.assignmentType = value
}
// Sets the endDateTime property value. Time that the roleAssignmentInstance will expire
// Parameters:
//  - value : Value to set for the endDateTime property.
func (m *UnifiedRoleAssignmentScheduleInstance) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// Sets the memberType property value. Membership type of the assignment. It can either be Inherited, Direct, or Group.
// Parameters:
//  - value : Value to set for the memberType property.
func (m *UnifiedRoleAssignmentScheduleInstance) SetMemberType(value *string)() {
    m.memberType = value
}
// Sets the roleAssignmentOriginId property value. ID of the roleAssignment in the directory
// Parameters:
//  - value : Value to set for the roleAssignmentOriginId property.
func (m *UnifiedRoleAssignmentScheduleInstance) SetRoleAssignmentOriginId(value *string)() {
    m.roleAssignmentOriginId = value
}
// Sets the roleAssignmentScheduleId property value. ID of the parent roleAssignmentSchedule for this instance
// Parameters:
//  - value : Value to set for the roleAssignmentScheduleId property.
func (m *UnifiedRoleAssignmentScheduleInstance) SetRoleAssignmentScheduleId(value *string)() {
    m.roleAssignmentScheduleId = value
}
// Sets the startDateTime property value. Time that the roleAssignmentInstance will start
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *UnifiedRoleAssignmentScheduleInstance) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
