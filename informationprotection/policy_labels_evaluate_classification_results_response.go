// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package informationprotection

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use PolicyLabelsEvaluateClassificationResultsPostResponseable instead.
type PolicyLabelsEvaluateClassificationResultsResponse struct {
    PolicyLabelsEvaluateClassificationResultsPostResponse
}
// NewPolicyLabelsEvaluateClassificationResultsResponse instantiates a new PolicyLabelsEvaluateClassificationResultsResponse and sets the default values.
func NewPolicyLabelsEvaluateClassificationResultsResponse()(*PolicyLabelsEvaluateClassificationResultsResponse) {
    m := &PolicyLabelsEvaluateClassificationResultsResponse{
        PolicyLabelsEvaluateClassificationResultsPostResponse: *NewPolicyLabelsEvaluateClassificationResultsPostResponse(),
    }
    return m
}
// CreatePolicyLabelsEvaluateClassificationResultsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePolicyLabelsEvaluateClassificationResultsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPolicyLabelsEvaluateClassificationResultsResponse(), nil
}
// Deprecated: This class is obsolete. Use PolicyLabelsEvaluateClassificationResultsPostResponseable instead.
type PolicyLabelsEvaluateClassificationResultsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyLabelsEvaluateClassificationResultsPostResponseable
}
