package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0a33d68030db60d4abf83d3cc44c2bd225ebb09fb65dabf238222f26a7b3aeca "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/acceptrecommendations"
    i157fdf67ef27b693ebfa46b836a81baa53b2fffe51cf8faa76f9ff1c2be3e533 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/contactedreviewers"
    i29e3e31b50fe024cbfcedaf85025abb3ad402366b9a1338c8b8bc0d397a92401 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/definition"
    i2bfb15080611c6b6292e90319c54737f8b19d678a61e405dd1bd25c0d3aa9c4e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages"
    i2c33bdddccade847db8f95ca5491389dc559779eb3443e04a0bfcb0e6e3fcc57 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/batchrecorddecisions"
    i2e9774ef796277adf2753e94982cf69a710728251ccf90259642dcf552172ca8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/resetdecisions"
    i4b5e7a16aef847860e0b992bfc76a3fe54ff29e5afb2bd1e2369911539e54814 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/applydecisions"
    i5b136c8ab95b62c794f13ee8a48c92d1c62d5dc6c75f5e4dc5bf16a4aeb7a432 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stop"
    i8ef41e575ff6010a2d837801cb7f0c7340ef83813e02674cb4ca79971d31a2a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/sendreminder"
    ibcccde7c49f5c5bbcb5ceea96760dde9fbef276461349ba11507419b629f04e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/decisions"
    i0b5d9b80b48fdacc03dda0d4899b47787af4703b66ab3b6c99f6c290394d6f15 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/decisions/item"
    i6af551f0724120edad9980686a562c6ee0ec8f50a04abb46cb03bf7eacb81ac5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/contactedreviewers/item"
    i7b1e883f98e741ec9b8bb6f1e9c089fcc58af710068f93b642d83ad60fa83c77 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item"
)

