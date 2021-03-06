package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
    i0597b69327ef47612760d3a06e5168d4357386c9da9233bf1fe1f95375bde771 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/custodiansources"
    i2c744ed9b76601b529ed40d577c487cc9b9cdba275f49afd6d038d45b275a597 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/additionalsources"
    i3a1a6840838c80f69f0c58a7e3911451029954fdbf7475cb9efe0917ce53d7b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/lastestimatestatisticsoperation"
    i3bd4cbfecafd766c2564c32940263a2048489063caa57ec31b0d0ad2df543409 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/addtoreviewsetoperation"
    i84c12d2b8523a8a3510e4a5f67b2e585d0ce763b1a2a1889cd8f30a1adc23b66 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/purgedata"
    i85c4be0475ccbd4aff4cb007fd0c2ec13881b05cf9e8cb5b59116c59ad86da86 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/estimatestatistics"
    ie844c3a4d7bb8cc79a58d398b260c44e966bd08e45f89594aa2bccaa7e8180e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/noncustodialsources"
    i68367f015eed592e4cb8c8ffbae9b14545f24cbd517163582e6f2b0ec13d383b "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/custodiansources/item"
    id3fbc20e2aaca189e5a1a34399bb63266b91e90673fe015d4987e17f6eff56c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/additionalsources/item"
    idb9eb7d4d565925a68fe0c31a5803668300d9310f130d86958e02c95d5021202 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/noncustodialsources/item"
)

