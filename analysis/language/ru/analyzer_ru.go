//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

// +build libstemmer full
// +build icu full

package ru

import (
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"

	"github.com/blevesearch/bleve/analysis/token_filters/lower_case_filter"
	"github.com/blevesearch/bleve/analysis/tokenizers/unicode_word_boundary"
)

const AnalyzerName = "ru"

func AnalyzerConstructor(config map[string]interface{}, cache *registry.Cache) (*analysis.Analyzer, error) {
	unicodeTokenizer, err := cache.TokenizerNamed(unicode_word_boundary.Name)
	if err != nil {
		return nil, err
	}
	toLowerFilter, err := cache.TokenFilterNamed(lower_case_filter.Name)
	if err != nil {
		return nil, err
	}
	stopRuFilter, err := cache.TokenFilterNamed(StopName)
	if err != nil {
		return nil, err
	}
	stemmerRuFilter, err := cache.TokenFilterNamed(StemmerName)
	if err != nil {
		return nil, err
	}
	rv := analysis.Analyzer{
		Tokenizer: unicodeTokenizer,
		TokenFilters: []analysis.TokenFilter{
			toLowerFilter,
			stopRuFilter,
			stemmerRuFilter,
		},
	}
	return &rv, nil
}

func init() {
	registry.RegisterAnalyzer(AnalyzerName, AnalyzerConstructor)
}
