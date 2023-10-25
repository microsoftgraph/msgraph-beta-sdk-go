package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsRequestBuilder provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
type MobileAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsRequestBuilderGetQueryParameters the mobile apps.
type MobileAppsRequestBuilderGetQueryParameters struct {
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
// MobileAppsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsRequestBuilderGetQueryParameters
}
// MobileAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMobileAppId provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
func (m *MobileAppsRequestBuilder) ByMobileAppId(mobileAppId string)(*MobileAppsMobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mobileAppId != "" {
        urlTplParams["mobileApp%2Did"] = mobileAppId
    }
    return NewMobileAppsMobileAppItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppsRequestBuilderInternal instantiates a new MobileAppsRequestBuilder and sets the default values.
func NewMobileAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsRequestBuilder) {
    m := &MobileAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMobileAppsRequestBuilder instantiates a new MobileAppsRequestBuilder and sets the default values.
func NewMobileAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *MobileAppsRequestBuilder) Count()(*MobileAppsCountRequestBuilder) {
    return NewMobileAppsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the mobile apps.
func (m *MobileAppsRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *MobileAppsRequestBuilder) GraphAndroidForWorkApp()(*MobileAppsGraphAndroidForWorkAppRequestBuilder) {
    return NewMobileAppsGraphAndroidForWorkAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAndroidLobApp casts the previous resource to androidLobApp.
func (m *MobileAppsRequestBuilder) GraphAndroidLobApp()(*MobileAppsGraphAndroidLobAppRequestBuilder) {
    return NewMobileAppsGraphAndroidLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAndroidManagedStoreApp casts the previous resource to androidManagedStoreApp.
func (m *MobileAppsRequestBuilder) GraphAndroidManagedStoreApp()(*MobileAppsGraphAndroidManagedStoreAppRequestBuilder) {
    return NewMobileAppsGraphAndroidManagedStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphAndroidStoreApp casts the previous resource to androidStoreApp.
func (m *MobileAppsRequestBuilder) GraphAndroidStoreApp()(*MobileAppsGraphAndroidStoreAppRequestBuilder) {
    return NewMobileAppsGraphAndroidStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosLobApp casts the previous resource to iosLobApp.
func (m *MobileAppsRequestBuilder) GraphIosLobApp()(*MobileAppsGraphIosLobAppRequestBuilder) {
    return NewMobileAppsGraphIosLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosStoreApp casts the previous resource to iosStoreApp.
func (m *MobileAppsRequestBuilder) GraphIosStoreApp()(*MobileAppsGraphIosStoreAppRequestBuilder) {
    return NewMobileAppsGraphIosStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphIosVppApp casts the previous resource to iosVppApp.
func (m *MobileAppsRequestBuilder) GraphIosVppApp()(*MobileAppsGraphIosVppAppRequestBuilder) {
    return NewMobileAppsGraphIosVppAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSDmgApp casts the previous resource to macOSDmgApp.
func (m *MobileAppsRequestBuilder) GraphMacOSDmgApp()(*MobileAppsGraphMacOSDmgAppRequestBuilder) {
    return NewMobileAppsGraphMacOSDmgAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSLobApp casts the previous resource to macOSLobApp.
func (m *MobileAppsRequestBuilder) GraphMacOSLobApp()(*MobileAppsGraphMacOSLobAppRequestBuilder) {
    return NewMobileAppsGraphMacOSLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMacOSPkgApp casts the previous resource to macOSPkgApp.
func (m *MobileAppsRequestBuilder) GraphMacOSPkgApp()(*MobileAppsGraphMacOSPkgAppRequestBuilder) {
    return NewMobileAppsGraphMacOSPkgAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedAndroidLobApp casts the previous resource to managedAndroidLobApp.
func (m *MobileAppsRequestBuilder) GraphManagedAndroidLobApp()(*MobileAppsGraphManagedAndroidLobAppRequestBuilder) {
    return NewMobileAppsGraphManagedAndroidLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedIOSLobApp casts the previous resource to managedIOSLobApp.
func (m *MobileAppsRequestBuilder) GraphManagedIOSLobApp()(*MobileAppsGraphManagedIOSLobAppRequestBuilder) {
    return NewMobileAppsGraphManagedIOSLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphManagedMobileLobApp casts the previous resource to managedMobileLobApp.
func (m *MobileAppsRequestBuilder) GraphManagedMobileLobApp()(*MobileAppsGraphManagedMobileLobAppRequestBuilder) {
    return NewMobileAppsGraphManagedMobileLobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphMicrosoftStoreForBusinessApp casts the previous resource to microsoftStoreForBusinessApp.
func (m *MobileAppsRequestBuilder) GraphMicrosoftStoreForBusinessApp()(*MobileAppsGraphMicrosoftStoreForBusinessAppRequestBuilder) {
    return NewMobileAppsGraphMicrosoftStoreForBusinessAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWin32LobApp casts the previous resource to win32LobApp.
func (m *MobileAppsRequestBuilder) GraphWin32LobApp()(*MobileAppsGraphWin32LobAppRequestBuilder) {
    return NewMobileAppsGraphWin32LobAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsAppX casts the previous resource to windowsAppX.
func (m *MobileAppsRequestBuilder) GraphWindowsAppX()(*MobileAppsGraphWindowsAppXRequestBuilder) {
    return NewMobileAppsGraphWindowsAppXRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsMobileMSI casts the previous resource to windowsMobileMSI.
func (m *MobileAppsRequestBuilder) GraphWindowsMobileMSI()(*MobileAppsGraphWindowsMobileMSIRequestBuilder) {
    return NewMobileAppsGraphWindowsMobileMSIRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsStoreApp casts the previous resource to windowsStoreApp.
func (m *MobileAppsRequestBuilder) GraphWindowsStoreApp()(*MobileAppsGraphWindowsStoreAppRequestBuilder) {
    return NewMobileAppsGraphWindowsStoreAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsUniversalAppX casts the previous resource to windowsUniversalAppX.
func (m *MobileAppsRequestBuilder) GraphWindowsUniversalAppX()(*MobileAppsGraphWindowsUniversalAppXRequestBuilder) {
    return NewMobileAppsGraphWindowsUniversalAppXRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWindowsWebApp casts the previous resource to windowsWebApp.
func (m *MobileAppsRequestBuilder) GraphWindowsWebApp()(*MobileAppsGraphWindowsWebAppRequestBuilder) {
    return NewMobileAppsGraphWindowsWebAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphWinGetApp casts the previous resource to winGetApp.
func (m *MobileAppsRequestBuilder) GraphWinGetApp()(*MobileAppsGraphWinGetAppRequestBuilder) {
    return NewMobileAppsGraphWinGetAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HasPayloadLinks provides operations to call the hasPayloadLinks method.
func (m *MobileAppsRequestBuilder) HasPayloadLinks()(*MobileAppsHasPayloadLinksRequestBuilder) {
    return NewMobileAppsHasPayloadLinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to mobileApps for deviceAppManagement
func (m *MobileAppsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *MobileAppsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *MobileAppsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to mobileApps for deviceAppManagement
func (m *MobileAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *MobileAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ValidateXml provides operations to call the validateXml method.
func (m *MobileAppsRequestBuilder) ValidateXml()(*MobileAppsValidateXmlRequestBuilder) {
    return NewMobileAppsValidateXmlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MobileAppsRequestBuilder) WithUrl(rawUrl string)(*MobileAppsRequestBuilder) {
    return NewMobileAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
