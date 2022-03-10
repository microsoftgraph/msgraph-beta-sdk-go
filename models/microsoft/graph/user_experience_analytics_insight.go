package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsInsight provides operations to manage the deviceManagement singleton.
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
    values []UserExperienceAnalyticsInsightValueable;
}
// NewUserExperienceAnalyticsInsight instantiates a new userExperienceAnalyticsInsight and sets the default values.
func NewUserExperienceAnalyticsInsight()(*UserExperienceAnalyticsInsight) {
    m := &UserExperienceAnalyticsInsight{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserExperienceAnalyticsInsightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsInsightFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserExperienceAnalyticsInsight(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsInsight) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
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
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsInsightValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsInsightValueable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsInsightValueable)
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
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
func (m *UserExperienceAnalyticsInsight) GetValues()([]UserExperienceAnalyticsInsightValueable) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *UserExperienceAnalyticsInsight) SetValues(value []UserExperienceAnalyticsInsightValueable)() {
    if m != nil {
        m.values = value
    }
}
