package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsInsight the user experience analytics insight is the recomendation to improve the user experience analytics score.
type UserExperienceAnalyticsInsight struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The unique identifier of the user experience analytics insight.
    insightId *string
    // The OdataType property
    odataType *string
    // The severity property
    severity *UserExperienceAnalyticsInsightSeverity
    // The unique identifier of the user experience analytics insight.
    userExperienceAnalyticsMetricId *string
    // The value of the user experience analytics insight.
    values []UserExperienceAnalyticsInsightValueable
}
// NewUserExperienceAnalyticsInsight instantiates a new userExperienceAnalyticsInsight and sets the default values.
func NewUserExperienceAnalyticsInsight()(*UserExperienceAnalyticsInsight) {
    m := &UserExperienceAnalyticsInsight{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.userExperienceAnalyticsInsight";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserExperienceAnalyticsInsightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsInsightFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsInsight(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsInsight) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsInsight) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["insightId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInsightId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["severity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseUserExperienceAnalyticsInsightSeverity , m.SetSeverity)
    res["userExperienceAnalyticsMetricId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserExperienceAnalyticsMetricId)
    res["values"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsInsightValueFromDiscriminatorValue , m.SetValues)
    return res
}
// GetInsightId gets the insightId property value. The unique identifier of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) GetInsightId()(*string) {
    return m.insightId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsInsight) GetOdataType()(*string) {
    return m.odataType
}
// GetSeverity gets the severity property value. The severity property
func (m *UserExperienceAnalyticsInsight) GetSeverity()(*UserExperienceAnalyticsInsightSeverity) {
    return m.severity
}
// GetUserExperienceAnalyticsMetricId gets the userExperienceAnalyticsMetricId property value. The unique identifier of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) GetUserExperienceAnalyticsMetricId()(*string) {
    return m.userExperienceAnalyticsMetricId
}
// GetValues gets the values property value. The value of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) GetValues()([]UserExperienceAnalyticsInsightValueable) {
    return m.values
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsInsight) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("insightId", m.GetInsightId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValues())
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
    m.additionalData = value
}
// SetInsightId sets the insightId property value. The unique identifier of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) SetInsightId(value *string)() {
    m.insightId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsInsight) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSeverity sets the severity property value. The severity property
func (m *UserExperienceAnalyticsInsight) SetSeverity(value *UserExperienceAnalyticsInsightSeverity)() {
    m.severity = value
}
// SetUserExperienceAnalyticsMetricId sets the userExperienceAnalyticsMetricId property value. The unique identifier of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) SetUserExperienceAnalyticsMetricId(value *string)() {
    m.userExperienceAnalyticsMetricId = value
}
// SetValues sets the values property value. The value of the user experience analytics insight.
func (m *UserExperienceAnalyticsInsight) SetValues(value []UserExperienceAnalyticsInsightValueable)() {
    m.values = value
}
