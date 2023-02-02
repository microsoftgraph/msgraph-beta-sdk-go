package tenantrelationships

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody 
type ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The tenantIds property
    tenantIds []string
}
// NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody instantiates a new ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody and sets the default values.
func NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody()(*ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody) {
    m := &ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["tenantIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTenantIds(res)
        }
        return nil
    }
    return res
}
// GetTenantIds gets the tenantIds property value. The tenantIds property
func (m *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody) GetTenantIds()([]string) {
    return m.tenantIds
}
// Serialize serializes information the current object
func (m *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTenantIds() != nil {
        err := writer.WriteCollectionOfStringValues("tenantIds", m.GetTenantIds())
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
func (m *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTenantIds sets the tenantIds property value. The tenantIds property
func (m *ManagedTenantsTenantTagsItemMicrosoftGraphManagedTenantsUnassignTagUnassignTagPostRequestBody) SetTenantIds(value []string)() {
    m.tenantIds = value
}
