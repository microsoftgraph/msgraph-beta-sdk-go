package getazureadfeatureusagewithperiod

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetAzureADFeatureUsageWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    featureName *string;
    snapshotDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    usage *int32;
}
func NewGetAzureADFeatureUsageWithPeriod()(*GetAzureADFeatureUsageWithPeriod) {
    m := &GetAzureADFeatureUsageWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetAzureADFeatureUsageWithPeriod) GetFeatureName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.featureName
    }
}
func (m *GetAzureADFeatureUsageWithPeriod) GetSnapshotDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.snapshotDateTime
    }
}
func (m *GetAzureADFeatureUsageWithPeriod) GetUsage()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.usage
    }
}
func (m *GetAzureADFeatureUsageWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["featureName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFeatureName(val)
        return nil
    }
    res["snapshotDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetSnapshotDateTime(val)
        return nil
    }
    res["usage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUsage(val)
        return nil
    }
    return res
}
func (m *GetAzureADFeatureUsageWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetAzureADFeatureUsageWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("featureName", m.GetFeatureName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("snapshotDateTime", m.GetSnapshotDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("usage", m.GetUsage())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetAzureADFeatureUsageWithPeriod) SetFeatureName(value *string)() {
    m.featureName = value
}
func (m *GetAzureADFeatureUsageWithPeriod) SetSnapshotDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.snapshotDateTime = value
}
func (m *GetAzureADFeatureUsageWithPeriod) SetUsage(value *int32)() {
    m.usage = value
}
