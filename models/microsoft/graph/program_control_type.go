package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProgramControlType 
type ProgramControlType struct {
    Entity
    // 
    controlTypeGroupId *string;
    // The name of the program control type
    displayName *string;
}
// NewProgramControlType instantiates a new programControlType and sets the default values.
func NewProgramControlType()(*ProgramControlType) {
    m := &ProgramControlType{
        Entity: *NewEntity(),
    }
    return m
}
// GetControlTypeGroupId gets the controlTypeGroupId property value. 
func (m *ProgramControlType) GetControlTypeGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.controlTypeGroupId
    }
}
// GetDisplayName gets the displayName property value. The name of the program control type
func (m *ProgramControlType) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProgramControlType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["controlTypeGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControlTypeGroupId(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
func (m *ProgramControlType) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ProgramControlType) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("controlTypeGroupId", m.GetControlTypeGroupId())
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
    return nil
}
// SetControlTypeGroupId sets the controlTypeGroupId property value. 
func (m *ProgramControlType) SetControlTypeGroupId(value *string)() {
    if m != nil {
        m.controlTypeGroupId = value
    }
}
// SetDisplayName sets the displayName property value. The name of the program control type
func (m *ProgramControlType) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
