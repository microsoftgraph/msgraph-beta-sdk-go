package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody provides operations to call the evaluateDynamicMembership method.
type ItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The memberId property
    memberId *string
    // The membershipRule property
    membershipRule *string
}
// NewItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody instantiates a new ItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody and sets the default values.
func NewItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody()(*ItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody) {
    m := &ItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemJoinedGroupsEvaluateDynamicMembershipPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedGroupsEvaluateDynamicMembershipPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody) GetMemberId()(*string) {
    return m.memberId
}
// GetMembershipRule gets the membershipRule property value. The membershipRule property
func (m *ItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody) GetMembershipRule()(*string) {
    return m.membershipRule
}
// Serialize serializes information the current object
func (m *ItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMemberId sets the memberId property value. The memberId property
func (m *ItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody) SetMemberId(value *string)() {
    m.memberId = value
}
// SetMembershipRule sets the membershipRule property value. The membershipRule property
func (m *ItemJoinedGroupsEvaluateDynamicMembershipPostRequestBody) SetMembershipRule(value *string)() {
    m.membershipRule = value
}
