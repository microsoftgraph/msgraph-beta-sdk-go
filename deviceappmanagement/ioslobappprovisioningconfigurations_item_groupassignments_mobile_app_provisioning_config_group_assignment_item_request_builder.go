package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder provides operations to manage the groupAssignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
type IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderGetQueryParameters the associated group assignments.
type IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderGetQueryParameters
}
// IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderInternal instantiates a new IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder and sets the default values.
func NewIoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) {
    m := &IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration%2Did}/groupAssignments/{mobileAppProvisioningConfigGroupAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder instantiates a new IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder and sets the default values.
func NewIoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property groupAssignments for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the associated group assignments.
// returns a MobileAppProvisioningConfigGroupAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppProvisioningConfigGroupAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppProvisioningConfigGroupAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppProvisioningConfigGroupAssignmentable), nil
}
// Patch update the navigation property groupAssignments in deviceAppManagement
// returns a MobileAppProvisioningConfigGroupAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppProvisioningConfigGroupAssignmentable, requestConfiguration *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppProvisioningConfigGroupAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppProvisioningConfigGroupAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppProvisioningConfigGroupAssignmentable), nil
}
// ToDeleteRequestInformation delete navigation property groupAssignments for deviceAppManagement
// returns a *RequestInformation when successful
func (m *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the associated group assignments.
// returns a *RequestInformation when successful
func (m *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property groupAssignments in deviceAppManagement
// returns a *RequestInformation when successful
func (m *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppProvisioningConfigGroupAssignmentable, requestConfiguration *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder when successful
func (m *IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*IoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) {
    return NewIoslobappprovisioningconfigurationsItemGroupassignmentsMobileAppProvisioningConfigGroupAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
