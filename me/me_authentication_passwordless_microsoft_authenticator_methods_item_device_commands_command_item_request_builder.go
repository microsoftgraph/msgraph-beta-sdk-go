package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder provides operations to manage the commands property of the microsoft.graph.device entity.
type MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderGetQueryParameters set of commands sent to this device.
type MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderGetQueryParameters
}
// MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderInternal instantiates a new CommandItemRequestBuilder and sets the default values.
func NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder) {
    m := &MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/commands/{command%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder instantiates a new CommandItemRequestBuilder and sets the default values.
func NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property commands for me
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation set of commands sent to this device.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property commands in me
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Commandable, requestConfiguration *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property commands for me
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get set of commands sent to this device.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Commandable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCommandFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Commandable), nil
}
// Patch update the navigation property commands in me
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Commandable, requestConfiguration *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Commandable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCommandFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Commandable), nil
}
// Responsepayload provides operations to manage the responsepayload property of the microsoft.graph.command entity.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsCommandItemRequestBuilder) Responsepayload()(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsItemResponsepayloadRequestBuilder) {
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCommandsItemResponsepayloadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
