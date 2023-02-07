package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilder provides operations to call the cancelSecurityAction method.
type SecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilderInternal instantiates a new MicrosoftGraphCancelSecurityActionRequestBuilder and sets the default values.
func NewSecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilder) {
    m := &SecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/securityActions/{securityAction%2Did}/microsoft.graph.cancelSecurityAction";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilder instantiates a new MicrosoftGraphCancelSecurityActionRequestBuilder and sets the default values.
func NewSecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancel a security operation.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/securityaction-cancelsecurityaction?view=graph-rest-1.0
func (m *SecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilder) Post(ctx context.Context, requestConfiguration *SecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation cancel a security operation.
func (m *SecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SecurityActionsItemMicrosoftGraphCancelSecurityActionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
