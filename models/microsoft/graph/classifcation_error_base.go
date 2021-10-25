package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ClassifcationErrorBase struct {
    additionalData map[string]interface{};
    code *string;
    innerError *ClassificationInnerError;
    message *string;
    target *string;
}
func NewClassifcationErrorBase()(*ClassifcationErrorBase) {
    m := &ClassifcationErrorBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ClassifcationErrorBase) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ClassifcationErrorBase) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
func (m *ClassifcationErrorBase) GetInnerError()(*ClassificationInnerError) {
    if m == nil {
        return nil
    } else {
        return m.innerError
    }
}
func (m *ClassifcationErrorBase) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
func (m *ClassifcationErrorBase) GetTarget()(*string) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *ClassifcationErrorBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCode(val)
        return nil
    }
    res["innerError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClassificationInnerError() })
        if err != nil {
            return err
        }
        m.SetInnerError(val.(*ClassificationInnerError))
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessage(val)
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTarget(val)
        return nil
    }
    return res
}
func (m *ClassifcationErrorBase) IsNil()(bool) {
    return m == nil
}
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
func (m *ClassifcationErrorBase) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ClassifcationErrorBase) SetCode(value *string)() {
    m.code = value
}
func (m *ClassifcationErrorBase) SetInnerError(value *ClassificationInnerError)() {
    m.innerError = value
}
func (m *ClassifcationErrorBase) SetMessage(value *string)() {
    m.message = value
}
func (m *ClassifcationErrorBase) SetTarget(value *string)() {
    m.target = value
}
