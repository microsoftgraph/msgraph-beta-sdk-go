package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.deviceManagement entity.
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ReportsRequestBuilderGetQueryParameters reports singleton
type ReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsRequestBuilderGetQueryParameters
}
// ReportsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CachedReportConfigurations provides operations to manage the cachedReportConfigurations property of the microsoft.graph.deviceManagementReports entity.
func (m *ReportsRequestBuilder) CachedReportConfigurations()(*ReportsCachedReportConfigurationsRequestBuilder) {
    return NewReportsCachedReportConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CachedReportConfigurationsById provides operations to manage the cachedReportConfigurations property of the microsoft.graph.deviceManagementReports entity.
func (m *ReportsRequestBuilder) CachedReportConfigurationsById(id string)(*ReportsCachedReportConfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementCachedReportConfiguration%2Did"] = id
    }
    return NewReportsCachedReportConfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property reports for deviceManagement
func (m *ReportsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ExportJobs provides operations to manage the exportJobs property of the microsoft.graph.deviceManagementReports entity.
func (m *ReportsRequestBuilder) ExportJobs()(*ReportsExportJobsRequestBuilder) {
    return NewReportsExportJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportJobsById provides operations to manage the exportJobs property of the microsoft.graph.deviceManagementReports entity.
func (m *ReportsRequestBuilder) ExportJobsById(id string)(*ReportsExportJobsDeviceManagementExportJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementExportJob%2Did"] = id
    }
    return NewReportsExportJobsDeviceManagementExportJobItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get reports singleton
func (m *ReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable), nil
}
// GetActiveMalwareReport provides operations to call the getActiveMalwareReport method.
func (m *ReportsRequestBuilder) GetActiveMalwareReport()(*ReportsGetActiveMalwareReportRequestBuilder) {
    return NewReportsGetActiveMalwareReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActiveMalwareSummaryReport provides operations to call the getActiveMalwareSummaryReport method.
func (m *ReportsRequestBuilder) GetActiveMalwareSummaryReport()(*ReportsGetActiveMalwareSummaryReportRequestBuilder) {
    return NewReportsGetActiveMalwareSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAllCertificatesReport provides operations to call the getAllCertificatesReport method.
func (m *ReportsRequestBuilder) GetAllCertificatesReport()(*ReportsGetAllCertificatesReportRequestBuilder) {
    return NewReportsGetAllCertificatesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAppsInstallSummaryReport provides operations to call the getAppsInstallSummaryReport method.
func (m *ReportsRequestBuilder) GetAppsInstallSummaryReport()(*ReportsGetAppsInstallSummaryReportRequestBuilder) {
    return NewReportsGetAppsInstallSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAppStatusOverviewReport provides operations to call the getAppStatusOverviewReport method.
func (m *ReportsRequestBuilder) GetAppStatusOverviewReport()(*ReportsGetAppStatusOverviewReportRequestBuilder) {
    return NewReportsGetAppStatusOverviewReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCachedReport provides operations to call the getCachedReport method.
func (m *ReportsRequestBuilder) GetCachedReport()(*ReportsGetCachedReportRequestBuilder) {
    return NewReportsGetCachedReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCertificatesReport provides operations to call the getCertificatesReport method.
func (m *ReportsRequestBuilder) GetCertificatesReport()(*ReportsGetCertificatesReportRequestBuilder) {
    return NewReportsGetCertificatesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePoliciesReportForDevice provides operations to call the getCompliancePoliciesReportForDevice method.
func (m *ReportsRequestBuilder) GetCompliancePoliciesReportForDevice()(*ReportsGetCompliancePoliciesReportForDeviceRequestBuilder) {
    return NewReportsGetCompliancePoliciesReportForDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyDevicesReport provides operations to call the getCompliancePolicyDevicesReport method.
func (m *ReportsRequestBuilder) GetCompliancePolicyDevicesReport()(*ReportsGetCompliancePolicyDevicesReportRequestBuilder) {
    return NewReportsGetCompliancePolicyDevicesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyDeviceSummaryReport provides operations to call the getCompliancePolicyDeviceSummaryReport method.
func (m *ReportsRequestBuilder) GetCompliancePolicyDeviceSummaryReport()(*ReportsGetCompliancePolicyDeviceSummaryReportRequestBuilder) {
    return NewReportsGetCompliancePolicyDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyNonComplianceReport provides operations to call the getCompliancePolicyNonComplianceReport method.
func (m *ReportsRequestBuilder) GetCompliancePolicyNonComplianceReport()(*ReportsGetCompliancePolicyNonComplianceReportRequestBuilder) {
    return NewReportsGetCompliancePolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyNonComplianceSummaryReport provides operations to call the getCompliancePolicyNonComplianceSummaryReport method.
func (m *ReportsRequestBuilder) GetCompliancePolicyNonComplianceSummaryReport()(*ReportsGetCompliancePolicyNonComplianceSummaryReportRequestBuilder) {
    return NewReportsGetCompliancePolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComplianceSettingDetailsReport provides operations to call the getComplianceSettingDetailsReport method.
func (m *ReportsRequestBuilder) GetComplianceSettingDetailsReport()(*ReportsGetComplianceSettingDetailsReportRequestBuilder) {
    return NewReportsGetComplianceSettingDetailsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComplianceSettingNonComplianceReport provides operations to call the getComplianceSettingNonComplianceReport method.
func (m *ReportsRequestBuilder) GetComplianceSettingNonComplianceReport()(*ReportsGetComplianceSettingNonComplianceReportRequestBuilder) {
    return NewReportsGetComplianceSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComplianceSettingsReport provides operations to call the getComplianceSettingsReport method.
func (m *ReportsRequestBuilder) GetComplianceSettingsReport()(*ReportsGetComplianceSettingsReportRequestBuilder) {
    return NewReportsGetComplianceSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigManagerDevicePolicyStatusReport provides operations to call the getConfigManagerDevicePolicyStatusReport method.
func (m *ReportsRequestBuilder) GetConfigManagerDevicePolicyStatusReport()(*ReportsGetConfigManagerDevicePolicyStatusReportRequestBuilder) {
    return NewReportsGetConfigManagerDevicePolicyStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPoliciesReportForDevice provides operations to call the getConfigurationPoliciesReportForDevice method.
func (m *ReportsRequestBuilder) GetConfigurationPoliciesReportForDevice()(*ReportsGetConfigurationPoliciesReportForDeviceRequestBuilder) {
    return NewReportsGetConfigurationPoliciesReportForDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyDevicesReport provides operations to call the getConfigurationPolicyDevicesReport method.
func (m *ReportsRequestBuilder) GetConfigurationPolicyDevicesReport()(*ReportsGetConfigurationPolicyDevicesReportRequestBuilder) {
    return NewReportsGetConfigurationPolicyDevicesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyDeviceSummaryReport provides operations to call the getConfigurationPolicyDeviceSummaryReport method.
func (m *ReportsRequestBuilder) GetConfigurationPolicyDeviceSummaryReport()(*ReportsGetConfigurationPolicyDeviceSummaryReportRequestBuilder) {
    return NewReportsGetConfigurationPolicyDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyNonComplianceReport provides operations to call the getConfigurationPolicyNonComplianceReport method.
func (m *ReportsRequestBuilder) GetConfigurationPolicyNonComplianceReport()(*ReportsGetConfigurationPolicyNonComplianceReportRequestBuilder) {
    return NewReportsGetConfigurationPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyNonComplianceSummaryReport provides operations to call the getConfigurationPolicyNonComplianceSummaryReport method.
func (m *ReportsRequestBuilder) GetConfigurationPolicyNonComplianceSummaryReport()(*ReportsGetConfigurationPolicyNonComplianceSummaryReportRequestBuilder) {
    return NewReportsGetConfigurationPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicySettingsDeviceSummaryReport provides operations to call the getConfigurationPolicySettingsDeviceSummaryReport method.
func (m *ReportsRequestBuilder) GetConfigurationPolicySettingsDeviceSummaryReport()(*ReportsGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder) {
    return NewReportsGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationSettingDetailsReport provides operations to call the getConfigurationSettingDetailsReport method.
func (m *ReportsRequestBuilder) GetConfigurationSettingDetailsReport()(*ReportsGetConfigurationSettingDetailsReportRequestBuilder) {
    return NewReportsGetConfigurationSettingDetailsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationSettingNonComplianceReport provides operations to call the getConfigurationSettingNonComplianceReport method.
func (m *ReportsRequestBuilder) GetConfigurationSettingNonComplianceReport()(*ReportsGetConfigurationSettingNonComplianceReportRequestBuilder) {
    return NewReportsGetConfigurationSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationSettingsReport provides operations to call the getConfigurationSettingsReport method.
func (m *ReportsRequestBuilder) GetConfigurationSettingsReport()(*ReportsGetConfigurationSettingsReportRequestBuilder) {
    return NewReportsGetConfigurationSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceConfigurationPolicySettingsSummaryReport provides operations to call the getDeviceConfigurationPolicySettingsSummaryReport method.
func (m *ReportsRequestBuilder) GetDeviceConfigurationPolicySettingsSummaryReport()(*ReportsGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) {
    return NewReportsGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceConfigurationPolicyStatusSummary provides operations to call the getDeviceConfigurationPolicyStatusSummary method.
func (m *ReportsRequestBuilder) GetDeviceConfigurationPolicyStatusSummary()(*ReportsGetDeviceConfigurationPolicyStatusSummaryRequestBuilder) {
    return NewReportsGetDeviceConfigurationPolicyStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceInstallStatusReport provides operations to call the getDeviceInstallStatusReport method.
func (m *ReportsRequestBuilder) GetDeviceInstallStatusReport()(*ReportsGetDeviceInstallStatusReportRequestBuilder) {
    return NewReportsGetDeviceInstallStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceManagementIntentPerSettingContributingProfiles provides operations to call the getDeviceManagementIntentPerSettingContributingProfiles method.
func (m *ReportsRequestBuilder) GetDeviceManagementIntentPerSettingContributingProfiles()(*ReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    return NewReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceManagementIntentSettingsReport provides operations to call the getDeviceManagementIntentSettingsReport method.
func (m *ReportsRequestBuilder) GetDeviceManagementIntentSettingsReport()(*ReportsGetDeviceManagementIntentSettingsReportRequestBuilder) {
    return NewReportsGetDeviceManagementIntentSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceNonComplianceReport provides operations to call the getDeviceNonComplianceReport method.
func (m *ReportsRequestBuilder) GetDeviceNonComplianceReport()(*ReportsGetDeviceNonComplianceReportRequestBuilder) {
    return NewReportsGetDeviceNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEncryptionReportForDevices provides operations to call the getEncryptionReportForDevices method.
func (m *ReportsRequestBuilder) GetEncryptionReportForDevices()(*ReportsGetEncryptionReportForDevicesRequestBuilder) {
    return NewReportsGetEncryptionReportForDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEnrollmentConfigurationPoliciesByDevice provides operations to call the getEnrollmentConfigurationPoliciesByDevice method.
func (m *ReportsRequestBuilder) GetEnrollmentConfigurationPoliciesByDevice()(*ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) {
    return NewReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFailedMobileAppsReport provides operations to call the getFailedMobileAppsReport method.
func (m *ReportsRequestBuilder) GetFailedMobileAppsReport()(*ReportsGetFailedMobileAppsReportRequestBuilder) {
    return NewReportsGetFailedMobileAppsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFailedMobileAppsSummaryReport provides operations to call the getFailedMobileAppsSummaryReport method.
func (m *ReportsRequestBuilder) GetFailedMobileAppsSummaryReport()(*ReportsGetFailedMobileAppsSummaryReportRequestBuilder) {
    return NewReportsGetFailedMobileAppsSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetGroupPolicySettingsDeviceSettingsReport provides operations to call the getGroupPolicySettingsDeviceSettingsReport method.
func (m *ReportsRequestBuilder) GetGroupPolicySettingsDeviceSettingsReport()(*ReportsGetGroupPolicySettingsDeviceSettingsReportRequestBuilder) {
    return NewReportsGetGroupPolicySettingsDeviceSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetHistoricalReport provides operations to call the getHistoricalReport method.
func (m *ReportsRequestBuilder) GetHistoricalReport()(*ReportsGetHistoricalReportRequestBuilder) {
    return NewReportsGetHistoricalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMalwareSummaryReport provides operations to call the getMalwareSummaryReport method.
func (m *ReportsRequestBuilder) GetMalwareSummaryReport()(*ReportsGetMalwareSummaryReportRequestBuilder) {
    return NewReportsGetMalwareSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMobileApplicationManagementAppConfigurationReport provides operations to call the getMobileApplicationManagementAppConfigurationReport method.
func (m *ReportsRequestBuilder) GetMobileApplicationManagementAppConfigurationReport()(*ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilder) {
    return NewReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMobileApplicationManagementAppRegistrationSummaryReport provides operations to call the getMobileApplicationManagementAppRegistrationSummaryReport method.
func (m *ReportsRequestBuilder) GetMobileApplicationManagementAppRegistrationSummaryReport()(*ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder) {
    return NewReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNoncompliantDevicesAndSettingsReport provides operations to call the getNoncompliantDevicesAndSettingsReport method.
func (m *ReportsRequestBuilder) GetNoncompliantDevicesAndSettingsReport()(*ReportsGetNoncompliantDevicesAndSettingsReportRequestBuilder) {
    return NewReportsGetNoncompliantDevicesAndSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPolicyNonComplianceMetadata provides operations to call the getPolicyNonComplianceMetadata method.
func (m *ReportsRequestBuilder) GetPolicyNonComplianceMetadata()(*ReportsGetPolicyNonComplianceMetadataRequestBuilder) {
    return NewReportsGetPolicyNonComplianceMetadataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPolicyNonComplianceReport provides operations to call the getPolicyNonComplianceReport method.
func (m *ReportsRequestBuilder) GetPolicyNonComplianceReport()(*ReportsGetPolicyNonComplianceReportRequestBuilder) {
    return NewReportsGetPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPolicyNonComplianceSummaryReport provides operations to call the getPolicyNonComplianceSummaryReport method.
func (m *ReportsRequestBuilder) GetPolicyNonComplianceSummaryReport()(*ReportsGetPolicyNonComplianceSummaryReportRequestBuilder) {
    return NewReportsGetPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetQuietTimePolicyUsersReport provides operations to call the getQuietTimePolicyUsersReport method.
func (m *ReportsRequestBuilder) GetQuietTimePolicyUsersReport()(*ReportsGetQuietTimePolicyUsersReportRequestBuilder) {
    return NewReportsGetQuietTimePolicyUsersReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetQuietTimePolicyUserSummaryReport provides operations to call the getQuietTimePolicyUserSummaryReport method.
func (m *ReportsRequestBuilder) GetQuietTimePolicyUserSummaryReport()(*ReportsGetQuietTimePolicyUserSummaryReportRequestBuilder) {
    return NewReportsGetQuietTimePolicyUserSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetRelatedAppsStatusReport provides operations to call the getRelatedAppsStatusReport method.
func (m *ReportsRequestBuilder) GetRelatedAppsStatusReport()(*ReportsGetRelatedAppsStatusReportRequestBuilder) {
    return NewReportsGetRelatedAppsStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetRemoteAssistanceSessionsReport provides operations to call the getRemoteAssistanceSessionsReport method.
func (m *ReportsRequestBuilder) GetRemoteAssistanceSessionsReport()(*ReportsGetRemoteAssistanceSessionsReportRequestBuilder) {
    return NewReportsGetRemoteAssistanceSessionsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetReportFilters provides operations to call the getReportFilters method.
func (m *ReportsRequestBuilder) GetReportFilters()(*ReportsGetReportFiltersRequestBuilder) {
    return NewReportsGetReportFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetSettingNonComplianceReport provides operations to call the getSettingNonComplianceReport method.
func (m *ReportsRequestBuilder) GetSettingNonComplianceReport()(*ReportsGetSettingNonComplianceReportRequestBuilder) {
    return NewReportsGetSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUnhealthyDefenderAgentsReport provides operations to call the getUnhealthyDefenderAgentsReport method.
func (m *ReportsRequestBuilder) GetUnhealthyDefenderAgentsReport()(*ReportsGetUnhealthyDefenderAgentsReportRequestBuilder) {
    return NewReportsGetUnhealthyDefenderAgentsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUnhealthyFirewallReport provides operations to call the getUnhealthyFirewallReport method.
func (m *ReportsRequestBuilder) GetUnhealthyFirewallReport()(*ReportsGetUnhealthyFirewallReportRequestBuilder) {
    return NewReportsGetUnhealthyFirewallReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUnhealthyFirewallSummaryReport provides operations to call the getUnhealthyFirewallSummaryReport method.
func (m *ReportsRequestBuilder) GetUnhealthyFirewallSummaryReport()(*ReportsGetUnhealthyFirewallSummaryReportRequestBuilder) {
    return NewReportsGetUnhealthyFirewallSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserInstallStatusReport provides operations to call the getUserInstallStatusReport method.
func (m *ReportsRequestBuilder) GetUserInstallStatusReport()(*ReportsGetUserInstallStatusReportRequestBuilder) {
    return NewReportsGetUserInstallStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReport provides operations to call the getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport method.
func (m *ReportsRequestBuilder) GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReport()(*ReportsGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    return NewReportsGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWindowsQualityUpdateAlertSummaryReport provides operations to call the getWindowsQualityUpdateAlertSummaryReport method.
func (m *ReportsRequestBuilder) GetWindowsQualityUpdateAlertSummaryReport()(*ReportsGetWindowsQualityUpdateAlertSummaryReportRequestBuilder) {
    return NewReportsGetWindowsQualityUpdateAlertSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWindowsUpdateAlertsPerPolicyPerDeviceReport provides operations to call the getWindowsUpdateAlertsPerPolicyPerDeviceReport method.
func (m *ReportsRequestBuilder) GetWindowsUpdateAlertsPerPolicyPerDeviceReport()(*ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    return NewReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWindowsUpdateAlertSummaryReport provides operations to call the getWindowsUpdateAlertSummaryReport method.
func (m *ReportsRequestBuilder) GetWindowsUpdateAlertSummaryReport()(*ReportsGetWindowsUpdateAlertSummaryReportRequestBuilder) {
    return NewReportsGetWindowsUpdateAlertSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetZebraFotaDeploymentReport provides operations to call the getZebraFotaDeploymentReport method.
func (m *ReportsRequestBuilder) GetZebraFotaDeploymentReport()(*ReportsGetZebraFotaDeploymentReportRequestBuilder) {
    return NewReportsGetZebraFotaDeploymentReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property reports in deviceManagement
func (m *ReportsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable), nil
}
// ToDeleteRequestInformation delete navigation property reports for deviceManagement
func (m *ReportsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation reports singleton
func (m *ReportsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property reports in deviceManagement
func (m *ReportsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
