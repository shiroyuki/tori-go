package re

import "fmt"
import "testing"
import tori "../"

func localTestReSearch(
    t            *testing.T,
    pattern      string,
    sample       string,
    expectedList []string,
    expectedDict map[string]string,
) {
    var index    int
    var key      string
    var actual   *string
    var expected string

    var assertion = tori.NewAssertion(t)
    var compiled  = Compile(pattern)
    var result    = compiled.Search(sample)

    var expectedListLength = len(expectedList)
    var expectedDictLength = len(expectedDict)
    var expectedTotalCount = expectedListLength + expectedDictLength

    if expectedTotalCount > 0 {
        assertion.IsTrue(result.HasAny(), "This should be found.")
    }

    assertion.Equals(
        expectedTotalCount,
        result.Count(),
        "Total count",
    )

    assertion.IsTrue(
        result.CountIndices() == expectedListLength,
        fmt.Sprintf("Expected item-list count: %d, given %d", expectedListLength, result.CountIndices()),
    )

    if expectedListLength > 0 {
        for index = range expectedList {
            actual   = result.Index(index)
            expected = expectedList[index]

            assertion.IsTrue(
                *actual == expected,
                fmt.Sprintf("Expected ItemList#%d is '%s', not '%s'.", index, expected, *actual),
            )
        }
    }

    assertion.IsTrue(
        result.CountKeys() == expectedDictLength,
        fmt.Sprintf("Expected dictionary count: %d, given %d", expectedDictLength, result.CountKeys()),
    )

    if expectedDictLength > 0 {
        for key, expected = range expectedDict {
            actual = result.Key(key)

            assertion.IsTrue(
                *actual == expected,
                fmt.Sprintf("Expected ItemList#%s is '%s', not '%s'.", key, expected, *actual),
            )
        }
    }
}

func TestReSearchBasic(t *testing.T) {
    localTestReSearch(
        t,
        "shiroyuki",
        "user-01-shiroyuki-as-admin",
        []string{ "shiroyuki" },
        make(map[string]string, 0),
    )
}

func TestReSearchAdvanced(t *testing.T) {
    var testPattern = "(/api)?/users/(?P<key>[^/]+)"

    localTestReSearch(
        t,
        testPattern,
        "/users",
        make([]string, 0),
        make(map[string]string, 0),
    )

    localTestReSearch(
        t,
        testPattern,
        "/users/shiroyuki",
        []string{
            "/users/shiroyuki",
            "",
        },
        map[string]string{
            "key": "shiroyuki",
        },
    )

    localTestReSearch(
        t,
        testPattern,
        "/api/users/shiroyuki",
        []string{
            "/api/users/shiroyuki",
            "/api",
        },
        map[string]string{
            "key": "shiroyuki",
        },
    )
}
