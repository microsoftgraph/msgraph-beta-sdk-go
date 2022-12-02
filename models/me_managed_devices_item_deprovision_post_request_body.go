package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeManagedDevicesItemDeprovisionPostRequestBody provides operations to call the deprovision method.
type MeManagedDevicesItemDeprovisionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deprovisionReason property
    deprovisionReason *string
}
// NewMeManagedDevicesItemDeprovisionPostRequestBody instantiates a new MeManagedDevicesItemDeprovisionPostRequestBody and sets the default values.
func NewMeManagedDevicesItemDeprovisionPostRequestBody()(*MeManagedDevicesItemDeprovisionPostRequestBody) {
    m := &MeManagedDevicesItemDeprovisionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeManagedDevicesItemDeprovisionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeManagedDevicesItemDeprovisionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeManagedDevicesItemDeprovisionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeManagedDevicesItemDeprovisionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeprovisionReason gets the deprovisionReason property value. The deprovisionReason property
func (m *MeManagedDevicesItemDeprovisionPostRequestBody) GetDeprovisionReason()(*string) {
    return m.deprovisionReason
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeManagedDevicesItemDeprovisionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deprovisionReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeprovisionReason(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MeManagedDevicesItemDeprovisionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deprovisionReason", m.GetDeprovisionReason())
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
func (m *MeManagedDevicesItemDeprovisionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeprovisionReason sets the deprovisionReason property value. The deprovisionReason property
func (m *MeManagedDevicesItemDeprovisionPostRequestBody) SetDeprovisionReason(value *string)() {
    m.deprovisionReason = value
}
