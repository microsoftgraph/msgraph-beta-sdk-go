package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileappsMobileAppsRequestBuilder provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
type MobileappsMobileAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsMobileAppsRequestBuilderGetQueryParameters the mobile apps.
type MobileappsMobileAppsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MobileappsMobileAppsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsMobileAppsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsMobileAppsRequestBuilderGetQueryParameters
}
// MobileappsMobileAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsMobileAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMobileAppId provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
// returns a *MobileappsMobileAppItemRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) ByMobileAppId(mobileAppId string)(*MobileappsMobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mobileAppId != "" {
        urlTplParams["mobileApp%2Did"] = mobileAppId
    }
    return NewMobileappsMobileAppItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsMobileAppsRequestBuilderInternal instantiates a new MobileappsMobileAppsRequestBuilder and sets the default values.
func NewMobileappsMobileAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsMobileAppsRequestBuilder) {
    m := &MobileappsMobileAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobileappsMobileAppsRequestBuilder instantiates a new MobileappsMobileAppsRequestBuilder and sets the default values.
func NewMobileappsMobileAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsMobileAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsMobileAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// ConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageId provides operations to call the convertFromMobileAppCatalogPackage method.
// returns a *MobileappsConvertfrommobileappcatalogpackagewithmobileappcatalogpackageidConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) ConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageId(mobileAppCatalogPackageId *string)(*MobileappsConvertfrommobileappcatalogpackagewithmobileappcatalogpackageidConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder) {
    return NewMobileappsConvertfrommobileappcatalogpackagewithmobileappcatalogpackageidConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, mobileAppCatalogPackageId)
}
// Count provides operations to count the resources in the collection.
// returns a *MobileappsCountRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) Count()(*MobileappsCountRequestBuilder) {
    return NewMobileappsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the mobile apps.
// returns a MobileAppCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsMobileAppsRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsMobileAppsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppCollectionResponseable), nil
}
// GraphAndroidForWorkApp casts the previous resource to androidForWorkApp.
// returns a *MobileappsGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphAndroidForWorkApp()(*MobileappsGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder) {
    return NewMobileappsGraphandroidforworkappGraphAndroidForWorkAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAndroidLobApp casts the previous resource to androidLobApp.
// returns a *MobileappsGraphandroidlobappGraphAndroidLobAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphAndroidLobApp()(*MobileappsGraphandroidlobappGraphAndroidLobAppRequestBuilder) {
    return NewMobileappsGraphandroidlobappGraphAndroidLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAndroidManagedStoreApp casts the previous resource to androidManagedStoreApp.
// returns a *MobileappsGraphandroidmanagedstoreappGraphAndroidManagedStoreAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphAndroidManagedStoreApp()(*MobileappsGraphandroidmanagedstoreappGraphAndroidManagedStoreAppRequestBuilder) {
    return NewMobileappsGraphandroidmanagedstoreappGraphAndroidManagedStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAndroidStoreApp casts the previous resource to androidStoreApp.
// returns a *MobileappsGraphandroidstoreappGraphAndroidStoreAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphAndroidStoreApp()(*MobileappsGraphandroidstoreappGraphAndroidStoreAppRequestBuilder) {
    return NewMobileappsGraphandroidstoreappGraphAndroidStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosLobApp casts the previous resource to iosLobApp.
// returns a *MobileappsGraphioslobappGraphIosLobAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphIosLobApp()(*MobileappsGraphioslobappGraphIosLobAppRequestBuilder) {
    return NewMobileappsGraphioslobappGraphIosLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosStoreApp casts the previous resource to iosStoreApp.
// returns a *MobileappsGraphiosstoreappGraphIosStoreAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphIosStoreApp()(*MobileappsGraphiosstoreappGraphIosStoreAppRequestBuilder) {
    return NewMobileappsGraphiosstoreappGraphIosStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosVppApp casts the previous resource to iosVppApp.
// returns a *MobileappsGraphiosvppappGraphIosVppAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphIosVppApp()(*MobileappsGraphiosvppappGraphIosVppAppRequestBuilder) {
    return NewMobileappsGraphiosvppappGraphIosVppAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSDmgApp casts the previous resource to macOSDmgApp.
// returns a *MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphMacOSDmgApp()(*MobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilder) {
    return NewMobileappsGraphmacosdmgappGraphMacOSDmgAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSLobApp casts the previous resource to macOSLobApp.
// returns a *MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphMacOSLobApp()(*MobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilder) {
    return NewMobileappsGraphmacoslobappGraphMacOSLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSPkgApp casts the previous resource to macOSPkgApp.
