package getsharepointsiteusagedetailwithdate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetSharePointSiteUsageDetailWithDate struct {
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
// Instantiates a new getSharePointSiteUsageDetailWithDate and sets the default values.
func NewGetSharePointSiteUsageDetailWithDate()(*GetSharePointSiteUsageDetailWithDate) {
    m := &GetSharePointSiteUsageDetailWithDate{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the activeFileCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetActiveFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.activeFileCount
    }
}
// Gets the anonymousLinkCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetAnonymousLinkCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.anonymousLinkCount
    }
}
// Gets the companyLinkCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetCompanyLinkCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.companyLinkCount
    }
}
// Gets the externalSharing property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetExternalSharing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.externalSharing
    }
}
// Gets the fileCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileCount
    }
}
// Gets the geolocation property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetGeolocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.geolocation
    }
}
// Gets the isDeleted property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// Gets the lastActivityDate property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the ownerDisplayName property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetOwnerDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerDisplayName
    }
}
// Gets the ownerPrincipalName property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetOwnerPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerPrincipalName
    }
}
// Gets the pageViewCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetPageViewCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.pageViewCount
    }
}
// Gets the reportPeriod property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the rootWebTemplate property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetRootWebTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rootWebTemplate
    }
}
// Gets the secureLinkForGuestCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetSecureLinkForGuestCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.secureLinkForGuestCount
    }
}
// Gets the secureLinkForMemberCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetSecureLinkForMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.secureLinkForMemberCount
    }
}
// Gets the siteId property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetSiteId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteId
    }
}
// Gets the siteSensitivityLabelId property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetSiteSensitivityLabelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteSensitivityLabelId
    }
}
// Gets the siteUrl property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteUrl
    }
}
// Gets the storageAllocatedInBytes property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetStorageAllocatedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageAllocatedInBytes
    }
}
// Gets the storageUsedInBytes property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageUsedInBytes
    }
}
// Gets the unmanagedDevicePolicy property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetUnmanagedDevicePolicy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.unmanagedDevicePolicy
    }
}
// Gets the visitedPageCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetVisitedPageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.visitedPageCount
    }
}
// The deserialization information for the current model
func (m *GetSharePointSiteUsageDetailWithDate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *GetSharePointSiteUsageDetailWithDate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetSharePointSiteUsageDetailWithDate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *GetSharePointSiteUsageDetailWithDate) SetActiveFileCount(value *int64)() {
    m.activeFileCount = value
}
// Sets the anonymousLinkCount property value. 
// Parameters:
//  - value : Value to set for the anonymousLinkCount property.
func (m *GetSharePointSiteUsageDetailWithDate) SetAnonymousLinkCount(value *int64)() {
    m.anonymousLinkCount = value
}
// Sets the companyLinkCount property value. 
// Parameters:
//  - value : Value to set for the companyLinkCount property.
func (m *GetSharePointSiteUsageDetailWithDate) SetCompanyLinkCount(value *int64)() {
    m.companyLinkCount = value
}
// Sets the externalSharing property value. 
// Parameters:
//  - value : Value to set for the externalSharing property.
func (m *GetSharePointSiteUsageDetailWithDate) SetExternalSharing(value *bool)() {
    m.externalSharing = value
}
// Sets the fileCount property value. 
// Parameters:
//  - value : Value to set for the fileCount property.
func (m *GetSharePointSiteUsageDetailWithDate) SetFileCount(value *int64)() {
    m.fileCount = value
}
// Sets the geolocation property value. 
// Parameters:
//  - value : Value to set for the geolocation property.
func (m *GetSharePointSiteUsageDetailWithDate) SetGeolocation(value *string)() {
    m.geolocation = value
}
// Sets the isDeleted property value. 
// Parameters:
//  - value : Value to set for the isDeleted property.
func (m *GetSharePointSiteUsageDetailWithDate) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// Sets the lastActivityDate property value. 
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetSharePointSiteUsageDetailWithDate) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the ownerDisplayName property value. 
// Parameters:
//  - value : Value to set for the ownerDisplayName property.
func (m *GetSharePointSiteUsageDetailWithDate) SetOwnerDisplayName(value *string)() {
    m.ownerDisplayName = value
}
// Sets the ownerPrincipalName property value. 
// Parameters:
//  - value : Value to set for the ownerPrincipalName property.
func (m *GetSharePointSiteUsageDetailWithDate) SetOwnerPrincipalName(value *string)() {
    m.ownerPrincipalName = value
}
// Sets the pageViewCount property value. 
// Parameters:
//  - value : Value to set for the pageViewCount property.
func (m *GetSharePointSiteUsageDetailWithDate) SetPageViewCount(value *int64)() {
    m.pageViewCount = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetSharePointSiteUsageDetailWithDate) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetSharePointSiteUsageDetailWithDate) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the rootWebTemplate property value. 
// Parameters:
//  - value : Value to set for the rootWebTemplate property.
func (m *GetSharePointSiteUsageDetailWithDate) SetRootWebTemplate(value *string)() {
    m.rootWebTemplate = value
}
// Sets the secureLinkForGuestCount property value. 
// Parameters:
//  - value : Value to set for the secureLinkForGuestCount property.
func (m *GetSharePointSiteUsageDetailWithDate) SetSecureLinkForGuestCount(value *int64)() {
    m.secureLinkForGuestCount = value
}
// Sets the secureLinkForMemberCount property value. 
// Parameters:
//  - value : Value to set for the secureLinkForMemberCount property.
func (m *GetSharePointSiteUsageDetailWithDate) SetSecureLinkForMemberCount(value *int64)() {
    m.secureLinkForMemberCount = value
}
// Sets the siteId property value. 
// Parameters:
//  - value : Value to set for the siteId property.
func (m *GetSharePointSiteUsageDetailWithDate) SetSiteId(value *string)() {
    m.siteId = value
}
// Sets the siteSensitivityLabelId property value. 
// Parameters:
//  - value : Value to set for the siteSensitivityLabelId property.
func (m *GetSharePointSiteUsageDetailWithDate) SetSiteSensitivityLabelId(value *string)() {
    m.siteSensitivityLabelId = value
}
// Sets the siteUrl property value. 
// Parameters:
//  - value : Value to set for the siteUrl property.
func (m *GetSharePointSiteUsageDetailWithDate) SetSiteUrl(value *string)() {
    m.siteUrl = value
}
// Sets the storageAllocatedInBytes property value. 
// Parameters:
//  - value : Value to set for the storageAllocatedInBytes property.
func (m *GetSharePointSiteUsageDetailWithDate) SetStorageAllocatedInBytes(value *int64)() {
    m.storageAllocatedInBytes = value
}
// Sets the storageUsedInBytes property value. 
// Parameters:
//  - value : Value to set for the storageUsedInBytes property.
func (m *GetSharePointSiteUsageDetailWithDate) SetStorageUsedInBytes(value *int64)() {
    m.storageUsedInBytes = value
}
// Sets the unmanagedDevicePolicy property value. 
// Parameters:
//  - value : Value to set for the unmanagedDevicePolicy property.
func (m *GetSharePointSiteUsageDetailWithDate) SetUnmanagedDevicePolicy(value *string)() {
    m.unmanagedDevicePolicy = value
}
// Sets the visitedPageCount property value. 
// Parameters:
//  - value : Value to set for the visitedPageCount property.
func (m *GetSharePointSiteUsageDetailWithDate) SetVisitedPageCount(value *int64)() {
    m.visitedPageCount = value
}
