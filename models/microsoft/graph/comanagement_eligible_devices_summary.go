package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ComanagementEligibleDevicesSummary struct {
    additionalData map[string]interface{};
    comanagedCount *int32;
    eligibleButNotAzureAdJoinedCount *int32;
    eligibleCount *int32;
    ineligibleCount *int32;
    needsOsUpdateCount *int32;
}
func NewComanagementEligibleDevicesSummary()(*ComanagementEligibleDevicesSummary) {
    m := &ComanagementEligibleDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ComanagementEligibleDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ComanagementEligibleDevicesSummary) GetComanagedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.comanagedCount
    }
}
func (m *ComanagementEligibleDevicesSummary) GetEligibleButNotAzureAdJoinedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.eligibleButNotAzureAdJoinedCount
    }
}
func (m *ComanagementEligibleDevicesSummary) GetEligibleCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.eligibleCount
    }
}
func (m *ComanagementEligibleDevicesSummary) GetIneligibleCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.ineligibleCount
    }
}
func (m *ComanagementEligibleDevicesSummary) GetNeedsOsUpdateCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.needsOsUpdateCount
    }
}
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
func (m *ComanagementEligibleDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ComanagementEligibleDevicesSummary) SetComanagedCount(value *int32)() {
    m.comanagedCount = value
}
func (m *ComanagementEligibleDevicesSummary) SetEligibleButNotAzureAdJoinedCount(value *int32)() {
    m.eligibleButNotAzureAdJoinedCount = value
}
func (m *ComanagementEligibleDevicesSummary) SetEligibleCount(value *int32)() {
    m.eligibleCount = value
}
func (m *ComanagementEligibleDevicesSummary) SetIneligibleCount(value *int32)() {
    m.ineligibleCount = value
}
func (m *ComanagementEligibleDevicesSummary) SetNeedsOsUpdateCount(value *int32)() {
    m.needsOsUpdateCount = value
}
