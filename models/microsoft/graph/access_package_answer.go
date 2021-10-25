package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessPackageAnswer struct {
    additionalData map[string]interface{};
    answeredQuestion *AccessPackageQuestion;
    displayValue *string;
}
func NewAccessPackageAnswer()(*AccessPackageAnswer) {
    m := &AccessPackageAnswer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AccessPackageAnswer) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AccessPackageAnswer) GetAnsweredQuestion()(*AccessPackageQuestion) {
    if m == nil {
        return nil
    } else {
        return m.answeredQuestion
    }
}
func (m *AccessPackageAnswer) GetDisplayValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayValue
    }
}
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
func (m *AccessPackageAnswer) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AccessPackageAnswer) SetAnsweredQuestion(value *AccessPackageQuestion)() {
    m.answeredQuestion = value
}
func (m *AccessPackageAnswer) SetDisplayValue(value *string)() {
    m.displayValue = value
}
