package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder provides operations to manage the createdBy property of the microsoft.graph.identityGovernance.customTaskExtension entity.
type LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetQueryParameters the unique identifier of the Microsoft Entra user that created the custom task extension.Supports $filter(eq, ne) and $expand.
type LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderInternal instantiates a new CreatedByRequestBuilder and sets the default values.
func NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder) {
    m := &LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/customTaskExtensions/{customTaskExtension%2Did}/createdBy{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder instantiates a new CreatedByRequestBuilder and sets the default values.
func NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the unique identifier of the Microsoft Entra user that created the custom task extension.Supports $filter(eq, ne) and $expand.
func (m *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder) MailboxSettings()(*LifecycleWorkflowsCustomTaskExtensionsItemCreatedByMailboxSettingsRequestBuilder) {
    return NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
func (m *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder) ServiceProvisioningErrors()(*LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder) {
    return NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the unique identifier of the Microsoft Entra user that created the custom task extension.Supports $filter(eq, ne) and $expand.
func (m *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder) WithUrl(rawUrl string)(*LifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder) {
    return NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
