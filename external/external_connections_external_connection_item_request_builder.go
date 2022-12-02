package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/externalconnectors"
)

// ExternalConnectionsExternalConnectionItemRequestBuilder provides operations to manage the connections property of the microsoft.graph.externalConnectors.external entity.
type ExternalConnectionsExternalConnectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ExternalConnectionsExternalConnectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExternalConnectionsExternalConnectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExternalConnectionsExternalConnectionItemRequestBuilderGetQueryParameters get connections from external
type ExternalConnectionsExternalConnectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExternalConnectionsExternalConnectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExternalConnectionsExternalConnectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExternalConnectionsExternalConnectionItemRequestBuilderGetQueryParameters
}
// ExternalConnectionsExternalConnectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExternalConnectionsExternalConnectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExternalConnectionsExternalConnectionItemRequestBuilderInternal instantiates a new ExternalConnectionItemRequestBuilder and sets the default values.
func NewExternalConnectionsExternalConnectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExternalConnectionsExternalConnectionItemRequestBuilder) {
    m := &ExternalConnectionsExternalConnectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/external/connections/{externalConnection%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExternalConnectionsExternalConnectionItemRequestBuilder instantiates a new ExternalConnectionItemRequestBuilder and sets the default values.
func NewExternalConnectionsExternalConnectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExternalConnectionsExternalConnectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExternalConnectionsExternalConnectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property connections for external
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ExternalConnectionsExternalConnectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get connections from external
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ExternalConnectionsExternalConnectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property connections in external
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable, requestConfiguration *ExternalConnectionsExternalConnectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property connections for external
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExternalConnectionsExternalConnectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get connections from external
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExternalConnectionsExternalConnectionItemRequestBuilderGetRequestConfiguration)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.CreateExternalConnectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable), nil
}
// Groups provides operations to manage the groups property of the microsoft.graph.externalConnectors.externalConnection entity.
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) Groups()(*ExternalConnectionsItemGroupsRequestBuilder) {
    return NewExternalConnectionsItemGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupsById provides operations to manage the groups property of the microsoft.graph.externalConnectors.externalConnection entity.
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) GroupsById(id string)(*ExternalConnectionsItemGroupsExternalGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalGroup%2Did"] = id
    }
    return NewExternalConnectionsItemGroupsExternalGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Items provides operations to manage the items property of the microsoft.graph.externalConnectors.externalConnection entity.
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) Items()(*ExternalConnectionsItemItemsRequestBuilder) {
    return NewExternalConnectionsItemItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById provides operations to manage the items property of the microsoft.graph.externalConnectors.externalConnection entity.
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) ItemsById(id string)(*ExternalConnectionsItemItemsExternalItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalItem%2Did"] = id
    }
    return NewExternalConnectionsItemItemsExternalItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.externalConnectors.externalConnection entity.
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) Operations()(*ExternalConnectionsItemOperationsRequestBuilder) {
    return NewExternalConnectionsItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.externalConnectors.externalConnection entity.
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) OperationsById(id string)(*ExternalConnectionsItemOperationsConnectionOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectionOperation%2Did"] = id
    }
    return NewExternalConnectionsItemOperationsConnectionOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property connections in external
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) Patch(ctx context.Context, body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable, requestConfiguration *ExternalConnectionsExternalConnectionItemRequestBuilderPatchRequestConfiguration)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.CreateExternalConnectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable), nil
}
// Quota provides operations to manage the quota property of the microsoft.graph.externalConnectors.externalConnection entity.
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) Quota()(*ExternalConnectionsItemQuotaRequestBuilder) {
    return NewExternalConnectionsItemQuotaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schema provides operations to manage the schema property of the microsoft.graph.externalConnectors.externalConnection entity.
func (m *ExternalConnectionsExternalConnectionItemRequestBuilder) Schema()(*ExternalConnectionsItemSchemaRequestBuilder) {
    return NewExternalConnectionsItemSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
