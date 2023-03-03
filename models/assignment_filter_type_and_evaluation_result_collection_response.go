package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterTypeAndEvaluationResultCollectionResponse 
type AssignmentFilterTypeAndEvaluationResultCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewAssignmentFilterTypeAndEvaluationResultCollectionResponse instantiates a new AssignmentFilterTypeAndEvaluationResultCollectionResponse and sets the default values.
func NewAssignmentFilterTypeAndEvaluationResultCollectionResponse()(*AssignmentFilterTypeAndEvaluationResultCollectionResponse) {
    m := &AssignmentFilterTypeAndEvaluationResultCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateAssignmentFilterTypeAndEvaluationResultCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterTypeAndEvaluationResultCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterTypeAndEvaluationResultCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterTypeAndEvaluationResultCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignmentFilterTypeAndEvaluationResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignmentFilterTypeAndEvaluationResultable, len(val))
            for i, v := range val {
                res[i] = v.(AssignmentFilterTypeAndEvaluationResultable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *AssignmentFilterTypeAndEvaluationResultCollectionResponse) GetValue()([]AssignmentFilterTypeAndEvaluationResultable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AssignmentFilterTypeAndEvaluationResultable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AssignmentFilterTypeAndEvaluationResultCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *AssignmentFilterTypeAndEvaluationResultCollectionResponse) SetValue(value []AssignmentFilterTypeAndEvaluationResultable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// AssignmentFilterTypeAndEvaluationResultCollectionResponseable 
type AssignmentFilterTypeAndEvaluationResultCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]AssignmentFilterTypeAndEvaluationResultable)
    SetValue(value []AssignmentFilterTypeAndEvaluationResultable)()
}
