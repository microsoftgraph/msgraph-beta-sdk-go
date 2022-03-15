package security

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// InformationProtection provides operations to manage the deviceManagement singleton.
type InformationProtection struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    labelPolicySettings InformationProtectionPolicySettingable;
    // 
    sensitivityLabels []SensitivityLabelable;
}
// NewInformationProtection instantiates a new informationProtection and sets the default values.
func NewInformationProtection()(*InformationProtection) {
    m := &InformationProtection{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// CreateInformationProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewInformationProtection(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["labelPolicySettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateInformationProtectionPolicySettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelPolicySettings(val.(InformationProtectionPolicySettingable))
        }
        return nil
    }
    res["sensitivityLabels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitivityLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitivityLabelable, len(val))
            for i, v := range val {
                res[i] = v.(SensitivityLabelable)
            }
            m.SetSensitivityLabels(res)
        }
        return nil
    }
    return res
}
// GetLabelPolicySettings gets the labelPolicySettings property value. 
func (m *InformationProtection) GetLabelPolicySettings()(InformationProtectionPolicySettingable) {
    if m == nil {
        return nil
    } else {
        return m.labelPolicySettings
    }
}
// GetSensitivityLabels gets the sensitivityLabels property value. 
func (m *InformationProtection) GetSensitivityLabels()([]SensitivityLabelable) {
    if m == nil {
        return nil
    } else {
        return m.sensitivityLabels
    }
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
        err = writer.WriteObjectValue("labelPolicySettings", m.GetLabelPolicySettings())
        if err != nil {
            return err
        }
    }
    if m.GetSensitivityLabels() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSensitivityLabels()))
        for i, v := range m.GetSensitivityLabels() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sensitivityLabels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLabelPolicySettings sets the labelPolicySettings property value. 
func (m *InformationProtection) SetLabelPolicySettings(value InformationProtectionPolicySettingable)() {
    if m != nil {
        m.labelPolicySettings = value
    }
}
// SetSensitivityLabels sets the sensitivityLabels property value. 
func (m *InformationProtection) SetSensitivityLabels(value []SensitivityLabelable)() {
    if m != nil {
        m.sensitivityLabels = value
    }
}
