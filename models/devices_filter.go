package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DevicesFilter 
type DevicesFilter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The mode property
    mode *CrossTenantAccessPolicyTargetConfigurationAccessType
    // The OdataType property
    odataType *string
    // The rule property
    rule *string
}
// NewDevicesFilter instantiates a new devicesFilter and sets the default values.
func NewDevicesFilter()(*DevicesFilter) {
    m := &DevicesFilter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.devicesFilter";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDevicesFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDevicesFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDevicesFilter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DevicesFilter) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DevicesFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCrossTenantAccessPolicyTargetConfigurationAccessType , m.SetMode)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["rule"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRule)
    return res
}
// GetMode gets the mode property value. The mode property
func (m *DevicesFilter) GetMode()(*CrossTenantAccessPolicyTargetConfigurationAccessType) {
    return m.mode
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DevicesFilter) GetOdataType()(*string) {
    return m.odataType
}
// GetRule gets the rule property value. The rule property
func (m *DevicesFilter) GetRule()(*string) {
    return m.rule
}
// Serialize serializes information the current object
func (m *DevicesFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMode() != nil {
        cast := (*m.GetMode()).String()
        err := writer.WriteStringValue("mode", &cast)
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
        err := writer.WriteStringValue("rule", m.GetRule())
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
func (m *DevicesFilter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMode sets the mode property value. The mode property
func (m *DevicesFilter) SetMode(value *CrossTenantAccessPolicyTargetConfigurationAccessType)() {
    m.mode = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DevicesFilter) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRule sets the rule property value. The rule property
func (m *DevicesFilter) SetRule(value *string)() {
    m.rule = value
}
