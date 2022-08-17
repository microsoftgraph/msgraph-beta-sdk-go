package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FeatureUsageDetail 
type FeatureUsageDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The featureName property
    featureName *string
    // The lastConfiguredDateTime property
    lastConfiguredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lastUsedDateTime property
    lastUsedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The licenseAssigned property
    licenseAssigned *AzureADLicenseType
    // The licenseRequired property
    licenseRequired *AzureADLicenseType
    // The OdataType property
    odataType *string
}
// NewFeatureUsageDetail instantiates a new featureUsageDetail and sets the default values.
func NewFeatureUsageDetail()(*FeatureUsageDetail) {
    m := &FeatureUsageDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.featureUsageDetail";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateFeatureUsageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFeatureUsageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFeatureUsageDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FeatureUsageDetail) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFeatureName gets the featureName property value. The featureName property
func (m *FeatureUsageDetail) GetFeatureName()(*string) {
    return m.featureName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FeatureUsageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["lastConfiguredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastConfiguredDateTime(val)
        }
        return nil
    }
    res["lastUsedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUsedDateTime(val)
        }
        return nil
    }
    res["licenseAssigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAzureADLicenseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseAssigned(val.(*AzureADLicenseType))
        }
        return nil
    }
    res["licenseRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAzureADLicenseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseRequired(val.(*AzureADLicenseType))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetLastConfiguredDateTime gets the lastConfiguredDateTime property value. The lastConfiguredDateTime property
func (m *FeatureUsageDetail) GetLastConfiguredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastConfiguredDateTime
}
// GetLastUsedDateTime gets the lastUsedDateTime property value. The lastUsedDateTime property
func (m *FeatureUsageDetail) GetLastUsedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUsedDateTime
}
// GetLicenseAssigned gets the licenseAssigned property value. The licenseAssigned property
func (m *FeatureUsageDetail) GetLicenseAssigned()(*AzureADLicenseType) {
    return m.licenseAssigned
}
// GetLicenseRequired gets the licenseRequired property value. The licenseRequired property
func (m *FeatureUsageDetail) GetLicenseRequired()(*AzureADLicenseType) {
    return m.licenseRequired
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *FeatureUsageDetail) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *FeatureUsageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("featureName", m.GetFeatureName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastConfiguredDateTime", m.GetLastConfiguredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastUsedDateTime", m.GetLastUsedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLicenseAssigned() != nil {
        cast := (*m.GetLicenseAssigned()).String()
        err := writer.WriteStringValue("licenseAssigned", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLicenseRequired() != nil {
        cast := (*m.GetLicenseRequired()).String()
        err := writer.WriteStringValue("licenseRequired", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FeatureUsageDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFeatureName sets the featureName property value. The featureName property
func (m *FeatureUsageDetail) SetFeatureName(value *string)() {
    m.featureName = value
}
// SetLastConfiguredDateTime sets the lastConfiguredDateTime property value. The lastConfiguredDateTime property
func (m *FeatureUsageDetail) SetLastConfiguredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastConfiguredDateTime = value
}
// SetLastUsedDateTime sets the lastUsedDateTime property value. The lastUsedDateTime property
func (m *FeatureUsageDetail) SetLastUsedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUsedDateTime = value
}
// SetLicenseAssigned sets the licenseAssigned property value. The licenseAssigned property
func (m *FeatureUsageDetail) SetLicenseAssigned(value *AzureADLicenseType)() {
    m.licenseAssigned = value
}
// SetLicenseRequired sets the licenseRequired property value. The licenseRequired property
func (m *FeatureUsageDetail) SetLicenseRequired(value *AzureADLicenseType)() {
    m.licenseRequired = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *FeatureUsageDetail) SetOdataType(value *string)() {
    m.odataType = value
}
