package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IdentityProtectionRoot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    riskDetections []RiskDetection;
    // 
    riskyUsers []RiskyUser;
}
// Instantiates a new IdentityProtectionRoot and sets the default values.
func NewIdentityProtectionRoot()(*IdentityProtectionRoot) {
    m := &IdentityProtectionRoot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityProtectionRoot) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the riskDetections property value. 
func (m *IdentityProtectionRoot) GetRiskDetections()([]RiskDetection) {
    if m == nil {
        return nil
    } else {
        return m.riskDetections
    }
}
// Gets the riskyUsers property value. 
func (m *IdentityProtectionRoot) GetRiskyUsers()([]RiskyUser) {
    if m == nil {
        return nil
    } else {
        return m.riskyUsers
    }
}
// The deserialization information for the current model
func (m *IdentityProtectionRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["riskDetections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRiskDetection() })
        if err != nil {
            return err
        }
        res := make([]RiskDetection, len(val))
        for i, v := range val {
            res[i] = *(v.(*RiskDetection))
        }
        m.SetRiskDetections(res)
        return nil
    }
    res["riskyUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRiskyUser() })
        if err != nil {
            return err
        }
        res := make([]RiskyUser, len(val))
        for i, v := range val {
            res[i] = *(v.(*RiskyUser))
        }
        m.SetRiskyUsers(res)
        return nil
    }
    return res
}
func (m *IdentityProtectionRoot) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IdentityProtectionRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRiskDetections()))
        for i, v := range m.GetRiskDetections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("riskDetections", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRiskyUsers()))
        for i, v := range m.GetRiskyUsers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("riskyUsers", cast)
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *IdentityProtectionRoot) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the riskDetections property value. 
// Parameters:
//  - value : Value to set for the riskDetections property.
func (m *IdentityProtectionRoot) SetRiskDetections(value []RiskDetection)() {
    m.riskDetections = value
}
// Sets the riskyUsers property value. 
// Parameters:
//  - value : Value to set for the riskyUsers property.
func (m *IdentityProtectionRoot) SetRiskyUsers(value []RiskyUser)() {
    m.riskyUsers = value
}
