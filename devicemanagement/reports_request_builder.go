package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.deviceManagement entity.
type ReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
// returns a *ReportsCachedReportConfigurationsRequestBuilder when successful
func (m *ReportsRequestBuilder) CachedReportConfigurations()(*ReportsCachedReportConfigurationsRequestBuilder) {
    return NewReportsCachedReportConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property reports for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ExportJobs provides operations to manage the exportJobs property of the microsoft.graph.deviceManagementReports entity.
// returns a *ReportsExportJobsRequestBuilder when successful
func (m *ReportsRequestBuilder) ExportJobs()(*ReportsExportJobsRequestBuilder) {
    return NewReportsExportJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get reports singleton
// returns a DeviceManagementReportsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable), nil
}
// GetActiveMalwareReport provides operations to call the getActiveMalwareReport method.
// returns a *ReportsGetActiveMalwareReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetActiveMalwareReport()(*ReportsGetActiveMalwareReportRequestBuilder) {
    return NewReportsGetActiveMalwareReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetActiveMalwareSummaryReport provides operations to call the getActiveMalwareSummaryReport method.
// returns a *ReportsGetActiveMalwareSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetActiveMalwareSummaryReport()(*ReportsGetActiveMalwareSummaryReportRequestBuilder) {
    return NewReportsGetActiveMalwareSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetAllCertificatesReport provides operations to call the getAllCertificatesReport method.
// returns a *ReportsGetAllCertificatesReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetAllCertificatesReport()(*ReportsGetAllCertificatesReportRequestBuilder) {
    return NewReportsGetAllCertificatesReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetAppsInstallSummaryReport provides operations to call the getAppsInstallSummaryReport method.
// returns a *ReportsGetAppsInstallSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetAppsInstallSummaryReport()(*ReportsGetAppsInstallSummaryReportRequestBuilder) {
    return NewReportsGetAppsInstallSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetAppStatusOverviewReport provides operations to call the getAppStatusOverviewReport method.
// returns a *ReportsGetAppStatusOverviewReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetAppStatusOverviewReport()(*ReportsGetAppStatusOverviewReportRequestBuilder) {
    return NewReportsGetAppStatusOverviewReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCachedReport provides operations to call the getCachedReport method.
// returns a *ReportsGetCachedReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetCachedReport()(*ReportsGetCachedReportRequestBuilder) {
    return NewReportsGetCachedReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCertificatesReport provides operations to call the getCertificatesReport method.
// returns a *ReportsGetCertificatesReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetCertificatesReport()(*ReportsGetCertificatesReportRequestBuilder) {
    return NewReportsGetCertificatesReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCompliancePoliciesReportForDevice provides operations to call the getCompliancePoliciesReportForDevice method.
// returns a *ReportsGetCompliancePoliciesReportForDeviceRequestBuilder when successful
func (m *ReportsRequestBuilder) GetCompliancePoliciesReportForDevice()(*ReportsGetCompliancePoliciesReportForDeviceRequestBuilder) {
    return NewReportsGetCompliancePoliciesReportForDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCompliancePolicyDevicesReport provides operations to call the getCompliancePolicyDevicesReport method.
// returns a *ReportsGetCompliancePolicyDevicesReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetCompliancePolicyDevicesReport()(*ReportsGetCompliancePolicyDevicesReportRequestBuilder) {
    return NewReportsGetCompliancePolicyDevicesReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCompliancePolicyDeviceSummaryReport provides operations to call the getCompliancePolicyDeviceSummaryReport method.
// returns a *ReportsGetCompliancePolicyDeviceSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetCompliancePolicyDeviceSummaryReport()(*ReportsGetCompliancePolicyDeviceSummaryReportRequestBuilder) {
    return NewReportsGetCompliancePolicyDeviceSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCompliancePolicyNonComplianceReport provides operations to call the getCompliancePolicyNonComplianceReport method.
// returns a *ReportsGetCompliancePolicyNonComplianceReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetCompliancePolicyNonComplianceReport()(*ReportsGetCompliancePolicyNonComplianceReportRequestBuilder) {
    return NewReportsGetCompliancePolicyNonComplianceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCompliancePolicyNonComplianceSummaryReport provides operations to call the getCompliancePolicyNonComplianceSummaryReport method.
// returns a *ReportsGetCompliancePolicyNonComplianceSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetCompliancePolicyNonComplianceSummaryReport()(*ReportsGetCompliancePolicyNonComplianceSummaryReportRequestBuilder) {
    return NewReportsGetCompliancePolicyNonComplianceSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetComplianceSettingDetailsReport provides operations to call the getComplianceSettingDetailsReport method.
// returns a *ReportsGetComplianceSettingDetailsReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetComplianceSettingDetailsReport()(*ReportsGetComplianceSettingDetailsReportRequestBuilder) {
    return NewReportsGetComplianceSettingDetailsReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetComplianceSettingNonComplianceReport provides operations to call the getComplianceSettingNonComplianceReport method.
// returns a *ReportsGetComplianceSettingNonComplianceReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetComplianceSettingNonComplianceReport()(*ReportsGetComplianceSettingNonComplianceReportRequestBuilder) {
    return NewReportsGetComplianceSettingNonComplianceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetComplianceSettingsReport provides operations to call the getComplianceSettingsReport method.
// returns a *ReportsGetComplianceSettingsReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetComplianceSettingsReport()(*ReportsGetComplianceSettingsReportRequestBuilder) {
    return NewReportsGetComplianceSettingsReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetConfigManagerDevicePolicyStatusReport provides operations to call the getConfigManagerDevicePolicyStatusReport method.
// returns a *ReportsGetConfigManagerDevicePolicyStatusReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetConfigManagerDevicePolicyStatusReport()(*ReportsGetConfigManagerDevicePolicyStatusReportRequestBuilder) {
    return NewReportsGetConfigManagerDevicePolicyStatusReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetConfigurationPoliciesReportForDevice provides operations to call the getConfigurationPoliciesReportForDevice method.
// returns a *ReportsGetConfigurationPoliciesReportForDeviceRequestBuilder when successful
func (m *ReportsRequestBuilder) GetConfigurationPoliciesReportForDevice()(*ReportsGetConfigurationPoliciesReportForDeviceRequestBuilder) {
    return NewReportsGetConfigurationPoliciesReportForDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetConfigurationPolicyDevicesReport provides operations to call the getConfigurationPolicyDevicesReport method.
// returns a *ReportsGetConfigurationPolicyDevicesReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetConfigurationPolicyDevicesReport()(*ReportsGetConfigurationPolicyDevicesReportRequestBuilder) {
    return NewReportsGetConfigurationPolicyDevicesReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetConfigurationPolicyDeviceSummaryReport provides operations to call the getConfigurationPolicyDeviceSummaryReport method.
// returns a *ReportsGetConfigurationPolicyDeviceSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetConfigurationPolicyDeviceSummaryReport()(*ReportsGetConfigurationPolicyDeviceSummaryReportRequestBuilder) {
    return NewReportsGetConfigurationPolicyDeviceSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetConfigurationPolicyNonComplianceReport provides operations to call the getConfigurationPolicyNonComplianceReport method.
// returns a *ReportsGetConfigurationPolicyNonComplianceReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetConfigurationPolicyNonComplianceReport()(*ReportsGetConfigurationPolicyNonComplianceReportRequestBuilder) {
    return NewReportsGetConfigurationPolicyNonComplianceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetConfigurationPolicyNonComplianceSummaryReport provides operations to call the getConfigurationPolicyNonComplianceSummaryReport method.
// returns a *ReportsGetConfigurationPolicyNonComplianceSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetConfigurationPolicyNonComplianceSummaryReport()(*ReportsGetConfigurationPolicyNonComplianceSummaryReportRequestBuilder) {
    return NewReportsGetConfigurationPolicyNonComplianceSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetConfigurationPolicySettingsDeviceSummaryReport provides operations to call the getConfigurationPolicySettingsDeviceSummaryReport method.
// returns a *ReportsGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetConfigurationPolicySettingsDeviceSummaryReport()(*ReportsGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder) {
    return NewReportsGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetConfigurationSettingDetailsReport provides operations to call the getConfigurationSettingDetailsReport method.
// returns a *ReportsGetConfigurationSettingDetailsReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetConfigurationSettingDetailsReport()(*ReportsGetConfigurationSettingDetailsReportRequestBuilder) {
    return NewReportsGetConfigurationSettingDetailsReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetConfigurationSettingNonComplianceReport provides operations to call the getConfigurationSettingNonComplianceReport method.
// returns a *ReportsGetConfigurationSettingNonComplianceReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetConfigurationSettingNonComplianceReport()(*ReportsGetConfigurationSettingNonComplianceReportRequestBuilder) {
    return NewReportsGetConfigurationSettingNonComplianceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetConfigurationSettingsReport provides operations to call the getConfigurationSettingsReport method.
// returns a *ReportsGetConfigurationSettingsReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetConfigurationSettingsReport()(*ReportsGetConfigurationSettingsReportRequestBuilder) {
    return NewReportsGetConfigurationSettingsReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDeviceConfigurationPolicySettingsSummaryReport provides operations to call the getDeviceConfigurationPolicySettingsSummaryReport method.
// returns a *ReportsGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDeviceConfigurationPolicySettingsSummaryReport()(*ReportsGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) {
    return NewReportsGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDeviceConfigurationPolicyStatusSummary provides operations to call the getDeviceConfigurationPolicyStatusSummary method.
// returns a *ReportsGetDeviceConfigurationPolicyStatusSummaryRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDeviceConfigurationPolicyStatusSummary()(*ReportsGetDeviceConfigurationPolicyStatusSummaryRequestBuilder) {
    return NewReportsGetDeviceConfigurationPolicyStatusSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDeviceInstallStatusReport provides operations to call the getDeviceInstallStatusReport method.
// returns a *ReportsGetDeviceInstallStatusReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDeviceInstallStatusReport()(*ReportsGetDeviceInstallStatusReportRequestBuilder) {
    return NewReportsGetDeviceInstallStatusReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDeviceManagementIntentPerSettingContributingProfiles provides operations to call the getDeviceManagementIntentPerSettingContributingProfiles method.
// returns a *ReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDeviceManagementIntentPerSettingContributingProfiles()(*ReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    return NewReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDeviceManagementIntentSettingsReport provides operations to call the getDeviceManagementIntentSettingsReport method.
// returns a *ReportsGetDeviceManagementIntentSettingsReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDeviceManagementIntentSettingsReport()(*ReportsGetDeviceManagementIntentSettingsReportRequestBuilder) {
    return NewReportsGetDeviceManagementIntentSettingsReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDeviceNonComplianceReport provides operations to call the getDeviceNonComplianceReport method.
// returns a *ReportsGetDeviceNonComplianceReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDeviceNonComplianceReport()(*ReportsGetDeviceNonComplianceReportRequestBuilder) {
    return NewReportsGetDeviceNonComplianceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDevicePoliciesComplianceReport provides operations to call the getDevicePoliciesComplianceReport method.
// returns a *ReportsGetDevicePoliciesComplianceReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDevicePoliciesComplianceReport()(*ReportsGetDevicePoliciesComplianceReportRequestBuilder) {
    return NewReportsGetDevicePoliciesComplianceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDevicePolicySettingsComplianceReport provides operations to call the getDevicePolicySettingsComplianceReport method.
// returns a *ReportsGetDevicePolicySettingsComplianceReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDevicePolicySettingsComplianceReport()(*ReportsGetDevicePolicySettingsComplianceReportRequestBuilder) {
    return NewReportsGetDevicePolicySettingsComplianceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDevicesStatusByPolicyPlatformComplianceReport provides operations to call the getDevicesStatusByPolicyPlatformComplianceReport method.
// returns a *ReportsGetDevicesStatusByPolicyPlatformComplianceReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDevicesStatusByPolicyPlatformComplianceReport()(*ReportsGetDevicesStatusByPolicyPlatformComplianceReportRequestBuilder) {
    return NewReportsGetDevicesStatusByPolicyPlatformComplianceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDevicesStatusBySettingReport provides operations to call the getDevicesStatusBySettingReport method.
// returns a *ReportsGetDevicesStatusBySettingReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDevicesStatusBySettingReport()(*ReportsGetDevicesStatusBySettingReportRequestBuilder) {
    return NewReportsGetDevicesStatusBySettingReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDeviceStatusByCompliacePolicyReport provides operations to call the getDeviceStatusByCompliacePolicyReport method.
// returns a *ReportsGetDeviceStatusByCompliacePolicyReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDeviceStatusByCompliacePolicyReport()(*ReportsGetDeviceStatusByCompliacePolicyReportRequestBuilder) {
    return NewReportsGetDeviceStatusByCompliacePolicyReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDeviceStatusByCompliancePolicySettingReport provides operations to call the getDeviceStatusByCompliancePolicySettingReport method.
// returns a *ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDeviceStatusByCompliancePolicySettingReport()(*ReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilder) {
    return NewReportsGetDeviceStatusByCompliancePolicySettingReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDeviceStatusSummaryByCompliacePolicyReport provides operations to call the getDeviceStatusSummaryByCompliacePolicyReport method.
// returns a *ReportsGetDeviceStatusSummaryByCompliacePolicyReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDeviceStatusSummaryByCompliacePolicyReport()(*ReportsGetDeviceStatusSummaryByCompliacePolicyReportRequestBuilder) {
    return NewReportsGetDeviceStatusSummaryByCompliacePolicyReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDeviceStatusSummaryByCompliancePolicySettingsReport provides operations to call the getDeviceStatusSummaryByCompliancePolicySettingsReport method.
// returns a *ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDeviceStatusSummaryByCompliancePolicySettingsReport()(*ReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilder) {
    return NewReportsGetDeviceStatusSummaryByCompliancePolicySettingsReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetDevicesWithoutCompliancePolicyReport provides operations to call the getDevicesWithoutCompliancePolicyReport method.
// returns a *ReportsGetDevicesWithoutCompliancePolicyReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetDevicesWithoutCompliancePolicyReport()(*ReportsGetDevicesWithoutCompliancePolicyReportRequestBuilder) {
    return NewReportsGetDevicesWithoutCompliancePolicyReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetEncryptionReportForDevices provides operations to call the getEncryptionReportForDevices method.
// returns a *ReportsGetEncryptionReportForDevicesRequestBuilder when successful
func (m *ReportsRequestBuilder) GetEncryptionReportForDevices()(*ReportsGetEncryptionReportForDevicesRequestBuilder) {
    return NewReportsGetEncryptionReportForDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetEnrollmentConfigurationPoliciesByDevice provides operations to call the getEnrollmentConfigurationPoliciesByDevice method.
// returns a *ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder when successful
func (m *ReportsRequestBuilder) GetEnrollmentConfigurationPoliciesByDevice()(*ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) {
    return NewReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFailedMobileAppsReport provides operations to call the getFailedMobileAppsReport method.
// returns a *ReportsGetFailedMobileAppsReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetFailedMobileAppsReport()(*ReportsGetFailedMobileAppsReportRequestBuilder) {
    return NewReportsGetFailedMobileAppsReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFailedMobileAppsSummaryReport provides operations to call the getFailedMobileAppsSummaryReport method.
// returns a *ReportsGetFailedMobileAppsSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetFailedMobileAppsSummaryReport()(*ReportsGetFailedMobileAppsSummaryReportRequestBuilder) {
    return NewReportsGetFailedMobileAppsSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetGroupPolicySettingsDeviceSettingsReport provides operations to call the getGroupPolicySettingsDeviceSettingsReport method.
// returns a *ReportsGetGroupPolicySettingsDeviceSettingsReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetGroupPolicySettingsDeviceSettingsReport()(*ReportsGetGroupPolicySettingsDeviceSettingsReportRequestBuilder) {
    return NewReportsGetGroupPolicySettingsDeviceSettingsReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetHistoricalReport provides operations to call the getHistoricalReport method.
// returns a *ReportsGetHistoricalReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetHistoricalReport()(*ReportsGetHistoricalReportRequestBuilder) {
    return NewReportsGetHistoricalReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetMalwareSummaryReport provides operations to call the getMalwareSummaryReport method.
// returns a *ReportsGetMalwareSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetMalwareSummaryReport()(*ReportsGetMalwareSummaryReportRequestBuilder) {
    return NewReportsGetMalwareSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetMobileApplicationManagementAppConfigurationReport provides operations to call the getMobileApplicationManagementAppConfigurationReport method.
// returns a *ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetMobileApplicationManagementAppConfigurationReport()(*ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilder) {
    return NewReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetMobileApplicationManagementAppRegistrationSummaryReport provides operations to call the getMobileApplicationManagementAppRegistrationSummaryReport method.
// returns a *ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetMobileApplicationManagementAppRegistrationSummaryReport()(*ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder) {
    return NewReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetNoncompliantDevicesAndSettingsReport provides operations to call the getNoncompliantDevicesAndSettingsReport method.
// returns a *ReportsGetNoncompliantDevicesAndSettingsReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetNoncompliantDevicesAndSettingsReport()(*ReportsGetNoncompliantDevicesAndSettingsReportRequestBuilder) {
    return NewReportsGetNoncompliantDevicesAndSettingsReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetPolicyNonComplianceMetadata provides operations to call the getPolicyNonComplianceMetadata method.
// returns a *ReportsGetPolicyNonComplianceMetadataRequestBuilder when successful
func (m *ReportsRequestBuilder) GetPolicyNonComplianceMetadata()(*ReportsGetPolicyNonComplianceMetadataRequestBuilder) {
    return NewReportsGetPolicyNonComplianceMetadataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetPolicyNonComplianceReport provides operations to call the getPolicyNonComplianceReport method.
// returns a *ReportsGetPolicyNonComplianceReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetPolicyNonComplianceReport()(*ReportsGetPolicyNonComplianceReportRequestBuilder) {
    return NewReportsGetPolicyNonComplianceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetPolicyNonComplianceSummaryReport provides operations to call the getPolicyNonComplianceSummaryReport method.
// returns a *ReportsGetPolicyNonComplianceSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetPolicyNonComplianceSummaryReport()(*ReportsGetPolicyNonComplianceSummaryReportRequestBuilder) {
    return NewReportsGetPolicyNonComplianceSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetQuietTimePolicyUsersReport provides operations to call the getQuietTimePolicyUsersReport method.
// returns a *ReportsGetQuietTimePolicyUsersReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetQuietTimePolicyUsersReport()(*ReportsGetQuietTimePolicyUsersReportRequestBuilder) {
    return NewReportsGetQuietTimePolicyUsersReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetQuietTimePolicyUserSummaryReport provides operations to call the getQuietTimePolicyUserSummaryReport method.
// returns a *ReportsGetQuietTimePolicyUserSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetQuietTimePolicyUserSummaryReport()(*ReportsGetQuietTimePolicyUserSummaryReportRequestBuilder) {
    return NewReportsGetQuietTimePolicyUserSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetRelatedAppsStatusReport provides operations to call the getRelatedAppsStatusReport method.
// returns a *ReportsGetRelatedAppsStatusReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetRelatedAppsStatusReport()(*ReportsGetRelatedAppsStatusReportRequestBuilder) {
    return NewReportsGetRelatedAppsStatusReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetRemoteAssistanceSessionsReport provides operations to call the getRemoteAssistanceSessionsReport method.
// returns a *ReportsGetRemoteAssistanceSessionsReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetRemoteAssistanceSessionsReport()(*ReportsGetRemoteAssistanceSessionsReportRequestBuilder) {
    return NewReportsGetRemoteAssistanceSessionsReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetReportFilters provides operations to call the getReportFilters method.
// returns a *ReportsGetReportFiltersRequestBuilder when successful
func (m *ReportsRequestBuilder) GetReportFilters()(*ReportsGetReportFiltersRequestBuilder) {
    return NewReportsGetReportFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetSettingNonComplianceReport provides operations to call the getSettingNonComplianceReport method.
// returns a *ReportsGetSettingNonComplianceReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetSettingNonComplianceReport()(*ReportsGetSettingNonComplianceReportRequestBuilder) {
    return NewReportsGetSettingNonComplianceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetUnhealthyDefenderAgentsReport provides operations to call the getUnhealthyDefenderAgentsReport method.
// returns a *ReportsGetUnhealthyDefenderAgentsReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetUnhealthyDefenderAgentsReport()(*ReportsGetUnhealthyDefenderAgentsReportRequestBuilder) {
    return NewReportsGetUnhealthyDefenderAgentsReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetUnhealthyFirewallReport provides operations to call the getUnhealthyFirewallReport method.
// returns a *ReportsGetUnhealthyFirewallReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetUnhealthyFirewallReport()(*ReportsGetUnhealthyFirewallReportRequestBuilder) {
    return NewReportsGetUnhealthyFirewallReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetUnhealthyFirewallSummaryReport provides operations to call the getUnhealthyFirewallSummaryReport method.
// returns a *ReportsGetUnhealthyFirewallSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetUnhealthyFirewallSummaryReport()(*ReportsGetUnhealthyFirewallSummaryReportRequestBuilder) {
    return NewReportsGetUnhealthyFirewallSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetUserInstallStatusReport provides operations to call the getUserInstallStatusReport method.
// returns a *ReportsGetUserInstallStatusReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetUserInstallStatusReport()(*ReportsGetUserInstallStatusReportRequestBuilder) {
    return NewReportsGetUserInstallStatusReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetWindowsDriverUpdateAlertsPerPolicyPerDeviceReport provides operations to call the getWindowsDriverUpdateAlertsPerPolicyPerDeviceReport method.
// returns a *ReportsGetWindowsDriverUpdateAlertsPerPolicyPerDeviceReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetWindowsDriverUpdateAlertsPerPolicyPerDeviceReport()(*ReportsGetWindowsDriverUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    return NewReportsGetWindowsDriverUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetWindowsDriverUpdateAlertSummaryReport provides operations to call the getWindowsDriverUpdateAlertSummaryReport method.
// returns a *ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetWindowsDriverUpdateAlertSummaryReport()(*ReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilder) {
    return NewReportsGetWindowsDriverUpdateAlertSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReport provides operations to call the getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport method.
// returns a *ReportsGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReport()(*ReportsGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    return NewReportsGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetWindowsQualityUpdateAlertSummaryReport provides operations to call the getWindowsQualityUpdateAlertSummaryReport method.
// returns a *ReportsGetWindowsQualityUpdateAlertSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetWindowsQualityUpdateAlertSummaryReport()(*ReportsGetWindowsQualityUpdateAlertSummaryReportRequestBuilder) {
    return NewReportsGetWindowsQualityUpdateAlertSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetWindowsUpdateAlertsPerPolicyPerDeviceReport provides operations to call the getWindowsUpdateAlertsPerPolicyPerDeviceReport method.
// returns a *ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetWindowsUpdateAlertsPerPolicyPerDeviceReport()(*ReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    return NewReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetWindowsUpdateAlertSummaryReport provides operations to call the getWindowsUpdateAlertSummaryReport method.
// returns a *ReportsGetWindowsUpdateAlertSummaryReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetWindowsUpdateAlertSummaryReport()(*ReportsGetWindowsUpdateAlertSummaryReportRequestBuilder) {
    return NewReportsGetWindowsUpdateAlertSummaryReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetZebraFotaDeploymentReport provides operations to call the getZebraFotaDeploymentReport method.
// returns a *ReportsGetZebraFotaDeploymentReportRequestBuilder when successful
func (m *ReportsRequestBuilder) GetZebraFotaDeploymentReport()(*ReportsGetZebraFotaDeploymentReportRequestBuilder) {
    return NewReportsGetZebraFotaDeploymentReportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property reports in deviceManagement
// returns a DeviceManagementReportsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable), nil
}
// ToDeleteRequestInformation delete navigation property reports for deviceManagement
// returns a *RequestInformation when successful
func (m *ReportsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation reports singleton
// returns a *RequestInformation when successful
func (m *ReportsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property reports in deviceManagement
// returns a *RequestInformation when successful
func (m *ReportsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ReportsRequestBuilder when successful
func (m *ReportsRequestBuilder) WithUrl(rawUrl string)(*ReportsRequestBuilder) {
    return NewReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
