package evaluatedynamicmembership

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EvaluateDynamicMembershipRequestBody provides operations to call the evaluateDynamicMembership method.
type EvaluateDynamicMembershipRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    memberId *string;
    // 
    membershipRule *string;
}
// NewEvaluateDynamicMembershipRequestBody instantiates a new evaluateDynamicMembershipRequestBody and sets the default values.
func NewEvaluateDynamicMembershipRequestBody()(*EvaluateDynamicMembershipRequestBody) {
    m := &EvaluateDynamicMembershipRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEvaluateDynamicMembershipRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEvaluateDynamicMembershipRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEvaluateDynamicMembershipRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateDynamicMembershipRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateDynamicMembershipRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["memberId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberId(val)
        }
        return nil
    }
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
    return res
}
// GetMemberId gets the memberId property value. 
func (m *EvaluateDynamicMembershipRequestBody) GetMemberId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberId
    }
}
// GetMembershipRule gets the membershipRule property value. 
func (m *EvaluateDynamicMembershipRequestBody) GetMembershipRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.membershipRule
    }
}
func (m *EvaluateDynamicMembershipRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EvaluateDynamicMembershipRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("memberId", m.GetMemberId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("membershipRule", m.GetMembershipRule())
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
func (m *EvaluateDynamicMembershipRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMemberId sets the memberId property value. 
func (m *EvaluateDynamicMembershipRequestBody) SetMemberId(value *string)() {
    if m != nil {
        m.memberId = value
    }
}
// SetMembershipRule sets the membershipRule property value. 
func (m *EvaluateDynamicMembershipRequestBody) SetMembershipRule(value *string)() {
    if m != nil {
        m.membershipRule = value
    }
}
