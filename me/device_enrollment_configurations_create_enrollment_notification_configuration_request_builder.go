package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder provides operations to call the createEnrollmentNotificationConfiguration method.
type DeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderInternal instantiates a new CreateEnrollmentNotificationConfigurationRequestBuilder and sets the default values.
func NewDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder) {
    m := &DeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/deviceEnrollmentConfigurations/microsoft.graph.createEnrollmentNotificationConfiguration";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder instantiates a new CreateEnrollmentNotificationConfigurationRequestBuilder and sets the default values.
func NewDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createEnrollmentNotificationConfiguration
func (m *DeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder) Post(ctx context.Context, body DeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBodyable, requestConfiguration *DeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action createEnrollmentNotificationConfiguration
func (m *DeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBodyable, requestConfiguration *DeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}