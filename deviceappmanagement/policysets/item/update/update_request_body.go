package update

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// UpdateRequestBody provides operations to call the update method.
type UpdateRequestBody struct {
    // The addedPolicySetItems property
    addedPolicySetItems []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetItemable;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The assignments property
    assignments []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetAssignmentable;
    // The deletedPolicySetItems property
    deletedPolicySetItems []string;
    // The updatedPolicySetItems property
    updatedPolicySetItems []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetItemable;
}
// NewUpdateRequestBody instantiates a new updateRequestBody and sets the default values.
func NewUpdateRequestBody()(*UpdateRequestBody) {
    m := &UpdateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUpdateRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateRequestBody(), nil
}
// GetAddedPolicySetItems gets the addedPolicySetItems property value. The addedPolicySetItems property
func (m *UpdateRequestBody) GetAddedPolicySetItems()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetItemable) {
    if m == nil {
        return nil
    } else {
        return m.addedPolicySetItems
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignments gets the assignments property value. The assignments property
func (m *UpdateRequestBody) GetAssignments()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetDeletedPolicySetItems gets the deletedPolicySetItems property value. The deletedPolicySetItems property
func (m *UpdateRequestBody) GetDeletedPolicySetItems()([]string) {
    if m == nil {
        return nil
    } else {
        return m.deletedPolicySetItems
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addedPolicySetItems"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePolicySetItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetItemable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetItemable)
            }
            m.SetAddedPolicySetItems(res)
        }
        return nil
    }
    res["assignments"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePolicySetAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["deletedPolicySetItems"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["updatedPolicySetItems"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePolicySetItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetItemable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetItemable)
            }
            m.SetUpdatedPolicySetItems(res)
        }
        return nil
    }
    return res
}
// GetUpdatedPolicySetItems gets the updatedPolicySetItems property value. The updatedPolicySetItems property
func (m *UpdateRequestBody) GetUpdatedPolicySetItems()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetItemable) {
    if m == nil {
        return nil
    } else {
        return m.updatedPolicySetItems
    }
}
// Serialize serializes information the current object
func (m *UpdateRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UpdateRequestBody) SetAddedPolicySetItems(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetItemable)() {
    if m != nil {
        m.addedPolicySetItems = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignments sets the assignments property value. The assignments property
func (m *UpdateRequestBody) SetAssignments(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetAssignmentable)() {
    if m != nil {
        m.assignments = value
    }
}
// SetDeletedPolicySetItems sets the deletedPolicySetItems property value. The deletedPolicySetItems property
func (m *UpdateRequestBody) SetDeletedPolicySetItems(value []string)() {
    if m != nil {
        m.deletedPolicySetItems = value
    }
}
// SetUpdatedPolicySetItems sets the updatedPolicySetItems property value. The updatedPolicySetItems property
func (m *UpdateRequestBody) SetUpdatedPolicySetItems(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetItemable)() {
    if m != nil {
        m.updatedPolicySetItems = value
    }
}
