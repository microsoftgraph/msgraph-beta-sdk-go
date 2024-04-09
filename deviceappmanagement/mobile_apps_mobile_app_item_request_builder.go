package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsMobileAppItemRequestBuilder provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
type MobileAppsMobileAppItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileAppsMobileAppItemRequestBuilderGetQueryParameters the mobile apps.
type MobileAppsMobileAppItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppsMobileAppItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsMobileAppItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsMobileAppItemRequestBuilderGetQueryParameters
}
// MobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *MobileAppsItemAssignRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) Assign()(*MobileAppsItemAssignRequestBuilder) {
    return NewMobileAppsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileAppsItemAssignmentsRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) Assignments()(*MobileAppsItemAssignmentsRequestBuilder) {
    return NewMobileAppsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileAppsItemCategoriesRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) Categories()(*MobileAppsItemCategoriesRequestBuilder) {
    return NewMobileAppsItemCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppsMobileAppItemRequestBuilderInternal instantiates a new MobileAppsMobileAppItemRequestBuilder and sets the default values.
func NewMobileAppsMobileAppItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsMobileAppItemRequestBuilder) {
    m := &MobileAppsMobileAppItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileAppsMobileAppItemRequestBuilder instantiates a new MobileAppsMobileAppItemRequestBuilder and sets the default values.
func NewMobileAppsMobileAppItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsMobileAppItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsMobileAppItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mobileApps for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsMobileAppItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the mobile apps.
// returns a MobileAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsMobileAppItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsMobileAppItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable), nil
}
// GraphAndroidForWorkApp casts the previous resource to androidForWorkApp.
// returns a *MobileAppsItemGraphAndroidForWorkAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphAndroidForWorkApp()(*MobileAppsItemGraphAndroidForWorkAppRequestBuilder) {
    return NewMobileAppsItemGraphAndroidForWorkAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAndroidLobApp casts the previous resource to androidLobApp.
// returns a *MobileAppsItemGraphAndroidLobAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphAndroidLobApp()(*MobileAppsItemGraphAndroidLobAppRequestBuilder) {
    return NewMobileAppsItemGraphAndroidLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAndroidManagedStoreApp casts the previous resource to androidManagedStoreApp.
// returns a *MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphAndroidManagedStoreApp()(*MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder) {
    return NewMobileAppsItemGraphAndroidManagedStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAndroidStoreApp casts the previous resource to androidStoreApp.
// returns a *MobileAppsItemGraphAndroidStoreAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphAndroidStoreApp()(*MobileAppsItemGraphAndroidStoreAppRequestBuilder) {
    return NewMobileAppsItemGraphAndroidStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosLobApp casts the previous resource to iosLobApp.
// returns a *MobileAppsItemGraphIosLobAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphIosLobApp()(*MobileAppsItemGraphIosLobAppRequestBuilder) {
    return NewMobileAppsItemGraphIosLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosStoreApp casts the previous resource to iosStoreApp.
// returns a *MobileAppsItemGraphIosStoreAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphIosStoreApp()(*MobileAppsItemGraphIosStoreAppRequestBuilder) {
    return NewMobileAppsItemGraphIosStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosVppApp casts the previous resource to iosVppApp.
// returns a *MobileAppsItemGraphIosVppAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphIosVppApp()(*MobileAppsItemGraphIosVppAppRequestBuilder) {
    return NewMobileAppsItemGraphIosVppAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSDmgApp casts the previous resource to macOSDmgApp.
// returns a *MobileAppsItemGraphMacOSDmgAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphMacOSDmgApp()(*MobileAppsItemGraphMacOSDmgAppRequestBuilder) {
    return NewMobileAppsItemGraphMacOSDmgAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSLobApp casts the previous resource to macOSLobApp.
// returns a *MobileAppsItemGraphMacOSLobAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphMacOSLobApp()(*MobileAppsItemGraphMacOSLobAppRequestBuilder) {
    return NewMobileAppsItemGraphMacOSLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSPkgApp casts the previous resource to macOSPkgApp.
// returns a *MobileAppsItemGraphMacOSPkgAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphMacOSPkgApp()(*MobileAppsItemGraphMacOSPkgAppRequestBuilder) {
    return NewMobileAppsItemGraphMacOSPkgAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedAndroidLobApp casts the previous resource to managedAndroidLobApp.
