package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
type IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetQueryParameters the associated group assignments for IosLobAppProvisioningConfiguration.
type IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetQueryParameters
}
// IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderInternal instantiates a new IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder and sets the default values.
func NewIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) {
    m := &IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration%2Did}/assignments/{iosLobAppProvisioningConfigurationAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder instantiates a new IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder and sets the default values.
func NewIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the associated group assignments for IosLobAppProvisioningConfiguration.
// returns a IosLobAppProvisioningConfigurationAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosLobAppProvisioningConfigurationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable), nil
}
// Patch update the navigation property assignments in deviceAppManagement
// returns a IosLobAppProvisioningConfigurationAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable, requestConfiguration *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosLobAppProvisioningConfigurationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable), nil
}
// ToDeleteRequestInformation delete navigation property assignments for deviceAppManagement
// returns a *RequestInformation when successful
func (m *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the associated group assignments for IosLobAppProvisioningConfiguration.
// returns a *RequestInformation when successful
func (m *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignments in deviceAppManagement
// returns a *RequestInformation when successful
func (m *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable, requestConfiguration *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder when successful
func (m *IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*IosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
