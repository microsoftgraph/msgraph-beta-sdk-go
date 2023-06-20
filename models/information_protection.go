package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtection 
type InformationProtection struct {
    Entity
}
// NewInformationProtection instantiates a new informationProtection and sets the default values.
func NewInformationProtection()(*InformationProtection) {
    m := &InformationProtection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInformationProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtection(), nil
}
// GetBitlocker gets the bitlocker property value. The bitlocker property
func (m *InformationProtection) GetBitlocker()(Bitlockerable) {
    val, err := m.GetBackingStore().Get("bitlocker")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Bitlockerable)
    }
    return nil
}
// GetDataLossPreventionPolicies gets the dataLossPreventionPolicies property value. The dataLossPreventionPolicies property
func (m *InformationProtection) GetDataLossPreventionPolicies()([]DataLossPreventionPolicyable) {
    val, err := m.GetBackingStore().Get("dataLossPreventionPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DataLossPreventionPolicyable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bitlocker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBitlockerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitlocker(val.(Bitlockerable))
        }
        return nil
    }
    res["dataLossPreventionPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDataLossPreventionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DataLossPreventionPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DataLossPreventionPolicyable)
                }
            }
            m.SetDataLossPreventionPolicies(res)
        }
        return nil
    }
    res["policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInformationProtectionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val.(InformationProtectionPolicyable))
        }
        return nil
    }
    res["sensitivityLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitivityLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitivityLabelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SensitivityLabelable)
                }
            }
            m.SetSensitivityLabels(res)
        }
        return nil
    }
    res["sensitivityPolicySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSensitivityPolicySettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivityPolicySettings(val.(SensitivityPolicySettingsable))
        }
        return nil
    }
    res["threatAssessmentRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateThreatAssessmentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ThreatAssessmentRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ThreatAssessmentRequestable)
                }
            }
            m.SetThreatAssessmentRequests(res)
        }
        return nil
    }
    return res
}
// GetPolicy gets the policy property value. The policy property
func (m *InformationProtection) GetPolicy()(InformationProtectionPolicyable) {
    val, err := m.GetBackingStore().Get("policy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(InformationProtectionPolicyable)
    }
    return nil
}
// GetSensitivityLabels gets the sensitivityLabels property value. The sensitivityLabels property
func (m *InformationProtection) GetSensitivityLabels()([]SensitivityLabelable) {
    val, err := m.GetBackingStore().Get("sensitivityLabels")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SensitivityLabelable)
    }
    return nil
}
// GetSensitivityPolicySettings gets the sensitivityPolicySettings property value. The sensitivityPolicySettings property
func (m *InformationProtection) GetSensitivityPolicySettings()(SensitivityPolicySettingsable) {
    val, err := m.GetBackingStore().Get("sensitivityPolicySettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SensitivityPolicySettingsable)
    }
    return nil
}
// GetThreatAssessmentRequests gets the threatAssessmentRequests property value. The threatAssessmentRequests property
func (m *InformationProtection) GetThreatAssessmentRequests()([]ThreatAssessmentRequestable) {
    val, err := m.GetBackingStore().Get("threatAssessmentRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ThreatAssessmentRequestable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *InformationProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("bitlocker", m.GetBitlocker())
        if err != nil {
            return err
        }
    }
    if m.GetDataLossPreventionPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDataLossPreventionPolicies()))
        for i, v := range m.GetDataLossPreventionPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("dataLossPreventionPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetSensitivityLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSensitivityLabels()))
        for i, v := range m.GetSensitivityLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sensitivityLabels", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sensitivityPolicySettings", m.GetSensitivityPolicySettings())
        if err != nil {
            return err
        }
    }
    if m.GetThreatAssessmentRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetThreatAssessmentRequests()))
        for i, v := range m.GetThreatAssessmentRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("threatAssessmentRequests", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBitlocker sets the bitlocker property value. The bitlocker property
func (m *InformationProtection) SetBitlocker(value Bitlockerable)() {
    err := m.GetBackingStore().Set("bitlocker", value)
    if err != nil {
        panic(err)
    }
}
// SetDataLossPreventionPolicies sets the dataLossPreventionPolicies property value. The dataLossPreventionPolicies property
func (m *InformationProtection) SetDataLossPreventionPolicies(value []DataLossPreventionPolicyable)() {
    err := m.GetBackingStore().Set("dataLossPreventionPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicy sets the policy property value. The policy property
func (m *InformationProtection) SetPolicy(value InformationProtectionPolicyable)() {
    err := m.GetBackingStore().Set("policy", value)
    if err != nil {
        panic(err)
    }
}
// SetSensitivityLabels sets the sensitivityLabels property value. The sensitivityLabels property
func (m *InformationProtection) SetSensitivityLabels(value []SensitivityLabelable)() {
    err := m.GetBackingStore().Set("sensitivityLabels", value)
    if err != nil {
        panic(err)
    }
}
// SetSensitivityPolicySettings sets the sensitivityPolicySettings property value. The sensitivityPolicySettings property
func (m *InformationProtection) SetSensitivityPolicySettings(value SensitivityPolicySettingsable)() {
    err := m.GetBackingStore().Set("sensitivityPolicySettings", value)
    if err != nil {
        panic(err)
    }
}
// SetThreatAssessmentRequests sets the threatAssessmentRequests property value. The threatAssessmentRequests property
func (m *InformationProtection) SetThreatAssessmentRequests(value []ThreatAssessmentRequestable)() {
    err := m.GetBackingStore().Set("threatAssessmentRequests", value)
    if err != nil {
        panic(err)
    }
}
// InformationProtectionable 
type InformationProtectionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBitlocker()(Bitlockerable)
    GetDataLossPreventionPolicies()([]DataLossPreventionPolicyable)
    GetPolicy()(InformationProtectionPolicyable)
    GetSensitivityLabels()([]SensitivityLabelable)
    GetSensitivityPolicySettings()(SensitivityPolicySettingsable)
    GetThreatAssessmentRequests()([]ThreatAssessmentRequestable)
    SetBitlocker(value Bitlockerable)()
    SetDataLossPreventionPolicies(value []DataLossPreventionPolicyable)()
    SetPolicy(value InformationProtectionPolicyable)()
    SetSensitivityLabels(value []SensitivityLabelable)()
    SetSensitivityPolicySettings(value SensitivityPolicySettingsable)()
    SetThreatAssessmentRequests(value []ThreatAssessmentRequestable)()
}
