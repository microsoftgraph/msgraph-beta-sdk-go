package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProjectParticipationCollectionResponse 
type ProjectParticipationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewProjectParticipationCollectionResponse instantiates a new ProjectParticipationCollectionResponse and sets the default values.
func NewProjectParticipationCollectionResponse()(*ProjectParticipationCollectionResponse) {
    m := &ProjectParticipationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateProjectParticipationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProjectParticipationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProjectParticipationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProjectParticipationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProjectParticipationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProjectParticipationable, len(val))
            for i, v := range val {
                res[i] = v.(ProjectParticipationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ProjectParticipationCollectionResponse) GetValue()([]ProjectParticipationable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ProjectParticipationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ProjectParticipationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ProjectParticipationCollectionResponse) SetValue(value []ProjectParticipationable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// ProjectParticipationCollectionResponseable 
type ProjectParticipationCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ProjectParticipationable)
    SetValue(value []ProjectParticipationable)()
}
