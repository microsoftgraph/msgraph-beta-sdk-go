package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TemplateParameter 
type TemplateParameter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The description for the template parameter. Optional. Read-only.
    description *string
    // The display name for the template parameter. Required. Read-only.
    displayName *string
    // The allowed values for the template parameter represented by a serialized string of JSON. Optional. Read-only.
    jsonAllowedValues *string
    // The default value for the template parameter represented by a serialized string of JSON. Required. Read-only.
    jsonDefaultValue *string
    // The data type for the template parameter.. Possible values are: string, integer, boolean, guid, stringCollection, integerCollection, booleanCollection, guidCollection, unknownFutureValue. Required. Read-only.
    valueType *ManagementParameterValueType
}
// NewTemplateParameter instantiates a new templateParameter and sets the default values.
func NewTemplateParameter()(*TemplateParameter) {
    m := &TemplateParameter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTemplateParameterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTemplateParameterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTemplateParameter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TemplateParameter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. The description for the template parameter. Optional. Read-only.
func (m *TemplateParameter) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name for the template parameter. Required. Read-only.
func (m *TemplateParameter) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TemplateParameter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["jsonAllowedValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJsonAllowedValues(val)
        }
        return nil
    }
    res["jsonDefaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJsonDefaultValue(val)
        }
        return nil
    }
    res["valueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementParameterValueType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueType(val.(*ManagementParameterValueType))
        }
        return nil
    }
    return res
}
// GetJsonAllowedValues gets the jsonAllowedValues property value. The allowed values for the template parameter represented by a serialized string of JSON. Optional. Read-only.
func (m *TemplateParameter) GetJsonAllowedValues()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jsonAllowedValues
    }
}
// GetJsonDefaultValue gets the jsonDefaultValue property value. The default value for the template parameter represented by a serialized string of JSON. Required. Read-only.
func (m *TemplateParameter) GetJsonDefaultValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jsonDefaultValue
    }
}
// GetValueType gets the valueType property value. The data type for the template parameter.. Possible values are: string, integer, boolean, guid, stringCollection, integerCollection, booleanCollection, guidCollection, unknownFutureValue. Required. Read-only.
func (m *TemplateParameter) GetValueType()(*ManagementParameterValueType) {
    if m == nil {
        return nil
    } else {
        return m.valueType
    }
}
// Serialize serializes information the current object
func (m *TemplateParameter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteStringValue("jsonAllowedValues", m.GetJsonAllowedValues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jsonDefaultValue", m.GetJsonDefaultValue())
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
func (m *TemplateParameter) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. The description for the template parameter. Optional. Read-only.
func (m *TemplateParameter) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the template parameter. Required. Read-only.
func (m *TemplateParameter) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetJsonAllowedValues sets the jsonAllowedValues property value. The allowed values for the template parameter represented by a serialized string of JSON. Optional. Read-only.
func (m *TemplateParameter) SetJsonAllowedValues(value *string)() {
    if m != nil {
        m.jsonAllowedValues = value
    }
}
// SetJsonDefaultValue sets the jsonDefaultValue property value. The default value for the template parameter represented by a serialized string of JSON. Required. Read-only.
func (m *TemplateParameter) SetJsonDefaultValue(value *string)() {
    if m != nil {
        m.jsonDefaultValue = value
    }
}
// SetValueType sets the valueType property value. The data type for the template parameter.. Possible values are: string, integer, boolean, guid, stringCollection, integerCollection, booleanCollection, guidCollection, unknownFutureValue. Required. Read-only.
func (m *TemplateParameter) SetValueType(value *ManagementParameterValueType)() {
    if m != nil {
        m.valueType = value
    }
}
