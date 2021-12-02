package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Shift 
type Shift struct {
    ChangeTrackedEntity
    // The draft version of this shift that is viewable by managers. Required.
    draftShift *ShiftItem;
    // 
    isStagedForDeletion *bool;
    // ID of the scheduling group the shift is part of. Required.
    schedulingGroupId *string;
    // The shared version of this shift that is viewable by both employees and managers. Required.
    sharedShift *ShiftItem;
    // ID of the user assigned to the shift. Required.
    userId *string;
}
// NewShift instantiates a new shift and sets the default values.
func NewShift()(*Shift) {
    m := &Shift{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// GetDraftShift gets the draftShift property value. The draft version of this shift that is viewable by managers. Required.
func (m *Shift) GetDraftShift()(*ShiftItem) {
    if m == nil {
        return nil
    } else {
        return m.draftShift
    }
}
// GetIsStagedForDeletion gets the isStagedForDeletion property value. 
func (m *Shift) GetIsStagedForDeletion()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isStagedForDeletion
    }
}
// GetSchedulingGroupId gets the schedulingGroupId property value. ID of the scheduling group the shift is part of. Required.
func (m *Shift) GetSchedulingGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schedulingGroupId
    }
}
// GetSharedShift gets the sharedShift property value. The shared version of this shift that is viewable by both employees and managers. Required.
func (m *Shift) GetSharedShift()(*ShiftItem) {
    if m == nil {
        return nil
    } else {
        return m.sharedShift
    }
}
// GetUserId gets the userId property value. ID of the user assigned to the shift. Required.
func (m *Shift) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Shift) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftShift"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShiftItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraftShift(val.(*ShiftItem))
        }
        return nil
    }
    res["isStagedForDeletion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsStagedForDeletion(val)
        }
        return nil
    }
    res["schedulingGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedulingGroupId(val)
        }
        return nil
    }
    res["sharedShift"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShiftItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedShift(val.(*ShiftItem))
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
func (m *Shift) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Shift) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("draftShift", m.GetDraftShift())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isStagedForDeletion", m.GetIsStagedForDeletion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("schedulingGroupId", m.GetSchedulingGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedShift", m.GetSharedShift())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDraftShift sets the draftShift property value. The draft version of this shift that is viewable by managers. Required.
func (m *Shift) SetDraftShift(value *ShiftItem)() {
    if m != nil {
        m.draftShift = value
    }
}
// SetIsStagedForDeletion sets the isStagedForDeletion property value. 
func (m *Shift) SetIsStagedForDeletion(value *bool)() {
    if m != nil {
        m.isStagedForDeletion = value
    }
}
// SetSchedulingGroupId sets the schedulingGroupId property value. ID of the scheduling group the shift is part of. Required.
func (m *Shift) SetSchedulingGroupId(value *string)() {
    if m != nil {
        m.schedulingGroupId = value
    }
}
// SetSharedShift sets the sharedShift property value. The shared version of this shift that is viewable by both employees and managers. Required.
func (m *Shift) SetSharedShift(value *ShiftItem)() {
    if m != nil {
        m.sharedShift = value
    }
}
// SetUserId sets the userId property value. ID of the user assigned to the shift. Required.
func (m *Shift) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
