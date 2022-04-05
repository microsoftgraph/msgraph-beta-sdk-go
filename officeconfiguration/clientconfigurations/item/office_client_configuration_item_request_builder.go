package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1068f00153d896756e5c75882c5dbee1a31ff4ad6c6cd1299946c4acf1e1932d "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/policypayload"
    i6ee98b33165e3fc15ddb585c37c29eb4d3887d300742016bd30e69c1dff7f0b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/userpreferencepayload"
    ib708431220be37864eb91ad6ad04c27754da18439883881ba9f54b62ae1d2206 "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/assign"
    ie1103875ba3eb7fa000dd34e6685f4073add80b4c464e341669e479f36621a4c "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/assignments"
    ie58d03e4a8a1ff62a18e2ec91d18b2410cce187a18440bbde6c2754055d5ee34 "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item/assignments/item"
)

// OfficeClientConfigurationItemRequestBuilder provides operations to manage the clientConfigurations property of the microsoft.graph.officeConfiguration entity.
type OfficeClientConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OfficeClientConfigurationItemRequestBuilderDeleteOptions options for Delete
type OfficeClientConfigurationItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OfficeClientConfigurationItemRequestBuilderGetOptions options for Get
type OfficeClientConfigurationItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *OfficeClientConfigurationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// OfficeClientConfigurationItemRequestBuilderGetQueryParameters list of office Client configuration.
type OfficeClientConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OfficeClientConfigurationItemRequestBuilderPatchOptions options for Patch
type OfficeClientConfigurationItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeClientConfigurationable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Assign the assign property
func (m *OfficeClientConfigurationItemRequestBuilder) Assign()(*ib708431220be37864eb91ad6ad04c27754da18439883881ba9f54b62ae1d2206.AssignRequestBuilder) {
    return ib708431220be37864eb91ad6ad04c27754da18439883881ba9f54b62ae1d2206.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *OfficeClientConfigurationItemRequestBuilder) Assignments()(*ie1103875ba3eb7fa000dd34e6685f4073add80b4c464e341669e479f36621a4c.AssignmentsRequestBuilder) {
    return ie1103875ba3eb7fa000dd34e6685f4073add80b4c464e341669e479f36621a4c.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.officeConfiguration.clientConfigurations.item.assignments.item collection
func (m *OfficeClientConfigurationItemRequestBuilder) AssignmentsById(id string)(*ie58d03e4a8a1ff62a18e2ec91d18b2410cce187a18440bbde6c2754055d5ee34.OfficeClientConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["officeClientConfigurationAssignment_id"] = id
    }
    return ie58d03e4a8a1ff62a18e2ec91d18b2410cce187a18440bbde6c2754055d5ee34.NewOfficeClientConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOfficeClientConfigurationItemRequestBuilderInternal instantiates a new OfficeClientConfigurationItemRequestBuilder and sets the default values.
func NewOfficeClientConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OfficeClientConfigurationItemRequestBuilder) {
    m := &OfficeClientConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/officeConfiguration/clientConfigurations/{officeClientConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOfficeClientConfigurationItemRequestBuilder instantiates a new OfficeClientConfigurationItemRequestBuilder and sets the default values.
func NewOfficeClientConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OfficeClientConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOfficeClientConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property clientConfigurations for officeConfiguration
func (m *OfficeClientConfigurationItemRequestBuilder) CreateDeleteRequestInformation(options *OfficeClientConfigurationItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation list of office Client configuration.
func (m *OfficeClientConfigurationItemRequestBuilder) CreateGetRequestInformation(options *OfficeClientConfigurationItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property clientConfigurations in officeConfiguration
func (m *OfficeClientConfigurationItemRequestBuilder) CreatePatchRequestInformation(options *OfficeClientConfigurationItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property clientConfigurations for officeConfiguration
func (m *OfficeClientConfigurationItemRequestBuilder) Delete(options *OfficeClientConfigurationItemRequestBuilderDeleteOptions)(error) {
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
// Get list of office Client configuration.
func (m *OfficeClientConfigurationItemRequestBuilder) Get(options *OfficeClientConfigurationItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeClientConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOfficeClientConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeClientConfigurationable), nil
}
// Patch update the navigation property clientConfigurations in officeConfiguration
func (m *OfficeClientConfigurationItemRequestBuilder) Patch(options *OfficeClientConfigurationItemRequestBuilderPatchOptions)(error) {
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
// PolicyPayload the policyPayload property
func (m *OfficeClientConfigurationItemRequestBuilder) PolicyPayload()(*i1068f00153d896756e5c75882c5dbee1a31ff4ad6c6cd1299946c4acf1e1932d.PolicyPayloadRequestBuilder) {
    return i1068f00153d896756e5c75882c5dbee1a31ff4ad6c6cd1299946c4acf1e1932d.NewPolicyPayloadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserPreferencePayload the userPreferencePayload property
func (m *OfficeClientConfigurationItemRequestBuilder) UserPreferencePayload()(*i6ee98b33165e3fc15ddb585c37c29eb4d3887d300742016bd30e69c1dff7f0b8.UserPreferencePayloadRequestBuilder) {
    return i6ee98b33165e3fc15ddb585c37c29eb4d3887d300742016bd30e69c1dff7f0b8.NewUserPreferencePayloadRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
