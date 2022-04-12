package activateserviceplan

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActivateServicePlanRequestBody provides operations to call the activateServicePlan method.
type ActivateServicePlanRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The servicePlanId property
    servicePlanId *string
    // The skuId property
    skuId *string
}
// NewActivateServicePlanRequestBody instantiates a new activateServicePlanRequestBody and sets the default values.
func NewActivateServicePlanRequestBody()(*ActivateServicePlanRequestBody) {
    m := &ActivateServicePlanRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateActivateServicePlanRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActivateServicePlanRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActivateServicePlanRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActivateServicePlanRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActivateServicePlanRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ActivateServicePlanRequestBody) GetServicePlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanId
    }
}
// GetSkuId gets the skuId property value. The skuId property
func (m *ActivateServicePlanRequestBody) GetSkuId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuId
    }
}
// Serialize serializes information the current object
func (m *ActivateServicePlanRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ActivateServicePlanRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetServicePlanId sets the servicePlanId property value. The servicePlanId property
func (m *ActivateServicePlanRequestBody) SetServicePlanId(value *string)() {
    if m != nil {
        m.servicePlanId = value
    }
}
// SetSkuId sets the skuId property value. The skuId property
func (m *ActivateServicePlanRequestBody) SetSkuId(value *string)() {
    if m != nil {
        m.skuId = value
    }
}
