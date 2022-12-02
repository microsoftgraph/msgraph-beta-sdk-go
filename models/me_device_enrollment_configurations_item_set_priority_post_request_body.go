package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody provides operations to call the setPriority method.
type MeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The priority property
    priority *int32
}
// NewMeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody instantiates a new MeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody and sets the default values.
func NewMeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody()(*MeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody) {
    m := &MeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    return res
}
// GetPriority gets the priority property value. The priority property
func (m *MeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody) GetPriority()(*int32) {
    return m.priority
}
// Serialize serializes information the current object
func (m *MeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
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
func (m *MeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPriority sets the priority property value. The priority property
func (m *MeDeviceEnrollmentConfigurationsItemSetPriorityPostRequestBody) SetPriority(value *int32)() {
    m.priority = value
}
