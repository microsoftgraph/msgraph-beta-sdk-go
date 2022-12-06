package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebPart 
type WebPart struct {
    Entity
    // The data property
    data SitePageDataable
    // The type property
    type_escaped *string
}
// NewWebPart instantiates a new webPart and sets the default values.
func NewWebPart()(*WebPart) {
    m := &WebPart{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWebPartFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebPartFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebPart(), nil
}
// GetData gets the data property value. The data property
func (m *WebPart) GetData()(SitePageDataable) {
    return m.data
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebPart) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["data"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSitePageDataFromDiscriminatorValue , m.SetData)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetType)
    return res
}
// GetType gets the type property value. The type property
func (m *WebPart) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *WebPart) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("data", m.GetData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetData sets the data property value. The data property
func (m *WebPart) SetData(value SitePageDataable)() {
    m.data = value
}
// SetType sets the type property value. The type property
func (m *WebPart) SetType(value *string)() {
    m.type_escaped = value
}