// returns a *MobileappsGraphmacospkgappGraphMacOSPkgAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphMacOSPkgApp()(*MobileappsGraphmacospkgappGraphMacOSPkgAppRequestBuilder) {
    return NewMobileappsGraphmacospkgappGraphMacOSPkgAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedAndroidLobApp casts the previous resource to managedAndroidLobApp.
// returns a *MobileappsGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphManagedAndroidLobApp()(*MobileappsGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilder) {
    return NewMobileappsGraphmanagedandroidlobappGraphManagedAndroidLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedIOSLobApp casts the previous resource to managedIOSLobApp.
// returns a *MobileappsGraphmanagedioslobappGraphManagedIOSLobAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphManagedIOSLobApp()(*MobileappsGraphmanagedioslobappGraphManagedIOSLobAppRequestBuilder) {
    return NewMobileappsGraphmanagedioslobappGraphManagedIOSLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedMobileLobApp casts the previous resource to managedMobileLobApp.
// returns a *MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphManagedMobileLobApp()(*MobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilder) {
    return NewMobileappsGraphmanagedmobilelobappGraphManagedMobileLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMicrosoftStoreForBusinessApp casts the previous resource to microsoftStoreForBusinessApp.
// returns a *MobileappsGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphMicrosoftStoreForBusinessApp()(*MobileappsGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    return NewMobileappsGraphmicrosoftstoreforbusinessappGraphMicrosoftStoreForBusinessAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWin32LobApp casts the previous resource to win32LobApp.
// returns a *MobileappsGraphwin32lobappGraphWin32LobAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphWin32LobApp()(*MobileappsGraphwin32lobappGraphWin32LobAppRequestBuilder) {
    return NewMobileappsGraphwin32lobappGraphWin32LobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsAppX casts the previous resource to windowsAppX.
// returns a *MobileappsGraphwindowsappxGraphWindowsAppXRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphWindowsAppX()(*MobileappsGraphwindowsappxGraphWindowsAppXRequestBuilder) {
    return NewMobileappsGraphwindowsappxGraphWindowsAppXRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsMobileMSI casts the previous resource to windowsMobileMSI.
// returns a *MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphWindowsMobileMSI()(*MobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilder) {
    return NewMobileappsGraphwindowsmobilemsiGraphWindowsMobileMSIRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsStoreApp casts the previous resource to windowsStoreApp.
// returns a *MobileappsGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphWindowsStoreApp()(*MobileappsGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder) {
    return NewMobileappsGraphwindowsstoreappGraphWindowsStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsUniversalAppX casts the previous resource to windowsUniversalAppX.
// returns a *MobileappsGraphwindowsuniversalappxGraphWindowsUniversalAppXRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphWindowsUniversalAppX()(*MobileappsGraphwindowsuniversalappxGraphWindowsUniversalAppXRequestBuilder) {
    return NewMobileappsGraphwindowsuniversalappxGraphWindowsUniversalAppXRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsWebApp casts the previous resource to windowsWebApp.
// returns a *MobileappsGraphwindowswebappGraphWindowsWebAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphWindowsWebApp()(*MobileappsGraphwindowswebappGraphWindowsWebAppRequestBuilder) {
    return NewMobileappsGraphwindowswebappGraphWindowsWebAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWinGetApp casts the previous resource to winGetApp.
// returns a *MobileappsGraphwingetappGraphWinGetAppRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) GraphWinGetApp()(*MobileappsGraphwingetappGraphWinGetAppRequestBuilder) {
    return NewMobileappsGraphwingetappGraphWinGetAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HasPayloadLinks provides operations to call the hasPayloadLinks method.
// returns a *MobileappsHaspayloadlinksHasPayloadLinksRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) HasPayloadLinks()(*MobileappsHaspayloadlinksHasPayloadLinksRequestBuilder) {
    return NewMobileappsHaspayloadlinksHasPayloadLinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to mobileApps for deviceAppManagement
// returns a MobileAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsMobileAppsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *MobileappsMobileAppsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the mobile apps.
// returns a *RequestInformation when successful
func (m *MobileappsMobileAppsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsMobileAppsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to mobileApps for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappsMobileAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *MobileappsMobileAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// ValidateXml provides operations to call the validateXml method.
// returns a *MobileappsValidatexmlValidateXmlRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) ValidateXml()(*MobileappsValidatexmlValidateXmlRequestBuilder) {
    return NewMobileappsValidatexmlValidateXmlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileappsMobileAppsRequestBuilder when successful
func (m *MobileappsMobileAppsRequestBuilder) WithUrl(rawUrl string)(*MobileappsMobileAppsRequestBuilder) {
    return NewMobileappsMobileAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
