package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody provides operations to call the restart method.
type ServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The criteria property
    criteria SynchronizationJobRestartCriteriaable
}
// NewServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody instantiates a new ServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody and sets the default values.
func NewServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody()(*ServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody) {
    m := &ServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCriteria gets the criteria property value. The criteria property
func (m *ServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody) GetCriteria()(SynchronizationJobRestartCriteriaable) {
    return m.criteria
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["criteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationJobRestartCriteriaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriteria(val.(SynchronizationJobRestartCriteriaable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("criteria", m.GetCriteria())
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
func (m *ServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCriteria sets the criteria property value. The criteria property
func (m *ServicePrincipalsItemSynchronizationJobsItemRestartPostRequestBody) SetCriteria(value SynchronizationJobRestartCriteriaable)() {
    m.criteria = value
}
