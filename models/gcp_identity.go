package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GcpIdentity 
type GcpIdentity struct {
    AuthorizationSystemIdentity
}
// NewGcpIdentity instantiates a new gcpIdentity and sets the default values.
func NewGcpIdentity()(*GcpIdentity) {
    m := &GcpIdentity{
        AuthorizationSystemIdentity: *NewAuthorizationSystemIdentity(),
    }
    odataTypeValue := "#microsoft.graph.gcpIdentity"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGcpIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGcpIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.gcpCloudFunction":
                        return NewGcpCloudFunction(), nil
                    case "#microsoft.graph.gcpGroup":
                        return NewGcpGroup(), nil
                    case "#microsoft.graph.gcpServiceAccount":
                        return NewGcpServiceAccount(), nil
                    case "#microsoft.graph.gcpUser":
                        return NewGcpUser(), nil
                }
            }
        }
    }
    return NewGcpIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GcpIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthorizationSystemIdentity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *GcpIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthorizationSystemIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// GcpIdentityable 
type GcpIdentityable interface {
    AuthorizationSystemIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
