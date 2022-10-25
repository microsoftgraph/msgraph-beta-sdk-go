package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
    i8e0839a962a6be8ea2c2523eda6c17f4e508aac07b0756d69521b54bc6b30589 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalertruledefinitions/item/alertrules"
    i4fd0c0bbec3ca7822b7b6c6477ea4293c0faf2ceaef18a16e41d253b868e37e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managedtenantalertruledefinitions/item/alertrules/item"
)

// ManagedTenantAlertRuleDefinitionItemRequestBuilder provides operations to manage the managedTenantAlertRuleDefinitions property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedTenantAlertRuleDefinitionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedTenantAlertRuleDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantAlertRuleDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedTenantAlertRuleDefinitionItemRequestBuilderGetQueryParameters get managedTenantAlertRuleDefinitions from tenantRelationships
type ManagedTenantAlertRuleDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantAlertRuleDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantAlertRuleDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantAlertRuleDefinitionItemRequestBuilderGetQueryParameters
}
// ManagedTenantAlertRuleDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantAlertRuleDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlertRules the alertRules property
func (m *ManagedTenantAlertRuleDefinitionItemRequestBuilder) AlertRules()(*i8e0839a962a6be8ea2c2523eda6c17f4e508aac07b0756d69521b54bc6b30589.AlertRulesRequestBuilder) {
    return i8e0839a962a6be8ea2c2523eda6c17f4e508aac07b0756d69521b54bc6b30589.NewAlertRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AlertRulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.managedTenantAlertRuleDefinitions.item.alertRules.item collection
func (m *ManagedTenantAlertRuleDefinitionItemRequestBuilder) AlertRulesById(id string)(*i4fd0c0bbec3ca7822b7b6c6477ea4293c0faf2ceaef18a16e41d253b868e37e9.ManagedTenantAlertRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertRule%2Did"] = id
    }
    return i4fd0c0bbec3ca7822b7b6c6477ea4293c0faf2ceaef18a16e41d253b868e37e9.NewManagedTenantAlertRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedTenantAlertRuleDefinitionItemRequestBuilderInternal instantiates a new ManagedTenantAlertRuleDefinitionItemRequestBuilder and sets the default values.
func NewManagedTenantAlertRuleDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantAlertRuleDefinitionItemRequestBuilder) {
    m := &ManagedTenantAlertRuleDefinitionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlertRuleDefinitions/{managedTenantAlertRuleDefinition%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedTenantAlertRuleDefinitionItemRequestBuilder instantiates a new ManagedTenantAlertRuleDefinitionItemRequestBuilder and sets the default values.
func NewManagedTenantAlertRuleDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantAlertRuleDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantAlertRuleDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedTenantAlertRuleDefinitions for tenantRelationships
func (m *ManagedTenantAlertRuleDefinitionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantAlertRuleDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get managedTenantAlertRuleDefinitions from tenantRelationships
func (m *ManagedTenantAlertRuleDefinitionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantAlertRuleDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property managedTenantAlertRuleDefinitions in tenantRelationships
func (m *ManagedTenantAlertRuleDefinitionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleDefinitionable, requestConfiguration *ManagedTenantAlertRuleDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property managedTenantAlertRuleDefinitions for tenantRelationships
func (m *ManagedTenantAlertRuleDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedTenantAlertRuleDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get managedTenantAlertRuleDefinitions from tenantRelationships
func (m *ManagedTenantAlertRuleDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantAlertRuleDefinitionItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleDefinitionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertRuleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleDefinitionable), nil
}
// Patch update the navigation property managedTenantAlertRuleDefinitions in tenantRelationships
func (m *ManagedTenantAlertRuleDefinitionItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleDefinitionable, requestConfiguration *ManagedTenantAlertRuleDefinitionItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleDefinitionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertRuleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleDefinitionable), nil
}
