package updateaudiencebyid

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateAudienceByIdRequestBody provides operations to call the updateAudienceById method.
type UpdateAudienceByIdRequestBody struct {
    // The addExclusions property
    addExclusions []string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The addMembers property
    addMembers []string
    // The memberEntityType property
    memberEntityType *string
    // The removeExclusions property
    removeExclusions []string
    // The removeMembers property
    removeMembers []string
}
// NewUpdateAudienceByIdRequestBody instantiates a new updateAudienceByIdRequestBody and sets the default values.
func NewUpdateAudienceByIdRequestBody()(*UpdateAudienceByIdRequestBody) {
    m := &UpdateAudienceByIdRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUpdateAudienceByIdRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateAudienceByIdRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateAudienceByIdRequestBody(), nil
}
// GetAddExclusions gets the addExclusions property value. The addExclusions property
func (m *UpdateAudienceByIdRequestBody) GetAddExclusions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.addExclusions
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateAudienceByIdRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddMembers gets the addMembers property value. The addMembers property
func (m *UpdateAudienceByIdRequestBody) GetAddMembers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.addMembers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateAudienceByIdRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addExclusions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAddExclusions(res)
        }
        return nil
    }
    res["addMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAddMembers(res)
        }
        return nil
    }
    res["memberEntityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberEntityType(val)
        }
        return nil
    }
    res["removeExclusions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRemoveExclusions(res)
        }
        return nil
    }
    res["removeMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRemoveMembers(res)
        }
        return nil
    }
    return res
}
// GetMemberEntityType gets the memberEntityType property value. The memberEntityType property
func (m *UpdateAudienceByIdRequestBody) GetMemberEntityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberEntityType
    }
}
// GetRemoveExclusions gets the removeExclusions property value. The removeExclusions property
func (m *UpdateAudienceByIdRequestBody) GetRemoveExclusions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.removeExclusions
    }
}
// GetRemoveMembers gets the removeMembers property value. The removeMembers property
func (m *UpdateAudienceByIdRequestBody) GetRemoveMembers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.removeMembers
    }
}
// Serialize serializes information the current object
func (m *UpdateAudienceByIdRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAddExclusions() != nil {
        err := writer.WriteCollectionOfStringValues("addExclusions", m.GetAddExclusions())
        if err != nil {
            return err
        }
    }
    if m.GetAddMembers() != nil {
        err := writer.WriteCollectionOfStringValues("addMembers", m.GetAddMembers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("memberEntityType", m.GetMemberEntityType())
        if err != nil {
            return err
        }
    }
    if m.GetRemoveExclusions() != nil {
        err := writer.WriteCollectionOfStringValues("removeExclusions", m.GetRemoveExclusions())
        if err != nil {
            return err
        }
    }
    if m.GetRemoveMembers() != nil {
        err := writer.WriteCollectionOfStringValues("removeMembers", m.GetRemoveMembers())
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
// SetAddExclusions sets the addExclusions property value. The addExclusions property
func (m *UpdateAudienceByIdRequestBody) SetAddExclusions(value []string)() {
    if m != nil {
        m.addExclusions = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateAudienceByIdRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAddMembers sets the addMembers property value. The addMembers property
func (m *UpdateAudienceByIdRequestBody) SetAddMembers(value []string)() {
    if m != nil {
        m.addMembers = value
    }
}
// SetMemberEntityType sets the memberEntityType property value. The memberEntityType property
func (m *UpdateAudienceByIdRequestBody) SetMemberEntityType(value *string)() {
    if m != nil {
        m.memberEntityType = value
    }
}
// SetRemoveExclusions sets the removeExclusions property value. The removeExclusions property
func (m *UpdateAudienceByIdRequestBody) SetRemoveExclusions(value []string)() {
    if m != nil {
        m.removeExclusions = value
    }
}
// SetRemoveMembers sets the removeMembers property value. The removeMembers property
func (m *UpdateAudienceByIdRequestBody) SetRemoveMembers(value []string)() {
    if m != nil {
        m.removeMembers = value
    }
}
