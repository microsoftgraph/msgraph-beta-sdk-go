package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InboundFileFlow 
type InboundFileFlow struct {
    InboundFlow
    // The OdataType property
    OdataType *string
}
// NewInboundFileFlow instantiates a new inboundFileFlow and sets the default values.
func NewInboundFileFlow()(*InboundFileFlow) {
    m := &InboundFileFlow{
        InboundFlow: *NewInboundFlow(),
    }
    odataTypeValue := "#microsoft.graph.industryData.inboundFileFlow"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateInboundFileFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInboundFileFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInboundFileFlow(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InboundFileFlow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InboundFlow.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *InboundFileFlow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InboundFlow.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// InboundFileFlowable 
type InboundFileFlowable interface {
    InboundFlowable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
