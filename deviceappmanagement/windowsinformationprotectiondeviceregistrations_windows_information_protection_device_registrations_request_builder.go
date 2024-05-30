package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.deviceAppManagement entity.
type WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetQueryParameters windows information protection device registrations that are not MDM enrolled.
type WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetQueryParameters
}
// WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWindowsInformationProtectionDeviceRegistrationId provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.deviceAppManagement entity.
// returns a *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder when successful
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) ByWindowsInformationProtectionDeviceRegistrationId(windowsInformationProtectionDeviceRegistrationId string)(*WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsInformationProtectionDeviceRegistrationId != "" {
        urlTplParams["windowsInformationProtectionDeviceRegistration%2Did"] = windowsInformationProtectionDeviceRegistrationId
    }
    return NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal instantiates a new WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder and sets the default values.
func NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    m := &WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsInformationProtectionDeviceRegistrations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder instantiates a new WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder and sets the default values.
func NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsinformationprotectiondeviceregistrationsCountRequestBuilder when successful
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) Count()(*WindowsinformationprotectiondeviceregistrationsCountRequestBuilder) {
    return NewWindowsinformationprotectiondeviceregistrationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get windows information protection device registrations that are not MDM enrolled.
// returns a WindowsInformationProtectionDeviceRegistrationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsInformationProtectionDeviceRegistrationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationCollectionResponseable), nil
}
// Post create new navigation property to windowsInformationProtectionDeviceRegistrations for deviceAppManagement
// returns a WindowsInformationProtectionDeviceRegistrationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationable, requestConfiguration *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation windows information protection device registrations that are not MDM enrolled.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to windowsInformationProtectionDeviceRegistrations for deviceAppManagement
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionDeviceRegistrationable, requestConfiguration *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder when successful
func (m *WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) WithUrl(rawUrl string)(*WindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    return NewWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
