package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AutoReviewSettings struct {
    additionalData map[string]interface{};
    notReviewedResult *string;
}
func NewAutoReviewSettings()(*AutoReviewSettings) {
    m := &AutoReviewSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AutoReviewSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AutoReviewSettings) GetNotReviewedResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notReviewedResult
    }
}
func (m *AutoReviewSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["notReviewedResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotReviewedResult(val)
        return nil
    }
    return res
}
func (m *AutoReviewSettings) IsNil()(bool) {
    return m == nil
}
func (m *AutoReviewSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("notReviewedResult", m.GetNotReviewedResult())
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
func (m *AutoReviewSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AutoReviewSettings) SetNotReviewedResult(value *string)() {
    m.notReviewedResult = value
}
