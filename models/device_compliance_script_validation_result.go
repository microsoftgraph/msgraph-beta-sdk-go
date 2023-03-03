package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceComplianceScriptValidationResult 
type DeviceComplianceScriptValidationResult struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceComplianceScriptValidationResult instantiates a new deviceComplianceScriptValidationResult and sets the default values.
func NewDeviceComplianceScriptValidationResult()(*DeviceComplianceScriptValidationResult) {
    m := &DeviceComplianceScriptValidationResult{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceComplianceScriptValidationResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComplianceScriptValidationResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceComplianceScriptValidationResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceComplianceScriptValidationResult) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *DeviceComplianceScriptValidationResult) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
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
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRuleErrors gets the ruleErrors property value. Errors in json for the script for rules.
func (m *DeviceComplianceScriptValidationResult) GetRuleErrors()([]DeviceComplianceScriptRuleErrorable) {
    val, err := m.GetBackingStore().Get("ruleErrors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceComplianceScriptRuleErrorable)
    }
    return nil
}
// GetRules gets the rules property value. Parsed rules from json.
func (m *DeviceComplianceScriptValidationResult) GetRules()([]DeviceComplianceScriptRuleable) {
    val, err := m.GetBackingStore().Get("rules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceComplianceScriptRuleable)
    }
    return nil
}
// GetScriptErrors gets the scriptErrors property value. Errors in json for the script.
func (m *DeviceComplianceScriptValidationResult) GetScriptErrors()([]DeviceComplianceScriptErrorable) {
    val, err := m.GetBackingStore().Get("scriptErrors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceComplianceScriptErrorable)
    }
    return nil
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
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *DeviceComplianceScriptValidationResult) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceComplianceScriptValidationResult) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRuleErrors sets the ruleErrors property value. Errors in json for the script for rules.
func (m *DeviceComplianceScriptValidationResult) SetRuleErrors(value []DeviceComplianceScriptRuleErrorable)() {
    err := m.GetBackingStore().Set("ruleErrors", value)
    if err != nil {
        panic(err)
    }
}
// SetRules sets the rules property value. Parsed rules from json.
func (m *DeviceComplianceScriptValidationResult) SetRules(value []DeviceComplianceScriptRuleable)() {
    err := m.GetBackingStore().Set("rules", value)
    if err != nil {
        panic(err)
    }
}
// SetScriptErrors sets the scriptErrors property value. Errors in json for the script.
func (m *DeviceComplianceScriptValidationResult) SetScriptErrors(value []DeviceComplianceScriptErrorable)() {
    err := m.GetBackingStore().Set("scriptErrors", value)
    if err != nil {
        panic(err)
    }
}
// DeviceComplianceScriptValidationResultable 
type DeviceComplianceScriptValidationResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetRuleErrors()([]DeviceComplianceScriptRuleErrorable)
    GetRules()([]DeviceComplianceScriptRuleable)
    GetScriptErrors()([]DeviceComplianceScriptErrorable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetRuleErrors(value []DeviceComplianceScriptRuleErrorable)()
    SetRules(value []DeviceComplianceScriptRuleable)()
    SetScriptErrors(value []DeviceComplianceScriptErrorable)()
}
