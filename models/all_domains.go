package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AllDomains struct {
    ValidatingDomains
}
// NewAllDomains instantiates a new AllDomains and sets the default values.
func NewAllDomains()(*AllDomains) {
    m := &AllDomains{
        ValidatingDomains: *NewValidatingDomains(),
    }
    odataTypeValue := "#microsoft.graph.allDomains"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAllDomainsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAllDomainsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAllDomains(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AllDomains) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ValidatingDomains.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AllDomains) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ValidatingDomains.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type AllDomainsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ValidatingDomainsable
}
