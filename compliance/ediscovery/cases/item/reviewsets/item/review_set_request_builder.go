package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i08b6f0b5a320f941f6b9e04f436d2deaf397e1da914e9bfb4309c34077c320f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reviewsets/item/addtoreviewset"
    i331d9fcb9fb3723ce0d6251992ccb1b62085e60afae20de233674ebeb7e41db8 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reviewsets/item/export"
    i95d988bf764f142af2cbc9b71e126590844a8f1556f57902f7965dddff422bb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reviewsets/item/queries"
    iddd703417744884f23660dc02635324f3f6311317b830c6b4be8a93262f68548 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reviewsets/item/queries/item"
)

// ReviewSetRequestBuilder builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\reviewSets\{reviewSet-id}
type ReviewSetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ReviewSetRequestBuilderDeleteOptions options for Delete
type ReviewSetRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ReviewSetRequestBuilderGetOptions options for Get
type ReviewSetRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ReviewSetRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ReviewSetRequestBuilderGetQueryParameters returns a list of reviewSet objects in the case. Read-only. Nullable.
type ReviewSetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ReviewSetRequestBuilderPatchOptions options for Patch
type ReviewSetRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ReviewSet;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ReviewSetRequestBuilder) AddToReviewSet()(*i08b6f0b5a320f941f6b9e04f436d2deaf397e1da914e9bfb4309c34077c320f6.AddToReviewSetRequestBuilder) {
    return i08b6f0b5a320f941f6b9e04f436d2deaf397e1da914e9bfb4309c34077c320f6.NewAddToReviewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewReviewSetRequestBuilderInternal instantiates a new ReviewSetRequestBuilder and sets the default values.
func NewReviewSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReviewSetRequestBuilder) {
    m := &ReviewSetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/reviewSets/{reviewSet_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReviewSetRequestBuilder instantiates a new ReviewSetRequestBuilder and sets the default values.
func NewReviewSetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReviewSetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReviewSetRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *ReviewSetRequestBuilder) CreateDeleteRequestInformation(options *ReviewSetRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *ReviewSetRequestBuilder) CreateGetRequestInformation(options *ReviewSetRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *ReviewSetRequestBuilder) CreatePatchRequestInformation(options *ReviewSetRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *ReviewSetRequestBuilder) Delete(options *ReviewSetRequestBuilderDeleteOptions)(error) {
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
func (m *ReviewSetRequestBuilder) Export()(*i331d9fcb9fb3723ce0d6251992ccb1b62085e60afae20de233674ebeb7e41db8.ExportRequestBuilder) {
    return i331d9fcb9fb3723ce0d6251992ccb1b62085e60afae20de233674ebeb7e41db8.NewExportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *ReviewSetRequestBuilder) Get(options *ReviewSetRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ReviewSet, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewReviewSet() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ReviewSet), nil
}
// Patch returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *ReviewSetRequestBuilder) Patch(options *ReviewSetRequestBuilderPatchOptions)(error) {
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
func (m *ReviewSetRequestBuilder) Queries()(*i95d988bf764f142af2cbc9b71e126590844a8f1556f57902f7965dddff422bb5.QueriesRequestBuilder) {
    return i95d988bf764f142af2cbc9b71e126590844a8f1556f57902f7965dddff422bb5.NewQueriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// QueriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.reviewSets.item.queries.item collection
func (m *ReviewSetRequestBuilder) QueriesById(id string)(*iddd703417744884f23660dc02635324f3f6311317b830c6b4be8a93262f68548.ReviewSetQueryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["reviewSetQuery_id"] = id
    }
    return iddd703417744884f23660dc02635324f3f6311317b830c6b4be8a93262f68548.NewReviewSetQueryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
