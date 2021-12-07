package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EvaluateDynamicMembershipResult 
type EvaluateDynamicMembershipResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If a group ID is provided, the value is the membership rule for the group. If a group ID is not provided, the value is the membership rule that was provided as a parameter. For more information, see Dynamic membership rules for groups in Azure Active Directory.
    membershipRule *string;
    // Provides a detailed anaylsis of the membership evaluation result.
    membershipRuleEvaluationDetails *ExpressionEvaluationDetails;
    // The value is true if the user or device is a member of the group. The value can also be true if a membership rule was provided and the user or device passes the rule evaluation; otherwise false.
    membershipRuleEvaluationResult *bool;
}
// NewEvaluateDynamicMembershipResult instantiates a new evaluateDynamicMembershipResult and sets the default values.
func NewEvaluateDynamicMembershipResult()(*EvaluateDynamicMembershipResult) {
    m := &EvaluateDynamicMembershipResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateDynamicMembershipResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMembershipRule gets the membershipRule property value. If a group ID is provided, the value is the membership rule for the group. If a group ID is not provided, the value is the membership rule that was provided as a parameter. For more information, see Dynamic membership rules for groups in Azure Active Directory.
func (m *EvaluateDynamicMembershipResult) GetMembershipRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.membershipRule
    }
}
// GetMembershipRuleEvaluationDetails gets the membershipRuleEvaluationDetails property value. Provides a detailed anaylsis of the membership evaluation result.
func (m *EvaluateDynamicMembershipResult) GetMembershipRuleEvaluationDetails()(*ExpressionEvaluationDetails) {
    if m == nil {
        return nil
    } else {
        return m.membershipRuleEvaluationDetails
    }
}
// GetMembershipRuleEvaluationResult gets the membershipRuleEvaluationResult property value. The value is true if the user or device is a member of the group. The value can also be true if a membership rule was provided and the user or device passes the rule evaluation; otherwise false.
func (m *EvaluateDynamicMembershipResult) GetMembershipRuleEvaluationResult()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.membershipRuleEvaluationResult
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateDynamicMembershipResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["membershipRule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRule(val)
        }
        return nil
    }
    res["membershipRuleEvaluationDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExpressionEvaluationDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRuleEvaluationDetails(val.(*ExpressionEvaluationDetails))
        }
        return nil
    }
    res["membershipRuleEvaluationResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRuleEvaluationResult(val)
        }
        return nil
    }
    return res
}
func (m *EvaluateDynamicMembershipResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EvaluateDynamicMembershipResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateDynamicMembershipResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMembershipRule sets the membershipRule property value. If a group ID is provided, the value is the membership rule for the group. If a group ID is not provided, the value is the membership rule that was provided as a parameter. For more information, see Dynamic membership rules for groups in Azure Active Directory.
func (m *EvaluateDynamicMembershipResult) SetMembershipRule(value *string)() {
    if m != nil {
        m.membershipRule = value
    }
}
// SetMembershipRuleEvaluationDetails sets the membershipRuleEvaluationDetails property value. Provides a detailed anaylsis of the membership evaluation result.
func (m *EvaluateDynamicMembershipResult) SetMembershipRuleEvaluationDetails(value *ExpressionEvaluationDetails)() {
    if m != nil {
        m.membershipRuleEvaluationDetails = value
    }
}
// SetMembershipRuleEvaluationResult sets the membershipRuleEvaluationResult property value. The value is true if the user or device is a member of the group. The value can also be true if a membership rule was provided and the user or device passes the rule evaluation; otherwise false.
func (m *EvaluateDynamicMembershipResult) SetMembershipRuleEvaluationResult(value *bool)() {
    if m != nil {
        m.membershipRuleEvaluationResult = value
    }
}
