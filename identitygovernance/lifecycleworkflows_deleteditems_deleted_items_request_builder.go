package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder provides operations to manage the deletedItems property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
type LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderGetQueryParameters deleted workflows in your lifecycle workflows instance.
type LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderInternal instantiates a new LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder) {
    m := &LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder instantiates a new LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deletedItems for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder) Delete(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get deleted workflows in your lifecycle workflows instance.
// returns a DeletedItemContainerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeletedItemContainerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeletedItemContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeletedItemContainerable), nil
}
// ToDeleteRequestInformation delete navigation property deletedItems for identityGovernance
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation deleted workflows in your lifecycle workflows instance.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Workflows provides operations to manage the workflows property of the microsoft.graph.deletedItemContainer entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder) Workflows()(*LifecycleworkflowsDeleteditemsWorkflowsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
