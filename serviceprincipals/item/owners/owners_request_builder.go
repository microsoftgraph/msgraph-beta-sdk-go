package owners

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i93ef759c0e28cdc3c60e0a07e1aa51f5753d0620552b81815b0f256154c44040 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/user"
    ib38e1ffd94712b28166b02f16b2e83d55d42618641811fe9972045de7fb9a7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/endpoint"
    ic260a1a9f0014e6557bfcccd22b9fcf95023f8eb041dd35c794d63c782327c02 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/serviceprincipal"
    id07c90722ad0688242e77859cd090ff8c47280fe234fd1a48f6bc0104f5cca15 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/count"
    if6165a897e433e529e846d0cb8fe2e9839649ad09b949ff71a3b21dea2dbf2d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/ref"
)

// OwnersRequestBuilder provides operations to manage the owners property of the microsoft.graph.servicePrincipal entity.
type OwnersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OwnersRequestBuilderGetQueryParameters directory objects that are owners of this servicePrincipal. The owners are a set of non-admin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand.
type OwnersRequestBuilderGetQueryParameters struct {
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
// OwnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OwnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OwnersRequestBuilderGetQueryParameters
}
// NewOwnersRequestBuilderInternal instantiates a new OwnersRequestBuilder and sets the default values.
func NewOwnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnersRequestBuilder) {
    m := &OwnersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/owners{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOwnersRequestBuilder instantiates a new OwnersRequestBuilder and sets the default values.
func NewOwnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOwnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *OwnersRequestBuilder) Count()(*id07c90722ad0688242e77859cd090ff8c47280fe234fd1a48f6bc0104f5cca15.CountRequestBuilder) {
    return id07c90722ad0688242e77859cd090ff8c47280fe234fd1a48f6bc0104f5cca15.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation directory objects that are owners of this servicePrincipal. The owners are a set of non-admin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand.
func (m *OwnersRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OwnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Endpoint the endpoint property
func (m *OwnersRequestBuilder) Endpoint()(*ib38e1ffd94712b28166b02f16b2e83d55d42618641811fe9972045de7fb9a7bc.EndpointRequestBuilder) {
    return ib38e1ffd94712b28166b02f16b2e83d55d42618641811fe9972045de7fb9a7bc.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get directory objects that are owners of this servicePrincipal. The owners are a set of non-admin users or servicePrincipals who are allowed to modify this object. Read-only. Nullable. Supports $expand.
func (m *OwnersRequestBuilder) Get(ctx context.Context, requestConfiguration *OwnersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// Ref the Ref property
func (m *OwnersRequestBuilder) Ref()(*if6165a897e433e529e846d0cb8fe2e9839649ad09b949ff71a3b21dea2dbf2d3.RefRequestBuilder) {
    return if6165a897e433e529e846d0cb8fe2e9839649ad09b949ff71a3b21dea2dbf2d3.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *OwnersRequestBuilder) ServicePrincipal()(*ic260a1a9f0014e6557bfcccd22b9fcf95023f8eb041dd35c794d63c782327c02.ServicePrincipalRequestBuilder) {
    return ic260a1a9f0014e6557bfcccd22b9fcf95023f8eb041dd35c794d63c782327c02.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *OwnersRequestBuilder) User()(*i93ef759c0e28cdc3c60e0a07e1aa51f5753d0620552b81815b0f256154c44040.UserRequestBuilder) {
    return i93ef759c0e28cdc3c60e0a07e1aa51f5753d0620552b81815b0f256154c44040.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
