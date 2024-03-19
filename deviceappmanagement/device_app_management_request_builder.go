package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementRequestBuilder provides operations to manage the deviceAppManagement singleton.
type DeviceAppManagementRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceAppManagementRequestBuilderGetQueryParameters get deviceAppManagement
type DeviceAppManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementRequestBuilderGetQueryParameters
}
// DeviceAppManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AndroidManagedAppProtections provides operations to manage the androidManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
// returns a *AndroidManagedAppProtectionsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) AndroidManagedAppProtections()(*AndroidManagedAppProtectionsRequestBuilder) {
    return NewAndroidManagedAppProtectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceAppManagementRequestBuilderInternal instantiates a new DeviceAppManagementRequestBuilder and sets the default values.
func NewDeviceAppManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementRequestBuilder) {
    m := &DeviceAppManagementRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceAppManagementRequestBuilder instantiates a new DeviceAppManagementRequestBuilder and sets the default values.
func NewDeviceAppManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// DefaultManagedAppProtections provides operations to manage the defaultManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
// returns a *DefaultManagedAppProtectionsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) DefaultManagedAppProtections()(*DefaultManagedAppProtectionsRequestBuilder) {
    return NewDefaultManagedAppProtectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceAppManagementTasks provides operations to manage the deviceAppManagementTasks property of the microsoft.graph.deviceAppManagement entity.
// returns a *DeviceAppManagementTasksRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) DeviceAppManagementTasks()(*DeviceAppManagementTasksRequestBuilder) {
    return NewDeviceAppManagementTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnterpriseCodeSigningCertificates provides operations to manage the enterpriseCodeSigningCertificates property of the microsoft.graph.deviceAppManagement entity.
// returns a *EnterpriseCodeSigningCertificatesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) EnterpriseCodeSigningCertificates()(*EnterpriseCodeSigningCertificatesRequestBuilder) {
    return NewEnterpriseCodeSigningCertificatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get deviceAppManagement
// returns a DeviceAppManagementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceAppManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAppManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable), nil
}
// IosLobAppProvisioningConfigurations provides operations to manage the iosLobAppProvisioningConfigurations property of the microsoft.graph.deviceAppManagement entity.
// returns a *IosLobAppProvisioningConfigurationsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) IosLobAppProvisioningConfigurations()(*IosLobAppProvisioningConfigurationsRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IosManagedAppProtections provides operations to manage the iosManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
// returns a *IosManagedAppProtectionsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) IosManagedAppProtections()(*IosManagedAppProtectionsRequestBuilder) {
    return NewIosManagedAppProtectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedAppPolicies provides operations to manage the managedAppPolicies property of the microsoft.graph.deviceAppManagement entity.
// returns a *ManagedAppPoliciesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) ManagedAppPolicies()(*ManagedAppPoliciesRequestBuilder) {
    return NewManagedAppPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedAppRegistrations provides operations to manage the managedAppRegistrations property of the microsoft.graph.deviceAppManagement entity.
// returns a *ManagedAppRegistrationsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) ManagedAppRegistrations()(*ManagedAppRegistrationsRequestBuilder) {
    return NewManagedAppRegistrationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedAppStatuses provides operations to manage the managedAppStatuses property of the microsoft.graph.deviceAppManagement entity.
// returns a *ManagedAppStatusesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) ManagedAppStatuses()(*ManagedAppStatusesRequestBuilder) {
    return NewManagedAppStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedEBookCategories provides operations to manage the managedEBookCategories property of the microsoft.graph.deviceAppManagement entity.
// returns a *ManagedEBookCategoriesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) ManagedEBookCategories()(*ManagedEBookCategoriesRequestBuilder) {
    return NewManagedEBookCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedEBooks provides operations to manage the managedEBooks property of the microsoft.graph.deviceAppManagement entity.
// returns a *ManagedEBooksRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) ManagedEBooks()(*ManagedEBooksRequestBuilder) {
    return NewManagedEBooksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MdmWindowsInformationProtectionPolicies provides operations to manage the mdmWindowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
// returns a *MdmWindowsInformationProtectionPoliciesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) MdmWindowsInformationProtectionPolicies()(*MdmWindowsInformationProtectionPoliciesRequestBuilder) {
    return NewMdmWindowsInformationProtectionPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileAppCatalogPackages provides operations to manage the mobileAppCatalogPackages property of the microsoft.graph.deviceAppManagement entity.
// returns a *MobileAppCatalogPackagesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) MobileAppCatalogPackages()(*MobileAppCatalogPackagesRequestBuilder) {
    return NewMobileAppCatalogPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileAppCategories provides operations to manage the mobileAppCategories property of the microsoft.graph.deviceAppManagement entity.
// returns a *MobileAppCategoriesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) MobileAppCategories()(*MobileAppCategoriesRequestBuilder) {
    return NewMobileAppCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileAppConfigurations provides operations to manage the mobileAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
// returns a *MobileAppConfigurationsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) MobileAppConfigurations()(*MobileAppConfigurationsRequestBuilder) {
    return NewMobileAppConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileApps provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
// returns a *MobileAppsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) MobileApps()(*MobileAppsRequestBuilder) {
    return NewMobileAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update deviceAppManagement
// returns a DeviceAppManagementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceAppManagementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable, requestConfiguration *DeviceAppManagementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAppManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable), nil
}
// PolicySets provides operations to manage the policySets property of the microsoft.graph.deviceAppManagement entity.
// returns a *PolicySetsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) PolicySets()(*PolicySetsRequestBuilder) {
    return NewPolicySetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SymantecCodeSigningCertificate provides operations to manage the symantecCodeSigningCertificate property of the microsoft.graph.deviceAppManagement entity.
// returns a *SymantecCodeSigningCertificateRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) SymantecCodeSigningCertificate()(*SymantecCodeSigningCertificateRequestBuilder) {
    return NewSymantecCodeSigningCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncMicrosoftStoreForBusinessApps provides operations to call the syncMicrosoftStoreForBusinessApps method.
// returns a *SyncMicrosoftStoreForBusinessAppsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) SyncMicrosoftStoreForBusinessApps()(*SyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    return NewSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TargetedManagedAppConfigurations provides operations to manage the targetedManagedAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
// returns a *TargetedManagedAppConfigurationsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) TargetedManagedAppConfigurations()(*TargetedManagedAppConfigurationsRequestBuilder) {
    return NewTargetedManagedAppConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get deviceAppManagement
// returns a *RequestInformation when successful
func (m *DeviceAppManagementRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update deviceAppManagement
// returns a *RequestInformation when successful
func (m *DeviceAppManagementRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable, requestConfiguration *DeviceAppManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceAppManagement", m.BaseRequestBuilder.PathParameters)
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
// VppTokens provides operations to manage the vppTokens property of the microsoft.graph.deviceAppManagement entity.
// returns a *VppTokensRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) VppTokens()(*VppTokensRequestBuilder) {
    return NewVppTokensRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WdacSupplementalPolicies provides operations to manage the wdacSupplementalPolicies property of the microsoft.graph.deviceAppManagement entity.
// returns a *WdacSupplementalPoliciesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WdacSupplementalPolicies()(*WdacSupplementalPoliciesRequestBuilder) {
    return NewWdacSupplementalPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsInformationProtectionDeviceRegistrations provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.deviceAppManagement entity.
// returns a *WindowsInformationProtectionDeviceRegistrationsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionDeviceRegistrations()(*WindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    return NewWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsInformationProtectionPolicies provides operations to manage the windowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
// returns a *WindowsInformationProtectionPoliciesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionPolicies()(*WindowsInformationProtectionPoliciesRequestBuilder) {
    return NewWindowsInformationProtectionPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsInformationProtectionWipeActions provides operations to manage the windowsInformationProtectionWipeActions property of the microsoft.graph.deviceAppManagement entity.
// returns a *WindowsInformationProtectionWipeActionsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionWipeActions()(*WindowsInformationProtectionWipeActionsRequestBuilder) {
    return NewWindowsInformationProtectionWipeActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsManagedAppProtections provides operations to manage the windowsManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
// returns a *WindowsManagedAppProtectionsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WindowsManagedAppProtections()(*WindowsManagedAppProtectionsRequestBuilder) {
    return NewWindowsManagedAppProtectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsManagementApp provides operations to manage the windowsManagementApp property of the microsoft.graph.deviceAppManagement entity.
// returns a *WindowsManagementAppRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WindowsManagementApp()(*WindowsManagementAppRequestBuilder) {
    return NewWindowsManagementAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceAppManagementRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WithUrl(rawUrl string)(*DeviceAppManagementRequestBuilder) {
    return NewDeviceAppManagementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
