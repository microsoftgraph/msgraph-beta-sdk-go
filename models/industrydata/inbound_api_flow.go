package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type InboundApiFlow struct {
    InboundFlow
}
// NewInboundApiFlow instantiates a new InboundApiFlow and sets the default values.
func NewInboundApiFlow()(*InboundApiFlow) {
    m := &InboundApiFlow{
        InboundFlow: *NewInboundFlow(),
    }
    odataTypeValue := "#microsoft.graph.industryData.inboundApiFlow"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateInboundApiFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInboundApiFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInboundApiFlow(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InboundApiFlow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InboundFlow.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *InboundApiFlow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InboundFlow.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type InboundApiFlowable interface {
    InboundFlowable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
