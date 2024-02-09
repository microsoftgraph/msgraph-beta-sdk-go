package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type RuleCondition struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewRuleCondition instantiates a new RuleCondition and sets the default values.
func NewRuleCondition()(*RuleCondition) {
    m := &RuleCondition{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRuleConditionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRuleConditionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRuleCondition(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RuleCondition) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAggregation gets the aggregation property value. The built-in aggregation method for the rule condition. The possible values are: count, percentage, affectedCloudPcCount, affectedCloudPcPercentage, unknownFutureValue.
// returns a *AggregationType when successful
func (m *RuleCondition) GetAggregation()(*AggregationType) {
    val, err := m.GetBackingStore().Get("aggregation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AggregationType)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *RuleCondition) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetConditionCategory gets the conditionCategory property value. The property that the rule condition monitors. Possible values are:  provisionFailures, imageUploadFailures, azureNetworkConnectionCheckFailures, cloudPcInGracePeriod, frontlineInsufficientLicenses, cloudPcConnectionErrors, cloudPcHostHealthCheckFailures, cloudPcZoneOutage, unknownFutureValue.
// returns a *ConditionCategory when successful
func (m *RuleCondition) GetConditionCategory()(*ConditionCategory) {
    val, err := m.GetBackingStore().Get("conditionCategory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConditionCategory)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RuleCondition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aggregation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAggregationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAggregation(val.(*AggregationType))
        }
        return nil
    }
    res["conditionCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionCategory(val.(*ConditionCategory))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperatorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val.(*OperatorType))
        }
        return nil
    }
    res["relationshipType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRelationshipType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelationshipType(val.(*RelationshipType))
        }
        return nil
    }
    res["thresholdValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThresholdValue(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *RuleCondition) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperator gets the operator property value. The built-in operator for the rule condition. The possible values are: greaterOrEqual, equal, greater, less, lessOrEqual, notEqual, unknownFutureValue.
// returns a *OperatorType when successful
func (m *RuleCondition) GetOperator()(*OperatorType) {
    val, err := m.GetBackingStore().Get("operator")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OperatorType)
    }
    return nil
}
// GetRelationshipType gets the relationshipType property value. The relationship type.  Possible values are: and, or.
// returns a *RelationshipType when successful
func (m *RuleCondition) GetRelationshipType()(*RelationshipType) {
    val, err := m.GetBackingStore().Get("relationshipType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RelationshipType)
    }
    return nil
}
// GetThresholdValue gets the thresholdValue property value. The threshold value of the alert condition. The threshold value can be a number in string form or string like 'WestUS'.
// returns a *string when successful
func (m *RuleCondition) GetThresholdValue()(*string) {
    val, err := m.GetBackingStore().Get("thresholdValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RuleCondition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAggregation() != nil {
        cast := (*m.GetAggregation()).String()
        err := writer.WriteStringValue("aggregation", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionCategory() != nil {
        cast := (*m.GetConditionCategory()).String()
        err := writer.WriteStringValue("conditionCategory", &cast)
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
    if m.GetRelationshipType() != nil {
        cast := (*m.GetRelationshipType()).String()
        err := writer.WriteStringValue("relationshipType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("thresholdValue", m.GetThresholdValue())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RuleCondition) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAggregation sets the aggregation property value. The built-in aggregation method for the rule condition. The possible values are: count, percentage, affectedCloudPcCount, affectedCloudPcPercentage, unknownFutureValue.
func (m *RuleCondition) SetAggregation(value *AggregationType)() {
    err := m.GetBackingStore().Set("aggregation", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *RuleCondition) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetConditionCategory sets the conditionCategory property value. The property that the rule condition monitors. Possible values are:  provisionFailures, imageUploadFailures, azureNetworkConnectionCheckFailures, cloudPcInGracePeriod, frontlineInsufficientLicenses, cloudPcConnectionErrors, cloudPcHostHealthCheckFailures, cloudPcZoneOutage, unknownFutureValue.
func (m *RuleCondition) SetConditionCategory(value *ConditionCategory)() {
    err := m.GetBackingStore().Set("conditionCategory", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RuleCondition) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOperator sets the operator property value. The built-in operator for the rule condition. The possible values are: greaterOrEqual, equal, greater, less, lessOrEqual, notEqual, unknownFutureValue.
func (m *RuleCondition) SetOperator(value *OperatorType)() {
    err := m.GetBackingStore().Set("operator", value)
    if err != nil {
        panic(err)
    }
}
// SetRelationshipType sets the relationshipType property value. The relationship type.  Possible values are: and, or.
func (m *RuleCondition) SetRelationshipType(value *RelationshipType)() {
    err := m.GetBackingStore().Set("relationshipType", value)
    if err != nil {
        panic(err)
    }
}
// SetThresholdValue sets the thresholdValue property value. The threshold value of the alert condition. The threshold value can be a number in string form or string like 'WestUS'.
func (m *RuleCondition) SetThresholdValue(value *string)() {
    err := m.GetBackingStore().Set("thresholdValue", value)
    if err != nil {
        panic(err)
    }
}
type RuleConditionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAggregation()(*AggregationType)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetConditionCategory()(*ConditionCategory)
    GetOdataType()(*string)
    GetOperator()(*OperatorType)
    GetRelationshipType()(*RelationshipType)
    GetThresholdValue()(*string)
    SetAggregation(value *AggregationType)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetConditionCategory(value *ConditionCategory)()
    SetOdataType(value *string)()
    SetOperator(value *OperatorType)()
    SetRelationshipType(value *RelationshipType)()
    SetThresholdValue(value *string)()
}
