package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody 
type TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description property
    description *string
    // The displayName property
    displayName *string
    // The roleScopeTagIds property
    roleScopeTagIds []string
    // The settingsDelta property
    settingsDelta []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingInstanceable
}
// NewTemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody instantiates a new TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody and sets the default values.
func NewTemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody()(*TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody) {
    m := &TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description property
func (m *TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["settingsDelta"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingInstanceable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingInstanceable)
            }
            m.SetSettingsDelta(res)
        }
        return nil
    }
    return res
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. The roleScopeTagIds property
func (m *TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetSettingsDelta gets the settingsDelta property value. The settingsDelta property
func (m *TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody) GetSettingsDelta()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingInstanceable) {
    return m.settingsDelta
}
// Serialize serializes information the current object
func (m *TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err := writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    if m.GetSettingsDelta() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettingsDelta()))
        for i, v := range m.GetSettingsDelta() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("settingsDelta", cast)
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
func (m *TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description property
func (m *TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. The roleScopeTagIds property
func (m *TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetSettingsDelta sets the settingsDelta property value. The settingsDelta property
func (m *TemplatesItemMicrosoftGraphCreateInstanceCreateInstancePostRequestBody) SetSettingsDelta(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementSettingInstanceable)() {
    m.settingsDelta = value
}
