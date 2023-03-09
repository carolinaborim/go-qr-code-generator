#### This project is a collaboration between Abdallah and Carolina
* abdallah.hodieb@hashicorp.com
* carolina.borim@hashicorp.com


## TODO

### Phase 1

CLI application that takes a url ? and generates a Qr code

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
