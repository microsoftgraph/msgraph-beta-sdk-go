package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.deviceAppManagement entity.
type WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetQueryParameters windows information protection device registrations that are not MDM enrolled.
type WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetQueryParameters
}
// WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal instantiates a new WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    m := &WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsInformationProtectionDeviceRegistrations/{windowsInformationProtectionDeviceRegistration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder instantiates a new WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsInformationProtectionDeviceRegistrations for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get windows information protection device registrations that are not MDM enrolled.
// returns a WindowsInformationProtectionDeviceRegistrationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsInformationProtectionDeviceRegistrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationable), nil
}
// Patch update the navigation property windowsInformationProtectionDeviceRegistrations in deviceAppManagement
// returns a WindowsInformationProtectionDeviceRegistrationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationable, requestConfiguration *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsInformationProtectionDeviceRegistrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationable), nil
}
// ToDeleteRequestInformation delete navigation property windowsInformationProtectionDeviceRegistrations for deviceAppManagement
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation windows information protection device registrations that are not MDM enrolled.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsInformationProtectionDeviceRegistrations in deviceAppManagement
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationable, requestConfiguration *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Wipe provides operations to call the wipe method.
// returns a *WindowsinformationprotectiondeviceregistrationsItemWipeRequestBuilder when successful
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) Wipe()(*WindowsinformationprotectiondeviceregistrationsItemWipeRequestBuilder) {
    return NewWindowsinformationprotectiondeviceregistrationsItemWipeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder when successful
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) WithUrl(rawUrl string)(*WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    return NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
