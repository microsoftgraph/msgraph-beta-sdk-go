package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i7cca032aba7817ce573097d7ffcda7fe783e88b72b9694a811a3aea09de04886 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/journals/item/account"
    ic60d487aadcbb1c18a607e1ab79a5536b13c438e08f09a3b9d0d1a41c5a73a9f "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/journals/item/journallines"
    if8219da59c227d3b7810fa9438c16450ddb8cb36ddb34695d876359a43e349fc "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/journals/item/post"
    i22450797f858de15c310b162f7fd2d8ec26e3cdc4d1257d1de65a2cc0de71163 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/journals/item/journallines/item"
)

// JournalRequestBuilder builds and executes requests for operations under \financials\companies\{company-id}\journals\{journal-id}
type JournalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// JournalRequestBuilderDeleteOptions options for Delete
type JournalRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// JournalRequestBuilderGetOptions options for Get
type JournalRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *JournalRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// JournalRequestBuilderGetQueryParameters get journals from financials
type JournalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// JournalRequestBuilderPatchOptions options for Patch
type JournalRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Journal;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *JournalRequestBuilder) Account()(*i7cca032aba7817ce573097d7ffcda7fe783e88b72b9694a811a3aea09de04886.AccountRequestBuilder) {
    return i7cca032aba7817ce573097d7ffcda7fe783e88b72b9694a811a3aea09de04886.NewAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewJournalRequestBuilderInternal instantiates a new JournalRequestBuilder and sets the default values.
func NewJournalRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*JournalRequestBuilder) {
    m := &JournalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/journals/{journal_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewJournalRequestBuilder instantiates a new JournalRequestBuilder and sets the default values.
func NewJournalRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*JournalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewJournalRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property journals for financials
func (m *JournalRequestBuilder) CreateDeleteRequestInformation(options *JournalRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get journals from financials
func (m *JournalRequestBuilder) CreateGetRequestInformation(options *JournalRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property journals in financials
func (m *JournalRequestBuilder) CreatePatchRequestInformation(options *JournalRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property journals for financials
func (m *JournalRequestBuilder) Delete(options *JournalRequestBuilderDeleteOptions)(error) {
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
// Get get journals from financials
func (m *JournalRequestBuilder) Get(options *JournalRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Journal, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewJournal() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Journal), nil
}
func (m *JournalRequestBuilder) JournalLines()(*ic60d487aadcbb1c18a607e1ab79a5536b13c438e08f09a3b9d0d1a41c5a73a9f.JournalLinesRequestBuilder) {
    return ic60d487aadcbb1c18a607e1ab79a5536b13c438e08f09a3b9d0d1a41c5a73a9f.NewJournalLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JournalLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.journals.item.journalLines.item collection
func (m *JournalRequestBuilder) JournalLinesById(id string)(*i22450797f858de15c310b162f7fd2d8ec26e3cdc4d1257d1de65a2cc0de71163.JournalLineRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["journalLine_id"] = id
    }
    return i22450797f858de15c310b162f7fd2d8ec26e3cdc4d1257d1de65a2cc0de71163.NewJournalLineRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property journals in financials
func (m *JournalRequestBuilder) Patch(options *JournalRequestBuilderPatchOptions)(error) {
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
func (m *JournalRequestBuilder) Post()(*if8219da59c227d3b7810fa9438c16450ddb8cb36ddb34695d876359a43e349fc.PostRequestBuilder) {
    return if8219da59c227d3b7810fa9438c16450ddb8cb36ddb34695d876359a43e349fc.NewPostRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
