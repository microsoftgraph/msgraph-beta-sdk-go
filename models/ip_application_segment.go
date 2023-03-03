package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IpApplicationSegment 
type IpApplicationSegment struct {
    ApplicationSegment
}
// NewIpApplicationSegment instantiates a new IpApplicationSegment and sets the default values.
func NewIpApplicationSegment()(*IpApplicationSegment) {
    m := &IpApplicationSegment{
        ApplicationSegment: *NewApplicationSegment(),
    }
    odataTypeValue := "#microsoft.graph.ipApplicationSegment"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIpApplicationSegmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIpApplicationSegmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIpApplicationSegment(), nil
}
// GetDestinationHost gets the destinationHost property value. The destinationHost property
func (m *IpApplicationSegment) GetDestinationHost()(*string) {
    val, err := m.GetBackingStore().Get("destinationHost")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IpApplicationSegment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ApplicationSegment.GetFieldDeserializers()
    res["destinationHost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationHost(val)
        }
        return nil
    }
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    return res
}
// GetPort gets the port property value. The port property
func (m *IpApplicationSegment) GetPort()(*int32) {
    val, err := m.GetBackingStore().Get("port")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IpApplicationSegment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ApplicationSegment.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("destinationHost", m.GetDestinationHost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDestinationHost sets the destinationHost property value. The destinationHost property
func (m *IpApplicationSegment) SetDestinationHost(value *string)() {
    err := m.GetBackingStore().Set("destinationHost", value)
    if err != nil {
        panic(err)
    }
}
// SetPort sets the port property value. The port property
func (m *IpApplicationSegment) SetPort(value *int32)() {
    err := m.GetBackingStore().Set("port", value)
    if err != nil {
        panic(err)
    }
}
// IpApplicationSegmentable 
type IpApplicationSegmentable interface {
    ApplicationSegmentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDestinationHost()(*string)
    GetPort()(*int32)
    SetDestinationHost(value *string)()
    SetPort(value *int32)()
}
