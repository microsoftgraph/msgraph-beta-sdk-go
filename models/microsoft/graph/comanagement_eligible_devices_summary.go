package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ComanagementEligibleDevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Count of devices already Co-Managed
    comanagedCount *int32;
    // Count of devices eligible for Co-Management but not yet joined to Azure Active Directory
    eligibleButNotAzureAdJoinedCount *int32;
    // Count of devices fully eligible for Co-Management
    eligibleCount *int32;
    // Count of devices ineligible for Co-Management
    ineligibleCount *int32;
    // Count of devices that will be eligible for Co-Management after an OS update
    needsOsUpdateCount *int32;
}
// Instantiates a new comanagementEligibleDevicesSummary and sets the default values.
func NewComanagementEligibleDevicesSummary()(*ComanagementEligibleDevicesSummary) {
    m := &ComanagementEligibleDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagementEligibleDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the comanagedCount property value. Count of devices already Co-Managed
func (m *ComanagementEligibleDevicesSummary) GetComanagedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.comanagedCount
    }
}
// Gets the eligibleButNotAzureAdJoinedCount property value. Count of devices eligible for Co-Management but not yet joined to Azure Active Directory
func (m *ComanagementEligibleDevicesSummary) GetEligibleButNotAzureAdJoinedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.eligibleButNotAzureAdJoinedCount
    }
}
// Gets the eligibleCount property value. Count of devices fully eligible for Co-Management
func (m *ComanagementEligibleDevicesSummary) GetEligibleCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.eligibleCount
    }
}
// Gets the ineligibleCount property value. Count of devices ineligible for Co-Management
func (m *ComanagementEligibleDevicesSummary) GetIneligibleCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.ineligibleCount
    }
}
// Gets the needsOsUpdateCount property value. Count of devices that will be eligible for Co-Management after an OS update
func (m *ComanagementEligibleDevicesSummary) GetNeedsOsUpdateCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.needsOsUpdateCount
    }
}
// The deserialization information for the current model
func (m *ComanagementEligibleDevicesSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["comanagedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetComanagedCount(val)
        return nil
    }
    res["eligibleButNotAzureAdJoinedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetEligibleButNotAzureAdJoinedCount(val)
        return nil
    }
    res["eligibleCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetEligibleCount(val)
        return nil
    }
    res["ineligibleCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIneligibleCount(val)
        return nil
    }
    res["needsOsUpdateCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNeedsOsUpdateCount(val)
        return nil
    }
    return res
}
func (m *ComanagementEligibleDevicesSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ComanagementEligibleDevicesSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("comanagedCount", m.GetComanagedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("eligibleButNotAzureAdJoinedCount", m.GetEligibleButNotAzureAdJoinedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("eligibleCount", m.GetEligibleCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ineligibleCount", m.GetIneligibleCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("needsOsUpdateCount", m.GetNeedsOsUpdateCount())
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
func (m *ComanagementEligibleDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the comanagedCount property value. Count of devices already Co-Managed
// Parameters:
//  - value : Value to set for the comanagedCount property.
func (m *ComanagementEligibleDevicesSummary) SetComanagedCount(value *int32)() {
    m.comanagedCount = value
}
// Sets the eligibleButNotAzureAdJoinedCount property value. Count of devices eligible for Co-Management but not yet joined to Azure Active Directory
// Parameters:
//  - value : Value to set for the eligibleButNotAzureAdJoinedCount property.
func (m *ComanagementEligibleDevicesSummary) SetEligibleButNotAzureAdJoinedCount(value *int32)() {
    m.eligibleButNotAzureAdJoinedCount = value
}
// Sets the eligibleCount property value. Count of devices fully eligible for Co-Management
// Parameters:
//  - value : Value to set for the eligibleCount property.
func (m *ComanagementEligibleDevicesSummary) SetEligibleCount(value *int32)() {
    m.eligibleCount = value
}
// Sets the ineligibleCount property value. Count of devices ineligible for Co-Management
// Parameters:
//  - value : Value to set for the ineligibleCount property.
func (m *ComanagementEligibleDevicesSummary) SetIneligibleCount(value *int32)() {
    m.ineligibleCount = value
}
// Sets the needsOsUpdateCount property value. Count of devices that will be eligible for Co-Management after an OS update
// Parameters:
//  - value : Value to set for the needsOsUpdateCount property.
func (m *ComanagementEligibleDevicesSummary) SetNeedsOsUpdateCount(value *int32)() {
    m.needsOsUpdateCount = value
}
