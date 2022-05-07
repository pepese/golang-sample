package app

import (
	"testing"

	. "github.com/go-playground/assert"
	"github.com/golang/mock/gomock"
	"github.com/pepese/golang-sample/unit-testing/app/mock_app"
)

func TestSample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_app.NewMockSampleInterfacer(ctrl)

	m.
		EXPECT().
		SampleFunc(1).
		Return(101)

	Equal(t, Execute(m), true) // assert
}
