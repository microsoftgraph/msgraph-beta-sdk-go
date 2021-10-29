package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ParentLabelDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The color that the user interface should display for the label, if configured.
    color *string;
    // The admin-defined description for the label.
    description *string;
    // The label ID is a globally unique identifier (GUID).
    id *string;
    // Indicates whether the label is active or not. Active labels should be hidden or disabled in user interfaces.
    isActive *bool;
    // The plaintext name of the label.
    name *string;
    // 
    parent *ParentLabelDetails;
    // The sensitivity value of the label, where lower is less sensitive.
    sensitivity *int32;
    // The tooltip that should be displayed for the label in a user interface.
    tooltip *string;
}
// Instantiates a new parentLabelDetails and sets the default values.
func NewParentLabelDetails()(*ParentLabelDetails) {
    m := &ParentLabelDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParentLabelDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the color property value. The color that the user interface should display for the label, if configured.
func (m *ParentLabelDetails) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// Gets the description property value. The admin-defined description for the label.
func (m *ParentLabelDetails) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the id property value. The label ID is a globally unique identifier (GUID).
func (m *ParentLabelDetails) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in user interfaces.
func (m *ParentLabelDetails) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// Gets the name property value. The plaintext name of the label.
func (m *ParentLabelDetails) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the parent property value. 
func (m *ParentLabelDetails) GetParent()(*ParentLabelDetails) {
    if m == nil {
        return nil
    } else {
        return m.parent
    }
}
// Gets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *ParentLabelDetails) GetSensitivity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
// Gets the tooltip property value. The tooltip that should be displayed for the label in a user interface.
func (m *ParentLabelDetails) GetTooltip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tooltip
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ParentLabelDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the color property value. The color that the user interface should display for the label, if configured.
// Parameters:
//  - value : Value to set for the color property.
func (m *ParentLabelDetails) SetColor(value *string)() {
    m.color = value
}
// Sets the description property value. The admin-defined description for the label.
// Parameters:
//  - value : Value to set for the description property.
func (m *ParentLabelDetails) SetDescription(value *string)() {
    m.description = value
}
// Sets the id property value. The label ID is a globally unique identifier (GUID).
// Parameters:
//  - value : Value to set for the id property.
func (m *ParentLabelDetails) SetId(value *string)() {
    m.id = value
}
// Sets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in user interfaces.
// Parameters:
//  - value : Value to set for the isActive property.
func (m *ParentLabelDetails) SetIsActive(value *bool)() {
    m.isActive = value
}
// Sets the name property value. The plaintext name of the label.
// Parameters:
//  - value : Value to set for the name property.
func (m *ParentLabelDetails) SetName(value *string)() {
    m.name = value
}
// Sets the parent property value. 
// Parameters:
//  - value : Value to set for the parent property.
func (m *ParentLabelDetails) SetParent(value *ParentLabelDetails)() {
    m.parent = value
}
// Sets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
// Parameters:
//  - value : Value to set for the sensitivity property.
func (m *ParentLabelDetails) SetSensitivity(value *int32)() {
    m.sensitivity = value
}
// Sets the tooltip property value. The tooltip that should be displayed for the label in a user interface.
// Parameters:
//  - value : Value to set for the tooltip property.
func (m *ParentLabelDetails) SetTooltip(value *string)() {
    m.tooltip = value
}
