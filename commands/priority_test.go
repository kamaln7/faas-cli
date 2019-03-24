package commands

import "testing"

func Test_getTemplateStoreURL(t *testing.T) {
	tests := []struct {
		title       string
		envURL      string
		defaultURL  string
		argURL      string
		expectedURL string
	}{
		{
			title:       "Environmental variable is set and argument equals defaultURL which should be priority",
			envURL:      "https://github.com/custom/url",
			defaultURL:  DefaultTemplatesStore,
			argURL:      DefaultTemplatesStore,
			expectedURL: "https://github.com/custom/url",
		},
		{
			title:       "Environmental variable is unset and argument is unset which falls back to default store",
			envURL:      "",
			defaultURL:  DefaultTemplatesStore,
			argURL:      DefaultTemplatesStore,
			expectedURL: DefaultTemplatesStore,
		},
		{
			title:       "Environmental variable is unset but argument is set which should set URL as argument",
			envURL:      "",
			defaultURL:  DefaultTemplatesStore,
			argURL:      "https://github.com/openfaas/store/official",
			expectedURL: "https://github.com/openfaas/store/official",
		},
		{
			title:       "Environmental variable is set and argument is set which should set URL as argument",
			envURL:      "https://github.com/custom/url",
			defaultURL:  DefaultTemplatesStore,
			argURL:      "https://github.com/openfaas/store/official",
			expectedURL: "https://github.com/openfaas/store/official",
		},
	}
	// defaultURL is always present that is why we don't test that case
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			storeURL := getTemplateStoreURL(test.argURL, test.envURL, test.defaultURL)
			if storeURL != test.expectedURL {
				t.Errorf("expected store URL: `%s` got: `%s`", test.expectedURL, storeURL)
			}
		})
	}
}

func Test_getImagePrefix(t *testing.T) {
	tests := []struct {
		title          string
		argPrefix      string
		envPrefix      string
		yamlPrefix     string
		expectedPrefix string
	}{
		{
			title:          "No prefixes are set which should return a blank prefix",
			expectedPrefix: "",
		},
		{
			title:          "Environment prefix should take priority over YAML prefix",
			envPrefix:      "envuser",
			yamlPrefix:     "yamluser",
			expectedPrefix: "envuser",
		},
		{
			title:          "Argument prefix should take priority over both environment and YAML prefixes",
			envPrefix:      "envuser",
			argPrefix:      "arguser",
			yamlPrefix:     "yamluser",
			expectedPrefix: "arguser",
		},
		{
			title:          "YAML prefix should be used if no others are set",
			yamlPrefix:     "yamluser",
			expectedPrefix: "yamluser",
		},
	}
	// defaultURL is always present that is why we don't test that case
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			prefix := getImagePrefix(test.argPrefix, test.envPrefix, test.yamlPrefix)
			if prefix != test.expectedPrefix {
				t.Errorf("expected image prefix: `%s` got: `%s`", test.expectedPrefix, prefix)
			}
		})
	}
}
