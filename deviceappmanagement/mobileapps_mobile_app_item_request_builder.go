package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileappsMobileAppItemRequestBuilder provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
type MobileappsMobileAppItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsMobileAppItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsMobileAppItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileappsMobileAppItemRequestBuilderGetQueryParameters the mobile apps.
type MobileappsMobileAppItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsMobileAppItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsMobileAppItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsMobileAppItemRequestBuilderGetQueryParameters
}
// MobileappsMobileAppItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsMobileAppItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *MobileappsItemAssignRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) Assign()(*MobileappsItemAssignRequestBuilder) {
    return NewMobileappsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemAssignmentsRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) Assignments()(*MobileappsItemAssignmentsRequestBuilder) {
    return NewMobileappsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemCategoriesRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) Categories()(*MobileappsItemCategoriesRequestBuilder) {
    return NewMobileappsItemCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsMobileAppItemRequestBuilderInternal instantiates a new MobileappsMobileAppItemRequestBuilder and sets the default values.
func NewMobileappsMobileAppItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsMobileAppItemRequestBuilder) {
    m := &MobileappsMobileAppItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsMobileAppItemRequestBuilder instantiates a new MobileappsMobileAppItemRequestBuilder and sets the default values.
func NewMobileappsMobileAppItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsMobileAppItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsMobileAppItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mobileApps for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsMobileAppItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileappsMobileAppItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *MobileappsMobileAppItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsMobileAppItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, error) {
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
// returns a *MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphAndroidForWorkApp()(*MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder) {
    return NewMobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAndroidLobApp casts the previous resource to androidLobApp.
// returns a *MobileappsItemGraphandroidlobappGraphAndroidLobAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphAndroidLobApp()(*MobileappsItemGraphandroidlobappGraphAndroidLobAppRequestBuilder) {
    return NewMobileappsItemGraphandroidlobappGraphAndroidLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAndroidManagedStoreApp casts the previous resource to androidManagedStoreApp.
// returns a *MobileappsItemGraphandroidmanagedstoreappGraphAndroidManagedStoreAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphAndroidManagedStoreApp()(*MobileappsItemGraphandroidmanagedstoreappGraphAndroidManagedStoreAppRequestBuilder) {
    return NewMobileappsItemGraphandroidmanagedstoreappGraphAndroidManagedStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAndroidStoreApp casts the previous resource to androidStoreApp.
// returns a *MobileappsItemGraphandroidstoreappGraphAndroidStoreAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphAndroidStoreApp()(*MobileappsItemGraphandroidstoreappGraphAndroidStoreAppRequestBuilder) {
    return NewMobileappsItemGraphandroidstoreappGraphAndroidStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosLobApp casts the previous resource to iosLobApp.
// returns a *MobileappsItemGraphioslobappGraphIosLobAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphIosLobApp()(*MobileappsItemGraphioslobappGraphIosLobAppRequestBuilder) {
    return NewMobileappsItemGraphioslobappGraphIosLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosStoreApp casts the previous resource to iosStoreApp.
// returns a *MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphIosStoreApp()(*MobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilder) {
    return NewMobileappsItemGraphiosstoreappGraphIosStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosVppApp casts the previous resource to iosVppApp.
// returns a *MobileappsItemGraphiosvppappGraphIosVppAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphIosVppApp()(*MobileappsItemGraphiosvppappGraphIosVppAppRequestBuilder) {
    return NewMobileappsItemGraphiosvppappGraphIosVppAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSDmgApp casts the previous resource to macOSDmgApp.
// returns a *MobileappsItemGraphmacosdmgappGraphMacOSDmgAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphMacOSDmgApp()(*MobileappsItemGraphmacosdmgappGraphMacOSDmgAppRequestBuilder) {
    return NewMobileappsItemGraphmacosdmgappGraphMacOSDmgAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSLobApp casts the previous resource to macOSLobApp.
// returns a *MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphMacOSLobApp()(*MobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilder) {
    return NewMobileappsItemGraphmacoslobappGraphMacOSLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSPkgApp casts the previous resource to macOSPkgApp.
// returns a *MobileappsItemGraphmacospkgappGraphMacOSPkgAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphMacOSPkgApp()(*MobileappsItemGraphmacospkgappGraphMacOSPkgAppRequestBuilder) {
    return NewMobileappsItemGraphmacospkgappGraphMacOSPkgAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedAndroidLobApp casts the previous resource to managedAndroidLobApp.
// returns a *MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphManagedAndroidLobApp()(*MobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder) {
    return NewMobileappsItemGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedIOSLobApp casts the previous resource to managedIOSLobApp.
// returns a *MobileappsItemGraphmanagedioslobappGraphManagedIOSLobAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphManagedIOSLobApp()(*MobileappsItemGraphmanagedioslobappGraphManagedIOSLobAppRequestBuilder) {
    return NewMobileappsItemGraphmanagedioslobappGraphManagedIOSLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedMobileLobApp casts the previous resource to managedMobileLobApp.
// returns a *MobileappsItemGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphManagedMobileLobApp()(*MobileappsItemGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder) {
    return NewMobileappsItemGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMicrosoftStoreForBusinessApp casts the previous resource to microsoftStoreForBusinessApp.
// returns a *MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphMicrosoftStoreForBusinessApp()(*MobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    return NewMobileappsItemGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWin32LobApp casts the previous resource to win32LobApp.
// returns a *MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphWin32LobApp()(*MobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilder) {
    return NewMobileappsItemGraphwin32lobappGraphWin32LobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsAppX casts the previous resource to windowsAppX.
// returns a *MobileappsItemGraphwindowsappxGraphWindowsAppXRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphWindowsAppX()(*MobileappsItemGraphwindowsappxGraphWindowsAppXRequestBuilder) {
    return NewMobileappsItemGraphwindowsappxGraphWindowsAppXRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsMobileMSI casts the previous resource to windowsMobileMSI.
// returns a *MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphWindowsMobileMSI()(*MobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) {
    return NewMobileappsItemGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsStoreApp casts the previous resource to windowsStoreApp.
// returns a *MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphWindowsStoreApp()(*MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder) {
    return NewMobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsUniversalAppX casts the previous resource to windowsUniversalAppX.
// returns a *MobileappsItemGraphwindowsuniversalappxGraphWindowsUniversalAppXRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphWindowsUniversalAppX()(*MobileappsItemGraphwindowsuniversalappxGraphWindowsUniversalAppXRequestBuilder) {
    return NewMobileappsItemGraphwindowsuniversalappxGraphWindowsUniversalAppXRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsWebApp casts the previous resource to windowsWebApp.
// returns a *MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphWindowsWebApp()(*MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder) {
    return NewMobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWinGetApp casts the previous resource to winGetApp.
// returns a *MobileappsItemGraphwingetappGraphWinGetAppRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) GraphWinGetApp()(*MobileappsItemGraphwingetappGraphWinGetAppRequestBuilder) {
    return NewMobileappsItemGraphwingetappGraphWinGetAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property mobileApps in deviceAppManagement
// returns a MobileAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsMobileAppItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *MobileappsMobileAppItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, error) {
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
// returns a *MobileappsItemRelationshipsRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) Relationships()(*MobileappsItemRelationshipsRequestBuilder) {
    return NewMobileappsItemRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property mobileApps for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappsMobileAppItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileappsMobileAppItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MobileappsMobileAppItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsMobileAppItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MobileappsMobileAppItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *MobileappsMobileAppItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemUpdaterelationshipsUpdateRelationshipsRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) UpdateRelationships()(*MobileappsItemUpdaterelationshipsUpdateRelationshipsRequestBuilder) {
    return NewMobileappsItemUpdaterelationshipsUpdateRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileappsMobileAppItemRequestBuilder when successful
func (m *MobileappsMobileAppItemRequestBuilder) WithUrl(rawUrl string)(*MobileappsMobileAppItemRequestBuilder) {
    return NewMobileappsMobileAppItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
