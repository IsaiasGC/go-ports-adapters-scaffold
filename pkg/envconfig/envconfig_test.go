package envconfig

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testConfig struct {
	Var1             string `env:"VAR1,required"`
	Var2             string `env:"VAR2"`
	Var3             string `env:"VAR3,required"`
	Var4             string `env:"VAR4,required"`
	Var5             int    `env:"VAR5,required"`
	NamedStruct      myStruct
	AnonymousStruct1 struct {
		Var8             bool `env:"VAR8,required"`
		Var9             int  `env:"VAR9,required"`
		AnonymousStruct2 struct {
			Var10 string `env:"VAR10,required"`
		}
	}
}

type myStruct struct {
	Var6 bool `env:"VAR6,required"`
	Var7 int  `env:"VAR7,required"`
}

func TestGetMissingEnvVars(t *testing.T) {
	// Create a testConfig instance
	config := &testConfig{}

	// Set up environment variables
	setEnv("VAR1", "value1", config)
	setEnv("VAR2", "", config)
	setEnv("VAR3", "", config)
	setEnv("VAR6", "false", config)
	setEnv("VAR8", "", config)
	setEnv("VAR9", "17", config)

	// Call GetMissingEnvVars
	missingVars := GetMissingEnvVars(config)

	// Check that the missing variable is correctly identified
	assert.Equal(t, 6, len(missingVars))
	assert.Contains(t, missingVars, "VAR3")
	assert.Contains(t, missingVars, "VAR4")
	assert.Contains(t, missingVars, "VAR5")
	assert.Contains(t, missingVars, "VAR7")
	assert.Contains(t, missingVars, "VAR8")
	assert.Contains(t, missingVars, "VAR10")

	// Clean up
	os.Unsetenv("VAR1")
	os.Unsetenv("VAR2")
	os.Unsetenv("VAR3")
	os.Unsetenv("VAR6")
	os.Unsetenv("VAR8")
	os.Unsetenv("VAR9")
}

func setEnv(key, value string, config *testConfig) {
	os.Setenv(key, value)
	switch key {
	case "VAR1":
		config.Var1 = value
	case "VAR2":
		config.Var2 = value
	case "VAR3":
		config.Var3 = value
	case "VAR6":
		config.NamedStruct.Var6 = value == "true"
	case "VAR9":
		if intValue, err := strconv.Atoi(value); err == nil {
			config.AnonymousStruct1.Var9 = intValue
		}
	}
}
