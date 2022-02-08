package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// Setting 
type Setting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The display name for the setting. Required. Read-only.
    displayName *string;
    // The value for the setting serialized as string of JSON. Required. Read-only.
    jsonValue *string;
    // A flag indicating whether the setting can be override existing configurations when applied. Required. Read-only.
    overwriteAllowed *bool;
    // 
    settingId *string;
    // The data type for the setting. Possible values are: string, integer, boolean, guid, stringCollection, integerCollection, booleanCollection, guidCollection, unknownFutureValue. Required. Read-only.
    valueType *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType;
}
// NewSetting instantiates a new setting and sets the default values.
func NewSetting()(*Setting) {
    m := &Setting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Setting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. The display name for the setting. Required. Read-only.
func (m *Setting) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetJsonValue gets the jsonValue property value. The value for the setting serialized as string of JSON. Required. Read-only.
func (m *Setting) GetJsonValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jsonValue
    }
}
// GetOverwriteAllowed gets the overwriteAllowed property value. A flag indicating whether the setting can be override existing configurations when applied. Required. Read-only.
func (m *Setting) GetOverwriteAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.overwriteAllowed
    }
}
// GetSettingId gets the settingId property value. 
func (m *Setting) GetSettingId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingId
    }
}
// GetValueType gets the valueType property value. The data type for the setting. Possible values are: string, integer, boolean, guid, stringCollection, integerCollection, booleanCollection, guidCollection, unknownFutureValue. Required. Read-only.
func (m *Setting) GetValueType()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType) {
    if m == nil {
        return nil
    } else {
        return m.valueType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Setting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["jsonValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJsonValue(val)
        }
        return nil
    }
    res["overwriteAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverwriteAllowed(val)
        }
        return nil
    }
    res["settingId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingId(val)
        }
        return nil
    }
    res["valueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseManagementParameterValueType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueType(val.(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType))
        }
        return nil
    }
    return res
}
func (m *Setting) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Setting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jsonValue", m.GetJsonValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("overwriteAllowed", m.GetOverwriteAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingId", m.GetSettingId())
        if err != nil {
            return err
        }
    }
    if m.GetValueType() != nil {
        cast := (*m.GetValueType()).String()
        err := writer.WriteStringValue("valueType", &cast)
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
func (m *Setting) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the setting. Required. Read-only.
func (m *Setting) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetJsonValue sets the jsonValue property value. The value for the setting serialized as string of JSON. Required. Read-only.
func (m *Setting) SetJsonValue(value *string)() {
    if m != nil {
        m.jsonValue = value
    }
}
// SetOverwriteAllowed sets the overwriteAllowed property value. A flag indicating whether the setting can be override existing configurations when applied. Required. Read-only.
func (m *Setting) SetOverwriteAllowed(value *bool)() {
    if m != nil {
        m.overwriteAllowed = value
    }
}
// SetSettingId sets the settingId property value. 
func (m *Setting) SetSettingId(value *string)() {
    if m != nil {
        m.settingId = value
    }
}
// SetValueType sets the valueType property value. The data type for the setting. Possible values are: string, integer, boolean, guid, stringCollection, integerCollection, booleanCollection, guidCollection, unknownFutureValue. Required. Read-only.
func (m *Setting) SetValueType(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementParameterValueType)() {
    if m != nil {
        m.valueType = value
    }
}
