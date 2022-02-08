package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsInsight 
type UserExperienceAnalyticsInsight struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The unique identifier of the user experience analytics insight.
    insightId *string;
    // The value of the user experience analytics insight. Possible values are: none, informational, warning, error.
    severity *UserExperienceAnalyticsInsightSeverity;
    // The unique identifier of the user experience analytics insight.
    userExperienceAnalyticsMetricId *string;
    // The value of the user experience analytics insight.
    values []UserExperienceAnalyticsInsightValue;
}
// NewUserExperienceAnalyticsInsight instantiates a new userExperienceAnalyticsInsight and sets the default values.
func NewUserExperienceAnalyticsInsight()(*UserExperienceAnalyticsInsight) {
    m := &UserExperienceAnalyticsInsight{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsInsight) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetInsightId gets the insightId property value. The unique identifier of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) GetInsightId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.insightId
    }
}
// GetSeverity gets the severity property value. The value of the user experience analytics insight. Possible values are: none, informational, warning, error.
func (m *UserExperienceAnalyticsInsight) GetSeverity()(*UserExperienceAnalyticsInsightSeverity) {
    if m == nil {
        return nil
    } else {
        return m.severity
    }
}
// GetUserExperienceAnalyticsMetricId gets the userExperienceAnalyticsMetricId property value. The unique identifier of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) GetUserExperienceAnalyticsMetricId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsMetricId
    }
}
// GetValues gets the values property value. The value of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) GetValues()([]UserExperienceAnalyticsInsightValue) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsInsight) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["insightId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInsightId(val)
        }
        return nil
    }
    res["severity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsInsightSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*UserExperienceAnalyticsInsightSeverity))
        }
        return nil
    }
    res["userExperienceAnalyticsMetricId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsMetricId(val)
        }
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsInsightValue() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsInsightValue, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserExperienceAnalyticsInsightValue))
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsInsight) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsInsight) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("insightId", m.GetInsightId())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
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
    if m.GetValues() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsInsight) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetInsightId sets the insightId property value. The unique identifier of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) SetInsightId(value *string)() {
    if m != nil {
        m.insightId = value
    }
}
// SetSeverity sets the severity property value. The value of the user experience analytics insight. Possible values are: none, informational, warning, error.
func (m *UserExperienceAnalyticsInsight) SetSeverity(value *UserExperienceAnalyticsInsightSeverity)() {
    if m != nil {
        m.severity = value
    }
}
// SetUserExperienceAnalyticsMetricId sets the userExperienceAnalyticsMetricId property value. The unique identifier of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) SetUserExperienceAnalyticsMetricId(value *string)() {
    if m != nil {
        m.userExperienceAnalyticsMetricId = value
    }
}
// SetValues sets the values property value. The value of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) SetValues(value []UserExperienceAnalyticsInsightValue)() {
    if m != nil {
        m.values = value
    }
}
