package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i29529418372c31fe11888bc42b9e925b44a4f44a558e637673b91be3e193f02d "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchdatastores/item/lookup"
    i2c50a6af5b17b681669f493c4dd54cac6a789f62e7a1e7e75c246aa3738d6530 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchdatastores/item/sessions"
    i1dcb0390b62e6994f2eedf162f9337be1ed916a12c0cd74064045e244b407892 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchdatastores/item/sessions/item"
)

// ExactMatchDataStoreRequestBuilder builds and executes requests for operations under \dataClassification\exactMatchDataStores\{exactMatchDataStore-id}
type ExactMatchDataStoreRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ExactMatchDataStoreRequestBuilderDeleteOptions options for Delete
type ExactMatchDataStoreRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExactMatchDataStoreRequestBuilderGetOptions options for Get
type ExactMatchDataStoreRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ExactMatchDataStoreRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExactMatchDataStoreRequestBuilderGetQueryParameters get exactMatchDataStores from dataClassification
type ExactMatchDataStoreRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ExactMatchDataStoreRequestBuilderPatchOptions options for Patch
type ExactMatchDataStoreRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExactMatchDataStore;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewExactMatchDataStoreRequestBuilderInternal instantiates a new ExactMatchDataStoreRequestBuilder and sets the default values.
func NewExactMatchDataStoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExactMatchDataStoreRequestBuilder) {
    m := &ExactMatchDataStoreRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExactMatchDataStoreRequestBuilder instantiates a new ExactMatchDataStoreRequestBuilder and sets the default values.
func NewExactMatchDataStoreRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExactMatchDataStoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExactMatchDataStoreRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property exactMatchDataStores for dataClassification
func (m *ExactMatchDataStoreRequestBuilder) CreateDeleteRequestInformation(options *ExactMatchDataStoreRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get exactMatchDataStores from dataClassification
func (m *ExactMatchDataStoreRequestBuilder) CreateGetRequestInformation(options *ExactMatchDataStoreRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property exactMatchDataStores in dataClassification
func (m *ExactMatchDataStoreRequestBuilder) CreatePatchRequestInformation(options *ExactMatchDataStoreRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property exactMatchDataStores for dataClassification
func (m *ExactMatchDataStoreRequestBuilder) Delete(options *ExactMatchDataStoreRequestBuilderDeleteOptions)(error) {
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
// Get get exactMatchDataStores from dataClassification
func (m *ExactMatchDataStoreRequestBuilder) Get(options *ExactMatchDataStoreRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExactMatchDataStore, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewExactMatchDataStore() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExactMatchDataStore), nil
}
func (m *ExactMatchDataStoreRequestBuilder) Lookup()(*i29529418372c31fe11888bc42b9e925b44a4f44a558e637673b91be3e193f02d.LookupRequestBuilder) {
    return i29529418372c31fe11888bc42b9e925b44a4f44a558e637673b91be3e193f02d.NewLookupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property exactMatchDataStores in dataClassification
func (m *ExactMatchDataStoreRequestBuilder) Patch(options *ExactMatchDataStoreRequestBuilderPatchOptions)(error) {
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
func (m *ExactMatchDataStoreRequestBuilder) Sessions()(*i2c50a6af5b17b681669f493c4dd54cac6a789f62e7a1e7e75c246aa3738d6530.SessionsRequestBuilder) {
    return i2c50a6af5b17b681669f493c4dd54cac6a789f62e7a1e7e75c246aa3738d6530.NewSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SessionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.dataClassification.exactMatchDataStores.item.sessions.item collection
func (m *ExactMatchDataStoreRequestBuilder) SessionsById(id string)(*i1dcb0390b62e6994f2eedf162f9337be1ed916a12c0cd74064045e244b407892.ExactMatchSessionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["exactMatchSession_id"] = id
    }
    return i1dcb0390b62e6994f2eedf162f9337be1ed916a12c0cd74064045e244b407892.NewExactMatchSessionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
