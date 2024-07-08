package testMethod

import "testing"

func SetupEnv(t *testing.T, envs map[string]string) {
	for k, v := range envs {
		t.Setenv(k, v)
	}
}
