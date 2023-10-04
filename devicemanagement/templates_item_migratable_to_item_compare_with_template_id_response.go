package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TemplatesItemMigratableToItemCompareWithTemplateIdResponse 
// Deprecated: This class is obsolete. Use compareWithTemplateIdGetResponse instead.
type TemplatesItemMigratableToItemCompareWithTemplateIdResponse struct {
    TemplatesItemMigratableToItemCompareWithTemplateIdGetResponse
}
// NewTemplatesItemMigratableToItemCompareWithTemplateIdResponse instantiates a new TemplatesItemMigratableToItemCompareWithTemplateIdResponse and sets the default values.
func NewTemplatesItemMigratableToItemCompareWithTemplateIdResponse()(*TemplatesItemMigratableToItemCompareWithTemplateIdResponse) {
    m := &TemplatesItemMigratableToItemCompareWithTemplateIdResponse{
        TemplatesItemMigratableToItemCompareWithTemplateIdGetResponse: *NewTemplatesItemMigratableToItemCompareWithTemplateIdGetResponse(),
    }
    return m
}
// CreateTemplatesItemMigratableToItemCompareWithTemplateIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTemplatesItemMigratableToItemCompareWithTemplateIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTemplatesItemMigratableToItemCompareWithTemplateIdResponse(), nil
}
// TemplatesItemMigratableToItemCompareWithTemplateIdResponseable 
// Deprecated: This class is obsolete. Use compareWithTemplateIdGetResponse instead.
type TemplatesItemMigratableToItemCompareWithTemplateIdResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TemplatesItemMigratableToItemCompareWithTemplateIdGetResponseable
}
