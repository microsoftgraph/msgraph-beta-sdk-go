package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder provides operations to manage the lastModifiedBy property of the microsoft.graph.identityGovernance.workflowBase entity.
type LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilderGetQueryParameters the user who last modified the workflow.
type LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilderInternal instantiates a new LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder) {
    m := &LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/versions/{workflowVersion%2DversionNumber}/lastModifiedBy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder instantiates a new LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the user who last modified the workflow.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// MailboxSettings the mailboxSettings property
// returns a *LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByMailboxSettingsRequestBuilder when successful
func (m *LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder) MailboxSettings()(*LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByMailboxSettingsRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByServiceProvisioningErrorsRequestBuilder when successful
func (m *LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder) ServiceProvisioningErrors()(*LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByServiceProvisioningErrorsRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the user who last modified the workflow.
// returns a *RequestInformation when successful
func (m *LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder when successful
func (m *LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder) WithUrl(rawUrl string)(*LifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemVersionsItemLastModifiedByRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
