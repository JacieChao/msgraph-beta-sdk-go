package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder provides operations to call the updateGlobalScript method.
type DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderInternal instantiates a new UpdateGlobalScriptRequestBuilder and sets the default values.
func NewDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) {
    m := &DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript%2Did}/updateGlobalScript", pathParameters),
    }
    return m
}
// NewDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder instantiates a new UpdateGlobalScriptRequestBuilder and sets the default values.
func NewDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the Proprietary Device Health Script
// Deprecated: This method is obsolete. Use PostAsUpdateGlobalScriptPostResponse instead.
func (m *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) Post(ctx context.Context, body DeviceHealthScriptsItemUpdateGlobalScriptPostRequestBodyable, requestConfiguration *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderPostRequestConfiguration)(DeviceHealthScriptsItemUpdateGlobalScriptResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceHealthScriptsItemUpdateGlobalScriptResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceHealthScriptsItemUpdateGlobalScriptResponseable), nil
}
// PostAsUpdateGlobalScriptPostResponse update the Proprietary Device Health Script
func (m *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) PostAsUpdateGlobalScriptPostResponse(ctx context.Context, body DeviceHealthScriptsItemUpdateGlobalScriptPostRequestBodyable, requestConfiguration *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderPostRequestConfiguration)(DeviceHealthScriptsItemUpdateGlobalScriptPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceHealthScriptsItemUpdateGlobalScriptPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceHealthScriptsItemUpdateGlobalScriptPostResponseable), nil
}
// ToPostRequestInformation update the Proprietary Device Health Script
func (m *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeviceHealthScriptsItemUpdateGlobalScriptPostRequestBodyable, requestConfiguration *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) WithUrl(rawUrl string)(*DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) {
    return NewDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
