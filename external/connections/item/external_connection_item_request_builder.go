package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/externalconnectors"
    i37bbb931ba7a99580029f5e59fceec6eb8640b3b273e993b0da516dcc40deb00 "github.com/microsoftgraph/msgraph-beta-sdk-go/external/connections/item/operations"
    i65e452495f83460495ba06e038372fc125514330984811a66430c11e50b0a903 "github.com/microsoftgraph/msgraph-beta-sdk-go/external/connections/item/schema"
    i8c82bcb72ad581e1896086d0a0ba519177dc48e3fabef71c43f877aa89d0d5bb "github.com/microsoftgraph/msgraph-beta-sdk-go/external/connections/item/quota"
    ic21306f170b512d1bf604d2d8a2934279112fc43e76ee2e1190304bfcb9affaf "github.com/microsoftgraph/msgraph-beta-sdk-go/external/connections/item/items"
    iffd6530bca8d8e49e6ef2ef7466d12e5661125b7576d14dcf6d7effb88eecfab "github.com/microsoftgraph/msgraph-beta-sdk-go/external/connections/item/groups"
    i69c1e05ad8c29369be49685d62438c214dc9c439a6a4fbbe7ed3e2fe0eda2729 "github.com/microsoftgraph/msgraph-beta-sdk-go/external/connections/item/groups/item"
    ia3b063108d214ef75fea5b818cfe6b67e41a7d601f088df5d728f5455baccf43 "github.com/microsoftgraph/msgraph-beta-sdk-go/external/connections/item/items/item"
    ib6f02e4874285f867a5cd99416af186a892f455b055bf8cf850ece628119315b "github.com/microsoftgraph/msgraph-beta-sdk-go/external/connections/item/operations/item"
)

// ExternalConnectionItemRequestBuilder provides operations to manage the connections property of the microsoft.graph.externalConnectors.external entity.
type ExternalConnectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ExternalConnectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExternalConnectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExternalConnectionItemRequestBuilderGetQueryParameters get connections from external
type ExternalConnectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExternalConnectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExternalConnectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExternalConnectionItemRequestBuilderGetQueryParameters
}
// ExternalConnectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExternalConnectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExternalConnectionItemRequestBuilderInternal instantiates a new ExternalConnectionItemRequestBuilder and sets the default values.
func NewExternalConnectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExternalConnectionItemRequestBuilder) {
    m := &ExternalConnectionItemRequestBuilder{
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
// NewExternalConnectionItemRequestBuilder instantiates a new ExternalConnectionItemRequestBuilder and sets the default values.
func NewExternalConnectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExternalConnectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExternalConnectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property connections for external
func (m *ExternalConnectionItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property connections for external
func (m *ExternalConnectionItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ExternalConnectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ExternalConnectionItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get connections from external
func (m *ExternalConnectionItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ExternalConnectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ExternalConnectionItemRequestBuilder) CreatePatchRequestInformation(body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property connections in external
func (m *ExternalConnectionItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable, requestConfiguration *ExternalConnectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property connections for external
func (m *ExternalConnectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExternalConnectionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *ExternalConnectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExternalConnectionItemRequestBuilderGetRequestConfiguration)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Groups the groups property
func (m *ExternalConnectionItemRequestBuilder) Groups()(*iffd6530bca8d8e49e6ef2ef7466d12e5661125b7576d14dcf6d7effb88eecfab.GroupsRequestBuilder) {
    return iffd6530bca8d8e49e6ef2ef7466d12e5661125b7576d14dcf6d7effb88eecfab.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.external.connections.item.groups.item collection
func (m *ExternalConnectionItemRequestBuilder) GroupsById(id string)(*i69c1e05ad8c29369be49685d62438c214dc9c439a6a4fbbe7ed3e2fe0eda2729.ExternalGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalGroup%2Did"] = id
    }
    return i69c1e05ad8c29369be49685d62438c214dc9c439a6a4fbbe7ed3e2fe0eda2729.NewExternalGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Items the items property
func (m *ExternalConnectionItemRequestBuilder) Items()(*ic21306f170b512d1bf604d2d8a2934279112fc43e76ee2e1190304bfcb9affaf.ItemsRequestBuilder) {
    return ic21306f170b512d1bf604d2d8a2934279112fc43e76ee2e1190304bfcb9affaf.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.external.connections.item.items.item collection
func (m *ExternalConnectionItemRequestBuilder) ItemsById(id string)(*ia3b063108d214ef75fea5b818cfe6b67e41a7d601f088df5d728f5455baccf43.ExternalItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalItem%2Did"] = id
    }
    return ia3b063108d214ef75fea5b818cfe6b67e41a7d601f088df5d728f5455baccf43.NewExternalItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *ExternalConnectionItemRequestBuilder) Operations()(*i37bbb931ba7a99580029f5e59fceec6eb8640b3b273e993b0da516dcc40deb00.OperationsRequestBuilder) {
    return i37bbb931ba7a99580029f5e59fceec6eb8640b3b273e993b0da516dcc40deb00.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.external.connections.item.operations.item collection
func (m *ExternalConnectionItemRequestBuilder) OperationsById(id string)(*ib6f02e4874285f867a5cd99416af186a892f455b055bf8cf850ece628119315b.ConnectionOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectionOperation%2Did"] = id
    }
    return ib6f02e4874285f867a5cd99416af186a892f455b055bf8cf850ece628119315b.NewConnectionOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property connections in external
func (m *ExternalConnectionItemRequestBuilder) Patch(ctx context.Context, body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable, requestConfiguration *ExternalConnectionItemRequestBuilderPatchRequestConfiguration)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Quota the quota property
func (m *ExternalConnectionItemRequestBuilder) Quota()(*i8c82bcb72ad581e1896086d0a0ba519177dc48e3fabef71c43f877aa89d0d5bb.QuotaRequestBuilder) {
    return i8c82bcb72ad581e1896086d0a0ba519177dc48e3fabef71c43f877aa89d0d5bb.NewQuotaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schema the schema property
func (m *ExternalConnectionItemRequestBuilder) Schema()(*i65e452495f83460495ba06e038372fc125514330984811a66430c11e50b0a903.SchemaRequestBuilder) {
    return i65e452495f83460495ba06e038372fc125514330984811a66430c11e50b0a903.NewSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
