package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
type MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetQueryParameters the list of group assignemenets for app configration.
type MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetQueryParameters
}
// MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderInternal instantiates a new MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder and sets the default values.
func NewMobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) {
    m := &MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration%2Did}/assignments/{managedDeviceMobileAppConfigurationAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder instantiates a new MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder and sets the default values.
func NewMobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of group assignemenets for app configration.
// returns a ManagedDeviceMobileAppConfigurationAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationAssignmentable), nil
}
// Patch update the navigation property assignments in deviceAppManagement
// returns a ManagedDeviceMobileAppConfigurationAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationAssignmentable, requestConfiguration *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationAssignmentable), nil
}
// ToDeleteRequestInformation delete navigation property assignments for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration%2Did}/assignments/{managedDeviceMobileAppConfigurationAssignment%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of group assignemenets for app configration.
// returns a *RequestInformation when successful
func (m *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationAssignmentable, requestConfiguration *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration%2Did}/assignments/{managedDeviceMobileAppConfigurationAssignment%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder when successful
func (m *MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*MobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder) {
    return NewMobileAppConfigurationsItemAssignmentsManagedDeviceMobileAppConfigurationAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
