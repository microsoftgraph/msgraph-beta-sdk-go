package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody 
type MicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The memberId property
    memberId *string
    // The membershipRule property
    membershipRule *string
}
// NewMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody instantiates a new MicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody and sets the default values.
func NewMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody()(*MicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) {
    m := &MicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetMemberId gets the memberId property value. The memberId property
func (m *MicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) GetMemberId()(*string) {
    return m.memberId
}
// GetMembershipRule gets the membershipRule property value. The membershipRule property
func (m *MicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) GetMembershipRule()(*string) {
    return m.membershipRule
}
// Serialize serializes information the current object
func (m *MicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMemberId sets the memberId property value. The memberId property
func (m *MicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) SetMemberId(value *string)() {
    m.memberId = value
}
// SetMembershipRule sets the membershipRule property value. The membershipRule property
func (m *MicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) SetMembershipRule(value *string)() {
    m.membershipRule = value
}
