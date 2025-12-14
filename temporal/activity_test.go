package temporal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/testsuite"
)

func SquareActivity(ctx context.Context, num int) (int, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("SquareActivity", "num", num)
	return num * num, nil
}

func Test_SquareActivityOne(t *testing.T) {
	testSuit := &testsuite.WorkflowTestSuite{}
	env := testSuit.NewTestActivityEnvironment()
	env.RegisterActivity(SquareActivity)

	future, err := env.ExecuteActivity(SquareActivity, 10)
	assert.NoError(t, err)

	var res int
	assert.NoError(t, future.Get(&res))

	assert.Equal(t, 100, res)
}

type User struct {
	Name string
	Age  int
}

type PersonGetter interface {
}
