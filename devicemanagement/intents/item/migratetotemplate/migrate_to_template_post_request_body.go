package migratetotemplate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MigrateToTemplatePostRequestBody provides operations to call the migrateToTemplate method.
type MigrateToTemplatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The newTemplateId property
    newTemplateId *string
    // The preserveCustomValues property
    preserveCustomValues *bool
}
// NewMigrateToTemplatePostRequestBody instantiates a new migrateToTemplatePostRequestBody and sets the default values.
func NewMigrateToTemplatePostRequestBody()(*MigrateToTemplatePostRequestBody) {
    m := &MigrateToTemplatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMigrateToTemplatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMigrateToTemplatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMigrateToTemplatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MigrateToTemplatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MigrateToTemplatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newTemplateId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNewTemplateId)
    res["preserveCustomValues"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPreserveCustomValues)
    return res
}
// GetNewTemplateId gets the newTemplateId property value. The newTemplateId property
func (m *MigrateToTemplatePostRequestBody) GetNewTemplateId()(*string) {
    return m.newTemplateId
}
// GetPreserveCustomValues gets the preserveCustomValues property value. The preserveCustomValues property
func (m *MigrateToTemplatePostRequestBody) GetPreserveCustomValues()(*bool) {
    return m.preserveCustomValues
}
// Serialize serializes information the current object
func (m *MigrateToTemplatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newTemplateId", m.GetNewTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("preserveCustomValues", m.GetPreserveCustomValues())
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
func (m *MigrateToTemplatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNewTemplateId sets the newTemplateId property value. The newTemplateId property
func (m *MigrateToTemplatePostRequestBody) SetNewTemplateId(value *string)() {
    m.newTemplateId = value
}
// SetPreserveCustomValues sets the preserveCustomValues property value. The preserveCustomValues property
func (m *MigrateToTemplatePostRequestBody) SetPreserveCustomValues(value *bool)() {
    m.preserveCustomValues = value
}
