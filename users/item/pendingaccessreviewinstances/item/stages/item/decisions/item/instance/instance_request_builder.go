package instance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2b6e639b85012916127dc0d07246eb1ae49de6d165778c5f395e700affdd9ef1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/definition"
    i4201497ed84875eef39ef0e9f79d2e857848f718dba808ab7cdb040dfdd886e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/stop"
    i782c49834b7b944c9386d26b1447e1203c6c78b9faf1b852a88f039a1d2dcd01 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/batchrecorddecisions"
    i7f164e80748f8254d77062eb04b9fc70d24acae1f8ff55025019ff39b93f807f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/acceptrecommendations"
    ib32fea3d4d08f37811492a78820bd10fd5ed6a2b6ebe324905ab1721dcf9468d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/applydecisions"
    ibc590043d3d2dc48e7d0a988e65b8d6ce44a64055edec92bbfb7f460a581677d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/resetdecisions"
    id55ed28e16edbddd855986652f8486f24b2ce428790808ae55a7be780db9fb0e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/sendreminder"
    ie8f8ab4cd8e535acc957f1d84dfe870b17611bf063d1c7001c04bc5926b90f0a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/contactedreviewers"
    if1bc878fe4076da3e59089857683f38dcd1d7348a735eaf6d2fa215d83f96017 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/decisions"
    i029ca4e7f214bf663eda08e436807f3035ef57d28cc79e27c6d721f043965f7c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/decisions/item"
    iebc8a00efa2b50bd2e1e9d3b9cf13585be8d95411d38a44a5f752a1cdd32ea2b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/contactedreviewers/item"
)

// InstanceRequestBuilder provides operations to manage the instance property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
type InstanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InstanceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InstanceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InstanceRequestBuilderGetQueryParameters there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
type InstanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// InstanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InstanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InstanceRequestBuilderGetQueryParameters
}
// InstanceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InstanceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptRecommendations provides operations to call the acceptRecommendations method.
func (m *InstanceRequestBuilder) AcceptRecommendations()(*i7f164e80748f8254d77062eb04b9fc70d24acae1f8ff55025019ff39b93f807f.AcceptRecommendationsRequestBuilder) {
    return i7f164e80748f8254d77062eb04b9fc70d24acae1f8ff55025019ff39b93f807f.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyDecisions provides operations to call the applyDecisions method.
func (m *InstanceRequestBuilder) ApplyDecisions()(*ib32fea3d4d08f37811492a78820bd10fd5ed6a2b6ebe324905ab1721dcf9468d.ApplyDecisionsRequestBuilder) {
    return ib32fea3d4d08f37811492a78820bd10fd5ed6a2b6ebe324905ab1721dcf9468d.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BatchRecordDecisions provides operations to call the batchRecordDecisions method.
func (m *InstanceRequestBuilder) BatchRecordDecisions()(*i782c49834b7b944c9386d26b1447e1203c6c78b9faf1b852a88f039a1d2dcd01.BatchRecordDecisionsRequestBuilder) {
    return i782c49834b7b944c9386d26b1447e1203c6c78b9faf1b852a88f039a1d2dcd01.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewInstanceRequestBuilderInternal instantiates a new InstanceRequestBuilder and sets the default values.
func NewInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstanceRequestBuilder) {
    m := &InstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInstanceRequestBuilder instantiates a new InstanceRequestBuilder and sets the default values.
func NewInstanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstanceRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *InstanceRequestBuilder) ContactedReviewers()(*ie8f8ab4cd8e535acc957f1d84dfe870b17611bf063d1c7001c04bc5926b90f0a.ContactedReviewersRequestBuilder) {
    return ie8f8ab4cd8e535acc957f1d84dfe870b17611bf063d1c7001c04bc5926b90f0a.NewContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
func (m *InstanceRequestBuilder) ContactedReviewersById(id string)(*iebc8a00efa2b50bd2e1e9d3b9cf13585be8d95411d38a44a5f752a1cdd32ea2b.AccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return iebc8a00efa2b50bd2e1e9d3b9cf13585be8d95411d38a44a5f752a1cdd32ea2b.NewAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property instance for users
func (m *InstanceRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *InstanceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *InstanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property instance in users
func (m *InstanceRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *InstanceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Decisions provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
func (m *InstanceRequestBuilder) Decisions()(*if1bc878fe4076da3e59089857683f38dcd1d7348a735eaf6d2fa215d83f96017.DecisionsRequestBuilder) {
    return if1bc878fe4076da3e59089857683f38dcd1d7348a735eaf6d2fa215d83f96017.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
func (m *InstanceRequestBuilder) DecisionsById(id string)(*i029ca4e7f214bf663eda08e436807f3035ef57d28cc79e27c6d721f043965f7c.AccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem%2Did1"] = id
    }
    return i029ca4e7f214bf663eda08e436807f3035ef57d28cc79e27c6d721f043965f7c.NewAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Definition provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
func (m *InstanceRequestBuilder) Definition()(*i2b6e639b85012916127dc0d07246eb1ae49de6d165778c5f395e700affdd9ef1.DefinitionRequestBuilder) {
    return i2b6e639b85012916127dc0d07246eb1ae49de6d165778c5f395e700affdd9ef1.NewDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instance for users
func (m *InstanceRequestBuilder) Delete(ctx context.Context, requestConfiguration *InstanceRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) Get(ctx context.Context, requestConfiguration *InstanceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable), nil
}
// Patch update the navigation property instance in users
func (m *InstanceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *InstanceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable), nil
}
// ResetDecisions provides operations to call the resetDecisions method.
func (m *InstanceRequestBuilder) ResetDecisions()(*ibc590043d3d2dc48e7d0a988e65b8d6ce44a64055edec92bbfb7f460a581677d.ResetDecisionsRequestBuilder) {
    return ibc590043d3d2dc48e7d0a988e65b8d6ce44a64055edec92bbfb7f460a581677d.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendReminder provides operations to call the sendReminder method.
func (m *InstanceRequestBuilder) SendReminder()(*id55ed28e16edbddd855986652f8486f24b2ce428790808ae55a7be780db9fb0e.SendReminderRequestBuilder) {
    return id55ed28e16edbddd855986652f8486f24b2ce428790808ae55a7be780db9fb0e.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stop provides operations to call the stop method.
func (m *InstanceRequestBuilder) Stop()(*i4201497ed84875eef39ef0e9f79d2e857848f718dba808ab7cdb040dfdd886e8.StopRequestBuilder) {
    return i4201497ed84875eef39ef0e9f79d2e857848f718dba808ab7cdb040dfdd886e8.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
