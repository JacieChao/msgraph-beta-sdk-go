package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder provides operations to call the removeHold method.
type EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderInternal instantiates a new MicrosoftGraphEdiscoveryRemoveHoldRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) {
    m := &EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/noncustodialDataSources/microsoft.graph.ediscovery.removeHold", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder instantiates a new MicrosoftGraphEdiscoveryRemoveHoldRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action removeHold
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) Post(ctx context.Context, body EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRemoveHoldPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action removeHold
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, body EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRemoveHoldPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) {
    return NewEdiscoveryCasesItemNoncustodialDataSourcesMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
