package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder provides operations to call the getOrder method.
type B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilderInternal instantiates a new B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder and sets the default values.
func NewB2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder) {
    m := &B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/userAttributeAssignments/getOrder()", pathParameters),
    }
    return m
}
// NewB2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder instantiates a new B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder and sets the default values.
func NewB2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the order of identityUserFlowAttributeAssignments being collected within a user flow.
// returns a AssignmentOrderable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattributeassignment-getorder?view=graph-rest-beta
func (m *B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder) Get(ctx context.Context, requestConfiguration *B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAssignmentOrderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable), nil
}
// ToGetRequestInformation get the order of identityUserFlowAttributeAssignments being collected within a user flow.
// returns a *RequestInformation when successful
func (m *B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder when successful
func (m *B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder) WithUrl(rawUrl string)(*B2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder) {
    return NewB2cuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
