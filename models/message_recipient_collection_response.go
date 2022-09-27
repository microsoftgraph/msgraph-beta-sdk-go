package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageRecipientCollectionResponse 
type MessageRecipientCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []MessageRecipientable
}
// NewMessageRecipientCollectionResponse instantiates a new MessageRecipientCollectionResponse and sets the default values.
func NewMessageRecipientCollectionResponse()(*MessageRecipientCollectionResponse) {
    m := &MessageRecipientCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMessageRecipientCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageRecipientCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageRecipientCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageRecipientCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMessageRecipientFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *MessageRecipientCollectionResponse) GetValue()([]MessageRecipientable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MessageRecipientCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *MessageRecipientCollectionResponse) SetValue(value []MessageRecipientable)() {
    m.value = value
}
