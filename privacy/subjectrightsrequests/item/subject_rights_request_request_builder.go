package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i06e33bb086cbb5f9333cc1c8bb09529f00623800d8f148e19c13fab399be0e21 "github.com/microsoftgraph/msgraph-beta-sdk-go/privacy/subjectrightsrequests/item/notes"
    i2bf59a4d763ab3355a4526b9ff2ddcae384f1210dd464ff5940ff3a054d5f26e "github.com/microsoftgraph/msgraph-beta-sdk-go/privacy/subjectrightsrequests/item/getfinalattachment"
    i958489264196dd08f36dd719f907da913cf1a6b619228f3d6bc8da52997d8449 "github.com/microsoftgraph/msgraph-beta-sdk-go/privacy/subjectrightsrequests/item/getfinalreport"
    ia790ce0fe221c2e579ab27030d34081e5415f38f47af2263081eb369b2561ccd "github.com/microsoftgraph/msgraph-beta-sdk-go/privacy/subjectrightsrequests/item/team"
    ibe7a5a69eee3e39fd9fd52c318c683abd357fbe438629376f053fbf8dafee6d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/privacy/subjectrightsrequests/item/notes/item"
)

// SubjectRightsRequestRequestBuilder builds and executes requests for operations under \privacy\subjectRightsRequests\{subjectRightsRequest-id}
type SubjectRightsRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SubjectRightsRequestRequestBuilderDeleteOptions options for Delete
type SubjectRightsRequestRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SubjectRightsRequestRequestBuilderGetOptions options for Get
type SubjectRightsRequestRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SubjectRightsRequestRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SubjectRightsRequestRequestBuilderGetQueryParameters get subjectRightsRequests from privacy
type SubjectRightsRequestRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SubjectRightsRequestRequestBuilderPatchOptions options for Patch
type SubjectRightsRequestRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SubjectRightsRequest;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSubjectRightsRequestRequestBuilderInternal instantiates a new SubjectRightsRequestRequestBuilder and sets the default values.
func NewSubjectRightsRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SubjectRightsRequestRequestBuilder) {
    m := &SubjectRightsRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privacy/subjectRightsRequests/{subjectRightsRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSubjectRightsRequestRequestBuilder instantiates a new SubjectRightsRequestRequestBuilder and sets the default values.
func NewSubjectRightsRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SubjectRightsRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectRightsRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property subjectRightsRequests for privacy
func (m *SubjectRightsRequestRequestBuilder) CreateDeleteRequestInformation(options *SubjectRightsRequestRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get subjectRightsRequests from privacy
func (m *SubjectRightsRequestRequestBuilder) CreateGetRequestInformation(options *SubjectRightsRequestRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property subjectRightsRequests in privacy
func (m *SubjectRightsRequestRequestBuilder) CreatePatchRequestInformation(options *SubjectRightsRequestRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property subjectRightsRequests for privacy
func (m *SubjectRightsRequestRequestBuilder) Delete(options *SubjectRightsRequestRequestBuilderDeleteOptions)(error) {
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
// Get get subjectRightsRequests from privacy
func (m *SubjectRightsRequestRequestBuilder) Get(options *SubjectRightsRequestRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SubjectRightsRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSubjectRightsRequest() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SubjectRightsRequest), nil
}
// GetFinalAttachment builds and executes requests for operations under \privacy\subjectRightsRequests\{subjectRightsRequest-id}\microsoft.graph.getFinalAttachment()
func (m *SubjectRightsRequestRequestBuilder) GetFinalAttachment()(*i2bf59a4d763ab3355a4526b9ff2ddcae384f1210dd464ff5940ff3a054d5f26e.GetFinalAttachmentRequestBuilder) {
    return i2bf59a4d763ab3355a4526b9ff2ddcae384f1210dd464ff5940ff3a054d5f26e.NewGetFinalAttachmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFinalReport builds and executes requests for operations under \privacy\subjectRightsRequests\{subjectRightsRequest-id}\microsoft.graph.getFinalReport()
func (m *SubjectRightsRequestRequestBuilder) GetFinalReport()(*i958489264196dd08f36dd719f907da913cf1a6b619228f3d6bc8da52997d8449.GetFinalReportRequestBuilder) {
    return i958489264196dd08f36dd719f907da913cf1a6b619228f3d6bc8da52997d8449.NewGetFinalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SubjectRightsRequestRequestBuilder) Notes()(*i06e33bb086cbb5f9333cc1c8bb09529f00623800d8f148e19c13fab399be0e21.NotesRequestBuilder) {
    return i06e33bb086cbb5f9333cc1c8bb09529f00623800d8f148e19c13fab399be0e21.NewNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privacy.subjectRightsRequests.item.notes.item collection
func (m *SubjectRightsRequestRequestBuilder) NotesById(id string)(*ibe7a5a69eee3e39fd9fd52c318c683abd357fbe438629376f053fbf8dafee6d9.AuthoredNoteRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authoredNote_id"] = id
    }
    return ibe7a5a69eee3e39fd9fd52c318c683abd357fbe438629376f053fbf8dafee6d9.NewAuthoredNoteRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property subjectRightsRequests in privacy
func (m *SubjectRightsRequestRequestBuilder) Patch(options *SubjectRightsRequestRequestBuilderPatchOptions)(error) {
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
func (m *SubjectRightsRequestRequestBuilder) Team()(*ia790ce0fe221c2e579ab27030d34081e5415f38f47af2263081eb369b2561ccd.TeamRequestBuilder) {
    return ia790ce0fe221c2e579ab27030d34081e5415f38f47af2263081eb369b2561ccd.NewTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
