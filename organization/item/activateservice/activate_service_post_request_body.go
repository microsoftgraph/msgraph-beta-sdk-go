package activateservice

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActivateServicePostRequestBody provides operations to call the activateService method.
type ActivateServicePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The service property
    service *string
    // The servicePlanId property
    servicePlanId *string
    // The skuId property
    skuId *string
}
// NewActivateServicePostRequestBody instantiates a new activateServicePostRequestBody and sets the default values.
func NewActivateServicePostRequestBody()(*ActivateServicePostRequestBody) {
    m := &ActivateServicePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateActivateServicePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActivateServicePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActivateServicePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActivateServicePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActivateServicePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["service"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetService)
    res["servicePlanId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetServicePlanId)
    res["skuId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSkuId)
    return res
}
// GetService gets the service property value. The service property
func (m *ActivateServicePostRequestBody) GetService()(*string) {
    return m.service
}
// GetServicePlanId gets the servicePlanId property value. The servicePlanId property
func (m *ActivateServicePostRequestBody) GetServicePlanId()(*string) {
    return m.servicePlanId
}
// GetSkuId gets the skuId property value. The skuId property
func (m *ActivateServicePostRequestBody) GetSkuId()(*string) {
    return m.skuId
}
// Serialize serializes information the current object
func (m *ActivateServicePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePlanId", m.GetServicePlanId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("skuId", m.GetSkuId())
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
func (m *ActivateServicePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetService sets the service property value. The service property
func (m *ActivateServicePostRequestBody) SetService(value *string)() {
    m.service = value
}
// SetServicePlanId sets the servicePlanId property value. The servicePlanId property
func (m *ActivateServicePostRequestBody) SetServicePlanId(value *string)() {
    m.servicePlanId = value
}
// SetSkuId sets the skuId property value. The skuId property
func (m *ActivateServicePostRequestBody) SetSkuId(value *string)() {
    m.skuId = value
}
