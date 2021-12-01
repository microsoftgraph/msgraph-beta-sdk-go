package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CustomQuestionAnswer 
type CustomQuestionAnswer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Display name of the custom registration question. Read-only.
    displayName *string;
    // ID the custom registration question. Read-only.
    questionId *string;
    // Answer to the custom registration question.
    value *string;
}
// NewCustomQuestionAnswer instantiates a new customQuestionAnswer and sets the default values.
func NewCustomQuestionAnswer()(*CustomQuestionAnswer) {
    m := &CustomQuestionAnswer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomQuestionAnswer) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Display name of the custom registration question. Read-only.
func (m *CustomQuestionAnswer) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetQuestionId gets the questionId property value. ID the custom registration question. Read-only.
func (m *CustomQuestionAnswer) GetQuestionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.questionId
    }
}
// GetValue gets the value property value. Answer to the custom registration question.
func (m *CustomQuestionAnswer) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomQuestionAnswer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["questionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuestionId(val)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
func (m *CustomQuestionAnswer) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CustomQuestionAnswer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("questionId", m.GetQuestionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomQuestionAnswer) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the custom registration question. Read-only.
func (m *CustomQuestionAnswer) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetQuestionId sets the questionId property value. ID the custom registration question. Read-only.
func (m *CustomQuestionAnswer) SetQuestionId(value *string)() {
    if m != nil {
        m.questionId = value
    }
}
// SetValue sets the value property value. Answer to the custom registration question.
func (m *CustomQuestionAnswer) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
