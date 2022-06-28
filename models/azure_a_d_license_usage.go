package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureADLicenseUsage 
type AzureADLicenseUsage struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The licenseInfoDetails property
    licenseInfoDetails []LicenseInfoDetailable
    // The snapshotDateTime property
    snapshotDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewAzureADLicenseUsage instantiates a new AzureADLicenseUsage and sets the default values.
func NewAzureADLicenseUsage()(*AzureADLicenseUsage) {
    m := &AzureADLicenseUsage{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAzureADLicenseUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureADLicenseUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureADLicenseUsage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AzureADLicenseUsage) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureADLicenseUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["licenseInfoDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLicenseInfoDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LicenseInfoDetailable, len(val))
            for i, v := range val {
                res[i] = v.(LicenseInfoDetailable)
            }
            m.SetLicenseInfoDetails(res)
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
    return res
}
// GetLicenseInfoDetails gets the licenseInfoDetails property value. The licenseInfoDetails property
func (m *AzureADLicenseUsage) GetLicenseInfoDetails()([]LicenseInfoDetailable) {
    if m == nil {
        return nil
    } else {
        return m.licenseInfoDetails
    }
}
// GetSnapshotDateTime gets the snapshotDateTime property value. The snapshotDateTime property
func (m *AzureADLicenseUsage) GetSnapshotDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.snapshotDateTime
    }
}
// Serialize serializes information the current object
func (m *AzureADLicenseUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLicenseInfoDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLicenseInfoDetails()))
        for i, v := range m.GetLicenseInfoDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("licenseInfoDetails", cast)
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
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AzureADLicenseUsage) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLicenseInfoDetails sets the licenseInfoDetails property value. The licenseInfoDetails property
func (m *AzureADLicenseUsage) SetLicenseInfoDetails(value []LicenseInfoDetailable)() {
    if m != nil {
        m.licenseInfoDetails = value
    }
}
// SetSnapshotDateTime sets the snapshotDateTime property value. The snapshotDateTime property
func (m *AzureADLicenseUsage) SetSnapshotDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.snapshotDateTime = value
    }
}
