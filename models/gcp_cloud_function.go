package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GcpCloudFunction struct {
    GcpIdentity
}
// NewGcpCloudFunction instantiates a new GcpCloudFunction and sets the default values.
func NewGcpCloudFunction()(*GcpCloudFunction) {
    m := &GcpCloudFunction{
        GcpIdentity: *NewGcpIdentity(),
    }
    odataTypeValue := "#microsoft.graph.gcpCloudFunction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGcpCloudFunctionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGcpCloudFunctionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGcpCloudFunction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GcpCloudFunction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GcpIdentity.GetFieldDeserializers()
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGcpAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(GcpAuthorizationSystemResourceable))
        }
        return nil
    }
    return res
}
// GetResource gets the resource property value. Represents the resources in an authorization system..
// returns a GcpAuthorizationSystemResourceable when successful
func (m *GcpCloudFunction) GetResource()(GcpAuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("resource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GcpAuthorizationSystemResourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GcpCloudFunction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GcpIdentity.Serialize(writer)
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
// SetResource sets the resource property value. Represents the resources in an authorization system..
func (m *GcpCloudFunction) SetResource(value GcpAuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("resource", value)
    if err != nil {
        panic(err)
    }
}
type GcpCloudFunctionable interface {
    GcpIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResource()(GcpAuthorizationSystemResourceable)
    SetResource(value GcpAuthorizationSystemResourceable)()
}
