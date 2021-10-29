package getonedriveusageaccountdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetOneDriveUsageAccountDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    activeFileCount *int64;
    // 
    fileCount *int64;
    // 
    isDeleted *bool;
    // 
    lastActivityDate *string;
    // 
    ownerDisplayName *string;
    // 
    ownerPrincipalName *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    siteUrl *string;
    // 
    storageAllocatedInBytes *int64;
    // 
    storageUsedInBytes *int64;
}
// Instantiates a new getOneDriveUsageAccountDetailWithPeriod and sets the default values.
func NewGetOneDriveUsageAccountDetailWithPeriod()(*GetOneDriveUsageAccountDetailWithPeriod) {
    m := &GetOneDriveUsageAccountDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the activeFileCount property value. 
func (m *GetOneDriveUsageAccountDetailWithPeriod) GetActiveFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.activeFileCount
    }
}
// Gets the fileCount property value. 
func (m *GetOneDriveUsageAccountDetailWithPeriod) GetFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileCount
    }
}
// Gets the isDeleted property value. 
func (m *GetOneDriveUsageAccountDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// Gets the lastActivityDate property value. 
func (m *GetOneDriveUsageAccountDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the ownerDisplayName property value. 
func (m *GetOneDriveUsageAccountDetailWithPeriod) GetOwnerDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerDisplayName
    }
}
// Gets the ownerPrincipalName property value. 
func (m *GetOneDriveUsageAccountDetailWithPeriod) GetOwnerPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerPrincipalName
    }
}
// Gets the reportPeriod property value. 
func (m *GetOneDriveUsageAccountDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetOneDriveUsageAccountDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the siteUrl property value. 
func (m *GetOneDriveUsageAccountDetailWithPeriod) GetSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteUrl
    }
}
// Gets the storageAllocatedInBytes property value. 
func (m *GetOneDriveUsageAccountDetailWithPeriod) GetStorageAllocatedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageAllocatedInBytes
    }
}
// Gets the storageUsedInBytes property value. 
func (m *GetOneDriveUsageAccountDetailWithPeriod) GetStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageUsedInBytes
    }
}
// The deserialization information for the current model
func (m *GetOneDriveUsageAccountDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetActiveFileCount(val)
        return nil
    }
    res["fileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetFileCount(val)
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeleted(val)
        return nil
    }
    res["lastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastActivityDate(val)
        return nil
    }
    res["ownerDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOwnerDisplayName(val)
        return nil
    }
    res["ownerPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOwnerPrincipalName(val)
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportPeriod(val)
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportRefreshDate(val)
        return nil
    }
    res["siteUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSiteUrl(val)
        return nil
    }
    res["storageAllocatedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetStorageAllocatedInBytes(val)
        return nil
    }
    res["storageUsedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetStorageUsedInBytes(val)
        return nil
    }
    return res
}
func (m *GetOneDriveUsageAccountDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetOneDriveUsageAccountDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("activeFileCount", m.GetActiveFileCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("fileCount", m.GetFileCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastActivityDate", m.GetLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerDisplayName", m.GetOwnerDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerPrincipalName", m.GetOwnerPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("siteUrl", m.GetSiteUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("storageAllocatedInBytes", m.GetStorageAllocatedInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("storageUsedInBytes", m.GetStorageUsedInBytes())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the activeFileCount property value. 
// Parameters:
//  - value : Value to set for the activeFileCount property.
func (m *GetOneDriveUsageAccountDetailWithPeriod) SetActiveFileCount(value *int64)() {
    m.activeFileCount = value
}
// Sets the fileCount property value. 
// Parameters:
//  - value : Value to set for the fileCount property.
func (m *GetOneDriveUsageAccountDetailWithPeriod) SetFileCount(value *int64)() {
    m.fileCount = value
}
// Sets the isDeleted property value. 
// Parameters:
//  - value : Value to set for the isDeleted property.
func (m *GetOneDriveUsageAccountDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// Sets the lastActivityDate property value. 
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetOneDriveUsageAccountDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the ownerDisplayName property value. 
// Parameters:
//  - value : Value to set for the ownerDisplayName property.
func (m *GetOneDriveUsageAccountDetailWithPeriod) SetOwnerDisplayName(value *string)() {
    m.ownerDisplayName = value
}
// Sets the ownerPrincipalName property value. 
// Parameters:
//  - value : Value to set for the ownerPrincipalName property.
func (m *GetOneDriveUsageAccountDetailWithPeriod) SetOwnerPrincipalName(value *string)() {
    m.ownerPrincipalName = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetOneDriveUsageAccountDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetOneDriveUsageAccountDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the siteUrl property value. 
// Parameters:
//  - value : Value to set for the siteUrl property.
func (m *GetOneDriveUsageAccountDetailWithPeriod) SetSiteUrl(value *string)() {
    m.siteUrl = value
}
// Sets the storageAllocatedInBytes property value. 
// Parameters:
//  - value : Value to set for the storageAllocatedInBytes property.
func (m *GetOneDriveUsageAccountDetailWithPeriod) SetStorageAllocatedInBytes(value *int64)() {
    m.storageAllocatedInBytes = value
}
// Sets the storageUsedInBytes property value. 
// Parameters:
//  - value : Value to set for the storageUsedInBytes property.
func (m *GetOneDriveUsageAccountDetailWithPeriod) SetStorageUsedInBytes(value *int64)() {
    m.storageUsedInBytes = value
}
