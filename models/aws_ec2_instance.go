package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AwsEc2Instance 
type AwsEc2Instance struct {
    AwsIdentity
}
// NewAwsEc2Instance instantiates a new awsEc2Instance and sets the default values.
func NewAwsEc2Instance()(*AwsEc2Instance) {
    m := &AwsEc2Instance{
        AwsIdentity: *NewAwsIdentity(),
    }
    odataTypeValue := "#microsoft.graph.awsEc2Instance"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAwsEc2InstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAwsEc2InstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsEc2Instance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AwsEc2Instance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AwsIdentity.GetFieldDeserializers()
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAwsAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(AwsAuthorizationSystemResourceable))
        }
        return nil
    }
    return res
}
// GetResource gets the resource property value. The resource property
func (m *AwsEc2Instance) GetResource()(AwsAuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("resource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AwsAuthorizationSystemResourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AwsEc2Instance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AwsIdentity.Serialize(writer)
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
// SetResource sets the resource property value. The resource property
func (m *AwsEc2Instance) SetResource(value AwsAuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("resource", value)
    if err != nil {
        panic(err)
    }
}
// AwsEc2Instanceable 
type AwsEc2Instanceable interface {
    AwsIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResource()(AwsAuthorizationSystemResourceable)
    SetResource(value AwsAuthorizationSystemResourceable)()
}
