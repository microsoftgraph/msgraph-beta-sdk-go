package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody provides operations to call the provisionOnDemand method.
type ServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The parameters property
    parameters []SynchronizationJobApplicationParametersable
}
// NewServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody instantiates a new ServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody and sets the default values.
func NewServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody()(*ServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody) {
    m := &ServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationJobApplicationParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationJobApplicationParametersable, len(val))
            for i, v := range val {
                res[i] = v.(SynchronizationJobApplicationParametersable)
            }
            m.SetParameters(res)
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
func (m *ServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody) GetParameters()([]SynchronizationJobApplicationParametersable) {
    return m.parameters
}
// Serialize serializes information the current object
func (m *ServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParameters()))
        for i, v := range m.GetParameters() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("parameters", cast)
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
func (m *ServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *ServicePrincipalsItemSynchronizationJobsItemProvisionOnDemandPostRequestBody) SetParameters(value []SynchronizationJobApplicationParametersable)() {
    m.parameters = value
}
