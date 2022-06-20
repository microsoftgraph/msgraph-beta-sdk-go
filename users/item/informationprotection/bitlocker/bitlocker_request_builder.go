package bitlocker

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i6500822905fab5c3cf933e09dbf8ac1bb5df488cf21800ea5df6d0a2fa809263 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/bitlocker/recoverykeys"
    i1bab938e73efbf18be0fa79b7277a62cc62b99cf543603847896ed64d30417de "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/informationprotection/bitlocker/recoverykeys/item"
)

// BitlockerRequestBuilder provides operations to manage the bitlocker property of the microsoft.graph.informationProtection entity.
type BitlockerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BitlockerRequestBuilderGetQueryParameters get bitlocker from users
type BitlockerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BitlockerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BitlockerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BitlockerRequestBuilderGetQueryParameters
}
// BitlockerRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BitlockerRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBitlockerRequestBuilderInternal instantiates a new BitlockerRequestBuilder and sets the default values.
func NewBitlockerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRequestBuilder) {
    m := &BitlockerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/informationProtection/bitlocker{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBitlockerRequestBuilder instantiates a new BitlockerRequestBuilder and sets the default values.
func NewBitlockerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBitlockerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get bitlocker from users
func (m *BitlockerRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get bitlocker from users
func (m *BitlockerRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *BitlockerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property bitlocker in users
func (m *BitlockerRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Bitlockerable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property bitlocker in users
func (m *BitlockerRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Bitlockerable, requestConfiguration *BitlockerRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get bitlocker from users
func (m *BitlockerRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Bitlockerable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get bitlocker from users
func (m *BitlockerRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *BitlockerRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Bitlockerable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBitlockerFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Bitlockerable), nil
}
// Patch update the navigation property bitlocker in users
func (m *BitlockerRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Bitlockerable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property bitlocker in users
func (m *BitlockerRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Bitlockerable, requestConfiguration *BitlockerRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RecoveryKeys the recoveryKeys property
func (m *BitlockerRequestBuilder) RecoveryKeys()(*i6500822905fab5c3cf933e09dbf8ac1bb5df488cf21800ea5df6d0a2fa809263.RecoveryKeysRequestBuilder) {
    return i6500822905fab5c3cf933e09dbf8ac1bb5df488cf21800ea5df6d0a2fa809263.NewRecoveryKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecoveryKeysById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.informationProtection.bitlocker.recoveryKeys.item collection
func (m *BitlockerRequestBuilder) RecoveryKeysById(id string)(*i1bab938e73efbf18be0fa79b7277a62cc62b99cf543603847896ed64d30417de.BitlockerRecoveryKeyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bitlockerRecoveryKey%2Did"] = id
    }
    return i1bab938e73efbf18be0fa79b7277a62cc62b99cf543603847896ed64d30417de.NewBitlockerRecoveryKeyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
