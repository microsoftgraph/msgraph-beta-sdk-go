package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder provides operations to call the setOrder method.
type B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderInternal instantiates a new B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder and sets the default values.
func NewB2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder) {
    m := &B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/userAttributeAssignments/setOrder", pathParameters),
    }
    return m
}
// NewB2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder instantiates a new B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder and sets the default values.
func NewB2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the order of identityUserFlowAttributeAssignments being collected within a user flow.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattributeassignment-setorder?view=graph-rest-beta
func (m *B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder) Post(ctx context.Context, body B2xuserflowsItemUserattributeassignmentsSetorderSetOrderPostRequestBodyable, requestConfiguration *B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderPostRequestConfiguration)(error) {
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
func (m *B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder) ToPostRequestInformation(ctx context.Context, body B2xuserflowsItemUserattributeassignmentsSetorderSetOrderPostRequestBodyable, requestConfiguration *B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder when successful
func (m *B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder) WithUrl(rawUrl string)(*B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder) {
    return NewB2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
