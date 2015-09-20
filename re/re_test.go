package re

import "fmt"
import "testing"
import "../tameshigiri"

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

    var assertion = tameshigiri.NewAssertion(t)
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

    assertion.Equals(
        expectedListLength,
        result.CountIndices(),
        "Item-list count",
    )

    if expectedListLength > 0 {
        for index = range expectedList {
            actual   = result.Index(index)
            expected = expectedList[index]

            assertion.Equals(
                expected,
                *actual,
                fmt.Sprintf("List#%d", index),
            )
        }
    }

    assertion.Equals(
        expectedDictLength,
        result.CountKeys(),
        "Dictionary count",
    )

    if expectedDictLength > 0 {
        for key, expected = range expectedDict {
            actual = result.Key(key)

            assertion.Equals(
                expected,
                *actual,
                fmt.Sprintf("Dictionary[%s]", key),
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
