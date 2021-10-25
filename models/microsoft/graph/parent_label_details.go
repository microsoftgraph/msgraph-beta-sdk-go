package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ParentLabelDetails struct {
    additionalData map[string]interface{};
    color *string;
    description *string;
    id *string;
    isActive *bool;
    name *string;
    parent *ParentLabelDetails;
    sensitivity *int32;
    tooltip *string;
}
func NewParentLabelDetails()(*ParentLabelDetails) {
    m := &ParentLabelDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ParentLabelDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ParentLabelDetails) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
func (m *ParentLabelDetails) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *ParentLabelDetails) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *ParentLabelDetails) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
func (m *ParentLabelDetails) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *ParentLabelDetails) GetParent()(*ParentLabelDetails) {
    if m == nil {
        return nil
    } else {
        return m.parent
    }
}
func (m *ParentLabelDetails) GetSensitivity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
func (m *ParentLabelDetails) GetTooltip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tooltip
    }
}
func (m *ParentLabelDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
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
func (m *ParentLabelDetails) IsNil()(bool) {
    return m == nil
}
func (m *ParentLabelDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("sensitivity", m.GetSensitivity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tooltip", m.GetTooltip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ParentLabelDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ParentLabelDetails) SetColor(value *string)() {
    m.color = value
}
func (m *ParentLabelDetails) SetDescription(value *string)() {
    m.description = value
}
func (m *ParentLabelDetails) SetId(value *string)() {
    m.id = value
}
func (m *ParentLabelDetails) SetIsActive(value *bool)() {
    m.isActive = value
}
func (m *ParentLabelDetails) SetName(value *string)() {
    m.name = value
}
func (m *ParentLabelDetails) SetParent(value *ParentLabelDetails)() {
    m.parent = value
}
func (m *ParentLabelDetails) SetSensitivity(value *int32)() {
    m.sensitivity = value
}
func (m *ParentLabelDetails) SetTooltip(value *string)() {
    m.tooltip = value
}
