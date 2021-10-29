package evaluatedynamicmembership

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EvaluateDynamicMembershipRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    memberId *string;
}
// Instantiates a new evaluateDynamicMembershipRequestBody and sets the default values.
func NewEvaluateDynamicMembershipRequestBody()(*EvaluateDynamicMembershipRequestBody) {
    m := &EvaluateDynamicMembershipRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateDynamicMembershipRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the memberId property value. 
func (m *EvaluateDynamicMembershipRequestBody) GetMemberId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberId
    }
}
// The deserialization information for the current model
func (m *EvaluateDynamicMembershipRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["memberId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMemberId(val)
        return nil
    }
    return res
}
func (m *EvaluateDynamicMembershipRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EvaluateDynamicMembershipRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("memberId", m.GetMemberId())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *EvaluateDynamicMembershipRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the memberId property value. 
// Parameters:
//  - value : Value to set for the memberId property.
func (m *EvaluateDynamicMembershipRequestBody) SetMemberId(value *string)() {
    m.memberId = value
}