// AccessReviewInstanceRequestBuilder builds and executes requests for operations under \users\{user-id}\pendingAccessReviewInstances\{accessReviewInstance-id}
type AccessReviewInstanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessReviewInstanceRequestBuilderDeleteOptions options for Delete
type AccessReviewInstanceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewInstanceRequestBuilderGetOptions options for Get
type AccessReviewInstanceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessReviewInstanceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessReviewInstanceRequestBuilderGetQueryParameters navigation property to get list of access reviews pending approval by reviewer.
type AccessReviewInstanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessReviewInstanceRequestBuilderPatchOptions options for Patch
type AccessReviewInstanceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstance;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessReviewInstanceRequestBuilder) AcceptRecommendations()(*i0a33d68030db60d4abf83d3cc44c2bd225ebb09fb65dabf238222f26a7b3aeca.AcceptRecommendationsRequestBuilder) {
    return i0a33d68030db60d4abf83d3cc44c2bd225ebb09fb65dabf238222f26a7b3aeca.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) ApplyDecisions()(*i4b5e7a16aef847860e0b992bfc76a3fe54ff29e5afb2bd1e2369911539e54814.ApplyDecisionsRequestBuilder) {
    return i4b5e7a16aef847860e0b992bfc76a3fe54ff29e5afb2bd1e2369911539e54814.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) BatchRecordDecisions()(*i2c33bdddccade847db8f95ca5491389dc559779eb3443e04a0bfcb0e6e3fcc57.BatchRecordDecisionsRequestBuilder) {
    return i2c33bdddccade847db8f95ca5491389dc559779eb3443e04a0bfcb0e6e3fcc57.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessReviewInstanceRequestBuilderInternal instantiates a new AccessReviewInstanceRequestBuilder and sets the default values.
func NewAccessReviewInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewInstanceRequestBuilder) {
    m := &AccessReviewInstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/pendingAccessReviewInstances/{accessReviewInstance_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessReviewInstanceRequestBuilder instantiates a new AccessReviewInstanceRequestBuilder and sets the default values.
func NewAccessReviewInstanceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewInstanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewInstanceRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AccessReviewInstanceRequestBuilder) ContactedReviewers()(*i157fdf67ef27b693ebfa46b836a81baa53b2fffe51cf8faa76f9ff1c2be3e533.ContactedReviewersRequestBuilder) {
    return i157fdf67ef27b693ebfa46b836a81baa53b2fffe51cf8faa76f9ff1c2be3e533.NewContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.pendingAccessReviewInstances.item.contactedReviewers.item collection
func (m *AccessReviewInstanceRequestBuilder) ContactedReviewersById(id string)(*i6af551f0724120edad9980686a562c6ee0ec8f50a04abb46cb03bf7eacb81ac5.AccessReviewReviewerRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer_id"] = id
    }
    return i6af551f0724120edad9980686a562c6ee0ec8f50a04abb46cb03bf7eacb81ac5.NewAccessReviewReviewerRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation navigation property to get list of access reviews pending approval by reviewer.
func (m *AccessReviewInstanceRequestBuilder) CreateDeleteRequestInformation(options *AccessReviewInstanceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation navigation property to get list of access reviews pending approval by reviewer.
func (m *AccessReviewInstanceRequestBuilder) CreateGetRequestInformation(options *AccessReviewInstanceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation navigation property to get list of access reviews pending approval by reviewer.
func (m *AccessReviewInstanceRequestBuilder) CreatePatchRequestInformation(options *AccessReviewInstanceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AccessReviewInstanceRequestBuilder) Decisions()(*ibcccde7c49f5c5bbcb5ceea96760dde9fbef276461349ba11507419b629f04e6.DecisionsRequestBuilder) {
    return ibcccde7c49f5c5bbcb5ceea96760dde9fbef276461349ba11507419b629f04e6.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.pendingAccessReviewInstances.item.decisions.item collection
func (m *AccessReviewInstanceRequestBuilder) DecisionsById(id string)(*i0b5d9b80b48fdacc03dda0d4899b47787af4703b66ab3b6c99f6c290394d6f15.AccessReviewInstanceDecisionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem_id"] = id
    }
    return i0b5d9b80b48fdacc03dda0d4899b47787af4703b66ab3b6c99f6c290394d6f15.NewAccessReviewInstanceDecisionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) Definition()(*i29e3e31b50fe024cbfcedaf85025abb3ad402366b9a1338c8b8bc0d397a92401.DefinitionRequestBuilder) {
    return i29e3e31b50fe024cbfcedaf85025abb3ad402366b9a1338c8b8bc0d397a92401.NewDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete navigation property to get list of access reviews pending approval by reviewer.
func (m *AccessReviewInstanceRequestBuilder) Delete(options *AccessReviewInstanceRequestBuilderDeleteOptions)(error) {
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
// Get navigation property to get list of access reviews pending approval by reviewer.
func (m *AccessReviewInstanceRequestBuilder) Get(options *AccessReviewInstanceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstance, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessReviewInstance() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstance), nil
}
// Patch navigation property to get list of access reviews pending approval by reviewer.
func (m *AccessReviewInstanceRequestBuilder) Patch(options *AccessReviewInstanceRequestBuilderPatchOptions)(error) {
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
func (m *AccessReviewInstanceRequestBuilder) ResetDecisions()(*i2e9774ef796277adf2753e94982cf69a710728251ccf90259642dcf552172ca8.ResetDecisionsRequestBuilder) {
    return i2e9774ef796277adf2753e94982cf69a710728251ccf90259642dcf552172ca8.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) SendReminder()(*i8ef41e575ff6010a2d837801cb7f0c7340ef83813e02674cb4ca79971d31a2a2.SendReminderRequestBuilder) {
    return i8ef41e575ff6010a2d837801cb7f0c7340ef83813e02674cb4ca79971d31a2a2.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) Stages()(*i2bfb15080611c6b6292e90319c54737f8b19d678a61e405dd1bd25c0d3aa9c4e.StagesRequestBuilder) {
    return i2bfb15080611c6b6292e90319c54737f8b19d678a61e405dd1bd25c0d3aa9c4e.NewStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.pendingAccessReviewInstances.item.stages.item collection
func (m *AccessReviewInstanceRequestBuilder) StagesById(id string)(*i7b1e883f98e741ec9b8bb6f1e9c089fcc58af710068f93b642d83ad60fa83c77.AccessReviewStageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage_id"] = id
    }
    return i7b1e883f98e741ec9b8bb6f1e9c089fcc58af710068f93b642d83ad60fa83c77.NewAccessReviewStageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) Stop()(*i5b136c8ab95b62c794f13ee8a48c92d1c62d5dc6c75f5e4dc5bf16a4aeb7a432.StopRequestBuilder) {
    return i5b136c8ab95b62c794f13ee8a48c92d1c62d5dc6c75f5e4dc5bf16a4aeb7a432.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
