package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsInsight struct {
    additionalData map[string]interface{};
    insightId *string;
    severity *UserExperienceAnalyticsInsightSeverity;
    userExperienceAnalyticsMetricId *string;
    values []UserExperienceAnalyticsInsightValue;
}
func NewUserExperienceAnalyticsInsight()(*UserExperienceAnalyticsInsight) {
    m := &UserExperienceAnalyticsInsight{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserExperienceAnalyticsInsight) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserExperienceAnalyticsInsight) GetInsightId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.insightId
    }
}
func (m *UserExperienceAnalyticsInsight) GetSeverity()(*UserExperienceAnalyticsInsightSeverity) {
    if m == nil {
        return nil
    } else {
        return m.severity
    }
}
func (m *UserExperienceAnalyticsInsight) GetUserExperienceAnalyticsMetricId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsMetricId
    }
}
func (m *UserExperienceAnalyticsInsight) GetValues()([]UserExperienceAnalyticsInsightValue) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
func (m *UserExperienceAnalyticsInsight) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["insightId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInsightId(val)
        return nil
    }
    res["severity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsInsightSeverity)
        if err != nil {
            return err
        }
        cast := val.(UserExperienceAnalyticsInsightSeverity)
        m.SetSeverity(&cast)
        return nil
    }
    res["userExperienceAnalyticsMetricId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserExperienceAnalyticsMetricId(val)
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsInsightValue() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsInsightValue, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsInsightValue))
        }
        m.SetValues(res)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsInsight) IsNil()(bool) {
    return m == nil
}
func (m *UserExperienceAnalyticsInsight) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("insightId", m.GetInsightId())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := m.GetSeverity().String()
        err := writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userExperienceAnalyticsMetricId", m.GetUserExperienceAnalyticsMetricId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValues()))
        for i, v := range m.GetValues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("values", cast)
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
func (m *UserExperienceAnalyticsInsight) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserExperienceAnalyticsInsight) SetInsightId(value *string)() {
    m.insightId = value
}
func (m *UserExperienceAnalyticsInsight) SetSeverity(value *UserExperienceAnalyticsInsightSeverity)() {
    m.severity = value
}
func (m *UserExperienceAnalyticsInsight) SetUserExperienceAnalyticsMetricId(value *string)() {
    m.userExperienceAnalyticsMetricId = value
}
func (m *UserExperienceAnalyticsInsight) SetValues(value []UserExperienceAnalyticsInsightValue)() {
    m.values = value
}
