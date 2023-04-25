package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder provides operations to call the setOrder method.
type B2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderInternal instantiates a new SetOrderRequestBuilder and sets the default values.
func NewB2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder) {
    m := &B2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2cUserFlows/{b2cIdentityUserFlow%2Did}/userAttributeAssignments/setOrder", pathParameters),
    }
    return m
}
// NewB2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder instantiates a new SetOrderRequestBuilder and sets the default values.
func NewB2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setOrder
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder) Post(ctx context.Context, body B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBodyable, requestConfiguration *B2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action setOrder
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilder) ToPostRequestInformation(ctx context.Context, body B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBodyable, requestConfiguration *B2cUserFlowsItemUserAttributeAssignmentsSetOrderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
