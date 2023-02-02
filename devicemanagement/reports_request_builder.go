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
    return NewReportsCachedReportConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CachedReportConfigurationsById provides operations to manage the cachedReportConfigurations property of the microsoft.graph.deviceManagementReports entity.
func (m *ReportsRequestBuilder) CachedReportConfigurationsById(id string)(*ReportsCachedReportConfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewReportsCachedReportConfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
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
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ExportJobs provides operations to manage the exportJobs property of the microsoft.graph.deviceManagementReports entity.
func (m *ReportsRequestBuilder) ExportJobs()(*ReportsExportJobsRequestBuilder) {
    return NewReportsExportJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExportJobsById provides operations to manage the exportJobs property of the microsoft.graph.deviceManagementReports entity.
func (m *ReportsRequestBuilder) ExportJobsById(id string)(*ReportsExportJobsDeviceManagementExportJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewReportsExportJobsDeviceManagementExportJobItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable), nil
}
// MicrosoftGraphGetActiveMalwareReport provides operations to call the getActiveMalwareReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetActiveMalwareReport()(*ReportsMicrosoftGraphGetActiveMalwareReportGetActiveMalwareReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetActiveMalwareReportGetActiveMalwareReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetActiveMalwareSummaryReport provides operations to call the getActiveMalwareSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetActiveMalwareSummaryReport()(*ReportsMicrosoftGraphGetActiveMalwareSummaryReportGetActiveMalwareSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetActiveMalwareSummaryReportGetActiveMalwareSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAllCertificatesReport provides operations to call the getAllCertificatesReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAllCertificatesReport()(*ReportsMicrosoftGraphGetAllCertificatesReportGetAllCertificatesReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetAllCertificatesReportGetAllCertificatesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAppsInstallSummaryReport provides operations to call the getAppsInstallSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAppsInstallSummaryReport()(*ReportsMicrosoftGraphGetAppsInstallSummaryReportGetAppsInstallSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetAppsInstallSummaryReportGetAppsInstallSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAppStatusOverviewReport provides operations to call the getAppStatusOverviewReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAppStatusOverviewReport()(*ReportsMicrosoftGraphGetAppStatusOverviewReportGetAppStatusOverviewReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetAppStatusOverviewReportGetAppStatusOverviewReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCachedReport provides operations to call the getCachedReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCachedReport()(*ReportsMicrosoftGraphGetCachedReportGetCachedReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetCachedReportGetCachedReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCertificatesReport provides operations to call the getCertificatesReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCertificatesReport()(*ReportsMicrosoftGraphGetCertificatesReportGetCertificatesReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetCertificatesReportGetCertificatesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCompliancePoliciesReportForDevice provides operations to call the getCompliancePoliciesReportForDevice method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCompliancePoliciesReportForDevice()(*ReportsMicrosoftGraphGetCompliancePoliciesReportForDeviceGetCompliancePoliciesReportForDeviceRequestBuilder) {
    return NewReportsMicrosoftGraphGetCompliancePoliciesReportForDeviceGetCompliancePoliciesReportForDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCompliancePolicyDevicesReport provides operations to call the getCompliancePolicyDevicesReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCompliancePolicyDevicesReport()(*ReportsMicrosoftGraphGetCompliancePolicyDevicesReportGetCompliancePolicyDevicesReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetCompliancePolicyDevicesReportGetCompliancePolicyDevicesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCompliancePolicyDeviceSummaryReport provides operations to call the getCompliancePolicyDeviceSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCompliancePolicyDeviceSummaryReport()(*ReportsMicrosoftGraphGetCompliancePolicyDeviceSummaryReportGetCompliancePolicyDeviceSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetCompliancePolicyDeviceSummaryReportGetCompliancePolicyDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCompliancePolicyNonComplianceReport provides operations to call the getCompliancePolicyNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCompliancePolicyNonComplianceReport()(*ReportsMicrosoftGraphGetCompliancePolicyNonComplianceReportGetCompliancePolicyNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetCompliancePolicyNonComplianceReportGetCompliancePolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCompliancePolicyNonComplianceSummaryReport provides operations to call the getCompliancePolicyNonComplianceSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCompliancePolicyNonComplianceSummaryReport()(*ReportsMicrosoftGraphGetCompliancePolicyNonComplianceSummaryReportGetCompliancePolicyNonComplianceSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetCompliancePolicyNonComplianceSummaryReportGetCompliancePolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetComplianceSettingDetailsReport provides operations to call the getComplianceSettingDetailsReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetComplianceSettingDetailsReport()(*ReportsMicrosoftGraphGetComplianceSettingDetailsReportGetComplianceSettingDetailsReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetComplianceSettingDetailsReportGetComplianceSettingDetailsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetComplianceSettingNonComplianceReport provides operations to call the getComplianceSettingNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetComplianceSettingNonComplianceReport()(*ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetComplianceSettingsReport provides operations to call the getComplianceSettingsReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetComplianceSettingsReport()(*ReportsMicrosoftGraphGetComplianceSettingsReportGetComplianceSettingsReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetComplianceSettingsReportGetComplianceSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetConfigManagerDevicePolicyStatusReport provides operations to call the getConfigManagerDevicePolicyStatusReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetConfigManagerDevicePolicyStatusReport()(*ReportsMicrosoftGraphGetConfigManagerDevicePolicyStatusReportGetConfigManagerDevicePolicyStatusReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetConfigManagerDevicePolicyStatusReportGetConfigManagerDevicePolicyStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetConfigurationPoliciesReportForDevice provides operations to call the getConfigurationPoliciesReportForDevice method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetConfigurationPoliciesReportForDevice()(*ReportsMicrosoftGraphGetConfigurationPoliciesReportForDeviceGetConfigurationPoliciesReportForDeviceRequestBuilder) {
    return NewReportsMicrosoftGraphGetConfigurationPoliciesReportForDeviceGetConfigurationPoliciesReportForDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetConfigurationPolicyDevicesReport provides operations to call the getConfigurationPolicyDevicesReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetConfigurationPolicyDevicesReport()(*ReportsMicrosoftGraphGetConfigurationPolicyDevicesReportGetConfigurationPolicyDevicesReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetConfigurationPolicyDevicesReportGetConfigurationPolicyDevicesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetConfigurationPolicyDeviceSummaryReport provides operations to call the getConfigurationPolicyDeviceSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetConfigurationPolicyDeviceSummaryReport()(*ReportsMicrosoftGraphGetConfigurationPolicyDeviceSummaryReportGetConfigurationPolicyDeviceSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetConfigurationPolicyDeviceSummaryReportGetConfigurationPolicyDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetConfigurationPolicyNonComplianceReport provides operations to call the getConfigurationPolicyNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetConfigurationPolicyNonComplianceReport()(*ReportsMicrosoftGraphGetConfigurationPolicyNonComplianceReportGetConfigurationPolicyNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetConfigurationPolicyNonComplianceReportGetConfigurationPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetConfigurationPolicyNonComplianceSummaryReport provides operations to call the getConfigurationPolicyNonComplianceSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetConfigurationPolicyNonComplianceSummaryReport()(*ReportsMicrosoftGraphGetConfigurationPolicyNonComplianceSummaryReportGetConfigurationPolicyNonComplianceSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetConfigurationPolicyNonComplianceSummaryReportGetConfigurationPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetConfigurationPolicySettingsDeviceSummaryReport provides operations to call the getConfigurationPolicySettingsDeviceSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetConfigurationPolicySettingsDeviceSummaryReport()(*ReportsMicrosoftGraphGetConfigurationPolicySettingsDeviceSummaryReportGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetConfigurationPolicySettingsDeviceSummaryReportGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetConfigurationSettingDetailsReport provides operations to call the getConfigurationSettingDetailsReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetConfigurationSettingDetailsReport()(*ReportsMicrosoftGraphGetConfigurationSettingDetailsReportGetConfigurationSettingDetailsReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetConfigurationSettingDetailsReportGetConfigurationSettingDetailsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetConfigurationSettingNonComplianceReport provides operations to call the getConfigurationSettingNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetConfigurationSettingNonComplianceReport()(*ReportsMicrosoftGraphGetConfigurationSettingNonComplianceReportGetConfigurationSettingNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetConfigurationSettingNonComplianceReportGetConfigurationSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetConfigurationSettingsReport provides operations to call the getConfigurationSettingsReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetConfigurationSettingsReport()(*ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetDeviceConfigurationPolicySettingsSummaryReport provides operations to call the getDeviceConfigurationPolicySettingsSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetDeviceConfigurationPolicySettingsSummaryReport()(*ReportsMicrosoftGraphGetDeviceConfigurationPolicySettingsSummaryReportGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetDeviceConfigurationPolicySettingsSummaryReportGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetDeviceConfigurationPolicyStatusSummary provides operations to call the getDeviceConfigurationPolicyStatusSummary method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetDeviceConfigurationPolicyStatusSummary()(*ReportsMicrosoftGraphGetDeviceConfigurationPolicyStatusSummaryGetDeviceConfigurationPolicyStatusSummaryRequestBuilder) {
    return NewReportsMicrosoftGraphGetDeviceConfigurationPolicyStatusSummaryGetDeviceConfigurationPolicyStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetDeviceInstallStatusReport provides operations to call the getDeviceInstallStatusReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetDeviceInstallStatusReport()(*ReportsMicrosoftGraphGetDeviceInstallStatusReportGetDeviceInstallStatusReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetDeviceInstallStatusReportGetDeviceInstallStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfiles provides operations to call the getDeviceManagementIntentPerSettingContributingProfiles method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfiles()(*ReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    return NewReportsMicrosoftGraphGetDeviceManagementIntentPerSettingContributingProfilesGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetDeviceManagementIntentSettingsReport provides operations to call the getDeviceManagementIntentSettingsReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetDeviceManagementIntentSettingsReport()(*ReportsMicrosoftGraphGetDeviceManagementIntentSettingsReportGetDeviceManagementIntentSettingsReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetDeviceManagementIntentSettingsReportGetDeviceManagementIntentSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetDeviceNonComplianceReport provides operations to call the getDeviceNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetDeviceNonComplianceReport()(*ReportsMicrosoftGraphGetDeviceNonComplianceReportGetDeviceNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetDeviceNonComplianceReportGetDeviceNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetDevicesWithoutCompliancePolicyReport provides operations to call the getDevicesWithoutCompliancePolicyReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetDevicesWithoutCompliancePolicyReport()(*ReportsMicrosoftGraphGetDevicesWithoutCompliancePolicyReportGetDevicesWithoutCompliancePolicyReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetDevicesWithoutCompliancePolicyReportGetDevicesWithoutCompliancePolicyReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetEncryptionReportForDevices provides operations to call the getEncryptionReportForDevices method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEncryptionReportForDevices()(*ReportsMicrosoftGraphGetEncryptionReportForDevicesGetEncryptionReportForDevicesRequestBuilder) {
    return NewReportsMicrosoftGraphGetEncryptionReportForDevicesGetEncryptionReportForDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetEnrollmentConfigurationPoliciesByDevice provides operations to call the getEnrollmentConfigurationPoliciesByDevice method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEnrollmentConfigurationPoliciesByDevice()(*ReportsMicrosoftGraphGetEnrollmentConfigurationPoliciesByDeviceGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) {
    return NewReportsMicrosoftGraphGetEnrollmentConfigurationPoliciesByDeviceGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetFailedMobileAppsReport provides operations to call the getFailedMobileAppsReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFailedMobileAppsReport()(*ReportsMicrosoftGraphGetFailedMobileAppsReportGetFailedMobileAppsReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetFailedMobileAppsReportGetFailedMobileAppsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetFailedMobileAppsSummaryReport provides operations to call the getFailedMobileAppsSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFailedMobileAppsSummaryReport()(*ReportsMicrosoftGraphGetFailedMobileAppsSummaryReportGetFailedMobileAppsSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetFailedMobileAppsSummaryReportGetFailedMobileAppsSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetGroupPolicySettingsDeviceSettingsReport provides operations to call the getGroupPolicySettingsDeviceSettingsReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetGroupPolicySettingsDeviceSettingsReport()(*ReportsMicrosoftGraphGetGroupPolicySettingsDeviceSettingsReportGetGroupPolicySettingsDeviceSettingsReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetGroupPolicySettingsDeviceSettingsReportGetGroupPolicySettingsDeviceSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetHistoricalReport provides operations to call the getHistoricalReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetHistoricalReport()(*ReportsMicrosoftGraphGetHistoricalReportGetHistoricalReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetHistoricalReportGetHistoricalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMalwareSummaryReport provides operations to call the getMalwareSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMalwareSummaryReport()(*ReportsMicrosoftGraphGetMalwareSummaryReportGetMalwareSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetMalwareSummaryReportGetMalwareSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMobileApplicationManagementAppConfigurationReport provides operations to call the getMobileApplicationManagementAppConfigurationReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMobileApplicationManagementAppConfigurationReport()(*ReportsMicrosoftGraphGetMobileApplicationManagementAppConfigurationReportGetMobileApplicationManagementAppConfigurationReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetMobileApplicationManagementAppConfigurationReportGetMobileApplicationManagementAppConfigurationReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetMobileApplicationManagementAppRegistrationSummaryReport provides operations to call the getMobileApplicationManagementAppRegistrationSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMobileApplicationManagementAppRegistrationSummaryReport()(*ReportsMicrosoftGraphGetMobileApplicationManagementAppRegistrationSummaryReportGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetMobileApplicationManagementAppRegistrationSummaryReportGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetNoncompliantDevicesAndSettingsReport provides operations to call the getNoncompliantDevicesAndSettingsReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetNoncompliantDevicesAndSettingsReport()(*ReportsMicrosoftGraphGetNoncompliantDevicesAndSettingsReportGetNoncompliantDevicesAndSettingsReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetNoncompliantDevicesAndSettingsReportGetNoncompliantDevicesAndSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetPolicyNonComplianceMetadata provides operations to call the getPolicyNonComplianceMetadata method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetPolicyNonComplianceMetadata()(*ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilder) {
    return NewReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetPolicyNonComplianceReport provides operations to call the getPolicyNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetPolicyNonComplianceReport()(*ReportsMicrosoftGraphGetPolicyNonComplianceReportGetPolicyNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetPolicyNonComplianceReportGetPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetPolicyNonComplianceSummaryReport provides operations to call the getPolicyNonComplianceSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetPolicyNonComplianceSummaryReport()(*ReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportGetPolicyNonComplianceSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetPolicyNonComplianceSummaryReportGetPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetQuietTimePolicyUsersReport provides operations to call the getQuietTimePolicyUsersReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetQuietTimePolicyUsersReport()(*ReportsMicrosoftGraphGetQuietTimePolicyUsersReportGetQuietTimePolicyUsersReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetQuietTimePolicyUsersReportGetQuietTimePolicyUsersReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetQuietTimePolicyUserSummaryReport provides operations to call the getQuietTimePolicyUserSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetQuietTimePolicyUserSummaryReport()(*ReportsMicrosoftGraphGetQuietTimePolicyUserSummaryReportGetQuietTimePolicyUserSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetQuietTimePolicyUserSummaryReportGetQuietTimePolicyUserSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetRelatedAppsStatusReport provides operations to call the getRelatedAppsStatusReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetRelatedAppsStatusReport()(*ReportsMicrosoftGraphGetRelatedAppsStatusReportGetRelatedAppsStatusReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetRelatedAppsStatusReportGetRelatedAppsStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetRemoteAssistanceSessionsReport provides operations to call the getRemoteAssistanceSessionsReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetRemoteAssistanceSessionsReport()(*ReportsMicrosoftGraphGetRemoteAssistanceSessionsReportGetRemoteAssistanceSessionsReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetRemoteAssistanceSessionsReportGetRemoteAssistanceSessionsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetReportFilters provides operations to call the getReportFilters method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetReportFilters()(*ReportsMicrosoftGraphGetReportFiltersGetReportFiltersRequestBuilder) {
    return NewReportsMicrosoftGraphGetReportFiltersGetReportFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetSettingNonComplianceReport provides operations to call the getSettingNonComplianceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSettingNonComplianceReport()(*ReportsMicrosoftGraphGetSettingNonComplianceReportGetSettingNonComplianceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetSettingNonComplianceReportGetSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetUnhealthyDefenderAgentsReport provides operations to call the getUnhealthyDefenderAgentsReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetUnhealthyDefenderAgentsReport()(*ReportsMicrosoftGraphGetUnhealthyDefenderAgentsReportGetUnhealthyDefenderAgentsReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetUnhealthyDefenderAgentsReportGetUnhealthyDefenderAgentsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetUnhealthyFirewallReport provides operations to call the getUnhealthyFirewallReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetUnhealthyFirewallReport()(*ReportsMicrosoftGraphGetUnhealthyFirewallReportGetUnhealthyFirewallReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetUnhealthyFirewallReportGetUnhealthyFirewallReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetUnhealthyFirewallSummaryReport provides operations to call the getUnhealthyFirewallSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetUnhealthyFirewallSummaryReport()(*ReportsMicrosoftGraphGetUnhealthyFirewallSummaryReportGetUnhealthyFirewallSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetUnhealthyFirewallSummaryReportGetUnhealthyFirewallSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetUserInstallStatusReport provides operations to call the getUserInstallStatusReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetUserInstallStatusReport()(*ReportsMicrosoftGraphGetUserInstallStatusReportGetUserInstallStatusReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetUserInstallStatusReportGetUserInstallStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReport provides operations to call the getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReport()(*ReportsMicrosoftGraphGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetWindowsQualityUpdateAlertSummaryReport provides operations to call the getWindowsQualityUpdateAlertSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetWindowsQualityUpdateAlertSummaryReport()(*ReportsMicrosoftGraphGetWindowsQualityUpdateAlertSummaryReportGetWindowsQualityUpdateAlertSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetWindowsQualityUpdateAlertSummaryReportGetWindowsQualityUpdateAlertSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetWindowsUpdateAlertsPerPolicyPerDeviceReport provides operations to call the getWindowsUpdateAlertsPerPolicyPerDeviceReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetWindowsUpdateAlertsPerPolicyPerDeviceReport()(*ReportsMicrosoftGraphGetWindowsUpdateAlertsPerPolicyPerDeviceReportGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetWindowsUpdateAlertsPerPolicyPerDeviceReportGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetWindowsUpdateAlertSummaryReport provides operations to call the getWindowsUpdateAlertSummaryReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetWindowsUpdateAlertSummaryReport()(*ReportsMicrosoftGraphGetWindowsUpdateAlertSummaryReportGetWindowsUpdateAlertSummaryReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetWindowsUpdateAlertSummaryReportGetWindowsUpdateAlertSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetZebraFotaDeploymentReport provides operations to call the getZebraFotaDeploymentReport method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetZebraFotaDeploymentReport()(*ReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportRequestBuilder) {
    return NewReportsMicrosoftGraphGetZebraFotaDeploymentReportGetZebraFotaDeploymentReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
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
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementReportsFromDiscriminatorValue, errorMapping)
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
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
