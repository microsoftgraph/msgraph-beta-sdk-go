package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementRequestBuilder provides operations to manage the deviceAppManagement singleton.
type DeviceAppManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
func (m *DeviceAppManagementRequestBuilder) AndroidManagedAppProtections()(*AndroidManagedAppProtectionsRequestBuilder) {
    return NewAndroidManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AndroidManagedAppProtectionsById provides operations to manage the androidManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) AndroidManagedAppProtectionsById(id string)(*AndroidManagedAppProtectionsAndroidManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewAndroidManagedAppProtectionsAndroidManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// NewDeviceAppManagementRequestBuilderInternal instantiates a new DeviceAppManagementRequestBuilder and sets the default values.
func NewDeviceAppManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementRequestBuilder) {
    m := &DeviceAppManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeviceAppManagementRequestBuilder instantiates a new DeviceAppManagementRequestBuilder and sets the default values.
func NewDeviceAppManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// DefaultManagedAppProtections provides operations to manage the defaultManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) DefaultManagedAppProtections()(*DefaultManagedAppProtectionsRequestBuilder) {
    return NewDefaultManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DefaultManagedAppProtectionsById provides operations to manage the defaultManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) DefaultManagedAppProtectionsById(id string)(*DefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceAppManagementTasks provides operations to manage the deviceAppManagementTasks property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) DeviceAppManagementTasks()(*DeviceAppManagementTasksRequestBuilder) {
    return NewDeviceAppManagementTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceAppManagementTasksById provides operations to manage the deviceAppManagementTasks property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) DeviceAppManagementTasksById(id string)(*DeviceAppManagementTasksDeviceAppManagementTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceAppManagementTasksDeviceAppManagementTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// EnterpriseCodeSigningCertificates provides operations to manage the enterpriseCodeSigningCertificates property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) EnterpriseCodeSigningCertificates()(*EnterpriseCodeSigningCertificatesRequestBuilder) {
    return NewEnterpriseCodeSigningCertificatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EnterpriseCodeSigningCertificatesById provides operations to manage the enterpriseCodeSigningCertificates property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) EnterpriseCodeSigningCertificatesById(id string)(*EnterpriseCodeSigningCertificatesEnterpriseCodeSigningCertificateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEnterpriseCodeSigningCertificatesEnterpriseCodeSigningCertificateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Get get deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAppManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable), nil
}
// IosLobAppProvisioningConfigurations provides operations to manage the iosLobAppProvisioningConfigurations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) IosLobAppProvisioningConfigurations()(*IosLobAppProvisioningConfigurationsRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IosLobAppProvisioningConfigurationsById provides operations to manage the iosLobAppProvisioningConfigurations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) IosLobAppProvisioningConfigurationsById(id string)(*IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// IosManagedAppProtections provides operations to manage the iosManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) IosManagedAppProtections()(*IosManagedAppProtectionsRequestBuilder) {
    return NewIosManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IosManagedAppProtectionsById provides operations to manage the iosManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) IosManagedAppProtectionsById(id string)(*IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewIosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ManagedAppPolicies provides operations to manage the managedAppPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedAppPolicies()(*ManagedAppPoliciesRequestBuilder) {
    return NewManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedAppPoliciesById provides operations to manage the managedAppPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedAppPoliciesById(id string)(*ManagedAppPoliciesManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewManagedAppPoliciesManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ManagedAppRegistrations provides operations to manage the managedAppRegistrations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedAppRegistrations()(*ManagedAppRegistrationsRequestBuilder) {
    return NewManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedAppRegistrationsById provides operations to manage the managedAppRegistrations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedAppRegistrationsById(id string)(*ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ManagedAppStatuses provides operations to manage the managedAppStatuses property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedAppStatuses()(*ManagedAppStatusesRequestBuilder) {
    return NewManagedAppStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedAppStatusesById provides operations to manage the managedAppStatuses property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedAppStatusesById(id string)(*ManagedAppStatusesManagedAppStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewManagedAppStatusesManagedAppStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ManagedEBookCategories provides operations to manage the managedEBookCategories property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedEBookCategories()(*ManagedEBookCategoriesRequestBuilder) {
    return NewManagedEBookCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedEBookCategoriesById provides operations to manage the managedEBookCategories property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedEBookCategoriesById(id string)(*ManagedEBookCategoriesManagedEBookCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewManagedEBookCategoriesManagedEBookCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ManagedEBooks provides operations to manage the managedEBooks property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedEBooks()(*ManagedEBooksRequestBuilder) {
    return NewManagedEBooksRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedEBooksById provides operations to manage the managedEBooks property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedEBooksById(id string)(*ManagedEBooksManagedEBookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewManagedEBooksManagedEBookItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MdmWindowsInformationProtectionPolicies provides operations to manage the mdmWindowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MdmWindowsInformationProtectionPolicies()(*MdmWindowsInformationProtectionPoliciesRequestBuilder) {
    return NewMdmWindowsInformationProtectionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MdmWindowsInformationProtectionPoliciesById provides operations to manage the mdmWindowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MdmWindowsInformationProtectionPoliciesById(id string)(*MdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MicrosoftGraphSyncMicrosoftStoreForBusinessApps provides operations to call the syncMicrosoftStoreForBusinessApps method.
func (m *DeviceAppManagementRequestBuilder) MicrosoftGraphSyncMicrosoftStoreForBusinessApps()(*MicrosoftGraphSyncMicrosoftStoreForBusinessAppsSyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    return NewMicrosoftGraphSyncMicrosoftStoreForBusinessAppsSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MobileAppCategories provides operations to manage the mobileAppCategories property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MobileAppCategories()(*MobileAppCategoriesRequestBuilder) {
    return NewMobileAppCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MobileAppCategoriesById provides operations to manage the mobileAppCategories property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MobileAppCategoriesById(id string)(*MobileAppCategoriesMobileAppCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMobileAppCategoriesMobileAppCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MobileAppConfigurations provides operations to manage the mobileAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MobileAppConfigurations()(*MobileAppConfigurationsRequestBuilder) {
    return NewMobileAppConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MobileAppConfigurationsById provides operations to manage the mobileAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MobileAppConfigurationsById(id string)(*MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MobileApps provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MobileApps()(*MobileAppsRequestBuilder) {
    return NewMobileAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MobileAppsById provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MobileAppsById(id string)(*MobileAppsMobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMobileAppsMobileAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Patch update deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable, requestConfiguration *DeviceAppManagementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAppManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable), nil
}
// PolicySets provides operations to manage the policySets property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) PolicySets()(*PolicySetsRequestBuilder) {
    return NewPolicySetsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PolicySetsById provides operations to manage the policySets property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) PolicySetsById(id string)(*PolicySetsPolicySetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewPolicySetsPolicySetItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// SymantecCodeSigningCertificate provides operations to manage the symantecCodeSigningCertificate property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) SymantecCodeSigningCertificate()(*SymantecCodeSigningCertificateRequestBuilder) {
    return NewSymantecCodeSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TargetedManagedAppConfigurations provides operations to manage the targetedManagedAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) TargetedManagedAppConfigurations()(*TargetedManagedAppConfigurationsRequestBuilder) {
    return NewTargetedManagedAppConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TargetedManagedAppConfigurationsById provides operations to manage the targetedManagedAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) TargetedManagedAppConfigurationsById(id string)(*TargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ToGetRequestInformation get deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable, requestConfiguration *DeviceAppManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// VppTokens provides operations to manage the vppTokens property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) VppTokens()(*VppTokensRequestBuilder) {
    return NewVppTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// VppTokensById provides operations to manage the vppTokens property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) VppTokensById(id string)(*VppTokensVppTokenItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewVppTokensVppTokenItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WdacSupplementalPolicies provides operations to manage the wdacSupplementalPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) WdacSupplementalPolicies()(*WdacSupplementalPoliciesRequestBuilder) {
    return NewWdacSupplementalPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WdacSupplementalPoliciesById provides operations to manage the wdacSupplementalPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) WdacSupplementalPoliciesById(id string)(*WdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWdacSupplementalPoliciesWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WindowsInformationProtectionDeviceRegistrations provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionDeviceRegistrations()(*WindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    return NewWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsInformationProtectionDeviceRegistrationsById provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionDeviceRegistrationsById(id string)(*WindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WindowsInformationProtectionPolicies provides operations to manage the windowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionPolicies()(*WindowsInformationProtectionPoliciesRequestBuilder) {
    return NewWindowsInformationProtectionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsInformationProtectionPoliciesById provides operations to manage the windowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionPoliciesById(id string)(*WindowsInformationProtectionPoliciesWindowsInformationProtectionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWindowsInformationProtectionPoliciesWindowsInformationProtectionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WindowsInformationProtectionWipeActions provides operations to manage the windowsInformationProtectionWipeActions property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionWipeActions()(*WindowsInformationProtectionWipeActionsRequestBuilder) {
    return NewWindowsInformationProtectionWipeActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsInformationProtectionWipeActionsById provides operations to manage the windowsInformationProtectionWipeActions property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionWipeActionsById(id string)(*WindowsInformationProtectionWipeActionsWindowsInformationProtectionWipeActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWindowsInformationProtectionWipeActionsWindowsInformationProtectionWipeActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WindowsManagedAppProtections provides operations to manage the windowsManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) WindowsManagedAppProtections()(*WindowsManagedAppProtectionsRequestBuilder) {
    return NewWindowsManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsManagedAppProtectionsById provides operations to manage the windowsManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) WindowsManagedAppProtectionsById(id string)(*WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WindowsManagementApp provides operations to manage the windowsManagementApp property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) WindowsManagementApp()(*WindowsManagementAppRequestBuilder) {
    return NewWindowsManagementAppRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
