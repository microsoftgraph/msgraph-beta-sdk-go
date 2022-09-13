package registeredowners

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i491a5dda2581c32fb991ffa3306204abde4a1de5081c6a90f87ca62e4f4dee77 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners/serviceprincipal"
    i8e92484cb0a38af48f0db62f2949c2333937d97001d44a01b2b04dccccb081be "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners/user"
    ib674e7b0208ee4ccebca4f4c07f6b7f8ca4d15e3735a196c9d0d5aaac889db6f "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners/count"
    ieb7c0b5bc4163d6c3c875ff53f15a4dd88b951baee3d68093c9407119ce6a4a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners/endpoint"
    if6a775172c9b3745e19a07f2cfd4ae36fd8ec055443f76d3cb90ecc478b009b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners/ref"
)

// RegisteredOwnersRequestBuilder provides operations to manage the registeredOwners property of the microsoft.graph.device entity.
type RegisteredOwnersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RegisteredOwnersRequestBuilderGetQueryParameters the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
type RegisteredOwnersRequestBuilderGetQueryParameters struct {
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
// RegisteredOwnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RegisteredOwnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RegisteredOwnersRequestBuilderGetQueryParameters
}
// NewRegisteredOwnersRequestBuilderInternal instantiates a new RegisteredOwnersRequestBuilder and sets the default values.
func NewRegisteredOwnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegisteredOwnersRequestBuilder) {
    m := &RegisteredOwnersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/devices/{device%2Did}/registeredOwners{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRegisteredOwnersRequestBuilder instantiates a new RegisteredOwnersRequestBuilder and sets the default values.
func NewRegisteredOwnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RegisteredOwnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRegisteredOwnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *RegisteredOwnersRequestBuilder) Count()(*ib674e7b0208ee4ccebca4f4c07f6b7f8ca4d15e3735a196c9d0d5aaac889db6f.CountRequestBuilder) {
    return ib674e7b0208ee4ccebca4f4c07f6b7f8ca4d15e3735a196c9d0d5aaac889db6f.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *RegisteredOwnersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *RegisteredOwnersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RegisteredOwnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RegisteredOwnersRequestBuilder) Endpoint()(*ieb7c0b5bc4163d6c3c875ff53f15a4dd88b951baee3d68093c9407119ce6a4a5.EndpointRequestBuilder) {
    return ieb7c0b5bc4163d6c3c875ff53f15a4dd88b951baee3d68093c9407119ce6a4a5.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *RegisteredOwnersRequestBuilder) Get(ctx context.Context, requestConfiguration *RegisteredOwnersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
func (m *RegisteredOwnersRequestBuilder) Ref()(*if6a775172c9b3745e19a07f2cfd4ae36fd8ec055443f76d3cb90ecc478b009b5.RefRequestBuilder) {
    return if6a775172c9b3745e19a07f2cfd4ae36fd8ec055443f76d3cb90ecc478b009b5.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *RegisteredOwnersRequestBuilder) ServicePrincipal()(*i491a5dda2581c32fb991ffa3306204abde4a1de5081c6a90f87ca62e4f4dee77.ServicePrincipalRequestBuilder) {
    return i491a5dda2581c32fb991ffa3306204abde4a1de5081c6a90f87ca62e4f4dee77.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *RegisteredOwnersRequestBuilder) User()(*i8e92484cb0a38af48f0db62f2949c2333937d97001d44a01b2b04dccccb081be.UserRequestBuilder) {
    return i8e92484cb0a38af48f0db62f2949c2333937d97001d44a01b2b04dccccb081be.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
