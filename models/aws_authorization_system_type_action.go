package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AwsAuthorizationSystemTypeAction 
type AwsAuthorizationSystemTypeAction struct {
    AuthorizationSystemTypeAction
}
// NewAwsAuthorizationSystemTypeAction instantiates a new awsAuthorizationSystemTypeAction and sets the default values.
func NewAwsAuthorizationSystemTypeAction()(*AwsAuthorizationSystemTypeAction) {
    m := &AwsAuthorizationSystemTypeAction{
        AuthorizationSystemTypeAction: *NewAuthorizationSystemTypeAction(),
    }
    return m
}
// CreateAwsAuthorizationSystemTypeActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAwsAuthorizationSystemTypeActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsAuthorizationSystemTypeAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AwsAuthorizationSystemTypeAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *AwsAuthorizationSystemTypeAction) GetService()(AuthorizationSystemTypeServiceable) {
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
func (m *AwsAuthorizationSystemTypeAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AwsAuthorizationSystemTypeAction) SetService(value AuthorizationSystemTypeServiceable)() {
    err := m.GetBackingStore().Set("service", value)
    if err != nil {
        panic(err)
    }
}
// AwsAuthorizationSystemTypeActionable 
type AwsAuthorizationSystemTypeActionable interface {
    AuthorizationSystemTypeActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetService()(AuthorizationSystemTypeServiceable)
    SetService(value AuthorizationSystemTypeServiceable)()
}