// SourceCollectionItemRequestBuilder provides operations to manage the sourceCollections property of the microsoft.graph.ediscovery.case entity.
type SourceCollectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SourceCollectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SourceCollectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SourceCollectionItemRequestBuilderGetQueryParameters returns a list of sourceCollection objects associated with this case.
type SourceCollectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SourceCollectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SourceCollectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SourceCollectionItemRequestBuilderGetQueryParameters
}
// SourceCollectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SourceCollectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdditionalSources the additionalSources property
func (m *SourceCollectionItemRequestBuilder) AdditionalSources()(*i2c744ed9b76601b529ed40d577c487cc9b9cdba275f49afd6d038d45b275a597.AdditionalSourcesRequestBuilder) {
    return i2c744ed9b76601b529ed40d577c487cc9b9cdba275f49afd6d038d45b275a597.NewAdditionalSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AdditionalSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.sourceCollections.item.additionalSources.item collection
func (m *SourceCollectionItemRequestBuilder) AdditionalSourcesById(id string)(*id3fbc20e2aaca189e5a1a34399bb63266b91e90673fe015d4987e17f6eff56c3.DataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataSource%2Did"] = id
    }
    return id3fbc20e2aaca189e5a1a34399bb63266b91e90673fe015d4987e17f6eff56c3.NewDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AddToReviewSetOperation the addToReviewSetOperation property
func (m *SourceCollectionItemRequestBuilder) AddToReviewSetOperation()(*i3bd4cbfecafd766c2564c32940263a2048489063caa57ec31b0d0ad2df543409.AddToReviewSetOperationRequestBuilder) {
    return i3bd4cbfecafd766c2564c32940263a2048489063caa57ec31b0d0ad2df543409.NewAddToReviewSetOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSourceCollectionItemRequestBuilderInternal instantiates a new SourceCollectionItemRequestBuilder and sets the default values.
func NewSourceCollectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SourceCollectionItemRequestBuilder) {
    m := &SourceCollectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/sourceCollections/{sourceCollection%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSourceCollectionItemRequestBuilder instantiates a new SourceCollectionItemRequestBuilder and sets the default values.
func NewSourceCollectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SourceCollectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSourceCollectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sourceCollections for compliance
func (m *SourceCollectionItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property sourceCollections for compliance
func (m *SourceCollectionItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *SourceCollectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation returns a list of sourceCollection objects associated with this case.
func (m *SourceCollectionItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration returns a list of sourceCollection objects associated with this case.
func (m *SourceCollectionItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SourceCollectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sourceCollections in compliance
func (m *SourceCollectionItemRequestBuilder) CreatePatchRequestInformation(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property sourceCollections in compliance
func (m *SourceCollectionItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, requestConfiguration *SourceCollectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CustodianSources the custodianSources property
func (m *SourceCollectionItemRequestBuilder) CustodianSources()(*i0597b69327ef47612760d3a06e5168d4357386c9da9233bf1fe1f95375bde771.CustodianSourcesRequestBuilder) {
    return i0597b69327ef47612760d3a06e5168d4357386c9da9233bf1fe1f95375bde771.NewCustodianSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustodianSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.sourceCollections.item.custodianSources.item collection
func (m *SourceCollectionItemRequestBuilder) CustodianSourcesById(id string)(*i68367f015eed592e4cb8c8ffbae9b14545f24cbd517163582e6f2b0ec13d383b.DataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataSource%2Did"] = id
    }
    return i68367f015eed592e4cb8c8ffbae9b14545f24cbd517163582e6f2b0ec13d383b.NewDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property sourceCollections for compliance
func (m *SourceCollectionItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property sourceCollections for compliance
func (m *SourceCollectionItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *SourceCollectionItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// EstimateStatistics the estimateStatistics property
func (m *SourceCollectionItemRequestBuilder) EstimateStatistics()(*i85c4be0475ccbd4aff4cb007fd0c2ec13881b05cf9e8cb5b59116c59ad86da86.EstimateStatisticsRequestBuilder) {
    return i85c4be0475ccbd4aff4cb007fd0c2ec13881b05cf9e8cb5b59116c59ad86da86.NewEstimateStatisticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get returns a list of sourceCollection objects associated with this case.
func (m *SourceCollectionItemRequestBuilder) Get()(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler returns a list of sourceCollection objects associated with this case.
func (m *SourceCollectionItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *SourceCollectionItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateSourceCollectionFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable), nil
}
// LastEstimateStatisticsOperation the lastEstimateStatisticsOperation property
func (m *SourceCollectionItemRequestBuilder) LastEstimateStatisticsOperation()(*i3a1a6840838c80f69f0c58a7e3911451029954fdbf7475cb9efe0917ce53d7b3.LastEstimateStatisticsOperationRequestBuilder) {
    return i3a1a6840838c80f69f0c58a7e3911451029954fdbf7475cb9efe0917ce53d7b3.NewLastEstimateStatisticsOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NoncustodialSources the noncustodialSources property
func (m *SourceCollectionItemRequestBuilder) NoncustodialSources()(*ie844c3a4d7bb8cc79a58d398b260c44e966bd08e45f89594aa2bccaa7e8180e9.NoncustodialSourcesRequestBuilder) {
    return ie844c3a4d7bb8cc79a58d398b260c44e966bd08e45f89594aa2bccaa7e8180e9.NewNoncustodialSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NoncustodialSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.sourceCollections.item.noncustodialSources.item collection
func (m *SourceCollectionItemRequestBuilder) NoncustodialSourcesById(id string)(*idb9eb7d4d565925a68fe0c31a5803668300d9310f130d86958e02c95d5021202.NoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["noncustodialDataSource%2Did"] = id
    }
    return idb9eb7d4d565925a68fe0c31a5803668300d9310f130d86958e02c95d5021202.NewNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property sourceCollections in compliance
func (m *SourceCollectionItemRequestBuilder) Patch(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property sourceCollections in compliance
func (m *SourceCollectionItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.SourceCollectionable, requestConfiguration *SourceCollectionItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// PurgeData the purgeData property
func (m *SourceCollectionItemRequestBuilder) PurgeData()(*i84c12d2b8523a8a3510e4a5f67b2e585d0ce763b1a2a1889cd8f30a1adc23b66.PurgeDataRequestBuilder) {
    return i84c12d2b8523a8a3510e4a5f67b2e585d0ce763b1a2a1889cd8f30a1adc23b66.NewPurgeDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
