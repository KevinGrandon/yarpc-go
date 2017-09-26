// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package transport

import (
	"strconv"
	"strings"
)

const (
	// FeatureThriftApplicationError says that the client can handle thrift
	// application errors returned over the wire, not just in the thrift
	// envelope. This may involve a transport-specific header being returned
	// on the response to indicate if an error is an application error or not.
	FeatureThriftApplicationError = 1
)

var (
	// AllFeatures contains all features.
	AllFeatures = []Feature{
		FeatureThriftApplicationError,
	}

	_featureToString = map[Feature]string{
		FeatureThriftApplicationError: "1",
	}
	_stringToFeature = map[string]Feature{
		"1": FeatureThriftApplicationError,
	}
)

// Feature is a feature that the client can support.
//
// This makes it easier to add new features to YARPC in a backwards-compatible
// manner so that servers know how to construct responses.
type Feature int

// In returns true if f is in the given Features slice.
func (f Feature) In(features []Feature) bool {
	for _, feature := range features {
		if f == feature {
			return true
		}
	}
	return false
}

// String returns the the string representation of the Feature.
//
// This is just for printing, use the safer ToString method
// for transport implementations.
//
// Strings will be all lowercase and not contain commas.
func (f Feature) String() string {
	s, ok := _featureToString[f]
	if ok {
		return s
	}
	return strconv.Itoa(int(f))
}

// ToString returns the the string representation of the Feature, or false
// if the Feature is not known
//
// Strings will be all lowercase and not contain commas.
func (f Feature) ToString() (string, bool) {
	s, ok := _featureToString[f]
	if !ok {
		return "", false
	}
	return s, true
}

// FeatureFromString returns the Feature for the string, or false
// if the Feature is not known.
func FeatureFromString(s string) (Feature, bool) {
	f, ok := _stringToFeature[strings.ToLower(s)]
	if !ok {
		return Feature(0), false
	}
	return f, true
}

// FeaturesToString returns a comma-separated list of the string
// representations of the given Features.
//
// Unknown features will not be included.
func FeaturesToString(features []Feature) string {
	if len(features) == 0 {
		return ""
	}
	var featureStrings []string
	for _, feature := range features {
		s, ok := feature.ToString()
		if ok {
			featureStrings = append(featureStrings, s)
		}
	}
	return strings.Join(featureStrings, ",")
}

// FeaturesFromString returns a slice of Features for the given
// comma-separated string representation.
//
// Unknown features will not be included.
func FeaturesFromString(s string) []Feature {
	if s == "" {
		return nil
	}
	var features []Feature
	for _, e := range strings.Split(s, ",") {
		feature, ok := FeatureFromString(e)
		if ok {
			features = append(features, feature)
		}
	}
	return features
}
