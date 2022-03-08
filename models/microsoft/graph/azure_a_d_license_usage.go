package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AzureADLicenseUsage provides operations to call the getAzureADLicenseUsage method.
type AzureADLicenseUsage struct {
    Entity
    // 
    licenseInfoDetails []LicenseInfoDetailable;
    // 
    snapshotDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewAzureADLicenseUsage instantiates a new azureADLicenseUsage and sets the default values.
func NewAzureADLicenseUsage()(*AzureADLicenseUsage) {
    m := &AzureADLicenseUsage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAzureADLicenseUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureADLicenseUsageFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAzureADLicenseUsage(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureADLicenseUsage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["licenseInfoDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    return res
}
// GetLicenseInfoDetails gets the licenseInfoDetails property value. 
func (m *AzureADLicenseUsage) GetLicenseInfoDetails()([]LicenseInfoDetailable) {
    if m == nil {
        return nil
    } else {
        return m.licenseInfoDetails
    }
}
// GetSnapshotDateTime gets the snapshotDateTime property value. 
func (m *AzureADLicenseUsage) GetSnapshotDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.snapshotDateTime
    }
}
func (m *AzureADLicenseUsage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AzureADLicenseUsage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLicenseInfoDetails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLicenseInfoDetails()))
        for i, v := range m.GetLicenseInfoDetails() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    return nil
}
// SetLicenseInfoDetails sets the licenseInfoDetails property value. 
func (m *AzureADLicenseUsage) SetLicenseInfoDetails(value []LicenseInfoDetailable)() {
    if m != nil {
        m.licenseInfoDetails = value
    }
}
// SetSnapshotDateTime sets the snapshotDateTime property value. 
func (m *AzureADLicenseUsage) SetSnapshotDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.snapshotDateTime = value
    }
}
