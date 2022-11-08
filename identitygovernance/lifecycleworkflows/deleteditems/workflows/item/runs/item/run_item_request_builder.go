package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
    i36650ffab2f715179aa89bbac2ecb39a161291d8e05211b444272567779eba02 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/runs/item/userprocessingresults"
    if44b034c7603f2c8d8627f9a3705d5a8013bc24940dbd491f63b6bf2ce7f39d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/runs/item/taskprocessingresults"
    i63a0705ecd39d5be87c44dce4a7a01d9cae21b55c0c17f321d406fcc802f9cd5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/runs/item/userprocessingresults/item"
    i73be248b707e6c0c2e221897e1d6a3e6e9f381cfe607563d3fe05def99298abe "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/lifecycleworkflows/deleteditems/workflows/item/runs/item/taskprocessingresults/item"
)

// RunItemRequestBuilder provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
type RunItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RunItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RunItemRequestBuilderGetQueryParameters get runs from identityGovernance
type RunItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RunItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RunItemRequestBuilderGetQueryParameters
}
// RunItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RunItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRunItemRequestBuilderInternal instantiates a new RunItemRequestBuilder and sets the default values.
func NewRunItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemRequestBuilder) {
    m := &RunItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/runs/{run%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRunItemRequestBuilder instantiates a new RunItemRequestBuilder and sets the default values.
func NewRunItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RunItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRunItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property runs for identityGovernance
func (m *RunItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *RunItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get runs from identityGovernance
func (m *RunItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *RunItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property runs in identityGovernance
func (m *RunItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Runable, requestConfiguration *RunItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property runs for identityGovernance
func (m *RunItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RunItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get runs from identityGovernance
func (m *RunItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RunItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Runable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateRunFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Runable), nil
}
// Patch update the navigation property runs in identityGovernance
func (m *RunItemRequestBuilder) Patch(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Runable, requestConfiguration *RunItemRequestBuilderPatchRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Runable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateRunFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Runable), nil
}
// TaskProcessingResults provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.run entity.
func (m *RunItemRequestBuilder) TaskProcessingResults()(*if44b034c7603f2c8d8627f9a3705d5a8013bc24940dbd491f63b6bf2ce7f39d3.TaskProcessingResultsRequestBuilder) {
    return if44b034c7603f2c8d8627f9a3705d5a8013bc24940dbd491f63b6bf2ce7f39d3.NewTaskProcessingResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskProcessingResultsById provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.run entity.
func (m *RunItemRequestBuilder) TaskProcessingResultsById(id string)(*i73be248b707e6c0c2e221897e1d6a3e6e9f381cfe607563d3fe05def99298abe.TaskProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taskProcessingResult%2Did"] = id
    }
    return i73be248b707e6c0c2e221897e1d6a3e6e9f381cfe607563d3fe05def99298abe.NewTaskProcessingResultItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserProcessingResults provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.run entity.
func (m *RunItemRequestBuilder) UserProcessingResults()(*i36650ffab2f715179aa89bbac2ecb39a161291d8e05211b444272567779eba02.UserProcessingResultsRequestBuilder) {
    return i36650ffab2f715179aa89bbac2ecb39a161291d8e05211b444272567779eba02.NewUserProcessingResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserProcessingResultsById provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.run entity.
func (m *RunItemRequestBuilder) UserProcessingResultsById(id string)(*i63a0705ecd39d5be87c44dce4a7a01d9cae21b55c0c17f321d406fcc802f9cd5.UserProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userProcessingResult%2Did"] = id
    }
    return i63a0705ecd39d5be87c44dce4a7a01d9cae21b55c0c17f321d406fcc802f9cd5.NewUserProcessingResultItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
