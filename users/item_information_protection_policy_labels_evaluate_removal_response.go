package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInformationProtectionPolicyLabelsEvaluateRemovalResponse 
// Deprecated: This class is obsolete. Use evaluateRemovalPostResponse instead.
type ItemInformationProtectionPolicyLabelsEvaluateRemovalResponse struct {
    ItemInformationProtectionPolicyLabelsEvaluateRemovalPostResponse
}
// NewItemInformationProtectionPolicyLabelsEvaluateRemovalResponse instantiates a new ItemInformationProtectionPolicyLabelsEvaluateRemovalResponse and sets the default values.
func NewItemInformationProtectionPolicyLabelsEvaluateRemovalResponse()(*ItemInformationProtectionPolicyLabelsEvaluateRemovalResponse) {
    m := &ItemInformationProtectionPolicyLabelsEvaluateRemovalResponse{
        ItemInformationProtectionPolicyLabelsEvaluateRemovalPostResponse: *NewItemInformationProtectionPolicyLabelsEvaluateRemovalPostResponse(),
    }
    return m
}
// CreateItemInformationProtectionPolicyLabelsEvaluateRemovalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInformationProtectionPolicyLabelsEvaluateRemovalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInformationProtectionPolicyLabelsEvaluateRemovalResponse(), nil
}
// ItemInformationProtectionPolicyLabelsEvaluateRemovalResponseable 
// Deprecated: This class is obsolete. Use evaluateRemovalPostResponse instead.
type ItemInformationProtectionPolicyLabelsEvaluateRemovalResponseable interface {
    ItemInformationProtectionPolicyLabelsEvaluateRemovalPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
