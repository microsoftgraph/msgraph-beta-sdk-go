package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptParameter base properties of the script parameter.
type DeviceHealthScriptParameter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Whether Apply DefaultValue When Not Assigned
    applyDefaultValueWhenNotAssigned *bool
    // The description of the param
    description *string
    // Whether the param is required
    isRequired *bool
    // The name of the param
    name *string
    // The OdataType property
    odataType *string
}
// NewDeviceHealthScriptParameter instantiates a new deviceHealthScriptParameter and sets the default values.
func NewDeviceHealthScriptParameter()(*DeviceHealthScriptParameter) {
    m := &DeviceHealthScriptParameter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.deviceHealthScriptParameter";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceHealthScriptParameterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptParameterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.deviceHealthScriptBooleanParameter":
                        return NewDeviceHealthScriptBooleanParameter(), nil
                    case "#microsoft.graph.deviceHealthScriptIntegerParameter":
                        return NewDeviceHealthScriptIntegerParameter(), nil
                    case "#microsoft.graph.deviceHealthScriptStringParameter":
                        return NewDeviceHealthScriptStringParameter(), nil
                }
            }
        }
    }
    return NewDeviceHealthScriptParameter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptParameter) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApplyDefaultValueWhenNotAssigned gets the applyDefaultValueWhenNotAssigned property value. Whether Apply DefaultValue When Not Assigned
func (m *DeviceHealthScriptParameter) GetApplyDefaultValueWhenNotAssigned()(*bool) {
    return m.applyDefaultValueWhenNotAssigned
}
// GetDescription gets the description property value. The description of the param
func (m *DeviceHealthScriptParameter) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptParameter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applyDefaultValueWhenNotAssigned"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetApplyDefaultValueWhenNotAssigned)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["isRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsRequired)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetIsRequired gets the isRequired property value. Whether the param is required
func (m *DeviceHealthScriptParameter) GetIsRequired()(*bool) {
    return m.isRequired
}
// GetName gets the name property value. The name of the param
func (m *DeviceHealthScriptParameter) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptParameter) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptParameter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("applyDefaultValueWhenNotAssigned", m.GetApplyDefaultValueWhenNotAssigned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *DeviceHealthScriptParameter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApplyDefaultValueWhenNotAssigned sets the applyDefaultValueWhenNotAssigned property value. Whether Apply DefaultValue When Not Assigned
func (m *DeviceHealthScriptParameter) SetApplyDefaultValueWhenNotAssigned(value *bool)() {
    m.applyDefaultValueWhenNotAssigned = value
}
// SetDescription sets the description property value. The description of the param
func (m *DeviceHealthScriptParameter) SetDescription(value *string)() {
    m.description = value
}
// SetIsRequired sets the isRequired property value. Whether the param is required
func (m *DeviceHealthScriptParameter) SetIsRequired(value *bool)() {
    m.isRequired = value
}
// SetName sets the name property value. The name of the param
func (m *DeviceHealthScriptParameter) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptParameter) SetOdataType(value *string)() {
    m.odataType = value
}
