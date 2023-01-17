package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody 
type GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody struct {
    // The added property
    added []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deletedIds property
    deletedIds []string
    // The updated property
    updated []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable
}
// NewGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody instantiates a new GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody and sets the default values.
func NewGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody()(*GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) {
    m := &GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody(), nil
}
// GetAdded gets the added property value. The added property
func (m *GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) GetAdded()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable) {
    return m.added
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeletedIds gets the deletedIds property value. The deletedIds property
func (m *GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) GetDeletedIds()([]string) {
    return m.deletedIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["added"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)
            }
            m.SetAdded(res)
        }
        return nil
    }
    res["deletedIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDeletedIds(res)
        }
        return nil
    }
    res["updated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)
            }
            m.SetUpdated(res)
        }
        return nil
    }
    return res
}
// GetUpdated gets the updated property value. The updated property
func (m *GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) GetUpdated()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable) {
    return m.updated
}
// Serialize serializes information the current object
func (m *GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdded() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdded()))
        for i, v := range m.GetAdded() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("added", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedIds() != nil {
        err := writer.WriteCollectionOfStringValues("deletedIds", m.GetDeletedIds())
        if err != nil {
            return err
        }
    }
    if m.GetUpdated() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUpdated()))
        for i, v := range m.GetUpdated() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("updated", cast)
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
// SetAdded sets the added property value. The added property
func (m *GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) SetAdded(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)() {
    m.added = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeletedIds sets the deletedIds property value. The deletedIds property
func (m *GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) SetDeletedIds(value []string)() {
    m.deletedIds = value
}
// SetUpdated sets the updated property value. The updated property
func (m *GroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) SetUpdated(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable)() {
    m.updated = value
}
