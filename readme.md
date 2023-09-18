#### This project is a collaboration between Abdallah and Carolina
* abdallah.hodieb@hashicorp.com
* carolina.borim@hashicorp.com


## TODO

### Phase 1

CLI application that takes an url ? and generates a Qr code

  ```
  qr -o qr.png "https:google.com"
  ```

* create CLI
* Parse flags
* Generate a qr (qr library)
* save to file

### Phase 2

Evolve this to a HTTP server which exposes an endpoint that takes the same url and returns an image of the qr code
* HTTP server 
* URL parsing
* Maybe add a router ?
* Extract parameter
* Return an image 

```
qr -u "https://google.com" -o image.png

qr -server 
```


#### End goal
* CLI has a new flag that runs in server mode
* Server mode exposes an endpoint that takes a url and brings back a QR image
* We want the API responses to be similar to https://placehold.co/ i.e if you put the url in an <img> tag the browser should be able to render the qr code just fine


#### Breakdown 

* ✅ Creating an HTTP server  
* ✅ Extract the url from the endpoint.
* ✅ Adding additional "server-mode" flag to the CLI.
* ✅ Reuse the existing qr generation logic to return the image.
* ✅ Proper test http (instead of using phone).
* ✅ Add Tests & re-structure the code to make it more testable.
* ✅ Refactor to split things in different files with a logical flow.
* Embed html for a  UI for adding the URL.

### Phase 3
* Define the service GRPC
