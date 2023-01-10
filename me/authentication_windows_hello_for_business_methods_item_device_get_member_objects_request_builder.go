package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilder provides operations to call the getMemberObjects method.
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilderInternal instantiates a new GetMemberObjectsRequestBuilder and sets the default values.
func NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilder) {
    m := &AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/microsoft.graph.getMemberObjects";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilder instantiates a new GetMemberObjectsRequestBuilder and sets the default values.
func NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getMemberObjects
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilder) Post(ctx context.Context, body AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBodyable, requestConfiguration *AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilderPostRequestConfiguration)(AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateAuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsResponseable), nil
}
// ToPostRequestInformation invoke action getMemberObjects
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsPostRequestBodyable, requestConfiguration *AuthenticationWindowsHelloForBusinessMethodsItemDeviceGetMemberObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
