package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// EvaluateDynamicMembershipResult 
type EvaluateDynamicMembershipResult struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEvaluateDynamicMembershipResult instantiates a new evaluateDynamicMembershipResult and sets the default values.
func NewEvaluateDynamicMembershipResult()(*EvaluateDynamicMembershipResult) {
    m := &EvaluateDynamicMembershipResult{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEvaluateDynamicMembershipResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEvaluateDynamicMembershipResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEvaluateDynamicMembershipResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateDynamicMembershipResult) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *EvaluateDynamicMembershipResult) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateDynamicMembershipResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["membershipRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRule(val)
        }
        return nil
    }
    res["membershipRuleEvaluationDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExpressionEvaluationDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRuleEvaluationDetails(val.(ExpressionEvaluationDetailsable))
        }
        return nil
    }
    res["membershipRuleEvaluationResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRuleEvaluationResult(val)
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
    return res
}
// GetMembershipRule gets the membershipRule property value. If a group ID is provided, the value is the membership rule for the group. If a group ID isn't provided, the value is the membership rule that was provided as a parameter. For more information, see Dynamic membership rules for groups in Azure Active Directory.
func (m *EvaluateDynamicMembershipResult) GetMembershipRule()(*string) {
    val, err := m.GetBackingStore().Get("membershipRule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMembershipRuleEvaluationDetails gets the membershipRuleEvaluationDetails property value. Provides a detailed analysis of the membership evaluation result.
func (m *EvaluateDynamicMembershipResult) GetMembershipRuleEvaluationDetails()(ExpressionEvaluationDetailsable) {
    val, err := m.GetBackingStore().Get("membershipRuleEvaluationDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ExpressionEvaluationDetailsable)
    }
    return nil
}
// GetMembershipRuleEvaluationResult gets the membershipRuleEvaluationResult property value. The value is true if the user or device is a member of the group. The value can also be true if a membership rule was provided and the user or device passes the rule evaluation; otherwise false.
func (m *EvaluateDynamicMembershipResult) GetMembershipRuleEvaluationResult()(*bool) {
    val, err := m.GetBackingStore().Get("membershipRuleEvaluationResult")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EvaluateDynamicMembershipResult) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EvaluateDynamicMembershipResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("membershipRule", m.GetMembershipRule())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("membershipRuleEvaluationDetails", m.GetMembershipRuleEvaluationDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("membershipRuleEvaluationResult", m.GetMembershipRuleEvaluationResult())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateDynamicMembershipResult) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *EvaluateDynamicMembershipResult) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetMembershipRule sets the membershipRule property value. If a group ID is provided, the value is the membership rule for the group. If a group ID isn't provided, the value is the membership rule that was provided as a parameter. For more information, see Dynamic membership rules for groups in Azure Active Directory.
func (m *EvaluateDynamicMembershipResult) SetMembershipRule(value *string)() {
    err := m.GetBackingStore().Set("membershipRule", value)
    if err != nil {
        panic(err)
    }
}
// SetMembershipRuleEvaluationDetails sets the membershipRuleEvaluationDetails property value. Provides a detailed analysis of the membership evaluation result.
func (m *EvaluateDynamicMembershipResult) SetMembershipRuleEvaluationDetails(value ExpressionEvaluationDetailsable)() {
    err := m.GetBackingStore().Set("membershipRuleEvaluationDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetMembershipRuleEvaluationResult sets the membershipRuleEvaluationResult property value. The value is true if the user or device is a member of the group. The value can also be true if a membership rule was provided and the user or device passes the rule evaluation; otherwise false.
func (m *EvaluateDynamicMembershipResult) SetMembershipRuleEvaluationResult(value *bool)() {
    err := m.GetBackingStore().Set("membershipRuleEvaluationResult", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EvaluateDynamicMembershipResult) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// EvaluateDynamicMembershipResultable 
type EvaluateDynamicMembershipResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetMembershipRule()(*string)
    GetMembershipRuleEvaluationDetails()(ExpressionEvaluationDetailsable)
    GetMembershipRuleEvaluationResult()(*bool)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetMembershipRule(value *string)()
    SetMembershipRuleEvaluationDetails(value ExpressionEvaluationDetailsable)()
    SetMembershipRuleEvaluationResult(value *bool)()
    SetOdataType(value *string)()
}
