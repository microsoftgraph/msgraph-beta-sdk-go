package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureAuthorizationSystemResource 
type AzureAuthorizationSystemResource struct {
    AuthorizationSystemResource
}
// NewAzureAuthorizationSystemResource instantiates a new azureAuthorizationSystemResource and sets the default values.
func NewAzureAuthorizationSystemResource()(*AzureAuthorizationSystemResource) {
    m := &AzureAuthorizationSystemResource{
        AuthorizationSystemResource: *NewAuthorizationSystemResource(),
    }
    return m
}
// CreateAzureAuthorizationSystemResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureAuthorizationSystemResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureAuthorizationSystemResource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureAuthorizationSystemResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetService gets the service property value. The service property
func (m *AzureAuthorizationSystemResource) GetService()(AuthorizationSystemTypeServiceable) {
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
func (m *AzureAuthorizationSystemResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetService sets the service property value. The service property
func (m *AzureAuthorizationSystemResource) SetService(value AuthorizationSystemTypeServiceable)() {
    err := m.GetBackingStore().Set("service", value)
    if err != nil {
        panic(err)
    }
}
// AzureAuthorizationSystemResourceable 
type AzureAuthorizationSystemResourceable interface {
    AuthorizationSystemResourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetService()(AuthorizationSystemTypeServiceable)
    SetService(value AuthorizationSystemTypeServiceable)()
}
