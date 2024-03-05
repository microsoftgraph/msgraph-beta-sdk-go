package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder provides operations to manage the matchedDevices property of the microsoft.graph.windowsUpdates.applicableContent entity.
type WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderGetQueryParameters collection of devices and recommendations for applicable catalog content.
type WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderGetQueryParameters
}
// WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderInternal instantiates a new WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder) {
    m := &WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/applicableContent/{applicableContent%2DcatalogEntryId}/matchedDevices/{applicableContentDeviceMatch%2DdeviceId}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder instantiates a new WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property matchedDevices for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of devices and recommendations for applicable catalog content.
// returns a ApplicableContentDeviceMatchable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentDeviceMatchable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateApplicableContentDeviceMatchFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentDeviceMatchable), nil
}
// Patch update the navigation property matchedDevices in admin
// returns a ApplicableContentDeviceMatchable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentDeviceMatchable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentDeviceMatchable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateApplicableContentDeviceMatchFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentDeviceMatchable), nil
}
// ToDeleteRequestInformation delete navigation property matchedDevices for admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/applicableContent/{applicableContent%2DcatalogEntryId}/matchedDevices/{applicableContentDeviceMatch%2DdeviceId}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of devices and recommendations for applicable catalog content.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property matchedDevices in admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentDeviceMatchable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/applicableContent/{applicableContent%2DcatalogEntryId}/matchedDevices/{applicableContentDeviceMatch%2DdeviceId}", m.BaseRequestBuilder.PathParameters)
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
// returns a *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesApplicableContentDeviceMatchDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
