package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AssignmentFilterEvaluationStatusDetailsCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewAssignmentFilterEvaluationStatusDetailsCollectionResponse instantiates a new AssignmentFilterEvaluationStatusDetailsCollectionResponse and sets the default values.
func NewAssignmentFilterEvaluationStatusDetailsCollectionResponse()(*AssignmentFilterEvaluationStatusDetailsCollectionResponse) {
    m := &AssignmentFilterEvaluationStatusDetailsCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateAssignmentFilterEvaluationStatusDetailsCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAssignmentFilterEvaluationStatusDetailsCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterEvaluationStatusDetailsCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AssignmentFilterEvaluationStatusDetailsCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignmentFilterEvaluationStatusDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AssignmentFilterEvaluationStatusDetailsable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []AssignmentFilterEvaluationStatusDetailsable when successful
func (m *AssignmentFilterEvaluationStatusDetailsCollectionResponse) GetValue()([]AssignmentFilterEvaluationStatusDetailsable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AssignmentFilterEvaluationStatusDetailsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AssignmentFilterEvaluationStatusDetailsCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *AssignmentFilterEvaluationStatusDetailsCollectionResponse) SetValue(value []AssignmentFilterEvaluationStatusDetailsable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type AssignmentFilterEvaluationStatusDetailsCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]AssignmentFilterEvaluationStatusDetailsable)
    SetValue(value []AssignmentFilterEvaluationStatusDetailsable)()
}
