package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceAccessReviewsRequestBuilder provides operations to manage the accessReviews property of the microsoft.graph.identityGovernance entity.
type IdentityGovernanceAccessReviewsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceAccessReviewsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceAccessReviewsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceAccessReviewsRequestBuilderGetQueryParameters get accessReviews from identityGovernance
type IdentityGovernanceAccessReviewsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceAccessReviewsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceAccessReviewsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceAccessReviewsRequestBuilderGetQueryParameters
}
// IdentityGovernanceAccessReviewsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceAccessReviewsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityGovernanceAccessReviewsRequestBuilderInternal instantiates a new AccessReviewsRequestBuilder and sets the default values.
func NewIdentityGovernanceAccessReviewsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceAccessReviewsRequestBuilder) {
    m := &IdentityGovernanceAccessReviewsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceAccessReviewsRequestBuilder instantiates a new AccessReviewsRequestBuilder and sets the default values.
func NewIdentityGovernanceAccessReviewsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceAccessReviewsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceAccessReviewsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessReviews for identityGovernance
func (m *IdentityGovernanceAccessReviewsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceAccessReviewsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get accessReviews from identityGovernance
func (m *IdentityGovernanceAccessReviewsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceAccessReviewsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessReviews in identityGovernance
func (m *IdentityGovernanceAccessReviewsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewSetable, requestConfiguration *IdentityGovernanceAccessReviewsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Decisions provides operations to manage the decisions property of the microsoft.graph.accessReviewSet entity.
func (m *IdentityGovernanceAccessReviewsRequestBuilder) Decisions()(*IdentityGovernanceAccessReviewsDecisionsRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById provides operations to manage the decisions property of the microsoft.graph.accessReviewSet entity.
func (m *IdentityGovernanceAccessReviewsRequestBuilder) DecisionsById(id string)(*IdentityGovernanceAccessReviewsDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem%2Did"] = id
    }
    return NewIdentityGovernanceAccessReviewsDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Definitions provides operations to manage the definitions property of the microsoft.graph.accessReviewSet entity.
func (m *IdentityGovernanceAccessReviewsRequestBuilder) Definitions()(*IdentityGovernanceAccessReviewsDefinitionsRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefinitionsById provides operations to manage the definitions property of the microsoft.graph.accessReviewSet entity.
func (m *IdentityGovernanceAccessReviewsRequestBuilder) DefinitionsById(id string)(*IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewScheduleDefinition%2Did"] = id
    }
    return NewIdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property accessReviews for identityGovernance
func (m *IdentityGovernanceAccessReviewsRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceAccessReviewsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get accessReviews from identityGovernance
func (m *IdentityGovernanceAccessReviewsRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceAccessReviewsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewSetable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewSetable), nil
}
// HistoryDefinitions provides operations to manage the historyDefinitions property of the microsoft.graph.accessReviewSet entity.
func (m *IdentityGovernanceAccessReviewsRequestBuilder) HistoryDefinitions()(*IdentityGovernanceAccessReviewsHistoryDefinitionsRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsHistoryDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HistoryDefinitionsById provides operations to manage the historyDefinitions property of the microsoft.graph.accessReviewSet entity.
func (m *IdentityGovernanceAccessReviewsRequestBuilder) HistoryDefinitionsById(id string)(*IdentityGovernanceAccessReviewsHistoryDefinitionsAccessReviewHistoryDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewHistoryDefinition%2Did"] = id
    }
    return NewIdentityGovernanceAccessReviewsHistoryDefinitionsAccessReviewHistoryDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property accessReviews in identityGovernance
func (m *IdentityGovernanceAccessReviewsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewSetable, requestConfiguration *IdentityGovernanceAccessReviewsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewSetable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewSetable), nil
}
// Policy provides operations to manage the policy property of the microsoft.graph.accessReviewSet entity.
func (m *IdentityGovernanceAccessReviewsRequestBuilder) Policy()(*IdentityGovernanceAccessReviewsPolicyRequestBuilder) {
    return NewIdentityGovernanceAccessReviewsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
