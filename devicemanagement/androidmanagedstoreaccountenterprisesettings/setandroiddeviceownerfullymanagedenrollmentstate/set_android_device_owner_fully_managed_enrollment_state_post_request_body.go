package setandroiddeviceownerfullymanagedenrollmentstate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody provides operations to call the setAndroidDeviceOwnerFullyManagedEnrollmentState method.
type SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enabled property
    enabled *bool
}
// NewSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody instantiates a new setAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody and sets the default values.
func NewSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody()(*SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) {
    m := &SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. The enabled property
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnabled)
    return res
}
// Serialize serializes information the current object
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
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
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) SetEnabled(value *bool)() {
    m.enabled = value
}
