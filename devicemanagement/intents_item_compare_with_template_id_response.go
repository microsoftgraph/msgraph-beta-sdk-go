// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use IntentsItemCompareWithTemplateIdGetResponseable instead.
type IntentsItemCompareWithTemplateIdResponse struct {
    IntentsItemCompareWithTemplateIdGetResponse
}
// NewIntentsItemCompareWithTemplateIdResponse instantiates a new IntentsItemCompareWithTemplateIdResponse and sets the default values.
func NewIntentsItemCompareWithTemplateIdResponse()(*IntentsItemCompareWithTemplateIdResponse) {
    m := &IntentsItemCompareWithTemplateIdResponse{
        IntentsItemCompareWithTemplateIdGetResponse: *NewIntentsItemCompareWithTemplateIdGetResponse(),
    }
    return m
}
// CreateIntentsItemCompareWithTemplateIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntentsItemCompareWithTemplateIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntentsItemCompareWithTemplateIdResponse(), nil
}
// Deprecated: This class is obsolete. Use IntentsItemCompareWithTemplateIdGetResponseable instead.
type IntentsItemCompareWithTemplateIdResponseable interface {
    IntentsItemCompareWithTemplateIdGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
