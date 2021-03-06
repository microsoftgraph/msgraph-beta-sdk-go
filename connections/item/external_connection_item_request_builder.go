package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/externalconnectors"
    i16aec0d979679a78de026ea357af647676e9bd2f7c8534da4da2643e5d6f81b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/operations"
    ia229bb0ea4c647efe1ef7f3f3d49c1c41e13acf3b845ba20fa1cc1c6f2730bdd "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/schema"
    ia8a9fb046e54d24538d7c6d25cc69813af40692eee89a9d2eb29dc5570b69c1d "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/quota"
    ic4ecba2a36e9429147361e0927b616e9b14c0047f0c81f4d383857826bf622b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/groups"
    ieb276fe54a3917eca6ba7222449ff2e6a5642f1c6bebeaabc0427903311a92df "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/items"
    i20ab49c75285e4cd8a2b4e4645b09f58717fe4efdaf039c7973589d395d6d3a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/groups/item"
    i38a4dbb9e776d366f238688a9031f7f382abfd310dedd2207973525e4859aebd "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/operations/item"
    ifb639436ebf63dc6cfd09347d97c3adc1ba4e054cb845ef13d79cea6a7aece70 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/items/item"
)

// ExternalConnectionItemRequestBuilder provides operations to manage the collection of externalConnection entities.
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
// ExternalConnectionItemRequestBuilderGetQueryParameters get entity from connections by key
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
    m.urlTemplate = "{+baseurl}/connections/{externalConnection%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete entity from connections
func (m *ExternalConnectionItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete entity from connections
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
// CreateGetRequestInformation get entity from connections by key
func (m *ExternalConnectionItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get entity from connections by key
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
// CreatePatchRequestInformation update entity in connections
func (m *ExternalConnectionItemRequestBuilder) CreatePatchRequestInformation(body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update entity in connections
func (m *ExternalConnectionItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable, requestConfiguration *ExternalConnectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete entity from connections
func (m *ExternalConnectionItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete entity from connections
func (m *ExternalConnectionItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ExternalConnectionItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from connections by key
func (m *ExternalConnectionItemRequestBuilder) Get()(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get entity from connections by key
func (m *ExternalConnectionItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ExternalConnectionItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.CreateExternalConnectionFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable), nil
}
// Groups the groups property
func (m *ExternalConnectionItemRequestBuilder) Groups()(*ic4ecba2a36e9429147361e0927b616e9b14c0047f0c81f4d383857826bf622b0.GroupsRequestBuilder) {
    return ic4ecba2a36e9429147361e0927b616e9b14c0047f0c81f4d383857826bf622b0.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.connections.item.groups.item collection
func (m *ExternalConnectionItemRequestBuilder) GroupsById(id string)(*i20ab49c75285e4cd8a2b4e4645b09f58717fe4efdaf039c7973589d395d6d3a7.ExternalGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalGroup%2Did"] = id
    }
    return i20ab49c75285e4cd8a2b4e4645b09f58717fe4efdaf039c7973589d395d6d3a7.NewExternalGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Items the items property
func (m *ExternalConnectionItemRequestBuilder) Items()(*ieb276fe54a3917eca6ba7222449ff2e6a5642f1c6bebeaabc0427903311a92df.ItemsRequestBuilder) {
    return ieb276fe54a3917eca6ba7222449ff2e6a5642f1c6bebeaabc0427903311a92df.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.connections.item.items.item collection
func (m *ExternalConnectionItemRequestBuilder) ItemsById(id string)(*ifb639436ebf63dc6cfd09347d97c3adc1ba4e054cb845ef13d79cea6a7aece70.ExternalItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalItem%2Did"] = id
    }
    return ifb639436ebf63dc6cfd09347d97c3adc1ba4e054cb845ef13d79cea6a7aece70.NewExternalItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *ExternalConnectionItemRequestBuilder) Operations()(*i16aec0d979679a78de026ea357af647676e9bd2f7c8534da4da2643e5d6f81b1.OperationsRequestBuilder) {
    return i16aec0d979679a78de026ea357af647676e9bd2f7c8534da4da2643e5d6f81b1.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.connections.item.operations.item collection
func (m *ExternalConnectionItemRequestBuilder) OperationsById(id string)(*i38a4dbb9e776d366f238688a9031f7f382abfd310dedd2207973525e4859aebd.ConnectionOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectionOperation%2Did"] = id
    }
    return i38a4dbb9e776d366f238688a9031f7f382abfd310dedd2207973525e4859aebd.NewConnectionOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in connections
func (m *ExternalConnectionItemRequestBuilder) Patch(body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update entity in connections
func (m *ExternalConnectionItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalConnectionable, requestConfiguration *ExternalConnectionItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Quota the quota property
func (m *ExternalConnectionItemRequestBuilder) Quota()(*ia8a9fb046e54d24538d7c6d25cc69813af40692eee89a9d2eb29dc5570b69c1d.QuotaRequestBuilder) {
    return ia8a9fb046e54d24538d7c6d25cc69813af40692eee89a9d2eb29dc5570b69c1d.NewQuotaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schema the schema property
func (m *ExternalConnectionItemRequestBuilder) Schema()(*ia229bb0ea4c647efe1ef7f3f3d49c1c41e13acf3b845ba20fa1cc1c6f2730bdd.SchemaRequestBuilder) {
    return ia229bb0ea4c647efe1ef7f3f3d49c1c41e13acf3b845ba20fa1cc1c6f2730bdd.NewSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
