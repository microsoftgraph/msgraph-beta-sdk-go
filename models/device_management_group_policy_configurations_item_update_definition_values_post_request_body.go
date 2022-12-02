package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody provides operations to call the updateDefinitionValues method.
type DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody struct {
    // The added property
    added []GroupPolicyDefinitionValueable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deletedIds property
    deletedIds []string
    // The updated property
    updated []GroupPolicyDefinitionValueable
}
// NewDeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody instantiates a new DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody and sets the default values.
func NewDeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody()(*DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) {
    m := &DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody(), nil
}
// GetAdded gets the added property value. The added property
func (m *DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) GetAdded()([]GroupPolicyDefinitionValueable) {
    return m.added
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeletedIds gets the deletedIds property value. The deletedIds property
func (m *DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) GetDeletedIds()([]string) {
    return m.deletedIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["added"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyDefinitionValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyDefinitionValueable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyDefinitionValueable)
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
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyDefinitionValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyDefinitionValueable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyDefinitionValueable)
            }
            m.SetUpdated(res)
        }
        return nil
    }
    return res
}
// GetUpdated gets the updated property value. The updated property
func (m *DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) GetUpdated()([]GroupPolicyDefinitionValueable) {
    return m.updated
}
// Serialize serializes information the current object
func (m *DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) SetAdded(value []GroupPolicyDefinitionValueable)() {
    m.added = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeletedIds sets the deletedIds property value. The deletedIds property
func (m *DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) SetDeletedIds(value []string)() {
    m.deletedIds = value
}
// SetUpdated sets the updated property value. The updated property
func (m *DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesPostRequestBody) SetUpdated(value []GroupPolicyDefinitionValueable)() {
    m.updated = value
}
