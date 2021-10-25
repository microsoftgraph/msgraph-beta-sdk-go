package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EvaluateDynamicMembershipResult struct {
    additionalData map[string]interface{};
    membershipRule *string;
    membershipRuleEvaluationDetails *ExpressionEvaluationDetails;
    membershipRuleEvaluationResult *bool;
}
func NewEvaluateDynamicMembershipResult()(*EvaluateDynamicMembershipResult) {
    m := &EvaluateDynamicMembershipResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EvaluateDynamicMembershipResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *EvaluateDynamicMembershipResult) GetMembershipRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.membershipRule
    }
}
func (m *EvaluateDynamicMembershipResult) GetMembershipRuleEvaluationDetails()(*ExpressionEvaluationDetails) {
    if m == nil {
        return nil
    } else {
        return m.membershipRuleEvaluationDetails
    }
}
func (m *EvaluateDynamicMembershipResult) GetMembershipRuleEvaluationResult()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.membershipRuleEvaluationResult
    }
}
func (m *EvaluateDynamicMembershipResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["membershipRule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMembershipRule(val)
        return nil
    }
    res["membershipRuleEvaluationDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExpressionEvaluationDetails() })
        if err != nil {
            return err
        }
        m.SetMembershipRuleEvaluationDetails(val.(*ExpressionEvaluationDetails))
        return nil
    }
    res["membershipRuleEvaluationResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMembershipRuleEvaluationResult(val)
        return nil
    }
    return res
}
func (m *EvaluateDynamicMembershipResult) IsNil()(bool) {
    return m == nil
}
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
func (m *EvaluateDynamicMembershipResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EvaluateDynamicMembershipResult) SetMembershipRule(value *string)() {
    m.membershipRule = value
}
func (m *EvaluateDynamicMembershipResult) SetMembershipRuleEvaluationDetails(value *ExpressionEvaluationDetails)() {
    m.membershipRuleEvaluationDetails = value
}
func (m *EvaluateDynamicMembershipResult) SetMembershipRuleEvaluationResult(value *bool)() {
    m.membershipRuleEvaluationResult = value
}
