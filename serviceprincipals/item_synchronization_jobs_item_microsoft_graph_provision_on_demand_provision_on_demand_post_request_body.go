package serviceprincipals

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody 
type ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parameters property
    parameters []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobApplicationParametersable
}
// NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody instantiates a new ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody and sets the default values.
func NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody()(*ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody) {
    m := &ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSynchronizationJobApplicationParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobApplicationParametersable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobApplicationParametersable)
            }
            m.SetParameters(res)
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody) GetParameters()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobApplicationParametersable) {
    return m.parameters
}
// Serialize serializes information the current object
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody) SetParameters(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobApplicationParametersable)() {
    m.parameters = value
}
