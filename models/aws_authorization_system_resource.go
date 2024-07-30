package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AwsAuthorizationSystemResource struct {
    AuthorizationSystemResource
}
// NewAwsAuthorizationSystemResource instantiates a new AwsAuthorizationSystemResource and sets the default values.
func NewAwsAuthorizationSystemResource()(*AwsAuthorizationSystemResource) {
    m := &AwsAuthorizationSystemResource{
        AuthorizationSystemResource: *NewAuthorizationSystemResource(),
    }
    odataTypeValue := "#microsoft.graph.awsAuthorizationSystemResource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAwsAuthorizationSystemResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAwsAuthorizationSystemResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsAuthorizationSystemResource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AwsAuthorizationSystemResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthorizationSystemResource.GetFieldDeserializers()
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
// GetService gets the service property value. The service associated with the resource in an AWS authorization system. This is autoexpanded.
// returns a AuthorizationSystemTypeServiceable when successful
func (m *AwsAuthorizationSystemResource) GetService()(AuthorizationSystemTypeServiceable) {
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
func (m *AwsAuthorizationSystemResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthorizationSystemResource.Serialize(writer)
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
// SetService sets the service property value. The service associated with the resource in an AWS authorization system. This is autoexpanded.
func (m *AwsAuthorizationSystemResource) SetService(value AuthorizationSystemTypeServiceable)() {
    err := m.GetBackingStore().Set("service", value)
    if err != nil {
        panic(err)
    }
}
type AwsAuthorizationSystemResourceable interface {
    AuthorizationSystemResourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetService()(AuthorizationSystemTypeServiceable)
    SetService(value AuthorizationSystemTypeServiceable)()
}
