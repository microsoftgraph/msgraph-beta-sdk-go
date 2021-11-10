package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessPackageQuestion struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // ID of the question.
    id *string;
    // 
    isAnswerEditable *bool;
    // Whether the requestor is required to supply an answer or not.
    isRequired *bool;
    // Relative position of this question when displaying a list of questions to the requestor.
    sequence *int32;
    // The text of the question to show to the requestor.
    text *AccessPackageLocalizedContent;
}
// Instantiates a new accessPackageQuestion and sets the default values.
func NewAccessPackageQuestion()(*AccessPackageQuestion) {
    m := &AccessPackageQuestion{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageQuestion) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the id property value. ID of the question.
func (m *AccessPackageQuestion) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the isAnswerEditable property value. 
func (m *AccessPackageQuestion) GetIsAnswerEditable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAnswerEditable
    }
}
// Gets the isRequired property value. Whether the requestor is required to supply an answer or not.
func (m *AccessPackageQuestion) GetIsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequired
    }
}
// Gets the sequence property value. Relative position of this question when displaying a list of questions to the requestor.
func (m *AccessPackageQuestion) GetSequence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sequence
    }
}
// Gets the text property value. The text of the question to show to the requestor.
func (m *AccessPackageQuestion) GetText()(*AccessPackageLocalizedContent) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// The deserialization information for the current model
func (m *AccessPackageQuestion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["isAnswerEditable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAnswerEditable(val)
        }
        return nil
    }
    res["isRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequired(val)
        }
        return nil
    }
    res["sequence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSequence(val)
        }
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageLocalizedContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val.(*AccessPackageLocalizedContent))
        }
        return nil
    }
    return res
}
func (m *AccessPackageQuestion) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessPackageQuestion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAnswerEditable", m.GetIsAnswerEditable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("sequence", m.GetSequence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("text", m.GetText())
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
func (m *AccessPackageQuestion) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the id property value. ID of the question.
// Parameters:
//  - value : Value to set for the id property.
func (m *AccessPackageQuestion) SetId(value *string)() {
    m.id = value
}
// Sets the isAnswerEditable property value. 
// Parameters:
//  - value : Value to set for the isAnswerEditable property.
func (m *AccessPackageQuestion) SetIsAnswerEditable(value *bool)() {
    m.isAnswerEditable = value
}
// Sets the isRequired property value. Whether the requestor is required to supply an answer or not.
// Parameters:
//  - value : Value to set for the isRequired property.
func (m *AccessPackageQuestion) SetIsRequired(value *bool)() {
    m.isRequired = value
}
// Sets the sequence property value. Relative position of this question when displaying a list of questions to the requestor.
// Parameters:
//  - value : Value to set for the sequence property.
func (m *AccessPackageQuestion) SetSequence(value *int32)() {
    m.sequence = value
}
// Sets the text property value. The text of the question to show to the requestor.
// Parameters:
//  - value : Value to set for the text property.
func (m *AccessPackageQuestion) SetText(value *AccessPackageLocalizedContent)() {
    m.text = value
}
