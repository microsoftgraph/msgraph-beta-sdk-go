package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingComparison entity representing setting comparison result
type DeviceManagementSettingComparison struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Setting comparison result type
    comparisonResult *DeviceManagementComparisonResult
    // JSON representation of current intent (or) template setting's value
    currentValueJson *string
    // The ID of the setting definition for this instance
    definitionId *string
    // The setting's display name
    displayName *string
    // The setting ID
    id *string
    // JSON representation of new template setting's value
    newValueJson *string
    // The OdataType property
    odataType *string
}
// NewDeviceManagementSettingComparison instantiates a new deviceManagementSettingComparison and sets the default values.
func NewDeviceManagementSettingComparison()(*DeviceManagementSettingComparison) {
    m := &DeviceManagementSettingComparison{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementSettingComparisonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingComparisonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingComparison(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettingComparison) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetComparisonResult gets the comparisonResult property value. Setting comparison result type
func (m *DeviceManagementSettingComparison) GetComparisonResult()(*DeviceManagementComparisonResult) {
    return m.comparisonResult
}
// GetCurrentValueJson gets the currentValueJson property value. JSON representation of current intent (or) template setting's value
func (m *DeviceManagementSettingComparison) GetCurrentValueJson()(*string) {
    return m.currentValueJson
}
// GetDefinitionId gets the definitionId property value. The ID of the setting definition for this instance
func (m *DeviceManagementSettingComparison) GetDefinitionId()(*string) {
    return m.definitionId
}
// GetDisplayName gets the displayName property value. The setting's display name
func (m *DeviceManagementSettingComparison) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingComparison) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comparisonResult"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementComparisonResult , m.SetComparisonResult)
    res["currentValueJson"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCurrentValueJson)
    res["definitionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefinitionId)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["id"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetId)
    res["newValueJson"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNewValueJson)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetId gets the id property value. The setting ID
func (m *DeviceManagementSettingComparison) GetId()(*string) {
    return m.id
}
// GetNewValueJson gets the newValueJson property value. JSON representation of new template setting's value
func (m *DeviceManagementSettingComparison) GetNewValueJson()(*string) {
    return m.newValueJson
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementSettingComparison) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingComparison) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetComparisonResult() != nil {
        cast := (*m.GetComparisonResult()).String()
        err := writer.WriteStringValue("comparisonResult", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("currentValueJson", m.GetCurrentValueJson())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("definitionId", m.GetDefinitionId())
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
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("newValueJson", m.GetNewValueJson())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *DeviceManagementSettingComparison) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetComparisonResult sets the comparisonResult property value. Setting comparison result type
func (m *DeviceManagementSettingComparison) SetComparisonResult(value *DeviceManagementComparisonResult)() {
    m.comparisonResult = value
}
// SetCurrentValueJson sets the currentValueJson property value. JSON representation of current intent (or) template setting's value
func (m *DeviceManagementSettingComparison) SetCurrentValueJson(value *string)() {
    m.currentValueJson = value
}
// SetDefinitionId sets the definitionId property value. The ID of the setting definition for this instance
func (m *DeviceManagementSettingComparison) SetDefinitionId(value *string)() {
    m.definitionId = value
}
// SetDisplayName sets the displayName property value. The setting's display name
func (m *DeviceManagementSettingComparison) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. The setting ID
func (m *DeviceManagementSettingComparison) SetId(value *string)() {
    m.id = value
}
// SetNewValueJson sets the newValueJson property value. JSON representation of new template setting's value
func (m *DeviceManagementSettingComparison) SetNewValueJson(value *string)() {
    m.newValueJson = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementSettingComparison) SetOdataType(value *string)() {
    m.odataType = value
}
