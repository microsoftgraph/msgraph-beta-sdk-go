package governanceroleassignments

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExportResponse 
// Deprecated: This class is obsolete. Use exportGetResponse instead.
type ExportResponse struct {
    ExportGetResponse
}
// NewExportResponse instantiates a new ExportResponse and sets the default values.
func NewExportResponse()(*ExportResponse) {
    m := &ExportResponse{
        ExportGetResponse: *NewExportGetResponse(),
    }
    return m
}
// CreateExportResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExportResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExportResponse(), nil
}
// ExportResponseable 
// Deprecated: This class is obsolete. Use exportGetResponse instead.
type ExportResponseable interface {
    ExportGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
