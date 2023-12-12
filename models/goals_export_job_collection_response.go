package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GoalsExportJobCollectionResponse 
type GoalsExportJobCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewGoalsExportJobCollectionResponse instantiates a new goalsExportJobCollectionResponse and sets the default values.
func NewGoalsExportJobCollectionResponse()(*GoalsExportJobCollectionResponse) {
    m := &GoalsExportJobCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateGoalsExportJobCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGoalsExportJobCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGoalsExportJobCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GoalsExportJobCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGoalsExportJobFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GoalsExportJobable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GoalsExportJobable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *GoalsExportJobCollectionResponse) GetValue()([]GoalsExportJobable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GoalsExportJobable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GoalsExportJobCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GoalsExportJobCollectionResponse) SetValue(value []GoalsExportJobable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// GoalsExportJobCollectionResponseable 
type GoalsExportJobCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]GoalsExportJobable)
    SetValue(value []GoalsExportJobable)()
}
