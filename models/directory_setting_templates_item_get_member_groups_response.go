package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectorySettingTemplatesItemGetMemberGroupsResponse provides operations to call the getMemberGroups method.
type DirectorySettingTemplatesItemGetMemberGroupsResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []string
}
// NewDirectorySettingTemplatesItemGetMemberGroupsResponse instantiates a new DirectorySettingTemplatesItemGetMemberGroupsResponse and sets the default values.
func NewDirectorySettingTemplatesItemGetMemberGroupsResponse()(*DirectorySettingTemplatesItemGetMemberGroupsResponse) {
    m := &DirectorySettingTemplatesItemGetMemberGroupsResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDirectorySettingTemplatesItemGetMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectorySettingTemplatesItemGetMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectorySettingTemplatesItemGetMemberGroupsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectorySettingTemplatesItemGetMemberGroupsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *DirectorySettingTemplatesItemGetMemberGroupsResponse) GetValue()([]string) {
    return m.value
}
// Serialize serializes information the current object
func (m *DirectorySettingTemplatesItemGetMemberGroupsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        err = writer.WriteCollectionOfStringValues("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *DirectorySettingTemplatesItemGetMemberGroupsResponse) SetValue(value []string)() {
    m.value = value
}
