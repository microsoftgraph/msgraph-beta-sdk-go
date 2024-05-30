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
// returns a *AndroidmanagedappprotectionsAndroidManagedAppProtectionsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) AndroidManagedAppProtections()(*AndroidmanagedappprotectionsAndroidManagedAppProtectionsRequestBuilder) {
    return NewAndroidmanagedappprotectionsAndroidManagedAppProtectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
// returns a *DefaultmanagedappprotectionsDefaultManagedAppProtectionsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) DefaultManagedAppProtections()(*DefaultmanagedappprotectionsDefaultManagedAppProtectionsRequestBuilder) {
    return NewDefaultmanagedappprotectionsDefaultManagedAppProtectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceAppManagementTasks provides operations to manage the deviceAppManagementTasks property of the microsoft.graph.deviceAppManagement entity.
// returns a *DeviceappmanagementtasksDeviceAppManagementTasksRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) DeviceAppManagementTasks()(*DeviceappmanagementtasksDeviceAppManagementTasksRequestBuilder) {
    return NewDeviceappmanagementtasksDeviceAppManagementTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnterpriseCodeSigningCertificates provides operations to manage the enterpriseCodeSigningCertificates property of the microsoft.graph.deviceAppManagement entity.
// returns a *EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) EnterpriseCodeSigningCertificates()(*EnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilder) {
    return NewEnterprisecodesigningcertificatesEnterpriseCodeSigningCertificatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
// returns a *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) IosLobAppProvisioningConfigurations()(*IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationsRequestBuilder) {
    return NewIoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IosManagedAppProtections provides operations to manage the iosManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
// returns a *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) IosManagedAppProtections()(*IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder) {
    return NewIosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedAppPolicies provides operations to manage the managedAppPolicies property of the microsoft.graph.deviceAppManagement entity.
// returns a *ManagedapppoliciesManagedAppPoliciesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) ManagedAppPolicies()(*ManagedapppoliciesManagedAppPoliciesRequestBuilder) {
    return NewManagedapppoliciesManagedAppPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedAppRegistrations provides operations to manage the managedAppRegistrations property of the microsoft.graph.deviceAppManagement entity.
// returns a *ManagedappregistrationsManagedAppRegistrationsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) ManagedAppRegistrations()(*ManagedappregistrationsManagedAppRegistrationsRequestBuilder) {
    return NewManagedappregistrationsManagedAppRegistrationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedAppStatuses provides operations to manage the managedAppStatuses property of the microsoft.graph.deviceAppManagement entity.
// returns a *ManagedappstatusesManagedAppStatusesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) ManagedAppStatuses()(*ManagedappstatusesManagedAppStatusesRequestBuilder) {
    return NewManagedappstatusesManagedAppStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedEBookCategories provides operations to manage the managedEBookCategories property of the microsoft.graph.deviceAppManagement entity.
// returns a *ManagedebookcategoriesManagedEBookCategoriesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) ManagedEBookCategories()(*ManagedebookcategoriesManagedEBookCategoriesRequestBuilder) {
    return NewManagedebookcategoriesManagedEBookCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedEBooks provides operations to manage the managedEBooks property of the microsoft.graph.deviceAppManagement entity.
// returns a *ManagedebooksManagedEBooksRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) ManagedEBooks()(*ManagedebooksManagedEBooksRequestBuilder) {
    return NewManagedebooksManagedEBooksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MdmWindowsInformationProtectionPolicies provides operations to manage the mdmWindowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
// returns a *MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPoliciesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) MdmWindowsInformationProtectionPolicies()(*MdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPoliciesRequestBuilder) {
    return NewMdmwindowsinformationprotectionpoliciesMdmWindowsInformationProtectionPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileAppCatalogPackages provides operations to manage the mobileAppCatalogPackages property of the microsoft.graph.deviceAppManagement entity.
// returns a *MobileappcatalogpackagesMobileAppCatalogPackagesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) MobileAppCatalogPackages()(*MobileappcatalogpackagesMobileAppCatalogPackagesRequestBuilder) {
    return NewMobileappcatalogpackagesMobileAppCatalogPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileAppCategories provides operations to manage the mobileAppCategories property of the microsoft.graph.deviceAppManagement entity.
// returns a *MobileappcategoriesMobileAppCategoriesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) MobileAppCategories()(*MobileappcategoriesMobileAppCategoriesRequestBuilder) {
    return NewMobileappcategoriesMobileAppCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileAppConfigurations provides operations to manage the mobileAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
// returns a *MobileappconfigurationsMobileAppConfigurationsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) MobileAppConfigurations()(*MobileappconfigurationsMobileAppConfigurationsRequestBuilder) {
    return NewMobileappconfigurationsMobileAppConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileApps provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
// returns a *MobileappsMobileAppsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) MobileApps()(*MobileappsMobileAppsRequestBuilder) {
    return NewMobileappsMobileAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
// returns a *PolicysetsPolicySetsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) PolicySets()(*PolicysetsPolicySetsRequestBuilder) {
    return NewPolicysetsPolicySetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SymantecCodeSigningCertificate provides operations to manage the symantecCodeSigningCertificate property of the microsoft.graph.deviceAppManagement entity.
// returns a *SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) SymantecCodeSigningCertificate()(*SymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilder) {
    return NewSymanteccodesigningcertificateSymantecCodeSigningCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncMicrosoftStoreForBusinessApps provides operations to call the syncMicrosoftStoreForBusinessApps method.
// returns a *SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) SyncMicrosoftStoreForBusinessApps()(*SyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    return NewSyncmicrosoftstoreforbusinessappsSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TargetedManagedAppConfigurations provides operations to manage the targetedManagedAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
// returns a *TargetedmanagedappconfigurationsTargetedManagedAppConfigurationsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) TargetedManagedAppConfigurations()(*TargetedmanagedappconfigurationsTargetedManagedAppConfigurationsRequestBuilder) {
    return NewTargetedmanagedappconfigurationsTargetedManagedAppConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
// VppTokens provides operations to manage the vppTokens property of the microsoft.graph.deviceAppManagement entity.
// returns a *VpptokensVppTokensRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) VppTokens()(*VpptokensVppTokensRequestBuilder) {
    return NewVpptokensVppTokensRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WdacSupplementalPolicies provides operations to manage the wdacSupplementalPolicies property of the microsoft.graph.deviceAppManagement entity.
// returns a *WdacsupplementalpoliciesWdacSupplementalPoliciesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WdacSupplementalPolicies()(*WdacsupplementalpoliciesWdacSupplementalPoliciesRequestBuilder) {
    return NewWdacsupplementalpoliciesWdacSupplementalPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsInformationProtectionDeviceRegistrations provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.deviceAppManagement entity.
// returns a *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionDeviceRegistrations()(*WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    return NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsInformationProtectionPolicies provides operations to manage the windowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
// returns a *WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionPolicies()(*WindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilder) {
    return NewWindowsinformationprotectionpoliciesWindowsInformationProtectionPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsInformationProtectionWipeActions provides operations to manage the windowsInformationProtectionWipeActions property of the microsoft.graph.deviceAppManagement entity.
// returns a *WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionWipeActions()(*WindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionsRequestBuilder) {
    return NewWindowsinformationprotectionwipeactionsWindowsInformationProtectionWipeActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsManagedAppProtections provides operations to manage the windowsManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
// returns a *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WindowsManagedAppProtections()(*WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder) {
    return NewWindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsManagementApp provides operations to manage the windowsManagementApp property of the microsoft.graph.deviceAppManagement entity.
// returns a *WindowsmanagementappWindowsManagementAppRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WindowsManagementApp()(*WindowsmanagementappWindowsManagementAppRequestBuilder) {
    return NewWindowsmanagementappWindowsManagementAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceAppManagementRequestBuilder when successful
func (m *DeviceAppManagementRequestBuilder) WithUrl(rawUrl string)(*DeviceAppManagementRequestBuilder) {
    return NewDeviceAppManagementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
