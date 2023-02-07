package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody 
type DeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deviceManagementScriptAssignments property
    deviceManagementScriptAssignments []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptAssignmentable
    // The deviceManagementScriptGroupAssignments property
    deviceManagementScriptGroupAssignments []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable
}
// NewDeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody instantiates a new DeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody and sets the default values.
func NewDeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody()(*DeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody) {
    m := &DeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceManagementScriptAssignments gets the deviceManagementScriptAssignments property value. The deviceManagementScriptAssignments property
func (m *DeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody) GetDeviceManagementScriptAssignments()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptAssignmentable) {
    return m.deviceManagementScriptAssignments
}
// GetDeviceManagementScriptGroupAssignments gets the deviceManagementScriptGroupAssignments property value. The deviceManagementScriptGroupAssignments property
func (m *DeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody) GetDeviceManagementScriptGroupAssignments()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable) {
    return m.deviceManagementScriptGroupAssignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceManagementScriptAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptAssignmentable)
            }
            m.SetDeviceManagementScriptAssignments(res)
        }
        return nil
    }
    res["deviceManagementScriptGroupAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptGroupAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable)
            }
            m.SetDeviceManagementScriptGroupAssignments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceManagementScriptAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceManagementScriptAssignments()))
        for i, v := range m.GetDeviceManagementScriptAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("deviceManagementScriptAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceManagementScriptGroupAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceManagementScriptGroupAssignments()))
        for i, v := range m.GetDeviceManagementScriptGroupAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("deviceManagementScriptGroupAssignments", cast)
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
func (m *DeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceManagementScriptAssignments sets the deviceManagementScriptAssignments property value. The deviceManagementScriptAssignments property
func (m *DeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody) SetDeviceManagementScriptAssignments(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptAssignmentable)() {
    m.deviceManagementScriptAssignments = value
}
// SetDeviceManagementScriptGroupAssignments sets the deviceManagementScriptGroupAssignments property value. The deviceManagementScriptGroupAssignments property
func (m *DeviceManagementScriptsItemMicrosoftGraphAssignAssignPostRequestBody) SetDeviceManagementScriptGroupAssignments(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptGroupAssignmentable)() {
    m.deviceManagementScriptGroupAssignments = value
}
