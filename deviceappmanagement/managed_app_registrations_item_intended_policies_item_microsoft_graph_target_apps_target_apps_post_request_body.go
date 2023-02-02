package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody 
type ManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The apps property
    apps []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedMobileAppable
}
// NewManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody instantiates a new ManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody and sets the default values.
func NewManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody()(*ManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) {
    m := &ManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApps gets the apps property value. The apps property
func (m *ManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) GetApps()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedMobileAppable) {
    return m.apps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedMobileAppable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedMobileAppable)
            }
            m.SetApps(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("apps", cast)
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
func (m *ManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApps sets the apps property value. The apps property
func (m *ManagedAppRegistrationsItemIntendedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBody) SetApps(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedMobileAppable)() {
    m.apps = value
}