// returns a *MobileAppsItemGraphManagedAndroidLobAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphManagedAndroidLobApp()(*MobileAppsItemGraphManagedAndroidLobAppRequestBuilder) {
    return NewMobileAppsItemGraphManagedAndroidLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedIOSLobApp casts the previous resource to managedIOSLobApp.
// returns a *MobileAppsItemGraphManagedIOSLobAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphManagedIOSLobApp()(*MobileAppsItemGraphManagedIOSLobAppRequestBuilder) {
    return NewMobileAppsItemGraphManagedIOSLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedMobileLobApp casts the previous resource to managedMobileLobApp.
// returns a *MobileAppsItemGraphManagedMobileLobAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphManagedMobileLobApp()(*MobileAppsItemGraphManagedMobileLobAppRequestBuilder) {
    return NewMobileAppsItemGraphManagedMobileLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMicrosoftStoreForBusinessApp casts the previous resource to microsoftStoreForBusinessApp.
// returns a *MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphMicrosoftStoreForBusinessApp()(*MobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    return NewMobileAppsItemGraphMicrosoftStoreForBusinessAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWin32LobApp casts the previous resource to win32LobApp.
// returns a *MobileAppsItemGraphWin32LobAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphWin32LobApp()(*MobileAppsItemGraphWin32LobAppRequestBuilder) {
    return NewMobileAppsItemGraphWin32LobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsAppX casts the previous resource to windowsAppX.
// returns a *MobileAppsItemGraphWindowsAppXRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphWindowsAppX()(*MobileAppsItemGraphWindowsAppXRequestBuilder) {
    return NewMobileAppsItemGraphWindowsAppXRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsMobileMSI casts the previous resource to windowsMobileMSI.
// returns a *MobileAppsItemGraphWindowsMobileMSIRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphWindowsMobileMSI()(*MobileAppsItemGraphWindowsMobileMSIRequestBuilder) {
    return NewMobileAppsItemGraphWindowsMobileMSIRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsStoreApp casts the previous resource to windowsStoreApp.
// returns a *MobileAppsItemGraphWindowsStoreAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphWindowsStoreApp()(*MobileAppsItemGraphWindowsStoreAppRequestBuilder) {
    return NewMobileAppsItemGraphWindowsStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsUniversalAppX casts the previous resource to windowsUniversalAppX.
// returns a *MobileAppsItemGraphWindowsUniversalAppXRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphWindowsUniversalAppX()(*MobileAppsItemGraphWindowsUniversalAppXRequestBuilder) {
    return NewMobileAppsItemGraphWindowsUniversalAppXRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsWebApp casts the previous resource to windowsWebApp.
// returns a *MobileAppsItemGraphWindowsWebAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphWindowsWebApp()(*MobileAppsItemGraphWindowsWebAppRequestBuilder) {
    return NewMobileAppsItemGraphWindowsWebAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWinGetApp casts the previous resource to winGetApp.
// returns a *MobileAppsItemGraphWinGetAppRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) GraphWinGetApp()(*MobileAppsItemGraphWinGetAppRequestBuilder) {
    return NewMobileAppsItemGraphWinGetAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property mobileApps in deviceAppManagement
// returns a MobileAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsMobileAppItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *MobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable), nil
}
// Relationships provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
// returns a *MobileAppsItemRelationshipsRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) Relationships()(*MobileAppsItemRelationshipsRequestBuilder) {
    return NewMobileAppsItemRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property mobileApps for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileAppsMobileAppItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the mobile apps.
// returns a *RequestInformation when successful
func (m *MobileAppsMobileAppItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsMobileAppItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mobileApps in deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileAppsMobileAppItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *MobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UpdateRelationships provides operations to call the updateRelationships method.
// returns a *MobileAppsItemUpdateRelationshipsRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) UpdateRelationships()(*MobileAppsItemUpdateRelationshipsRequestBuilder) {
    return NewMobileAppsItemUpdateRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileAppsMobileAppItemRequestBuilder when successful
func (m *MobileAppsMobileAppItemRequestBuilder) WithUrl(rawUrl string)(*MobileAppsMobileAppItemRequestBuilder) {
    return NewMobileAppsMobileAppItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
