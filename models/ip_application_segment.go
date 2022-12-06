package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IpApplicationSegment 
type IpApplicationSegment struct {
    ApplicationSegment
    // The destinationHost property
    destinationHost *string
    // The port property
    port *int32
}
// NewIpApplicationSegment instantiates a new IpApplicationSegment and sets the default values.
func NewIpApplicationSegment()(*IpApplicationSegment) {
    m := &IpApplicationSegment{
        ApplicationSegment: *NewApplicationSegment(),
    }
    odataTypeValue := "#microsoft.graph.ipApplicationSegment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIpApplicationSegmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIpApplicationSegmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIpApplicationSegment(), nil
}
// GetDestinationHost gets the destinationHost property value. The destinationHost property
func (m *IpApplicationSegment) GetDestinationHost()(*string) {
    return m.destinationHost
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IpApplicationSegment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ApplicationSegment.GetFieldDeserializers()
    res["destinationHost"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDestinationHost)
    res["port"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPort)
    return res
}
// GetPort gets the port property value. The port property
func (m *IpApplicationSegment) GetPort()(*int32) {
    return m.port
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
    m.destinationHost = value
}
// SetPort sets the port property value. The port property
func (m *IpApplicationSegment) SetPort(value *int32)() {
    m.port = value
}
