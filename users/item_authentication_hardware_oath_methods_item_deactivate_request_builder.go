package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder provides operations to call the deactivate method.
type ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilderInternal instantiates a new ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder and sets the default values.
func NewItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder) {
    m := &ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/hardwareOathMethods/{hardwareOathAuthenticationMethod%2Did}/deactivate", pathParameters),
    }
    return m
}
// NewItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder instantiates a new ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder and sets the default values.
func NewItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post deactive a hardware OATH token. It remains assigned to a user.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/hardwareoathauthenticationmethod-deactivate?view=graph-rest-beta
func (m *ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation deactive a hardware OATH token. It remains assigned to a user.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder when successful
func (m *ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder) {
    return NewItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
