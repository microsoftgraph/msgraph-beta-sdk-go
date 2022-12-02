package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityProtectionRiskyServicePrincipalsDismissPostRequestBody provides operations to call the dismiss method.
type IdentityProtectionRiskyServicePrincipalsDismissPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The servicePrincipalIds property
    servicePrincipalIds []string
}
// NewIdentityProtectionRiskyServicePrincipalsDismissPostRequestBody instantiates a new IdentityProtectionRiskyServicePrincipalsDismissPostRequestBody and sets the default values.
func NewIdentityProtectionRiskyServicePrincipalsDismissPostRequestBody()(*IdentityProtectionRiskyServicePrincipalsDismissPostRequestBody) {
    m := &IdentityProtectionRiskyServicePrincipalsDismissPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityProtectionRiskyServicePrincipalsDismissPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityProtectionRiskyServicePrincipalsDismissPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityProtectionRiskyServicePrincipalsDismissPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityProtectionRiskyServicePrincipalsDismissPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityProtectionRiskyServicePrincipalsDismissPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["servicePrincipalIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetServicePrincipalIds(res)
        }
        return nil
    }
    return res
}
// GetServicePrincipalIds gets the servicePrincipalIds property value. The servicePrincipalIds property
func (m *IdentityProtectionRiskyServicePrincipalsDismissPostRequestBody) GetServicePrincipalIds()([]string) {
    return m.servicePrincipalIds
}
// Serialize serializes information the current object
func (m *IdentityProtectionRiskyServicePrincipalsDismissPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetServicePrincipalIds() != nil {
        err := writer.WriteCollectionOfStringValues("servicePrincipalIds", m.GetServicePrincipalIds())
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
func (m *IdentityProtectionRiskyServicePrincipalsDismissPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetServicePrincipalIds sets the servicePrincipalIds property value. The servicePrincipalIds property
func (m *IdentityProtectionRiskyServicePrincipalsDismissPostRequestBody) SetServicePrincipalIds(value []string)() {
    m.servicePrincipalIds = value
}
