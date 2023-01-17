package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody 
type GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The roleScopeTagIds property
    roleScopeTagIds []string
}
// NewGroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody instantiates a new GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody and sets the default values.
func NewGroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody()(*GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody) {
    m := &GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateGroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    return res
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. The roleScopeTagIds property
func (m *GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// Serialize serializes information the current object
func (m *GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRoleScopeTagIds() != nil {
        err := writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
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
func (m *GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. The roleScopeTagIds property
func (m *GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBody) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
