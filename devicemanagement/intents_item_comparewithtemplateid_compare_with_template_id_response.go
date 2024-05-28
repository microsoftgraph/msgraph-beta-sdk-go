package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use IntentsItemComparewithtemplateidCompareWithTemplateIdGetResponseable instead.
type IntentsItemComparewithtemplateidCompareWithTemplateIdResponse struct {
    IntentsItemComparewithtemplateidCompareWithTemplateIdGetResponse
}
// NewIntentsItemComparewithtemplateidCompareWithTemplateIdResponse instantiates a new IntentsItemComparewithtemplateidCompareWithTemplateIdResponse and sets the default values.
func NewIntentsItemComparewithtemplateidCompareWithTemplateIdResponse()(*IntentsItemComparewithtemplateidCompareWithTemplateIdResponse) {
    m := &IntentsItemComparewithtemplateidCompareWithTemplateIdResponse{
        IntentsItemComparewithtemplateidCompareWithTemplateIdGetResponse: *NewIntentsItemComparewithtemplateidCompareWithTemplateIdGetResponse(),
    }
    return m
}
// CreateIntentsItemComparewithtemplateidCompareWithTemplateIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntentsItemComparewithtemplateidCompareWithTemplateIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntentsItemComparewithtemplateidCompareWithTemplateIdResponse(), nil
}
// Deprecated: This class is obsolete. Use IntentsItemComparewithtemplateidCompareWithTemplateIdGetResponseable instead.
type IntentsItemComparewithtemplateidCompareWithTemplateIdResponseable interface {
    IntentsItemComparewithtemplateidCompareWithTemplateIdGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
