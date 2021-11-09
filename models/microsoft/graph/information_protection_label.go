package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new informationProtectionLabel and sets the default values.
func NewInformationProtectionLabel()(*InformationProtectionLabel) {
    m := &InformationProtectionLabel{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the color property value. The color that the UI should display for the label, if configured.
func (m *InformationProtectionLabel) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// Gets the description property value. The admin-defined description for the label.
func (m *InformationProtectionLabel) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in UI.
func (m *InformationProtectionLabel) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// Gets the name property value. The plaintext name of the label.
func (m *InformationProtectionLabel) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the parent property value. The parent label associated with a child label. Null if label has no parent.
func (m *InformationProtectionLabel) GetParent()(*ParentLabelDetails) {
    if m == nil {
        return nil
    } else {
        return m.parent
    }
}
// Gets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *InformationProtectionLabel) GetSensitivity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
// Gets the tooltip property value. The tooltip that should be displayed for the label in a UI.
func (m *InformationProtectionLabel) GetTooltip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tooltip
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the color property value. The color that the UI should display for the label, if configured.
// Parameters:
//  - value : Value to set for the color property.
func (m *InformationProtectionLabel) SetColor(value *string)() {
    m.color = value
}
// Sets the description property value. The admin-defined description for the label.
// Parameters:
//  - value : Value to set for the description property.
func (m *InformationProtectionLabel) SetDescription(value *string)() {
    m.description = value
}
// Sets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in UI.
// Parameters:
//  - value : Value to set for the isActive property.
func (m *InformationProtectionLabel) SetIsActive(value *bool)() {
    m.isActive = value
}
// Sets the name property value. The plaintext name of the label.
// Parameters:
//  - value : Value to set for the name property.
func (m *InformationProtectionLabel) SetName(value *string)() {
    m.name = value
}
// Sets the parent property value. The parent label associated with a child label. Null if label has no parent.
// Parameters:
//  - value : Value to set for the parent property.
func (m *InformationProtectionLabel) SetParent(value *ParentLabelDetails)() {
    m.parent = value
}
// Sets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
// Parameters:
//  - value : Value to set for the sensitivity property.
func (m *InformationProtectionLabel) SetSensitivity(value *int32)() {
    m.sensitivity = value
}
// Sets the tooltip property value. The tooltip that should be displayed for the label in a UI.
// Parameters:
//  - value : Value to set for the tooltip property.
func (m *InformationProtectionLabel) SetTooltip(value *string)() {
    m.tooltip = value
}
