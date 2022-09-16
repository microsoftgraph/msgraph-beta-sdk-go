package owners

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i10f832541960426e3ed7a66cfc5be95d2306c71ced17d6fc71ab18a0c0c5b4c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/owners/user"
    i5a1da955cc48ec0732456545bba0802e31506c787c6f77e70d6f7611c21351b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/owners/serviceprincipal"
    i906675c7ba5aa8b4867252f606e20b495748a780f537372ccfee50eaed4ad798 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/owners/ref"
    i9f7f3411a83113cc0d5bb69e3f5b6fe5d10845cc71e07818449ad638ae8520ec "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/owners/count"
    ieacb60fe51334c0dbe81e09e8b0549581ff1d791be3271be1f719508fa2ec76a "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/owners/endpoint"
)

// OwnersRequestBuilder provides operations to manage the owners property of the microsoft.graph.application entity.
type OwnersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OwnersRequestBuilderGetQueryParameters directory objects that are owners of the application. Read-only. Nullable. Supports $expand and $filter (eq and ne when counting empty collections).
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
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/owners{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
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
func (m *OwnersRequestBuilder) Count()(*i9f7f3411a83113cc0d5bb69e3f5b6fe5d10845cc71e07818449ad638ae8520ec.CountRequestBuilder) {
    return i9f7f3411a83113cc0d5bb69e3f5b6fe5d10845cc71e07818449ad638ae8520ec.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation directory objects that are owners of the application. Read-only. Nullable. Supports $expand and $filter (eq and ne when counting empty collections).
func (m *OwnersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration directory objects that are owners of the application. Read-only. Nullable. Supports $expand and $filter (eq and ne when counting empty collections).
func (m *OwnersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OwnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OwnersRequestBuilder) Endpoint()(*ieacb60fe51334c0dbe81e09e8b0549581ff1d791be3271be1f719508fa2ec76a.EndpointRequestBuilder) {
    return ieacb60fe51334c0dbe81e09e8b0549581ff1d791be3271be1f719508fa2ec76a.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get directory objects that are owners of the application. Read-only. Nullable. Supports $expand and $filter (eq and ne when counting empty collections).
func (m *OwnersRequestBuilder) Get(ctx context.Context, requestConfiguration *OwnersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *OwnersRequestBuilder) Ref()(*i906675c7ba5aa8b4867252f606e20b495748a780f537372ccfee50eaed4ad798.RefRequestBuilder) {
    return i906675c7ba5aa8b4867252f606e20b495748a780f537372ccfee50eaed4ad798.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *OwnersRequestBuilder) ServicePrincipal()(*i5a1da955cc48ec0732456545bba0802e31506c787c6f77e70d6f7611c21351b1.ServicePrincipalRequestBuilder) {
    return i5a1da955cc48ec0732456545bba0802e31506c787c6f77e70d6f7611c21351b1.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *OwnersRequestBuilder) User()(*i10f832541960426e3ed7a66cfc5be95d2306c71ced17d6fc71ab18a0c0c5b4c7.UserRequestBuilder) {
    return i10f832541960426e3ed7a66cfc5be95d2306c71ced17d6fc71ab18a0c0c5b4c7.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
