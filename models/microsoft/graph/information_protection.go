package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InformationProtection 
type InformationProtection struct {
    Entity
    // 
    bitlocker *Bitlocker;
    // 
    dataLossPreventionPolicies []DataLossPreventionPolicy;
    // 
    policy *InformationProtectionPolicy;
    // 
    sensitivityLabels []SensitivityLabel;
    // 
    sensitivityPolicySettings *SensitivityPolicySettings;
    // 
    threatAssessmentRequests []ThreatAssessmentRequest;
}
// NewInformationProtection instantiates a new informationProtection and sets the default values.
func NewInformationProtection()(*InformationProtection) {
    m := &InformationProtection{
        Entity: *NewEntity(),
    }
    return m
}
// GetBitlocker gets the bitlocker property value. 
func (m *InformationProtection) GetBitlocker()(*Bitlocker) {
    if m == nil {
        return nil
    } else {
        return m.bitlocker
    }
}
// GetDataLossPreventionPolicies gets the dataLossPreventionPolicies property value. 
func (m *InformationProtection) GetDataLossPreventionPolicies()([]DataLossPreventionPolicy) {
    if m == nil {
        return nil
    } else {
        return m.dataLossPreventionPolicies
    }
}
// GetPolicy gets the policy property value. 
func (m *InformationProtection) GetPolicy()(*InformationProtectionPolicy) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
// GetSensitivityLabels gets the sensitivityLabels property value. 
func (m *InformationProtection) GetSensitivityLabels()([]SensitivityLabel) {
    if m == nil {
        return nil
    } else {
        return m.sensitivityLabels
    }
}
// GetSensitivityPolicySettings gets the sensitivityPolicySettings property value. 
func (m *InformationProtection) GetSensitivityPolicySettings()(*SensitivityPolicySettings) {
    if m == nil {
        return nil
    } else {
        return m.sensitivityPolicySettings
    }
}
// GetThreatAssessmentRequests gets the threatAssessmentRequests property value. 
func (m *InformationProtection) GetThreatAssessmentRequests()([]ThreatAssessmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.threatAssessmentRequests
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bitlocker"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBitlocker() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitlocker(val.(*Bitlocker))
        }
        return nil
    }
    res["dataLossPreventionPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDataLossPreventionPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DataLossPreventionPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*DataLossPreventionPolicy))
            }
            m.SetDataLossPreventionPolicies(res)
        }
        return nil
    }
    res["policy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInformationProtectionPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val.(*InformationProtectionPolicy))
        }
        return nil
    }
    res["sensitivityLabels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitivityLabel() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitivityLabel, len(val))
            for i, v := range val {
                res[i] = *(v.(*SensitivityLabel))
            }
            m.SetSensitivityLabels(res)
        }
        return nil
    }
    res["sensitivityPolicySettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitivityPolicySettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivityPolicySettings(val.(*SensitivityPolicySettings))
        }
        return nil
    }
    res["threatAssessmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThreatAssessmentRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ThreatAssessmentRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*ThreatAssessmentRequest))
            }
            m.SetThreatAssessmentRequests(res)
        }
        return nil
    }
    return res
}
func (m *InformationProtection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InformationProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDataLossPreventionPolicies()))
        for i, v := range m.GetDataLossPreventionPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSensitivityLabels()))
        for i, v := range m.GetSensitivityLabels() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetThreatAssessmentRequests()))
        for i, v := range m.GetThreatAssessmentRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("threatAssessmentRequests", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBitlocker sets the bitlocker property value. 
func (m *InformationProtection) SetBitlocker(value *Bitlocker)() {
    if m != nil {
        m.bitlocker = value
    }
}
// SetDataLossPreventionPolicies sets the dataLossPreventionPolicies property value. 
func (m *InformationProtection) SetDataLossPreventionPolicies(value []DataLossPreventionPolicy)() {
    if m != nil {
        m.dataLossPreventionPolicies = value
    }
}
// SetPolicy sets the policy property value. 
func (m *InformationProtection) SetPolicy(value *InformationProtectionPolicy)() {
    if m != nil {
        m.policy = value
    }
}
// SetSensitivityLabels sets the sensitivityLabels property value. 
func (m *InformationProtection) SetSensitivityLabels(value []SensitivityLabel)() {
    if m != nil {
        m.sensitivityLabels = value
    }
}
// SetSensitivityPolicySettings sets the sensitivityPolicySettings property value. 
func (m *InformationProtection) SetSensitivityPolicySettings(value *SensitivityPolicySettings)() {
    if m != nil {
        m.sensitivityPolicySettings = value
    }
}
// SetThreatAssessmentRequests sets the threatAssessmentRequests property value. 
func (m *InformationProtection) SetThreatAssessmentRequests(value []ThreatAssessmentRequest)() {
    if m != nil {
        m.threatAssessmentRequests = value
    }
}
