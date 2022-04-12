package setandroiddeviceownerfullymanagedenrollmentstate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody provides operations to call the setAndroidDeviceOwnerFullyManagedEnrollmentState method.
type SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enabled property
    enabled *bool
}
// NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody instantiates a new setAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody and sets the default values.
func NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody()(*SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) {
    m := &SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnabled gets the enabled property value. The enabled property
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnabled sets the enabled property value. The enabled property
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) SetEnabled(value *bool)() {
    if m != nil {
        m.enabled = value
    }
}
