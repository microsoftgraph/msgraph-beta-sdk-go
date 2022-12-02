package officeconfiguration

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder provides operations to manage the clientConfigurations property of the microsoft.graph.officeConfiguration entity.
type OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderGetQueryParameters list of office Client configuration.
type OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderGetQueryParameters
}
// OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder) Assign()(*OfficeConfigurationClientConfigurationsItemAssignRequestBuilder) {
    return NewOfficeConfigurationClientConfigurationsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.officeClientConfiguration entity.
func (m *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder) Assignments()(*OfficeConfigurationClientConfigurationsItemAssignmentsRequestBuilder) {
    return NewOfficeConfigurationClientConfigurationsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.officeClientConfiguration entity.
func (m *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder) AssignmentsById(id string)(*OfficeConfigurationClientConfigurationsItemAssignmentsOfficeClientConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["officeClientConfigurationAssignment%2Did"] = id
    }
    return NewOfficeConfigurationClientConfigurationsItemAssignmentsOfficeClientConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderInternal instantiates a new OfficeClientConfigurationItemRequestBuilder and sets the default values.
func NewOfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder) {
    m := &OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/officeConfiguration/clientConfigurations/{officeClientConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder instantiates a new OfficeClientConfigurationItemRequestBuilder and sets the default values.
func NewOfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property clientConfigurations for officeConfiguration
func (m *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation list of office Client configuration.
func (m *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property clientConfigurations in officeConfiguration
func (m *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeClientConfigurationable, requestConfiguration *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property clientConfigurations for officeConfiguration
func (m *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of office Client configuration.
func (m *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeClientConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOfficeClientConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeClientConfigurationable), nil
}
// Patch update the navigation property clientConfigurations in officeConfiguration
func (m *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeClientConfigurationable, requestConfiguration *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeClientConfigurationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOfficeClientConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeClientConfigurationable), nil
}
// PolicyPayload provides operations to manage the media for the officeConfiguration entity.
func (m *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder) PolicyPayload()(*OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilder) {
    return NewOfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserPreferencePayload provides operations to manage the media for the officeConfiguration entity.
func (m *OfficeConfigurationClientConfigurationsOfficeClientConfigurationItemRequestBuilder) UserPreferencePayload()(*OfficeConfigurationClientConfigurationsItemUserPreferencePayloadRequestBuilder) {
    return NewOfficeConfigurationClientConfigurationsItemUserPreferencePayloadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
