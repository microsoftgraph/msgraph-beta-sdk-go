package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// VirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody 
type VirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The assignments property
    assignments []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcUserSettingAssignmentable
}
// NewVirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody instantiates a new VirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody and sets the default values.
func NewVirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody()(*VirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody) {
    m := &VirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignments gets the assignments property value. The assignments property
func (m *VirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody) GetAssignments()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcUserSettingAssignmentable) {
    return m.assignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcUserSettingAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcUserSettingAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcUserSettingAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *VirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assignments", cast)
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
func (m *VirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignments sets the assignments property value. The assignments property
func (m *VirtualEndpointUserSettingsItemMicrosoftGraphAssignAssignPostRequestBody) SetAssignments(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcUserSettingAssignmentable)() {
    m.assignments = value
}
