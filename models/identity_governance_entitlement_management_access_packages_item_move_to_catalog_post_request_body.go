package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody provides operations to call the moveToCatalog method.
type IdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The catalogId property
    catalogId *string
}
// NewIdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody instantiates a new IdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody()(*IdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody) {
    m := &IdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCatalogId gets the catalogId property value. The catalogId property
func (m *IdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody) GetCatalogId()(*string) {
    return m.catalogId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["catalogId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *IdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("catalogId", m.GetCatalogId())
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
func (m *IdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCatalogId sets the catalogId property value. The catalogId property
func (m *IdentityGovernanceEntitlementManagementAccessPackagesItemMoveToCatalogPostRequestBody) SetCatalogId(value *string)() {
    m.catalogId = value
}
