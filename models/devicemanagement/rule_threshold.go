package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RuleThreshold 
type RuleThreshold struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The aggregation property
    aggregation *AggregationType
    // The OdataType property
    odataType *string
    // The operator property
    operator *OperatorType
    // The target property
    target *int32
}
// NewRuleThreshold instantiates a new ruleThreshold and sets the default values.
func NewRuleThreshold()(*RuleThreshold) {
    m := &RuleThreshold{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.deviceManagement.ruleThreshold";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRuleThresholdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRuleThresholdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRuleThreshold(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RuleThreshold) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAggregation gets the aggregation property value. The aggregation property
func (m *RuleThreshold) GetAggregation()(*AggregationType) {
    return m.aggregation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RuleThreshold) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aggregation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAggregationType , m.SetAggregation)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["operator"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseOperatorType , m.SetOperator)
    res["target"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTarget)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RuleThreshold) GetOdataType()(*string) {
    return m.odataType
}
// GetOperator gets the operator property value. The operator property
func (m *RuleThreshold) GetOperator()(*OperatorType) {
    return m.operator
}
// GetTarget gets the target property value. The target property
func (m *RuleThreshold) GetTarget()(*int32) {
    return m.target
}
// Serialize serializes information the current object
func (m *RuleThreshold) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAggregation() != nil {
        cast := (*m.GetAggregation()).String()
        err := writer.WriteStringValue("aggregation", &cast)
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
    if m.GetOperator() != nil {
        cast := (*m.GetOperator()).String()
        err := writer.WriteStringValue("operator", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("target", m.GetTarget())
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
func (m *RuleThreshold) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAggregation sets the aggregation property value. The aggregation property
func (m *RuleThreshold) SetAggregation(value *AggregationType)() {
    m.aggregation = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RuleThreshold) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperator sets the operator property value. The operator property
func (m *RuleThreshold) SetOperator(value *OperatorType)() {
    m.operator = value
}
// SetTarget sets the target property value. The target property
func (m *RuleThreshold) SetTarget(value *int32)() {
    m.target = value
}
