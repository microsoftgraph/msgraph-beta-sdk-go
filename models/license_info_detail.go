package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LicenseInfoDetail 
type LicenseInfoDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The licenseType property
    licenseType *AzureADLicenseType
    // The totalAssignedCount property
    totalAssignedCount *int32
    // The totalLicenseCount property
    totalLicenseCount *int32
    // The totalUsageCount property
    totalUsageCount *int32
}
// NewLicenseInfoDetail instantiates a new licenseInfoDetail and sets the default values.
func NewLicenseInfoDetail()(*LicenseInfoDetail) {
    m := &LicenseInfoDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLicenseInfoDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLicenseInfoDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLicenseInfoDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LicenseInfoDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LicenseInfoDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["licenseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAzureADLicenseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseType(val.(*AzureADLicenseType))
        }
        return nil
    }
    res["totalAssignedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAssignedCount(val)
        }
        return nil
    }
    res["totalLicenseCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLicenseCount(val)
        }
        return nil
    }
    res["totalUsageCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUsageCount(val)
        }
        return nil
    }
    return res
}
// GetLicenseType gets the licenseType property value. The licenseType property
func (m *LicenseInfoDetail) GetLicenseType()(*AzureADLicenseType) {
    if m == nil {
        return nil
    } else {
        return m.licenseType
    }
}
// GetTotalAssignedCount gets the totalAssignedCount property value. The totalAssignedCount property
func (m *LicenseInfoDetail) GetTotalAssignedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalAssignedCount
    }
}
// GetTotalLicenseCount gets the totalLicenseCount property value. The totalLicenseCount property
func (m *LicenseInfoDetail) GetTotalLicenseCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalLicenseCount
    }
}
// GetTotalUsageCount gets the totalUsageCount property value. The totalUsageCount property
func (m *LicenseInfoDetail) GetTotalUsageCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalUsageCount
    }
}
// Serialize serializes information the current object
func (m *LicenseInfoDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLicenseType() != nil {
        cast := (*m.GetLicenseType()).String()
        err := writer.WriteStringValue("licenseType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalAssignedCount", m.GetTotalAssignedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalLicenseCount", m.GetTotalLicenseCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalUsageCount", m.GetTotalUsageCount())
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
func (m *LicenseInfoDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLicenseType sets the licenseType property value. The licenseType property
func (m *LicenseInfoDetail) SetLicenseType(value *AzureADLicenseType)() {
    if m != nil {
        m.licenseType = value
    }
}
// SetTotalAssignedCount sets the totalAssignedCount property value. The totalAssignedCount property
func (m *LicenseInfoDetail) SetTotalAssignedCount(value *int32)() {
    if m != nil {
        m.totalAssignedCount = value
    }
}
// SetTotalLicenseCount sets the totalLicenseCount property value. The totalLicenseCount property
func (m *LicenseInfoDetail) SetTotalLicenseCount(value *int32)() {
    if m != nil {
        m.totalLicenseCount = value
    }
}
// SetTotalUsageCount sets the totalUsageCount property value. The totalUsageCount property
func (m *LicenseInfoDetail) SetTotalUsageCount(value *int32)() {
    if m != nil {
        m.totalUsageCount = value
    }
}
