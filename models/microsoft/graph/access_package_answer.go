package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessPackageAnswer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The question the answer is for. Required and Read-only.
    answeredQuestion *AccessPackageQuestion;
    // The display value of the answer. Required.
    displayValue *string;
}
// Instantiates a new accessPackageAnswer and sets the default values.
func NewAccessPackageAnswer()(*AccessPackageAnswer) {
    m := &AccessPackageAnswer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAnswer) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the answeredQuestion property value. The question the answer is for. Required and Read-only.
func (m *AccessPackageAnswer) GetAnsweredQuestion()(*AccessPackageQuestion) {
    if m == nil {
        return nil
    } else {
        return m.answeredQuestion
    }
}
// Gets the displayValue property value. The display value of the answer. Required.
func (m *AccessPackageAnswer) GetDisplayValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayValue
    }
}
// The deserialization information for the current model
func (m *AccessPackageAnswer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["answeredQuestion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageQuestion() })
        if err != nil {
            return err
        }
        m.SetAnsweredQuestion(val.(*AccessPackageQuestion))
        return nil
    }
    res["displayValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayValue(val)
        return nil
    }
    return res
}
func (m *AccessPackageAnswer) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessPackageAnswer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("answeredQuestion", m.GetAnsweredQuestion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayValue", m.GetDisplayValue())
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
func (m *AccessPackageAnswer) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the answeredQuestion property value. The question the answer is for. Required and Read-only.
// Parameters:
//  - value : Value to set for the answeredQuestion property.
func (m *AccessPackageAnswer) SetAnsweredQuestion(value *AccessPackageQuestion)() {
    m.answeredQuestion = value
}
// Sets the displayValue property value. The display value of the answer. Required.
// Parameters:
//  - value : Value to set for the displayValue property.
func (m *AccessPackageAnswer) SetDisplayValue(value *string)() {
    m.displayValue = value
}
