package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GcpAuthorizationSystemTypeAction 
type GcpAuthorizationSystemTypeAction struct {
    AuthorizationSystemTypeAction
}
// NewGcpAuthorizationSystemTypeAction instantiates a new gcpAuthorizationSystemTypeAction and sets the default values.
func NewGcpAuthorizationSystemTypeAction()(*GcpAuthorizationSystemTypeAction) {
    m := &GcpAuthorizationSystemTypeAction{
        AuthorizationSystemTypeAction: *NewAuthorizationSystemTypeAction(),
    }
    return m
}
// CreateGcpAuthorizationSystemTypeActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGcpAuthorizationSystemTypeActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGcpAuthorizationSystemTypeAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GcpAuthorizationSystemTypeAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthorizationSystemTypeAction.GetFieldDeserializers()
    res["service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationSystemTypeServiceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val.(AuthorizationSystemTypeServiceable))
        }
        return nil
    }
    return res
}
// GetService gets the service property value. The service property
func (m *GcpAuthorizationSystemTypeAction) GetService()(AuthorizationSystemTypeServiceable) {
    val, err := m.GetBackingStore().Get("service")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemTypeServiceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GcpAuthorizationSystemTypeAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthorizationSystemTypeAction.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetService sets the service property value. The service property
func (m *GcpAuthorizationSystemTypeAction) SetService(value AuthorizationSystemTypeServiceable)() {
    err := m.GetBackingStore().Set("service", value)
    if err != nil {
        panic(err)
    }
}
// GcpAuthorizationSystemTypeActionable 
type GcpAuthorizationSystemTypeActionable interface {
    AuthorizationSystemTypeActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetService()(AuthorizationSystemTypeServiceable)
    SetService(value AuthorizationSystemTypeServiceable)()
}
