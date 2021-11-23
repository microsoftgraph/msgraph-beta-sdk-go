package getonedriveusageaccountdetailwithdate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// getOneDriveUsageAccountDetailWithDate 
type GetOneDriveUsageAccountDetailWithDate struct {
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
// NewGetOneDriveUsageAccountDetailWithDate instantiates a new getOneDriveUsageAccountDetailWithDate and sets the default values.
func NewGetOneDriveUsageAccountDetailWithDate()(*GetOneDriveUsageAccountDetailWithDate) {
    m := &GetOneDriveUsageAccountDetailWithDate{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetActiveFileCount gets the activeFileCount property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) GetActiveFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.activeFileCount
    }
}
// GetFileCount gets the fileCount property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) GetFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileCount
    }
}
// GetIsDeleted gets the isDeleted property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// GetLastActivityDate gets the lastActivityDate property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// GetOwnerDisplayName gets the ownerDisplayName property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) GetOwnerDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerDisplayName
    }
}
// GetOwnerPrincipalName gets the ownerPrincipalName property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) GetOwnerPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerPrincipalName
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetSiteUrl gets the siteUrl property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) GetSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteUrl
    }
}
// GetStorageAllocatedInBytes gets the storageAllocatedInBytes property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) GetStorageAllocatedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageAllocatedInBytes
    }
}
// GetStorageUsedInBytes gets the storageUsedInBytes property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) GetStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageUsedInBytes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetOneDriveUsageAccountDetailWithDate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveFileCount(val)
        }
        return nil
    }
    res["fileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileCount(val)
        }
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["lastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivityDate(val)
        }
        return nil
    }
    res["ownerDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerDisplayName(val)
        }
        return nil
    }
    res["ownerPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerPrincipalName(val)
        }
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportPeriod(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["siteUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteUrl(val)
        }
        return nil
    }
    res["storageAllocatedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageAllocatedInBytes(val)
        }
        return nil
    }
    res["storageUsedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageUsedInBytes(val)
        }
        return nil
    }
    return res
}
func (m *GetOneDriveUsageAccountDetailWithDate) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetOneDriveUsageAccountDetailWithDate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetActiveFileCount sets the activeFileCount property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) SetActiveFileCount(value *int64)() {
    m.activeFileCount = value
}
// SetFileCount sets the fileCount property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) SetFileCount(value *int64)() {
    m.fileCount = value
}
// SetIsDeleted sets the isDeleted property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// SetLastActivityDate sets the lastActivityDate property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// SetOwnerDisplayName sets the ownerDisplayName property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) SetOwnerDisplayName(value *string)() {
    m.ownerDisplayName = value
}
// SetOwnerPrincipalName sets the ownerPrincipalName property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) SetOwnerPrincipalName(value *string)() {
    m.ownerPrincipalName = value
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// SetSiteUrl sets the siteUrl property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) SetSiteUrl(value *string)() {
    m.siteUrl = value
}
// SetStorageAllocatedInBytes sets the storageAllocatedInBytes property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) SetStorageAllocatedInBytes(value *int64)() {
    m.storageAllocatedInBytes = value
}
// SetStorageUsedInBytes sets the storageUsedInBytes property value. 
func (m *GetOneDriveUsageAccountDetailWithDate) SetStorageUsedInBytes(value *int64)() {
    m.storageUsedInBytes = value
}
