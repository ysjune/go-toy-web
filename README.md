# go-toy-web
clone tucker go programming

## Testing

### [goconvey](https://github.com/smartystreets/goconvey)
테스트 자동화

```bash
go get github.com/smartystreets/goconvey

goconvey
```

### [testify](https://github.com/stretchr/testify)

테스팅 모듈 (assert 패키지만 사용)

```bash
go get github.com/stretchr/testify/assert
```

#### usage
```go
import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

    assert.Equal(t, 123, 123, "they should be equal")

    //or

    assert := assert.New(t)
	assert.Equal(123, 123, "they should be equal")
}
```

## rest api

### [gorilla mux](https://github.com/gorilla/mux) (archived)

리퀘스트 라우터 핸들러

```
go get -u github.com/gorilla/mux
```