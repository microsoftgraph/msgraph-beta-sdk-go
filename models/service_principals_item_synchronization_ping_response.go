package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalsItemSynchronizationPingResponse provides operations to call the Ping method.
type ServicePrincipalsItemSynchronizationPingResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The value property
    value *string
}
// NewServicePrincipalsItemSynchronizationPingResponse instantiates a new ServicePrincipalsItemSynchronizationPingResponse and sets the default values.
func NewServicePrincipalsItemSynchronizationPingResponse()(*ServicePrincipalsItemSynchronizationPingResponse) {
    m := &ServicePrincipalsItemSynchronizationPingResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateServicePrincipalsItemSynchronizationPingResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalsItemSynchronizationPingResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalsItemSynchronizationPingResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServicePrincipalsItemSynchronizationPingResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalsItemSynchronizationPingResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ServicePrincipalsItemSynchronizationPingResponse) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ServicePrincipalsItemSynchronizationPingResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *ServicePrincipalsItemSynchronizationPingResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *ServicePrincipalsItemSynchronizationPingResponse) SetValue(value *string)() {
    m.value = value
}
