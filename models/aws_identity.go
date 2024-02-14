package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AwsIdentity struct {
    AuthorizationSystemIdentity
}
// NewAwsIdentity instantiates a new AwsIdentity and sets the default values.
func NewAwsIdentity()(*AwsIdentity) {
    m := &AwsIdentity{
        AuthorizationSystemIdentity: *NewAuthorizationSystemIdentity(),
    }
    odataTypeValue := "#microsoft.graph.awsIdentity"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAwsIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAwsIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.awsAccessKey":
                        return NewAwsAccessKey(), nil
                    case "#microsoft.graph.awsEc2Instance":
                        return NewAwsEc2Instance(), nil
                    case "#microsoft.graph.awsGroup":
                        return NewAwsGroup(), nil
                    case "#microsoft.graph.awsLambda":
                        return NewAwsLambda(), nil
                    case "#microsoft.graph.awsRole":
                        return NewAwsRole(), nil
                    case "#microsoft.graph.awsUser":
                        return NewAwsUser(), nil
                }
            }
        }
    }
    return NewAwsIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AwsIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthorizationSystemIdentity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AwsIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthorizationSystemIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type AwsIdentityable interface {
    AuthorizationSystemIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
