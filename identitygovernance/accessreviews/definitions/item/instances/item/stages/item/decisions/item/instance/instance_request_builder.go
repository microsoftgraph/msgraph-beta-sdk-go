package instance

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1ba863fa5586aac600b8b6f0cec6f50920eeeb6658de97b61ec0da2b0e5045d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/stages/item/decisions/item/instance/stop"
    i9370ff97a612bfd87d39f181601e42335b643df8ea16eb203974508bfe70fe00 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/stages/item/decisions/item/instance/sendreminder"
    ib8ae3d5b6c556d579f0709acab9f33396c56b3c7693a49fb4b9308c5eae886b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/stages/item/decisions/item/instance/acceptrecommendations"
    id7265d114b68cbea5226c0b83a4a8eb32944aca353e79212da9dbfdf60309090 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/stages/item/decisions/item/instance/applydecisions"
    if2f9c146f2936ccff0dadf479287443cb43e958bf7fb5235e568fcf8b21071a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/stages/item/decisions/item/instance/resetdecisions"
    if7183e43d3bd020a7357ce03103e242823973522d5fd91e375eb7a9d2f57deb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/stages/item/decisions/item/instance/batchrecorddecisions"
)

// InstanceRequestBuilder builds and executes requests for operations under \identityGovernance\accessReviews\definitions\{accessReviewScheduleDefinition-id}\instances\{accessReviewInstance-id}\stages\{accessReviewStage-id}\decisions\{accessReviewInstanceDecisionItem-id}\instance
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
func (m *InstanceRequestBuilder) AcceptRecommendations()(*ib8ae3d5b6c556d579f0709acab9f33396c56b3c7693a49fb4b9308c5eae886b5.AcceptRecommendationsRequestBuilder) {
    return ib8ae3d5b6c556d579f0709acab9f33396c56b3c7693a49fb4b9308c5eae886b5.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) ApplyDecisions()(*id7265d114b68cbea5226c0b83a4a8eb32944aca353e79212da9dbfdf60309090.ApplyDecisionsRequestBuilder) {
    return id7265d114b68cbea5226c0b83a4a8eb32944aca353e79212da9dbfdf60309090.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) BatchRecordDecisions()(*if7183e43d3bd020a7357ce03103e242823973522d5fd91e375eb7a9d2f57deb0.BatchRecordDecisionsRequestBuilder) {
    return if7183e43d3bd020a7357ce03103e242823973522d5fd91e375eb7a9d2f57deb0.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewInstanceRequestBuilderInternal instantiates a new InstanceRequestBuilder and sets the default values.
func NewInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InstanceRequestBuilder) {
    m := &InstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition_id}/instances/{accessReviewInstance_id}/stages/{accessReviewStage_id}/decisions/{accessReviewInstanceDecisionItem_id}/instance{?select,expand}";
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
func (m *InstanceRequestBuilder) ResetDecisions()(*if2f9c146f2936ccff0dadf479287443cb43e958bf7fb5235e568fcf8b21071a7.ResetDecisionsRequestBuilder) {
    return if2f9c146f2936ccff0dadf479287443cb43e958bf7fb5235e568fcf8b21071a7.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) SendReminder()(*i9370ff97a612bfd87d39f181601e42335b643df8ea16eb203974508bfe70fe00.SendReminderRequestBuilder) {
    return i9370ff97a612bfd87d39f181601e42335b643df8ea16eb203974508bfe70fe00.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) Stop()(*i1ba863fa5586aac600b8b6f0cec6f50920eeeb6658de97b61ec0da2b0e5045d3.StopRequestBuilder) {
    return i1ba863fa5586aac600b8b6f0cec6f50920eeeb6658de97b61ec0da2b0e5045d3.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
