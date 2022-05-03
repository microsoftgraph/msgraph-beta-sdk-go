package areglobalscriptsavailable

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AreGlobalScriptsAvailableRequestBuilder provides operations to call the areGlobalScriptsAvailable method.
type AreGlobalScriptsAvailableRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AreGlobalScriptsAvailableResponse union type wrapper for classes globalDeviceHealthScriptState
type AreGlobalScriptsAvailableResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Union type representation for type globalDeviceHealthScriptState
    globalDeviceHealthScriptState *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GlobalDeviceHealthScriptState
}
// NewAreGlobalScriptsAvailableResponse instantiates a new areGlobalScriptsAvailableResponse and sets the default values.
func NewAreGlobalScriptsAvailableResponse()(*AreGlobalScriptsAvailableResponse) {
    m := &AreGlobalScriptsAvailableResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateAreGlobalScriptsAvailableResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAreGlobalScriptsAvailableResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AreGlobalScriptsAvailableResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AreGlobalScriptsAvailableResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["globalDeviceHealthScriptState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseGlobalDeviceHealthScriptState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGlobalDeviceHealthScriptState(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GlobalDeviceHealthScriptState))
        }
        return nil
    }
    return res
}
// GetGlobalDeviceHealthScriptState gets the globalDeviceHealthScriptState property value. Union type representation for type globalDeviceHealthScriptState
func (m *AreGlobalScriptsAvailableResponse) GetGlobalDeviceHealthScriptState()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GlobalDeviceHealthScriptState) {
    if m == nil {
        return nil
    } else {
        return m.globalDeviceHealthScriptState
    }
}
// Serialize serializes information the current object
func (m *AreGlobalScriptsAvailableResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGlobalDeviceHealthScriptState() != nil {
        cast := (*m.GetGlobalDeviceHealthScriptState()).String()
        err := writer.WriteStringValue("globalDeviceHealthScriptState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AreGlobalScriptsAvailableResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGlobalDeviceHealthScriptState sets the globalDeviceHealthScriptState property value. Union type representation for type globalDeviceHealthScriptState
func (m *AreGlobalScriptsAvailableResponse) SetGlobalDeviceHealthScriptState(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GlobalDeviceHealthScriptState)() {
    if m != nil {
        m.globalDeviceHealthScriptState = value
    }
}
// AreGlobalScriptsAvailableResponseable 
type AreGlobalScriptsAvailableResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGlobalDeviceHealthScriptState()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GlobalDeviceHealthScriptState)
    SetGlobalDeviceHealthScriptState(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GlobalDeviceHealthScriptState)()
}
// NewAreGlobalScriptsAvailableRequestBuilderInternal instantiates a new AreGlobalScriptsAvailableRequestBuilder and sets the default values.
func NewAreGlobalScriptsAvailableRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AreGlobalScriptsAvailableRequestBuilder) {
    m := &AreGlobalScriptsAvailableRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts/microsoft.graph.areGlobalScriptsAvailable()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAreGlobalScriptsAvailableRequestBuilder instantiates a new AreGlobalScriptsAvailableRequestBuilder and sets the default values.
func NewAreGlobalScriptsAvailableRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AreGlobalScriptsAvailableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAreGlobalScriptsAvailableRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function areGlobalScriptsAvailable
func (m *AreGlobalScriptsAvailableRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function areGlobalScriptsAvailable
func (m *AreGlobalScriptsAvailableRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// GetWithResponseHandler invoke function areGlobalScriptsAvailable
func (m *AreGlobalScriptsAvailableRequestBuilder) GetWithResponseHandler(requestConfiguration *AreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration)(AreGlobalScriptsAvailableResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function areGlobalScriptsAvailable
func (m *AreGlobalScriptsAvailableRequestBuilder) GetWithResponseHandler(requestConfiguration *AreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(AreGlobalScriptsAvailableResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateAreGlobalScriptsAvailableResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(AreGlobalScriptsAvailableResponseable), nil
}
