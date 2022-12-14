package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilder provides operations to call the getMemberObjects method.
type AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilderInternal instantiates a new GetMemberObjectsRequestBuilder and sets the default values.
func NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilder) {
    m := &AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/microsoft.graph.getMemberObjects";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilder instantiates a new GetMemberObjectsRequestBuilder and sets the default values.
func NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getMemberObjects
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBodyable, requestConfiguration *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action getMemberObjects
func (m *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilder) Post(ctx context.Context, body AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsPostRequestBodyable, requestConfiguration *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsRequestBuilderPostRequestConfiguration)(AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberObjectsResponseable), nil
}
