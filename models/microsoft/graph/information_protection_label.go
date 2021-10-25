package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InformationProtectionLabel struct {
    Entity
    color *string;
    description *string;
    isActive *bool;
    name *string;
    parent *ParentLabelDetails;
    sensitivity *int32;
    tooltip *string;
}
func NewInformationProtectionLabel()(*InformationProtectionLabel) {
    m := &InformationProtectionLabel{
        Entity: *NewEntity(),
    }
    return m
}
func (m *InformationProtectionLabel) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
func (m *InformationProtectionLabel) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *InformationProtectionLabel) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
func (m *InformationProtectionLabel) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *InformationProtectionLabel) GetParent()(*ParentLabelDetails) {
    if m == nil {
        return nil
    } else {
        return m.parent
    }
}
func (m *InformationProtectionLabel) GetSensitivity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
func (m *InformationProtectionLabel) GetTooltip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tooltip
    }
}
func (m *InformationProtectionLabel) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetColor(val)
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
    res["isActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsActive(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["parent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewParentLabelDetails() })
        if err != nil {
            return err
        }
        m.SetParent(val.(*ParentLabelDetails))
        return nil
    }
    res["sensitivity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSensitivity(val)
        return nil
    }
    res["tooltip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTooltip(val)
        return nil
    }
    return res
}
func (m *InformationProtectionLabel) IsNil()(bool) {
    return m == nil
}
func (m *InformationProtectionLabel) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
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
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("sensitivity", m.GetSensitivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tooltip", m.GetTooltip())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *InformationProtectionLabel) SetColor(value *string)() {
    m.color = value
}
func (m *InformationProtectionLabel) SetDescription(value *string)() {
    m.description = value
}
func (m *InformationProtectionLabel) SetIsActive(value *bool)() {
    m.isActive = value
}
func (m *InformationProtectionLabel) SetName(value *string)() {
    m.name = value
}
func (m *InformationProtectionLabel) SetParent(value *ParentLabelDetails)() {
    m.parent = value
}
func (m *InformationProtectionLabel) SetSensitivity(value *int32)() {
    m.sensitivity = value
}
func (m *InformationProtectionLabel) SetTooltip(value *string)() {
    m.tooltip = value
}
