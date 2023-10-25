package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureIdentity 
type AzureIdentity struct {
    AuthorizationSystemIdentity
}
// NewAzureIdentity instantiates a new azureIdentity and sets the default values.
func NewAzureIdentity()(*AzureIdentity) {
    m := &AzureIdentity{
        AuthorizationSystemIdentity: *NewAuthorizationSystemIdentity(),
    }
    odataTypeValue := "#microsoft.graph.azureIdentity"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.azureGroup":
                        return NewAzureGroup(), nil
                    case "#microsoft.graph.azureManagedIdentity":
                        return NewAzureManagedIdentity(), nil
                    case "#microsoft.graph.azureServerlessFunction":
                        return NewAzureServerlessFunction(), nil
                    case "#microsoft.graph.azureServicePrincipal":
                        return NewAzureServicePrincipal(), nil
                    case "#microsoft.graph.azureUser":
                        return NewAzureUser(), nil
                }
            }
        }
    }
    return NewAzureIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthorizationSystemIdentity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AzureIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthorizationSystemIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// AzureIdentityable 
type AzureIdentityable interface {
    AuthorizationSystemIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
