package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EvaluateDynamicMembershipResult 
type EvaluateDynamicMembershipResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // If a group ID is provided, the value is the membership rule for the group. If a group ID is not provided, the value is the membership rule that was provided as a parameter. For more information, see Dynamic membership rules for groups in Azure Active Directory.
    membershipRule *string
    // Provides a detailed anaylsis of the membership evaluation result.
    membershipRuleEvaluationDetails ExpressionEvaluationDetailsable
    // The value is true if the user or device is a member of the group. The value can also be true if a membership rule was provided and the user or device passes the rule evaluation; otherwise false.
    membershipRuleEvaluationResult *bool
    // The OdataType property
    odataType *string
}
// NewEvaluateDynamicMembershipResult instantiates a new evaluateDynamicMembershipResult and sets the default values.
func NewEvaluateDynamicMembershipResult()(*EvaluateDynamicMembershipResult) {
    m := &EvaluateDynamicMembershipResult{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateEvaluateDynamicMembershipResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEvaluateDynamicMembershipResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEvaluateDynamicMembershipResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateDynamicMembershipResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
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
// GetMembershipRule gets the membershipRule property value. If a group ID is provided, the value is the membership rule for the group. If a group ID is not provided, the value is the membership rule that was provided as a parameter. For more information, see Dynamic membership rules for groups in Azure Active Directory.
func (m *EvaluateDynamicMembershipResult) GetMembershipRule()(*string) {
    return m.membershipRule
}
// GetMembershipRuleEvaluationDetails gets the membershipRuleEvaluationDetails property value. Provides a detailed anaylsis of the membership evaluation result.
func (m *EvaluateDynamicMembershipResult) GetMembershipRuleEvaluationDetails()(ExpressionEvaluationDetailsable) {
    return m.membershipRuleEvaluationDetails
}
// GetMembershipRuleEvaluationResult gets the membershipRuleEvaluationResult property value. The value is true if the user or device is a member of the group. The value can also be true if a membership rule was provided and the user or device passes the rule evaluation; otherwise false.
func (m *EvaluateDynamicMembershipResult) GetMembershipRuleEvaluationResult()(*bool) {
    return m.membershipRuleEvaluationResult
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EvaluateDynamicMembershipResult) GetOdataType()(*string) {
    return m.odataType
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateDynamicMembershipResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMembershipRule sets the membershipRule property value. If a group ID is provided, the value is the membership rule for the group. If a group ID is not provided, the value is the membership rule that was provided as a parameter. For more information, see Dynamic membership rules for groups in Azure Active Directory.
func (m *EvaluateDynamicMembershipResult) SetMembershipRule(value *string)() {
    m.membershipRule = value
}
// SetMembershipRuleEvaluationDetails sets the membershipRuleEvaluationDetails property value. Provides a detailed anaylsis of the membership evaluation result.
func (m *EvaluateDynamicMembershipResult) SetMembershipRuleEvaluationDetails(value ExpressionEvaluationDetailsable)() {
    m.membershipRuleEvaluationDetails = value
}
// SetMembershipRuleEvaluationResult sets the membershipRuleEvaluationResult property value. The value is true if the user or device is a member of the group. The value can also be true if a membership rule was provided and the user or device passes the rule evaluation; otherwise false.
func (m *EvaluateDynamicMembershipResult) SetMembershipRuleEvaluationResult(value *bool)() {
    m.membershipRuleEvaluationResult = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EvaluateDynamicMembershipResult) SetOdataType(value *string)() {
    m.odataType = value
}
