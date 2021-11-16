package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AttributeSet struct {
    Entity
    // 
    description *string;
    // 
    maxAttributesPerSet *int32;
}
// Instantiates a new attributeSet and sets the default values.
func NewAttributeSet()(*AttributeSet) {
    m := &AttributeSet{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the description property value. 
func (m *AttributeSet) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the maxAttributesPerSet property value. 
func (m *AttributeSet) GetMaxAttributesPerSet()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxAttributesPerSet
    }
}
// The deserialization information for the current model
func (m *AttributeSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["maxAttributesPerSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxAttributesPerSet(val)
        }
        return nil
    }
    return res
}
func (m *AttributeSet) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AttributeSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maxAttributesPerSet", m.GetMaxAttributesPerSet())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the description property value. 
// Parameters:
//  - value : Value to set for the description property.
func (m *AttributeSet) SetDescription(value *string)() {
    m.description = value
}
// Sets the maxAttributesPerSet property value. 
// Parameters:
//  - value : Value to set for the maxAttributesPerSet property.
func (m *AttributeSet) SetMaxAttributesPerSet(value *int32)() {
    m.maxAttributesPerSet = value
}
