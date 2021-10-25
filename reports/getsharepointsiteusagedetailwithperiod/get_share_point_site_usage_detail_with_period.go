package getsharepointsiteusagedetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetSharePointSiteUsageDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    activeFileCount *int64;
    anonymousLinkCount *int64;
    companyLinkCount *int64;
    externalSharing *bool;
    fileCount *int64;
    geolocation *string;
    isDeleted *bool;
    lastActivityDate *string;
    ownerDisplayName *string;
    ownerPrincipalName *string;
    pageViewCount *int64;
    reportPeriod *string;
    reportRefreshDate *string;
    rootWebTemplate *string;
    secureLinkForGuestCount *int64;
    secureLinkForMemberCount *int64;
    siteId *string;
    siteSensitivityLabelId *string;
    siteUrl *string;
    storageAllocatedInBytes *int64;
    storageUsedInBytes *int64;
    unmanagedDevicePolicy *string;
    visitedPageCount *int64;
}
func NewGetSharePointSiteUsageDetailWithPeriod()(*GetSharePointSiteUsageDetailWithPeriod) {
    m := &GetSharePointSiteUsageDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetActiveFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.activeFileCount
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetAnonymousLinkCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.anonymousLinkCount
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetCompanyLinkCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.companyLinkCount
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetExternalSharing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.externalSharing
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileCount
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetGeolocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.geolocation
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetOwnerDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerDisplayName
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetOwnerPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerPrincipalName
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetPageViewCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.pageViewCount
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetRootWebTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rootWebTemplate
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSecureLinkForGuestCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.secureLinkForGuestCount
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSecureLinkForMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.secureLinkForMemberCount
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSiteId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteId
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSiteSensitivityLabelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteSensitivityLabelId
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteUrl
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetStorageAllocatedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageAllocatedInBytes
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageUsedInBytes
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetUnmanagedDevicePolicy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.unmanagedDevicePolicy
    }
}
func (m *GetSharePointSiteUsageDetailWithPeriod) GetVisitedPageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.visitedPageCount
    }
}
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
func (m *GetSharePointSiteUsageDetailWithPeriod) SetActiveFileCount(value *int64)() {
    m.activeFileCount = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetAnonymousLinkCount(value *int64)() {
    m.anonymousLinkCount = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetCompanyLinkCount(value *int64)() {
    m.companyLinkCount = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetExternalSharing(value *bool)() {
    m.externalSharing = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetFileCount(value *int64)() {
    m.fileCount = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetGeolocation(value *string)() {
    m.geolocation = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetOwnerDisplayName(value *string)() {
    m.ownerDisplayName = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetOwnerPrincipalName(value *string)() {
    m.ownerPrincipalName = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetPageViewCount(value *int64)() {
    m.pageViewCount = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetRootWebTemplate(value *string)() {
    m.rootWebTemplate = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSecureLinkForGuestCount(value *int64)() {
    m.secureLinkForGuestCount = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSecureLinkForMemberCount(value *int64)() {
    m.secureLinkForMemberCount = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSiteId(value *string)() {
    m.siteId = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSiteSensitivityLabelId(value *string)() {
    m.siteSensitivityLabelId = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetSiteUrl(value *string)() {
    m.siteUrl = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetStorageAllocatedInBytes(value *int64)() {
    m.storageAllocatedInBytes = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetStorageUsedInBytes(value *int64)() {
    m.storageUsedInBytes = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetUnmanagedDevicePolicy(value *string)() {
    m.unmanagedDevicePolicy = value
}
func (m *GetSharePointSiteUsageDetailWithPeriod) SetVisitedPageCount(value *int64)() {
    m.visitedPageCount = value
}
