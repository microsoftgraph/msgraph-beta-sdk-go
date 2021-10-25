package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SuggestedEnrollmentLimit struct {
    additionalData map[string]interface{};
    suggestedDailyLimit *int32;
}
func NewSuggestedEnrollmentLimit()(*SuggestedEnrollmentLimit) {
    m := &SuggestedEnrollmentLimit{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SuggestedEnrollmentLimit) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SuggestedEnrollmentLimit) GetSuggestedDailyLimit()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.suggestedDailyLimit
    }
}
func (m *SuggestedEnrollmentLimit) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["suggestedDailyLimit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSuggestedDailyLimit(val)
        return nil
    }
    return res
}
func (m *SuggestedEnrollmentLimit) IsNil()(bool) {
    return m == nil
}
func (m *SuggestedEnrollmentLimit) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("suggestedDailyLimit", m.GetSuggestedDailyLimit())
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
func (m *SuggestedEnrollmentLimit) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SuggestedEnrollmentLimit) SetSuggestedDailyLimit(value *int32)() {
    m.suggestedDailyLimit = value
}
