package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemActivateServicePlanPostRequestBody provides operations to call the activateServicePlan method.
type UsersItemActivateServicePlanPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The servicePlanId property
    servicePlanId *string
    // The skuId property
    skuId *string
}
// NewUsersItemActivateServicePlanPostRequestBody instantiates a new UsersItemActivateServicePlanPostRequestBody and sets the default values.
func NewUsersItemActivateServicePlanPostRequestBody()(*UsersItemActivateServicePlanPostRequestBody) {
    m := &UsersItemActivateServicePlanPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemActivateServicePlanPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemActivateServicePlanPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemActivateServicePlanPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemActivateServicePlanPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemActivateServicePlanPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["servicePlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanId(val)
        }
        return nil
    }
    res["skuId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuId(val)
        }
        return nil
    }
    return res
}
// GetServicePlanId gets the servicePlanId property value. The servicePlanId property
func (m *UsersItemActivateServicePlanPostRequestBody) GetServicePlanId()(*string) {
    return m.servicePlanId
}
// GetSkuId gets the skuId property value. The skuId property
func (m *UsersItemActivateServicePlanPostRequestBody) GetSkuId()(*string) {
    return m.skuId
}
// Serialize serializes information the current object
func (m *UsersItemActivateServicePlanPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UsersItemActivateServicePlanPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetServicePlanId sets the servicePlanId property value. The servicePlanId property
func (m *UsersItemActivateServicePlanPostRequestBody) SetServicePlanId(value *string)() {
    m.servicePlanId = value
}
// SetSkuId sets the skuId property value. The skuId property
func (m *UsersItemActivateServicePlanPostRequestBody) SetSkuId(value *string)() {
    m.skuId = value
}
