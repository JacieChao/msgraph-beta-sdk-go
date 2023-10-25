package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ZebraFotaConnectorHasActiveDeploymentsRequestBuilder provides operations to call the hasActiveDeployments method.
type ZebraFotaConnectorHasActiveDeploymentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ZebraFotaConnectorHasActiveDeploymentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ZebraFotaConnectorHasActiveDeploymentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewZebraFotaConnectorHasActiveDeploymentsRequestBuilderInternal instantiates a new HasActiveDeploymentsRequestBuilder and sets the default values.
func NewZebraFotaConnectorHasActiveDeploymentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorHasActiveDeploymentsRequestBuilder) {
    m := &ZebraFotaConnectorHasActiveDeploymentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/zebraFotaConnector/hasActiveDeployments", pathParameters),
    }
    return m
}
// NewZebraFotaConnectorHasActiveDeploymentsRequestBuilder instantiates a new HasActiveDeploymentsRequestBuilder and sets the default values.
func NewZebraFotaConnectorHasActiveDeploymentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ZebraFotaConnectorHasActiveDeploymentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewZebraFotaConnectorHasActiveDeploymentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action hasActiveDeployments
// Deprecated: This method is obsolete. Use PostAsHasActiveDeploymentsPostResponse instead.
func (m *ZebraFotaConnectorHasActiveDeploymentsRequestBuilder) Post(ctx context.Context, requestConfiguration *ZebraFotaConnectorHasActiveDeploymentsRequestBuilderPostRequestConfiguration)(ZebraFotaConnectorHasActiveDeploymentsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebraFotaConnectorHasActiveDeploymentsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebraFotaConnectorHasActiveDeploymentsResponseable), nil
}
// PostAsHasActiveDeploymentsPostResponse invoke action hasActiveDeployments
func (m *ZebraFotaConnectorHasActiveDeploymentsRequestBuilder) PostAsHasActiveDeploymentsPostResponse(ctx context.Context, requestConfiguration *ZebraFotaConnectorHasActiveDeploymentsRequestBuilderPostRequestConfiguration)(ZebraFotaConnectorHasActiveDeploymentsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateZebraFotaConnectorHasActiveDeploymentsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ZebraFotaConnectorHasActiveDeploymentsPostResponseable), nil
}
// ToPostRequestInformation invoke action hasActiveDeployments
func (m *ZebraFotaConnectorHasActiveDeploymentsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ZebraFotaConnectorHasActiveDeploymentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ZebraFotaConnectorHasActiveDeploymentsRequestBuilder) WithUrl(rawUrl string)(*ZebraFotaConnectorHasActiveDeploymentsRequestBuilder) {
    return NewZebraFotaConnectorHasActiveDeploymentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
