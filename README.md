# YMuse API Client for Go

This is an experimental YMuse API client for Go. Hopefully, there will
be some documentation some day. Until then, you can enjoy it by:

```go
package main

import (
	"ymuse.com/gomuse"
)

func main(){
	cfg := gomuse.ClientConfig{
		ClientId:     "YOUR CLIENT ID",
		ClientSecret: "YOUR CLIENT SECRET",
		Debug:        false,
	}
	client := gomuse.NewClient(cfg)
	content, _ := client.GetFeaturedContent()
	//Do something with the content etc...
}
```

## License

This client uses the [napping](https://github.com/jmcvetta/napping/)
library which is licensed under [GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html). Therefore this library is also
licensed with GPLv3.
