package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i04990ccdcd19a339f57b3ab646603ac886a33f74477ab8d057f344a6c0b4713d "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchdatastores/item/sessions/item/uploadagent"
    ib8b448373dd317e53420a2b8e592c4264bb8430dd86f1f8af1ff13cf0b07a30f "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchdatastores/item/sessions/item/renew"
    ic6bdbe766584845439738d0d52856be660bd81209b30e8e329715ddadc6fbfd4 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchdatastores/item/sessions/item/cancel"
    ie50d1404accfd9f1f2d095ca342c1d50a359ffc201c488ebc36605b1a52c8fe8 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchdatastores/item/sessions/item/commit"
)

// ExactMatchSessionRequestBuilder builds and executes requests for operations under \dataClassification\exactMatchDataStores\{exactMatchDataStore-id}\sessions\{exactMatchSession-id}
type ExactMatchSessionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ExactMatchSessionRequestBuilderDeleteOptions options for Delete
type ExactMatchSessionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExactMatchSessionRequestBuilderGetOptions options for Get
type ExactMatchSessionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ExactMatchSessionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExactMatchSessionRequestBuilderGetQueryParameters get sessions from dataClassification
type ExactMatchSessionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ExactMatchSessionRequestBuilderPatchOptions options for Patch
type ExactMatchSessionRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExactMatchSession;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ExactMatchSessionRequestBuilder) Cancel()(*ic6bdbe766584845439738d0d52856be660bd81209b30e8e329715ddadc6fbfd4.CancelRequestBuilder) {
    return ic6bdbe766584845439738d0d52856be660bd81209b30e8e329715ddadc6fbfd4.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ExactMatchSessionRequestBuilder) Commit()(*ie50d1404accfd9f1f2d095ca342c1d50a359ffc201c488ebc36605b1a52c8fe8.CommitRequestBuilder) {
    return ie50d1404accfd9f1f2d095ca342c1d50a359ffc201c488ebc36605b1a52c8fe8.NewCommitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewExactMatchSessionRequestBuilderInternal instantiates a new ExactMatchSessionRequestBuilder and sets the default values.
func NewExactMatchSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExactMatchSessionRequestBuilder) {
    m := &ExactMatchSessionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/dataClassification/exactMatchDataStores/{exactMatchDataStore_id}/sessions/{exactMatchSession_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExactMatchSessionRequestBuilder instantiates a new ExactMatchSessionRequestBuilder and sets the default values.
func NewExactMatchSessionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExactMatchSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExactMatchSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sessions for dataClassification
func (m *ExactMatchSessionRequestBuilder) CreateDeleteRequestInformation(options *ExactMatchSessionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get sessions from dataClassification
func (m *ExactMatchSessionRequestBuilder) CreateGetRequestInformation(options *ExactMatchSessionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sessions in dataClassification
func (m *ExactMatchSessionRequestBuilder) CreatePatchRequestInformation(options *ExactMatchSessionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property sessions for dataClassification
func (m *ExactMatchSessionRequestBuilder) Delete(options *ExactMatchSessionRequestBuilderDeleteOptions)(error) {
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
// Get get sessions from dataClassification
func (m *ExactMatchSessionRequestBuilder) Get(options *ExactMatchSessionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExactMatchSession, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewExactMatchSession() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExactMatchSession), nil
}
// Patch update the navigation property sessions in dataClassification
func (m *ExactMatchSessionRequestBuilder) Patch(options *ExactMatchSessionRequestBuilderPatchOptions)(error) {
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
func (m *ExactMatchSessionRequestBuilder) Renew()(*ib8b448373dd317e53420a2b8e592c4264bb8430dd86f1f8af1ff13cf0b07a30f.RenewRequestBuilder) {
    return ib8b448373dd317e53420a2b8e592c4264bb8430dd86f1f8af1ff13cf0b07a30f.NewRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ExactMatchSessionRequestBuilder) UploadAgent()(*i04990ccdcd19a339f57b3ab646603ac886a33f74477ab8d057f344a6c0b4713d.UploadAgentRequestBuilder) {
    return i04990ccdcd19a339f57b3ab646603ac886a33f74477ab8d057f344a6c0b4713d.NewUploadAgentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
