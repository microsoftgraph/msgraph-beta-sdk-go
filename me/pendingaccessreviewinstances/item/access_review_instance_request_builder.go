package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i226aac2d9e753eec2396e7a04e06f10afaa0784792e00d5fe99fea4069422656 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/batchrecorddecisions"
    i28cbfe53788555c1e8fb2bcbb1f7b5ee3a25273656373f13f5ff7b954d036e04 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/contactedreviewers"
    i5a33caa17867d8a52a52d95e786e9714ce0e9a6b07f5e9aa14f00f581f95c45f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/stop"
    i6be43a98526e3cd95557ccc422ad7cdaa8c43585327deaffa99f4367f522a108 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/stages"
    i78cfe79793d77a5fede4967db4de1713d8be67261324f74669c2af685e4134f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/applydecisions"
    i857369f7b2494f5fc984108c41f95436f6b0349231c74000340f9f1547e30300 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions"
    i978a7302921342acb736ff821abff5eb57189f5f20e0b9d44a2800bd91c9bc2e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/resetdecisions"
    ia0d94a184a000bb845f57559a39262d0cde1e4cbd62ba5b70e7f517eff968628 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/acceptrecommendations"
    ia8e711f9bce6e1855eb0b9ef5889095dbc6f1dff81d7c5cb1530e08f1e8af277 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/definition"
    icbbe17c4811e8e4855cc7841cf8914b2f2a20b1d41e16779ac8b42dd4ad533a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/sendreminder"
    i92a83f24221d12aedaef668e1fbb993200a107c50dd873c65f56c93ab8abaa7e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item"
    ia273f0c967440d14bb94006b177f70e6b83152c6fcad8ccbf86dc8739482bc95 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/contactedreviewers/item"
    ie9afc88f4d0ebd933223975c8172435f458acf5c8c922ff0f1d55f973e0ed66f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/stages/item"
)

// AccessReviewInstanceRequestBuilder builds and executes requests for operations under \me\pendingAccessReviewInstances\{accessReviewInstance-id}
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
func (m *AccessReviewInstanceRequestBuilder) AcceptRecommendations()(*ia0d94a184a000bb845f57559a39262d0cde1e4cbd62ba5b70e7f517eff968628.AcceptRecommendationsRequestBuilder) {
    return ia0d94a184a000bb845f57559a39262d0cde1e4cbd62ba5b70e7f517eff968628.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) ApplyDecisions()(*i78cfe79793d77a5fede4967db4de1713d8be67261324f74669c2af685e4134f8.ApplyDecisionsRequestBuilder) {
    return i78cfe79793d77a5fede4967db4de1713d8be67261324f74669c2af685e4134f8.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) BatchRecordDecisions()(*i226aac2d9e753eec2396e7a04e06f10afaa0784792e00d5fe99fea4069422656.BatchRecordDecisionsRequestBuilder) {
    return i226aac2d9e753eec2396e7a04e06f10afaa0784792e00d5fe99fea4069422656.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessReviewInstanceRequestBuilderInternal instantiates a new AccessReviewInstanceRequestBuilder and sets the default values.
func NewAccessReviewInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewInstanceRequestBuilder) {
    m := &AccessReviewInstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/pendingAccessReviewInstances/{accessReviewInstance_id}{?select,expand}";
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
func (m *AccessReviewInstanceRequestBuilder) ContactedReviewers()(*i28cbfe53788555c1e8fb2bcbb1f7b5ee3a25273656373f13f5ff7b954d036e04.ContactedReviewersRequestBuilder) {
    return i28cbfe53788555c1e8fb2bcbb1f7b5ee3a25273656373f13f5ff7b954d036e04.NewContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.pendingAccessReviewInstances.item.contactedReviewers.item collection
func (m *AccessReviewInstanceRequestBuilder) ContactedReviewersById(id string)(*ia273f0c967440d14bb94006b177f70e6b83152c6fcad8ccbf86dc8739482bc95.AccessReviewReviewerRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer_id"] = id
    }
    return ia273f0c967440d14bb94006b177f70e6b83152c6fcad8ccbf86dc8739482bc95.NewAccessReviewReviewerRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *AccessReviewInstanceRequestBuilder) Decisions()(*i857369f7b2494f5fc984108c41f95436f6b0349231c74000340f9f1547e30300.DecisionsRequestBuilder) {
    return i857369f7b2494f5fc984108c41f95436f6b0349231c74000340f9f1547e30300.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.pendingAccessReviewInstances.item.decisions.item collection
func (m *AccessReviewInstanceRequestBuilder) DecisionsById(id string)(*i92a83f24221d12aedaef668e1fbb993200a107c50dd873c65f56c93ab8abaa7e.AccessReviewInstanceDecisionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem_id"] = id
    }
    return i92a83f24221d12aedaef668e1fbb993200a107c50dd873c65f56c93ab8abaa7e.NewAccessReviewInstanceDecisionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) Definition()(*ia8e711f9bce6e1855eb0b9ef5889095dbc6f1dff81d7c5cb1530e08f1e8af277.DefinitionRequestBuilder) {
    return ia8e711f9bce6e1855eb0b9ef5889095dbc6f1dff81d7c5cb1530e08f1e8af277.NewDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *AccessReviewInstanceRequestBuilder) ResetDecisions()(*i978a7302921342acb736ff821abff5eb57189f5f20e0b9d44a2800bd91c9bc2e.ResetDecisionsRequestBuilder) {
    return i978a7302921342acb736ff821abff5eb57189f5f20e0b9d44a2800bd91c9bc2e.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) SendReminder()(*icbbe17c4811e8e4855cc7841cf8914b2f2a20b1d41e16779ac8b42dd4ad533a6.SendReminderRequestBuilder) {
    return icbbe17c4811e8e4855cc7841cf8914b2f2a20b1d41e16779ac8b42dd4ad533a6.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) Stages()(*i6be43a98526e3cd95557ccc422ad7cdaa8c43585327deaffa99f4367f522a108.StagesRequestBuilder) {
    return i6be43a98526e3cd95557ccc422ad7cdaa8c43585327deaffa99f4367f522a108.NewStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.pendingAccessReviewInstances.item.stages.item collection
func (m *AccessReviewInstanceRequestBuilder) StagesById(id string)(*ie9afc88f4d0ebd933223975c8172435f458acf5c8c922ff0f1d55f973e0ed66f.AccessReviewStageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage_id"] = id
    }
    return ie9afc88f4d0ebd933223975c8172435f458acf5c8c922ff0f1d55f973e0ed66f.NewAccessReviewStageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) Stop()(*i5a33caa17867d8a52a52d95e786e9714ce0e9a6b07f5e9aa14f00f581f95c45f.StopRequestBuilder) {
    return i5a33caa17867d8a52a52d95e786e9714ce0e9a6b07f5e9aa14f00f581f95c45f.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
