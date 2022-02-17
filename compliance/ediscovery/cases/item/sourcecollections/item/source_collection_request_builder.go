package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0597b69327ef47612760d3a06e5168d4357386c9da9233bf1fe1f95375bde771 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/custodiansources"
    i2c744ed9b76601b529ed40d577c487cc9b9cdba275f49afd6d038d45b275a597 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/additionalsources"
    i3a1a6840838c80f69f0c58a7e3911451029954fdbf7475cb9efe0917ce53d7b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/lastestimatestatisticsoperation"
    i3bd4cbfecafd766c2564c32940263a2048489063caa57ec31b0d0ad2df543409 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/addtoreviewsetoperation"
    i85c4be0475ccbd4aff4cb007fd0c2ec13881b05cf9e8cb5b59116c59ad86da86 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/estimatestatistics"
    ie844c3a4d7bb8cc79a58d398b260c44e966bd08e45f89594aa2bccaa7e8180e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/noncustodialsources"
    id3fbc20e2aaca189e5a1a34399bb63266b91e90673fe015d4987e17f6eff56c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item/additionalsources/item"
)

// SourceCollectionRequestBuilder builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\sourceCollections\{sourceCollection-id}
type SourceCollectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SourceCollectionRequestBuilderDeleteOptions options for Delete
type SourceCollectionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SourceCollectionRequestBuilderGetOptions options for Get
type SourceCollectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SourceCollectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SourceCollectionRequestBuilderGetQueryParameters returns a list of sourceCollection objects associated with this case.
type SourceCollectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SourceCollectionRequestBuilderPatchOptions options for Patch
type SourceCollectionRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SourceCollectionRequestBuilder) AdditionalSources()(*i2c744ed9b76601b529ed40d577c487cc9b9cdba275f49afd6d038d45b275a597.AdditionalSourcesRequestBuilder) {
    return i2c744ed9b76601b529ed40d577c487cc9b9cdba275f49afd6d038d45b275a597.NewAdditionalSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AdditionalSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.sourceCollections.item.additionalSources.item collection
func (m *SourceCollectionRequestBuilder) AdditionalSourcesById(id string)(*id3fbc20e2aaca189e5a1a34399bb63266b91e90673fe015d4987e17f6eff56c3.DataSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataSource_id"] = id
    }
    return id3fbc20e2aaca189e5a1a34399bb63266b91e90673fe015d4987e17f6eff56c3.NewDataSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SourceCollectionRequestBuilder) AddToReviewSetOperation()(*i3bd4cbfecafd766c2564c32940263a2048489063caa57ec31b0d0ad2df543409.AddToReviewSetOperationRequestBuilder) {
    return i3bd4cbfecafd766c2564c32940263a2048489063caa57ec31b0d0ad2df543409.NewAddToReviewSetOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSourceCollectionRequestBuilderInternal instantiates a new SourceCollectionRequestBuilder and sets the default values.
func NewSourceCollectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SourceCollectionRequestBuilder) {
    m := &SourceCollectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/sourceCollections/{sourceCollection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSourceCollectionRequestBuilder instantiates a new SourceCollectionRequestBuilder and sets the default values.
func NewSourceCollectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SourceCollectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSourceCollectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation returns a list of sourceCollection objects associated with this case.
func (m *SourceCollectionRequestBuilder) CreateDeleteRequestInformation(options *SourceCollectionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SourceCollectionRequestBuilder) CreateGetRequestInformation(options *SourceCollectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation returns a list of sourceCollection objects associated with this case.
func (m *SourceCollectionRequestBuilder) CreatePatchRequestInformation(options *SourceCollectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SourceCollectionRequestBuilder) CustodianSources()(*i0597b69327ef47612760d3a06e5168d4357386c9da9233bf1fe1f95375bde771.CustodianSourcesRequestBuilder) {
    return i0597b69327ef47612760d3a06e5168d4357386c9da9233bf1fe1f95375bde771.NewCustodianSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete returns a list of sourceCollection objects associated with this case.
func (m *SourceCollectionRequestBuilder) Delete(options *SourceCollectionRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *SourceCollectionRequestBuilder) EstimateStatistics()(*i85c4be0475ccbd4aff4cb007fd0c2ec13881b05cf9e8cb5b59116c59ad86da86.EstimateStatisticsRequestBuilder) {
    return i85c4be0475ccbd4aff4cb007fd0c2ec13881b05cf9e8cb5b59116c59ad86da86.NewEstimateStatisticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get returns a list of sourceCollection objects associated with this case.
func (m *SourceCollectionRequestBuilder) Get(options *SourceCollectionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSourceCollection() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollection), nil
}
func (m *SourceCollectionRequestBuilder) LastEstimateStatisticsOperation()(*i3a1a6840838c80f69f0c58a7e3911451029954fdbf7475cb9efe0917ce53d7b3.LastEstimateStatisticsOperationRequestBuilder) {
    return i3a1a6840838c80f69f0c58a7e3911451029954fdbf7475cb9efe0917ce53d7b3.NewLastEstimateStatisticsOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SourceCollectionRequestBuilder) NoncustodialSources()(*ie844c3a4d7bb8cc79a58d398b260c44e966bd08e45f89594aa2bccaa7e8180e9.NoncustodialSourcesRequestBuilder) {
    return ie844c3a4d7bb8cc79a58d398b260c44e966bd08e45f89594aa2bccaa7e8180e9.NewNoncustodialSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch returns a list of sourceCollection objects associated with this case.
func (m *SourceCollectionRequestBuilder) Patch(options *SourceCollectionRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
