package managedtenants

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
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
    // The OdataType property
    odataType *string
    // The valueType property
    valueType *ManagementParameterValueType
}
// NewTemplateParameter instantiates a new templateParameter and sets the default values.
func NewTemplateParameter()(*TemplateParameter) {
    m := &TemplateParameter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.managedTenants.templateParameter";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTemplateParameterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTemplateParameterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTemplateParameter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TemplateParameter) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDescription gets the description property value. The description for the template parameter. Optional. Read-only.
func (m *TemplateParameter) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name for the template parameter. Required. Read-only.
func (m *TemplateParameter) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TemplateParameter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["jsonAllowedValues"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJsonAllowedValues)
    res["jsonDefaultValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJsonDefaultValue)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["valueType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagementParameterValueType , m.SetValueType)
    return res
}
// GetJsonAllowedValues gets the jsonAllowedValues property value. The allowed values for the template parameter represented by a serialized string of JSON. Optional. Read-only.
func (m *TemplateParameter) GetJsonAllowedValues()(*string) {
    return m.jsonAllowedValues
}
// GetJsonDefaultValue gets the jsonDefaultValue property value. The default value for the template parameter represented by a serialized string of JSON. Required. Read-only.
func (m *TemplateParameter) GetJsonDefaultValue()(*string) {
    return m.jsonDefaultValue
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TemplateParameter) GetOdataType()(*string) {
    return m.odataType
}
// GetValueType gets the valueType property value. The valueType property
func (m *TemplateParameter) GetValueType()(*ManagementParameterValueType) {
    return m.valueType
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
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    m.additionalData = value
}
// SetDescription sets the description property value. The description for the template parameter. Optional. Read-only.
func (m *TemplateParameter) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name for the template parameter. Required. Read-only.
func (m *TemplateParameter) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetJsonAllowedValues sets the jsonAllowedValues property value. The allowed values for the template parameter represented by a serialized string of JSON. Optional. Read-only.
func (m *TemplateParameter) SetJsonAllowedValues(value *string)() {
    m.jsonAllowedValues = value
}
// SetJsonDefaultValue sets the jsonDefaultValue property value. The default value for the template parameter represented by a serialized string of JSON. Required. Read-only.
func (m *TemplateParameter) SetJsonDefaultValue(value *string)() {
    m.jsonDefaultValue = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TemplateParameter) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValueType sets the valueType property value. The valueType property
func (m *TemplateParameter) SetValueType(value *ManagementParameterValueType)() {
    m.valueType = value
}
