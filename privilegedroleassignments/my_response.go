package privilegedroleassignments

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use MyGetResponseable instead.
type MyResponse struct {
    MyGetResponse
}
// NewMyResponse instantiates a new MyResponse and sets the default values.
func NewMyResponse()(*MyResponse) {
    m := &MyResponse{
        MyGetResponse: *NewMyGetResponse(),
    }
    return m
}
// CreateMyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMyResponse(), nil
}
// Deprecated: This class is obsolete. Use MyGetResponseable instead.
type MyResponseable interface {
    MyGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
