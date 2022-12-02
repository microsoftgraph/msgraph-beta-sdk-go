package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementPolicySetsItemUpdatePostRequestBody provides operations to call the update method.
type DeviceAppManagementPolicySetsItemUpdatePostRequestBody struct {
    // The addedPolicySetItems property
    addedPolicySetItems []PolicySetItemable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignments property
    assignments []PolicySetAssignmentable
    // The deletedPolicySetItems property
    deletedPolicySetItems []string
    // The updatedPolicySetItems property
    updatedPolicySetItems []PolicySetItemable
}
// NewDeviceAppManagementPolicySetsItemUpdatePostRequestBody instantiates a new DeviceAppManagementPolicySetsItemUpdatePostRequestBody and sets the default values.
func NewDeviceAppManagementPolicySetsItemUpdatePostRequestBody()(*DeviceAppManagementPolicySetsItemUpdatePostRequestBody) {
    m := &DeviceAppManagementPolicySetsItemUpdatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceAppManagementPolicySetsItemUpdatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementPolicySetsItemUpdatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagementPolicySetsItemUpdatePostRequestBody(), nil
}
// GetAddedPolicySetItems gets the addedPolicySetItems property value. The addedPolicySetItems property
func (m *DeviceAppManagementPolicySetsItemUpdatePostRequestBody) GetAddedPolicySetItems()([]PolicySetItemable) {
    return m.addedPolicySetItems
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAppManagementPolicySetsItemUpdatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignments gets the assignments property value. The assignments property
func (m *DeviceAppManagementPolicySetsItemUpdatePostRequestBody) GetAssignments()([]PolicySetAssignmentable) {
    return m.assignments
}
// GetDeletedPolicySetItems gets the deletedPolicySetItems property value. The deletedPolicySetItems property
func (m *DeviceAppManagementPolicySetsItemUpdatePostRequestBody) GetDeletedPolicySetItems()([]string) {
    return m.deletedPolicySetItems
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementPolicySetsItemUpdatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addedPolicySetItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePolicySetItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PolicySetItemable, len(val))
            for i, v := range val {
                res[i] = v.(PolicySetItemable)
            }
            m.SetAddedPolicySetItems(res)
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePolicySetAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PolicySetAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(PolicySetAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["deletedPolicySetItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDeletedPolicySetItems(res)
        }
        return nil
    }
    res["updatedPolicySetItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePolicySetItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PolicySetItemable, len(val))
            for i, v := range val {
                res[i] = v.(PolicySetItemable)
            }
            m.SetUpdatedPolicySetItems(res)
        }
        return nil
    }
    return res
}
// GetUpdatedPolicySetItems gets the updatedPolicySetItems property value. The updatedPolicySetItems property
func (m *DeviceAppManagementPolicySetsItemUpdatePostRequestBody) GetUpdatedPolicySetItems()([]PolicySetItemable) {
    return m.updatedPolicySetItems
}
// Serialize serializes information the current object
func (m *DeviceAppManagementPolicySetsItemUpdatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAddedPolicySetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddedPolicySetItems()))
        for i, v := range m.GetAddedPolicySetItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("addedPolicySetItems", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedPolicySetItems() != nil {
        err := writer.WriteCollectionOfStringValues("deletedPolicySetItems", m.GetDeletedPolicySetItems())
        if err != nil {
            return err
        }
    }
    if m.GetUpdatedPolicySetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUpdatedPolicySetItems()))
        for i, v := range m.GetUpdatedPolicySetItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("updatedPolicySetItems", cast)
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
// SetAddedPolicySetItems sets the addedPolicySetItems property value. The addedPolicySetItems property
func (m *DeviceAppManagementPolicySetsItemUpdatePostRequestBody) SetAddedPolicySetItems(value []PolicySetItemable)() {
    m.addedPolicySetItems = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAppManagementPolicySetsItemUpdatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignments sets the assignments property value. The assignments property
func (m *DeviceAppManagementPolicySetsItemUpdatePostRequestBody) SetAssignments(value []PolicySetAssignmentable)() {
    m.assignments = value
}
// SetDeletedPolicySetItems sets the deletedPolicySetItems property value. The deletedPolicySetItems property
func (m *DeviceAppManagementPolicySetsItemUpdatePostRequestBody) SetDeletedPolicySetItems(value []string)() {
    m.deletedPolicySetItems = value
}
// SetUpdatedPolicySetItems sets the updatedPolicySetItems property value. The updatedPolicySetItems property
func (m *DeviceAppManagementPolicySetsItemUpdatePostRequestBody) SetUpdatedPolicySetItems(value []PolicySetItemable)() {
    m.updatedPolicySetItems = value
}
