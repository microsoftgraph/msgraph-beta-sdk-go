package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilder provides operations to call the setOrder method.
type B2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// B2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilderInternal instantiates a new MicrosoftGraphSetOrderRequestBuilder and sets the default values.
func NewB2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilder) {
    m := &B2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/userAttributeAssignments/microsoft.graph.setOrder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewB2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilder instantiates a new MicrosoftGraphSetOrderRequestBuilder and sets the default values.
func NewB2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the order of identityUserFlowAttributeAssignments being collected within a user flow.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/identityuserflowattributeassignment-setorder?view=graph-rest-1.0
func (m *B2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilder) Post(ctx context.Context, body B2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBodyable, requestConfiguration *B2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation set the order of identityUserFlowAttributeAssignments being collected within a user flow.
func (m *B2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilder) ToPostRequestInformation(ctx context.Context, body B2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBodyable, requestConfiguration *B2cUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
