package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerProgressTaskBoardTaskFormatCollectionResponse 
type PlannerProgressTaskBoardTaskFormatCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewPlannerProgressTaskBoardTaskFormatCollectionResponse instantiates a new PlannerProgressTaskBoardTaskFormatCollectionResponse and sets the default values.
func NewPlannerProgressTaskBoardTaskFormatCollectionResponse()(*PlannerProgressTaskBoardTaskFormatCollectionResponse) {
    m := &PlannerProgressTaskBoardTaskFormatCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreatePlannerProgressTaskBoardTaskFormatCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerProgressTaskBoardTaskFormatCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerProgressTaskBoardTaskFormatCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerProgressTaskBoardTaskFormatCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerProgressTaskBoardTaskFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerProgressTaskBoardTaskFormatable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerProgressTaskBoardTaskFormatable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *PlannerProgressTaskBoardTaskFormatCollectionResponse) GetValue()([]PlannerProgressTaskBoardTaskFormatable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PlannerProgressTaskBoardTaskFormatable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerProgressTaskBoardTaskFormatCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PlannerProgressTaskBoardTaskFormatCollectionResponse) SetValue(value []PlannerProgressTaskBoardTaskFormatable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// PlannerProgressTaskBoardTaskFormatCollectionResponseable 
type PlannerProgressTaskBoardTaskFormatCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]PlannerProgressTaskBoardTaskFormatable)
    SetValue(value []PlannerProgressTaskBoardTaskFormatable)()
}
