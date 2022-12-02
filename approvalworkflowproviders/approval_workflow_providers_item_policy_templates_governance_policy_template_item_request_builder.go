package approvalworkflowproviders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder provides operations to manage the policyTemplates property of the microsoft.graph.approvalWorkflowProvider entity.
type ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderGetQueryParameters get policyTemplates from approvalWorkflowProviders
type ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderGetQueryParameters
}
// ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderInternal instantiates a new GovernancePolicyTemplateItemRequestBuilder and sets the default values.
func NewApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder) {
    m := &ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/approvalWorkflowProviders/{approvalWorkflowProvider%2Did}/policyTemplates/{governancePolicyTemplate%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder instantiates a new GovernancePolicyTemplateItemRequestBuilder and sets the default values.
func NewApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property policyTemplates for approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get policyTemplates from approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property policyTemplates in approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernancePolicyTemplateable, requestConfiguration *ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property policyTemplates for approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get policyTemplates from approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernancePolicyTemplateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernancePolicyTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernancePolicyTemplateable), nil
}
// Patch update the navigation property policyTemplates in approvalWorkflowProviders
func (m *ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernancePolicyTemplateable, requestConfiguration *ApprovalWorkflowProvidersItemPolicyTemplatesGovernancePolicyTemplateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernancePolicyTemplateable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernancePolicyTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernancePolicyTemplateable), nil
}
