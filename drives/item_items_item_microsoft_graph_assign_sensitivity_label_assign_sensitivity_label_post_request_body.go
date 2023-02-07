package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody 
type ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The assignmentMethod property
    assignmentMethod *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelAssignmentMethod
    // The justificationText property
    justificationText *string
    // The sensitivityLabelId property
    sensitivityLabelId *string
}
// NewItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody instantiates a new ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody and sets the default values.
func NewItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody()(*ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody) {
    m := &ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignmentMethod gets the assignmentMethod property value. The assignmentMethod property
func (m *ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody) GetAssignmentMethod()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelAssignmentMethod) {
    return m.assignmentMethod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseSensitivityLabelAssignmentMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentMethod(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelAssignmentMethod))
        }
        return nil
    }
    res["justificationText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustificationText(val)
        }
        return nil
    }
    res["sensitivityLabelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivityLabelId(val)
        }
        return nil
    }
    return res
}
// GetJustificationText gets the justificationText property value. The justificationText property
func (m *ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody) GetJustificationText()(*string) {
    return m.justificationText
}
// GetSensitivityLabelId gets the sensitivityLabelId property value. The sensitivityLabelId property
func (m *ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody) GetSensitivityLabelId()(*string) {
    return m.sensitivityLabelId
}
// Serialize serializes information the current object
func (m *ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentMethod() != nil {
        cast := (*m.GetAssignmentMethod()).String()
        err := writer.WriteStringValue("assignmentMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("justificationText", m.GetJustificationText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sensitivityLabelId", m.GetSensitivityLabelId())
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
func (m *ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignmentMethod sets the assignmentMethod property value. The assignmentMethod property
func (m *ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody) SetAssignmentMethod(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelAssignmentMethod)() {
    m.assignmentMethod = value
}
// SetJustificationText sets the justificationText property value. The justificationText property
func (m *ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody) SetJustificationText(value *string)() {
    m.justificationText = value
}
// SetSensitivityLabelId sets the sensitivityLabelId property value. The sensitivityLabelId property
func (m *ItemItemsItemMicrosoftGraphAssignSensitivityLabelAssignSensitivityLabelPostRequestBody) SetSensitivityLabelId(value *string)() {
    m.sensitivityLabelId = value
}
