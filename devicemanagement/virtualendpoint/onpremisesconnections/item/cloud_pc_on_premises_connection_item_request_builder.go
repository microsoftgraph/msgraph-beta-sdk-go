package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i747483174c0e61425ce1a618854ba8ca6d202a33166c1334d75d025ab520f858 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/onpremisesconnections/item/runhealthchecks"
    if44fdd10cc1f6f4d9584fbcc1424ba5392d6f3eda6f11879e38028e5a1ea53bf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/onpremisesconnections/item/updateaddomainpassword"
)

// CloudPcOnPremisesConnectionItemRequestBuilder provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
type CloudPcOnPremisesConnectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CloudPcOnPremisesConnectionItemRequestBuilderDeleteOptions options for Delete
type CloudPcOnPremisesConnectionItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CloudPcOnPremisesConnectionItemRequestBuilderGetOptions options for Get
type CloudPcOnPremisesConnectionItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters a defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
type CloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CloudPcOnPremisesConnectionItemRequestBuilderPatchOptions options for Patch
type CloudPcOnPremisesConnectionItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewCloudPcOnPremisesConnectionItemRequestBuilderInternal instantiates a new CloudPcOnPremisesConnectionItemRequestBuilder and sets the default values.
func NewCloudPcOnPremisesConnectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPcOnPremisesConnectionItemRequestBuilder) {
    m := &CloudPcOnPremisesConnectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/onPremisesConnections/{cloudPcOnPremisesConnection%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCloudPcOnPremisesConnectionItemRequestBuilder instantiates a new CloudPcOnPremisesConnectionItemRequestBuilder and sets the default values.
func NewCloudPcOnPremisesConnectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPcOnPremisesConnectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPcOnPremisesConnectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property onPremisesConnections for deviceManagement
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) CreateDeleteRequestInformation(options *CloudPcOnPremisesConnectionItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) CreateGetRequestInformation(options *CloudPcOnPremisesConnectionItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property onPremisesConnections in deviceManagement
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) CreatePatchRequestInformation(options *CloudPcOnPremisesConnectionItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property onPremisesConnections for deviceManagement
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) Delete(options *CloudPcOnPremisesConnectionItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) Get(options *CloudPcOnPremisesConnectionItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcOnPremisesConnectionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable), nil
}
// Patch update the navigation property onPremisesConnections in deviceManagement
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) Patch(options *CloudPcOnPremisesConnectionItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RunHealthChecks the runHealthChecks property
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) RunHealthChecks()(*i747483174c0e61425ce1a618854ba8ca6d202a33166c1334d75d025ab520f858.RunHealthChecksRequestBuilder) {
    return i747483174c0e61425ce1a618854ba8ca6d202a33166c1334d75d025ab520f858.NewRunHealthChecksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateAdDomainPassword the updateAdDomainPassword property
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) UpdateAdDomainPassword()(*if44fdd10cc1f6f4d9584fbcc1424ba5392d6f3eda6f11879e38028e5a1ea53bf.UpdateAdDomainPasswordRequestBuilder) {
    return if44fdd10cc1f6f4d9584fbcc1424ba5392d6f3eda6f11879e38028e5a1ea53bf.NewUpdateAdDomainPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
