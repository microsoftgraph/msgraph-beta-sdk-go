package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosSingleSignOnExtension an abstract base class for all iOS-specific single sign-on extension types.
type IosSingleSignOnExtension struct {
    SingleSignOnExtension
}
// NewIosSingleSignOnExtension instantiates a new iosSingleSignOnExtension and sets the default values.
func NewIosSingleSignOnExtension()(*IosSingleSignOnExtension) {
    m := &IosSingleSignOnExtension{
        SingleSignOnExtension: *NewSingleSignOnExtension(),
    }
    return m
}
// CreateIosSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosSingleSignOnExtension(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SingleSignOnExtension.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *IosSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SingleSignOnExtension.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
