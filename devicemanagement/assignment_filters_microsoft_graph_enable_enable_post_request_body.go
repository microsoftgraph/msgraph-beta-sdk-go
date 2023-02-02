package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody 
type AssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enable property
    enable *bool
}
// NewAssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody instantiates a new AssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody and sets the default values.
func NewAssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody()(*AssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody) {
    m := &AssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAssignmentFiltersMicrosoftGraphEnableEnablePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFiltersMicrosoftGraphEnableEnablePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnable gets the enable property value. The enable property
func (m *AssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody) GetEnable()(*bool) {
    return m.enable
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnable(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enable", m.GetEnable())
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
func (m *AssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnable sets the enable property value. The enable property
func (m *AssignmentFiltersMicrosoftGraphEnableEnablePostRequestBody) SetEnable(value *bool)() {
    m.enable = value
}
