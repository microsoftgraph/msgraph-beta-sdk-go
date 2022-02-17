package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i34029e4c2d0a245f6ca7822cbbd442b2a54ddd3d92e7f8f9797638a7756d56d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/stages"
    i4167933608f10c80b058ba8a45ffcaff5a3638d6f3458cded3059b347646eabe "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/sendreminder"
    i6a0a160cb4443ebb8d098e51ce55956f64d15462282145c7d7333c345b943534 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/resetdecisions"
    i9327ed35d3b3a7e13ba1ba8f389450b462683ccdaccff95c747cb5f750f0d4d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/decisions"
    iaee1412a0b517074ea75eb9acdb529e3a9e5a31a64757270c95df5946da814b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/stop"
    ic0e8dea6d9a1d97ce97cabb68e83dd9aa75ea89e6f734ec7f271d9ba27f052ce "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/contactedreviewers"
    ic519225da39bf9a83bff86a837186671f7908ceabe699cf8fab5579f59cb50c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/batchrecorddecisions"
    ic8e74603f5ee41faf457cafb716b509d060bc44d1f5078afd66ec5bf53e64729 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/applydecisions"
    id859fb99d92f5a45bd2ce6145abca710efa33e2a5b1fe88b4ff16f1f437a67c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/acceptrecommendations"
    ie25b10e68fa81f5005eea5f527a61e7a311f8b60257495559d75d76510909fe3 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/definition"
    i425fb16e7ed420101998d1e876028364d5747a9897e3d5604f4a824ec3e04ffb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/decisions/item"
    i778254787cb626d9fe348e2d305e8501a888ae97b99b9154fb0f2bee3401ccea "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/contactedreviewers/item"
    i83c0dcbc9b5cafef3dad8f565d30a946fd80a79a1ba13afeef9c56a7f634d98e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/stages/item"
)

// AccessReviewInstanceRequestBuilder builds and executes requests for operations under \identityGovernance\accessReviews\definitions\{accessReviewScheduleDefinition-id}\instances\{accessReviewInstance-id}
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
// AccessReviewInstanceRequestBuilderGetQueryParameters if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
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
func (m *AccessReviewInstanceRequestBuilder) AcceptRecommendations()(*id859fb99d92f5a45bd2ce6145abca710efa33e2a5b1fe88b4ff16f1f437a67c5.AcceptRecommendationsRequestBuilder) {
    return id859fb99d92f5a45bd2ce6145abca710efa33e2a5b1fe88b4ff16f1f437a67c5.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) ApplyDecisions()(*ic8e74603f5ee41faf457cafb716b509d060bc44d1f5078afd66ec5bf53e64729.ApplyDecisionsRequestBuilder) {
    return ic8e74603f5ee41faf457cafb716b509d060bc44d1f5078afd66ec5bf53e64729.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) BatchRecordDecisions()(*ic519225da39bf9a83bff86a837186671f7908ceabe699cf8fab5579f59cb50c3.BatchRecordDecisionsRequestBuilder) {
    return ic519225da39bf9a83bff86a837186671f7908ceabe699cf8fab5579f59cb50c3.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessReviewInstanceRequestBuilderInternal instantiates a new AccessReviewInstanceRequestBuilder and sets the default values.
func NewAccessReviewInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewInstanceRequestBuilder) {
    m := &AccessReviewInstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition_id}/instances/{accessReviewInstance_id}{?select,expand}";
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
func (m *AccessReviewInstanceRequestBuilder) ContactedReviewers()(*ic0e8dea6d9a1d97ce97cabb68e83dd9aa75ea89e6f734ec7f271d9ba27f052ce.ContactedReviewersRequestBuilder) {
    return ic0e8dea6d9a1d97ce97cabb68e83dd9aa75ea89e6f734ec7f271d9ba27f052ce.NewContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.accessReviews.definitions.item.instances.item.contactedReviewers.item collection
func (m *AccessReviewInstanceRequestBuilder) ContactedReviewersById(id string)(*i778254787cb626d9fe348e2d305e8501a888ae97b99b9154fb0f2bee3401ccea.AccessReviewReviewerRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer_id"] = id
    }
    return i778254787cb626d9fe348e2d305e8501a888ae97b99b9154fb0f2bee3401ccea.NewAccessReviewReviewerRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
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
// CreateGetRequestInformation if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
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
// CreatePatchRequestInformation if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
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
func (m *AccessReviewInstanceRequestBuilder) Decisions()(*i9327ed35d3b3a7e13ba1ba8f389450b462683ccdaccff95c747cb5f750f0d4d8.DecisionsRequestBuilder) {
    return i9327ed35d3b3a7e13ba1ba8f389450b462683ccdaccff95c747cb5f750f0d4d8.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.accessReviews.definitions.item.instances.item.decisions.item collection
func (m *AccessReviewInstanceRequestBuilder) DecisionsById(id string)(*i425fb16e7ed420101998d1e876028364d5747a9897e3d5604f4a824ec3e04ffb.AccessReviewInstanceDecisionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem_id"] = id
    }
    return i425fb16e7ed420101998d1e876028364d5747a9897e3d5604f4a824ec3e04ffb.NewAccessReviewInstanceDecisionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) Definition()(*ie25b10e68fa81f5005eea5f527a61e7a311f8b60257495559d75d76510909fe3.DefinitionRequestBuilder) {
    return ie25b10e68fa81f5005eea5f527a61e7a311f8b60257495559d75d76510909fe3.NewDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
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
// Get if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
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
// Patch if the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
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
func (m *AccessReviewInstanceRequestBuilder) ResetDecisions()(*i6a0a160cb4443ebb8d098e51ce55956f64d15462282145c7d7333c345b943534.ResetDecisionsRequestBuilder) {
    return i6a0a160cb4443ebb8d098e51ce55956f64d15462282145c7d7333c345b943534.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) SendReminder()(*i4167933608f10c80b058ba8a45ffcaff5a3638d6f3458cded3059b347646eabe.SendReminderRequestBuilder) {
    return i4167933608f10c80b058ba8a45ffcaff5a3638d6f3458cded3059b347646eabe.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) Stages()(*i34029e4c2d0a245f6ca7822cbbd442b2a54ddd3d92e7f8f9797638a7756d56d6.StagesRequestBuilder) {
    return i34029e4c2d0a245f6ca7822cbbd442b2a54ddd3d92e7f8f9797638a7756d56d6.NewStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.accessReviews.definitions.item.instances.item.stages.item collection
func (m *AccessReviewInstanceRequestBuilder) StagesById(id string)(*i83c0dcbc9b5cafef3dad8f565d30a946fd80a79a1ba13afeef9c56a7f634d98e.AccessReviewStageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage_id"] = id
    }
    return i83c0dcbc9b5cafef3dad8f565d30a946fd80a79a1ba13afeef9c56a7f634d98e.NewAccessReviewStageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessReviewInstanceRequestBuilder) Stop()(*iaee1412a0b517074ea75eb9acdb529e3a9e5a31a64757270c95df5946da814b4.StopRequestBuilder) {
    return iaee1412a0b517074ea75eb9acdb529e3a9e5a31a64757270c95df5946da814b4.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
