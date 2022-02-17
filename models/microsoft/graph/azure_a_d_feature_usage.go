package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AzureADFeatureUsage 
type AzureADFeatureUsage struct {
    Entity
    // 
    featureName *string;
    // 
    snapshotDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    usage *int32;
}
// NewAzureADFeatureUsage instantiates a new AzureADFeatureUsage and sets the default values.
func NewAzureADFeatureUsage()(*AzureADFeatureUsage) {
    m := &AzureADFeatureUsage{
        Entity: *NewEntity(),
    }
    return m
}
// GetFeatureName gets the featureName property value. 
func (m *AzureADFeatureUsage) GetFeatureName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.featureName
    }
}
// GetSnapshotDateTime gets the snapshotDateTime property value. 
func (m *AzureADFeatureUsage) GetSnapshotDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.snapshotDateTime
    }
}
// GetUsage gets the usage property value. 
func (m *AzureADFeatureUsage) GetUsage()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.usage
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureADFeatureUsage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["featureName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureName(val)
        }
        return nil
    }
    res["snapshotDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSnapshotDateTime(val)
        }
        return nil
    }
    res["usage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsage(val)
        }
        return nil
    }
    return res
}
func (m *AzureADFeatureUsage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AzureADFeatureUsage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetFeatureName sets the featureName property value. 
func (m *AzureADFeatureUsage) SetFeatureName(value *string)() {
    if m != nil {
        m.featureName = value
    }
}
// SetSnapshotDateTime sets the snapshotDateTime property value. 
func (m *AzureADFeatureUsage) SetSnapshotDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.snapshotDateTime = value
    }
}
// SetUsage sets the usage property value. 
func (m *AzureADFeatureUsage) SetUsage(value *int32)() {
    if m != nil {
        m.usage = value
    }
}
