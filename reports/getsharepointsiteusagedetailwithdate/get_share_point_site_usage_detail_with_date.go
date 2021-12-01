package getsharepointsiteusagedetailwithdate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetSharePointSiteUsageDetailWithDate 
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
// NewGetSharePointSiteUsageDetailWithDate instantiates a new getSharePointSiteUsageDetailWithDate and sets the default values.
func NewGetSharePointSiteUsageDetailWithDate()(*GetSharePointSiteUsageDetailWithDate) {
    m := &GetSharePointSiteUsageDetailWithDate{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetActiveFileCount gets the activeFileCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetActiveFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.activeFileCount
    }
}
// GetAnonymousLinkCount gets the anonymousLinkCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetAnonymousLinkCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.anonymousLinkCount
    }
}
// GetCompanyLinkCount gets the companyLinkCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetCompanyLinkCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.companyLinkCount
    }
}
// GetExternalSharing gets the externalSharing property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetExternalSharing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.externalSharing
    }
}
// GetFileCount gets the fileCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileCount
    }
}
// GetGeolocation gets the geolocation property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetGeolocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.geolocation
    }
}
// GetIsDeleted gets the isDeleted property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// GetLastActivityDate gets the lastActivityDate property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// GetOwnerDisplayName gets the ownerDisplayName property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetOwnerDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerDisplayName
    }
}
// GetOwnerPrincipalName gets the ownerPrincipalName property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetOwnerPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerPrincipalName
    }
}
// GetPageViewCount gets the pageViewCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetPageViewCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.pageViewCount
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetRootWebTemplate gets the rootWebTemplate property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetRootWebTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rootWebTemplate
    }
}
// GetSecureLinkForGuestCount gets the secureLinkForGuestCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetSecureLinkForGuestCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.secureLinkForGuestCount
    }
}
// GetSecureLinkForMemberCount gets the secureLinkForMemberCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetSecureLinkForMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.secureLinkForMemberCount
    }
}
// GetSiteId gets the siteId property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetSiteId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteId
    }
}
// GetSiteSensitivityLabelId gets the siteSensitivityLabelId property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetSiteSensitivityLabelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteSensitivityLabelId
    }
}
// GetSiteUrl gets the siteUrl property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteUrl
    }
}
// GetStorageAllocatedInBytes gets the storageAllocatedInBytes property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetStorageAllocatedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageAllocatedInBytes
    }
}
// GetStorageUsedInBytes gets the storageUsedInBytes property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageUsedInBytes
    }
}
// GetUnmanagedDevicePolicy gets the unmanagedDevicePolicy property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetUnmanagedDevicePolicy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.unmanagedDevicePolicy
    }
}
// GetVisitedPageCount gets the visitedPageCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) GetVisitedPageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.visitedPageCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetSharePointSiteUsageDetailWithDate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *GetSharePointSiteUsageDetailWithDate) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetActiveFileCount sets the activeFileCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetActiveFileCount(value *int64)() {
    if m != nil {
        m.activeFileCount = value
    }
}
// SetAnonymousLinkCount sets the anonymousLinkCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetAnonymousLinkCount(value *int64)() {
    if m != nil {
        m.anonymousLinkCount = value
    }
}
// SetCompanyLinkCount sets the companyLinkCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetCompanyLinkCount(value *int64)() {
    if m != nil {
        m.companyLinkCount = value
    }
}
// SetExternalSharing sets the externalSharing property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetExternalSharing(value *bool)() {
    if m != nil {
        m.externalSharing = value
    }
}
// SetFileCount sets the fileCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetFileCount(value *int64)() {
    if m != nil {
        m.fileCount = value
    }
}
// SetGeolocation sets the geolocation property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetGeolocation(value *string)() {
    if m != nil {
        m.geolocation = value
    }
}
// SetIsDeleted sets the isDeleted property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetIsDeleted(value *bool)() {
    if m != nil {
        m.isDeleted = value
    }
}
// SetLastActivityDate sets the lastActivityDate property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetLastActivityDate(value *string)() {
    if m != nil {
        m.lastActivityDate = value
    }
}
// SetOwnerDisplayName sets the ownerDisplayName property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetOwnerDisplayName(value *string)() {
    if m != nil {
        m.ownerDisplayName = value
    }
}
// SetOwnerPrincipalName sets the ownerPrincipalName property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetOwnerPrincipalName(value *string)() {
    if m != nil {
        m.ownerPrincipalName = value
    }
}
// SetPageViewCount sets the pageViewCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetPageViewCount(value *int64)() {
    if m != nil {
        m.pageViewCount = value
    }
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetReportRefreshDate(value *string)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetRootWebTemplate sets the rootWebTemplate property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetRootWebTemplate(value *string)() {
    if m != nil {
        m.rootWebTemplate = value
    }
}
// SetSecureLinkForGuestCount sets the secureLinkForGuestCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetSecureLinkForGuestCount(value *int64)() {
    if m != nil {
        m.secureLinkForGuestCount = value
    }
}
// SetSecureLinkForMemberCount sets the secureLinkForMemberCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetSecureLinkForMemberCount(value *int64)() {
    if m != nil {
        m.secureLinkForMemberCount = value
    }
}
// SetSiteId sets the siteId property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetSiteId(value *string)() {
    if m != nil {
        m.siteId = value
    }
}
// SetSiteSensitivityLabelId sets the siteSensitivityLabelId property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetSiteSensitivityLabelId(value *string)() {
    if m != nil {
        m.siteSensitivityLabelId = value
    }
}
// SetSiteUrl sets the siteUrl property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetSiteUrl(value *string)() {
    if m != nil {
        m.siteUrl = value
    }
}
// SetStorageAllocatedInBytes sets the storageAllocatedInBytes property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetStorageAllocatedInBytes(value *int64)() {
    if m != nil {
        m.storageAllocatedInBytes = value
    }
}
// SetStorageUsedInBytes sets the storageUsedInBytes property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetStorageUsedInBytes(value *int64)() {
    if m != nil {
        m.storageUsedInBytes = value
    }
}
// SetUnmanagedDevicePolicy sets the unmanagedDevicePolicy property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetUnmanagedDevicePolicy(value *string)() {
    if m != nil {
        m.unmanagedDevicePolicy = value
    }
}
// SetVisitedPageCount sets the visitedPageCount property value. 
func (m *GetSharePointSiteUsageDetailWithDate) SetVisitedPageCount(value *int64)() {
    if m != nil {
        m.visitedPageCount = value
    }
}
