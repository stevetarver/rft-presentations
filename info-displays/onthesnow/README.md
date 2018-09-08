Each numbered directory shows incremental evolution. Check the file header comments for what has changed.

Install "Gapplin" from the Mac App Store to make svg viewing easy.


To run a `go` file in a directory and produce the svg:

```
./run.sh 'dir number'
```

To refresh data:

```
./get_data.sh
```

## Serving images

Once you have a program to draw your SVG, and it has live data, you can set up a simple web server to provide realtime images to clients.

```go
package main
	
import (
	"log"
	"github.com/ajstarks/svgo"
	"net/http"
)
	
func main() {
	http.Handle("/circle", http.HandlerFunc(circle))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
	
func circle(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "image/svg+xml")
  s := svg.New(w)
  s.Start(500, 500)
  s.Circle(250, 250, 125, "fill:none;stroke:black")
  s.End()
}
```

_NOTE:_ You could set up a `time.Ticker` to regen images on some period.

