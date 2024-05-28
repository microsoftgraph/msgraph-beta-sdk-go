package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder provides operations to call the setOrder method.
type B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderInternal instantiates a new B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder and sets the default values.
func NewB2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder) {
    m := &B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/userAttributeAssignments/setOrder", pathParameters),
    }
    return m
}
// NewB2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder instantiates a new B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder and sets the default values.
func NewB2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the order of identityUserFlowAttributeAssignments being collected within a user flow.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattributeassignment-setorder?view=graph-rest-beta
func (m *B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder) Post(ctx context.Context, body B2cuserflowsItemUserattributeassignmentsSetorderSetOrderPostRequestBodyable, requestConfiguration *B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation set the order of identityUserFlowAttributeAssignments being collected within a user flow.
// returns a *RequestInformation when successful
func (m *B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder) ToPostRequestInformation(ctx context.Context, body B2cuserflowsItemUserattributeassignmentsSetorderSetOrderPostRequestBodyable, requestConfiguration *B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder when successful
func (m *B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder) WithUrl(rawUrl string)(*B2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder) {
    return NewB2cuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
