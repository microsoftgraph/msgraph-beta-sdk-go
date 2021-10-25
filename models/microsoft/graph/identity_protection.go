package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type IdentityProtection struct {
    additionalData map[string]interface{};
    riskDetections []RiskDetection;
    riskyUsers []RiskyUser;
}
func NewIdentityProtection()(*IdentityProtection) {
    m := &IdentityProtection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *IdentityProtection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *IdentityProtection) GetRiskDetections()([]RiskDetection) {
    if m == nil {
        return nil
    } else {
        return m.riskDetections
    }
}
func (m *IdentityProtection) GetRiskyUsers()([]RiskyUser) {
    if m == nil {
        return nil
    } else {
        return m.riskyUsers
    }
}
func (m *IdentityProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *IdentityProtection) IsNil()(bool) {
    return m == nil
}
func (m *IdentityProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *IdentityProtection) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *IdentityProtection) SetRiskDetections(value []RiskDetection)() {
    m.riskDetections = value
}
func (m *IdentityProtection) SetRiskyUsers(value []RiskyUser)() {
    m.riskyUsers = value
}
