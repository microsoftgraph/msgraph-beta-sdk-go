package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody 
type TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The groupIds property
    groupIds []string
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody()(*TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupIds(res)
        }
        return nil
    }
    return res
}
// GetGroupIds gets the groupIds property value. The groupIds property
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody) GetGroupIds()([]string) {
    return m.groupIds
}
// Serialize serializes information the current object
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGroupIds() != nil {
        err := writer.WriteCollectionOfStringValues("groupIds", m.GetGroupIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroupIds sets the groupIds property value. The groupIds property
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody) SetGroupIds(value []string)() {
    m.groupIds = value
}
