package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationItemActivateServicePostRequestBody provides operations to call the activateService method.
type OrganizationItemActivateServicePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The service property
    service *string
    // The servicePlanId property
    servicePlanId *string
    // The skuId property
    skuId *string
}
// NewOrganizationItemActivateServicePostRequestBody instantiates a new OrganizationItemActivateServicePostRequestBody and sets the default values.
func NewOrganizationItemActivateServicePostRequestBody()(*OrganizationItemActivateServicePostRequestBody) {
    m := &OrganizationItemActivateServicePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOrganizationItemActivateServicePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationItemActivateServicePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationItemActivateServicePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationItemActivateServicePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationItemActivateServicePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
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
// GetService gets the service property value. The service property
func (m *OrganizationItemActivateServicePostRequestBody) GetService()(*string) {
    return m.service
}
// GetServicePlanId gets the servicePlanId property value. The servicePlanId property
func (m *OrganizationItemActivateServicePostRequestBody) GetServicePlanId()(*string) {
    return m.servicePlanId
}
// GetSkuId gets the skuId property value. The skuId property
func (m *OrganizationItemActivateServicePostRequestBody) GetSkuId()(*string) {
    return m.skuId
}
// Serialize serializes information the current object
func (m *OrganizationItemActivateServicePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OrganizationItemActivateServicePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetService sets the service property value. The service property
func (m *OrganizationItemActivateServicePostRequestBody) SetService(value *string)() {
    m.service = value
}
// SetServicePlanId sets the servicePlanId property value. The servicePlanId property
func (m *OrganizationItemActivateServicePostRequestBody) SetServicePlanId(value *string)() {
    m.servicePlanId = value
}
// SetSkuId sets the skuId property value. The skuId property
func (m *OrganizationItemActivateServicePostRequestBody) SetSkuId(value *string)() {
    m.skuId = value
}
