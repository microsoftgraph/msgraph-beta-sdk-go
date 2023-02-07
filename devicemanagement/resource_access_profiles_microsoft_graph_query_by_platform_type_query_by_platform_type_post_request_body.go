package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody 
type ResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Supported platform types for policies.
    platformType *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyPlatformType
}
// NewResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody instantiates a new ResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody and sets the default values.
func NewResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody()(*ResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody) {
    m := &ResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["platformType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyPlatformType))
        }
        return nil
    }
    return res
}
// GetPlatformType gets the platformType property value. Supported platform types for policies.
func (m *ResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody) GetPlatformType()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyPlatformType) {
    return m.platformType
}
// Serialize serializes information the current object
func (m *ResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPlatformType() != nil {
        cast := (*m.GetPlatformType()).String()
        err := writer.WriteStringValue("platformType", &cast)
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
func (m *ResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPlatformType sets the platformType property value. Supported platform types for policies.
func (m *ResourceAccessProfilesMicrosoftGraphQueryByPlatformTypeQueryByPlatformTypePostRequestBody) SetPlatformType(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyPlatformType)() {
    m.platformType = value
}
