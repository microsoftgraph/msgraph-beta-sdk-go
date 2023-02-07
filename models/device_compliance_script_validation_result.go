package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceScriptValidationResult 
type DeviceComplianceScriptValidationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Errors in json for the script for rules.
    ruleErrors []DeviceComplianceScriptRuleErrorable
    // Parsed rules from json.
    rules []DeviceComplianceScriptRuleable
    // Errors in json for the script.
    scriptErrors []DeviceComplianceScriptErrorable
}
// NewDeviceComplianceScriptValidationResult instantiates a new deviceComplianceScriptValidationResult and sets the default values.
func NewDeviceComplianceScriptValidationResult()(*DeviceComplianceScriptValidationResult) {
    m := &DeviceComplianceScriptValidationResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceComplianceScriptValidationResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComplianceScriptValidationResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceComplianceScriptValidationResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceComplianceScriptValidationResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceScriptValidationResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["ruleErrors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceComplianceScriptRuleErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptRuleErrorable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceComplianceScriptRuleErrorable)
            }
            m.SetRuleErrors(res)
        }
        return nil
    }
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceComplianceScriptRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptRuleable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceComplianceScriptRuleable)
            }
            m.SetRules(res)
        }
        return nil
    }
    res["scriptErrors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceComplianceScriptErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptErrorable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceComplianceScriptErrorable)
            }
            m.SetScriptErrors(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceComplianceScriptValidationResult) GetOdataType()(*string) {
    return m.odataType
}
// GetRuleErrors gets the ruleErrors property value. Errors in json for the script for rules.
func (m *DeviceComplianceScriptValidationResult) GetRuleErrors()([]DeviceComplianceScriptRuleErrorable) {
    return m.ruleErrors
}
// GetRules gets the rules property value. Parsed rules from json.
func (m *DeviceComplianceScriptValidationResult) GetRules()([]DeviceComplianceScriptRuleable) {
    return m.rules
}
// GetScriptErrors gets the scriptErrors property value. Errors in json for the script.
func (m *DeviceComplianceScriptValidationResult) GetScriptErrors()([]DeviceComplianceScriptErrorable) {
    return m.scriptErrors
}
// Serialize serializes information the current object
func (m *DeviceComplianceScriptValidationResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRuleErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRuleErrors()))
        for i, v := range m.GetRuleErrors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("ruleErrors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("rules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetScriptErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetScriptErrors()))
        for i, v := range m.GetScriptErrors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("scriptErrors", cast)
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
func (m *DeviceComplianceScriptValidationResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceComplianceScriptValidationResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRuleErrors sets the ruleErrors property value. Errors in json for the script for rules.
func (m *DeviceComplianceScriptValidationResult) SetRuleErrors(value []DeviceComplianceScriptRuleErrorable)() {
    m.ruleErrors = value
}
// SetRules sets the rules property value. Parsed rules from json.
func (m *DeviceComplianceScriptValidationResult) SetRules(value []DeviceComplianceScriptRuleable)() {
    m.rules = value
}
// SetScriptErrors sets the scriptErrors property value. Errors in json for the script.
func (m *DeviceComplianceScriptValidationResult) SetScriptErrors(value []DeviceComplianceScriptErrorable)() {
    m.scriptErrors = value
}
