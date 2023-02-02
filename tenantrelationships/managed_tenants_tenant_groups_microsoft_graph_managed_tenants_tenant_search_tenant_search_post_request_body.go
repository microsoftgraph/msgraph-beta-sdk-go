package tenantrelationships

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody 
type ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The tenantId property
    tenantId *string
}
// NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody instantiates a new ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody and sets the default values.
func NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody()(*ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody) {
    m := &ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
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
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBody) SetTenantId(value *string)() {
    m.tenantId = value
}
