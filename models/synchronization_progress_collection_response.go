package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationProgressCollectionResponse 
type SynchronizationProgressCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewSynchronizationProgressCollectionResponse instantiates a new SynchronizationProgressCollectionResponse and sets the default values.
func NewSynchronizationProgressCollectionResponse()(*SynchronizationProgressCollectionResponse) {
    m := &SynchronizationProgressCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSynchronizationProgressCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationProgressCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationProgressCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationProgressCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationProgressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationProgressable, len(val))
            for i, v := range val {
                res[i] = v.(SynchronizationProgressable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *SynchronizationProgressCollectionResponse) GetValue()([]SynchronizationProgressable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SynchronizationProgressable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SynchronizationProgressCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SynchronizationProgressCollectionResponse) SetValue(value []SynchronizationProgressable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// SynchronizationProgressCollectionResponseable 
type SynchronizationProgressCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]SynchronizationProgressable)
    SetValue(value []SynchronizationProgressable)()
}
