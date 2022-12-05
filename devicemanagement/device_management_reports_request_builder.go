package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.deviceManagement entity.
type DeviceManagementReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementReportsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReportsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementReportsRequestBuilderGetQueryParameters reports singleton
type DeviceManagementReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementReportsRequestBuilderGetQueryParameters
}
// DeviceManagementReportsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReportsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CachedReportConfigurations provides operations to manage the cachedReportConfigurations property of the microsoft.graph.deviceManagementReports entity.
func (m *DeviceManagementReportsRequestBuilder) CachedReportConfigurations()(*DeviceManagementReportsCachedReportConfigurationsRequestBuilder) {
    return NewDeviceManagementReportsCachedReportConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CachedReportConfigurationsById provides operations to manage the cachedReportConfigurations property of the microsoft.graph.deviceManagementReports entity.
func (m *DeviceManagementReportsRequestBuilder) CachedReportConfigurationsById(id string)(*DeviceManagementReportsCachedReportConfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementCachedReportConfiguration%2Did"] = id
    }
    return NewDeviceManagementReportsCachedReportConfigurationsDeviceManagementCachedReportConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewDeviceManagementReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementReportsRequestBuilder) {
    m := &DeviceManagementReportsRequestBuilder{
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
// NewDeviceManagementReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewDeviceManagementReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reports for deviceManagement
func (m *DeviceManagementReportsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation reports singleton
func (m *DeviceManagementReportsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property reports in deviceManagement
func (m *DeviceManagementReportsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, requestConfiguration *DeviceManagementReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property reports for deviceManagement
func (m *DeviceManagementReportsRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementReportsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
func (m *DeviceManagementReportsRequestBuilder) ExportJobs()(*DeviceManagementReportsExportJobsRequestBuilder) {
    return NewDeviceManagementReportsExportJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportJobsById provides operations to manage the exportJobs property of the microsoft.graph.deviceManagementReports entity.
func (m *DeviceManagementReportsRequestBuilder) ExportJobsById(id string)(*DeviceManagementReportsExportJobsDeviceManagementExportJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementExportJob%2Did"] = id
    }
    return NewDeviceManagementReportsExportJobsDeviceManagementExportJobItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get reports singleton
func (m *DeviceManagementReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
func (m *DeviceManagementReportsRequestBuilder) GetActiveMalwareReport()(*DeviceManagementReportsGetActiveMalwareReportRequestBuilder) {
    return NewDeviceManagementReportsGetActiveMalwareReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetActiveMalwareSummaryReport provides operations to call the getActiveMalwareSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetActiveMalwareSummaryReport()(*DeviceManagementReportsGetActiveMalwareSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetActiveMalwareSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAllCertificatesReport provides operations to call the getAllCertificatesReport method.
func (m *DeviceManagementReportsRequestBuilder) GetAllCertificatesReport()(*DeviceManagementReportsGetAllCertificatesReportRequestBuilder) {
    return NewDeviceManagementReportsGetAllCertificatesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAppsInstallSummaryReport provides operations to call the getAppsInstallSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetAppsInstallSummaryReport()(*DeviceManagementReportsGetAppsInstallSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetAppsInstallSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAppStatusOverviewReport provides operations to call the getAppStatusOverviewReport method.
func (m *DeviceManagementReportsRequestBuilder) GetAppStatusOverviewReport()(*DeviceManagementReportsGetAppStatusOverviewReportRequestBuilder) {
    return NewDeviceManagementReportsGetAppStatusOverviewReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCachedReport provides operations to call the getCachedReport method.
func (m *DeviceManagementReportsRequestBuilder) GetCachedReport()(*DeviceManagementReportsGetCachedReportRequestBuilder) {
    return NewDeviceManagementReportsGetCachedReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCertificatesReport provides operations to call the getCertificatesReport method.
func (m *DeviceManagementReportsRequestBuilder) GetCertificatesReport()(*DeviceManagementReportsGetCertificatesReportRequestBuilder) {
    return NewDeviceManagementReportsGetCertificatesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePoliciesReportForDevice provides operations to call the getCompliancePoliciesReportForDevice method.
func (m *DeviceManagementReportsRequestBuilder) GetCompliancePoliciesReportForDevice()(*DeviceManagementReportsGetCompliancePoliciesReportForDeviceRequestBuilder) {
    return NewDeviceManagementReportsGetCompliancePoliciesReportForDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyDevicesReport provides operations to call the getCompliancePolicyDevicesReport method.
func (m *DeviceManagementReportsRequestBuilder) GetCompliancePolicyDevicesReport()(*DeviceManagementReportsGetCompliancePolicyDevicesReportRequestBuilder) {
    return NewDeviceManagementReportsGetCompliancePolicyDevicesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyDeviceSummaryReport provides operations to call the getCompliancePolicyDeviceSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetCompliancePolicyDeviceSummaryReport()(*DeviceManagementReportsGetCompliancePolicyDeviceSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetCompliancePolicyDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyNonComplianceReport provides operations to call the getCompliancePolicyNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetCompliancePolicyNonComplianceReport()(*DeviceManagementReportsGetCompliancePolicyNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetCompliancePolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyNonComplianceSummaryReport provides operations to call the getCompliancePolicyNonComplianceSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetCompliancePolicyNonComplianceSummaryReport()(*DeviceManagementReportsGetCompliancePolicyNonComplianceSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetCompliancePolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComplianceSettingDetailsReport provides operations to call the getComplianceSettingDetailsReport method.
func (m *DeviceManagementReportsRequestBuilder) GetComplianceSettingDetailsReport()(*DeviceManagementReportsGetComplianceSettingDetailsReportRequestBuilder) {
    return NewDeviceManagementReportsGetComplianceSettingDetailsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComplianceSettingNonComplianceReport provides operations to call the getComplianceSettingNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetComplianceSettingNonComplianceReport()(*DeviceManagementReportsGetComplianceSettingNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetComplianceSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComplianceSettingsReport provides operations to call the getComplianceSettingsReport method.
func (m *DeviceManagementReportsRequestBuilder) GetComplianceSettingsReport()(*DeviceManagementReportsGetComplianceSettingsReportRequestBuilder) {
    return NewDeviceManagementReportsGetComplianceSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigManagerDevicePolicyStatusReport provides operations to call the getConfigManagerDevicePolicyStatusReport method.
func (m *DeviceManagementReportsRequestBuilder) GetConfigManagerDevicePolicyStatusReport()(*DeviceManagementReportsGetConfigManagerDevicePolicyStatusReportRequestBuilder) {
    return NewDeviceManagementReportsGetConfigManagerDevicePolicyStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPoliciesReportForDevice provides operations to call the getConfigurationPoliciesReportForDevice method.
func (m *DeviceManagementReportsRequestBuilder) GetConfigurationPoliciesReportForDevice()(*DeviceManagementReportsGetConfigurationPoliciesReportForDeviceRequestBuilder) {
    return NewDeviceManagementReportsGetConfigurationPoliciesReportForDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyDevicesReport provides operations to call the getConfigurationPolicyDevicesReport method.
func (m *DeviceManagementReportsRequestBuilder) GetConfigurationPolicyDevicesReport()(*DeviceManagementReportsGetConfigurationPolicyDevicesReportRequestBuilder) {
    return NewDeviceManagementReportsGetConfigurationPolicyDevicesReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyDeviceSummaryReport provides operations to call the getConfigurationPolicyDeviceSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetConfigurationPolicyDeviceSummaryReport()(*DeviceManagementReportsGetConfigurationPolicyDeviceSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetConfigurationPolicyDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyNonComplianceReport provides operations to call the getConfigurationPolicyNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetConfigurationPolicyNonComplianceReport()(*DeviceManagementReportsGetConfigurationPolicyNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetConfigurationPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyNonComplianceSummaryReport provides operations to call the getConfigurationPolicyNonComplianceSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetConfigurationPolicyNonComplianceSummaryReport()(*DeviceManagementReportsGetConfigurationPolicyNonComplianceSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetConfigurationPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicySettingsDeviceSummaryReport provides operations to call the getConfigurationPolicySettingsDeviceSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetConfigurationPolicySettingsDeviceSummaryReport()(*DeviceManagementReportsGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetConfigurationPolicySettingsDeviceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationSettingDetailsReport provides operations to call the getConfigurationSettingDetailsReport method.
func (m *DeviceManagementReportsRequestBuilder) GetConfigurationSettingDetailsReport()(*DeviceManagementReportsGetConfigurationSettingDetailsReportRequestBuilder) {
    return NewDeviceManagementReportsGetConfigurationSettingDetailsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationSettingNonComplianceReport provides operations to call the getConfigurationSettingNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetConfigurationSettingNonComplianceReport()(*DeviceManagementReportsGetConfigurationSettingNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetConfigurationSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationSettingsReport provides operations to call the getConfigurationSettingsReport method.
func (m *DeviceManagementReportsRequestBuilder) GetConfigurationSettingsReport()(*DeviceManagementReportsGetConfigurationSettingsReportRequestBuilder) {
    return NewDeviceManagementReportsGetConfigurationSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceConfigurationPolicySettingsSummaryReport provides operations to call the getDeviceConfigurationPolicySettingsSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetDeviceConfigurationPolicySettingsSummaryReport()(*DeviceManagementReportsGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceConfigurationPolicyStatusSummary provides operations to call the getDeviceConfigurationPolicyStatusSummary method.
func (m *DeviceManagementReportsRequestBuilder) GetDeviceConfigurationPolicyStatusSummary()(*DeviceManagementReportsGetDeviceConfigurationPolicyStatusSummaryRequestBuilder) {
    return NewDeviceManagementReportsGetDeviceConfigurationPolicyStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceInstallStatusReport provides operations to call the getDeviceInstallStatusReport method.
func (m *DeviceManagementReportsRequestBuilder) GetDeviceInstallStatusReport()(*DeviceManagementReportsGetDeviceInstallStatusReportRequestBuilder) {
    return NewDeviceManagementReportsGetDeviceInstallStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceManagementIntentPerSettingContributingProfiles provides operations to call the getDeviceManagementIntentPerSettingContributingProfiles method.
func (m *DeviceManagementReportsRequestBuilder) GetDeviceManagementIntentPerSettingContributingProfiles()(*DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    return NewDeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceManagementIntentSettingsReport provides operations to call the getDeviceManagementIntentSettingsReport method.
func (m *DeviceManagementReportsRequestBuilder) GetDeviceManagementIntentSettingsReport()(*DeviceManagementReportsGetDeviceManagementIntentSettingsReportRequestBuilder) {
    return NewDeviceManagementReportsGetDeviceManagementIntentSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceNonComplianceReport provides operations to call the getDeviceNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetDeviceNonComplianceReport()(*DeviceManagementReportsGetDeviceNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetDeviceNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEncryptionReportForDevices provides operations to call the getEncryptionReportForDevices method.
func (m *DeviceManagementReportsRequestBuilder) GetEncryptionReportForDevices()(*DeviceManagementReportsGetEncryptionReportForDevicesRequestBuilder) {
    return NewDeviceManagementReportsGetEncryptionReportForDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEnrollmentConfigurationPoliciesByDevice provides operations to call the getEnrollmentConfigurationPoliciesByDevice method.
func (m *DeviceManagementReportsRequestBuilder) GetEnrollmentConfigurationPoliciesByDevice()(*DeviceManagementReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) {
    return NewDeviceManagementReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFailedMobileAppsReport provides operations to call the getFailedMobileAppsReport method.
func (m *DeviceManagementReportsRequestBuilder) GetFailedMobileAppsReport()(*DeviceManagementReportsGetFailedMobileAppsReportRequestBuilder) {
    return NewDeviceManagementReportsGetFailedMobileAppsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFailedMobileAppsSummaryReport provides operations to call the getFailedMobileAppsSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetFailedMobileAppsSummaryReport()(*DeviceManagementReportsGetFailedMobileAppsSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetFailedMobileAppsSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetGroupPolicySettingsDeviceSettingsReport provides operations to call the getGroupPolicySettingsDeviceSettingsReport method.
func (m *DeviceManagementReportsRequestBuilder) GetGroupPolicySettingsDeviceSettingsReport()(*DeviceManagementReportsGetGroupPolicySettingsDeviceSettingsReportRequestBuilder) {
    return NewDeviceManagementReportsGetGroupPolicySettingsDeviceSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetHistoricalReport provides operations to call the getHistoricalReport method.
func (m *DeviceManagementReportsRequestBuilder) GetHistoricalReport()(*DeviceManagementReportsGetHistoricalReportRequestBuilder) {
    return NewDeviceManagementReportsGetHistoricalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMalwareSummaryReport provides operations to call the getMalwareSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetMalwareSummaryReport()(*DeviceManagementReportsGetMalwareSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetMalwareSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMobileApplicationManagementAppConfigurationReport provides operations to call the getMobileApplicationManagementAppConfigurationReport method.
func (m *DeviceManagementReportsRequestBuilder) GetMobileApplicationManagementAppConfigurationReport()(*DeviceManagementReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilder) {
    return NewDeviceManagementReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMobileApplicationManagementAppRegistrationSummaryReport provides operations to call the getMobileApplicationManagementAppRegistrationSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetMobileApplicationManagementAppRegistrationSummaryReport()(*DeviceManagementReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNoncompliantDevicesAndSettingsReport provides operations to call the getNoncompliantDevicesAndSettingsReport method.
func (m *DeviceManagementReportsRequestBuilder) GetNoncompliantDevicesAndSettingsReport()(*DeviceManagementReportsGetNoncompliantDevicesAndSettingsReportRequestBuilder) {
    return NewDeviceManagementReportsGetNoncompliantDevicesAndSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPolicyNonComplianceMetadata provides operations to call the getPolicyNonComplianceMetadata method.
func (m *DeviceManagementReportsRequestBuilder) GetPolicyNonComplianceMetadata()(*DeviceManagementReportsGetPolicyNonComplianceMetadataRequestBuilder) {
    return NewDeviceManagementReportsGetPolicyNonComplianceMetadataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPolicyNonComplianceReport provides operations to call the getPolicyNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetPolicyNonComplianceReport()(*DeviceManagementReportsGetPolicyNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPolicyNonComplianceSummaryReport provides operations to call the getPolicyNonComplianceSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetPolicyNonComplianceSummaryReport()(*DeviceManagementReportsGetPolicyNonComplianceSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetQuietTimePolicyUsersReport provides operations to call the getQuietTimePolicyUsersReport method.
func (m *DeviceManagementReportsRequestBuilder) GetQuietTimePolicyUsersReport()(*DeviceManagementReportsGetQuietTimePolicyUsersReportRequestBuilder) {
    return NewDeviceManagementReportsGetQuietTimePolicyUsersReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetQuietTimePolicyUserSummaryReport provides operations to call the getQuietTimePolicyUserSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetQuietTimePolicyUserSummaryReport()(*DeviceManagementReportsGetQuietTimePolicyUserSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetQuietTimePolicyUserSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetRelatedAppsStatusReport provides operations to call the getRelatedAppsStatusReport method.
func (m *DeviceManagementReportsRequestBuilder) GetRelatedAppsStatusReport()(*DeviceManagementReportsGetRelatedAppsStatusReportRequestBuilder) {
    return NewDeviceManagementReportsGetRelatedAppsStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetRemoteAssistanceSessionsReport provides operations to call the getRemoteAssistanceSessionsReport method.
func (m *DeviceManagementReportsRequestBuilder) GetRemoteAssistanceSessionsReport()(*DeviceManagementReportsGetRemoteAssistanceSessionsReportRequestBuilder) {
    return NewDeviceManagementReportsGetRemoteAssistanceSessionsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetReportFilters provides operations to call the getReportFilters method.
func (m *DeviceManagementReportsRequestBuilder) GetReportFilters()(*DeviceManagementReportsGetReportFiltersRequestBuilder) {
    return NewDeviceManagementReportsGetReportFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetSettingNonComplianceReport provides operations to call the getSettingNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetSettingNonComplianceReport()(*DeviceManagementReportsGetSettingNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUnhealthyDefenderAgentsReport provides operations to call the getUnhealthyDefenderAgentsReport method.
func (m *DeviceManagementReportsRequestBuilder) GetUnhealthyDefenderAgentsReport()(*DeviceManagementReportsGetUnhealthyDefenderAgentsReportRequestBuilder) {
    return NewDeviceManagementReportsGetUnhealthyDefenderAgentsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUnhealthyFirewallReport provides operations to call the getUnhealthyFirewallReport method.
func (m *DeviceManagementReportsRequestBuilder) GetUnhealthyFirewallReport()(*DeviceManagementReportsGetUnhealthyFirewallReportRequestBuilder) {
    return NewDeviceManagementReportsGetUnhealthyFirewallReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUnhealthyFirewallSummaryReport provides operations to call the getUnhealthyFirewallSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetUnhealthyFirewallSummaryReport()(*DeviceManagementReportsGetUnhealthyFirewallSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetUnhealthyFirewallSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserInstallStatusReport provides operations to call the getUserInstallStatusReport method.
func (m *DeviceManagementReportsRequestBuilder) GetUserInstallStatusReport()(*DeviceManagementReportsGetUserInstallStatusReportRequestBuilder) {
    return NewDeviceManagementReportsGetUserInstallStatusReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReport provides operations to call the getWindowsQualityUpdateAlertsPerPolicyPerDeviceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetWindowsQualityUpdateAlertsPerPolicyPerDeviceReport()(*DeviceManagementReportsGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    return NewDeviceManagementReportsGetWindowsQualityUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWindowsQualityUpdateAlertSummaryReport provides operations to call the getWindowsQualityUpdateAlertSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetWindowsQualityUpdateAlertSummaryReport()(*DeviceManagementReportsGetWindowsQualityUpdateAlertSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetWindowsQualityUpdateAlertSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWindowsUpdateAlertsPerPolicyPerDeviceReport provides operations to call the getWindowsUpdateAlertsPerPolicyPerDeviceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetWindowsUpdateAlertsPerPolicyPerDeviceReport()(*DeviceManagementReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilder) {
    return NewDeviceManagementReportsGetWindowsUpdateAlertsPerPolicyPerDeviceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWindowsUpdateAlertSummaryReport provides operations to call the getWindowsUpdateAlertSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetWindowsUpdateAlertSummaryReport()(*DeviceManagementReportsGetWindowsUpdateAlertSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetWindowsUpdateAlertSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetZebraFotaDeploymentReport provides operations to call the getZebraFotaDeploymentReport method.
func (m *DeviceManagementReportsRequestBuilder) GetZebraFotaDeploymentReport()(*DeviceManagementReportsGetZebraFotaDeploymentReportRequestBuilder) {
    return NewDeviceManagementReportsGetZebraFotaDeploymentReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property reports in deviceManagement
func (m *DeviceManagementReportsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, requestConfiguration *DeviceManagementReportsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReportsable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
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
