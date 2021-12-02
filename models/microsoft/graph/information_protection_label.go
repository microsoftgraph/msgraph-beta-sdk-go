package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InformationProtectionLabel 
type InformationProtectionLabel struct {
    Entity
    // The color that the UI should display for the label, if configured.
    color *string;
    // The admin-defined description for the label.
    description *string;
    // Indicates whether the label is active or not. Active labels should be hidden or disabled in UI.
    isActive *bool;
    // The plaintext name of the label.
    name *string;
    // The parent label associated with a child label. Null if label has no parent.
    parent *ParentLabelDetails;
    // The sensitivity value of the label, where lower is less sensitive.
    sensitivity *int32;
    // The tooltip that should be displayed for the label in a UI.
    tooltip *string;
}
// NewInformationProtectionLabel instantiates a new informationProtectionLabel and sets the default values.
func NewInformationProtectionLabel()(*InformationProtectionLabel) {
    m := &InformationProtectionLabel{
        Entity: *NewEntity(),
    }
    return m
}
// GetColor gets the color property value. The color that the UI should display for the label, if configured.
func (m *InformationProtectionLabel) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// GetDescription gets the description property value. The admin-defined description for the label.
func (m *InformationProtectionLabel) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetIsActive gets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in UI.
func (m *InformationProtectionLabel) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// GetName gets the name property value. The plaintext name of the label.
func (m *InformationProtectionLabel) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetParent gets the parent property value. The parent label associated with a child label. Null if label has no parent.
func (m *InformationProtectionLabel) GetParent()(*ParentLabelDetails) {
    if m == nil {
        return nil
    } else {
        return m.parent
    }
}
// GetSensitivity gets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *InformationProtectionLabel) GetSensitivity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
// GetTooltip gets the tooltip property value. The tooltip that should be displayed for the label in a UI.
func (m *InformationProtectionLabel) GetTooltip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tooltip
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionLabel) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["isActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["parent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewParentLabelDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(*ParentLabelDetails))
        }
        return nil
    }
    res["sensitivity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivity(val)
        }
        return nil
    }
    res["tooltip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTooltip(val)
        }
        return nil
    }
    return res
}
func (m *InformationProtectionLabel) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetColor sets the color property value. The color that the UI should display for the label, if configured.
func (m *InformationProtectionLabel) SetColor(value *string)() {
    if m != nil {
        m.color = value
    }
}
// SetDescription sets the description property value. The admin-defined description for the label.
func (m *InformationProtectionLabel) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetIsActive sets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in UI.
func (m *InformationProtectionLabel) SetIsActive(value *bool)() {
    if m != nil {
        m.isActive = value
    }
}
// SetName sets the name property value. The plaintext name of the label.
func (m *InformationProtectionLabel) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetParent sets the parent property value. The parent label associated with a child label. Null if label has no parent.
func (m *InformationProtectionLabel) SetParent(value *ParentLabelDetails)() {
    if m != nil {
        m.parent = value
    }
}
// SetSensitivity sets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *InformationProtectionLabel) SetSensitivity(value *int32)() {
    if m != nil {
        m.sensitivity = value
    }
}
// SetTooltip sets the tooltip property value. The tooltip that should be displayed for the label in a UI.
func (m *InformationProtectionLabel) SetTooltip(value *string)() {
    if m != nil {
        m.tooltip = value
    }
}
