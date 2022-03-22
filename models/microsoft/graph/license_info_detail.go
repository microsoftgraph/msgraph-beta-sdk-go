package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LicenseInfoDetail 
type LicenseInfoDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    licenseType *AzureADLicenseType;
    // 
    totalAssignedCount *int32;
    // 
    totalLicenseCount *int32;
    // 
    totalUsageCount *int32;
}
// NewLicenseInfoDetail instantiates a new licenseInfoDetail and sets the default values.
func NewLicenseInfoDetail()(*LicenseInfoDetail) {
    m := &LicenseInfoDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLicenseInfoDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLicenseInfoDetailFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
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
func (m *LicenseInfoDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["licenseType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAzureADLicenseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseType(val.(*AzureADLicenseType))
        }
        return nil
    }
    res["totalAssignedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAssignedCount(val)
        }
        return nil
    }
    res["totalLicenseCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLicenseCount(val)
        }
        return nil
    }
    res["totalUsageCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
// GetLicenseType gets the licenseType property value. 
func (m *LicenseInfoDetail) GetLicenseType()(*AzureADLicenseType) {
    if m == nil {
        return nil
    } else {
        return m.licenseType
    }
}
// GetTotalAssignedCount gets the totalAssignedCount property value. 
func (m *LicenseInfoDetail) GetTotalAssignedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalAssignedCount
    }
}
// GetTotalLicenseCount gets the totalLicenseCount property value. 
func (m *LicenseInfoDetail) GetTotalLicenseCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalLicenseCount
    }
}
// GetTotalUsageCount gets the totalUsageCount property value. 
func (m *LicenseInfoDetail) GetTotalUsageCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalUsageCount
    }
}
// Serialize serializes information the current object
func (m *LicenseInfoDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetLicenseType sets the licenseType property value. 
func (m *LicenseInfoDetail) SetLicenseType(value *AzureADLicenseType)() {
    if m != nil {
        m.licenseType = value
    }
}
// SetTotalAssignedCount sets the totalAssignedCount property value. 
func (m *LicenseInfoDetail) SetTotalAssignedCount(value *int32)() {
    if m != nil {
        m.totalAssignedCount = value
    }
}
// SetTotalLicenseCount sets the totalLicenseCount property value. 
func (m *LicenseInfoDetail) SetTotalLicenseCount(value *int32)() {
    if m != nil {
        m.totalLicenseCount = value
    }
}
// SetTotalUsageCount sets the totalUsageCount property value. 
func (m *LicenseInfoDetail) SetTotalUsageCount(value *int32)() {
    if m != nil {
        m.totalUsageCount = value
    }
}
