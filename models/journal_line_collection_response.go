package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JournalLineCollectionResponse 
type JournalLineCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewJournalLineCollectionResponse instantiates a new journalLineCollectionResponse and sets the default values.
func NewJournalLineCollectionResponse()(*JournalLineCollectionResponse) {
    m := &JournalLineCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateJournalLineCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJournalLineCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJournalLineCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JournalLineCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJournalLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JournalLineable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(JournalLineable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *JournalLineCollectionResponse) GetValue()([]JournalLineable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]JournalLineable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *JournalLineCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *JournalLineCollectionResponse) SetValue(value []JournalLineable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// JournalLineCollectionResponseable 
type JournalLineCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]JournalLineable)
    SetValue(value []JournalLineable)()
}
