package getsharepointsiteusagedetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetSharePointSiteUsageDetailWithPeriod 
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
// NewGetSharePointSiteUsageDetailWithPeriod instantiates a new getSharePointSiteUsageDetailWithPeriod and sets the default values.
func NewGetSharePointSiteUsageDetailWithPeriod()(*GetSharePointSiteUsageDetailWithPeriod) {
    m := &GetSharePointSiteUsageDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetActiveFileCount gets the activeFileCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetActiveFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.activeFileCount
    }
}
// GetAnonymousLinkCount gets the anonymousLinkCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetAnonymousLinkCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.anonymousLinkCount
    }
}
// GetCompanyLinkCount gets the companyLinkCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetCompanyLinkCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.companyLinkCount
    }
}
// GetExternalSharing gets the externalSharing property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetExternalSharing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.externalSharing
    }
}
// GetFileCount gets the fileCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileCount
    }
}
// GetGeolocation gets the geolocation property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetGeolocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.geolocation
    }
}
// GetIsDeleted gets the isDeleted property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// GetLastActivityDate gets the lastActivityDate property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// GetOwnerDisplayName gets the ownerDisplayName property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetOwnerDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerDisplayName
    }
}
// GetOwnerPrincipalName gets the ownerPrincipalName property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetOwnerPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerPrincipalName
    }
}
// GetPageViewCount gets the pageViewCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetPageViewCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.pageViewCount
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetRootWebTemplate gets the rootWebTemplate property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetRootWebTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rootWebTemplate
    }
}
// GetSecureLinkForGuestCount gets the secureLinkForGuestCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSecureLinkForGuestCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.secureLinkForGuestCount
    }
}
// GetSecureLinkForMemberCount gets the secureLinkForMemberCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSecureLinkForMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.secureLinkForMemberCount
    }
}
// GetSiteId gets the siteId property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSiteId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteId
    }
}
// GetSiteSensitivityLabelId gets the siteSensitivityLabelId property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSiteSensitivityLabelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteSensitivityLabelId
    }
}
// GetSiteUrl gets the siteUrl property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteUrl
    }
}
// GetStorageAllocatedInBytes gets the storageAllocatedInBytes property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetStorageAllocatedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageAllocatedInBytes
    }
}
// GetStorageUsedInBytes gets the storageUsedInBytes property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageUsedInBytes
    }
}
// GetUnmanagedDevicePolicy gets the unmanagedDevicePolicy property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetUnmanagedDevicePolicy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.unmanagedDevicePolicy
    }
}
// GetVisitedPageCount gets the visitedPageCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) GetVisitedPageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.visitedPageCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetSharePointSiteUsageDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["anonymousLinkCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnonymousLinkCount(val)
        }
        return nil
    }
    res["companyLinkCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyLinkCount(val)
        }
        return nil
    }
    res["externalSharing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalSharing(val)
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
    res["geolocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeolocation(val)
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
    res["pageViewCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPageViewCount(val)
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
    res["rootWebTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootWebTemplate(val)
        }
        return nil
    }
    res["secureLinkForGuestCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureLinkForGuestCount(val)
        }
        return nil
    }
    res["secureLinkForMemberCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureLinkForMemberCount(val)
        }
        return nil
    }
    res["siteId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteId(val)
        }
        return nil
    }
    res["siteSensitivityLabelId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteSensitivityLabelId(val)
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
    res["unmanagedDevicePolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnmanagedDevicePolicy(val)
        }
        return nil
    }
    res["visitedPageCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisitedPageCount(val)
        }
        return nil
    }
    return res
}
func (m *GetSharePointSiteUsageDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetActiveFileCount sets the activeFileCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetActiveFileCount(value *int64)() {
    m.activeFileCount = value
}
// SetAnonymousLinkCount sets the anonymousLinkCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetAnonymousLinkCount(value *int64)() {
    m.anonymousLinkCount = value
}
// SetCompanyLinkCount sets the companyLinkCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetCompanyLinkCount(value *int64)() {
    m.companyLinkCount = value
}
// SetExternalSharing sets the externalSharing property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetExternalSharing(value *bool)() {
    m.externalSharing = value
}
// SetFileCount sets the fileCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetFileCount(value *int64)() {
    m.fileCount = value
}
// SetGeolocation sets the geolocation property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetGeolocation(value *string)() {
    m.geolocation = value
}
// SetIsDeleted sets the isDeleted property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// SetLastActivityDate sets the lastActivityDate property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// SetOwnerDisplayName sets the ownerDisplayName property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetOwnerDisplayName(value *string)() {
    m.ownerDisplayName = value
}
// SetOwnerPrincipalName sets the ownerPrincipalName property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetOwnerPrincipalName(value *string)() {
    m.ownerPrincipalName = value
}
// SetPageViewCount sets the pageViewCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetPageViewCount(value *int64)() {
    m.pageViewCount = value
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// SetRootWebTemplate sets the rootWebTemplate property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetRootWebTemplate(value *string)() {
    m.rootWebTemplate = value
}
// SetSecureLinkForGuestCount sets the secureLinkForGuestCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSecureLinkForGuestCount(value *int64)() {
    m.secureLinkForGuestCount = value
}
// SetSecureLinkForMemberCount sets the secureLinkForMemberCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSecureLinkForMemberCount(value *int64)() {
    m.secureLinkForMemberCount = value
}
// SetSiteId sets the siteId property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSiteId(value *string)() {
    m.siteId = value
}
// SetSiteSensitivityLabelId sets the siteSensitivityLabelId property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSiteSensitivityLabelId(value *string)() {
    m.siteSensitivityLabelId = value
}
// SetSiteUrl sets the siteUrl property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSiteUrl(value *string)() {
    m.siteUrl = value
}
// SetStorageAllocatedInBytes sets the storageAllocatedInBytes property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetStorageAllocatedInBytes(value *int64)() {
    m.storageAllocatedInBytes = value
}
// SetStorageUsedInBytes sets the storageUsedInBytes property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetStorageUsedInBytes(value *int64)() {
    m.storageUsedInBytes = value
}
// SetUnmanagedDevicePolicy sets the unmanagedDevicePolicy property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetUnmanagedDevicePolicy(value *string)() {
    m.unmanagedDevicePolicy = value
}
// SetVisitedPageCount sets the visitedPageCount property value. 
func (m *GetSharePointSiteUsageDetailWithPeriod) SetVisitedPageCount(value *int64)() {
    m.visitedPageCount = value
}
