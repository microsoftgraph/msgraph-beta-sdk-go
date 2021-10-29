package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UnifiedRoleEligibilitySchedule struct {
    UnifiedRoleScheduleBase
    // Membership type of the eligible assignment. It can either be Inherited, Direct, or Group.
    memberType *string;
    // The schedule object of the eligible role assignment request.
    scheduleInfo *RequestSchedule;
}
// Instantiates a new unifiedRoleEligibilitySchedule and sets the default values.
func NewUnifiedRoleEligibilitySchedule()(*UnifiedRoleEligibilitySchedule) {
    m := &UnifiedRoleEligibilitySchedule{
        UnifiedRoleScheduleBase: *NewUnifiedRoleScheduleBase(),
    }
    return m
}
// Gets the memberType property value. Membership type of the eligible assignment. It can either be Inherited, Direct, or Group.
func (m *UnifiedRoleEligibilitySchedule) GetMemberType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberType
    }
}
// Gets the scheduleInfo property value. The schedule object of the eligible role assignment request.
func (m *UnifiedRoleEligibilitySchedule) GetScheduleInfo()(*RequestSchedule) {
    if m == nil {
        return nil
    } else {
        return m.scheduleInfo
    }
}
// The deserialization information for the current model
func (m *UnifiedRoleEligibilitySchedule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.UnifiedRoleScheduleBase.GetFieldDeserializers()
    res["memberType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMemberType(val)
        return nil
    }
    res["scheduleInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRequestSchedule() })
        if err != nil {
            return err
        }
        m.SetScheduleInfo(val.(*RequestSchedule))
        return nil
    }
    return res
}
func (m *UnifiedRoleEligibilitySchedule) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UnifiedRoleEligibilitySchedule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.UnifiedRoleScheduleBase.Serialize(writer)
    if err != nil {
        return err
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
// Sets the memberType property value. Membership type of the eligible assignment. It can either be Inherited, Direct, or Group.
// Parameters:
//  - value : Value to set for the memberType property.
func (m *UnifiedRoleEligibilitySchedule) SetMemberType(value *string)() {
    m.memberType = value
}
// Sets the scheduleInfo property value. The schedule object of the eligible role assignment request.
// Parameters:
//  - value : Value to set for the scheduleInfo property.
func (m *UnifiedRoleEligibilitySchedule) SetScheduleInfo(value *RequestSchedule)() {
    m.scheduleInfo = value
}
