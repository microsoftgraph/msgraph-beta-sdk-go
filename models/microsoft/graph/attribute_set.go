package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttributeSet 
type AttributeSet struct {
    Entity
    // Description of the attribute set. Can be up to 128 characters long and include Unicode characters. Can be changed later.
    description *string;
    // Maximum number of custom security attributes that can be defined in this attribute set. Default value is null. If not specified, the administrator can add up to the maximum of 500 active attributes per tenant. Can be changed later.
    maxAttributesPerSet *int32;
}
// NewAttributeSet instantiates a new attributeSet and sets the default values.
func NewAttributeSet()(*AttributeSet) {
    m := &AttributeSet{
        Entity: *NewEntity(),
    }
    return m
}
// GetDescription gets the description property value. Description of the attribute set. Can be up to 128 characters long and include Unicode characters. Can be changed later.
func (m *AttributeSet) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetMaxAttributesPerSet gets the maxAttributesPerSet property value. Maximum number of custom security attributes that can be defined in this attribute set. Default value is null. If not specified, the administrator can add up to the maximum of 500 active attributes per tenant. Can be changed later.
func (m *AttributeSet) GetMaxAttributesPerSet()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxAttributesPerSet
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetDescription sets the description property value. Description of the attribute set. Can be up to 128 characters long and include Unicode characters. Can be changed later.
func (m *AttributeSet) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetMaxAttributesPerSet sets the maxAttributesPerSet property value. Maximum number of custom security attributes that can be defined in this attribute set. Default value is null. If not specified, the administrator can add up to the maximum of 500 active attributes per tenant. Can be changed later.
func (m *AttributeSet) SetMaxAttributesPerSet(value *int32)() {
    if m != nil {
        m.maxAttributesPerSet = value
    }
}
