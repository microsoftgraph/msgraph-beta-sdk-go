package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
type UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilderGetQueryParameters collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
type UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilderGetQueryParameters struct {
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
// UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilderGetQueryParameters
}
// NewUsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilderInternal instantiates a new RegisteredUsersRequestBuilder and sets the default values.
func NewUsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder) {
    m := &UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredUsers{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder instantiates a new RegisteredUsersRequestBuilder and sets the default values.
func NewUsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder) Count()(*UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersCountRequestBuilder) {
    return NewUsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Endpoint casts the previous resource to endpoint.
func (m *UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder) Endpoint()(*UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersEndpointRequestBuilder) {
    return NewUsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder) ServicePrincipal()(*UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersServicePrincipalRequestBuilder) {
    return NewUsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersRequestBuilder) User()(*UsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersUserRequestBuilder) {
    return NewUsersItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredUsersUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
