package devicemanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlertImpact 
type AlertImpact struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The aggregation type of the impact. The possible values are: count, percentage, affectedCloudPcCount, affectedCloudPcPercentage, unknownFutureValue.
    aggregationType *AggregationType
    // The OdataType property
    odataType *string
    // The number value of the impact.
    value *int32
}
// NewAlertImpact instantiates a new alertImpact and sets the default values.
func NewAlertImpact()(*AlertImpact) {
    m := &AlertImpact{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAlertImpactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertImpactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertImpact(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertImpact) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAggregationType gets the aggregationType property value. The aggregation type of the impact. The possible values are: count, percentage, affectedCloudPcCount, affectedCloudPcPercentage, unknownFutureValue.
func (m *AlertImpact) GetAggregationType()(*AggregationType) {
    return m.aggregationType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertImpact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aggregationType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAggregationType , m.SetAggregationType)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetValue)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AlertImpact) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. The number value of the impact.
func (m *AlertImpact) GetValue()(*int32) {
    return m.value
}
// Serialize serializes information the current object
func (m *AlertImpact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAggregationType() != nil {
        cast := (*m.GetAggregationType()).String()
        err := writer.WriteStringValue("aggregationType", &cast)
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
    {
        err := writer.WriteInt32Value("value", m.GetValue())
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
func (m *AlertImpact) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAggregationType sets the aggregationType property value. The aggregation type of the impact. The possible values are: count, percentage, affectedCloudPcCount, affectedCloudPcPercentage, unknownFutureValue.
func (m *AlertImpact) SetAggregationType(value *AggregationType)() {
    m.aggregationType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AlertImpact) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. The number value of the impact.
func (m *AlertImpact) SetValue(value *int32)() {
    m.value = value
}
