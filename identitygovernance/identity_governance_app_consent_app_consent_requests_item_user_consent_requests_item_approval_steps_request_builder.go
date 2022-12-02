package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilder provides operations to manage the steps property of the microsoft.graph.approval entity.
type IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderGetQueryParameters in Azure AD entitlement management, lists the approvalStep objects associated with an approval object.  This call can be made by an approver, providing the identifier of the access package assignment request.
type IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderGetQueryParameters
}
// IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderInternal instantiates a new StepsRequestBuilder and sets the default values.
func NewIdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilder) {
    m := &IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/appConsent/appConsentRequests/{appConsentRequest%2Did}/userConsentRequests/{userConsentRequest%2Did}/approval/steps{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilder instantiates a new StepsRequestBuilder and sets the default values.
func NewIdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilder) Count()(*IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsCountRequestBuilder) {
    return NewIdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation in Azure AD entitlement management, lists the approvalStep objects associated with an approval object.  This call can be made by an approver, providing the identifier of the access package assignment request.
func (m *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to steps for identityGovernance
func (m *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, requestConfiguration *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get in Azure AD entitlement management, lists the approvalStep objects associated with an approval object.  This call can be made by an approver, providing the identifier of the access package assignment request.
func (m *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalStepCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepCollectionResponseable), nil
}
// Post create new navigation property to steps for identityGovernance
func (m *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, requestConfiguration *IdentityGovernanceAppConsentAppConsentRequestsItemUserConsentRequestsItemApprovalStepsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalStepFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable), nil
}
