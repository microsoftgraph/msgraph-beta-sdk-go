package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleAppListItem represents an app in the list of managed Apple applications
type AppleAppListItem struct {
    AppListItem
}
// NewAppleAppListItem instantiates a new appleAppListItem and sets the default values.
func NewAppleAppListItem()(*AppleAppListItem) {
    m := &AppleAppListItem{
        AppListItem: *NewAppListItem(),
    }
    return m
}
// CreateAppleAppListItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleAppListItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppleAppListItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleAppListItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AppListItem.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AppleAppListItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AppListItem.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// AppleAppListItemable 
type AppleAppListItemable interface {
    AppListItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
