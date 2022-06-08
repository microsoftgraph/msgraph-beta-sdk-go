package registeredowners

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0e2674dd665949cd53a6a6a0910fd717c1455e7e871234d67c9c22d39862627e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/serviceprincipal"
    i63b29febff3bd41b27b8d48d19f6d06247e03d34e374c55fbeef70e90509659c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/endpoint"
    i86079b6385bb86fa78556c1bda63af69f6237b4c6cae56fa7c471a72692cd7b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/count"
    ie0b2e56ba24bb1848e9c7ec5406b036e51eddbbed4bf1aaf952ed0fbca9524d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/user"
    if6321cd5e3ae020ce8fd0661bb5b4f154f6429fbfc02006ebbec59f639fe4349 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/ref"
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
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
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
// Count the count property
func (m *RegisteredOwnersRequestBuilder) Count()(*i86079b6385bb86fa78556c1bda63af69f6237b4c6cae56fa7c471a72692cd7b4.CountRequestBuilder) {
    return i86079b6385bb86fa78556c1bda63af69f6237b4c6cae56fa7c471a72692cd7b4.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *RegisteredOwnersRequestBuilder) Endpoint()(*i63b29febff3bd41b27b8d48d19f6d06247e03d34e374c55fbeef70e90509659c.EndpointRequestBuilder) {
    return i63b29febff3bd41b27b8d48d19f6d06247e03d34e374c55fbeef70e90509659c.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *RegisteredOwnersRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *RegisteredOwnersRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *RegisteredOwnersRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// Ref the ref property
func (m *RegisteredOwnersRequestBuilder) Ref()(*if6321cd5e3ae020ce8fd0661bb5b4f154f6429fbfc02006ebbec59f639fe4349.RefRequestBuilder) {
    return if6321cd5e3ae020ce8fd0661bb5b4f154f6429fbfc02006ebbec59f639fe4349.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *RegisteredOwnersRequestBuilder) ServicePrincipal()(*i0e2674dd665949cd53a6a6a0910fd717c1455e7e871234d67c9c22d39862627e.ServicePrincipalRequestBuilder) {
    return i0e2674dd665949cd53a6a6a0910fd717c1455e7e871234d67c9c22d39862627e.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *RegisteredOwnersRequestBuilder) User()(*ie0b2e56ba24bb1848e9c7ec5406b036e51eddbbed4bf1aaf952ed0fbca9524d5.UserRequestBuilder) {
    return ie0b2e56ba24bb1848e9c7ec5406b036e51eddbbed4bf1aaf952ed0fbca9524d5.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
