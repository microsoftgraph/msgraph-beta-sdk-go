package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TermColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies whether the column will allow more than one value.
    allowMultipleValues *bool;
    // Specifies whether to display the entire term path or only the term label.
    showFullyQualifiedName *bool;
}
// Instantiates a new termColumn and sets the default values.
func NewTermColumn()(*TermColumn) {
    m := &TermColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TermColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowMultipleValues property value. Specifies whether the column will allow more than one value.
func (m *TermColumn) GetAllowMultipleValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowMultipleValues
    }
}
// Gets the showFullyQualifiedName property value. Specifies whether to display the entire term path or only the term label.
func (m *TermColumn) GetShowFullyQualifiedName()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showFullyQualifiedName
    }
}
// The deserialization information for the current model
func (m *TermColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowMultipleValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowMultipleValues(val)
        return nil
    }
    res["showFullyQualifiedName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowFullyQualifiedName(val)
        return nil
    }
    return res
}
func (m *TermColumn) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TermColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleValues", m.GetAllowMultipleValues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showFullyQualifiedName", m.GetShowFullyQualifiedName())
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
func (m *TermColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowMultipleValues property value. Specifies whether the column will allow more than one value.
// Parameters:
//  - value : Value to set for the allowMultipleValues property.
func (m *TermColumn) SetAllowMultipleValues(value *bool)() {
    m.allowMultipleValues = value
}
// Sets the showFullyQualifiedName property value. Specifies whether to display the entire term path or only the term label.
// Parameters:
//  - value : Value to set for the showFullyQualifiedName property.
func (m *TermColumn) SetShowFullyQualifiedName(value *bool)() {
    m.showFullyQualifiedName = value
}
