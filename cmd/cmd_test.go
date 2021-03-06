// Copyright 2017 the lostromos Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestSetupLogger(t *testing.T) {
	viper.Set("logging.debug", false)
	viper.Set("logging.pretty", false)
	setupLogging()
	assert.Equal(t, false, logger.Desugar().Core().Enabled(zap.DebugLevel), "debug logging should be false")

	viper.Set("logging.debug", false)
	viper.Set("logging.pretty", true)
	setupLogging()
	assert.Equal(t, true, logger.Desugar().Core().Enabled(zap.InfoLevel), "debug logging should be false")
}
