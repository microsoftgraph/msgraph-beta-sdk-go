package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder provides operations to manage the sensitivityLabels property of the microsoft.graph.informationProtection entity.
type InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderGetQueryParameters get sensitivityLabels from informationProtection
type InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderGetQueryParameters
}
// InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderInternal instantiates a new SensitivityLabelItemRequestBuilder and sets the default values.
func NewInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) {
    m := &InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/informationProtection/sensitivityLabels/{sensitivityLabel%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder instantiates a new SensitivityLabelItemRequestBuilder and sets the default values.
func NewInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sensitivityLabels for informationProtection
func (m *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get sensitivityLabels from informationProtection
func (m *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sensitivityLabels in informationProtection
func (m *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable, requestConfiguration *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property sensitivityLabels for informationProtection
func (m *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get sensitivityLabels from informationProtection
func (m *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) Get(ctx context.Context, requestConfiguration *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSensitivityLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable), nil
}
// Patch update the navigation property sensitivityLabels in informationProtection
func (m *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable, requestConfiguration *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSensitivityLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitivityLabelable), nil
}
// Sublabels provides operations to manage the sublabels property of the microsoft.graph.sensitivityLabel entity.
func (m *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) Sublabels()(*InformationProtectionSensitivityLabelsItemSublabelsRequestBuilder) {
    return NewInformationProtectionSensitivityLabelsItemSublabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SublabelsById provides operations to manage the sublabels property of the microsoft.graph.sensitivityLabel entity.
func (m *InformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) SublabelsById(id string)(*InformationProtectionSensitivityLabelsItemSublabelsSensitivityLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitivityLabel%2Did1"] = id
    }
    return NewInformationProtectionSensitivityLabelsItemSublabelsSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
