package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmailActivityStatisticsCollectionResponse 
type EmailActivityStatisticsCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewEmailActivityStatisticsCollectionResponse instantiates a new EmailActivityStatisticsCollectionResponse and sets the default values.
func NewEmailActivityStatisticsCollectionResponse()(*EmailActivityStatisticsCollectionResponse) {
    m := &EmailActivityStatisticsCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateEmailActivityStatisticsCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailActivityStatisticsCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmailActivityStatisticsCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailActivityStatisticsCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmailActivityStatisticsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmailActivityStatisticsable, len(val))
            for i, v := range val {
                res[i] = v.(EmailActivityStatisticsable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *EmailActivityStatisticsCollectionResponse) GetValue()([]EmailActivityStatisticsable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EmailActivityStatisticsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EmailActivityStatisticsCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *EmailActivityStatisticsCollectionResponse) SetValue(value []EmailActivityStatisticsable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// EmailActivityStatisticsCollectionResponseable 
type EmailActivityStatisticsCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]EmailActivityStatisticsable)
    SetValue(value []EmailActivityStatisticsable)()
}
