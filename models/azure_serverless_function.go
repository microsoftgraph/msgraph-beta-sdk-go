package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureServerlessFunction 
type AzureServerlessFunction struct {
    AzureIdentity
}
// NewAzureServerlessFunction instantiates a new azureServerlessFunction and sets the default values.
func NewAzureServerlessFunction()(*AzureServerlessFunction) {
    m := &AzureServerlessFunction{
        AzureIdentity: *NewAzureIdentity(),
    }
    odataTypeValue := "#microsoft.graph.azureServerlessFunction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureServerlessFunctionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureServerlessFunctionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureServerlessFunction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureServerlessFunction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AzureIdentity.GetFieldDeserializers()
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAzureAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(AzureAuthorizationSystemResourceable))
        }
        return nil
    }
    return res
}
// GetResource gets the resource property value. Represents the resources in an authorization system.
func (m *AzureServerlessFunction) GetResource()(AzureAuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("resource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AzureAuthorizationSystemResourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AzureServerlessFunction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AzureIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResource sets the resource property value. Represents the resources in an authorization system.
func (m *AzureServerlessFunction) SetResource(value AzureAuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("resource", value)
    if err != nil {
        panic(err)
    }
}
// AzureServerlessFunctionable 
type AzureServerlessFunctionable interface {
    AzureIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResource()(AzureAuthorizationSystemResourceable)
    SetResource(value AzureAuthorizationSystemResourceable)()
}
