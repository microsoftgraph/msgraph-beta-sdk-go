package getauditactivitytypeswithcategory

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetAuditActivityTypesWithCategoryRequestBuilder provides operations to call the getAuditActivityTypes method.
type GetAuditActivityTypesWithCategoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetAuditActivityTypesWithCategoryRequestBuilderGetOptions options for Get
type GetAuditActivityTypesWithCategoryRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetAuditActivityTypesWithCategoryRequestBuilderInternal instantiates a new GetAuditActivityTypesWithCategoryRequestBuilder and sets the default values.
func NewGetAuditActivityTypesWithCategoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, category *string)(*GetAuditActivityTypesWithCategoryRequestBuilder) {
    m := &GetAuditActivityTypesWithCategoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/auditEvents/microsoft.graph.getAuditActivityTypes(category='{category}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if category != nil {
        urlTplParams[""] = *category
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetAuditActivityTypesWithCategoryRequestBuilder instantiates a new GetAuditActivityTypesWithCategoryRequestBuilder and sets the default values.
func NewGetAuditActivityTypesWithCategoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAuditActivityTypesWithCategoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAuditActivityTypesWithCategoryRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getAuditActivityTypes
func (m *GetAuditActivityTypesWithCategoryRequestBuilder) CreateGetRequestInformation(options *GetAuditActivityTypesWithCategoryRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
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
// Get invoke function getAuditActivityTypes
func (m *GetAuditActivityTypesWithCategoryRequestBuilder) Get(options *GetAuditActivityTypesWithCategoryRequestBuilderGetOptions)(GetAuditActivityTypesWithCategoryResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetAuditActivityTypesWithCategoryResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetAuditActivityTypesWithCategoryResponseable), nil
}
