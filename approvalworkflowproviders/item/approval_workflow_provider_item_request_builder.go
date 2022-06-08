package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i00b6268726a46fcf8e15e0ca3c5f8889dfdd53aaf48fe4ce2411243820496cf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/businessflows"
    i30206120795d8fda84ba9e547a109b67da276500cec13099585adedbf956e532 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/businessflowswithrequestsawaitingmydecision"
    i31436781316332b768fad45035d32f5befd814379c0ccbd0c58fb23e0e08e107 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/policytemplates"
    iad8cec4aa318577e66fff21478bdf279aabc1597d98a1031afe46ae7262c9daa "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/policytemplates/item"
    icdd8feb64f3d3f11cf1c42a81a0ec83a900ebf29181eae132e78625db81a1fa7 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/businessflows/item"
    id3e3e89c17dc1203ba2dba0acce89454c2eacc6063d8b15b83854c381c75fc28 "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders/item/businessflowswithrequestsawaitingmydecision/item"
)

// ApprovalWorkflowProviderItemRequestBuilder provides operations to manage the collection of approvalWorkflowProvider entities.
type ApprovalWorkflowProviderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ApprovalWorkflowProviderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProviderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApprovalWorkflowProviderItemRequestBuilderGetQueryParameters get entity from approvalWorkflowProviders by key
type ApprovalWorkflowProviderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ApprovalWorkflowProviderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProviderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ApprovalWorkflowProviderItemRequestBuilderGetQueryParameters
}
// ApprovalWorkflowProviderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalWorkflowProviderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BusinessFlows the businessFlows property
func (m *ApprovalWorkflowProviderItemRequestBuilder) BusinessFlows()(*i00b6268726a46fcf8e15e0ca3c5f8889dfdd53aaf48fe4ce2411243820496cf2.BusinessFlowsRequestBuilder) {
    return i00b6268726a46fcf8e15e0ca3c5f8889dfdd53aaf48fe4ce2411243820496cf2.NewBusinessFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BusinessFlowsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.approvalWorkflowProviders.item.businessFlows.item collection
func (m *ApprovalWorkflowProviderItemRequestBuilder) BusinessFlowsById(id string)(*icdd8feb64f3d3f11cf1c42a81a0ec83a900ebf29181eae132e78625db81a1fa7.BusinessFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["businessFlow%2Did"] = id
    }
    return icdd8feb64f3d3f11cf1c42a81a0ec83a900ebf29181eae132e78625db81a1fa7.NewBusinessFlowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// BusinessFlowsWithRequestsAwaitingMyDecision the businessFlowsWithRequestsAwaitingMyDecision property
func (m *ApprovalWorkflowProviderItemRequestBuilder) BusinessFlowsWithRequestsAwaitingMyDecision()(*i30206120795d8fda84ba9e547a109b67da276500cec13099585adedbf956e532.BusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilder) {
    return i30206120795d8fda84ba9e547a109b67da276500cec13099585adedbf956e532.NewBusinessFlowsWithRequestsAwaitingMyDecisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BusinessFlowsWithRequestsAwaitingMyDecisionById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.approvalWorkflowProviders.item.businessFlowsWithRequestsAwaitingMyDecision.item collection
func (m *ApprovalWorkflowProviderItemRequestBuilder) BusinessFlowsWithRequestsAwaitingMyDecisionById(id string)(*id3e3e89c17dc1203ba2dba0acce89454c2eacc6063d8b15b83854c381c75fc28.BusinessFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["businessFlow%2Did"] = id
    }
    return id3e3e89c17dc1203ba2dba0acce89454c2eacc6063d8b15b83854c381c75fc28.NewBusinessFlowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewApprovalWorkflowProviderItemRequestBuilderInternal instantiates a new ApprovalWorkflowProviderItemRequestBuilder and sets the default values.
func NewApprovalWorkflowProviderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalWorkflowProviderItemRequestBuilder) {
    m := &ApprovalWorkflowProviderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/approvalWorkflowProviders/{approvalWorkflowProvider%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApprovalWorkflowProviderItemRequestBuilder instantiates a new ApprovalWorkflowProviderItemRequestBuilder and sets the default values.
func NewApprovalWorkflowProviderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalWorkflowProviderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApprovalWorkflowProviderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from approvalWorkflowProviders
func (m *ApprovalWorkflowProviderItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete entity from approvalWorkflowProviders
func (m *ApprovalWorkflowProviderItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ApprovalWorkflowProviderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from approvalWorkflowProviders by key
func (m *ApprovalWorkflowProviderItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get entity from approvalWorkflowProviders by key
func (m *ApprovalWorkflowProviderItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ApprovalWorkflowProviderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in approvalWorkflowProviders
func (m *ApprovalWorkflowProviderItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update entity in approvalWorkflowProviders
func (m *ApprovalWorkflowProviderItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable, requestConfiguration *ApprovalWorkflowProviderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete entity from approvalWorkflowProviders
func (m *ApprovalWorkflowProviderItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete entity from approvalWorkflowProviders
func (m *ApprovalWorkflowProviderItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ApprovalWorkflowProviderItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from approvalWorkflowProviders by key
func (m *ApprovalWorkflowProviderItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get entity from approvalWorkflowProviders by key
func (m *ApprovalWorkflowProviderItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ApprovalWorkflowProviderItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalWorkflowProviderFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable), nil
}
// Patch update entity in approvalWorkflowProviders
func (m *ApprovalWorkflowProviderItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update entity in approvalWorkflowProviders
func (m *ApprovalWorkflowProviderItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalWorkflowProviderable, requestConfiguration *ApprovalWorkflowProviderItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// PolicyTemplates the policyTemplates property
func (m *ApprovalWorkflowProviderItemRequestBuilder) PolicyTemplates()(*i31436781316332b768fad45035d32f5befd814379c0ccbd0c58fb23e0e08e107.PolicyTemplatesRequestBuilder) {
    return i31436781316332b768fad45035d32f5befd814379c0ccbd0c58fb23e0e08e107.NewPolicyTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PolicyTemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.approvalWorkflowProviders.item.policyTemplates.item collection
func (m *ApprovalWorkflowProviderItemRequestBuilder) PolicyTemplatesById(id string)(*iad8cec4aa318577e66fff21478bdf279aabc1597d98a1031afe46ae7262c9daa.GovernancePolicyTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governancePolicyTemplate%2Did"] = id
    }
    return iad8cec4aa318577e66fff21478bdf279aabc1597d98a1031afe46ae7262c9daa.NewGovernancePolicyTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
