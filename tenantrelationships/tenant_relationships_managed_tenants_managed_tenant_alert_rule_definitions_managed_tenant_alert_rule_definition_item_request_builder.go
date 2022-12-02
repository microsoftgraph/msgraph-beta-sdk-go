package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder provides operations to manage the managedTenantAlertRuleDefinitions property of the microsoft.graph.managedTenants.managedTenant entity.
type TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderGetQueryParameters get managedTenantAlertRuleDefinitions from tenantRelationships
type TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderGetQueryParameters
}
// TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AlertRules provides operations to manage the alertRules property of the microsoft.graph.managedTenants.managedTenantAlertRuleDefinition entity.
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder) AlertRules()(*TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsItemAlertRulesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsItemAlertRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AlertRulesById provides operations to manage the alertRules property of the microsoft.graph.managedTenants.managedTenantAlertRuleDefinition entity.
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder) AlertRulesById(id string)(*TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsItemAlertRulesManagedTenantAlertRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertRule%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsItemAlertRulesManagedTenantAlertRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderInternal instantiates a new ManagedTenantAlertRuleDefinitionItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder) {
    m := &TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder{
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
// NewTenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder instantiates a new ManagedTenantAlertRuleDefinitionItemRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedTenantAlertRuleDefinitions for tenantRelationships
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleDefinitionable, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleDefinitionable, error) {
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
func (m *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleDefinitionable, requestConfiguration *TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertRuleDefinitionable, error) {
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
