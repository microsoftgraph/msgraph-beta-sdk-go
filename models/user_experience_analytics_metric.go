package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsMetric the user experience analytics metric contains the score and units of a metric of a user experience anlaytics category.
type UserExperienceAnalyticsMetric struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The unit of the user experience analytics metric.
    unit *string
    // The value of the user experience analytics metric.
    value *float64
}
// NewUserExperienceAnalyticsMetric instantiates a new userExperienceAnalyticsMetric and sets the default values.
func NewUserExperienceAnalyticsMetric()(*UserExperienceAnalyticsMetric) {
    m := &UserExperienceAnalyticsMetric{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserExperienceAnalyticsMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsMetricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsMetric(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsMetric) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsMetric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["unit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnit(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
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
// GetUnit gets the unit property value. The unit of the user experience analytics metric.
func (m *UserExperienceAnalyticsMetric) GetUnit()(*string) {
    if m == nil {
        return nil
    } else {
        return m.unit
    }
}
// GetValue gets the value property value. The value of the user experience analytics metric.
func (m *UserExperienceAnalyticsMetric) GetValue()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsMetric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("unit", m.GetUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsMetric) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUnit sets the unit property value. The unit of the user experience analytics metric.
func (m *UserExperienceAnalyticsMetric) SetUnit(value *string)() {
    if m != nil {
        m.unit = value
    }
}
// SetValue sets the value property value. The value of the user experience analytics metric.
func (m *UserExperienceAnalyticsMetric) SetValue(value *float64)() {
    if m != nil {
        m.value = value
    }
}
