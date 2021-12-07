package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ClassifcationErrorBase 
type ClassifcationErrorBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    code *string;
    // 
    innerError *ClassificationInnerError;
    // 
    message *string;
    // 
    target *string;
}
// NewClassifcationErrorBase instantiates a new classifcationErrorBase and sets the default values.
func NewClassifcationErrorBase()(*ClassifcationErrorBase) {
    m := &ClassifcationErrorBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifcationErrorBase) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCode gets the code property value. 
func (m *ClassifcationErrorBase) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// GetInnerError gets the innerError property value. 
func (m *ClassifcationErrorBase) GetInnerError()(*ClassificationInnerError) {
    if m == nil {
        return nil
    } else {
        return m.innerError
    }
}
// GetMessage gets the message property value. 
func (m *ClassifcationErrorBase) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// GetTarget gets the target property value. 
func (m *ClassifcationErrorBase) GetTarget()(*string) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassifcationErrorBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["innerError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClassificationInnerError() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerError(val.(*ClassificationInnerError))
        }
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    return res
}
func (m *ClassifcationErrorBase) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ClassifcationErrorBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("innerError", m.GetInnerError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target", m.GetTarget())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifcationErrorBase) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCode sets the code property value. 
func (m *ClassifcationErrorBase) SetCode(value *string)() {
    if m != nil {
        m.code = value
    }
}
// SetInnerError sets the innerError property value. 
func (m *ClassifcationErrorBase) SetInnerError(value *ClassificationInnerError)() {
    if m != nil {
        m.innerError = value
    }
}
// SetMessage sets the message property value. 
func (m *ClassifcationErrorBase) SetMessage(value *string)() {
    if m != nil {
        m.message = value
    }
}
// SetTarget sets the target property value. 
func (m *ClassifcationErrorBase) SetTarget(value *string)() {
    if m != nil {
        m.target = value
    }
}
