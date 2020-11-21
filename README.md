# Issue 2: Variadic

PostMessage takes a variadic parameter of `slack.MsgOption` This seems to prevent
us from testing it's usage inside wrapped worker code.

```bash
go generate
go test
```

Yields:

```text
--- FAIL: TestPostMessage (0.00s)
    testing_t_support.go:41: 
                ${HOME}/go/pkg/mod/github.com/petergtz/pegomock@v2.8.0+incompatible/testing_t_support.go:40 +0x5e
        github.com/petergtz/pegomock.(*GenericMock).Verify(0xc00000e4a0, 0x0, 0x12cae40, 0xc00007ef00, 0x127a246, 0xb, 0xc00000e500, 0x2, 0x2, 0xc000059ed0, ...)
                ${HOME}/go/pkg/mod/github.com/petergtz/pegomock@v2.8.0+incompatible/dsl.go:153 +0x5be
        awesomeProject/mocks.(*VerifierMockSlackClient).PostMessage(0xc000059f48, 0x1279af1, 0x9, 0xc000059f40, 0x1, 0x1, 0x0)
                ${HOME}/go/src/awesomeProject/mocks/mock_slack_client.go:96 +0x267
        awesomeProject_test.TestPostMessage(0xc000001b00)
                ${HOME}/go/src/awesomeProject/slack_test.go:17 +0x1b7
        testing.tRunner(0xc000001b00, 0x128d170)
                /usr/local/go/src/testing/testing.go:1123 +0xef
        created by testing.(*T).Run
                /usr/local/go/src/testing/testing.go:1168 +0x2b3
        
        Mock invocation count for PostMessage("something", (slack.MsgOption)(0x11f3f40)) does not match expectation.
        
                Expected: 1; but got: 0
        
                But other interactions with this mock were:
                PostMessage("something", (slack.MsgOption)(0x11f3f40))
FAIL
exit status 1
FAIL    awesomeProject  0.284s
```
