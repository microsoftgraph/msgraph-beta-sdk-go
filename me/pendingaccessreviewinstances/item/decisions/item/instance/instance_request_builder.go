package instance

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i03da4522f6ecfec1c06db2ac24e7752b8ccd635602f5fe82922c3a68e1d0dd7d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/resetdecisions"
    i33c7c5706a7b5c0bf13b40b705217aed59b2d5208c669fd999e36faa62df7d8d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/acceptrecommendations"
    i74deae20af1cfa651b11f8e4117763fc91498e69766fedcf61e65d3873b58bd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/stop"
    ia755ae82953cc4dc7511248e4432b31518b4ceb68464488197f488ec3a27c58f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/sendreminder"
    icd1d27bb795cd0bade6075661b23050cc531c7e953d93767d6867d925acd1153 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/batchrecorddecisions"
    iebc762487b2890223499b50b54324ea41c534597ca8a925c094747c4cc9293ba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/applydecisions"
)

// InstanceRequestBuilder builds and executes requests for operations under \me\pendingAccessReviewInstances\{accessReviewInstance-id}\decisions\{accessReviewInstanceDecisionItem-id}\instance
type InstanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// InstanceRequestBuilderDeleteOptions options for Delete
type InstanceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// InstanceRequestBuilderGetOptions options for Get
type InstanceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *InstanceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// InstanceRequestBuilderGetQueryParameters there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
type InstanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// InstanceRequestBuilderPatchOptions options for Patch
type InstanceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstance;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *InstanceRequestBuilder) AcceptRecommendations()(*i33c7c5706a7b5c0bf13b40b705217aed59b2d5208c669fd999e36faa62df7d8d.AcceptRecommendationsRequestBuilder) {
    return i33c7c5706a7b5c0bf13b40b705217aed59b2d5208c669fd999e36faa62df7d8d.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) ApplyDecisions()(*iebc762487b2890223499b50b54324ea41c534597ca8a925c094747c4cc9293ba.ApplyDecisionsRequestBuilder) {
    return iebc762487b2890223499b50b54324ea41c534597ca8a925c094747c4cc9293ba.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) BatchRecordDecisions()(*icd1d27bb795cd0bade6075661b23050cc531c7e953d93767d6867d925acd1153.BatchRecordDecisionsRequestBuilder) {
    return icd1d27bb795cd0bade6075661b23050cc531c7e953d93767d6867d925acd1153.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewInstanceRequestBuilderInternal instantiates a new InstanceRequestBuilder and sets the default values.
func NewInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InstanceRequestBuilder) {
    m := &InstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/pendingAccessReviewInstances/{accessReviewInstance_id}/decisions/{accessReviewInstanceDecisionItem_id}/instance{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInstanceRequestBuilder instantiates a new InstanceRequestBuilder and sets the default values.
func NewInstanceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InstanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstanceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) CreateDeleteRequestInformation(options *InstanceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) CreateGetRequestInformation(options *InstanceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) CreatePatchRequestInformation(options *InstanceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) Delete(options *InstanceRequestBuilderDeleteOptions)(error) {
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
// Get there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) Get(options *InstanceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewInstance, error) {
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
// Patch there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) Patch(options *InstanceRequestBuilderPatchOptions)(error) {
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
func (m *InstanceRequestBuilder) ResetDecisions()(*i03da4522f6ecfec1c06db2ac24e7752b8ccd635602f5fe82922c3a68e1d0dd7d.ResetDecisionsRequestBuilder) {
    return i03da4522f6ecfec1c06db2ac24e7752b8ccd635602f5fe82922c3a68e1d0dd7d.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) SendReminder()(*ia755ae82953cc4dc7511248e4432b31518b4ceb68464488197f488ec3a27c58f.SendReminderRequestBuilder) {
    return ia755ae82953cc4dc7511248e4432b31518b4ceb68464488197f488ec3a27c58f.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) Stop()(*i74deae20af1cfa651b11f8e4117763fc91498e69766fedcf61e65d3873b58bd6.StopRequestBuilder) {
    return i74deae20af1cfa651b11f8e4117763fc91498e69766fedcf61e65d3873b58bd6.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
