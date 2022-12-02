package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemFindRoomListsResponse provides operations to call the findRoomLists method.
type UsersItemFindRoomListsResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []EmailAddressable
}
// NewUsersItemFindRoomListsResponse instantiates a new UsersItemFindRoomListsResponse and sets the default values.
func NewUsersItemFindRoomListsResponse()(*UsersItemFindRoomListsResponse) {
    m := &UsersItemFindRoomListsResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateUsersItemFindRoomListsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemFindRoomListsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemFindRoomListsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemFindRoomListsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmailAddressable, len(val))
            for i, v := range val {
                res[i] = v.(EmailAddressable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *UsersItemFindRoomListsResponse) GetValue()([]EmailAddressable) {
    return m.value
}
// Serialize serializes information the current object
func (m *UsersItemFindRoomListsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UsersItemFindRoomListsResponse) SetValue(value []EmailAddressable)() {
    m.value = value
}
