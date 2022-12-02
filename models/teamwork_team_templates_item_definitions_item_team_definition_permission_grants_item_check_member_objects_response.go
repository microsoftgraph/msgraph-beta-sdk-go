package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponse provides operations to call the checkMemberObjects method.
type TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []string
}
// NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponse instantiates a new TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponse and sets the default values.
func NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponse()(*TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponse) {
    m := &TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponse) GetValue()([]string) {
    return m.value
}
// Serialize serializes information the current object
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TeamworkTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsResponse) SetValue(value []string)() {
    m.value = value
}
