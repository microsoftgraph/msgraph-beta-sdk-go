package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder provides operations to manage the files property of the microsoft.graph.mobileAppContent entity.
type MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderGetQueryParameters the list of files for this app content version.
type MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderGetQueryParameters struct {
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
// MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderGetQueryParameters
}
// MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMobileAppContentFileId provides operations to manage the files property of the microsoft.graph.mobileAppContent entity.
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder) ByMobileAppContentFileId(mobileAppContentFileId string)(*MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mobileAppContentFileId != "" {
        urlTplParams["mobileAppContentFile%2Did"] = mobileAppContentFileId
    }
    return NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderInternal instantiates a new FilesRequestBuilder and sets the default values.
func NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder) {
    m := &MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.macOSPkgApp/contentVersions/{mobileAppContent%2Did}/files{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder instantiates a new FilesRequestBuilder and sets the default values.
func NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder) Count()(*MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesCountRequestBuilder) {
    return NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of files for this app content version.
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentFileCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppContentFileCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentFileCollectionResponseable), nil
}
// Post create new navigation property to files for deviceAppManagement
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentFileable, requestConfiguration *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentFileable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppContentFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentFileable), nil
}
// ToGetRequestInformation the list of files for this app content version.
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to files for deviceAppManagement
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentFileable, requestConfiguration *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder) {
    return NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
