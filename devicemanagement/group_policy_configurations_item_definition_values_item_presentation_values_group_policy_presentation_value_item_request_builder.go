package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder provides operations to manage the presentationValues property of the microsoft.graph.groupPolicyDefinitionValue entity.
type GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetQueryParameters the associated group policy presentation values with the definition value.
type GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetQueryParameters
}
// GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderInternal instantiates a new GroupPolicyPresentationValueItemRequestBuilder and sets the default values.
func NewGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) {
    m := &GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration%2Did}/definitionValues/{groupPolicyDefinitionValue%2Did}/presentationValues/{groupPolicyPresentationValue%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder instantiates a new GroupPolicyPresentationValueItemRequestBuilder and sets the default values.
func NewGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property presentationValues for deviceManagement
func (m *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the associated group policy presentation values with the definition value.
func (m *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property presentationValues in deviceManagement
func (m *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable, requestConfiguration *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DefinitionValue provides operations to manage the definitionValue property of the microsoft.graph.groupPolicyPresentationValue entity.
func (m *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) DefinitionValue()(*GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilder) {
    return NewGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property presentationValues for deviceManagement
func (m *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the associated group policy presentation values with the definition value.
func (m *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable), nil
}
// Patch update the navigation property presentationValues in deviceManagement
func (m *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable, requestConfiguration *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable), nil
}
// Presentation provides operations to manage the presentation property of the microsoft.graph.groupPolicyPresentationValue entity.
func (m *GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) Presentation()(*GroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemPresentationRequestBuilder) {
    return NewGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemPresentationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
