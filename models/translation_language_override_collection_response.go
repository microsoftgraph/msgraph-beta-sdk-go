package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TranslationLanguageOverrideCollectionResponse 
type TranslationLanguageOverrideCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewTranslationLanguageOverrideCollectionResponse instantiates a new TranslationLanguageOverrideCollectionResponse and sets the default values.
func NewTranslationLanguageOverrideCollectionResponse()(*TranslationLanguageOverrideCollectionResponse) {
    m := &TranslationLanguageOverrideCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateTranslationLanguageOverrideCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTranslationLanguageOverrideCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTranslationLanguageOverrideCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TranslationLanguageOverrideCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTranslationLanguageOverrideFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TranslationLanguageOverrideable, len(val))
            for i, v := range val {
                res[i] = v.(TranslationLanguageOverrideable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *TranslationLanguageOverrideCollectionResponse) GetValue()([]TranslationLanguageOverrideable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TranslationLanguageOverrideable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TranslationLanguageOverrideCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TranslationLanguageOverrideCollectionResponse) SetValue(value []TranslationLanguageOverrideable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// TranslationLanguageOverrideCollectionResponseable 
type TranslationLanguageOverrideCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]TranslationLanguageOverrideable)
    SetValue(value []TranslationLanguageOverrideable)()
}
