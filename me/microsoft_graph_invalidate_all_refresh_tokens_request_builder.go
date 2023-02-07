package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftGraphInvalidateAllRefreshTokensRequestBuilder provides operations to call the invalidateAllRefreshTokens method.
type MicrosoftGraphInvalidateAllRefreshTokensRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphInvalidateAllRefreshTokensRequestBuilderInternal instantiates a new MicrosoftGraphInvalidateAllRefreshTokensRequestBuilder and sets the default values.
func NewMicrosoftGraphInvalidateAllRefreshTokensRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphInvalidateAllRefreshTokensRequestBuilder) {
    m := &MicrosoftGraphInvalidateAllRefreshTokensRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.invalidateAllRefreshTokens";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphInvalidateAllRefreshTokensRequestBuilder instantiates a new MicrosoftGraphInvalidateAllRefreshTokensRequestBuilder and sets the default values.
func NewMicrosoftGraphInvalidateAllRefreshTokensRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphInvalidateAllRefreshTokensRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphInvalidateAllRefreshTokensRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invalidates all of the user's refresh tokens issued to applications (as well as session cookies in a user's browser), by resetting the **refreshTokensValidFromDateTime** user property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the user has a lost or stolen device.  This operation would prevent access to any of the organization's data accessed through applications on the device without the user first being required to sign in again. In fact, this operation would force the user to sign in again for all applications that they have previously consented to, independent of device. For developers, if the application attempts to redeem a delegated access token for this user by using an invalidated refresh token, the application will get an error. If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint, which will force the user to sign in.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-invalidateallrefreshtokens?view=graph-rest-1.0
func (m *MicrosoftGraphInvalidateAllRefreshTokensRequestBuilder) Post(ctx context.Context, requestConfiguration *MicrosoftGraphInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration)(MicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponseable), nil
}
// ToPostRequestInformation invalidates all of the user's refresh tokens issued to applications (as well as session cookies in a user's browser), by resetting the **refreshTokensValidFromDateTime** user property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the user has a lost or stolen device.  This operation would prevent access to any of the organization's data accessed through applications on the device without the user first being required to sign in again. In fact, this operation would force the user to sign in again for all applications that they have previously consented to, independent of device. For developers, if the application attempts to redeem a delegated access token for this user by using an invalidated refresh token, the application will get an error. If this happens, the application will need to acquire a new refresh token by making a request to the authorize endpoint, which will force the user to sign in.
func (m *MicrosoftGraphInvalidateAllRefreshTokensRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphInvalidateAllRefreshTokensRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
