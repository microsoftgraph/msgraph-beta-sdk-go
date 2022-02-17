package instance

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i112309bcdf0c75639081cbf60ec6f9c39955c11f95422199c3e1aa9fd9b60ec4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/definition"
    i3a336a3f181639d0b400b4f90499430d40c47aec7a3f3f5b8992a6d831caaba2 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/contactedreviewers"
    i5cd91d67258652f5b2854613fdc8776da7412faf0be7b08d61715e958f802163 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/stop"
    i634ca318551c7e45221adbee2751c7e66b920ae332f90c95ef4ac4389fa7d373 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/stages"
    i7ee65cca46e5c3dd41258be0d0dfef78edbf481c1ca6ee31687f5ebe440c81cf "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/resetdecisions"
    i94c2546ce2c1ca2aaff4bfb460192a04c9360e5018d95ff5816cb03f92cbd9c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/batchrecorddecisions"
    i9eb2370a85ce8df5e8f50351054e9a4bdc9aa063da2ef2ea98a9f58a43b4045c "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/acceptrecommendations"
    ibf0291b18a35977b898932783008edde4b8fbfb3b1ac342362af4075f07e314e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/sendreminder"
    ibf4f3d33172cc835cda4e47809b585e37190bab1e76e467b086a125d159a0ce3 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/decisions"
    ifa7d755cd637f40d9c35f384521d33ee07417fba93337db7889f9e370af0bdca "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/applydecisions"
    i2cf7965e5bcacbaf398832a7f2f84cae59e8d9123c23a4d7df880c767c813358 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/contactedreviewers/item"
    i8f76eb834e079485d9bfcc198960c387e47a032c16f3c5e1a1d9779d19557d0e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/decisions/item"
    ie737e21101e5c8fcb439d2621c3dbdcc10511191b0a9e20174987172b109dc3c "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/decisions/item/instance/stages/item"
)

// InstanceRequestBuilder builds and executes requests for operations under \identityGovernance\accessReviews\decisions\{accessReviewInstanceDecisionItem-id}\instance
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
func (m *InstanceRequestBuilder) AcceptRecommendations()(*i9eb2370a85ce8df5e8f50351054e9a4bdc9aa063da2ef2ea98a9f58a43b4045c.AcceptRecommendationsRequestBuilder) {
    return i9eb2370a85ce8df5e8f50351054e9a4bdc9aa063da2ef2ea98a9f58a43b4045c.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) ApplyDecisions()(*ifa7d755cd637f40d9c35f384521d33ee07417fba93337db7889f9e370af0bdca.ApplyDecisionsRequestBuilder) {
    return ifa7d755cd637f40d9c35f384521d33ee07417fba93337db7889f9e370af0bdca.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) BatchRecordDecisions()(*i94c2546ce2c1ca2aaff4bfb460192a04c9360e5018d95ff5816cb03f92cbd9c9.BatchRecordDecisionsRequestBuilder) {
    return i94c2546ce2c1ca2aaff4bfb460192a04c9360e5018d95ff5816cb03f92cbd9c9.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewInstanceRequestBuilderInternal instantiates a new InstanceRequestBuilder and sets the default values.
func NewInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InstanceRequestBuilder) {
    m := &InstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/decisions/{accessReviewInstanceDecisionItem_id}/instance{?select,expand}";
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
func (m *InstanceRequestBuilder) ContactedReviewers()(*i3a336a3f181639d0b400b4f90499430d40c47aec7a3f3f5b8992a6d831caaba2.ContactedReviewersRequestBuilder) {
    return i3a336a3f181639d0b400b4f90499430d40c47aec7a3f3f5b8992a6d831caaba2.NewContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.accessReviews.decisions.item.instance.contactedReviewers.item collection
func (m *InstanceRequestBuilder) ContactedReviewersById(id string)(*i2cf7965e5bcacbaf398832a7f2f84cae59e8d9123c23a4d7df880c767c813358.AccessReviewReviewerRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer_id"] = id
    }
    return i2cf7965e5bcacbaf398832a7f2f84cae59e8d9123c23a4d7df880c767c813358.NewAccessReviewReviewerRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *InstanceRequestBuilder) Decisions()(*ibf4f3d33172cc835cda4e47809b585e37190bab1e76e467b086a125d159a0ce3.DecisionsRequestBuilder) {
    return ibf4f3d33172cc835cda4e47809b585e37190bab1e76e467b086a125d159a0ce3.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.accessReviews.decisions.item.instance.decisions.item collection
func (m *InstanceRequestBuilder) DecisionsById(id string)(*i8f76eb834e079485d9bfcc198960c387e47a032c16f3c5e1a1d9779d19557d0e.AccessReviewInstanceDecisionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem_id1"] = id
    }
    return i8f76eb834e079485d9bfcc198960c387e47a032c16f3c5e1a1d9779d19557d0e.NewAccessReviewInstanceDecisionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InstanceRequestBuilder) Definition()(*i112309bcdf0c75639081cbf60ec6f9c39955c11f95422199c3e1aa9fd9b60ec4.DefinitionRequestBuilder) {
    return i112309bcdf0c75639081cbf60ec6f9c39955c11f95422199c3e1aa9fd9b60ec4.NewDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *InstanceRequestBuilder) ResetDecisions()(*i7ee65cca46e5c3dd41258be0d0dfef78edbf481c1ca6ee31687f5ebe440c81cf.ResetDecisionsRequestBuilder) {
    return i7ee65cca46e5c3dd41258be0d0dfef78edbf481c1ca6ee31687f5ebe440c81cf.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) SendReminder()(*ibf0291b18a35977b898932783008edde4b8fbfb3b1ac342362af4075f07e314e.SendReminderRequestBuilder) {
    return ibf0291b18a35977b898932783008edde4b8fbfb3b1ac342362af4075f07e314e.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *InstanceRequestBuilder) Stages()(*i634ca318551c7e45221adbee2751c7e66b920ae332f90c95ef4ac4389fa7d373.StagesRequestBuilder) {
    return i634ca318551c7e45221adbee2751c7e66b920ae332f90c95ef4ac4389fa7d373.NewStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.accessReviews.decisions.item.instance.stages.item collection
func (m *InstanceRequestBuilder) StagesById(id string)(*ie737e21101e5c8fcb439d2621c3dbdcc10511191b0a9e20174987172b109dc3c.AccessReviewStageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage_id"] = id
    }
    return ie737e21101e5c8fcb439d2621c3dbdcc10511191b0a9e20174987172b109dc3c.NewAccessReviewStageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *InstanceRequestBuilder) Stop()(*i5cd91d67258652f5b2854613fdc8776da7412faf0be7b08d61715e958f802163.StopRequestBuilder) {
    return i5cd91d67258652f5b2854613fdc8776da7412faf0be7b08d61715e958f802163.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
