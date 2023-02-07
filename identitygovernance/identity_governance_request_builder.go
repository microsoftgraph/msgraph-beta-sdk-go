package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceRequestBuilder provides operations to manage the identityGovernance singleton.
type IdentityGovernanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceRequestBuilderGetQueryParameters get identityGovernance
type IdentityGovernanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceRequestBuilderGetQueryParameters
}
// IdentityGovernanceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviews provides operations to manage the accessReviews property of the microsoft.graph.identityGovernance entity.
func (m *IdentityGovernanceRequestBuilder) AccessReviews()(*AccessReviewsRequestBuilder) {
    return NewAccessReviewsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AppConsent provides operations to manage the appConsent property of the microsoft.graph.identityGovernance entity.
func (m *IdentityGovernanceRequestBuilder) AppConsent()(*AppConsentRequestBuilder) {
    return NewAppConsentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewIdentityGovernanceRequestBuilderInternal instantiates a new IdentityGovernanceRequestBuilder and sets the default values.
func NewIdentityGovernanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    m := &IdentityGovernanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewIdentityGovernanceRequestBuilder instantiates a new IdentityGovernanceRequestBuilder and sets the default values.
func NewIdentityGovernanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceRequestBuilderInternal(urlParams, requestAdapter)
}
// EntitlementManagement provides operations to manage the entitlementManagement property of the microsoft.graph.identityGovernance entity.
func (m *IdentityGovernanceRequestBuilder) EntitlementManagement()(*EntitlementManagementRequestBuilder) {
    return NewEntitlementManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get get identityGovernance
func (m *IdentityGovernanceRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityGovernanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable), nil
}
// LifecycleWorkflows provides operations to manage the lifecycleWorkflows property of the microsoft.graph.identityGovernance entity.
func (m *IdentityGovernanceRequestBuilder) LifecycleWorkflows()(*LifecycleWorkflowsRequestBuilder) {
    return NewLifecycleWorkflowsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update identityGovernance
func (m *IdentityGovernanceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable, requestConfiguration *IdentityGovernanceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityGovernanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable), nil
}
// PrivilegedAccess provides operations to manage the privilegedAccess property of the microsoft.graph.identityGovernance entity.
func (m *IdentityGovernanceRequestBuilder) PrivilegedAccess()(*PrivilegedAccessRequestBuilder) {
    return NewPrivilegedAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TermsOfUse provides operations to manage the termsOfUse property of the microsoft.graph.identityGovernance entity.
func (m *IdentityGovernanceRequestBuilder) TermsOfUse()(*TermsOfUseRequestBuilder) {
    return NewTermsOfUseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation get identityGovernance
func (m *IdentityGovernanceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update identityGovernance
func (m *IdentityGovernanceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable, requestConfiguration *IdentityGovernanceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
