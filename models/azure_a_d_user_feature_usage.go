package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureADUserFeatureUsage 
type AzureADUserFeatureUsage struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The featureUsageDetails property
    featureUsageDetails []FeatureUsageDetailable
    // The lastUpdatedDateTime property
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The licenseAssigned property
    licenseAssigned *AzureADLicenseType
    // The licenseRecommended property
    licenseRecommended *AzureADLicenseType
    // The userDisplayName property
    userDisplayName *string
    // The userId property
    userId *string
    // The userPrincipalName property
    userPrincipalName *string
}
// NewAzureADUserFeatureUsage instantiates a new AzureADUserFeatureUsage and sets the default values.
func NewAzureADUserFeatureUsage()(*AzureADUserFeatureUsage) {
    m := &AzureADUserFeatureUsage{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAzureADUserFeatureUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureADUserFeatureUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureADUserFeatureUsage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AzureADUserFeatureUsage) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFeatureUsageDetails gets the featureUsageDetails property value. The featureUsageDetails property
func (m *AzureADUserFeatureUsage) GetFeatureUsageDetails()([]FeatureUsageDetailable) {
    if m == nil {
        return nil
    } else {
        return m.featureUsageDetails
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureADUserFeatureUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["featureUsageDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFeatureUsageDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FeatureUsageDetailable, len(val))
            for i, v := range val {
                res[i] = v.(FeatureUsageDetailable)
            }
            m.SetFeatureUsageDetails(res)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
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
    res["licenseRecommended"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAzureADLicenseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseRecommended(val.(*AzureADLicenseType))
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
func (m *AzureADUserFeatureUsage) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// GetLicenseAssigned gets the licenseAssigned property value. The licenseAssigned property
func (m *AzureADUserFeatureUsage) GetLicenseAssigned()(*AzureADLicenseType) {
    if m == nil {
        return nil
    } else {
        return m.licenseAssigned
    }
}
// GetLicenseRecommended gets the licenseRecommended property value. The licenseRecommended property
func (m *AzureADUserFeatureUsage) GetLicenseRecommended()(*AzureADLicenseType) {
    if m == nil {
        return nil
    } else {
        return m.licenseRecommended
    }
}
// GetUserDisplayName gets the userDisplayName property value. The userDisplayName property
func (m *AzureADUserFeatureUsage) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// GetUserId gets the userId property value. The userId property
func (m *AzureADUserFeatureUsage) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *AzureADUserFeatureUsage) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Serialize serializes information the current object
func (m *AzureADUserFeatureUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetFeatureUsageDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFeatureUsageDetails()))
        for i, v := range m.GetFeatureUsageDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("featureUsageDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLicenseAssigned() != nil {
        cast := (*m.GetLicenseAssigned()).String()
        err = writer.WriteStringValue("licenseAssigned", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLicenseRecommended() != nil {
        cast := (*m.GetLicenseRecommended()).String()
        err = writer.WriteStringValue("licenseRecommended", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AzureADUserFeatureUsage) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFeatureUsageDetails sets the featureUsageDetails property value. The featureUsageDetails property
func (m *AzureADUserFeatureUsage) SetFeatureUsageDetails(value []FeatureUsageDetailable)() {
    if m != nil {
        m.featureUsageDetails = value
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
func (m *AzureADUserFeatureUsage) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdatedDateTime = value
    }
}
// SetLicenseAssigned sets the licenseAssigned property value. The licenseAssigned property
func (m *AzureADUserFeatureUsage) SetLicenseAssigned(value *AzureADLicenseType)() {
    if m != nil {
        m.licenseAssigned = value
    }
}
// SetLicenseRecommended sets the licenseRecommended property value. The licenseRecommended property
func (m *AzureADUserFeatureUsage) SetLicenseRecommended(value *AzureADLicenseType)() {
    if m != nil {
        m.licenseRecommended = value
    }
}
// SetUserDisplayName sets the userDisplayName property value. The userDisplayName property
func (m *AzureADUserFeatureUsage) SetUserDisplayName(value *string)() {
    if m != nil {
        m.userDisplayName = value
    }
}
// SetUserId sets the userId property value. The userId property
func (m *AzureADUserFeatureUsage) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *AzureADUserFeatureUsage) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
