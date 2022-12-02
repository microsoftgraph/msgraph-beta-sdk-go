package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemInvalidateAllRefreshTokensRequestBuilder provides operations to call the invalidateAllRefreshTokens method.
type UsersItemInvalidateAllRefreshTokensRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemInvalidateAllRefreshTokensRequestBuilderInternal instantiates a new InvalidateAllRefreshTokensRequestBuilder and sets the default values.
func NewUsersItemInvalidateAllRefreshTokensRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemInvalidateAllRefreshTokensRequestBuilder) {
    m := &UsersItemInvalidateAllRefreshTokensRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/microsoft.graph.invalidateAllRefreshTokens";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemInvalidateAllRefreshTokensRequestBuilder instantiates a new InvalidateAllRefreshTokensRequestBuilder and sets the default values.
func NewUsersItemInvalidateAllRefreshTokensRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemInvalidateAllRefreshTokensRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemInvalidateAllRefreshTokensRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invalidates all of the user's refresh tokens issued to applications (as well as session cookies in a user's browser), by resetting the **refreshTokensValidFromDateTime** user property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the user has a lost or stolen device.  This operation would prevent access to any of the organization's data accessed through applications on the device without the user first being required to sign in again. In fact, this operation would force the user to sign in again for all applications that they have previously consented to, independent of device. For developers, if the application attempts to redeem a delegated access token for this user by using an invalidated refresh token, the application will get an error. If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint, which will force the user to sign in.
func (m *UsersItemInvalidateAllRefreshTokensRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *UsersItemInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invalidates all of the user's refresh tokens issued to applications (as well as session cookies in a user's browser), by resetting the **refreshTokensValidFromDateTime** user property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the user has a lost or stolen device.  This operation would prevent access to any of the organization's data accessed through applications on the device without the user first being required to sign in again. In fact, this operation would force the user to sign in again for all applications that they have previously consented to, independent of device. For developers, if the application attempts to redeem a delegated access token for this user by using an invalidated refresh token, the application will get an error. If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint, which will force the user to sign in.
func (m *UsersItemInvalidateAllRefreshTokensRequestBuilder) Post(ctx context.Context, requestConfiguration *UsersItemInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemInvalidateAllRefreshTokensResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUsersItemInvalidateAllRefreshTokensResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemInvalidateAllRefreshTokensResponseable), nil
}
