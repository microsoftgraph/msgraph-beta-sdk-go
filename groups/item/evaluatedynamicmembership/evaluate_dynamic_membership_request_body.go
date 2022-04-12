package evaluatedynamicmembership

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EvaluateDynamicMembershipRequestBody provides operations to call the evaluateDynamicMembership method.
type EvaluateDynamicMembershipRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The memberId property
    memberId *string
}
// NewEvaluateDynamicMembershipRequestBody instantiates a new evaluateDynamicMembershipRequestBody and sets the default values.
func NewEvaluateDynamicMembershipRequestBody()(*EvaluateDynamicMembershipRequestBody) {
    m := &EvaluateDynamicMembershipRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEvaluateDynamicMembershipRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEvaluateDynamicMembershipRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
func (m *EvaluateDynamicMembershipRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["memberId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberId(val)
        }
        return nil
    }
    return res
}
// GetMemberId gets the memberId property value. The memberId property
func (m *EvaluateDynamicMembershipRequestBody) GetMemberId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberId
    }
}
// Serialize serializes information the current object
func (m *EvaluateDynamicMembershipRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateDynamicMembershipRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMemberId sets the memberId property value. The memberId property
func (m *EvaluateDynamicMembershipRequestBody) SetMemberId(value *string)() {
    if m != nil {
        m.memberId = value
    }
}
