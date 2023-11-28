package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GsuiteSource 
type GsuiteSource struct {
    AuthorizationSystemIdentitySource
}
// NewGsuiteSource instantiates a new gsuiteSource and sets the default values.
func NewGsuiteSource()(*GsuiteSource) {
    m := &GsuiteSource{
        AuthorizationSystemIdentitySource: *NewAuthorizationSystemIdentitySource(),
    }
    odataTypeValue := "#microsoft.graph.gsuiteSource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGsuiteSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGsuiteSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGsuiteSource(), nil
}
// GetDomain gets the domain property value. Domain name
func (m *GsuiteSource) GetDomain()(*string) {
    val, err := m.GetBackingStore().Get("domain")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GsuiteSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthorizationSystemIdentitySource.GetFieldDeserializers()
    res["domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomain(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GsuiteSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthorizationSystemIdentitySource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("domain", m.GetDomain())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDomain sets the domain property value. Domain name
func (m *GsuiteSource) SetDomain(value *string)() {
    err := m.GetBackingStore().Set("domain", value)
    if err != nil {
        panic(err)
    }
}
// GsuiteSourceable 
type GsuiteSourceable interface {
    AuthorizationSystemIdentitySourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDomain()(*string)
    SetDomain(value *string)()
}
