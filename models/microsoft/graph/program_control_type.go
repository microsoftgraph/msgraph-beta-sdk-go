package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ProgramControlType struct {
    Entity
    controlTypeGroupId *string;
    displayName *string;
}
func NewProgramControlType()(*ProgramControlType) {
    m := &ProgramControlType{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ProgramControlType) GetControlTypeGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.controlTypeGroupId
    }
}
func (m *ProgramControlType) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *ProgramControlType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["controlTypeGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetControlTypeGroupId(val)
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
    return res
}
func (m *ProgramControlType) IsNil()(bool) {
    return m == nil
}
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
func (m *ProgramControlType) SetControlTypeGroupId(value *string)() {
    m.controlTypeGroupId = value
}
func (m *ProgramControlType) SetDisplayName(value *string)() {
    m.displayName = value
}
