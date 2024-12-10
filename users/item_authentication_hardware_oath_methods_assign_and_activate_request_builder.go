package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder provides operations to call the assignAndActivate method.
type ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilderInternal instantiates a new ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder and sets the default values.
func NewItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder) {
    m := &ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/hardwareOathMethods/assignAndActivate", pathParameters),
    }
    return m
}
// NewItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder instantiates a new ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder and sets the default values.
func NewItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post assign and activate a hardware token at the same time. This operation requires the device ID to activate it.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/hardwareoathauthenticationmethod-assignandactivate?view=graph-rest-beta
func (m *ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder) Post(ctx context.Context, body ItemAuthenticationHardwareOathMethodsAssignAndActivatePostRequestBodyable, requestConfiguration *ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation assign and activate a hardware token at the same time. This operation requires the device ID to activate it.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAuthenticationHardwareOathMethodsAssignAndActivatePostRequestBodyable, requestConfiguration *ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder when successful
func (m *ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder) {
    return NewItemAuthenticationHardwareOathMethodsAssignAndActivateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
