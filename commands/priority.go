// Copyright (c) Alex Ellis 2017. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package commands

import (
	"fmt"
	"strings"
)

const (
	openFaaSURLEnvironment      = "OPENFAAS_URL"
	templateURLEnvironment      = "OPENFAAS_TEMPLATE_URL"
	templateStoreURLEnvironment = "OPENFAAS_TEMPLATE_STORE_URL"
	imagePrefixEnvironment      = "OPENFAAS_IMAGE_PREFIX"
)

func getGatewayURL(argumentURL, defaultURL, yamlURL, environmentURL string) string {
	var gatewayURL string

	if len(argumentURL) > 0 && argumentURL != defaultURL {
		gatewayURL = argumentURL
	} else if len(yamlURL) > 0 && yamlURL != defaultURL {
		gatewayURL = yamlURL
	} else if len(environmentURL) > 0 {
		gatewayURL = environmentURL
	} else {
		gatewayURL = defaultURL
	}

	gatewayURL = strings.ToLower(strings.TrimRight(gatewayURL, "/"))
	if !strings.HasPrefix(gatewayURL, "http") {
		gatewayURL = fmt.Sprintf("http://%s", gatewayURL)
	}

	return gatewayURL
}

func getTemplateURL(argumentURL, environmentURL, defaultURL string) string {
	var templateURL string

	if len(argumentURL) > 0 && argumentURL != defaultURL {
		templateURL = argumentURL
	} else if len(environmentURL) > 0 {
		templateURL = environmentURL
	} else {
		templateURL = defaultURL
	}

	return templateURL
}

func getTemplateStoreURL(argumentURL, environmentURL, defaultURL string) string {
	if argumentURL != defaultURL {
		return argumentURL
	} else if len(environmentURL) > 0 {
		return environmentURL
	} else {
		return defaultURL
	}
}

func getImagePrefix(argumentPrefix, environmentPrefix, yamlPrefix string) string {
	// default prefix is blank/no prefix
	var prefix string

	if len(yamlPrefix) > 0 {
		prefix = yamlPrefix
	}

	if len(environmentPrefix) > 0 {
		prefix = environmentPrefix
	}

	if len(argumentPrefix) > 0 {
		prefix = argumentPrefix
	}

	return prefix
}
