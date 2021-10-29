package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ReferencedObject struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the referenced object. Must match one of the objects in the directory definition.
    referencedObjectName *string;
    // Currently not supported. Name of the property in the referenced object, the value for which is used as the reference.
    referencedProperty *string;
}
// Instantiates a new referencedObject and sets the default values.
func NewReferencedObject()(*ReferencedObject) {
    m := &ReferencedObject{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReferencedObject) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the referencedObjectName property value. Name of the referenced object. Must match one of the objects in the directory definition.
func (m *ReferencedObject) GetReferencedObjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referencedObjectName
    }
}
// Gets the referencedProperty property value. Currently not supported. Name of the property in the referenced object, the value for which is used as the reference.
func (m *ReferencedObject) GetReferencedProperty()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referencedProperty
    }
}
// The deserialization information for the current model
func (m *ReferencedObject) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["referencedObjectName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReferencedObjectName(val)
        return nil
    }
    res["referencedProperty"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReferencedProperty(val)
        return nil
    }
    return res
}
func (m *ReferencedObject) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ReferencedObject) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("referencedObjectName", m.GetReferencedObjectName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("referencedProperty", m.GetReferencedProperty())
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
func (m *ReferencedObject) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the referencedObjectName property value. Name of the referenced object. Must match one of the objects in the directory definition.
// Parameters:
//  - value : Value to set for the referencedObjectName property.
func (m *ReferencedObject) SetReferencedObjectName(value *string)() {
    m.referencedObjectName = value
}
// Sets the referencedProperty property value. Currently not supported. Name of the property in the referenced object, the value for which is used as the reference.
// Parameters:
//  - value : Value to set for the referencedProperty property.
func (m *ReferencedObject) SetReferencedProperty(value *string)() {
    m.referencedProperty = value
}
