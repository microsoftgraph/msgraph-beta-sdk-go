package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementSettingComparison provides operations to call the compare method.
type DeviceManagementSettingComparison struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Setting comparison result. Possible values are: unknown, equal, notEqual, added, removed.
    comparisonResult *DeviceManagementComparisonResult;
    // JSON representation of current intent (or) template setting's value
    currentValueJson *string;
    // The ID of the setting definition for this instance
    definitionId *string;
    // The setting's display name
    displayName *string;
    // The setting ID
    id *string;
    // JSON representation of new template setting's value
    newValueJson *string;
}
// NewDeviceManagementSettingComparison instantiates a new deviceManagementSettingComparison and sets the default values.
func NewDeviceManagementSettingComparison()(*DeviceManagementSettingComparison) {
    m := &DeviceManagementSettingComparison{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementSettingComparisonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingComparisonFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceManagementSettingComparison(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettingComparison) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetComparisonResult gets the comparisonResult property value. Setting comparison result. Possible values are: unknown, equal, notEqual, added, removed.
func (m *DeviceManagementSettingComparison) GetComparisonResult()(*DeviceManagementComparisonResult) {
    if m == nil {
        return nil
    } else {
        return m.comparisonResult
    }
}
// GetCurrentValueJson gets the currentValueJson property value. JSON representation of current intent (or) template setting's value
func (m *DeviceManagementSettingComparison) GetCurrentValueJson()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currentValueJson
    }
}
// GetDefinitionId gets the definitionId property value. The ID of the setting definition for this instance
func (m *DeviceManagementSettingComparison) GetDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.definitionId
    }
}
// GetDisplayName gets the displayName property value. The setting's display name
func (m *DeviceManagementSettingComparison) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingComparison) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["comparisonResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementComparisonResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComparisonResult(val.(*DeviceManagementComparisonResult))
        }
        return nil
    }
    res["currentValueJson"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentValueJson(val)
        }
        return nil
    }
    res["definitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionId(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["newValueJson"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewValueJson(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The setting ID
func (m *DeviceManagementSettingComparison) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetNewValueJson gets the newValueJson property value. JSON representation of new template setting's value
func (m *DeviceManagementSettingComparison) GetNewValueJson()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newValueJson
    }
}
func (m *DeviceManagementSettingComparison) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingComparison) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettingComparison) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetComparisonResult sets the comparisonResult property value. Setting comparison result. Possible values are: unknown, equal, notEqual, added, removed.
func (m *DeviceManagementSettingComparison) SetComparisonResult(value *DeviceManagementComparisonResult)() {
    if m != nil {
        m.comparisonResult = value
    }
}
// SetCurrentValueJson sets the currentValueJson property value. JSON representation of current intent (or) template setting's value
func (m *DeviceManagementSettingComparison) SetCurrentValueJson(value *string)() {
    if m != nil {
        m.currentValueJson = value
    }
}
// SetDefinitionId sets the definitionId property value. The ID of the setting definition for this instance
func (m *DeviceManagementSettingComparison) SetDefinitionId(value *string)() {
    if m != nil {
        m.definitionId = value
    }
}
// SetDisplayName sets the displayName property value. The setting's display name
func (m *DeviceManagementSettingComparison) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetId sets the id property value. The setting ID
func (m *DeviceManagementSettingComparison) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetNewValueJson sets the newValueJson property value. JSON representation of new template setting's value
func (m *DeviceManagementSettingComparison) SetNewValueJson(value *string)() {
    if m != nil {
        m.newValueJson = value
    }
}
