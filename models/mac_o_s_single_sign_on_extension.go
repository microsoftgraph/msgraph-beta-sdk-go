package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSingleSignOnExtension an abstract base class for all macOS-specific single sign-on extension types.
type MacOSSingleSignOnExtension struct {
    SingleSignOnExtension
}
// NewMacOSSingleSignOnExtension instantiates a new macOSSingleSignOnExtension and sets the default values.
func NewMacOSSingleSignOnExtension()(*MacOSSingleSignOnExtension) {
    m := &MacOSSingleSignOnExtension{
        SingleSignOnExtension: *NewSingleSignOnExtension(),
    }
    return m
}
// CreateMacOSSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSSingleSignOnExtension(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SingleSignOnExtension.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *MacOSSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SingleSignOnExtension.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
