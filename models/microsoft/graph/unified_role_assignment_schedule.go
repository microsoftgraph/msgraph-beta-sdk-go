package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UnifiedRoleAssignmentSchedule struct {
    UnifiedRoleScheduleBase
    // If the roleAssignmentSchedule is activated by a roleEligibilitySchedule, this is the link to that schedule.
    activatedUsing *UnifiedRoleEligibilitySchedule;
    // Type of the assignment. It can either be Assigned or Activated.
    assignmentType *string;
    // Membership type of the assignment. It can either be Inherited, Direct, or Group.
    memberType *string;
    // The schedule object of the role assignment request.
    scheduleInfo *RequestSchedule;
}
// Instantiates a new unifiedRoleAssignmentSchedule and sets the default values.
func NewUnifiedRoleAssignmentSchedule()(*UnifiedRoleAssignmentSchedule) {
    m := &UnifiedRoleAssignmentSchedule{
        UnifiedRoleScheduleBase: *NewUnifiedRoleScheduleBase(),
    }
    return m
}
// Gets the activatedUsing property value. If the roleAssignmentSchedule is activated by a roleEligibilitySchedule, this is the link to that schedule.
func (m *UnifiedRoleAssignmentSchedule) GetActivatedUsing()(*UnifiedRoleEligibilitySchedule) {
    if m == nil {
        return nil
    } else {
        return m.activatedUsing
    }
}
// Gets the assignmentType property value. Type of the assignment. It can either be Assigned or Activated.
func (m *UnifiedRoleAssignmentSchedule) GetAssignmentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentType
    }
}
// Gets the memberType property value. Membership type of the assignment. It can either be Inherited, Direct, or Group.
func (m *UnifiedRoleAssignmentSchedule) GetMemberType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberType
    }
}
// Gets the scheduleInfo property value. The schedule object of the role assignment request.
func (m *UnifiedRoleAssignmentSchedule) GetScheduleInfo()(*RequestSchedule) {
    if m == nil {
        return nil
    } else {
        return m.scheduleInfo
    }
}
// The deserialization information for the current model
func (m *UnifiedRoleAssignmentSchedule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.UnifiedRoleScheduleBase.GetFieldDeserializers()
    res["activatedUsing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleEligibilitySchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivatedUsing(val.(*UnifiedRoleEligibilitySchedule))
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
    res["scheduleInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRequestSchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleInfo(val.(*RequestSchedule))
        }
        return nil
    }
    return res
}
func (m *UnifiedRoleAssignmentSchedule) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UnifiedRoleAssignmentSchedule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.UnifiedRoleScheduleBase.Serialize(writer)
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
        err = writer.WriteStringValue("memberType", m.GetMemberType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scheduleInfo", m.GetScheduleInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the activatedUsing property value. If the roleAssignmentSchedule is activated by a roleEligibilitySchedule, this is the link to that schedule.
// Parameters:
//  - value : Value to set for the activatedUsing property.
func (m *UnifiedRoleAssignmentSchedule) SetActivatedUsing(value *UnifiedRoleEligibilitySchedule)() {
    m.activatedUsing = value
}
// Sets the assignmentType property value. Type of the assignment. It can either be Assigned or Activated.
// Parameters:
//  - value : Value to set for the assignmentType property.
func (m *UnifiedRoleAssignmentSchedule) SetAssignmentType(value *string)() {
    m.assignmentType = value
}
// Sets the memberType property value. Membership type of the assignment. It can either be Inherited, Direct, or Group.
// Parameters:
//  - value : Value to set for the memberType property.
func (m *UnifiedRoleAssignmentSchedule) SetMemberType(value *string)() {
    m.memberType = value
}
// Sets the scheduleInfo property value. The schedule object of the role assignment request.
// Parameters:
//  - value : Value to set for the scheduleInfo property.
func (m *UnifiedRoleAssignmentSchedule) SetScheduleInfo(value *RequestSchedule)() {
    m.scheduleInfo = value
}
