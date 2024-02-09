package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DlpEvaluatePoliciesJobResponse struct {
    JobResponseBase
}
// NewDlpEvaluatePoliciesJobResponse instantiates a new DlpEvaluatePoliciesJobResponse and sets the default values.
func NewDlpEvaluatePoliciesJobResponse()(*DlpEvaluatePoliciesJobResponse) {
    m := &DlpEvaluatePoliciesJobResponse{
        JobResponseBase: *NewJobResponseBase(),
    }
    return m
}
// CreateDlpEvaluatePoliciesJobResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDlpEvaluatePoliciesJobResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDlpEvaluatePoliciesJobResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DlpEvaluatePoliciesJobResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.JobResponseBase.GetFieldDeserializers()
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDlpPoliciesJobResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(DlpPoliciesJobResultable))
        }
        return nil
    }
    return res
}
// GetResult gets the result property value. The result property
// returns a DlpPoliciesJobResultable when successful
func (m *DlpEvaluatePoliciesJobResponse) GetResult()(DlpPoliciesJobResultable) {
    val, err := m.GetBackingStore().Get("result")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DlpPoliciesJobResultable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DlpEvaluatePoliciesJobResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DlpEvaluatePoliciesJobResponse) SetResult(value DlpPoliciesJobResultable)() {
    err := m.GetBackingStore().Set("result", value)
    if err != nil {
        panic(err)
    }
}
type DlpEvaluatePoliciesJobResponseable interface {
    JobResponseBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResult()(DlpPoliciesJobResultable)
    SetResult(value DlpPoliciesJobResultable)()
}
