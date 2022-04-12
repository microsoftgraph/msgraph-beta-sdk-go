package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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

// AccessReviewInstanceItemRequestBuilder provides operations to manage the pendingAccessReviewInstances property of the microsoft.graph.user entity.
type AccessReviewInstanceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewInstanceItemRequestBuilderDeleteOptions options for Delete
type AccessReviewInstanceItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AccessReviewInstanceItemRequestBuilderGetOptions options for Get
type AccessReviewInstanceItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessReviewInstanceItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AccessReviewInstanceItemRequestBuilderGetQueryParameters navigation property to get list of access reviews pending approval by reviewer.
type AccessReviewInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessReviewInstanceItemRequestBuilderPatchOptions options for Patch
type AccessReviewInstanceItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AcceptRecommendations the acceptRecommendations property
func (m *AccessReviewInstanceItemRequestBuilder) AcceptRecommendations()(*i0a33d68030db60d4abf83d3cc44c2bd225ebb09fb65dabf238222f26a7b3aeca.AcceptRecommendationsRequestBuilder) {
    return i0a33d68030db60d4abf83d3cc44c2bd225ebb09fb65dabf238222f26a7b3aeca.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyDecisions the applyDecisions property
func (m *AccessReviewInstanceItemRequestBuilder) ApplyDecisions()(*i4b5e7a16aef847860e0b992bfc76a3fe54ff29e5afb2bd1e2369911539e54814.ApplyDecisionsRequestBuilder) {
    return i4b5e7a16aef847860e0b992bfc76a3fe54ff29e5afb2bd1e2369911539e54814.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BatchRecordDecisions the batchRecordDecisions property
func (m *AccessReviewInstanceItemRequestBuilder) BatchRecordDecisions()(*i2c33bdddccade847db8f95ca5491389dc559779eb3443e04a0bfcb0e6e3fcc57.BatchRecordDecisionsRequestBuilder) {
    return i2c33bdddccade847db8f95ca5491389dc559779eb3443e04a0bfcb0e6e3fcc57.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessReviewInstanceItemRequestBuilderInternal instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewAccessReviewInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewInstanceItemRequestBuilder) {
    m := &AccessReviewInstanceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessReviewInstanceItemRequestBuilder instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewAccessReviewInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers the contactedReviewers property
func (m *AccessReviewInstanceItemRequestBuilder) ContactedReviewers()(*i157fdf67ef27b693ebfa46b836a81baa53b2fffe51cf8faa76f9ff1c2be3e533.ContactedReviewersRequestBuilder) {
    return i157fdf67ef27b693ebfa46b836a81baa53b2fffe51cf8faa76f9ff1c2be3e533.NewContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.pendingAccessReviewInstances.item.contactedReviewers.item collection
func (m *AccessReviewInstanceItemRequestBuilder) ContactedReviewersById(id string)(*i6af551f0724120edad9980686a562c6ee0ec8f50a04abb46cb03bf7eacb81ac5.AccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return i6af551f0724120edad9980686a562c6ee0ec8f50a04abb46cb03bf7eacb81ac5.NewAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property pendingAccessReviewInstances for users
func (m *AccessReviewInstanceItemRequestBuilder) CreateDeleteRequestInformation(options *AccessReviewInstanceItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation navigation property to get list of access reviews pending approval by reviewer.
func (m *AccessReviewInstanceItemRequestBuilder) CreateGetRequestInformation(options *AccessReviewInstanceItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property pendingAccessReviewInstances in users
func (m *AccessReviewInstanceItemRequestBuilder) CreatePatchRequestInformation(options *AccessReviewInstanceItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Decisions the decisions property
func (m *AccessReviewInstanceItemRequestBuilder) Decisions()(*ibcccde7c49f5c5bbcb5ceea96760dde9fbef276461349ba11507419b629f04e6.DecisionsRequestBuilder) {
    return ibcccde7c49f5c5bbcb5ceea96760dde9fbef276461349ba11507419b629f04e6.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.pendingAccessReviewInstances.item.decisions.item collection
func (m *AccessReviewInstanceItemRequestBuilder) DecisionsById(id string)(*i0b5d9b80b48fdacc03dda0d4899b47787af4703b66ab3b6c99f6c290394d6f15.AccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem%2Did"] = id
    }
    return i0b5d9b80b48fdacc03dda0d4899b47787af4703b66ab3b6c99f6c290394d6f15.NewAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Definition the definition property
func (m *AccessReviewInstanceItemRequestBuilder) Definition()(*i29e3e31b50fe024cbfcedaf85025abb3ad402366b9a1338c8b8bc0d397a92401.DefinitionRequestBuilder) {
    return i29e3e31b50fe024cbfcedaf85025abb3ad402366b9a1338c8b8bc0d397a92401.NewDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property pendingAccessReviewInstances for users
func (m *AccessReviewInstanceItemRequestBuilder) Delete(options *AccessReviewInstanceItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get navigation property to get list of access reviews pending approval by reviewer.
func (m *AccessReviewInstanceItemRequestBuilder) Get(options *AccessReviewInstanceItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable), nil
}
// Patch update the navigation property pendingAccessReviewInstances in users
func (m *AccessReviewInstanceItemRequestBuilder) Patch(options *AccessReviewInstanceItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ResetDecisions the resetDecisions property
func (m *AccessReviewInstanceItemRequestBuilder) ResetDecisions()(*i2e9774ef796277adf2753e94982cf69a710728251ccf90259642dcf552172ca8.ResetDecisionsRequestBuilder) {
    return i2e9774ef796277adf2753e94982cf69a710728251ccf90259642dcf552172ca8.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendReminder the sendReminder property
func (m *AccessReviewInstanceItemRequestBuilder) SendReminder()(*i8ef41e575ff6010a2d837801cb7f0c7340ef83813e02674cb4ca79971d31a2a2.SendReminderRequestBuilder) {
    return i8ef41e575ff6010a2d837801cb7f0c7340ef83813e02674cb4ca79971d31a2a2.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stages the stages property
func (m *AccessReviewInstanceItemRequestBuilder) Stages()(*i2bfb15080611c6b6292e90319c54737f8b19d678a61e405dd1bd25c0d3aa9c4e.StagesRequestBuilder) {
    return i2bfb15080611c6b6292e90319c54737f8b19d678a61e405dd1bd25c0d3aa9c4e.NewStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.pendingAccessReviewInstances.item.stages.item collection
func (m *AccessReviewInstanceItemRequestBuilder) StagesById(id string)(*i7b1e883f98e741ec9b8bb6f1e9c089fcc58af710068f93b642d83ad60fa83c77.AccessReviewStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage%2Did"] = id
    }
    return i7b1e883f98e741ec9b8bb6f1e9c089fcc58af710068f93b642d83ad60fa83c77.NewAccessReviewStageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Stop the stop property
func (m *AccessReviewInstanceItemRequestBuilder) Stop()(*i5b136c8ab95b62c794f13ee8a48c92d1c62d5dc6c75f5e4dc5bf16a4aeb7a432.StopRequestBuilder) {
    return i5b136c8ab95b62c794f13ee8a48c92d1c62d5dc6c75f5e4dc5bf16a4aeb7a432.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
