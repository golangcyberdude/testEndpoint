# testEndpoint
This is a demonstration of how to conditionally test live endpoint services within go test

    =->export TESTENDPOINT=true

    =->go test --run TestExample --debug
    Found environment variable TESTENDPOINT, setting *testEndpoint = true
    running with debug mode [true]
    running with test-endpoint mode [true]
    in TestExample
    PASS
    ok      testEndpoint    0.007s


