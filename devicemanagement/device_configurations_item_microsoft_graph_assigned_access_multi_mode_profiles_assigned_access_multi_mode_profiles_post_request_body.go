package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody 
type DeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The assignedAccessMultiModeProfiles property
    assignedAccessMultiModeProfiles []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable
}
// NewDeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody instantiates a new DeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody and sets the default values.
func NewDeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody()(*DeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody) {
    m := &DeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignedAccessMultiModeProfiles gets the assignedAccessMultiModeProfiles property value. The assignedAccessMultiModeProfiles property
func (m *DeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody) GetAssignedAccessMultiModeProfiles()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable) {
    return m.assignedAccessMultiModeProfiles
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedAccessMultiModeProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAssignedAccessProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable)
            }
            m.SetAssignedAccessMultiModeProfiles(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignedAccessMultiModeProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedAccessMultiModeProfiles()))
        for i, v := range m.GetAssignedAccessMultiModeProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assignedAccessMultiModeProfiles", cast)
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
func (m *DeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignedAccessMultiModeProfiles sets the assignedAccessMultiModeProfiles property value. The assignedAccessMultiModeProfiles property
func (m *DeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesPostRequestBody) SetAssignedAccessMultiModeProfiles(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable)() {
    m.assignedAccessMultiModeProfiles = value
}
