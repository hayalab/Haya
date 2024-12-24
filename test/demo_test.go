package main

import (
	"testing"

	"github.com/melodydev777/Melody/models"
)

func TestData(t *testing.T) {
	initTester()
	result, err := models.GetUserInfoById(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestDemo(t *testing.T) {
	t.Log("test")
}
