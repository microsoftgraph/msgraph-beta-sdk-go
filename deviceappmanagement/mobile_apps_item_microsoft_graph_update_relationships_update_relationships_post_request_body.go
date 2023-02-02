package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// MobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody 
type MobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The relationships property
    relationships []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable
}
// NewMobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody instantiates a new MobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody and sets the default values.
func NewMobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody()(*MobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody) {
    m := &MobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["relationships"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppRelationshipFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable)
            }
            m.SetRelationships(res)
        }
        return nil
    }
    return res
}
// GetRelationships gets the relationships property value. The relationships property
func (m *MobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody) GetRelationships()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable) {
    return m.relationships
}
// Serialize serializes information the current object
func (m *MobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRelationships() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRelationships()))
        for i, v := range m.GetRelationships() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("relationships", cast)
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
func (m *MobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRelationships sets the relationships property value. The relationships property
func (m *MobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsPostRequestBody) SetRelationships(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppRelationshipable)() {
    m.relationships = value
}
