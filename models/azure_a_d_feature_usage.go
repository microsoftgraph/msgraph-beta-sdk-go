package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureADFeatureUsage 
type AzureADFeatureUsage struct {
    Entity
    // The featureName property
    featureName *string
    // The snapshotDateTime property
    snapshotDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The usage property
    usage *int32
}
// NewAzureADFeatureUsage instantiates a new AzureADFeatureUsage and sets the default values.
func NewAzureADFeatureUsage()(*AzureADFeatureUsage) {
    m := &AzureADFeatureUsage{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.azureADFeatureUsage";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAzureADFeatureUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureADFeatureUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureADFeatureUsage(), nil
}
// GetFeatureName gets the featureName property value. The featureName property
func (m *AzureADFeatureUsage) GetFeatureName()(*string) {
    return m.featureName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureADFeatureUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["featureName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureName(val)
        }
        return nil
    }
    res["snapshotDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSnapshotDateTime(val)
        }
        return nil
    }
    res["usage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetSnapshotDateTime gets the snapshotDateTime property value. The snapshotDateTime property
func (m *AzureADFeatureUsage) GetSnapshotDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.snapshotDateTime
}
// GetUsage gets the usage property value. The usage property
func (m *AzureADFeatureUsage) GetUsage()(*int32) {
    return m.usage
}
// Serialize serializes information the current object
func (m *AzureADFeatureUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetFeatureName sets the featureName property value. The featureName property
func (m *AzureADFeatureUsage) SetFeatureName(value *string)() {
    m.featureName = value
}
// SetSnapshotDateTime sets the snapshotDateTime property value. The snapshotDateTime property
func (m *AzureADFeatureUsage) SetSnapshotDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.snapshotDateTime = value
}
// SetUsage sets the usage property value. The usage property
func (m *AzureADFeatureUsage) SetUsage(value *int32)() {
    m.usage = value
}
