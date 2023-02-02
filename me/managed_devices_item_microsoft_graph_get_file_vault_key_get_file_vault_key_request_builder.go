package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilder provides operations to call the getFileVaultKey method.
type ManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilderInternal instantiates a new GetFileVaultKeyRequestBuilder and sets the default values.
func NewManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilder) {
    m := &ManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice%2Did}/microsoft.graph.getFileVaultKey()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilder instantiates a new GetFileVaultKeyRequestBuilder and sets the default values.
func NewManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getFileVaultKey
func (m *ManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilderGetRequestConfiguration)(ManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyResponseable), nil
}
// ToGetRequestInformation invoke function getFileVaultKey
func (m *ManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesItemMicrosoftGraphGetFileVaultKeyGetFileVaultKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
