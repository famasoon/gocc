package complexity

import "testing"

func TestComplexity(t *testing.T) {
  testcases := []struct{
    name string
    code string
    complexity int
  }{
    //TODO
  }

  for _, testcase := range testcases {
    t.Run(testcase.name, func(t *testing.T) {
      a := GetAST(t, testcase.code)
      c := Count(a)

      if c != testcase.complexity {
        t.Errorf("got=%d, want=%d", c, testcase.complexity)
      }
    })
  }
}