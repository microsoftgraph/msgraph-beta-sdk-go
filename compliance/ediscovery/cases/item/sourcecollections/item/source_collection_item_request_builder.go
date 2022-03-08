package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SourceCollectionItemRequestBuilderDeleteOptions options for Delete
type SourceCollectionItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SourceCollectionItemRequestBuilderGetOptions options for Get
type SourceCollectionItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SourceCollectionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SourceCollectionItemRequestBuilderGetQueryParameters returns a list of sourceCollection objects associated with this case.
type SourceCollectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SourceCollectionItemRequestBuilderPatchOptions options for Patch
type SourceCollectionItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollectionable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
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
        urlTplParams["dataSource_id"] = id
    }
    return id3fbc20e2aaca189e5a1a34399bb63266b91e90673fe015d4987e17f6eff56c3.NewDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SourceCollectionItemRequestBuilder) AddToReviewSetOperation()(*i3bd4cbfecafd766c2564c32940263a2048489063caa57ec31b0d0ad2df543409.AddToReviewSetOperationRequestBuilder) {
    return i3bd4cbfecafd766c2564c32940263a2048489063caa57ec31b0d0ad2df543409.NewAddToReviewSetOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSourceCollectionItemRequestBuilderInternal instantiates a new SourceCollectionItemRequestBuilder and sets the default values.
func NewSourceCollectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SourceCollectionItemRequestBuilder) {
    m := &SourceCollectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/sourceCollections/{sourceCollection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSourceCollectionItemRequestBuilder instantiates a new SourceCollectionItemRequestBuilder and sets the default values.
func NewSourceCollectionItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SourceCollectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSourceCollectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sourceCollections for compliance
func (m *SourceCollectionItemRequestBuilder) CreateDeleteRequestInformation(options *SourceCollectionItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation returns a list of sourceCollection objects associated with this case.
func (m *SourceCollectionItemRequestBuilder) CreateGetRequestInformation(options *SourceCollectionItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property sourceCollections in compliance
func (m *SourceCollectionItemRequestBuilder) CreatePatchRequestInformation(options *SourceCollectionItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
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
        urlTplParams["dataSource_id"] = id
    }
    return i68367f015eed592e4cb8c8ffbae9b14545f24cbd517163582e6f2b0ec13d383b.NewDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property sourceCollections for compliance
func (m *SourceCollectionItemRequestBuilder) Delete(options *SourceCollectionItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *SourceCollectionItemRequestBuilder) EstimateStatistics()(*i85c4be0475ccbd4aff4cb007fd0c2ec13881b05cf9e8cb5b59116c59ad86da86.EstimateStatisticsRequestBuilder) {
    return i85c4be0475ccbd4aff4cb007fd0c2ec13881b05cf9e8cb5b59116c59ad86da86.NewEstimateStatisticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get returns a list of sourceCollection objects associated with this case.
func (m *SourceCollectionItemRequestBuilder) Get(options *SourceCollectionItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateSourceCollectionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollectionable), nil
}
func (m *SourceCollectionItemRequestBuilder) LastEstimateStatisticsOperation()(*i3a1a6840838c80f69f0c58a7e3911451029954fdbf7475cb9efe0917ce53d7b3.LastEstimateStatisticsOperationRequestBuilder) {
    return i3a1a6840838c80f69f0c58a7e3911451029954fdbf7475cb9efe0917ce53d7b3.NewLastEstimateStatisticsOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["noncustodialDataSource_id"] = id
    }
    return idb9eb7d4d565925a68fe0c31a5803668300d9310f130d86958e02c95d5021202.NewNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property sourceCollections in compliance
func (m *SourceCollectionItemRequestBuilder) Patch(options *SourceCollectionItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *SourceCollectionItemRequestBuilder) PurgeData()(*i84c12d2b8523a8a3510e4a5f67b2e585d0ce763b1a2a1889cd8f30a1adc23b66.PurgeDataRequestBuilder) {
    return i84c12d2b8523a8a3510e4a5f67b2e585d0ce763b1a2a1889cd8f30a1adc23b66.NewPurgeDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
