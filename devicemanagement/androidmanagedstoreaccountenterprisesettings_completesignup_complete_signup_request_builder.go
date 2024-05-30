package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder provides operations to call the completeSignup method.
type AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilderInternal instantiates a new AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder and sets the default values.
func NewAndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder) {
    m := &AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings/completeSignup", pathParameters),
    }
    return m
}
// NewAndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder instantiates a new AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder and sets the default values.
func NewAndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action completeSignup
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder) Post(ctx context.Context, body AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupPostRequestBodyable, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action completeSignup
// returns a *RequestInformation when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder) ToPostRequestInformation(ctx context.Context, body AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupPostRequestBodyable, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder) WithUrl(rawUrl string)(*AndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder) {
    return NewAndroidmanagedstoreaccountenterprisesettingsCompletesignupCompleteSignupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
