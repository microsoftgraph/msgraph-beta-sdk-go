package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody 
type WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody struct {
    // The addExclusions property
    addExclusions []string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The addMembers property
    addMembers []string
    // The memberEntityType property
    memberEntityType *string
    // The removeExclusions property
    removeExclusions []string
    // The removeMembers property
    removeMembers []string
}
// NewWindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody instantiates a new WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody()(*WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) {
    m := &WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody(), nil
}
// GetAddExclusions gets the addExclusions property value. The addExclusions property
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetAddExclusions()([]string) {
    return m.addExclusions
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddMembers gets the addMembers property value. The addMembers property
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetAddMembers()([]string) {
    return m.addMembers
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetMemberEntityType()(*string) {
    return m.memberEntityType
}
// GetRemoveExclusions gets the removeExclusions property value. The removeExclusions property
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetRemoveExclusions()([]string) {
    return m.removeExclusions
}
// GetRemoveMembers gets the removeMembers property value. The removeMembers property
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) GetRemoveMembers()([]string) {
    return m.removeMembers
}
// Serialize serializes information the current object
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) SetAddExclusions(value []string)() {
    m.addExclusions = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddMembers sets the addMembers property value. The addMembers property
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) SetAddMembers(value []string)() {
    m.addMembers = value
}
// SetMemberEntityType sets the memberEntityType property value. The memberEntityType property
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) SetMemberEntityType(value *string)() {
    m.memberEntityType = value
}
// SetRemoveExclusions sets the removeExclusions property value. The removeExclusions property
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) SetRemoveExclusions(value []string)() {
    m.removeExclusions = value
}
// SetRemoveMembers sets the removeMembers property value. The removeMembers property
func (m *WindowsUpdatesDeploymentsItemAudienceMicrosoftGraphWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBody) SetRemoveMembers(value []string)() {
    m.removeMembers = value
}
