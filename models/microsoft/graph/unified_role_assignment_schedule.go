package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleAssignmentSchedule 
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
// NewUnifiedRoleAssignmentSchedule instantiates a new unifiedRoleAssignmentSchedule and sets the default values.
func NewUnifiedRoleAssignmentSchedule()(*UnifiedRoleAssignmentSchedule) {
    m := &UnifiedRoleAssignmentSchedule{
        UnifiedRoleScheduleBase: *NewUnifiedRoleScheduleBase(),
    }
    return m
}
// GetActivatedUsing gets the activatedUsing property value. If the roleAssignmentSchedule is activated by a roleEligibilitySchedule, this is the link to that schedule.
func (m *UnifiedRoleAssignmentSchedule) GetActivatedUsing()(*UnifiedRoleEligibilitySchedule) {
    if m == nil {
        return nil
    } else {
        return m.activatedUsing
    }
}
// GetAssignmentType gets the assignmentType property value. Type of the assignment. It can either be Assigned or Activated.
func (m *UnifiedRoleAssignmentSchedule) GetAssignmentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentType
    }
}
// GetMemberType gets the memberType property value. Membership type of the assignment. It can either be Inherited, Direct, or Group.
func (m *UnifiedRoleAssignmentSchedule) GetMemberType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberType
    }
}
// GetScheduleInfo gets the scheduleInfo property value. The schedule object of the role assignment request.
func (m *UnifiedRoleAssignmentSchedule) GetScheduleInfo()(*RequestSchedule) {
    if m == nil {
        return nil
    } else {
        return m.scheduleInfo
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetActivatedUsing sets the activatedUsing property value. If the roleAssignmentSchedule is activated by a roleEligibilitySchedule, this is the link to that schedule.
func (m *UnifiedRoleAssignmentSchedule) SetActivatedUsing(value *UnifiedRoleEligibilitySchedule)() {
    if m != nil {
        m.activatedUsing = value
    }
}
// SetAssignmentType sets the assignmentType property value. Type of the assignment. It can either be Assigned or Activated.
func (m *UnifiedRoleAssignmentSchedule) SetAssignmentType(value *string)() {
    if m != nil {
        m.assignmentType = value
    }
}
// SetMemberType sets the memberType property value. Membership type of the assignment. It can either be Inherited, Direct, or Group.
func (m *UnifiedRoleAssignmentSchedule) SetMemberType(value *string)() {
    if m != nil {
        m.memberType = value
    }
}
// SetScheduleInfo sets the scheduleInfo property value. The schedule object of the role assignment request.
func (m *UnifiedRoleAssignmentSchedule) SetScheduleInfo(value *RequestSchedule)() {
    if m != nil {
        m.scheduleInfo = value
    }
}
