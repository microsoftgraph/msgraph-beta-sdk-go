package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder provides operations to manage the temporaryAccessPassMethods property of the microsoft.graph.authentication entity.
type UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetQueryParameters represents a Temporary Access Pass registered to a user for authentication through time-limited passcodes.
type UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetQueryParameters
}
// NewUsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal instantiates a new TemporaryAccessPassAuthenticationMethodItemRequestBuilder and sets the default values.
func NewUsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) {
    m := &UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/temporaryAccessPassMethods/{temporaryAccessPassAuthenticationMethod%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder instantiates a new TemporaryAccessPassAuthenticationMethodItemRequestBuilder and sets the default values.
func NewUsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property temporaryAccessPassMethods for users
func (m *UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation represents a Temporary Access Pass registered to a user for authentication through time-limited passcodes.
func (m *UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property temporaryAccessPassMethods for users
func (m *UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get represents a Temporary Access Pass registered to a user for authentication through time-limited passcodes.
func (m *UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TemporaryAccessPassAuthenticationMethodable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTemporaryAccessPassAuthenticationMethodFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TemporaryAccessPassAuthenticationMethodable), nil
}
