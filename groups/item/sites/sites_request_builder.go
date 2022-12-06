package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i032f38299eb5e40683af93f6c784b710b3b5b086b76bd0ee59541f4d416334e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/remove"
    i94f4c3f907e63a555f457f8cef2f5323479fd57616082399625c4718c0769a01 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/count"
    iad4cabfe6e7b8c85c1ffb269d1fdb8abda5143efeab385fa69d6bf9a7bdc0205 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/add"
    id2234db9309f4c891a5adff1ae424657d5a8438f5b1cd8b39b3c365ca0f7dece "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/delta"
)

// SitesRequestBuilder provides operations to manage the sites property of the microsoft.graph.group entity.
type SitesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SitesRequestBuilderGetQueryParameters the list of SharePoint sites in this group. Access the default site with /sites/root.
type SitesRequestBuilderGetQueryParameters struct {
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
// SitesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SitesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SitesRequestBuilderGetQueryParameters
}
// Add provides operations to call the add method.
func (m *SitesRequestBuilder) Add()(*iad4cabfe6e7b8c85c1ffb269d1fdb8abda5143efeab385fa69d6bf9a7bdc0205.AddRequestBuilder) {
    return iad4cabfe6e7b8c85c1ffb269d1fdb8abda5143efeab385fa69d6bf9a7bdc0205.NewAddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSitesRequestBuilderInternal instantiates a new SitesRequestBuilder and sets the default values.
func NewSitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SitesRequestBuilder) {
    m := &SitesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSitesRequestBuilder instantiates a new SitesRequestBuilder and sets the default values.
func NewSitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *SitesRequestBuilder) Count()(*i94f4c3f907e63a555f457f8cef2f5323479fd57616082399625c4718c0769a01.CountRequestBuilder) {
    return i94f4c3f907e63a555f457f8cef2f5323479fd57616082399625c4718c0769a01.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *SitesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SitesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delta provides operations to call the delta method.
func (m *SitesRequestBuilder) Delta()(*id2234db9309f4c891a5adff1ae424657d5a8438f5b1cd8b39b3c365ca0f7dece.DeltaRequestBuilder) {
    return id2234db9309f4c891a5adff1ae424657d5a8438f5b1cd8b39b3c365ca0f7dece.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *SitesRequestBuilder) Get(ctx context.Context, requestConfiguration *SitesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteCollectionResponseable), nil
}
// Remove provides operations to call the remove method.
func (m *SitesRequestBuilder) Remove()(*i032f38299eb5e40683af93f6c784b710b3b5b086b76bd0ee59541f4d416334e7.RemoveRequestBuilder) {
    return i032f38299eb5e40683af93f6c784b710b3b5b086b76bd0ee59541f4d416334e7.NewRemoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
