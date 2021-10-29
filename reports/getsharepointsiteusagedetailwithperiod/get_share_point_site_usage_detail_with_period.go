package getsharepointsiteusagedetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetSharePointSiteUsageDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    activeFileCount *int64;
    // 
    anonymousLinkCount *int64;
    // 
    companyLinkCount *int64;
    // 
    externalSharing *bool;
    // 
    fileCount *int64;
    // 
    geolocation *string;
    // 
    isDeleted *bool;
    // 
    lastActivityDate *string;
    // 
    ownerDisplayName *string;
    // 
    ownerPrincipalName *string;
    // 
    pageViewCount *int64;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    rootWebTemplate *string;
    // 
    secureLinkForGuestCount *int64;
    // 
    secureLinkForMemberCount *int64;
    // 
    siteId *string;
    // 
    siteSensitivityLabelId *string;
    // 
    siteUrl *string;
    // 
    storageAllocatedInBytes *int64;
    // 
    storageUsedInBytes *int64;
    // 
    unmanagedDevicePolicy *string;
    // 
    visitedPageCount *int64;
}
// Instantiates a new getSharePointSiteUsageDetailWithPeriod and sets the default values.
func NewGetSharePointSiteUsageDetailWithPeriod()(*GetSharePointSiteUsageDetailWithPeriod) {
    m := &GetSharePointSiteUsageDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the activeFileCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetActiveFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.activeFileCount
    }
}
// Gets the anonymousLinkCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetAnonymousLinkCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.anonymousLinkCount
    }
}
// Gets the companyLinkCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetCompanyLinkCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.companyLinkCount
    }
}
// Gets the externalSharing property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetExternalSharing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.externalSharing
    }
}
// Gets the fileCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileCount
    }
}
// Gets the geolocation property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetGeolocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.geolocation
    }
}
// Gets the isDeleted property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// Gets the lastActivityDate property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the ownerDisplayName property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetOwnerDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerDisplayName
    }
}
// Gets the ownerPrincipalName property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetOwnerPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerPrincipalName
    }
}
// Gets the pageViewCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetPageViewCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.pageViewCount
    }
}
// Gets the reportPeriod property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the rootWebTemplate property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetRootWebTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rootWebTemplate
    }
}
// Gets the secureLinkForGuestCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSecureLinkForGuestCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.secureLinkForGuestCount
    }
}
// Gets the secureLinkForMemberCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSecureLinkForMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.secureLinkForMemberCount
    }
}
// Gets the siteId property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSiteId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteId
    }
}
// Gets the siteSensitivityLabelId property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSiteSensitivityLabelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteSensitivityLabelId
    }
}
// Gets the siteUrl property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteUrl
    }
}
// Gets the storageAllocatedInBytes property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetStorageAllocatedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageAllocatedInBytes
    }
}
// Gets the storageUsedInBytes property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageUsedInBytes
    }
}
// Gets the unmanagedDevicePolicy property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetUnmanagedDevicePolicy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.unmanagedDevicePolicy
    }
}
// Gets the visitedPageCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetVisitedPageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.visitedPageCount
    }
}
// The deserialization information for the current model
func (m *GetSharePointSiteUsageDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetActiveFileCount(val)
        return nil
    }
    res["anonymousLinkCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetAnonymousLinkCount(val)
        return nil
    }
    res["companyLinkCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCompanyLinkCount(val)
        return nil
    }
    res["externalSharing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetExternalSharing(val)
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
    res["geolocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGeolocation(val)
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
    res["pageViewCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPageViewCount(val)
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
    res["rootWebTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRootWebTemplate(val)
        return nil
    }
    res["secureLinkForGuestCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSecureLinkForGuestCount(val)
        return nil
    }
    res["secureLinkForMemberCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSecureLinkForMemberCount(val)
        return nil
    }
    res["siteId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSiteId(val)
        return nil
    }
    res["siteSensitivityLabelId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSiteSensitivityLabelId(val)
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
    res["unmanagedDevicePolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUnmanagedDevicePolicy(val)
        return nil
    }
    res["visitedPageCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetVisitedPageCount(val)
        return nil
    }
    return res
}
func (m *GetSharePointSiteUsageDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetSharePointSiteUsageDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt64Value("anonymousLinkCount", m.GetAnonymousLinkCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("companyLinkCount", m.GetCompanyLinkCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("externalSharing", m.GetExternalSharing())
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
        err = writer.WriteStringValue("geolocation", m.GetGeolocation())
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
        err = writer.WriteInt64Value("pageViewCount", m.GetPageViewCount())
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
        err = writer.WriteStringValue("rootWebTemplate", m.GetRootWebTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("secureLinkForGuestCount", m.GetSecureLinkForGuestCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("secureLinkForMemberCount", m.GetSecureLinkForMemberCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("siteId", m.GetSiteId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("siteSensitivityLabelId", m.GetSiteSensitivityLabelId())
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
    {
        err = writer.WriteStringValue("unmanagedDevicePolicy", m.GetUnmanagedDevicePolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("visitedPageCount", m.GetVisitedPageCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the activeFileCount property value. 
// Parameters:
//  - value : Value to set for the activeFileCount property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetActiveFileCount(value *int64)() {
    m.activeFileCount = value
}
// Sets the anonymousLinkCount property value. 
// Parameters:
//  - value : Value to set for the anonymousLinkCount property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetAnonymousLinkCount(value *int64)() {
    m.anonymousLinkCount = value
}
// Sets the companyLinkCount property value. 
// Parameters:
//  - value : Value to set for the companyLinkCount property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetCompanyLinkCount(value *int64)() {
    m.companyLinkCount = value
}
// Sets the externalSharing property value. 
// Parameters:
//  - value : Value to set for the externalSharing property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetExternalSharing(value *bool)() {
    m.externalSharing = value
}
// Sets the fileCount property value. 
// Parameters:
//  - value : Value to set for the fileCount property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetFileCount(value *int64)() {
    m.fileCount = value
}
// Sets the geolocation property value. 
// Parameters:
//  - value : Value to set for the geolocation property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetGeolocation(value *string)() {
    m.geolocation = value
}
// Sets the isDeleted property value. 
// Parameters:
//  - value : Value to set for the isDeleted property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// Sets the lastActivityDate property value. 
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the ownerDisplayName property value. 
// Parameters:
//  - value : Value to set for the ownerDisplayName property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetOwnerDisplayName(value *string)() {
    m.ownerDisplayName = value
}
// Sets the ownerPrincipalName property value. 
// Parameters:
//  - value : Value to set for the ownerPrincipalName property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetOwnerPrincipalName(value *string)() {
    m.ownerPrincipalName = value
}
// Sets the pageViewCount property value. 
// Parameters:
//  - value : Value to set for the pageViewCount property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetPageViewCount(value *int64)() {
    m.pageViewCount = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the rootWebTemplate property value. 
// Parameters:
//  - value : Value to set for the rootWebTemplate property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetRootWebTemplate(value *string)() {
    m.rootWebTemplate = value
}
// Sets the secureLinkForGuestCount property value. 
// Parameters:
//  - value : Value to set for the secureLinkForGuestCount property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSecureLinkForGuestCount(value *int64)() {
    m.secureLinkForGuestCount = value
}
// Sets the secureLinkForMemberCount property value. 
// Parameters:
//  - value : Value to set for the secureLinkForMemberCount property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSecureLinkForMemberCount(value *int64)() {
    m.secureLinkForMemberCount = value
}
// Sets the siteId property value. 
// Parameters:
//  - value : Value to set for the siteId property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSiteId(value *string)() {
    m.siteId = value
}
// Sets the siteSensitivityLabelId property value. 
// Parameters:
//  - value : Value to set for the siteSensitivityLabelId property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSiteSensitivityLabelId(value *string)() {
    m.siteSensitivityLabelId = value
}
// Sets the siteUrl property value. 
// Parameters:
//  - value : Value to set for the siteUrl property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSiteUrl(value *string)() {
    m.siteUrl = value
}
// Sets the storageAllocatedInBytes property value. 
// Parameters:
//  - value : Value to set for the storageAllocatedInBytes property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetStorageAllocatedInBytes(value *int64)() {
    m.storageAllocatedInBytes = value
}
// Sets the storageUsedInBytes property value. 
// Parameters:
//  - value : Value to set for the storageUsedInBytes property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetStorageUsedInBytes(value *int64)() {
    m.storageUsedInBytes = value
}
// Sets the unmanagedDevicePolicy property value. 
// Parameters:
//  - value : Value to set for the unmanagedDevicePolicy property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetUnmanagedDevicePolicy(value *string)() {
    m.unmanagedDevicePolicy = value
}
// Sets the visitedPageCount property value. 
// Parameters:
//  - value : Value to set for the visitedPageCount property.
func (m *GetSharePointSiteUsageDetailWithPeriod) SetVisitedPageCount(value *int64)() {
    m.visitedPageCount = value
}
