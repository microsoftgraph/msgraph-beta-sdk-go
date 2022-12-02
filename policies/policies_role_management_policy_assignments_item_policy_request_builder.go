package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilder provides operations to manage the policy property of the microsoft.graph.unifiedRoleManagementPolicyAssignment entity.
type PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilderGetQueryParameters the policy that's associated with a policy assignment. Supports $expand and a nested $expand of the rules and effectiveRules relationships for the policy.
type PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilderGetQueryParameters
}
// NewPoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilderInternal instantiates a new PolicyRequestBuilder and sets the default values.
func NewPoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilder) {
    m := &PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/roleManagementPolicyAssignments/{unifiedRoleManagementPolicyAssignment%2Did}/policy{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilder instantiates a new PolicyRequestBuilder and sets the default values.
func NewPoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the policy that's associated with a policy assignment. Supports $expand and a nested $expand of the rules and effectiveRules relationships for the policy.
func (m *PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get the policy that's associated with a policy assignment. Supports $expand and a nested $expand of the rules and effectiveRules relationships for the policy.
func (m *PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *PoliciesRoleManagementPolicyAssignmentsItemPolicyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyable), nil
}
