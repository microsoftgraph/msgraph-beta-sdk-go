package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomAppScopeAttributesDictionary 
type CustomAppScopeAttributesDictionary struct {
    Dictionary
    // The OdataType property
    OdataType *string
}
// NewCustomAppScopeAttributesDictionary instantiates a new customAppScopeAttributesDictionary and sets the default values.
func NewCustomAppScopeAttributesDictionary()(*CustomAppScopeAttributesDictionary) {
    m := &CustomAppScopeAttributesDictionary{
        Dictionary: *NewDictionary(),
    }
    return m
}
// CreateCustomAppScopeAttributesDictionaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomAppScopeAttributesDictionaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomAppScopeAttributesDictionary(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomAppScopeAttributesDictionary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Dictionary.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *CustomAppScopeAttributesDictionary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Dictionary.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// CustomAppScopeAttributesDictionaryable 
type CustomAppScopeAttributesDictionaryable interface {
    Dictionaryable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
