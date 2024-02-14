package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EvaluateLabelJobResponse struct {
    JobResponseBase
}
// NewEvaluateLabelJobResponse instantiates a new EvaluateLabelJobResponse and sets the default values.
func NewEvaluateLabelJobResponse()(*EvaluateLabelJobResponse) {
    m := &EvaluateLabelJobResponse{
        JobResponseBase: *NewJobResponseBase(),
    }
    return m
}
// CreateEvaluateLabelJobResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEvaluateLabelJobResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEvaluateLabelJobResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EvaluateLabelJobResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.JobResponseBase.GetFieldDeserializers()
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEvaluateLabelJobResultGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(EvaluateLabelJobResultGroupable))
        }
        return nil
    }
    return res
}
// GetResult gets the result property value. The result property
// returns a EvaluateLabelJobResultGroupable when successful
func (m *EvaluateLabelJobResponse) GetResult()(EvaluateLabelJobResultGroupable) {
    val, err := m.GetBackingStore().Get("result")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EvaluateLabelJobResultGroupable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EvaluateLabelJobResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.JobResponseBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("result", m.GetResult())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResult sets the result property value. The result property
func (m *EvaluateLabelJobResponse) SetResult(value EvaluateLabelJobResultGroupable)() {
    err := m.GetBackingStore().Set("result", value)
    if err != nil {
        panic(err)
    }
}
type EvaluateLabelJobResponseable interface {
    JobResponseBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResult()(EvaluateLabelJobResultGroupable)
    SetResult(value EvaluateLabelJobResultGroupable)()
}
