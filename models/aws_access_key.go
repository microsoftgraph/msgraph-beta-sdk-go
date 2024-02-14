package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AwsAccessKey struct {
    AwsIdentity
}
// NewAwsAccessKey instantiates a new AwsAccessKey and sets the default values.
func NewAwsAccessKey()(*AwsAccessKey) {
    m := &AwsAccessKey{
        AwsIdentity: *NewAwsIdentity(),
    }
    odataTypeValue := "#microsoft.graph.awsAccessKey"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAwsAccessKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAwsAccessKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsAccessKey(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AwsAccessKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AwsIdentity.GetFieldDeserializers()
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAwsUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(AwsUserable))
        }
        return nil
    }
    return res
}
// GetOwner gets the owner property value. Represents the owner of the access key.
// returns a AwsUserable when successful
func (m *AwsAccessKey) GetOwner()(AwsUserable) {
    val, err := m.GetBackingStore().Get("owner")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AwsUserable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AwsAccessKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AwsIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOwner sets the owner property value. Represents the owner of the access key.
func (m *AwsAccessKey) SetOwner(value AwsUserable)() {
    err := m.GetBackingStore().Set("owner", value)
    if err != nil {
        panic(err)
    }
}
type AwsAccessKeyable interface {
    AwsIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOwner()(AwsUserable)
    SetOwner(value AwsUserable)()
}
