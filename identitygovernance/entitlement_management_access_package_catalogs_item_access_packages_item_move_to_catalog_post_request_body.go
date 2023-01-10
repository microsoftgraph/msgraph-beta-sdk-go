package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody 
type EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The catalogId property
    catalogId *string
}
// NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody instantiates a new EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody and sets the default values.
func NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody()(*EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody) {
    m := &EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCatalogId gets the catalogId property value. The catalogId property
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody) GetCatalogId()(*string) {
    return m.catalogId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCatalogId sets the catalogId property value. The catalogId property
func (m *EntitlementManagementAccessPackageCatalogsItemAccessPackagesItemMoveToCatalogPostRequestBody) SetCatalogId(value *string)() {
    m.catalogId = value
}
