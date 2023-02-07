package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody 
type ItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The memberId property
    memberId *string
    // The membershipRule property
    membershipRule *string
}
// NewItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody instantiates a new ItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody and sets the default values.
func NewItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody()(*ItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) {
    m := &ItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) GetMemberId()(*string) {
    return m.memberId
}
// GetMembershipRule gets the membershipRule property value. The membershipRule property
func (m *ItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) GetMembershipRule()(*string) {
    return m.membershipRule
}
// Serialize serializes information the current object
func (m *ItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMemberId sets the memberId property value. The memberId property
func (m *ItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) SetMemberId(value *string)() {
    m.memberId = value
}
// SetMembershipRule sets the membershipRule property value. The membershipRule property
func (m *ItemJoinedGroupsMicrosoftGraphEvaluateDynamicMembershipEvaluateDynamicMembershipPostRequestBody) SetMembershipRule(value *string)() {
    m.membershipRule = value
}
