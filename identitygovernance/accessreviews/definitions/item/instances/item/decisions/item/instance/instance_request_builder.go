package instance

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0c9e12b891d2a58a7ca2affb5b2c40de052b15c0b7d7b1948efacead71cac455 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/decisions/item/instance/resetdecisions"
    i36e07bb4d46543d7080b81b6daa30da27915f513fb1feff2519ccd50d2a056b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/decisions/item/instance/stop"
    i3cbfca5fe6ae1be1ed3657c79b1f5eb405d013814412aa30d9acc50285f5948f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/decisions/item/instance/acceptrecommendations"
    i916700efa9c4b8dae6aa36d3da10f0750bdf50f89f822b4290c9fa42553faefe "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/decisions/item/instance/sendreminder"
    i9fb8201d892300f8a8b54fa1fddcf257f354421177fffc14511da85f2fe7aa27 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/decisions/item/instance/batchrecorddecisions"
    ie67965b8a5033c5b6f16edc76ea47ee2e403ca8880d5ad098498734f89a5619b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/decisions/item/instance/applydecisions"
)

// InstanceRequestBuilder builds and executes requests for operations under \identityGovernance\accessReviews\definitions\{accessReviewScheduleDefinition-id}\instances\{accessReviewInstance-id}\decisions\{accessReviewInstanceDecisionItem-id}\instance
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
func (m *InstanceRequestBuilder) AcceptRecommendations()(*i3cbfca5fe6ae1be1ed3657c79b1f5eb405d013814412aa30d9acc50285f5948f.AcceptRecommendationsRequestBuilder) {
    return i3cbfca5fe6ae1be1ed3657c79b1f5eb405d013814412aa30d9acc50285f5948f.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) ApplyDecisions()(*ie67965b8a5033c5b6f16edc76ea47ee2e403ca8880d5ad098498734f89a5619b.ApplyDecisionsRequestBuilder) {
    return ie67965b8a5033c5b6f16edc76ea47ee2e403ca8880d5ad098498734f89a5619b.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) BatchRecordDecisions()(*i9fb8201d892300f8a8b54fa1fddcf257f354421177fffc14511da85f2fe7aa27.BatchRecordDecisionsRequestBuilder) {
    return i9fb8201d892300f8a8b54fa1fddcf257f354421177fffc14511da85f2fe7aa27.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewInstanceRequestBuilderInternal instantiates a new InstanceRequestBuilder and sets the default values.
func NewInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InstanceRequestBuilder) {
    m := &InstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition_id}/instances/{accessReviewInstance_id}/decisions/{accessReviewInstanceDecisionItem_id}/instance{?select,expand}";
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
func (m *InstanceRequestBuilder) ResetDecisions()(*i0c9e12b891d2a58a7ca2affb5b2c40de052b15c0b7d7b1948efacead71cac455.ResetDecisionsRequestBuilder) {
    return i0c9e12b891d2a58a7ca2affb5b2c40de052b15c0b7d7b1948efacead71cac455.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) SendReminder()(*i916700efa9c4b8dae6aa36d3da10f0750bdf50f89f822b4290c9fa42553faefe.SendReminderRequestBuilder) {
    return i916700efa9c4b8dae6aa36d3da10f0750bdf50f89f822b4290c9fa42553faefe.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) Stop()(*i36e07bb4d46543d7080b81b6daa30da27915f513fb1feff2519ccd50d2a056b7.StopRequestBuilder) {
    return i36e07bb4d46543d7080b81b6daa30da27915f513fb1feff2519ccd50d2a056b7.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
