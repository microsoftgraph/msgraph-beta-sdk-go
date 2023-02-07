package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody 
type DeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deviceHealthScriptAssignments property
    deviceHealthScriptAssignments []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptAssignmentable
}
// NewDeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody instantiates a new DeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody and sets the default values.
func NewDeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody()(*DeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody) {
    m := &DeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceHealthScriptAssignments gets the deviceHealthScriptAssignments property value. The deviceHealthScriptAssignments property
func (m *DeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody) GetDeviceHealthScriptAssignments()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptAssignmentable) {
    return m.deviceHealthScriptAssignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceHealthScriptAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptAssignmentable)
            }
            m.SetDeviceHealthScriptAssignments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceHealthScriptAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceHealthScriptAssignments()))
        for i, v := range m.GetDeviceHealthScriptAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("deviceHealthScriptAssignments", cast)
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
func (m *DeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceHealthScriptAssignments sets the deviceHealthScriptAssignments property value. The deviceHealthScriptAssignments property
func (m *DeviceHealthScriptsItemMicrosoftGraphAssignAssignPostRequestBody) SetDeviceHealthScriptAssignments(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptAssignmentable)() {
    m.deviceHealthScriptAssignments = value
}
